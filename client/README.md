# Client things to take note

Run envoy so that the frontend can communicate with the go server:

```bash
cd ../envoy
docker-compose up
```

Run the development server:

```bash
npm run dev
```

To update the protobuf client stub:

```bash
cd ../proto
sh gen-ts-proto.sh
```

Make sure that you have protobuf installed before trying to update the protobuf client stubs. Go to this [link](http://google.github.io/proto-lens/installing-protoc.html) to install protoc and this [link](https://github.com/grpc/grpc-web#code-generator-plugin) to install the protoc-gen-grpc-web plugin.

To run the Cypress e2e test:

```bash
npx cypress open
```
