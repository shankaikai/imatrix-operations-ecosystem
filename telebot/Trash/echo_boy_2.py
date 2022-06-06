import contextvars
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

import MenuLocation

SAVE_PATH = "/files"
SAVE_PATH = "F:/Space/OneDrive/School/SUTD/Capstone/DLBot/files/"

TOKEN = "5267388669:AAEPBKhdSdB6F088WC_-q4X5JtRqKzB8Ccs"

updater = Updater(token=TOKEN, use_context=True)
dispatcher = updater.dispatcher


def start(update: Update, context: CallbackContext):
    context.bot.send_message(chat_id=update.effective_chat.id, text="Welcome!")
    MainMenuHandler(update, context)

def echo(update: Update, context: CallbackContext):
    text = update.message.text
    if text == "Back" or text == "Cancel":
        backHandler(update, context)
    elif text == "Attendance":
        AttendanceMenuHandler(update, context)
    elif text == "Reporting":
        ReportMenuHandler(update, context)
    elif text == "SOS":
        SOSMenuHandler(update, context)
    elif text == "Help":
        HelpMenuHandler(update, context)
    elif text == "User Guide":
        userguideHandler(update, context)
    elif text == "Bug Report":
        bugReportHandler(update, context)
    elif text == "HoTo" or text == "Faulty Equipment" or text == "Trespassing" or text == "Other Reports":
        reportMaker(update, context, text)
    else:
        nlpHandler(update, context)
        return
        
def backHandler(update: Update, context: CallbackContext):
    mLocation = MenuLocation.getLocation(context.user_data[MenuLocation.LOCATION_KEY])
    parent = mLocation.parent
    if parent == None:
        goToMenu(update, context, MenuLocation.LOC_MAINMENU.name)
        return
    goToMenu(update, context, parent.name)
    
def goToMenu(update: Update, context: CallbackContext, lName: str):
    if lName == MenuLocation.MAINMENU_NAME:
        MainMenuHandler(update, context)
    elif lName == MenuLocation.ATTENDANCEMENU_NAME:
        AttendanceMenuHandler(update, context)
    elif lName == MenuLocation.SOSMENU_NAME:
        SOSMenuHandler(update, context)
    elif lName == MenuLocation.REPORTMENU_NAME:
        ReportMenuHandler(update, context)
    elif lName == MenuLocation.HELPMENU_NAME:
        HelpMenuHandler(update, context)

# MENUS

def MainMenuHandler(update: Update, context: CallbackContext):
    context.user_data[MenuLocation.LOCATION_KEY] = MenuLocation.LOC_MAINMENU.name
    cKeyboardVals = [
        [KeyboardButton(text="Attendance")],
        [KeyboardButton(text="Reporting")],
        [KeyboardButton(text="SOS"), KeyboardButton(text="Help")]
    ]
    cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
    context.bot.send_message(chat_id=update.effective_chat.id, text=MenuLocation.LOC_MAINMENU.name, reply_markup=cKeyboard)

def AttendanceMenuHandler(update: Update, context: CallbackContext):
    context.user_data[MenuLocation.LOCATION_KEY] = MenuLocation.LOC_ATTENDANCEMENU.name
    cKeyboardVals = [
        [KeyboardButton(text="Cancel")]
    ]
    cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals, one_time_keyboard=True)
    context.bot.send_message(chat_id=update.effective_chat.id, text="Please send your photo.", reply_markup=cKeyboard)
    pass

def ReportMenuHandler(update: Update, context: CallbackContext):
    context.user_data[MenuLocation.LOCATION_KEY] = MenuLocation.LOC_REPORTMENU.name
    cKeyboardVals = [
        [KeyboardButton(text="HoTo")],
        [KeyboardButton(text="Faulty Equipment")],
        [KeyboardButton(text="Trespassing")],
        [KeyboardButton(text="Other Reports")],
        [KeyboardButton(text="Back")]
    ]
    cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
    context.bot.send_message(chat_id=update.effective_chat.id, text=MenuLocation.LOC_REPORTMENU.name, reply_markup=cKeyboard)
    pass

def reportMaker(update: Update, context: CallbackContext, typeOfForm: str):
    if typeOfForm == "Hoto":
        print("HoTo Form")
    elif typeOfForm == "Faulty Equipment":
        pass
    elif typeOfForm == "Trespassing":
        pass
    elif typeOfForm == "Other Reports":
        pass

def SOSMenuHandler(update: Update, context: CallbackContext):
    context.user_data[MenuLocation.LOCATION_KEY] = MenuLocation.LOC_SOSMENU.name
    cKeyboardVals = [
        [KeyboardButton(text="Cancel")]
    ]
    cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
    context.bot.send_message(chat_id=update.effective_chat.id, text="Dialing HQCC in 5 seconds!", reply_markup=cKeyboard)
    context.job_queue.run_once(_sosConfirmPrint, 5, context=(update, context))
    #time.sleep(5)
    #_sosConfirmPrint(update, context)
    pass

def _sosConfirmPrint(pContext: CallbackContext):
    update, context = pContext.job.context
    if context.user_data[MenuLocation.LOCATION_KEY] == MenuLocation.LOC_SOSMENU.name:
        context.bot.send_message(chat_id=update.effective_chat.id, text="Dialing HQCC now!")
        MainMenuHandler(update, context)

def HelpMenuHandler(update: Update, context: CallbackContext):
    context.user_data[MenuLocation.LOCATION_KEY] = MenuLocation.LOC_HELPMENU.name
    cKeyboardVals = [
        [KeyboardButton(text="User Guide")],
        [KeyboardButton(text="Bug Report")],
        [KeyboardButton(text="Back")]
    ]
    cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
    context.bot.send_message(chat_id=update.effective_chat.id, text=MenuLocation.LOC_HELPMENU.name, reply_markup=cKeyboard)
    pass


def userguideHandler(update: Update, context: CallbackContext):
    pass

def bugReportHandler(update: Update, context: CallbackContext):
    pass

def nlpHandler(update: Update, context: CallbackContext):
    pass

# <><><>

def attachmentHandler(update: Update, context: CallbackContext):
    attachment = update.effective_message.effective_attachment
    fileType = type(attachment)
    # Only photos should appear in a list, other attachment
    # types should be singular
    if fileType == type([]) and len(attachment) > 0:
        fileType = type(attachment[1])
    if fileType == PhotoSize:
        if context.user_data[MenuLocation.LOCATION_KEY] != MenuLocation.LOC_ATTENDANCEMENU.name:
            return
        fileName = update.effective_user.username + ".jpg"
        photo = attachment[1].get_file()
        photo.download(os.path.join(SAVE_PATH, fileName))
        print("Attendance photo for " + update.effective_user.username + " obtained.")
        context.bot.send_message(chat_id=update.effective_chat.id, text="Your attendance has been submitted.")
        MainMenuHandler(update, context)
        return

start_handler = CommandHandler('start', start)
echo_handler = MessageHandler(Filters.text & (~Filters.command), echo)
attachment_handler = MessageHandler(Filters.attachment & (~Filters.command), attachmentHandler)
dispatcher.add_handler(start_handler)
dispatcher.add_handler(echo_handler)
dispatcher.add_handler(attachment_handler)

updater.start_polling()
updater.idle()