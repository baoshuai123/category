

.PHONY: proto
proto:
	docker run --rm -v d:/GOLANG/src/taobao/category:/d/GOLANG/src/taobao/category -w /d/GOLANG/src/taobao/category  -e ICODE=2606C833CD172F4C cap1573/cap-protoc -I ./   --go_out=./ --micro_out=./ ./proto/category/category.proto

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o category-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t category-service:latest
