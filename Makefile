BINARY=engine
format:
	@find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/katsukiniwa/practical-go-programming"
verify:
	@staticcheck ./... && go vet ./...
unit-test:
	@go test ./... -coverprofile=./test_results/cover.out && go tool cover -html=./test_results/cover.out -o ./test_results/cover.html
serve:
	@docker-compose -f docker-compose.yml up
engine:
	go build -o ${BINARY} app/*.go
