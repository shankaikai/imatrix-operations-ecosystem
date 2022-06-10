from telegram import Chat, InlineKeyboardMarkup, InlineKeyboardButton, Update
from telegram.ext import Updater, ContextTypes
from google.protobuf.timestamp_pb2 import Timestamp

from Keyboard.keyboard_data import KeyboardData
from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2
from GrpcClient import broadcast_client

def send_roster_message(updater : Updater, message: str, chat_id: int, roster_assignment: operations_ecosys_pb2.RosterAssignement):
    print("sendBroadcastMessage", message, chat_id)
    keyboard_markup = InlineKeyboardMarkup(
            [[
                InlineKeyboardButton(
                    text="Acknowledge",
                    callback_data=str(KeyboardData(KeyboardData.ROSTERING_FEATURE, broadcast_recipient_id, chat_id))
                )
            ]]
    )

    formated_message = 23432

    updater.bot.send_message(chat_id=chat_id, text=message, reply_markup=keyboard_markup)

def acknowledge_broadcast(broadcast_recipient_id: int) -> bool:
    broadcast_recipient = operations_ecosys_pb2.BroadcastRecipient(
        broadcast_recipients_id = broadcast_recipient_id, 
        acknowledged = True, 
        rejected = False,
        last_replied = Timestamp().GetCurrentTime(),
    )
    updated = broadcast_client.update_broadcast_recipient(broadcast_recipient)
    return updated


# broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_ACK, strconv.FormatBool(bRec.Acknowledged), false))
# 	broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_REJECTION, strconv.FormatBool(bRec.Rejected), false))

# if bRec.LastReplied != nil {
#     broadcastTableFields = append(broadcastTableFields, formatFieldEqVal(BC_REC_DB_LAST_REPLIED, bRec.LastReplied.AsTime().Format(common.DATETIME_FORMAT), true))
# }