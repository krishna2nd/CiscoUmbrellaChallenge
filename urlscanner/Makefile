.PHONY: build
build: docker-compose.yml
	$(info --- Building source.)
	@env GOPATH=`pwd` go get -u github.com/go-redis/redis
	@env GOPATH=`pwd` GOOS=linux go build -o uscanner src/main.go
	$(info --- Building docker image)
	@docker-compose build  uscanner
	$(info --- Stopping running docker containers)
	@docker stop  $(shell docker ps -a -q --filter="name=urlscanner_uscanner")  &> /dev/null; true
	$(info --- Starting uscanner container)
	@docker-compose run --rm -d -p 9090:9090  uscanner

.PHONY: test
test: docker-compose.yml
	$(info --- Building docker image)
	docker-compose build test
	$(info --- Running test container)
	docker-compose run --rm test

.PHONY: run-test
run-test:
	@go test -v ucache uhandlers
