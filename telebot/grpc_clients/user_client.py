from __future__ import print_function

from grpc_clients import utils

import grpc

from typing import List

from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2

def get_webapp_nonce(user:operations_ecosys_pb2.User) -> str:
    stub = get_admin_stub()
    wanonce_res = stub.GetWANonce(user)
    if wanonce_res.response.type == operations_ecosys_pb2.Response.ACK:
        return wanonce_res.nonce
    else:
        print("Error obtaining webapp nonce:", wanonce_res)

def get_regcode(req:operations_ecosys_pb2.RegistrationCodeRequest) -> str:
    stub = get_admin_stub()
    regcode_res = stub.GetRegistrationCode(req)
    if regcode_res.response.type == operations_ecosys_pb2.Response.ACK:
        return regcode_res.code
    else:
        print("Error obtaining registration code:", regcode_res)

def get_users(user_query) -> List[operations_ecosys_pb2.User]:
    stub = get_admin_stub()
    userResponses = stub.FindUsers(user_query)
    
    users = []

    for userRes in userResponses:
        if userRes.response.type == operations_ecosys_pb2.Response.ACK:
            users.append(userRes.user)
        else:
            print("get_users", userRes)

    return users

def delete_user(user_id:int) -> bool:
    # return true if deletion was successful and false if otherwise
    user_id = str(user_id)
    print(user_id)
    filter_comparison = operations_ecosys_pb2.Filter(
        comparison=operations_ecosys_pb2.Filter.EQUAL, 
        value=user_id
    )
    query_filters = [operations_ecosys_pb2.UserFilter(
        field=operations_ecosys_pb2.UserFilter.USER_ID,
        comparisons=filter_comparison,
    )]
    user_query = operations_ecosys_pb2.UserQuery(filters=query_filters,limit=1)
    users = get_users(user_query)

    stub = get_admin_stub()
    res = stub.DeleteUser(users[0])
    
    if res.type == operations_ecosys_pb2.Response.ACK:
        return True
    else:
        print("delete_user ERROR:", delRes)
        return False

def lookup_users(telegram_username=None, name=None, \
        phone_num=None, email=None, telegram_user_id=None) -> List[operations_ecosys_pb2.User]:
    userFilters = []
    if telegram_username:
        userFilters.append(operations_ecosys_pb2.UserFilter(
            field=operations_ecosys_pb2.UserFilter.Field.TELEGRAM_HANDLE, 
            comparisons=operations_ecosys_pb2.Filter(
                comparison=operations_ecosys_pb2.Filter.Comparisons.CONTAINS, 
                value=telegram_username)
            )
        )
    if name:
        userFilters.append(operations_ecosys_pb2.UserFilter(
            field=operations_ecosys_pb2.UserFilter.Field.NAME, 
            comparisons=operations_ecosys_pb2.Filter(
                comparison=operations_ecosys_pb2.Filter.Comparisons.CONTAINS, 
                value=name)
            )
        )
    if phone_num:
        userFilters.append(operations_ecosys_pb2.UserFilter(
            field=operations_ecosys_pb2.UserFilter.Field.PHONE_NUMBER, 
            comparisons=operations_ecosys_pb2.Filter(
                comparison=operations_ecosys_pb2.Filter.Comparisons.EQUAL, 
                value=phone_num)
            )
        )
    if email:
        userFilters.append(operations_ecosys_pb2.UserFilter(
            field=operations_ecosys_pb2.UserFilter.Field.EMAIL, 
            comparisons=operations_ecosys_pb2.Filter(
                comparison=operations_ecosys_pb2.Filter.Comparisons.EQUAL, 
                value=email)
            )
        )
    if telegram_user_id:
        userFilters.append(operations_ecosys_pb2.UserFilter(
            field=operations_ecosys_pb2.UserFilter.Field.TELEGRAM_USER_ID, 
            comparisons=operations_ecosys_pb2.Filter(
                comparison=operations_ecosys_pb2.Filter.Comparisons.EQUAL, 
                value=telegram_user_id)
            )
        )
    userQuery = operations_ecosys_pb2.UserQuery(filters=userFilters)
    users = get_users(userQuery)
    return users

def update_user(user) -> bool:
    stub = get_admin_stub()
    res = stub.UpdateUser(user)
    return res.type == operations_ecosys_pb2.Response.ACK


# For registration
def check_reg_code(registration_code:str) -> str:
    req = operations_ecosys_pb2.RegistrationCode(code=registration_code)
    stub = get_admin_stub()
    security_string_res = stub.CheckRegistrationCode(req)
    if security_string_res.response.type == operations_ecosys_pb2.Response.ACK:
        return security_string_res.security_string
    else:
        print("Error obtaining registration code:", security_string_res)
        return None

# For linking
def link_via_telehandle(tele_handle:str, tele_user_id:int) -> bool:
    filter_comparison = operations_ecosys_pb2.Filter(
             comparison=operations_ecosys_pb2.Filter.EQUAL, 
             value=tele_handle,
        )

    query_filters = [operations_ecosys_pb2.UserFilter(
             field=operations_ecosys_pb2.UserFilter.TELEGRAM_HANDLE,
             comparisons=filter_comparison,
        )]

    user_query = operations_ecosys_pb2.UserQuery(filters=query_filters,limit=1)
    users = get_users(user_query)

    if len(users) != 1:
        return False

    users[0].tele_user_id = tele_user_id

    return update_user(users[0])
    

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


def get_admin_stub() -> operations_ecosys_pb2_grpc.AdminServicesStub:
    channel = grpc.insecure_channel('{}:{}'.format(utils.WEB_SERVER_ADDR, utils.WEB_SERVER_PORT))
    stub = operations_ecosys_pb2_grpc.AdminServicesStub(channel)
    return stub

