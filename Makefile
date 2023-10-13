.PHONY: api
api:
	protoc --proto_path=./pb \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./pb \
	       pb/*.proto

.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
