test:
	go test ./...
.PHONY: test

run:
	go run main.go
.PHONY: run

update-snapshot:
	UPDATE_SNAPSHOTS=true go test ./...