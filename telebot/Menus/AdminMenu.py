from __future__ import annotations
from telegram import Update, KeyboardButton, ReplyKeyboardMarkup, ReplyKeyboardRemove, InlineKeyboardMarkup, InlineKeyboardButton
from telegram.ext import CallbackContext
from telegram.files.photosize import PhotoSize

import os.path

from TelegramController import TelegramController, TelegramMenu, CMD_IDENTIFIER
from administration.TeleUser import TUser

from grpc_clients import user_client
from Protos import operations_ecosys_pb2

class AdminMenu(TelegramMenu):
    def __init__(self, parent=None, name="", triggerWords=[]):
        super().__init__(parent, name, triggerWords)
        self.name = "Admin Menu"
    def handler(self, update:Update, context:CallbackContext):
        if not (self.TController.user.user_type == operations_ecosys_pb2.User.UserType.MANAGER \
        or self.TController.user.user_type == operations_ecosys_pb2.User.UserType.CONTROLLER):
            context.bot.send_message(chat_id=update.effective_chat.id, text="You do not have permission to access this menu.")
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Create Registration Code")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Search User")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
        ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text=self.name, reply_markup=cKeyboard)

# Search
class AdminMenu_SearchUserMenu(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "Search User Menu"
    def handler(self, update:Update, context:CallbackContext):
        if not (self.TController.user.user_type == operations_ecosys_pb2.User.UserType.MANAGER \
        or self.TController.user.user_type == operations_ecosys_pb2.User.UserType.CONTROLLER):
            context.bot.send_message(chat_id=update.effective_chat.id, text="You do not have permission to do this.")
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Search by Telegram username")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Search by name")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Search by phone number")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Search by email")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Search by Telegram user ID")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
            ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text=self.name, reply_markup=cKeyboard)
        pass

class AdminMenu_SearchUserMenu_byTeleUsername(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "adminMenu_searchMenu_byTeleUsername"
    def handler(self, update:Update, context:CallbackContext):
        if not (self.TController.user.user_type == operations_ecosys_pb2.User.UserType.MANAGER \
        or self.TController.user.user_type == operations_ecosys_pb2.User.UserType.CONTROLLER):
            context.bot.send_message(chat_id=update.effective_chat.id, text="You do not have permission to do this.")
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
            ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text="Please enter a or a part of a Telegram username.", reply_markup=cKeyboard)
    def textHandler(self, update:Update, context:CallbackContext):
        searchValue = update.effective_message.text
        if len(searchValue) == 0:
            return
        foundUsers = user_client.lookup_users(telegram_username=searchValue)
        return _adminMenu_search_util_textHandler(self, update, context, foundUsers)
    def callbackqueryHandler(self, update:Update, context:CallbackContext):
        return _adminMenu_search_util_callbackqueryHandler(self, update, context)

class AdminMenu_SearchUserMenu_byName(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "adminMenu_searchMenu_byName"
    def handler(self, update:Update, context:CallbackContext):
        if not (self.TController.user.user_type == operations_ecosys_pb2.User.UserType.MANAGER \
        or self.TController.user.user_type == operations_ecosys_pb2.User.UserType.CONTROLLER):
            context.bot.send_message(chat_id=update.effective_chat.id, text="You do not have permission to do this.")
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
            ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text="Please enter a or a part of a name.", reply_markup=cKeyboard)
    def textHandler(self, update:Update, context:CallbackContext):
        searchValue = update.effective_message.text
        if len(searchValue) == 0:
            return
        foundUsers = user_client.lookup_users(name=searchValue)
        return _adminMenu_search_util_textHandler(self, update, context, foundUsers)
    def callbackqueryHandler(self, update:Update, context:CallbackContext):
        return _adminMenu_search_util_callbackqueryHandler(self, update, context)

class AdminMenu_SearchUserMenu_byPhoneNumber(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "adminMenu_searchMenu_byPhoneNumber"
    def handler(self, update:Update, context:CallbackContext):
        if not (self.TController.user.user_type == operations_ecosys_pb2.User.UserType.MANAGER \
        or self.TController.user.user_type == operations_ecosys_pb2.User.UserType.CONTROLLER):
            context.bot.send_message(chat_id=update.effective_chat.id, text="You do not have permission to do this.")
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
            ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text="Please enter a phone number.", reply_markup=cKeyboard)
    def textHandler(self, update:Update, context:CallbackContext):
        searchValue = update.effective_message.text
        if len(searchValue) == 0:
            return
        foundUsers = user_client.lookup_users(phone_num=searchValue)
        return _adminMenu_search_util_textHandler(self, update, context, foundUsers)
    def callbackqueryHandler(self, update:Update, context:CallbackContext):
        return _adminMenu_search_util_callbackqueryHandler(self, update, context)

class AdminMenu_SearchUserMenu_byEmail(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "adminMenu_searchMenu_byEmail"
    def handler(self, update:Update, context:CallbackContext):
        if not (self.TController.user.user_type == operations_ecosys_pb2.User.UserType.MANAGER \
        or self.TController.user.user_type == operations_ecosys_pb2.User.UserType.CONTROLLER):
            context.bot.send_message(chat_id=update.effective_chat.id, text="You do not have permission to do this.")
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
            ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text="Please enter an or a part of an email.", reply_markup=cKeyboard)
    def textHandler(self, update:Update, context:CallbackContext):
        searchValue = update.effective_message.text
        if len(searchValue) == 0:
            return
        foundUsers = user_client.lookup_users(email=searchValue)
        return _adminMenu_search_util_textHandler(self, update, context, foundUsers)
    def callbackqueryHandler(self, update:Update, context:CallbackContext):
        return _adminMenu_search_util_callbackqueryHandler(self, update, context)

class AdminMenu_SearchUserMenu_byTelegramUserID(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "adminMenu_searchMenu_byTelegramUserID"
    def handler(self, update:Update, context:CallbackContext):
        if not (self.TController.user.user_type == operations_ecosys_pb2.User.UserType.MANAGER \
        or self.TController.user.user_type == operations_ecosys_pb2.User.UserType.CONTROLLER):
            context.bot.send_message(chat_id=update.effective_chat.id, text="You do not have permission to do this.")
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Back")]
            ]
        cKeyboard = ReplyKeyboardMarkup(keyboard=cKeyboardVals)
        context.bot.send_message(chat_id=update.effective_chat.id, text="Please enter a Telegram user ID.", reply_markup=cKeyboard)
    def textHandler(self, update:Update, context:CallbackContext):
        searchValue = update.effective_message.text
        if len(searchValue) == 0:
            return
        foundUsers = user_client.lookup_users(telegram_user_id=searchValue)
        return _adminMenu_search_util_textHandler(self, update, context, foundUsers)
    def callbackqueryHandler(self, update:Update, context:CallbackContext):
        return _adminMenu_search_util_callbackqueryHandler(self, update, context)

# Search Utility
def _adminMenu_search_util_textHandler(tMenu:TelegramMenu, update:Update, context:CallbackContext, found_oes_users):
    tMenu.scratchData["found_users"] = []
    tMenu.scratchData["found_users_index"] = 0
    if len(found_oes_users) == 0:
        context.bot.send_message(chat_id=update.effective_chat.id, text="No users found.")
        return
    for found_oes_user in found_oes_users:
        tMenu.scratchData["found_users"].append(TUser.create_Tele_User_from_oes_user(found_oes_user))
    tMenu.callbackqueryHandler(update, context)   

def _adminMenu_search_util_callbackqueryHandler(tMenu:TelegramMenu, update:Update, context:CallbackContext):
    numUsers = len(tMenu.scratchData["found_users"])
    userIdx = tMenu.scratchData["found_users_index"]

    if update.callback_query != None:
        update.callback_query.answer()
        action = update.callback_query.data
        if action == "Prev":
            if userIdx > 0:
                userIdx = userIdx - 1
            else:
                return
            pass
        elif action == "Next":
            if userIdx < numUsers - 1:
                userIdx = userIdx + 1
            else:
                return
            pass
        elif action == "Confirm Delete User" and tMenu.TController.user.user_type == operations_ecosys_pb2.User.UserType.MANAGER:
            selectedUser = tMenu.scratchData["found_users"][userIdx]
            if user_client.delete_user(selectedUser.oes_user.user_id):
                numUsers = numUsers - 1
                if numUsers < 1:
                    update.callback_query.edit_message_text("User deleted. No results left.")
                    tMenu.TController.backHandler(update, context)
                    return
                tMenu.scratchData["found_users"].pop(userIdx)
                if userIdx == numUsers:
                    userIdx = userIdx - 1

        tMenu.scratchData["found_users_index"] = userIdx

    user_text = "<><> Result " + str(userIdx+1) + " of " + str(numUsers) + " <><>\n"
    user_text += tMenu.scratchData["found_users"][userIdx].__str__()

    inlineKBButtons = [
            [InlineKeyboardButton(text="Prev", callback_data="Prev"), InlineKeyboardButton(text="Next", callback_data="Next")],
            [InlineKeyboardButton(text="Exit", callback_data="Exit")]
        ]
    
    if tMenu.TController.user.user_type == operations_ecosys_pb2.User.UserType.MANAGER:
        action = ""
        if update.callback_query != None:
            update.callback_query.answer()
            action = update.callback_query.data
        if action == "Delete User":
            inlineKBButtons.insert(1, [InlineKeyboardButton(text="Confirm Delete User", callback_data="Confirm Delete User")])
        else:
            inlineKBButtons.insert(1, [InlineKeyboardButton(text="Delete User", callback_data="Delete User")])
        

    inlineKBs = InlineKeyboardMarkup(inlineKBButtons)

    if update.callback_query != None:
        update.callback_query.edit_message_text(user_text, reply_markup=inlineKBs)
    else:
        context.bot.send_message(chat_id=update.effective_chat.id, text=user_text, reply_markup=inlineKBs)
    pass

# Registration Code
class AdminMenu_RegistrationCodeMenu(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "Registration Code Menu"
    def handler(self, update:Update, context:CallbackContext):
        if self.TController.user.user_type != operations_ecosys_pb2.User.UserType.MANAGER:
            context.bot.send_message(chat_id=update.effective_chat.id, text="You do not have permission to do this.")
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        cKeyboardVals = [
            [KeyboardButton(text=CMD_IDENTIFIER+"Create I-Specialist Code")],
            [KeyboardButton(text=CMD_IDENTIFIER+"Create Security Guard Code")],
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
        if self.TController.user.user_type != operations_ecosys_pb2.User.UserType.MANAGER:
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        rCode = self.TController.user.getRegistrationCode(operations_ecosys_pb2.RegistrationCodeRequest.ISPECIALIST)
        cMessage = "I-Specialist registration code: " + rCode
        context.bot.send_message(chat_id=update.effective_chat.id, text=cMessage)
        self.TController.CurrentMenu = self.parent
        pass

class AdminMenu_RegistrationCodeMenu_SGuard(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "AdminMenu_RegistrationCodeMenu_SGuard"
    def handler(self, update:Update, context:CallbackContext):
        if self.TController.user.user_type != operations_ecosys_pb2.User.UserType.MANAGER:
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        rCode = self.TController.user.getRegistrationCode(operations_ecosys_pb2.RegistrationCodeRequest.SECURITYGUARD)
        cMessage = "Security Guard registration code: " + rCode
        context.bot.send_message(chat_id=update.effective_chat.id, text=cMessage)
        self.TController.CurrentMenu = self.parent
        pass

class AdminMenu_RegistrationCodeMenu_Controller(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "AdminMenu_RegistrationCodeMenu_Controller"
    def handler(self, update:Update, context:CallbackContext):
        if self.TController.user.user_type != operations_ecosys_pb2.User.UserType.MANAGER:
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        rCode = self.TController.user.getRegistrationCode(operations_ecosys_pb2.RegistrationCodeRequest.CONTROLLER)
        cMessage = "Controller registration code: " + rCode
        context.bot.send_message(chat_id=update.effective_chat.id, text=cMessage)
        self.TController.CurrentMenu = self.parent
        pass

class AdminMenu_RegistrationCodeMenu_Manager(TelegramMenu):
    def __init__(self, parent = None, name = "", triggerWords = []):
        super().__init__(parent, name, triggerWords)
        self.name = "AdminMenu_RegistrationCodeMenu_Manager"
    def handler(self, update:Update, context:CallbackContext):
        if self.TController.user.user_type != operations_ecosys_pb2.User.UserType.MANAGER:
            self.TController.CurrentMenu = self.TController.PreviousMenu
            return
        rCode = self.TController.user.getRegistrationCode(operations_ecosys_pb2.RegistrationCodeRequest.MANAGER)
        cMessage = "Manager registration code: " + rCode
        context.bot.send_message(chat_id=update.effective_chat.id, text=cMessage)
        self.TController.CurrentMenu = self.parent
        pass
