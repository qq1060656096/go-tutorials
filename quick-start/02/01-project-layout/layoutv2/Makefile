API_PROTO_FILES=$(shell find api -name *.proto)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)

.PHONY: api
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
 	       --openapi_out==paths=source_relative:. \
	       $(API_PROTO_FILES)
    protoc --proto_path=. \
    	       --proto_path=./third_party \
     	       --go_out=paths=source_relative:. \
    	       $(INTERNAL_PROTO_FILES)

.PHONY: generate
# generate
generate:
	go generate ./...


.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;
	make generate;

# show help
# help:
# 	@echo ''
# 	@echo 'Usage:'
# 	@echo ' make [target]'
# 	@echo ''
# 	@echo 'Targets:'
# 	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
# 	helpMessage = match(lastLine, /^# (.*)/); \
# 		if (helpMessage) { \
# 			helpCommand = substr($$1, 0, index($$1, ":")-1); \
# 			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
# 			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
# 		} \
# 	} \
# 	{ lastLine = $$0 }' $(MAKEFILE_LIST)
#
# .DEFAULT_GOAL := help
