from __future__ import annotations
from telegram import Update, KeyboardButton, ReplyKeyboardMarkup, ReplyKeyboardRemove, WebAppInfo
from telegram.ext import CallbackContext
from telegram.files.photosize import PhotoSize

import os.path

from TelegramController import TelegramController, TelegramMenu, CMD_IDENTIFIER

class MainMenu(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "Main Menu"
        
    def handler(self, update:Update, context:CallbackContext):
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Attendance")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Reporting")],
            [KeyboardButton(text=CMD_IDENTIFIER+"SOS"), KeyboardButton(text=CMD_IDENTIFIER+"Help"),],
        ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text=self.name, reply_markup=cKeyboard)
        pass