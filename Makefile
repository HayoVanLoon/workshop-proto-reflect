PROTOC := $(shell which protoc)
VIRTUALENV := $(shell which virtualenv)

VENV := .venv
V := . .venv/bin/activate

GO_OUT := go/lib
PY_OUT := py

PROTO_DIR := proto

FILES := $(shell find proto | grep -E "\.proto$$" | sort)

venv:
	$(VIRTUALENV) -p python3.8 $(VENV)
	$(V) && pip install protobuf

protoc: protoc-go protoc-py

protoc-go:
	rm -rf $(GO_OUT)
	mkdir -p $(GO_OUT)
	$(PROTOC) \
		--go_out=$(GO_OUT) \
		-I$(PROTO_DIR) \
		$(FILES)

protoc-py:
	rm -rf $(PY_OUT)/workshop
	$(PROTOC) \
		--python_out=$(PY_OUT) \
		-I$(PROTO_DIR) \
		$(FILES)

run-all: run-map run-struct run-reflect run-proto run-py

run-map:
	cd go && go run main.go map

run-struct:
	cd go && go run main.go struct

run-reflect:
	cd go && go run main.go reflect

run-proto:
	cd go && go run main.go proto

run-py:
	$(V) && cd py && ./main.py

bench:
	cd go && go test -bench . -benchmem
