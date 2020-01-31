FROM centos
USER root
RUN yum update -y && \
    yum install -y ant asciidoc cyrus-sasl-devel cyrus-sasl-gssapi cyrus-sasl-plain epel-release\
    gcc gcc-c++ krb5-devel libffi-devel libxml2-devel libxslt-devel make openldap-devel \
    python-devel sqlite-devel gmp-devel openssl-devel mysql mysql-devel wget git rsync maven \
    java-1.8.0-openjdk-devel && \
    yum clean all -y && \
    git clone https://github.com/cloudera/hue.git

WORKDIR hue

RUN yum install -y npm && \
    make apps && \
    useradd -ms /bin/bash hue

EXPOSE 8000

ENV DESKTOP_LOG_DIR=/tmp/logs


VOLUME /hue/desktop/
