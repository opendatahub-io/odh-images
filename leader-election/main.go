package main

import (
	"context"
	"flag"
	"sync/atomic"
	"time"

	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/apex/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

// album represents data about a record album.
type Leader struct {
	Name string `json:"name"`
}

var leaderElected = Leader{
	Name: "",
}

func main() {
	var kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	var nodeID = flag.String("node-id", "", "node id used for leader election")
	var namespace = flag.String("namespace", "default", "namespace used for leader election")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clientset, err := newClientset(*kubeconfig)
	if err != nil {
		log.WithError(err).Fatal("failed to connect to cluster")
	}

	lock := &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      "leader-election-lock",
			Namespace: *namespace,
		},
		Client: clientset.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: *nodeID,
		},
	}

	// Routing to check the leader elected
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.Default()
	router.GET("/", getLeaderElected)
	go router.Run("localhost:4040")

	var leading int32
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:            lock,
		ReleaseOnCancel: true,
		LeaseDuration:   15 * time.Second,
		RenewDeadline:   10 * time.Second,
		RetryPeriod:     2 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				atomic.StoreInt32(&leading, 1)
				log.WithField("node", *nodeID).Info("started leading")
			},
			OnStoppedLeading: func() {
				atomic.StoreInt32(&leading, 0)
				log.WithField("id", *nodeID).Info("stopped leading")
			},
			OnNewLeader: func(identity string) {
				leaderElected.Name = identity
				if identity == *nodeID {
					return
				}
				log.
					WithField("leader", identity).
					Info("new leader elected")
			},
		},
	})
}

func newClientset(filename string) (*kubernetes.Clientset, error) {
	config, err := getConfig(filename)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func getConfig(cfg string) (*rest.Config, error) {
	if cfg == "" {
		return rest.InClusterConfig()
	}
	return clientcmd.BuildConfigFromFlags("", cfg)
}

func getLeaderElected(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, leaderElected)
}
