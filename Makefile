build:
	@go build -o ff cmd/factsfood/main.go
run: build
	@./ff
