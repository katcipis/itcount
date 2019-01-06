all: app server

.PHONY: run
run: app server
	./server

.PHONY: server
server:
	go build -o server ./server.go

.PHONY: app
app:
	GOOS=js GOARCH=wasm go build -o html/itcount.wasm ./itcount.go

.PHONY: cleanup
clean:
	rm server
	rm ./html/*.wasm
