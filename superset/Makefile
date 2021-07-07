IMAGE_NAME?=centos/python-38-centos7

.PHONY: build
build:
	s2i build . centos/python-38-centos7 $(IMAGE_NAME) -E .s2i/environment

.PHONY: test
test:
	container-structure-test test -i $(IMAGE_NAME) -c test-config.yaml
