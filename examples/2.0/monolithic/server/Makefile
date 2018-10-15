petstore:
	CGO_ENABLED=0 GOOS=linux go build -i -v  -o $@ *.go

docker: petstore
	docker build -t soloio/petstore-example:latest .
