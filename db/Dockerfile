FROM postgres:16.0-alpine3.18

ENV POSTGRES_USER=qsadmin
ENV POSTGRES_PASSWORD=pass
ENV POSTGRES_DB=emoji

COPY migrations/000001_init.up.sql /docker-entrypoint-initdb.d