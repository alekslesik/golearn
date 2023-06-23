pwd = /root/go/src/github.com/alekslesik/golearn

# Include variables from the .envrc file
# include .envrc

#=====================================#
# DEVELOPMENT #
#=====================================#

## run: run the cmd/app application
.PHONY: run
run:
	go run .

#=====================================#
# HELPERS #
#=====================================#

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

#=====================================#
# DOCKER #
#=====================================#

## docker-up: build and run docker container
.PHONY: docker-up
docker-up: 
	docker build -t golearn .
	# docker run -it --rm --name golearn golearn

## docker-build: compile, but not run your app inside the Docker instance
.PHONY: docker-build
docker-build:
	docker run --rm -v $(pwd):$(pwd) -w $(pwd) golang:1.20 go build -v
	./golear

## docker-run: compile, and execite make run command inside container
.PHONY: docker-run
docker-run:
	docker container rm -f golearn
	docker run -d --rm --name golearn -p 9999:9999 -v $(pwd):$(pwd) -w $(pwd) golang:1.20 make run

## docker-restart: compile, and execite make run command inside container
.PHONY: docker-restart
docker-restart:
	docker container restart golearn

## docker-exec: compile, and execite make run command inside container
.PHONY: docker-exec
docker-exec:
	docker container exec -it golearn make run