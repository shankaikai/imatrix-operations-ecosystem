from . import broadcast_server as bc_server
from . import rostering_server as roster_server

from concurrent import futures
import logging
import math
import time

import grpc
from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2
from telegram.ext import Updater

GRPC_PORT = 9091

def serve(updater : Updater):
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  operations_ecosys_pb2_grpc.add_BroadcastServicesServicer_to_server(
      bc_server.BroadcastServicesServicer(updater), server)
  operations_ecosys_pb2_grpc.add_RosterServicesServicer_to_server(
      roster_server.RosterServicesServicer(updater), server)
  server.add_insecure_port('[::]:'+ str(GRPC_PORT))
  server.start()

  return server