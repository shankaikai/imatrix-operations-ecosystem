from __future__ import annotations
from telegram import Update, WebAppInfo, InlineKeyboardMarkup, InlineKeyboardButton
from telegram.ext import CallbackContext
import os

from grpc_clients import user_client

def try_link_via_telehandle(update:Update, context:CallbackContext):
    print("link")
    rawCommand = update.message.text
    tokenised = rawCommand.split()
    if len(tokenised) != 1:
        return
    #Try to update user at telehandle
    if not (user_client.link_via_telehandle(update.effective_user.username, update.effective_user.id)):
        context.bot.send_message(chat_id=update.effective_chat.id, text="No account exists.")
        return
    context.bot.send_message(chat_id=update.effective_chat.id, text="Account linked. Please proceed with /start")
    pass