# variable for common working directory and build cache arguments
docker_dir_args = -v $(PWD)/src:/usr/src/definition-graph -v $(PWD)/docker/.buildcache/pkg:/go/pkg -w /usr/src/definition-graph nextmetaphor/definition-graph-build:latest

.PHONY: help
help:	## show makefile help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build:	## build definition-graph using a docker build container
	docker build docker --tag nextmetaphor/definition-graph-build:latest -f docker/DockerfileBuild
	# optionally pass GOOS and GOARCH parameters e.g. make build GOOS=darwin GOARCH=amd64
	docker run --rm $(docker_dir_args) ./build.sh $(GOOS) $(GOARCH)

	# copy the built binary to the docker installation files
	cp src/definition-graph docker/utils

test:	## test definition-graph using a docker test container
	docker run --rm $(docker_dir_args) ./test.sh

docker-build: 	## build definition-graph docker image
	docker build --tag nextmetaphor/definition-graph:latest docker -f docker/DockerfileRun

docker-run: docker-build ## run definition-graph docker image
	docker run -it -p8080:8080 -v $(PWD)/definition:/home/dfngraph/definition nextmetaphor/definition-graph