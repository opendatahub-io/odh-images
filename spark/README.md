## Spark image build instructions

The spark image is located in [Quay.io](https://quay.io/repository/opendatahub/spark-cluster-image), but the image must be updated in order to upgrade
spark or hadoop version, or even enable different components like yarn integration.

### Build a custom Spark distribution

It is recommended to build a spark distribution from the source code, as the Hadoop jars embedded in spark 2.4 releases are very old. To build


* Clone the [git repo](https://github.com/apache/spark.git)
* Run `git checkout branch-2.4` to move to a 2.4 release branch
* Run the following command to build the custom distribution:

```
./dev/make-distribution.sh --name spark-hadoop28 --pip --tgz -Phadoop-2.8 -Dhadoop.version=2.8.4 -Phive -Phive-thriftserver -Pkubernetes -DskipTests
```

After the script execution, you will end up with a tgz file in the git repo dir. Now we can proceed with the image build.

### Build the Spark image

The Spark image is built from a tool created by the [RADAnalytics.io](https://radanalytics.io/) project. This will make our images easier to create as it uses a tool called [CEKit](https://docs.cekit.io/en/latest/) to modularize the steps to create the container image.

To build a Spark image, follow the [instructions](https://github.com/radanalyticsio/openshift-spark#image-completion) in RADAnalytics.io git repo to build the new image and push to your container registry.