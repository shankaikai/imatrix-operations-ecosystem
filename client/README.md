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
