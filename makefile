PROJECT_NAME ?= conscience-backend
NAMESPACE = default
DOCKER_NS = isnlan

BUILD_DIR ?= .build
BINARY_FILE = app

SERVICE_NAME = $(NAMESPACE)-$(PROJECT_NAME)
DOCKER_REGISTRY = registry.cn-hangzhou.aliyuncs.com
DOCKER_RUN_GOLANG_IMAGE = golang:1.12

IMAGE_NAME = $(DOCKER_NS)/$(PROJECT_NAME)
EXTRA_VERSION ?= $(shell git rev-parse --short HEAD)
IMAGE_FULL_NAME = $(DOCKER_REGISTRY)/$(IMAGE_NAME):$(EXTRA_VERSION)

USERID = $(shell id -u)
DRUN = docker run -i --rm --user=$(USERID):$(USERID) \
	-v $(abspath .):/go/src/$(PROJECT_NAME) \
	-e GOCACHE=/tmp/.cache \
	-w /go/src/$(PROJECT_NAME)

KUBERNETES_FILE = deployment-template.yaml service-template.yaml

CONFIG_TEST="config-test.yaml"
CONFIG_PROD="config-prod.yaml"

define deploy
	for item in $(KUBERNETES_FILE); do \
		cat .ci/.kubernetes/$$item | \
		sed -e 's|__APP_LABEL__|$(SERVICE_NAME)|g' | \
		sed -e 's|__IMAGE_FULL_NAME__|$(IMAGE_FULL_NAME)|g' | \
		sed -e 's|__CONTAINER_NAME__|$(SERVICE_NAME)|g' | \
		sed -e 's|__NAMESPACE__|$(NAMESPACE)|g' | \
		sed -e 's|__DEPLOY_NAME__|$(SERVICE_NAME)|g' | \
		sed -e 's|__SERVICE_NAME__|$(SERVICE_NAME)|g' | \
		sed -e 's|__CONFIG_FILE__|$(1)|g' | \
		kubectl apply --record -f - ; \
	done
endef


help:
	@echo
	@echo "帮助文档："
	@echo "  - make help              查看可用脚本"
	@echo "  - make dep               安装依赖"
	@echo "  - make build             编译可执行文件"
	@echo "  - make docker            编译Docker镜像"
	@echo "  - make deploy-test       部署测试环境"
	@echo "  - make deploy-prod       部署正式环境"
	@echo "  - make clean             清理.build"
	@echo

dep:
	@export GO111MODULE=on; go mod tidy; go mod vendor

build:
	$(DRUN) \
	  		-e CGO_ENABLED=0 -e GOOS=linux \
	  		$(DOCKER_RUN_GOLANG_IMAGE) \
       		go build -ldflags '-w -extldflags "-static"' -o $(BUILD_DIR)/$(BINARY_FILE) $(PROJECT_NAME)

docker: build
	@docker build -t $(IMAGE_FULL_NAME) . -f image/Dockerfile.in
	@docker push $(IMAGE_FULL_NAME)

deploy-test:
	$(call deploy, $(CONFIG_TEST))

deploy-prod:
	$(call deploy, $(CONFIG_PROD))

clean:
	@rm -rf .build

.PHONY: dep build deploy-test deploy-prod clean
