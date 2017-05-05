FROM alpine:3.5
RUN apk add --no-cache ca-certificates
COPY cmd/bolt-store/bolt-store /bin/bolt-store
COPY cmd/elwin/elwin /bin/elwin
COPY cmd/elwin-grpc-gateway/elwin-grpc-gateway /bin/elwin-grpc-gateway
COPY cmd/gen/gen /bin/gen
COPY cmd/houston/houston /bin/houston
COPY cmd/json-gateway/json-gateway /bin/json-gateway
COPY cmd/mongo-store/mongo-store /bin/mongo-store
CMD echo "You should supply the command that you want to run. [bolt-store, elwin, elwin-grpc-gateway, gen, houston, json-gateway, mongo-store]"; exit 1
