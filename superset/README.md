
# Container image for Apache superset

## Getting started  

### Creating the application image

It uses `make` to build the image.

    export IMAGE_NAME=<your-desired-image-name>
    make build

### Test the application image

It uses [`container-structure-test`](https://github.com/GoogleContainerTools/container-structure-test)
to test built application image:

    make test

### Running the application image
Running the application image is as simple as invoking the docker run command:
```
podman run -d -p 8088:8088 <your-desired-image-name>
```

The application should now be accessible at  [http://localhost:8088](http://localhost:8088).
