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
	
migration-up:
	migrate -path ./migrations/ -database 'postgres://admin:admin@localhost:5432/project?sslmode=disable' up
	
migration-down:
	migrate -path ./migrations/ -database 'postgres://admin:admin@localhost:5432/project?sslmode=disable' down
	
migration-force-1v:
	migrate -path ./migrations/ -database 'postgres://admin:admin@localhost:5432/project?sslmode=disable' force 1
	
	migrate -path ./migrations/ -database 'postgres://admin:admin@localhost:5432/project?sslmode=disable' up 1 01_add_column.up