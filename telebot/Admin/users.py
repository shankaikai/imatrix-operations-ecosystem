from __future__ import print_function

from GrpcClient import user_client

import logging
import random
import grpc

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2


def getUserFromHandle(user_tele_handle: str) -> operations_ecosys_pb2.User:
    filter_comparison = operations_ecosys_pb2.Filter(
        comparison=operations_ecosys_pb2.Filter.EQUAL, 
        value=user_tele_handle,
    )

    query_filters = [operations_ecosys_pb2.UserFilter(
        field=operations_ecosys_pb2.UserFilter.TELEGRAM_HANDLE,
        comparisons=filter_comparison,
    )]

    user_query = operations_ecosys_pb2.UserQuery(filters=query_filters,limit=1)
    users = user_client.get_users(user_query)
    
    return users[0] if len(users) >0 else None

def updateUserChatId(user_tele_handle: str, chat_id: int) -> bool:
    print("updateUserChatId", user_tele_handle, chat_id)
    user = getUserFromHandle(user_tele_handle)
    print("updateUserChatId", user)
    if user == None:
        return False
    user.tele_chat_id = chat_id
    return user_client.update_user(user)
