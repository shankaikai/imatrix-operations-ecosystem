from __future__ import annotations
from telegram import Update
from telegram.ext import CallbackContext
from telegram.ext import Updater, Dispatcher

from telegram.files.photosize import PhotoSize
from telegram import File as tFile

from typing import List, Dict
from abc import ABC, abstractclassmethod, abstractmethod

from subscriptions import subscription_message
from administration.TeleUser import TUser

import os.path
import time

from Protos import operations_ecosys_pb2

CMD_IDENTIFIER = "‎" 

class TelegramController:
    def __init__(self, updater:Updater, dispatcher:Dispatcher):
        self.updater:Updater = updater
        self.dispatcher:Dispatcher = dispatcher
        self.RootMenu:TelegramMenu = None
        self.Menus:list[TelegramMenu]= []
        self.CurrentMenu:TelegramMenu = None
        self.PreviousMenu:TelegramMenu = None
        self.user:TeleUser = None
        self.scratchData:Dict = {}
    def buildMenuTree(self, rootMenu:TelegramMenu):
        if self.Menus == None or len(self.Menus) == 0:
            print("Unable to build tree as no menus are given.")
            return
        if not (rootMenu in self.Menus):
            print("Unable to build tree as root menu is not in list of given menus.")
            return
        tempMenuHolder = {}
        for menu in self.Menus:
            menu.TController = self
            if menu == rootMenu:
                continue
            else:
                if menu.parent.name in tempMenuHolder:
                    tempMenuHolder[menu.parent.name].append(menu)
                else:
                    tempMenuHolder[menu.parent.name] = [menu]
        for menuName in tempMenuHolder:
            for menu in self.Menus:
                if menuName == menu.name:
                    menu.children = tempMenuHolder[menuName]

        self.RootMenu = rootMenu
    # Handlers that need to be attached directly to Tele Bot
    def mainHandler(self, update: Update, context: CallbackContext):
        if not self.ifPrivateChat(update, context):
            return
        text:str = update.message.text
        print(text)
        if len(text) > 1 and text[0] == CMD_IDENTIFIER and text[1] != CMD_IDENTIFIER:
            text = text[1:]
            if text == "Back" or text == "Cancel":
                self.CurrentMenu.backHandler(update, context)
                return
            for menu in self.CurrentMenu.children:
                if text in menu.triggerWords:
                    self.PreviousMenu = self.CurrentMenu
                    self.CurrentMenu = menu
                    menu.handler(update, context)
                    break
            self.catchAllHandler(update, context)
        else:
            self.CurrentMenu.textHandler(update, context)
    def startHandler(self, update: Update, context: CallbackContext):
        logging_in_user = TUser.create_Tele_User(update.effective_user.id)
        if not logging_in_user.oes_user:
            context.bot.send_message(chat_id=update.effective_chat.id, text="You are not authorised to use this bot. Please contact your administator if you believe otherwise.")
        elif update.effective_user.id != update.effective_chat.id:
            context.bot.send_message(chat_id=update.effective_chat.id, text="This bot can only be used in private chats.")
        else:
            logging_in_user.login(update)
            self.user = logging_in_user
            context.bot.send_message(chat_id=update.effective_chat.id, text="Welcome, " + self.user.oes_user.name + "!")
            self.PreviousMenu = None
            self.CurrentMenu = self.RootMenu
            self.RootMenu.handler(update, context)
        
    def attachmentHandler(self, update:Update, context:CallbackContext):
        if not self.ifPrivateChat(update, context):
            return
        self.CurrentMenu.attachmentHandler(update, context)
        pass
    # This handler is for in-line button presses
    def callbackqueryHandler(self, update:Update, context:CallbackContext):
        if not self.ifPrivateChat(update, context):
            return
        strData = update.callback_query.data
        # special case
        callback_type = strData.split(":")[0]
        print("HERE 101",callback_type,subscription_message.SubscriptionMessage.IDENTIFIER)
        if callback_type == subscription_message.SubscriptionMessage.IDENTIFIER:
            subscription_message.callbackqueryHandler(update, context)
            return
        # normal case
        if update.callback_query.data == "Back" \
            or update.callback_query.data == "Cancel" \
            or update.callback_query.data == "Exit":
            self.CurrentMenu.backHandler(update, context)
            return
        else:
            self.CurrentMenu.callbackqueryHandler(update, context)
        pass
    # Other handlers
    def backHandler(self, update:Update, context:CallbackContext):
        if(self.CurrentMenu.parent == None):
            return
        else:
            self.CurrentMenu = self.CurrentMenu.parent
            self.CurrentMenu.handler(update, context)
        pass
    def catchAllHandler(self, update: Update, context: CallbackContext):
        pass
    # Util
    def ifPrivateChat(self, update: Update, context: CallbackContext) -> bool:
        return update.effective_chat.id == update.effective_user.id

class TelegramMenu(ABC):
    def __init__(self, parent = None, name = "", triggerWords = []):
        self.parent:TelegramMenu = parent
        self.children:List[TelegramMenu] = []
        self.name:str = name
        self.triggerWords:List[str] = triggerWords
        self.TController:TelegramController = None
        self.scratchData = {}
    @abstractmethod
    def handler(self, update:Update, context:CallbackContext):
        pass
    def textHandler(self, update:Update, context:CallbackContext):
        pass
    def attachmentHandler(self, update:Update, context:CallbackContext):
        pass
    # Default backhandle - overwrite to change default behaviour
    # Note: To back up one level, always call TController
    def backHandler(self, update:Update, context:CallbackContext):
        self.TController.backHandler(update, context)
        pass