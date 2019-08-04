all: app server

.PHONY: run
run: app server
	./server

.PHONY: server
server:
	go build -o server ./server.go

.PHONY: app
app:
	GOOS=js GOARCH=wasm go build -o html/happycount.wasm ./happycount.go

.PHONY: cleanup
clean:
	rm server
	rm ./html/*.wasm
