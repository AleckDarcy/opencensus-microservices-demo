FROM golang:1.10-alpine as builder
RUN apk add --no-cache ca-certificates git && \
      wget -qO/go/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && \
      chmod +x /go/bin/dep

ENV PROJECT github.com/AleckDarcy/opencensus-microservices-demo/src/client
WORKDIR /go/src/$PROJECT

# restore dependencies
#COPY Gopkg.* ./
#RUN dep ensure --vendor-only -v
COPY . .
RUN go install .

FROM alpine as release
RUN apk add --no-cache ca-certificates \
    busybox-extras net-tools bind-tools
WORKDIR /client
COPY --from=builder /go/bin/client /client/server
COPY ./templates ./templates
EXPOSE 8081
ENTRYPOINT ["/client/server"]
