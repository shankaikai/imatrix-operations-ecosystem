from telegram import Chat, InlineKeyboardMarkup, InlineKeyboardButton, Update
from telegram.ext import Updater, ContextTypes
import Broadcast as bc
import Rostering as rostering

class KeyboardData: pass

class KeyboardData:
    BROADCAST_FEATURE = "broadcast"
    ROSTERING_ACCEPT_FEATURE = "rostering_accept"
    ROSTERING_REJECT_FEATURE = "rostering_reject"

    DELIMITER = "-"

    def __init__(self, feature : str, pb_msg_id : int, chat_id : int) -> None:
        self.feature = feature
        self.pb_msg_id = pb_msg_id
        self.chat_id = chat_id
    
    @staticmethod
    def get_keyboard_data(serialised_keyboard_data) -> KeyboardData:
        split_data = serialised_keyboard_data.split(KeyboardData.DELIMITER)
        if len(split_data) < 3:
            return KeyboardData("", -1, -1)
        feature = split_data[0]
        pb_msg_id = int(split_data[1])
        chat_id = int(split_data[2])

        return KeyboardData(feature, pb_msg_id, chat_id)

    def __str__(self) -> str:
        return "{}{}{}{}{}".format(self.feature, self.DELIMITER, self.pb_msg_id, self.DELIMITER, self.chat_id)



def keyboardCallback(update: Update, context: ContextTypes) -> None:
    """Parses the CallbackQuery and updates the message text."""
    query = update.callback_query

    # CallbackQueries need to be answered, even if no notification to the user is needed
    # Some clients may have trouble otherwise. See https://core.telegram.org/bots/api#callbackquery
    query.answer()

    keyboard_data = KeyboardData.get_keyboard_data(query.data)

    if keyboard_data.feature == KeyboardData.BROADCAST_FEATURE:
        print("found broadcast keyboard", query.to_dict())
        updated = bc.broadcast.acknowledge_broadcast(keyboard_data.pb_msg_id)
        if updated:
            query.edit_message_text(f"{query.message.text}\n\nAcknowledged")
        else:
            print("could not update broadcast message")
    elif keyboard_data.feature == KeyboardData.ROSTERING_ACCEPT_FEATURE:
        print("found roster accept keyboard", query.to_dict())
        updated = rostering.rostering.acknowledge_roster(keyboard_data.pb_msg_id)
        if updated:
            query.edit_message_text(f"{query.message.text}\n\nAcknowledged")
        else:
            print("could not update roster accept message")
    elif keyboard_data.feature == KeyboardData.ROSTERING_REJECT_FEATURE:
        print("found roster reject keyboard", query.to_dict())
        updated = rostering.rostering.reject_roster(keyboard_data.pb_msg_id)
        if updated:
            query.edit_message_text(f"{query.message.text}\n\nRejected")
        else:
            print("could not update roster reject message")