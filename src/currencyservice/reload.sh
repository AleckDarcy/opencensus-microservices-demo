# release

rm -f reload/grpc/src/server.js
cp $GOPATH/src/github.com/AleckDarcy/reload/reload/node/grpc/src/server.js reload/grpc/src/server.js

# debug

#rm node_modules/grpc/src/server.js
#ln -s $GOPATH/src/github.com/AleckDarcy/reload/reload/node/grpc/src/server.js node_modules/grpc/src/server.js
#
#rm node_modules/@grpc/proto-loader/build/src/index.js
#ln -s $GOPATH/src/github.com/AleckDarcy/reload/reload/node/@grpc/proto-loader/build/src/index.js node_modules/@grpc/proto-loader/build/src/index.js
