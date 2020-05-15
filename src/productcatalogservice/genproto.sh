#!/bin/bash -eu
#
# Copyright 2018 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#!/bin/bash -e

PATH=$PATH:$GOPATH/bin
#protodir=../../pb
protodir=genproto

#protoc --go_out=plugins=grpc:genproto -I $protodir $protodir/demo.proto

protoc --go_out=plugins=grpc:genproto -I $protodir -I./vendor $protodir/demo.proto

# replace some strings
sed -i '' 's/context \"context\"/context \"golang.org\/x\/net\/context\"/g' $protodir/demo.pb.go
sed -i '' 's/const _ = proto.ProtoPackageIsVersion3/const _ = proto.ProtoPackageIsVersion2/g' $protodir/demo.pb.go
sed -i '' 's/const _ = grpc.SupportPackageIsVersion6/const _ = grpc.SupportPackageIsVersion4/g' $protodir/demo.pb.go
sed -i '' 's/grpc.ClientConnInterface/*grpc.ClientConn/g' $protodir/demo.pb.go

