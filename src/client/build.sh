TMPPath=$GOPATH

export GOPATH=$(pwd)/../..

echo $GOPATH

dep ensure -update

export GOPATH=$TMPPath
echo $GOPATH