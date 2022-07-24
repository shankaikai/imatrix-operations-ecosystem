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
        # Report Web App stuff
        wa_url = os.getenv('WEBAPP_URL')
        params = "?twan=" + user_client.get_webapp_nonce(self.TController.user)
        params += "&user_id" + update.effective_user.id
        reportWAInfo = WebAppInfo(wa_url + params)
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Attendance")],
            # [KeyboardButton(text=CMD_IDENTIFIER+"Reporting")],
            [KeyboardButton(text="Reporting", web_app=reportWAInfo)],
            [KeyboardButton(text=CMD_IDENTIFIER+"SOS"), KeyboardButton(text=CMD_IDENTIFIER+"Help"),],
        ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text=self.name, reply_markup=cKeyboard)
        pass