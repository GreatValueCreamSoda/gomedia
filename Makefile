MakeFFMSBindings:
	go build -C ./build/clones/c-for-go -o ./cgo
	CC=clang ./build/clones/c-for-go/cgo -out ./ffms ./build/c-for-go-yaml/ffms.yaml
	rm ./build/clones/c-for-go/cgo