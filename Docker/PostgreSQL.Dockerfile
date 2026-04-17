FROM postgres:latest

EXPOSE 9999:9999

COPY ../PostgreSQL /usr/src/

WORKDIR /usr/src/PostgreSQL