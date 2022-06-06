from concurrent import futures

import grpc
from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2


class BroadcastServicesServicer(operations_ecosys_pb2_grpc.BroadcastServicesServicer):
    """Provides methods that implement functionality of broadcasting server."""
    
    # TODO: add broadcast method
    def AddBroadcast(self, request, context):
        print("Received Request:", request)
        res = operations_ecosys_pb2.Response(operations_ecosys_pb2.Response.ACK)
        return res