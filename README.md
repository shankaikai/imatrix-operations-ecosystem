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
```
#### Protocol Buffer Mock Class Generation
```
mockgen capstone.operations_ecosystem/backend/proto AdminServicesClient > ../mock_proto/admin_grpc_mock.go
mockgen capstone.operations_ecosystem/backend/proto BroadcastServicesClient > ../mock_proto/broadcast_grpc_mock.go
mockgen capstone.operations_ecosystem/backend/proto RosterServicesClient > ../mock_proto/roster_grpc_mock.go
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
python -m grpc_tools.protoc -I proto --python_out=telebot/Protos --mypy_out=telebot/Protos --grpc_python_out=telebot/Protos --mypy_grpc_out=telebot/Protos proto/operations_ecosys.proto
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
