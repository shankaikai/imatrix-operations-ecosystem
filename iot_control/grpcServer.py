from gate_control import gate_control_publisher
import iot_prototype_server

from concurrent import futures
import threading

import grpc
from Proto import iot_prototype_pb2_grpc

def serve(grpc_port : str):
  gate_lock = threading.Lock()
  fire_alarm_lock = threading.Lock()
  gate_publisher = gate_control_publisher.GateControlPublisher(gate_lock)
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  iot_prototype_pb2_grpc.add_IotControlPrototypeServiceServicer_to_server(
      iot_prototype_server.IotControlPrototypeServiceServicer(gate_lock, fire_alarm_lock, gate_publisher), server)
  server.add_insecure_port('[::]:'+ grpc_port)
  server.start()
  print("IoT Prototype Server started at port", grpc_port)
  server.wait_for_termination()