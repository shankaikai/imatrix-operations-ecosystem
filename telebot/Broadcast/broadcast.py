from telegram import Chat, InlineKeyboardMarkup, InlineKeyboardButton

async def sendBroadcastMessage(message: str, chat_id: str):
    chat = Chat(id=chat_id)
    keyboard_markup = InlineKeyboardMarkup(
            [
                InlineKeyboardButton(
                    text="Acknowledge",
                    callback_data="ack"
                )
            ]
    )
    await chat.send_message(text=message,reply_markup=keyboard_markup)
    
def query_yes(update, context):
    print(update.callback_query.data)
