FROM mysql:latest

COPY  scripts/db/*  /docker-entrypoint-initdb.d/