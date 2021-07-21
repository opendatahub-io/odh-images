Configmap Puller
----------

This repository contains code for the configmap-puller sidecar that would update the rules.toml file read by the traefik proxy whenever the configmap traefik-rules is updated by Jupyterhub.


### To Build the image run:
```
cd configmap-puller
DOCKER_IMAGE=quay.io/<repo>/configmap-puller:0.1 make build
```
