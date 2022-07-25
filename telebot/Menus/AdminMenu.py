from __future__ import annotations
from telegram import Update, KeyboardButton, ReplyKeyboardMarkup, ReplyKeyboardRemove
from telegram.ext import CallbackContext
from telegram.files.photosize import PhotoSize

import os.path

from TelegramController import TelegramController, TelegramMenu, CMD_IDENTIFIER

class AdminMenu(TelegramMenu):
    def __init__(self, parent=None, name="", triggerWords=[]):
        super().__init__(parent, name, triggerWords)
        self.name = "Admin Menu"
    def handler(self, update:Update, context:CallbackContext):
        if self.TController.user.user_type != 3:
            context.bot.send_message(chat_id=update.effective_chat.id, text="You do not have permission to access this menu.")
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Create Registration Code")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Search User")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Delete User")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
        ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text=self.name, reply_markup=cKeyboard)

# Registration Code
class AdminMenu_RegistrationCodeMenu(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "Registration Code Menu"
    def handler(self, update:Update, context:CallbackContext):
        # ToDo: Check with Hannah use enum instead of int
        if self.TController.user.user_type != 3:
            context.bot.send_message(chat_id=update.effective_chat.id, text="You do not have permission to do this.")
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Create I-Specialist Code")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Create Controller Code")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Create Manager Code")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
            ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text=self.name, reply_markup=cKeyboard)
        pass

class AdminMenu_RegistrationCodeMenu_ISpec(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "AdminMenu_RegistrationCodeMenu_ISpec"
    def handler(self, update:Update, context:CallbackContext):
        if self.TController.user.user_type != 3:
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        #ToDo: send the request
        rCode = "this_is_a_fake_code"
        cMessage = "I-Specialist registration code: " + rCode
        context.bot.send_message(chat_id=update.effective_chat.id, text=cMessage)
        self.TController.CurrentMenu = self.parent
        pass

class AdminMenu_RegistrationCodeMenu_Controller(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "AdminMenu_RegistrationCodeMenu_Controller"
    def handler(self, update:Update, context:CallbackContext):
        if self.TController.user.user_type != 3:
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        #ToDo: send the request
        rCode = "this_is_a_fake_code"
        cMessage = "Controller registration code: " + rCode
        context.bot.send_message(chat_id=update.effective_chat.id, text=cMessage)
        self.TController.CurrentMenu = self.parent
        pass

class AdminMenu_RegistrationCodeMenu_Manager(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "AdminMenu_RegistrationCodeMenu_Manager"
    def handler(self, update:Update, context:CallbackContext):
        if self.TController.user.user_type != 3:
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        #ToDo: send the request
        rCode = "this_is_a_fake_code"
        cMessage = "Manager registration code: " + rCode
        context.bot.send_message(chat_id=update.effective_chat.id, text=cMessage)
        self.TController.CurrentMenu = self.parent
        pass