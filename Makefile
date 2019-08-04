all: app

.PHONY: run
run: app
	echo "starting server listening on http://localhost:8080"
	shareh -port 8080

.PHONY: app
app: jsuport
	GOOS=js GOARCH=wasm go build -o happycount.wasm happycount.go

.PHONY: jsupport
jsuport:
	cp $(GOROOT)/misc/wasm/wasm_exec.js wasm_exec.js
