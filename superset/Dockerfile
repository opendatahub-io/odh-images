FROM docker.io/centos/python-36-centos7

# Superset version
ARG SUPERSET_VERSION=0.30

# Configure environment
ENV GUNICORN_BIND=0.0.0.0:8088 \
    GUNICORN_LIMIT_REQUEST_FIELD_SIZE=0 \
    GUNICORN_LIMIT_REQUEST_LINE=0 \
    GUNICORN_TIMEOUT=60 \
    GUNICORN_WORKERS=2 \
    PYTHONPATH=/etc/superset:/opt/superset/work-dir:$PYTHONPATH \
    SUPERSET_REPO=apache/incubator-superset \
    SUPERSET_VERSION=${SUPERSET_VERSION} \
    SUPERSET_HOME=/var/lib/superset
ENV GUNICORN_CMD_ARGS="--workers ${GUNICORN_WORKERS} --timeout ${GUNICORN_TIMEOUT} --bind ${GUNICORN_BIND} --limit-request-line ${GUNICORN_LIMIT_REQUEST_LINE} --limit-request-field_size ${GUNICORN_LIMIT_REQUEST_FIELD_SIZE}"

USER 0

COPY bin/superset-init /usr/local/bin
COPY bin/entrypoint.sh /usr/local/bin

# Create superset user & install dependencies
RUN mkdir /etc/superset  && \
    mkdir ${SUPERSET_HOME} && \
    chgrp -R 0 /etc/superset && chmod -R g=u /etc/superset && \
    chgrp -R 0 ${SUPERSET_HOME} && chmod -R g=u ${SUPERSET_HOME} && \
    chmod +x /usr/local/bin/superset-init && \
    chmod g=u /etc/passwd /etc/group && \
    yum update -y && \
    yum install cyrus-sasl-plain -y
RUN curl https://raw.githubusercontent.com/${SUPERSET_REPO}/${SUPERSET_VERSION}/requirements.txt -o requirements.txt && \
    pip install --no-cache-dir -r requirements.txt && \
        pip install --no-cache-dir \
        thrift==0.13.0 \
        thrift-sasl==0.3.0 \
        sasl==0.2.1 \
        Werkzeug==0.14.1 \
        flask-cors==3.0.3 \
        flask-mail==0.9.1 \
        flask-oauth==0.12 \
        flask_oauthlib==0.9.3 \
        gevent==1.2.2 \
        impyla==0.14.0 \
        infi.clickhouse-orm==1.0.2 \
        mysqlclient==1.4.4 \
        psycopg2==2.8.3 \
        pyathena==1.2.5 \
        pyhive==0.5.1 \
        pyldap==2.4.28 \
        pymssql==2.1.4 \
        redis==2.10.5 \
        sqlalchemy-clickhouse==0.1.5.post0 \
        sqlalchemy-redshift==0.7.1 \
        superset==${SUPERSET_VERSION} && \
    rm requirements.txt

VOLUME /etc/superset \
       /var/lib/superset

# Deploy application
EXPOSE 8088
HEALTHCHECK CMD ["curl", "-f", "http://localhost:8088/health"]
CMD ["gunicorn", "superset:app"]

RUN yum install cyrus-sasl-plain -y
USER 185
