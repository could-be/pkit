package templatevar

const (
	MakefileTemplate = `
APP={{.ProjectName}}
DOCKER_REPO=""
IMAGE_NAME="$(DOCKER_REPO)/$(APP)"
VERSION=0.0.1

.PHONY: all
all: build run

doc:
	swag init

.PHONY: src
src:
	pkit update

build:
	GOOS=linux go build -o ${APP}
	docker build -t ${IMAGE_NAME}:${VERSION} .

push:
	docker push ${IMAGE_NAME}:${VERSION}

.PHONY: local
local:
	go build -race
	ENV=dev \
	./$(APP) -v=4 -alsologtostderr=true -log_dir=./logs

.PHONY: run
run: build
	docker run -d \
	--mount source=,target=/opt \
	${APP}

test:
	go test -v ./...

clean:
	rm ${APP}

`
)
