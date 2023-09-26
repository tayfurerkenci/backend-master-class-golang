# Backend Master Class (Golang + Postgres + Kubernetes + gRPC) 
I'm following [TECH SCHOOL](https://www.udemy.com/user/tech-school/)'s [Udemy Course](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes) to learn Go.

## Project Setup

### Instal WSL on Windows

On Terminal:
```bash
wsl --install
```

### Instal Go, Sql and Docker Desktop

On Ubuntu Linux (wsl):
```bash
sudo apt update
sudo apt install make
sudo snap install go --classic
sudo snap install sqlc
```

On Terminal:
```bash
Start-Process "Docker Desktop Installer.exe" -Verb RunAs -Wait -ArgumentList "install --installation-dir=C:\Docker\"
docker pull postgres:15.4-alpine
```

### Working with docker containers
```bash
docker run --name <container_name> -e <environment_variable> -d <image>:<tag>

docker run --name postgres15.4 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=tayfur -d postgres:15.4-alpine

docker exec -it <container_name_or_id> <command> [args]
docker exec -it postgres15.4 psql -U root

docker logs postgres15.4

docker stop <container_name>
docker stop postgres15.4

docker start postgres15.4
docker exec -it postgres15.4 /bin/bash

docker exec -it postgres15.4 createdb --username=root --owner=root simple_bank
docker exec -it postgres15.4 dropdb simple_bank
docker exec -it postgres15.4 psql -U root simple_bank
```

### Golang-Migrate Database Migration
```bash
scoop install migrate
migrate create -ext sql -dir db/migration -seq init_schema

scoop install make
make postgres
make createdb
make migratedown
make migrateup
```

### Generate Go code from SQL with sqlc

```bash
sqlc init
make generate
```

