FROM node:8
RUN GRPC_HEALTH_PROBE_VERSION=v0.1.0-alpha.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

WORKDIR /usr/src/app

COPY . .

RUN npm install --only=production && \
    npm install sleep && \
    npm install uuid && \
    npm install log4js && \
    rm -f node_modules/grpc/src/server.js

COPY reload/grpc/src/server.js node_modules/grpc/src

EXPOSE 7000
CMD [ "node", "server.js" ]
