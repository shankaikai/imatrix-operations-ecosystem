# imatrix-operations-ecosystem
Capstone Project

## Set up
1. Create a local .env file using the structure in `.env_skeleton`. Ask for help if necessary. 
2. Run `pip install -r requirements.txt`
## Frontend
### Protocol Buffer Generation
```
```

## Backend
### Protocol Buffer Generation
```
protoc --go_out=backend --go_opt=paths=source_relative --go-grpc_out=backend --go-grpc_opt=paths=source_relative proto/operations_ecosys.proto

protoc --go_out=backend --go_opt=paths=source_relative --go-grpc_out=backend --go-grpc_opt=paths=source_relative proto/iot_prototype.proto

protoc --go_out=backend --go_opt=paths=source_relative --go-grpc_out=backend --go-grpc_opt=paths=source_relative proto/http_webapp.proto
```
#### Protocol Buffer Mock Class Generation
```
mockgen capstone.operations_ecosystem/backend/proto AdminServicesClient > ../mock_proto/admin_grpc_mock.go
mockgen capstone.operations_ecosystem/backend/proto BroadcastServicesClient > ../mock_proto/broadcast_grpc_mock.go
mockgen capstone.operations_ecosystem/backend/proto RosterServicesClient > ../mock_proto/roster_grpc_mock.go
```

### Making HTTP work with GRPC
1) Manually install http.proto and annotations.proto from https://github.com/googleapis/googleapis/blob/master/google/api to wherever your protoc include folder is
--> Tip: find the location of a file that's already roughly in the same place (find / -type f -name timestamp.proto)
2) Import google/api/annotations.proto, but not google/api/http.proto
3) Install protoc-gen-go-grpc, NOT protoc-gen-go
--> go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
--> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
4) Install grpc-gateway (https://github.com/grpc-ecosystem/grpc-gateway)
--> go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
5) Generate the gateway protoc files and the go_grpc files
```
protoc -I ./proto --grpc-gateway_out=backend/proto --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative proto/http_webapp.proto
```

## Telegram Bot
### Protocol Buffer Generation
#### Install Type Checking Protobuf Library
```
pip3 install mypy-protobuf
```
#### Generation
With Type Checking
```
python3 -m grpc_tools.protoc -I proto --python_out=telebot/Protos --mypy_out=telebot/Protos --grpc_python_out=telebot/Protos --mypy_grpc_out=telebot/Protos proto/iot_prototype.proto

python3 -m grpc_tools.protoc -I proto --python_out=telebot/Protos --mypy_out=telebot/Protos --grpc_python_out=telebot/Protos --mypy_grpc_out=telebot/Protos proto/operations_ecosys.proto

python3 -m grpc_tools.protoc -I proto --python_out=telebot/Protos --mypy_out=telebot/Protos --grpc_python_out=telebot/Protos --mypy_grpc_out=telebot/Protos proto/http_webapp.proto
```

Without Type Checking
```
python -m grpc_tools.protoc -I proto --python_out=telebot/Protos --grpc_python_out=telebot/Protos proto/operations_ecosys.proto
```

## IoT Gate Prototype
### Protocol Buffer Generation
#### Install Type Checking Protobuf Library
```
pip3 install mypy-protobuf
```
#### Generation
With Type Checking
```
python -m grpc_tools.protoc -I proto --python_out=iot_control/Proto --mypy_out=iot_control/Proto --grpc_python_out=iot_control/Proto --mypy_grpc_out=iot_control/Proto proto/iot_prototype.proto
```

# Docker

To start all services, run `docker-compose up -d`

to restart certain services after making your changes, run `docker-compose restart CONTAINER_NAME`.

restart client: `docker-compose restart client`
restart backend: `docker-compose retart backend`
restart client + backend: `docker-compose restart client backend`
restart all services: `docker-compose restart`

to stop all services, run `docker-compose down`

to view logs (e.g. print statements), run `docker logs CONTAINER_NAME`
