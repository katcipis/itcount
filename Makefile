all: app server

.PHONY: run
run: app jsuport
	echo "starting server listening on http://localhost:8080"
	cd html; shareh -port 8080

.PHONY: app
app:
	GOOS=js GOARCH=wasm go build -o html/happycount.wasm ./happycount.go

.PHONY: jsupport
jsuport:
	cp $(GOROOT)/misc/wasm/wasm_exec.js ./html/wasm_exec.js

.PHONY: cleanup
clean:
	rm server
	rm ./html/*.wasm
