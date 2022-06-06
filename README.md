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
```

## Telegram Bot
### Protocol Buffer Generation
```
python -m grpc_tools.protoc -I proto --python_out=telebot/Protos --grpc_python_out=telebot/Protos proto/operations_ecosys.proto
```