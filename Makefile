.PHONY: test fmt lint showcover

showcover:
	@go tool cover -html=cp.out
	
lint:
	@make fmt
	@golangci-lint --skip-dirs cdk.out run

test:
	@make lint && go test -coverprofile cp.out $$(go list ./...)
			
fmt:
	@gofmt -s -w .
