import argparse
from telnetlib import GA

import grpc

from Proto import iot_prototype_pb2_grpc, iot_prototype_pb2

GATE_ID = 1
FIRE_ALARM_ID = 1

def gate_state_iterator():
    print("0 for close, 1 for open")
    yield iot_prototype_pb2.GateState(
            id=GATE_ID, 
            state=iot_prototype_pb2.GateState.INITIAL
        )
    while True:
        user_input = input()
        if user_input == "0":
            state = iot_prototype_pb2.GateState(
                id=GATE_ID, 
                state=iot_prototype_pb2.GateState.CLOSED
            )
            yield state
        else:
            state = iot_prototype_pb2.GateState(
                id=GATE_ID, 
                state=iot_prototype_pb2.GateState.OPEN
            )
            yield state

def set_gate_states(addr, port):
    stub = get_iot_stub(addr, port)
    for gate_state in stub.SetGateState(gate_state_iterator()):
        print(gate_state)
 
def get_gate_states(addr, port):
    stub = get_iot_stub(addr, port)
    # itr = gate_state_iterator()
    responses =  stub.GetGateState(iot_prototype_pb2.Gate(id=GATE_ID))
    for gate_state in responses:
        print(gate_state)


def get_iot_stub(addr, port) -> iot_prototype_pb2_grpc.IotControlPrototypeServiceStub:
    channel = grpc.insecure_channel('{}:{}'.format(addr, port))
    stub = iot_prototype_pb2_grpc.IotControlPrototypeServiceStub(channel)
    return stub


def parse_arguments():
    parser = argparse.ArgumentParser(description='IoT Gate Control Prototype')
    parser.add_argument('--address', type=str, default="localhost",
                        help='Address of the IoT Gate Server')
    parser.add_argument('--port', type=str, default="9099",
                        help='Port of the IoT Gate Server')
    args = parser.parse_args()
    return args

def main():
    args = parse_arguments()
    # set_gate_states(args.address, args.port)
    get_gate_states(args.address, args.port)

if __name__ == "__main__":
    main()