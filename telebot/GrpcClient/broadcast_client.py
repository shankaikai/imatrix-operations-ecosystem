from __future__ import print_function

from GrpcClient import utils

import grpc

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2

def get_broadcast_recipient(broadcast_recipient_id: int) -> operations_ecosys_pb2.BroadcastRecipient:
    stub = get_broadcast_stub()
    broadcast_filter = operations_ecosys_pb2.BroadcastFilter(
        field=operations_ecosys_pb2.BroadcastFilter.BROADCAST_RECIPIENT_TABLE_ID,
        comparisons = operations_ecosys_pb2.Filter(
            comparison=operations_ecosys_pb2.Filter.EQUAL,
            value=str(broadcast_recipient_id)
        )
    )
    broadcastQuery = operations_ecosys_pb2.BroadcastQuery(
        filters = [broadcast_filter],
        limit = 1,
    )
    broadcastResponses = stub.FindBroadcasts(broadcastQuery)
    broadcastRes = None

    # There should only be at most one response because the limit was 1
    for res in broadcastResponses:
        broadcastRes = res
        break

    if res is None:
        print("No broadcasts returned")
        return None
    
    print("get_broadcast_recipient", broadcastRes.response)

    if broadcastRes.broadcast is None:
        return None

    if len(broadcastRes.broadcast.recipients) == 0:
        return None
    if len(broadcastRes.broadcast.recipients[0].recipient) == 0:
        return None
        
    return broadcastRes.broadcast.recipients[0].recipient[0]


def update_broadcast_recipient(broadcast_recipient) -> bool:
    stub = get_broadcast_stub()
    res = stub.UpdateBroadcastRecipient(broadcast_recipient)
    print("update_broadcast_recipient", res)
    return res.type == operations_ecosys_pb2.Response.ACK


def get_broadcast_stub() -> operations_ecosys_pb2_grpc.BroadcastServicesStub:
    channel = grpc.insecure_channel('{}:{}'.format(utils.WEB_SERVER_ADDR, utils.WEB_SERVER_PORT))
    stub = operations_ecosys_pb2_grpc.BroadcastServicesStub(channel)
    return stub
