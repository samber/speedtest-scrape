
all: get_vendor build

get_vendor:
	go get -u github.com/PuerkitoBio/goquery

build:
	go build main.go

run-dev:
	go run main.go

fmt:
	go fmt main.go
