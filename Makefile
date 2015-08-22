.PHONY: all clean

all: taipeiCity.go 
	@go generate
	@go build

clean:
	@go clean
	rm -f *_resource.go
	rm -f main.go
	rm -f web_service.go

cleandb: TaipeiCity.db
	rm -f TaipeiCity.db

migrate:
	./hackNTU2015 migratedb

run:
	./hackNTU2015 serve

test:
	@go test
