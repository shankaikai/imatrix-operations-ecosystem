from __future__ import annotations
from telegram import Update
from telegram.ext import CallbackContext
from telegram.ext import Updater, Dispatcher

from telegram.files.photosize import PhotoSize
from telegram import File as tFile

from typing import List
from abc import ABC, abstractclassmethod, abstractmethod

from subscriptions import subscription_message
from administration import user

import os.path
import time

CMD_IDENTIFIER = "‎" 

class TelegramController:
    def __init__(self, updater:Updater, dispatcher:Dispatcher):
        self.updater:Updater = updater
        self.dispatcher:Dispatcher = dispatcher
        self.RootMenu:TelegramMenu = None
        self.Menus:list[TelegramMenu]= []
        self.CurrentMenu:TelegramMenu = None
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
        text:str = update.message.text
        if len(text) > 1 and text[0] == CMD_IDENTIFIER and text[1] != CMD_IDENTIFIER:
            text = text[1:]
            if text == "Back" or text == "Cancel":
                self.CurrentMenu.backHandler(update, context)
                return
            for menu in self.CurrentMenu.children:
                if text in menu.triggerWords:
                    self.CurrentMenu = menu
                    menu.handler(update, context)
                    break
            self.catchAllHandler(update, context)
        else:
            self.CurrentMenu.textHandler(update, context)
    def startHandler(self, update: Update, context: CallbackContext):
        if user.login(update.effective_chat.username, update.effective_chat.id):
            context.bot.send_message(chat_id=update.effective_chat.id, text="Welcome!")
            self.RootMenu.handler(update, context)
        else:
            context.bot.send_message(chat_id=update.effective_chat.id, text="You are not authorised to use this bot. Please contact your administator if you believe otherwise.")
        
    def attachmentHandler(self, update:Update, context:CallbackContext):
        self.CurrentMenu.attachmentHandler(update, context)
        pass
    # This handler is for in-line button presses
    def callbackqueryHandler(self, update:Update, context:CallbackContext):
        strData = update.callback_query.data
        callback_type = strData.split(":")[0]
        if callback_type == subscription_message.IDENTIFIER:
            subscription_message.callbackqueryHandler(update, context)
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

class TelegramMenu(ABC):
    def __init__(self, parent = None, name = "", triggerWords = []):
        self.parent:TelegramMenu = parent
        self.children:List[TelegramMenu] = []
        self.name:str = name
        self.triggerWords:List[str] = triggerWords
        self.TController:TelegramController = None
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