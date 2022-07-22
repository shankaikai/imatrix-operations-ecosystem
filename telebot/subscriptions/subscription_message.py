from __future__ import annotations
from telegram import Chat, InlineKeyboardMarkup, InlineKeyboardButton, Update
from telegram.ext import Updater, ContextTypes
from . import subscription_modules

class SubscriptionMessage:
    DELIMITER = "-"
    IDENTIFIER = "SM:"

    def __init__(self, feature : str, pb_msg_id : int, chat_id : int) -> None:
        self.feature = feature
        self.pb_msg_id = pb_msg_id
        self.chat_id = chat_id
    
    @staticmethod
    def deserialise(serialised_data:str) -> SubscriptionMessage:
        split_data = serialised_data.split(SubscriptionMessage.DELIMITER)
        if len(split_data) < 3:
            return SubscriptionMessage("", -1, -1)
        feature = split_data[0]
        pb_msg_id = int(split_data[1])
        chat_id = int(split_data[2])

        return SubscriptionMessage(feature, pb_msg_id, chat_id)

    def __str__(self) -> str:
        return "{}{}{}{}{}{}".format(self.IDENTIFIER, self.feature, self.DELIMITER, self.pb_msg_id, self.DELIMITER, self.chat_id)



def callbackqueryHandler(update: Update, context: ContextTypes) -> None:
    """Parses the CallbackQuery and updates the message text."""
    query = update.callback_query

    # CallbackQueries need to be answered, even if no notification to the user is needed
    # Some clients may have trouble otherwise. See https://core.telegram.org/bots/api#callbackquery
    query.answer()

    sub_msg = SubscriptionMessage.deserialise(query.data)
    comment = "No subscription handle handles " + sub_msg.feature

    for mod in subscription_modules:
        if mod.IDENTIFIER in sub_msg.feature:
            isSuccessful, comment = mod.handle_sub_message(sub_msg)
            if isSuccessful:
                query.edit_message_text(f"{query.message.text}\n\n" + comment)
                break
            print("Error handling subscription message:", comment)
            print("Subscription message was:", query.to_dict())
