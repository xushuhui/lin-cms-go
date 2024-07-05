INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
BIZ_FILES=$(shell find internal/biz -name *.go)
.PHONY: init
init:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/favadi/protoc-go-inject-tag@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest

.PHONY: generate
# generate
generate:
	
	go generate ./...

.PHONY: mock
mock:
	mockgen -source=internal/biz/area.go -destination mocks/biz/area.go
	mockgen -source=internal/biz/report.go -destination mocks/biz/report.go	
	mockgen -source=internal/biz/comment.go -destination mocks/biz/comment.go
	mockgen -source=internal/biz/draft.go -destination mocks/biz/draft.go
	mockgen -source=internal/biz/follow.go -destination mocks/biz/follow.go
	mockgen -source=internal/biz/hotspot.go -destination mocks/biz/hotspot.go
	mockgen -source=internal/biz/like.go -destination mocks/biz/like.go
	mockgen -source=internal/biz/missionUser.go -destination mocks/biz/missionUser.go
	mockgen -source=internal/biz/notify.go -destination mocks/biz/notify.go
	mockgen -source=internal/biz/reply.go -destination mocks/biz/reply.go
	mockgen -source=internal/biz/report.go -destination mocks/biz/report.go
	mockgen -source=internal/biz/suggest.go -destination mocks/biz/suggest.go	
	mockgen -source=internal/biz/visit.go -destination mocks/biz/visit.go	
	
	

.PHONY: test
test:
	go test -cover lin-cms-go/internal/biz -coverprofile ./coverage.out 
	go tool cover -html=./coverage.out -o coverage.html

.PHONY: build
build:
	go build -toolexec="/apps/apm/skywalking-go/go-agent" -a -o test -ldflags="-s -w" -o ./bin/server ./cmd/server


.PHONY: swag
swag:
	swag init  -g ./cmd/server/main.go -o ./docs

.PHONY: config
config:
	protoc --proto_path=./internal \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
		   --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
		   --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       $(API_PROTO_FILES)
	protoc-go-inject-tag -input="./api/*.pb.go"
	ls ./api/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'


