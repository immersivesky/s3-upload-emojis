FROM postgres:16.0-alpine3.18

ENV POSTGRES_HOST=emo
ENV POSTGRES_USER=studio
ENV POSTGRES_PASSWORD=rootpassword
ENV POSTGRES_DB=emoji

COPY migrations/000001_init.up.sql /docker-entrypoint-initdb.d