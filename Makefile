.PHONY: all clean

all: dataSchema.go 


clean:
	@go clean
	rm -f *_resource.go
	rm -f main.go
	rm -f web_service.go

test:
	@go test
