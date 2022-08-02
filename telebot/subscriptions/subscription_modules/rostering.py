from typing import Tuple
from telegram import Chat, InlineKeyboardMarkup, InlineKeyboardButton, Message, ParseMode
from telegram.ext import Updater, ContextTypes, CallbackContext
from google.protobuf.timestamp_pb2 import Timestamp

from ..subscription_message import SubscriptionMessage
from Protos import operations_ecosys_pb2_grpc, operations_ecosys_pb2
from grpc_clients import rostering_client
from .. import reminders

from datetime import datetime

IDENTIFIER = "rostering"
ACCEPT_FEATURE = "_accept"
REJECT_FEATURE = "_reject"

# Boolean reflects if the handle supports the msg; str returns some comment/text
def handle_sub_message(sub_msg:SubscriptionMessage) -> Tuple[bool, str]:
    print(sub_msg.feature, IDENTIFIER + ACCEPT_FEATURE, IDENTIFIER + REJECT_FEATURE)
    if IDENTIFIER not in sub_msg.feature:
        return False, None
    if  IDENTIFIER + ACCEPT_FEATURE in sub_msg.feature:
        print("accept handler")
        if acknowledge_roster(sub_msg.pb_msg_id):
            return True, "Acknowledged"
        return False, "Error when updating roster accept."
    if IDENTIFIER + REJECT_FEATURE in sub_msg.feature:
        print("reject handler")
        if reject_roster(sub_msg.pb_msg_id):
            return True, "Rejected"
        return False, "Error when updating roster reject. "

def send_roster_message(updater : Updater, chat_id: int, 
        roster_assignment: operations_ecosys_pb2.RosterAssignement, 
        roster: operations_ecosys_pb2.Roster):

    print("send_roster_message", chat_id, roster_assignment.roster_assignment_id)
    
    keyboard_markup = InlineKeyboardMarkup(
            [[
                InlineKeyboardButton(
                    text="Acknowledge",
                    callback_data=str(SubscriptionMessage(
                        IDENTIFIER + ACCEPT_FEATURE, 
                        roster_assignment.roster_assignment_id, 
                        chat_id
                    ))
                ), 

                InlineKeyboardButton(
                    text="Reject",
                    callback_data=str(SubscriptionMessage(
                        IDENTIFIER + REJECT_FEATURE, 
                        roster_assignment.roster_assignment_id, 
                        chat_id
                    ))                
                ),
            ]]
    )

    format_data = "%d %b %Y %#I:%M%p" 

    message = """You have been assigned to a shift, please click to accept/reject.
    <b><u>Shift Details</u></b>
    <b>AIFS:</b> {}
    <b>Start Time:</b> {}
    <b>End Time:</b> {}""".format(roster.aifs_id,
    roster_assignment.custom_start_time.ToDatetime().strftime(format_data),
    roster_assignment.custom_end_time.ToDatetime().strftime(format_data),)

    msg = updater.bot.send_message(chat_id=chat_id, text=message, parse_mode=ParseMode.HTML, reply_markup=keyboard_markup)
    setup_roster_reminder(updater, chat_id, msg, roster_assignment.roster_assignment_id)

def acknowledge_roster(roster_assignment_id: int) -> bool:
    roster_assignment = operations_ecosys_pb2.RosterAssignement(
        roster_assignment_id = roster_assignment_id, 
        confirmed = True,
        attended = False,
        is_assigned = True, 
        rejected = False,
    )
    updated = rostering_client.update_rostering_assignment(roster_assignment)
    return updated


def reject_roster(roster_assignment_id: int) -> bool:
    roster_assignment = operations_ecosys_pb2.RosterAssignement(
        roster_assignment_id = roster_assignment_id, 
        confirmed = False,
        attended = False,
        is_assigned = True, 
        rejected = True,
    )
    updated = rostering_client.update_rostering_assignment(roster_assignment)
    return updated



def setup_roster_reminder(updater : Updater, chat_id: int, msg: Message, roster_assignment_id: int):
    updater.job_queue.run_once(
        callback=remind_roster, 
        when=reminders.ROSTER_REMINDER_SECONDS, 
        context={
            reminders.MSG_KEY:msg,
            reminders.CHAT_ID_KEY: chat_id, 
            reminders.ROSTER_ASGN_ID_KEY: roster_assignment_id,
            reminders.UPDATER_KEY: updater,
        },
    )

def remind_roster(context: CallbackContext):
    # print("remind_roster", context.job.context)

    # If already has a reply, do nothing
    roster_asgn = rostering_client.get_roster_assignment(context.job.context[reminders.ROSTER_ASGN_ID_KEY])

    # If the row in the db is already gone, it could have been deleted. Ignore
    if roster_asgn is None:
        return
    
    if roster_asgn.confirmed or roster_asgn.rejected:
        print("Roster already acknowledged")
        return

    # Resend the same message + keyboard
    reminders.resend_message(context, setup_roster_reminder, context.job.context[reminders.ROSTER_ASGN_ID_KEY])
    