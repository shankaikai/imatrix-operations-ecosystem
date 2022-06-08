from __future__ import print_function

from GrpcClient import utils

import grpc

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2

def get_users(user_query) -> list[operations_ecosys_pb2.User]:
    stub = get_admin_stub()
    userResponses = stub.FindUsers(user_query)
    
    users = []

    for userRes in userResponses:
        if userRes.response.type == operations_ecosys_pb2.Response.ACK:
            users.append(userRes.user)
        else:
            print("get_users", userRes)

    return users

def update_user(user) -> bool:
    stub = get_admin_stub()
    res = stub.UpdateUser(user)
    return res.type == operations_ecosys_pb2.Response.ACK


def get_admin_stub() -> operations_ecosys_pb2_grpc.AdminServicesStub:
    channel = grpc.insecure_channel('{}:{}'.format(utils.WEB_SERVER_ADDR, utils.WEB_SERVER_PORT))
    stub = operations_ecosys_pb2_grpc.AdminServicesStub(channel)
    return stub
