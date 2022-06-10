from __future__ import print_function

from GrpcClient import utils

import grpc

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2

def update_rostering_assignment(roster_assignment: operations_ecosys_pb2.RosterAssignement) -> bool:
    stub = get_roster_stub()
    res = stub.UpdateRosterAssignment(roster_assignment)
    print("update_rostering_recipient", res)
    return res.type == operations_ecosys_pb2.Response.ACK


def get_roster_stub() -> operations_ecosys_pb2_grpc.RosterServicesStub:
    channel = grpc.insecure_channel('{}:{}'.format(utils.WEB_SERVER_ADDR, utils.WEB_SERVER_PORT))
    stub = operations_ecosys_pb2_grpc.RosterServicesStub(channel)
    return stub
