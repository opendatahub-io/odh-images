SUPERSET_VERSION=1.5.0-ubi
IMAGE_NAME=quay.io/opendatahub/superset:${SUPERSET_VERSION}

BUILDER:=$(shell command -v podman 2> /dev/null || command -v docker 2> /dev/null || echo "docker or podman not found")

all: build push

build:
	echo "Building $(BUILDER) image..."
	$(BUILDER) build -t $(IMAGE_NAME) .

push:
	echo "Pushing $(IMAGE_NAME) image..."
	$(BUILDER) push $(IMAGE_NAME)

test:
	container-structure-test test -i $(IMAGE_NAME) -c test-config.yaml
