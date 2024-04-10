FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD root
ENV MYSQL_DATABASE "weather_database"

COPY  scripts/db/*  /docker-entrypoint-initdb.d/