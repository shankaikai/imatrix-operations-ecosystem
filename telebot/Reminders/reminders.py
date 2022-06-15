from telegram.ext import Updater, CallbackContext
from telegram import Chat, InlineKeyboardMarkup, InlineKeyboardButton, Update, ParseMode, Message

# Timings
BROADCAST_REMINDER_SECONDS = 5
ROSTER_REMINDER_SECONDS = 5# TODO 30

# Job context keys
# Common
CHAT_ID_KEY = "chat_id"
UPDATER_KEY = "updater"
MSG_KEY = "msg"

# DB Primary Keys
BROADCAST_REC_ID_KEY = "broadcast_recipient_id"
ROSTER_ASGN_ID_KEY = "roster_assignment_id"

# Resend the same message + keyboard
# Note: set_up_reminder_func(updater : Updater, chat_id: int, msg: Message, db_primary_key: int):
def resend_message(context: CallbackContext, set_up_reminder_func, db_pk: int):
    prev_msg = context.job.context[MSG_KEY]
    updater = context.job.context[UPDATER_KEY]

    msg = updater.bot.send_message(
        chat_id=context.job.context[CHAT_ID_KEY], 
        text=prev_msg.text, 
        parse_mode=ParseMode.HTML, 
        reply_markup=prev_msg.reply_markup, 
    )

    prev_msg.edit_text(text=prev_msg.text + "\n\nResent")
    set_up_reminder_func(updater, context.job.context[CHAT_ID_KEY], msg, db_pk)