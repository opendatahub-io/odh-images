DOCKER_COMMAND ?= sudo docker

DOCKER_IMAGE ?= quay.io/erdii/configmap-puller:dev

build:
	$(DOCKER_COMMAND) build -t $(DOCKER_IMAGE) .
.PHONY: build

push:
	$(DOCKER_COMMAND) push $(DOCKER_IMAGE)
.PHONY: push
