cover:
	go test -short -count=1 -race -coverprofile=coverage.out src/...
	go tool cover -html=coverage.out
	rm coverage.out

gen:
	/Users/ayto/go/bin/mockgen -source=src/service/interface.go \
            -destination=src/repository/mocks/mock_interface.go

.PHONY: gen cover