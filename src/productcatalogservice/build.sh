TMPPath=$GOPATH

export GOPATH=$(pwd)/../..

echo $GOPATH

dep ensure

export GOPATH=$TMPPath
echo $GOPATH

rm -rf vendor/github.com/AleckDarcy/reload
cp -r $GOPATH/src/github.com/AleckDarcy/reload/google.golang.org/grpc vendor/google.golang.org/grpc