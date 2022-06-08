from concurrent import futures

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2
from Broadcast import broadcast as bc

from telegram.ext import Updater

class BroadcastServicesServicer(operations_ecosys_pb2_grpc.BroadcastServicesServicer):
    """Provides methods that implement functionality of broadcasting server."""
    
    def __init__(self, updater : Updater):
        self.updater = updater

    def AddBroadcast(self, request, context):
        print("Received Request:", request)
        for aifsRecipient in request.recipients:
            for broadcastRecipient in aifsRecipient.recipient:
                print(broadcastRecipient.recipient)
                
                if broadcastRecipient.recipient.tele_chat_id == -1:
                    print("Broadcast recipient has no telegram chat id. User id:", broadcastRecipient.recipient.user_id)
                    continue
                
                print("Broadcast recipient has telegram chat id. User id:", broadcastRecipient.recipient.user_id)
                bc.send_broadcast_message(self.updater, request.content,
                    broadcastRecipient.recipient.tele_chat_id,
                    broadcastRecipient.broadcast_recipients_id,
                )
        res = operations_ecosys_pb2.Response(type=operations_ecosys_pb2.Response.ACK)
        return res