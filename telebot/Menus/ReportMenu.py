from __future__ import annotations
from typing import List
from telegram import InlineKeyboardButton, InlineKeyboardMarkup, Update, KeyboardButton, ReplyKeyboardMarkup, ReplyKeyboardRemove, WebAppInfo
from telegram.ext import CallbackContext
from telegram.files.photosize import PhotoSize

import os.path

from grpc_clients import user_client
from TelegramController import TelegramController, TelegramMenu, CMD_IDENTIFIER

class ReportMenu(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "Report Menu"
    def handler(self, update:Update, context:CallbackContext):
        # Web app stuff
        wa_url = os.getenv('WEBAPP_URL')
        params = "/?twan=" + user_client.get_webapp_nonce(self.TController.user)
        testWAInfo = WebAppInfo(wa_url + params)
        # end of web app stuff
        cKeyboardVals = [
        [KeyboardButton(text=CMD_IDENTIFIER+"HoTo")],
        [KeyboardButton(text=CMD_IDENTIFIER+"Faulty Equipment")],
        [KeyboardButton(text=CMD_IDENTIFIER+"Trespassing")],
        [KeyboardButton(text=CMD_IDENTIFIER+"Other Reports")],
        [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
        ]
        inlineKBs = InlineKeyboardMarkup([[InlineKeyboardButton(text="Report", web_app=testWAInfo)]])
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text=self.name, reply_markup=cKeyboard)
        context.bot.send_message(chat_id=update.effective_chat.id, text=self.name, reply_markup=inlineKBs)

class HoToReportMenu(TelegramMenu):
    def __init__(self, parent=None, name="", triggerWords=...):
        super().__init__(parent, name, triggerWords)
    def handler (self, update:Update, context:CallbackContext):
        # Need to send the keyboard that retracts the custom keyboard
        self.currentQuestion:str = "Qn1"
        self.textHandler(update, context)
        self.IsExpectingRawText = True
        self.answers:List[str] = []
        pass
        # Cannot append answers like this as each message is with a different
        # instance of HoToReportMenu
        # Need to store the informatoin in context.user_data instead
    def textHandler(self, update: Update, context: CallbackContext):
        if self.currentQuestion == "!Exit":
            context.bot.send_message(chat_id=update.effective_chat.id, text="HoTo report cancelled.")
            self.IsExpectingRawText = False
            self.backHandler(update, context)
        if self.currentQuestion == "Qn1":
            context.bot.send_message(chat_id=update.effective_chat.id, text="Starting HoTo report...")
            context.bot.send_message(chat_id=update.effective_chat.id, text="Type !Exit to cancel report making.")
            context.bot.send_message(chat_id=update.effective_chat.id, text="Who are you handing over to?")
            self.currentQuestion = "Qn2"
        elif self.currentQuestion == "Qn2":
            self.answers.append(update.effective_message.text)
            context.bot.send_message(chat_id=update.effective_chat.id, text="What is the fuel percentage?")
            self.currentQuestion = "Qn3"
        elif self.currentQuestion == "Qn3":
            self.answers.append(update.effective_message.text)
            replyText = "Handing over to " + self.answers[1] + " with fuel percentage at " + self.answers[2]
            context.bot.send_message(chat_id=update.effective_chat.id, text=replyText)
            self.IsExpectingRawText = False
            self.backHandler(update, context)
        pass