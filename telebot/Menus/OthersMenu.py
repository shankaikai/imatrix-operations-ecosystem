from __future__ import annotations
from telegram import Update, KeyboardButton, ReplyKeyboardMarkup, ReplyKeyboardRemove
from telegram.ext import CallbackContext
from telegram.files.photosize import PhotoSize

import os.path

from TelegramController import TelegramController, TelegramMenu, CMD_IDENTIFIER

from Protos import operations_ecosys_pb2


class OthersMenu(TelegramMenu):
    def __init__(self, parent=None, name="", triggerWords=[]):
        super().__init__(parent, name, triggerWords)
        self.name = "Others Menu"
    def handler(self, update:Update, context:CallbackContext):
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Quick Introduction")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Admin")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
        ]
        if not (self.TController.user.user_type == operations_ecosys_pb2.User.UserType.MANAGER \
        or self.TController.user.user_type == operations_ecosys_pb2.User.UserType.CONTROLLER):
            cKeyboardVals.pop(1)
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text=self.name, reply_markup=cKeyboard)

class OthersMenu_QuickIntro_Text(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "OthersMenu_QuickIntro_Text"
    def handler(self, update:Update, context:CallbackContext):
        quickIntroText = """
        This is the IM Bot, a Telegram Bot by iMatrix.
        'My  Assignments' shows you a list of your upcoming duties.
        'Make a Report' allows you to submit a report.
        For more inquiries, ask your system administrator.
        """
        context.bot.send_message(chat_id=update.effective_chat.id, text=quickIntroText)
        self.TController.CurrentMenu = self.parent
        pass