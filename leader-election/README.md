# Kubernetes Leader Election Informers

An example of a basic controller implementation based on informers based on the [go's leader election package](https://pkg.go.dev/k8s.io/client-go/tools/leaderelection).

## How does it work

Leader Election is a **distributed lock**, coordinated and managed by Kubernetes. It is implemented in High Availability deployments and lets your apps select one instance as a Leader, so only one can make changes.

## Implementation

Leader Election is a mechanism that could be implemented through different patterns. We are using here a basic controller lease approach.
Our image acts as a **sidecar container** of the HA image, and selects which of the instance will be the new leader. Once elected, the main container can check which pod is the leader by checking an open endpoint in `localhost:4040`.

An example deployment of the sidecar container could be:

```yaml
- name: leader-election
    image: [build image]
    args:
    - -node-id=$(POD_NAME)
    - -namespace=$(NAMESPACE)
    env:
    - name: POD_NAME
        valueFrom:
        fieldRef:
            apiVersion: v1
            fieldPath: metadata.name
    - name: NAMESPACE
        valueFrom:
        fieldRef:
            apiVersion: v1
            fieldPath: metadata.namespace
```

* **node-id** will determine the id used by the elected leader.
* **namespace** determines the OpenShift namespace in which the lease is deployed.

## Code

We mainly use three external libraries to deploy the code:

* [Gin](github.com/gin-gonic/gin): Dependencie to create the endpoint.
* [Apex](github.com/apex/log): For advanced logging.
* [client-go](k8s.io/client-go/tools): Library to interact with Kubernetes.
