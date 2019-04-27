wasm:
	GOARCH=wasm	GOOS=js	go build -o lib.wasm lib.go

server:
	go build -o bin/serve server.go

serve:
	./bin/serve

clean:
	rm -rf bin
	rm *.wasm
