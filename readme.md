# SimpleBank Project with Golang

I admire the Golang language and I learn it through various documents, Udemy and YouTube courses. In this context I am following [TECH SCHOOL's Udemy Course](https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes). You can click for [the original source from here](https://github.com/techschool/simplebank)

### Makefile

All necessary commands are in the Makefile.

### New Migration Command Example

```bash
migrate create -ext sql -dir db/migration -seq add_users
```

### Update golang-migrate

```bash
$env:GO111MODULE="on"; go get -u github.com/golang-migrate/migrate/v4/cmd/migrate
```

### Docker commands

```bash
docker images
docker ps -a
docker rmi <image_id>
docker start <name>
docker build -t simple-bank:latest
docker run --name simple-bank --network simple-network -p 8080:8080 -e GIN_MODE=release simple-bank:latest
docker container inspect <container_name>
docker network ls
docker network inspect <network_name>
docker network create <network_name>
```
