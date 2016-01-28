test:
	go test .
	go build demo/main.go
	rm -f ./main
	@echo OK
full: test
	go run demo/main.go