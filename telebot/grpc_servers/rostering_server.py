from concurrent import futures

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2
from subscriptions.subscription_modules import rostering

from telegram.ext import Updater

class RosterServicesServicer(operations_ecosys_pb2_grpc.RosterServicesServicer):
    """Provides methods that implement functionality of roster server."""
    
    def __init__(self, updater : Updater):
        self.updater = updater

    def AddRoster(self, request, context):
        print("Received Request:", request)
        for roster in request.rosters:
            print("Roster: ", roster)
            for guard in roster.guard_assigned:
                guard_user = guard.guard_assigned.employee
                # print("guard user", guard_user)

                if guard_user.tele_chat_id == -1:
                    # print("Roster recipient has no telegram chat id. User id:", guard_user.user_id)
                    continue
                
                # print("Roster recipient has telegram chat id. User id:", guard_user.user_id)
                rostering.send_roster_message(self.updater, guard_user.tele_chat_id,
                    guard, roster,
                )
        res = operations_ecosys_pb2.Response(type=operations_ecosys_pb2.Response.ACK)
        return res