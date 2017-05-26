## Dali Server Setup

Steps to get server built and run locally using Docker and Docker-Compose.

Prerequisite for OSX: [Docker for Mac](https://docs.docker.com/docker-for-mac/)

Prerequisite for Linux: [Docker](https://docs.docker.com/engine/installation/) and [Docker-Compose](https://docs.docker.com/compose/)

### Compile Dali and Build Containers

To compile the Dali Server and build the containers:

	./scripts/build_it.sh

### Run Dali Server

To run Dali Server and MongoDB locally:

	./scripts/run_it.sh

MongoDB is writing its data in to the `/tmp` folder of your Mac. This can be changed in the `docker-compose.yml` file.

Simple ping test using curl:

	curl -X GET -H "Accept: application/json" -H "Content-Type: application/json" 127.0.0.1:8085/admin/ping

Response from server:

	{"code":202,"message":"pong"}

### Stop Dali Server

To stop local Dali Server and MongoDB:

	./scripts/stop_it.sh