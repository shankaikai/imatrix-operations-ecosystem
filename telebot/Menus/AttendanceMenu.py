from __future__ import annotations
from telegram import Update, KeyboardButton, ReplyKeyboardMarkup, ReplyKeyboardRemove
from telegram.ext import CallbackContext
from telegram.files.photosize import PhotoSize

import os.path
import os



from TelegramController import TelegramController, TelegramMenu

class AttendanceMenu(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "Attendance Menu"
        self.attendance_save_path = os.getenv('SAVE_PATH')
    def handler(self, update:Update, context:CallbackContext):
        cKeyboardVals = [
            [KeyboardButton(text="Cancel")]
        ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text="Please send your photo.", reply_markup=cKeyboard)
        pass
    def attachmentHandler(self, update: Update, context: CallbackContext):
        attachment = update.effective_message.effective_attachment
        fileType = type(attachment)
        # Only photos should appear in a list, other attachment
        # types should be singular
        if fileType == type([]) and len(attachment) > 0:
            fileType = type(attachment[1])
        if fileType == PhotoSize:
            fileName = update.effective_user.username + ".jpg"
            photo = attachment[1].get_file()
            photo.download(os.path.join(self.attendance_save_path, fileName))
            print("Attendance photo for " + update.effective_user.username + " obtained.")
            context.bot.send_message(chat_id=update.effective_chat.id, text="Your attendance has been submitted.")
            self.TController.backHandler(update, context)
            return