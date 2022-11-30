FROM registry.access.redhat.com/ubi8/python-38:1-86

# Superset version
ARG SUPERSET_VERSION=1.5.2

# Configure environment
ENV GUNICORN_BIND=0.0.0.0:8088 \
    GUNICORN_LIMIT_REQUEST_FIELD_SIZE=0 \
    GUNICORN_LIMIT_REQUEST_LINE=0 \
    GUNICORN_TIMEOUT=60 \
    GUNICORN_WORKERS=2 \
    PYTHONPATH=/etc/superset:/opt/superset/work-dir:$PYTHONPATH \
    SUPERSET_REPO=apache/incubator-superset \
    SUPERSET_VERSION=${SUPERSET_VERSION} \
    SUPERSET_HOME=/var/lib/superset \
    FLASK_APP=superset
ENV GUNICORN_CMD_ARGS="-w ${GUNICORN_WORKERS} --timeout ${GUNICORN_TIMEOUT} -b ${GUNICORN_BIND} --limit-request-line ${GUNICORN_LIMIT_REQUEST_LINE} --limit-request-field_size ${GUNICORN_LIMIT_REQUEST_FIELD_SIZE}"

USER root

COPY bin/superset-init /usr/local/bin

RUN mkdir /etc/superset  && \
    mkdir ${SUPERSET_HOME} && \
    chgrp -R 0 /etc/superset && chmod -R g=u /etc/superset && \
    chgrp -R 0 ${SUPERSET_HOME} && chmod -R g=u ${SUPERSET_HOME} && \
    chmod +x /usr/local/bin/superset-init && \
    chmod g=u /etc/passwd /etc/group && \
    dnf update -y && \
    dnf install -y gcc \
                gcc-c++ \
                libffi-devel \
                python38-devel \
                python38-pip \
                python38-wheel \
                openssl-devel \
                cyrus-sasl-devel \
                openldap-devel && \
    yum clean all && \
    rm -rf /var/yum/cache/* && \
    pip install sqlalchemy-trino \
                pyhive \
                markupsafe==2.0.1 \
                psycopg2 \
                authlib==0.15.1 \
                python-ldap \
                apache-superset==${SUPERSET_VERSION}

VOLUME /etc/superset \
       /var/lib/superset

# Deploy application
EXPOSE 8088
HEALTHCHECK CMD ["curl", "-f", "http://localhost:8088/health"]

USER 185

CMD ["gunicorn", "superset.app:create_app()"]
