export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

.EXPORT_ALL_VARIABLES: ;

.PHONY: build
build:
	docker-compose -f docker-compose.build.yaml build

.PHONY: run-dev
run-dev:
	docker-compose -f docker-compose.yaml -f docker-compose.build.yaml up -d

.PHONY: lint
lint:
	golangci-lint run --fix --deadline=5m ./...

.PHONY: stop
stop:
	docker-compose stop

.PHONY: clean
clean:
	docker-compose down -v

swagger = docker run --rm -it -e GOPATH=/go \
			-u $(shell id -u ${USER}):$(shell id -g ${USER}) \
			-v $(shell pwd):/go/src \
			-w $(shell pwd)/src quay.io/goswagger/swagger

generate:
	rm -fr openbanking
	mkdir -p openbanking/accountinformation openbanking/paymentinitiation
	${swagger} generate client -f /go/src/api/accounts.yaml -A OpenbankingAccountsClient -t /go/src/openbanking/accountinformation
	${swagger} generate client -f /go/src/api/paymentinitiation.yaml -A OpenbankingPaymentsClient -t /go/src/openbanking/paymentinitiation
