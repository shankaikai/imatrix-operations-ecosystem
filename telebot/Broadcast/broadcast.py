from telegram import Chat, InlineKeyboardMarkup, InlineKeyboardButton
from telegram.ext import Updater

# TODO call back for ack button
def sendBroadcastMessage(updater : Updater, message: str, chat_id: int):
    print("sendBroadcastMessage", message, chat_id)
    keyboard_markup = InlineKeyboardMarkup(
            [[
                InlineKeyboardButton(
                    text="Acknowledge",
                    callback_data="ack"
                )
            ]]
    )

    updater.bot.send_message(chat_id=chat_id, text=message, reply_markup=keyboard_markup)
    
def query_yes(update, context):
    print(update.callback_query.data)
