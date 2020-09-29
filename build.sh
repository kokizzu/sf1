#!/usr/bin/zsh

swag init

#go build -tags netgo -a -v -o sf1.exe main.go
CGO_ENABLED=0 go build -o sf1.exe main.go 
#go build -ldflags '-linkmode external -w -extldflags "-static"' -o sf1.exe main.go
#go build -o sf1.exe main.go

(cd view; npm run build)

docker build --tag kokizzu/sf1:latest .


# docker run -it kokizzu/sf1 sh

# docker push kokizzu/sf1:latest
