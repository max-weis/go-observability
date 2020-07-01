default: run

VERSION ?= v0.1.0
APP ?= go-observability
DOCKER_IMAGE ?= maxweis/${APP}

ut:
	@go test ./... -tags=unit

it:
	@go test ./... -tags=integration

build:
	@go build main.go

run:
	@go run main.go

docker-build:
	@docker image build -t $(DOCKER_IMAGE):$(VERSION) .

docker-push:
	# Push to registry
	@docker image push $(DOCKER_IMAGE):latest
	@docker image push $(DOCKER_IMAGE):$(VERSION)
