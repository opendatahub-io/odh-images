# Trino image

This image is a customization of the original Trino image to add JMX exporter jars.

# Build instructions 

From the root folder, run the build command.

## Using podman

```
podman build -t quay.io/opendatahub/trino trino/
```

## Using docker

```
podman build -t quay.io/opendatahub/trino trino/
```
