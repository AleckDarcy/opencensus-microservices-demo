dep ensure

rm -rf vendor/google.golang.org/grpc
rm -rf vendor/github.com/AleckDarcy/reload
cp -r $GOPATH/src/github.com/AleckDarcy/reload vendor/github.com/AleckDarcy/reload
cp -r $GOPATH/src/github.com/AleckDarcy/reload/google.golang.org/grpc vendor/google.golang.org/grpc