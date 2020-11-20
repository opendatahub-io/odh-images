# Cloudera Hue

Produces the Hue image usable in ODH.

## Features

- Provides `centos:8` based build
- Hue version `4.8.0`
- S3 credential scripts that can be used in `hue.ini` to pull S3 credentials from environment variables
- Multistage minimal build, stripping all the build time requirements.
- Works well with `quay.io/opendatahub/spark-cluster-image:2.4.3-h2.7`.

## Build

Use podman/docker to build and release:

```sh
podman build . -t quay.io/opendatahub/hue:<TAG>
podman push quay.io/opendatahub/hue:<TAG>
```

## Downstream changes

Due to the nature of our Hue usage, we have to modify the upstream code a bit by cherry-picking out of release commits:

- [To make MySql connector build and work properly](https://docs.gethue.com/administrator/installation/dependencies/#mysql--mariadb) we cherry-pick `7a9100d4a7` and `e67c1105b8`.
