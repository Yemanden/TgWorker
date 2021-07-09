PROJECT_NAME := tg-worker
ORG_NAME := yemanden
REPO_PATH := $(ORG_NAME)/$(PROJECT_NAME)
IMAGE_NAME ?= $(REPO_PATH)/$(PROJECT_NAME)

build: Dockerfile
	docker build \
		--build-arg "APP_PKG_NAME=$(REPO_PATH)" \
		-t $(IMAGE_NAME) .