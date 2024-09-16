all: go python

go:
	protoc --go_out=. --go-grpc_out=. deep.proto

python:
	python -m grpc_tools.protoc -I. --python_out=client/proto --pyi_out=client/proto --grpc_python_out=. deep.proto