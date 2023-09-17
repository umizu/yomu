build: 
	@go build -o ./bin/yomu ./cmd/yomu/main.go

run: build
	@./bin/yomu
