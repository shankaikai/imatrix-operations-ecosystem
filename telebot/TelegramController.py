from __future__ import annotations
import contextvars
from turtle import update
from setuptools import Command
from telegram import Update, KeyboardButton, ReplyKeyboardMarkup, ReplyKeyboardRemove
from telegram.ext import CallbackContext
from telegram.ext import Updater
from telegram.ext import CommandHandler
from telegram.ext import MessageHandler, Filters
from telegram.files.photosize import PhotoSize
from telegram import File as tFile

import os.path
import time

class TelegramController:
    def __init__(self, updater, dispatcher):
        self.updated = updater
        self.dispatcher = dispatcher
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
        if self.CurrentMenu.IsExpectingRawText:
            self.CurrentMenu.textHandler(update, context)
        else:
            if text == "Back" or text == "Cancel":
                self.backHandler(update, context)
                return
            for menu in self.Menus:
                if text in menu.triggerWords:
                    self.CurrentMenu = menu
                    menu.handler(update, context)
                    break
            self.catchAllHandler(update, context)
    def startHandler(self, update: Update, context: CallbackContext):
        context.bot.send_message(chat_id=update.effective_chat.id, text="Welcome!")
        self.RootMenu.handler(update, context)
        pass
    def attachmentHandler(self, update:Update, context:CallbackContext):
        self.CurrentMenu.attachmentHandler(update, context)
        pass
    # Other handlers
    def backHandler(self, update:Update, context:CallbackContext):
        #Just for safety
        if self.CurrentMenu.IsExpectingRawText:
            self.CurrentMenu.IsExpectingRawText = False
            print("Check if this was intended.")
        if(self.CurrentMenu.parent == None):
            return
        else:
            self.CurrentMenu = self.CurrentMenu.parent
            self.CurrentMenu.handler(update, context)
        pass
    def catchAllHandler(self, update: Update, context: CallbackContext):
        pass

class TelegramMenu:
    def __init__(self, parent = None, name = "", triggerWords = []):
        self.parent = parent
        self.children = []
        self.name = name
        self.triggerWords = triggerWords
        self.TController:TelegramController = None
        self.IsExpectingRawText = False
    def handler(self, update:Update, context:CallbackContext):
        print(self.name + "'s handler is not implemented.")
        pass
    def textHandler(self, update:Update, context:CallbackContext):
        print(self.name + "'s text handler is not implemented.")
        pass
    def attachmentHandler(self, update:Update, context:CallbackContext):
        print(self.name + "'s attachment handler is ont implemented.")
        pass