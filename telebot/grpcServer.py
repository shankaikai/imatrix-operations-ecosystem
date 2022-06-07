from Broadcast import broadcast_grpc_server as bc_server

from concurrent import futures
import logging
import math
import time

import grpc
from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2

GRPC_PORT = 9091

def serve():
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  operations_ecosys_pb2_grpc.add_BroadcastServicesServicer_to_server(
      bc_server.BroadcastServicesServicer(), server)
  server.add_insecure_port('[::]:'+ str(GRPC_PORT))
  server.start()
  server.wait_for_termination()

