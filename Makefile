
.PHONY:build
build:
	# build app
	@CGO_ENABLED=0 go build -o ./jsoncli -a -ldflags '-s' -installsuffix cgo ./cmd/jsoncli/main.go

.PHONY:install
install:
	# install dependencies from gopkg file
	@go get -u golang.org/x/lint/golint
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure

.PHONY: test
test:
	# launch test across all project
	@go test -race ./...
	@go list ./... | grep -v /vendor/ | grep -v pb | xargs -L1 golint -set_exit_status
	@go vet `go list ./... | grep -v /vendor/`


.PHONY:coverage
coverage:
	# use go ability to generate an html with test coverage
	@go test `go list ./... | grep -v /vendor/` -cover -coverprofile=cover.out
	@go tool cover -html=cover.out

.PHONY:docker
docker:
	# build the docker image
	@docker build --tag=jsoncli .