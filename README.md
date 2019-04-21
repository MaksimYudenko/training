***

<h2>Golang web application

***

* MicroServices
* HTTP
* REST
* gRPC
* Swagger
* PostgresQL
***

<h3> Local running

* a better way - to use Makefile, but I did not make it in time (

* run `bash ./scripts/protoBuild.sh`

* create database by executing `./scripts/schema.sql`

* SERVER: `go run ./cmd/server/main.go` with appropriate args placed at the bottom of this file
* CLIENT-gRPC: `go run ./cmd/client-grpc/main.go` with appropriate args placed at the bottom of this file
* CLIENT-REST: `go run ./cmd/client-rest/main.go` with appropriate args placed at the bottom of this file

* primitive database tests are embedded

***

<h3> Docker

* a better way - to use docker-compose, but it doesn't work correctly (for now)

* to achieve predictable routing define your own bridge network `sudo docker network create raining_network`

* uncomment appropriate lines in Dockerfile (for now)

* to build image for server [client-grpc or client-rest] run `docker build -t server [client-grpc or client-rest] .` 
respectively

* to run server image do `docker run --rm --name serverC --network-alias mynet --network training_network -it server`

* to run client-grpc image do `docker run --rm --name grpcClientC --network training_network -it client-grpc`

* to run client-rest image do `docker run --rm --name restClientC --network training_network -it client-rest`

* to run database fetch postgres image
 and run `docker run --rm --name postgresC -e POSTGRES_USER=<user> -e POSTGRES_PASSWORD=<password> -e 
 POSTGRES_DB=training -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres`

***

<h3>Program arguments

* Server `-grpc-port=9090 -http-port=8080 -db-host=localhost -db-port=5432 -db-name=training -db-user=<user> 
-db-password=<password> -db-schema=public`

* gRPC-client `-server=mynet:9090`

* REST-client `-server=http://mynet:8080`

***

<h3>ToDo

* fix `failed to connect to database` while starting client-grpc (localhost might be a cause)

* fix `failed to connect to database` while starting client-rest (localhost might be a cause)
