from __future__ import annotations
from telegram import Update, KeyboardButton, ReplyKeyboardMarkup, ReplyKeyboardRemove, WebAppInfo, InlineKeyboardMarkup, InlineKeyboardButton
from telegram.ext import CallbackContext
from telegram.files.photosize import PhotoSize

import os.path

from grpc_clients import user_client

from TelegramController import TelegramController, TelegramMenu, CMD_IDENTIFIER

class MainMenu(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "Main Menu"
        
    def handler(self, update:Update, context:CallbackContext):
        # It is not a good idea to use web_app keyboards in the current implementation
        # as to do so would pull nonces excessively. 
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"My Assignments")],
            #[KeyboardButton(text="My Assignments", web_app=assignmentsWAInfo)],
            [KeyboardButton(text=CMD_IDENTIFIER+"Make a Report")],
            #[KeyboardButton(text="Make a Report", web_app=reportWAInfo)],
            [KeyboardButton(text=CMD_IDENTIFIER+"Others"),]
        ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text=self.name, reply_markup=cKeyboard)
        pass

class MainMenu_Attendance_WA(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "MainMenu_Attendance_WA"
    def handler(self, update:Update, context:CallbackContext):
        # Assignments Web App Stuff
        wa_url = os.getenv('WEBAPP_ASSIGNMENTS_URL')
        params = "?twan=" + user_client.get_webapp_nonce(self.TController.user.oes_user)
        params += "&user_id=" + str(update.effective_user.id)
        assignmentsWAInfo = WebAppInfo(wa_url + params)
        inlineKBs = InlineKeyboardMarkup([[InlineKeyboardButton(text="View", web_app=assignmentsWAInfo)]])
        context.bot.send_message(chat_id=update.effective_chat.id, text="Press 'View' to see your assignments.", reply_markup=inlineKBs)
        self.TController.CurrentMenu = self.parent
        pass

class MainMenu_Report_WA(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "MainMenu_Attendance_WA"
    def handler(self, update:Update, context:CallbackContext):
        # Report Web App Stuff
        wa_url = os.getenv('WEBAPP_REPORT_URL')
        params = "?twan=" + user_client.get_webapp_nonce(self.TController.user.oes_user)
        params += "&user_id=" + str(update.effective_user.id)
        reportWAInfo = WebAppInfo(wa_url + params)
        inlineKBs = InlineKeyboardMarkup([[InlineKeyboardButton(text="Go", web_app=reportWAInfo)]])
        context.bot.send_message(chat_id=update.effective_chat.id, text="Press 'Go' to make a report.", reply_markup=inlineKBs)
        self.TController.CurrentMenu = self.parent
        pass