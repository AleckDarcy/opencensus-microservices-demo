rm -rf vendor/google.golang.org/grpc

# release
#ln -s vendor/github.com/AleckDarcy/reload/google.golang.org/grpc vendor/google.golang.org/grpc


# debug
rm -rf vendor/github.com/AleckDarcy/reload
ln -s $GOPATH/src/github.com/AleckDarcy/reload vendor/github.com/AleckDarcy/reload
ln -s $GOPATH/src/github.com/AleckDarcy/reload/google.golang.org/grpc vendor/google.golang.org/grpc
