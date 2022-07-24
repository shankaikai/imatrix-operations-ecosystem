from __future__ import print_function

from telegram import Update

from grpc_clients import user_client

import logging
import random
import grpc

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2

class TUser:
    def __init__(self, tele_user_id:int, oes_user:operations_ecosys_pb2.User):
        self.tele_user_id = tele_user_id
        self.oes_user = oes_user
        pass

    @classmethod
    def create_Tele_User(cls, tele_user_id:int):
        oes_user = TUser.getOESUserFromUserID(tele_user_id)
        if oes_user == None:
            print("WARNING: Could not retrieve OES user for TUser with user_id:", tele_user_id)
        return TUser(tele_user_id, oes_user)

    @classmethod
    def getOESUserFromUserID(cls, tele_user_id:int):
        tele_user_id = str(tele_user_id)
        filter_comparison = operations_ecosys_pb2.Filter(
            comparison=operations_ecosys_pb2.Filter.EQUAL, 
            value=tele_user_id
        )
        query_filters = [operations_ecosys_pb2.UserFilter(
            field=operations_ecosys_pb2.UserFilter.TELEGRAM_USER_ID,
            comparisons=filter_comparison,
        )]
        user_query = operations_ecosys_pb2.UserQuery(filters=query_filters,limit=1)
        users = user_client.get_users(user_query)
        return users[0] if len(users) > 0 else None

    @classmethod
    def ifUserExists(cls, tele_user_id: str) -> bool:
        return getOESUserFromUserID(tele_user_id) != None

    # This method is important because it updates the database that the user "logged in".
    # Or at least, it should.
    def login(self, update:Update) -> bool:
        if self.oes_user == None:
            raise Exception("Unable to login TeleUser to DB because oes_user is None.")
            return
        self.oes_user.tele_user_id = self.tele_user_id
        self.oes_user.telegram_handle = update.effective_user.username
        self.oes_user.name = update.effective_user.first_name + " " if update.effective_user.first_name != None else ""
        self.oes_user.name += update.effective_user.last_name if update.effective_user.last_name != None else ""
        print("User Login:")
        print(self.oes_user)
        if not user_client.update_user(self.oes_user):
            print("Error updating and logging in user")
            return False
        return True

    # # JW does not recommend this because tele handles can change.
    # @classmethod
    # def get_from_handle(user_tele_handle: str) -> operations_ecosys_pb2.User:
    #     filter_comparison = operations_ecosys_pb2.Filter(
    #         comparison=operations_ecosys_pb2.Filter.EQUAL, 
    #         value=user_tele_handle,
    #     )

    #     query_filters = [operations_ecosys_pb2.UserFilter(
    #         field=operations_ecosys_pb2.UserFilter.TELEGRAM_HANDLE,
    #         comparisons=filter_comparison,
    #     )]

    #     user_query = operations_ecosys_pb2.UserQuery(filters=query_filters,limit=1)
    #     users = user_client.get_users(user_query)
        
    #     return users[0] if len(users) >0 else None

    

