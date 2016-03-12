BUILD_IMAGE=mgurov/tweens_build
test:
	go test .
	go build demo/main.go
	rm -f ./main
	@echo OK
full: test
	go run demo/main.go
dbuild:
	docker build -t $(BUILD_IMAGE) .
dtest:
	docker run -it --rm $(BUILD_IMAGE) go test ./...