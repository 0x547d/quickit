hub:=$(shell echo "${module}" | awk -F/ '{print $$2}')
app:=quickit
version:=latest
buildAt:=$(shell date "+%Y-%m-%d_%H:%M:%S")
commitId:=$(shell git rev-parse --short HEAD)
branch:=$(shell git symbolic-ref --short -q HEAD)

.phony: all
all: image

.phony: image
image:
	@-docker buildx rm tmp
	@-docker buildx create --name tmp --bootstrap --use
	docker buildx build -t ${hub}/${app}:${version} \
		--build-arg app=${app} \
		--build-arg commitId=${commitId} \
		--build-arg goProxy=${goProxy} \
		--platform linux/386,linux/amd64,linux/arm64 \
		--push \
		.
	@-docker buildx rm tmp

.phony: binary
binary:
	@echo ${hub}
	@rm -rf build/*
	CGO_ENABLED=0 go build -ldflags " \
		-X '${module}/pkg/version.AppName=${app}' \
		-X '${module}/pkg/version.AppCommitId=${commitId}' \
		-X '${module}/pkg/version.AppBranch=${branch}' \
		-X '${module}/pkg/version.AppVersion=${version}' \
		-X '${module}/pkg/version.AppBuildAt=${buildAt}' \
	" \
	-o build/${app} internal/*.go


