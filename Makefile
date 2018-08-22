test:
	@echo "Running Tests"
	@go test ./...

fixfmt:
	@echo "Fixing format"
	@gofmt -w -l .