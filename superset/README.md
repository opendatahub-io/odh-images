
# s2i image for Apache superset

## Getting started  

### Files and Directories  
| File                   | Required? | Description                                                  |
|------------------------|-----------|--------------------------------------------------------------|
| .s2i/bin/assemble      | Yes       | Script that builds the application                           |
| .s2i/bin/usage         | No        | Script that prints the usage of the builder                  |
| .s2i/bin/run           | Yes       | Script that runs the application                             |
| .s2i/bin/save-artifacts| No        | Script for incremental builds that saves the built artifacts |
| .test-config.yaml      | No        | Test config of container-structure-test                      |

### Creating the application image

It uses `s2i` to build the image. See this guide [How to Create an S2I Builder Image](https://www.openshift.com/blog/create-s2i-builder-image).

    export IMAGE_NAME=<your-desired-image-name>
    make build

### Test the application image

It uses [`container-structure-test`](https://github.com/GoogleContainerTools/container-structure-test)
to test built application image:

    make test

### Running the application image
Running the application image is as simple as invoking the docker run command:
```
docker run -d -p 8080:8080 <your-desired-image-name>
```

The application should now be accessible at  [http://localhost:8080](http://localhost:8080).
