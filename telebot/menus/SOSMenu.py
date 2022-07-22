from __future__ import annotations
from telegram import Update, KeyboardButton, ReplyKeyboardMarkup, ReplyKeyboardRemove
from telegram.ext import CallbackContext
from telegram.files.photosize import PhotoSize

import os.path

from TelegramController import TelegramController, TelegramMenu, CMD_IDENTIFIER

class SOSMenu(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "SOS Menu"
    def handler(self, update:Update, context:CallbackContext):
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Cancel")]
        ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text="Dialing HQCC in 5 seconds!", reply_markup=cKeyboard)
        context.job_queue.run_once(self._sosConfirmPrint, 5, context=(update, context))
    def _sosConfirmPrint(self, pContext: CallbackContext):
        update, context = pContext.job.context
        if self.TController.CurrentMenu == self:
            context.bot.send_message(chat_id=update.effective_chat.id, text="Dialing HQCC now!")
            self.TController.backHandler(update, context)