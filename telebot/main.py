from __future__ import annotations
import contextvars
from setuptools import Command
from telegram import Update, KeyboardButton, ReplyKeyboardMarkup, ReplyKeyboardRemove
from telegram.ext import CallbackContext, CallbackQueryHandler
from telegram.ext import Updater, Dispatcher
from telegram.ext import CommandHandler
from telegram.ext import MessageHandler, Filters
from telegram.files.photosize import PhotoSize
from telegram import File as tFile

import os.path
import time

from TelegramController import TelegramController, TelegramMenu
from Menus import MainMenu, OthersMenu, AdminMenu
from grpc_servers.grpc_server import serve
from Standalones import Registration

from dotenv import load_dotenv

load_dotenv()

# CONFIG Params
TOKEN = os.getenv('TOKEN')

# Get TeleBot updater and dispatcher
updater:Updater= Updater(token=TOKEN, use_context=True)
dispatcher:Dispatcher = updater.dispatcher
    
# Init and configure the TController
TController = TelegramController(updater, dispatcher)

# Create TelegramMenus
mainMenu = MainMenu.MainMenu()
mainMenuAWA = MainMenu.MainMenu_Attendance_WA(parent=mainMenu, triggerWords=["My Assignments"])
mainMenuRWA = MainMenu.MainMenu_Report_WA(parent=mainMenu, triggerWords=["Make a Report"])
#attendanceMenu = AttendanceMenu.AttendanceMenu(parent=mainMenu, triggerWords=["Attendance"])
#reportMenu = ReportMenu.ReportMenu(parent=mainMenu, triggerWords=["Reporting"])
#sosMenu = SOSMenu.SOSMenu(parent=mainMenu, triggerWords=["SOS"])
othersMenu = OthersMenu.OthersMenu(parent=mainMenu, triggerWords=["Others"])
othersMenuQI = OthersMenu.OthersMenu_QuickIntro_Text(parent=othersMenu, triggerWords=["Quick Introduction"])
adminMenu = AdminMenu.AdminMenu(parent=othersMenu, triggerWords=["Admin"])
adminMenuRCM = AdminMenu.AdminMenu_RegistrationCodeMenu(parent=adminMenu, triggerWords=["Create Registration Code"])
adminMenuRCM_IS = AdminMenu.AdminMenu_RegistrationCodeMenu_ISpec(parent=adminMenuRCM, triggerWords=["Create I-Specialist Code"])
adminMenuRCM_SG = AdminMenu.AdminMenu_RegistrationCodeMenu_SGuard(parent=adminMenuRCM, triggerWords=["Create Security Guard Code"])
adminMenuRCM_C = AdminMenu.AdminMenu_RegistrationCodeMenu_Controller(parent=adminMenuRCM, triggerWords=["Create Controller Code"])
adminMenuRCM_M = AdminMenu.AdminMenu_RegistrationCodeMenu_Manager(parent=adminMenuRCM, triggerWords=["Create Manager Code"])
adminMenuSU = AdminMenu.AdminMenu_SearchUserMenu(parent=adminMenu, triggerWords=["Search User"])
adminMenuSU_TUNAME = AdminMenu.AdminMenu_SearchUserMenu_byTeleUsername(parent=adminMenuSU, triggerWords=["Search by Telegram username"])
adminMenuSU_NAME = AdminMenu.AdminMenu_SearchUserMenu_byName(parent=adminMenuSU, triggerWords=["Search by name"])
adminMenuSU_PN = AdminMenu.AdminMenu_SearchUserMenu_byPhoneNumber(parent=adminMenuSU, triggerWords=["Search by phone number"])
adminMenuSU_EMAIL = AdminMenu.AdminMenu_SearchUserMenu_byEmail(parent=adminMenuSU, triggerWords=["Search by email"])
adminMenuSU_TUID = AdminMenu.AdminMenu_SearchUserMenu_byTelegramUserID(parent=adminMenuSU, triggerWords=["Search by Telegram user ID"])



# Add TelegramMenus to TController
menus = [
    mainMenu, mainMenuAWA, mainMenuRWA, 
    othersMenu, othersMenuQI, 
    adminMenu, adminMenuRCM, adminMenuRCM_IS, adminMenuRCM_SG, adminMenuRCM_C, adminMenuRCM_M,
    adminMenuSU, adminMenuSU_TUNAME, adminMenuSU_NAME, adminMenuSU_PN, adminMenuSU_EMAIL, adminMenuSU_TUID
    ]
TController.Menus = menus
TController.buildMenuTree(mainMenu)
TController.CurrentMenu = mainMenu

# Link TController with TeleBot 
start_handler = CommandHandler('start', TController.startHandler)
registration_handler = CommandHandler('register', Registration.try_open_registration_webapp)
echo_handler = MessageHandler(Filters.text & (~Filters.command), TController.mainHandler)
attachment_handler = MessageHandler(Filters.attachment & (~Filters.command), TController.attachmentHandler)
callbackquery_handler = CallbackQueryHandler(TController.callbackqueryHandler)
dispatcher.add_handler(start_handler)
dispatcher.add_handler(registration_handler)
dispatcher.add_handler(echo_handler)
dispatcher.add_handler(attachment_handler)
dispatcher.add_handler(callbackquery_handler)

# Start grpc server
grpc_server = serve(updater)
print("Started Tele Bot gRPC Server")

updater.start_polling(timeout=100)
print("Telegram bot started")

updater.idle()