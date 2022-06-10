from telegram import Chat, InlineKeyboardMarkup, InlineKeyboardButton, Update, ParseMode, Message
from telegram.ext import Updater, CallbackContext
from google.protobuf.timestamp_pb2 import Timestamp

from Keyboard.keyboard_data import KeyboardData
from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2
from GrpcClient import broadcast_client
from Reminders import reminders


def send_broadcast_message(updater : Updater, message: str, chat_id: int, broadcast_recipient_id: int, urgency):
    print("sendBroadcastMessage", message, chat_id)
    keyboard_markup = InlineKeyboardMarkup(
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
    
    msg = updater.bot.send_message(chat_id=chat_id, text=formatted_message, parse_mode=ParseMode.HTML, reply_markup=keyboard_markup)
    setup_broadcast_reminder(updater, chat_id, msg, broadcast_recipient_id)

def acknowledge_broadcast(broadcast_recipient_id: int) -> bool:
    broadcast_recipient = operations_ecosys_pb2.BroadcastRecipient(
        broadcast_recipients_id = broadcast_recipient_id, 
        acknowledged = True, 
        rejected = False,
        last_replied = Timestamp().GetCurrentTime(),
    )
    updated = broadcast_client.update_broadcast_recipient(broadcast_recipient)
    return updated

def setup_broadcast_reminder(updater : Updater, chat_id: int, msg: Message, broadcast_recipient_id: int):
    updater.job_queue.run_once(
        callback=remind_broadcast, 
        when=reminders.BROADCAST_REMINDER_SECONDS, 
        context={
            reminders.MSG_KEY:msg,
            reminders.CHAT_ID_KEY: chat_id, 
            reminders.BROADCAST_REC_ID_KEY: broadcast_recipient_id,
            reminders.UPDATER_KEY: updater,
        },
    )

def remind_broadcast(context: CallbackContext):
    # print("remind_broadcast", context.job.context)

    # If already has a reply, do nothing
    bc_recipient = broadcast_client.get_broadcast_recipient(context.job.context[reminders.BROADCAST_REC_ID_KEY])

    # If the row in the db is already gone, it could have been deleted. Ignore
    if bc_recipient is None:
        return
    
    if bc_recipient.acknowledged or bc_recipient.rejected:
        print("Broadcast already acknowledged")
        return

    # Resend the same message + keyboard
    reminders.resend_message(context, setup_broadcast_reminder, context.job.context[reminders.BROADCAST_REC_ID_KEY])
    