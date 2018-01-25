GOOS=linux GOARCH=386 \
go build -ldflags "-s" -a -installsuffix cgo -o grpcserver server/server.go