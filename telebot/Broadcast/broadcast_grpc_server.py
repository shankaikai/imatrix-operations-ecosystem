from concurrent import futures

import grpc
from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2
from Broadcast import broadcast as bc

class BroadcastServicesServicer(operations_ecosys_pb2_grpc.BroadcastServicesServicer):
    """Provides methods that implement functionality of broadcasting server."""
    
    # TODO: add broadcast method
    def AddBroadcast(self, request, context):
        # print("Received Request:", request)
        # TODO send to telegram user
        for recipient in request.recipients:
            print(recipient)
            # bc.sendBroadcastMessage(request.content, recipient.telegram_handle)
        res = operations_ecosys_pb2.Response(type=operations_ecosys_pb2.Response.ACK)
        return res