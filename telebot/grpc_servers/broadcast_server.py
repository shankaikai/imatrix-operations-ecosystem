from concurrent import futures

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2
from subscriptions.subscription_modules import broadcast as bc

from telegram.ext import Updater

class BroadcastServicesServicer(operations_ecosys_pb2_grpc.BroadcastServicesServicer):
    """Provides methods that implement functionality of broadcasting server."""
    
    def __init__(self, updater : Updater):
        self.updater = updater

    def AddBroadcast(self, request, context):
        print("Received Request:", request)
        # print stuff for testing
        print("urgency level: ", request.urgency)
        for aifsRecipient in request.recipients:
            for broadcastRecipient in aifsRecipient.recipient:
                # print(broadcastRecipient.recipient)
                
                if broadcastRecipient.recipient.tele_user_id == -1:
                    # print("Broadcast recipient has no telegram user id. User id:", broadcastRecipient.recipient.user_id)
                    continue
                
                # print("Broadcast recipient has telegram user id. User id:", broadcastRecipient.recipient.user_id)
                bc.send_broadcast_message(self.updater, request.content, #right click send_broadcast_message to go to function
                    broadcastRecipient.recipient.tele_user_id,
                    broadcastRecipient.broadcast_recipients_id, # update parameters here if needed
                    request.urgency
                )
        res = operations_ecosys_pb2.Response(type=operations_ecosys_pb2.Response.ACK)
        return res