BINARY=emqx-user-manager

default:
	@go build -o build/${BINARY}

docker:
	@docker build -t emqx-user-manager .

clean:
	@rm -r build/

.PHONY: default docker clean
