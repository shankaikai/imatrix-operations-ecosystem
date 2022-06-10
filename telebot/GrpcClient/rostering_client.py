from __future__ import print_function

from GrpcClient import utils

import grpc

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2


def get_roster_assignment(roster_assignment_id: int) -> operations_ecosys_pb2.BroadcastRecipient:
    print("get_roster_assignment", roster_assignment_id)
    stub = get_roster_stub()
    roster_filter = operations_ecosys_pb2.RosterFilter(
        field=operations_ecosys_pb2.RosterFilter.ROSTER_ASSIGNMENT_ID,
        comparisons = operations_ecosys_pb2.Filter(
            comparison=operations_ecosys_pb2.Filter.EQUAL,
            value=str(roster_assignment_id)
        )
    )
    roster_filter = operations_ecosys_pb2.RosterQuery(
        filters = [roster_filter],
        limit = 1,
    )
    roster_assgn_responses = stub.FindRosterAssignments(roster_filter)
    roster_res = None

    # There should only be at most one response because the limit was 1
    for res in roster_assgn_responses:
        roster_res = res
        break

    if roster_res is None:
        print("No roster assignments returned")
        return None
    
    print(roster_res.response)
    if roster_res.roster_assignment is None:
        return None
        
    return roster_res.roster_assignment


def update_rostering_assignment(roster_assignment: operations_ecosys_pb2.RosterAssignement) -> bool:
    stub = get_roster_stub()
    res = stub.UpdateRosterAssignment(roster_assignment)
    print("update_rostering_recipient", res)
    return res.type == operations_ecosys_pb2.Response.ACK


def get_roster_stub() -> operations_ecosys_pb2_grpc.RosterServicesStub:
    channel = grpc.insecure_channel('{}:{}'.format(utils.WEB_SERVER_ADDR, utils.WEB_SERVER_PORT))
    stub = operations_ecosys_pb2_grpc.RosterServicesStub(channel)
    return stub
