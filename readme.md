# Example for using RabbitMQ with Golang
This repository contains the example code for publishing and reading a message to and from an RabbitMQ message queue. I have also written a [blog](https://medium.com/@hinsulak/using-rabbitmq-with-golang-and-docker-e674831c959c) for the same, feel free to check it out. 
This repository also contains Dockerfile and docker-compose.yaml files for deploying these services using Docker.

## Run locally
### Producer
Follow these steps to run the producer service locally,
- At the repository root, execute the following command to open the producer directory.
`cd .\producer\`
- Get dependencies.
`go get .`
- Set the environment variables - use the sample .env file provided in the repository.
- Execute the `go run main.go` command to start the service.
### Consumer
Follow these steps to run the consumer service locally,
- At the repository root, execute the following command to open the consumer directory.
`cd .\consumer\`
- Get dependencies.
`go get .`
- Set the environment variables - use the sample .env file provided in the repository.
- Execute the `go run main.go` command to start the service.
### RabbitMQ
- There is a docker-compose.yaml file in the repository.
- Uncomment the `- "5000:5673"` to expose the RabbitMQ instance outside docker.
- [Optionally] Comment out the `producer` and `consumer` service specs so that they are not deployed in the docker.

## Running in Docker
- Running in Docker is very simple, the repo includes a `docker-compose.yaml` file.
- At the repository root, execute `docker-compose up` command to deploy the RabbitMQ instance, producer, and consumer services.
- The producer REST API will be available at `http://localhost:5050/v1/publish/example` if the default configuration is used.
