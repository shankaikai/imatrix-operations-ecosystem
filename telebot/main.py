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

from TelegramController import TelegramController, TelegramMenu
from Menus import MainMenu, AttendanceMenu, ReportMenu, SOSMenu, HelpMenu

from dotenv import load_dotenv

load_dotenv()

# CONFIG Params
TOKEN = os.getenv('TOKEN')

# Get TeleBot updated and dispatcher
updater = Updater(token=TOKEN, use_context=True)
dispatcher = updater.dispatcher
    
# Init and configure the TController
TController = TelegramController(updater, dispatcher)

# Create TelegramMenus
mainMenu = MainMenu.MainMenu()
attendanceMenu = AttendanceMenu.AttendanceMenu(parent=mainMenu, triggerWords=["Attendance"])
reportMenu = ReportMenu.ReportMenu(parent=mainMenu, triggerWords=["Reporting"])
sosMenu = SOSMenu.SOSMenu(parent=mainMenu, triggerWords=["SOS"])
helpMenu = HelpMenu.HelpMenu(parent=mainMenu, triggerWords=["Help"])

hotoReportMenu = ReportMenu.HoToReportMenu(parent=reportMenu, triggerWords=["HoTo"])

# Add TelegramMenus to TController
menus = [mainMenu, attendanceMenu, reportMenu, sosMenu, helpMenu, hotoReportMenu]
TController.Menus = menus
TController.buildMenuTree(mainMenu)
TController.CurrentMenu = mainMenu

# Link TController with TeleBot 
start_handler = CommandHandler('start', TController.startHandler)
echo_handler = MessageHandler(Filters.text & (~Filters.command), TController.mainHandler)
attachment_handler = MessageHandler(Filters.attachment & (~Filters.command), TController.attachmentHandler)
dispatcher.add_handler(start_handler)
dispatcher.add_handler(echo_handler)
dispatcher.add_handler(attachment_handler)

updater.start_polling()
updater.idle()