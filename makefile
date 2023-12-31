test:
	go test -v -count=1 ./...

test50:
	go test -v -count=50 ./...

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out