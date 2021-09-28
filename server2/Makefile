.PHONY: generate build run test serve

generate:
	@protoc --go_out="plugins=grpc:." *.proto
	@./../web/gen_js_lib2.sh
	@go mod tidy
	@go generate ./...
	@echo "[OK] Generate all completed!"

nowTime=$(shell date +00%y%m%d%H%M%S)
gitCID=$(shell git rev-parse HEAD)
gitTime=$(git log -1 --format=%at | xargs -I{} date -d @{} +%Y%m%d_%H%M%S)
build: generate
	@CGO_ENABLED=0 go build -ldflags "-s -w -X main.build=${nowTime}.${gitCID}" -o "./bin/webshell"
	@echo "[OK] App binary was created!"

run:
	@./bin/webshell

test: 
	go test -v ./...

serve: build run