CURRENT_DIR=$(shell pwd)
APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

gen-proto-module:
	sudo rm -rf ${CURRENT_DIR}/genproto
	./scripts/gen_proto.sh ${CURRENT_DIR}

clean-proto-module:
	sudo rm -rf ${CURRENT_DIR}/genproto

.PHONY: gen-proto-module clean-proto-module


swag-init:
	swag init -g api/router.go -o api/docs