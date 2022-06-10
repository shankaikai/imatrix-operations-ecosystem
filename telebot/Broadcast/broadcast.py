from telegram import Chat, InlineKeyboardMarkup, InlineKeyboardButton, Update, ParseMode
from telegram.ext import Updater, ContextTypes
from google.protobuf.timestamp_pb2 import Timestamp

from Keyboard.keyboard_data import KeyboardData
from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2
from GrpcClient import broadcast_client

def send_broadcast_message(updater : Updater, message: str, chat_id: int, broadcast_recipient_id: int, urgency): # can add urgency parameter here
    print("sendBroadcastMessage", message, chat_id)
    keyboard_markup = InlineKeyboardMarkup( #ignore
            [[
                InlineKeyboardButton(
                    text="Acknowledge",
                    callback_data=str(KeyboardData(KeyboardData.BROADCAST_FEATURE, broadcast_recipient_id, chat_id))
                )
            ]]
    )

    if urgency==2:
        formatted_message = "🔴 <b>High Urgency</b>\n" + message
    elif urgency==1:
        formatted_message = "🟠 <b>Medium Urgency</b>\n" + message
    elif urgency==0:
        formatted_message = "🟢 <b>Low Urgency</b>\n" + message
    else:
        formatted_message = message

    updater.bot.send_message(chat_id=chat_id, text=formatted_message, parse_mode=ParseMode.HTML, reply_markup=keyboard_markup) #keep

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