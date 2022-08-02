from __future__ import annotations
from telegram import Update, WebAppInfo, InlineKeyboardMarkup, InlineKeyboardButton
from telegram.ext import CallbackContext
import os

from grpc_clients import user_client

def try_open_registration_webapp(update:Update, context:CallbackContext):
    print("register")
    rawCommand = update.message.text
    tokenised = rawCommand.split()
    if len(tokenised) != 2:
        return
    registration_code = tokenised[1]
    #ToDo: Check with DB if registration code is valid
    security_string = user_client.check_reg_code(registration_code)
    if security_string == None:
        context.bot.send_message(chat_id=update.effective_chat.id, text="The provided code is invalid.")
        return
    wa_url = os.getenv('WEBAPP_REGISTRATION_URL')
    params = "?code=" + registration_code
    params += "&user_id=" + str(update.effective_user.id)
    params += "&tele_handle=" + str(update.effective_user.username)
    params += "&security_string=" + security_string
    WAInfo = WebAppInfo(wa_url + params)
    inlineKBs = InlineKeyboardMarkup([[InlineKeyboardButton(text="Create Account", web_app=WAInfo)]])
    context.bot.send_message(chat_id=update.effective_chat.id, text="The provided code is valid, please proceed with registration.", reply_markup=inlineKBs)
    pass