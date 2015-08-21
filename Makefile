.PHONY: all clean

all: TaipeiCompany.go 
	@go generate
	@go build

clean:
	@go clean
	rm -f *_resource.go
	rm -f main.go
	rm -f web_service.go

test:
	@go test
