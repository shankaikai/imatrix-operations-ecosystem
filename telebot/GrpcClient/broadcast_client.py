from __future__ import print_function

from GrpcClient import utils

import grpc

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2

def update_broadcast_recipient(broadcast_recipient) -> bool:
    stub = get_broadcast_stub()
    res = stub.UpdateBroadcastRecipient(broadcast_recipient)
    print("update_broadcast_recipient", res)
    return res.type == operations_ecosys_pb2.Response.ACK


def get_broadcast_stub() -> operations_ecosys_pb2_grpc.BroadcastServicesStub:
    channel = grpc.insecure_channel('{}:{}'.format(utils.WEB_SERVER_ADDR, utils.WEB_SERVER_PORT))
    stub = operations_ecosys_pb2_grpc.BroadcastServicesStub(channel)
    return stub
