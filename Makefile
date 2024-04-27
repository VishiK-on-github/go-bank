# for local application development
build_server:
		@go build -o bin/gobank

run_server: build_server
		 @./bin/gobank

test:
		@go test -v ./...

# for local postgres database
start_db:
		docker start gobank-postgres

run_db:
		docker run --name gobank-postgres -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres

stop_db:
		docker stop gobank-postgres