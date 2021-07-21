DOCKER_COMMAND ?= sudo docker

DOCKER_IMAGE ?= quay.io/internaldatahub/configmap-puller:0.1

build:
	$(DOCKER_COMMAND) build -t $(DOCKER_IMAGE) .
.PHONY: build

push:
	$(DOCKER_COMMAND) push $(DOCKER_IMAGE)
.PHONY: push
