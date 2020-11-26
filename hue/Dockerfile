FROM registry.centos.org/centos/centos:8 AS base

# Install runtime dependencies
RUN dnf install -y cyrus-sasl-gssapi cyrus-sasl cyrus-sasl-plain mysql mysql-devel openldap python2 sqlite libkadm5 libxslt java-1.8.0-openjdk && \
    dnf clean all -y && \
    rm -rf /var/cache/dnf

FROM base AS builder

# Install build dependencies
RUN dnf module enable -y nodejs:12 && \
    dnf install -y \
    ant asciidoc cyrus-sasl-devel openldap-devel python2-devel sqlite-devel gmp-devel \
    gcc gcc-c++ krb5-devel libffi-devel libxml2-devel libxslt-devel make java-1.8.0-openjdk-devel \
    git nodejs rsync

# Clone source and "make install"
ENV HUE_VERSION=4.8.0

RUN git clone https://github.com/cloudera/hue.git -b release-$HUE_VERSION /tmp/hue &&\
    cd /tmp/hue && \
    git cherry-pick -n 7a9100d4a7f38eaef7bd4bd7c715ac1f24a969a8 e67c1105b85b815346758ef1b9cd714dd91d7ea3 14dc72c4e087eaba8be1a84b340ad7e64d493133 && \
    rm /tmp/hue/desktop/conf/* && \
    PREFIX=/opt make install && \
    rm -rf /opt/hue/node_modules && \
    # Create scripts enabling Hue to load S3 credentials from environment, used in hue.ini
    echo 'echo $AWS_ACCESS_KEY_ID' > /opt/hue/aws_access_key_id.sh && \
    echo 'echo $AWS_SECRET_ACCESS_KEY' > /opt/hue/aws_secret_access_key.sh && \
    chmod +x /opt/hue/aws_access_key_id.sh /opt/hue/aws_secret_access_key.sh && \
    chown -R 1001 /opt/hue

# Install missing python packages
RUN /opt/hue/build/env/bin/pip install \
    redis django_redis flower pyhive gevent threadloop thrift-sasl==0.2.1

FROM base

# Copy the built application
COPY --from=builder /opt/hue /opt/hue
WORKDIR /opt/hue
USER 1001
ENV DESKTOP_LOG_DIR=/tmp/logs
VOLUME [ "/opt/hue/desktop/conf", "/etc/hive/conf" ]
EXPOSE 8000

ENTRYPOINT ["build/env/bin/hue" ]
CMD [ "runcpserver" ]
