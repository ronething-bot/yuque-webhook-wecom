# 构建脚本

.PHONY: deploy
deploy: build-linux upx-linux

.PHONY: build
build: set-env copy-config
	go build -v -o bin/yuque-webhook-wecom cmd/main.go
	@echo "build yuque-webhook-wecom success"

.PHONY: build-linux
build-linux: set-env copy-config
	GOOS=linux GOARCH=amd64 go build -v -o bin/yuque-webhook-wecom cmd/main.go
	@echo "build yuque-webhook-wecom success"

.PHONY: copy-config
copy-config:
	rm -rf bin && mkdir -p bin && cp config/*.yaml bin/
	@echo "copy config success"

.PHONY: set-env
set-env:
	export GO111MODULE=on
	export GOPROXY=https://goproxy.io
	@echo "set env success"

.PHONY: upx-linux
upx-linux:
	upx -v bin/yuque-webhook-wecom