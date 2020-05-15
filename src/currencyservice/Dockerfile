FROM node:8
RUN GRPC_HEALTH_PROBE_VERSION=v0.1.0-alpha.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
WORKDIR /usr/src/app

COPY package*.json ./
RUN npm install --only=production

RUN npm install sleep
RUN npm install uuid

COPY . .

#RUN rm -f node_modules/grpc/index.js
#COPY reload/index.js node_modules/grpc

RUN rm -f node_modules/grpc/src/server.js
COPY reload/grpc/src/server.js node_modules/grpc/src

#RUN rm -f node_modules/@grpc/proto-loader/build/src/index.js
#COPY reload/@grpc/proto-loader/build/src/server.js node_modules/@grpc/proto-loader/build/src

EXPOSE 7000
CMD [ "node", "server.js" ]
