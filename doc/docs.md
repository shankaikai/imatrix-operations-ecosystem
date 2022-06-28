# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [operations_ecosys.proto](#operations_ecosys-proto)
    - [AIFSBroadcastRecipient](#operations_ecosys-AIFSBroadcastRecipient)
    - [AIFSClientRoster](#operations_ecosys-AIFSClientRoster)
    - [AvailabilityFilter](#operations_ecosys-AvailabilityFilter)
    - [AvailabilityQuery](#operations_ecosys-AvailabilityQuery)
    - [Broadcast](#operations_ecosys-Broadcast)
    - [BroadcastFilter](#operations_ecosys-BroadcastFilter)
    - [BroadcastQuery](#operations_ecosys-BroadcastQuery)
    - [BroadcastRecipient](#operations_ecosys-BroadcastRecipient)
    - [BroadcastResponse](#operations_ecosys-BroadcastResponse)
    - [BulkRosters](#operations_ecosys-BulkRosters)
    - [Client](#operations_ecosys-Client)
    - [ClientFilter](#operations_ecosys-ClientFilter)
    - [ClientQuery](#operations_ecosys-ClientQuery)
    - [ClientResponse](#operations_ecosys-ClientResponse)
    - [EmployeeEvaluation](#operations_ecosys-EmployeeEvaluation)
    - [EmployeeEvaluationResponse](#operations_ecosys-EmployeeEvaluationResponse)
    - [Filter](#operations_ecosys-Filter)
    - [OrderByBroadcast](#operations_ecosys-OrderByBroadcast)
    - [OrderByClient](#operations_ecosys-OrderByClient)
    - [OrderByQuery](#operations_ecosys-OrderByQuery)
    - [OrderByRoster](#operations_ecosys-OrderByRoster)
    - [OrderByUser](#operations_ecosys-OrderByUser)
    - [Response](#operations_ecosys-Response)
    - [Roster](#operations_ecosys-Roster)
    - [RosterAssignement](#operations_ecosys-RosterAssignement)
    - [RosterAssignmentResponse](#operations_ecosys-RosterAssignmentResponse)
    - [RosterFilter](#operations_ecosys-RosterFilter)
    - [RosterQuery](#operations_ecosys-RosterQuery)
    - [RosterResponse](#operations_ecosys-RosterResponse)
    - [User](#operations_ecosys-User)
    - [UserFilter](#operations_ecosys-UserFilter)
    - [UserQuery](#operations_ecosys-UserQuery)
    - [UsersResponse](#operations_ecosys-UsersResponse)
  
    - [AvailabilityFilter.Field](#operations_ecosys-AvailabilityFilter-Field)
    - [Broadcast.BroadcastType](#operations_ecosys-Broadcast-BroadcastType)
    - [Broadcast.UrgencyType](#operations_ecosys-Broadcast-UrgencyType)
    - [BroadcastFilter.Field](#operations_ecosys-BroadcastFilter-Field)
    - [ClientFilter.Field](#operations_ecosys-ClientFilter-Field)
    - [Filter.Comparisons](#operations_ecosys-Filter-Comparisons)
    - [OrderBy](#operations_ecosys-OrderBy)
    - [Response.Type](#operations_ecosys-Response-Type)
    - [Roster.Status](#operations_ecosys-Roster-Status)
    - [RosterFilter.Field](#operations_ecosys-RosterFilter-Field)
    - [User.UserType](#operations_ecosys-User-UserType)
    - [UserFilter.Field](#operations_ecosys-UserFilter-Field)
  
    - [AdminServices](#operations_ecosys-AdminServices)
    - [BroadcastServices](#operations_ecosys-BroadcastServices)
    - [RosterServices](#operations_ecosys-RosterServices)
  
- [Scalar Value Types](#scalar-value-types)



<a name="operations_ecosys-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## operations_ecosys.proto



<a name="operations_ecosys-AIFSBroadcastRecipient"></a>

### AIFSBroadcastRecipient



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recipient | [BroadcastRecipient](#operations_ecosys-BroadcastRecipient) | repeated |  |
| aifs_id | [int64](#int64) |  |  |






<a name="operations_ecosys-AIFSClientRoster"></a>

### AIFSClientRoster



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aifs_client_roster_id | [int64](#int64) |  |  |
| client | [Client](#operations_ecosys-Client) |  |  |
| patrol_order | [int64](#int64) |  |  |






<a name="operations_ecosys-AvailabilityFilter"></a>

### AvailabilityFilter
Filter the types of availabilty to be returned.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [AvailabilityFilter.Field](#operations_ecosys-AvailabilityFilter-Field) |  |  |
| comparisons | [Filter](#operations_ecosys-Filter) |  |  |






<a name="operations_ecosys-AvailabilityQuery"></a>

### AvailabilityQuery
TODO: Find out if the user should be available
throughout the whole duration.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_time | [string](#string) |  |  |
| end_time | [string](#string) |  | If end time is null, the time period is taken to be from the start time &#43; 12h |
| limit | [int64](#int64) |  | Limit the number of objects being returned. If only 5 objects should be shown, limit = 5; |
| skip | [int64](#int64) |  | Skip n rows from the database |
| filters | [AvailabilityFilter](#operations_ecosys-AvailabilityFilter) | repeated | The client has no current need to fill this up. |
| order_by | [OrderByQuery](#operations_ecosys-OrderByQuery) |  |  |






<a name="operations_ecosys-Broadcast"></a>

### Broadcast
The default fields of a broadcast


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| broadcast_id | [int64](#int64) |  | Broadcast IDs are only useful for the backend database. |
| type | [Broadcast.BroadcastType](#operations_ecosys-Broadcast-BroadcastType) |  |  |
| content | [string](#string) |  |  |
| creation_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| deadline | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| creator | [User](#operations_ecosys-User) |  | The whole user fields does not need to be filled, as long as the user is identifiable. |
| recipients | [AIFSBroadcastRecipient](#operations_ecosys-AIFSBroadcastRecipient) | repeated |  |
| urgency | [Broadcast.UrgencyType](#operations_ecosys-Broadcast-UrgencyType) |  |  |






<a name="operations_ecosys-BroadcastFilter"></a>

### BroadcastFilter
Filter the types of broadcasts to be returned.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [BroadcastFilter.Field](#operations_ecosys-BroadcastFilter-Field) |  |  |
| comparisons | [Filter](#operations_ecosys-Filter) |  |  |






<a name="operations_ecosys-BroadcastQuery"></a>

### BroadcastQuery
Get specific types users as specified in the Filter. 
If one wants to get all objects, leave filters empty. 
A default limit of 10 will be used if the field is empty.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [BroadcastFilter](#operations_ecosys-BroadcastFilter) | repeated |  |
| limit | [int64](#int64) |  | Limit the number of objects being returned. If only 5 objects should be shown, limit = 5; |
| skip | [int64](#int64) |  | Skip n rows from the database |
| order_by | [OrderByBroadcast](#operations_ecosys-OrderByBroadcast) |  | Order the queries, by default the order is desc by creation date |






<a name="operations_ecosys-BroadcastRecipient"></a>

### BroadcastRecipient



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| broadcast_recipients_id | [int64](#int64) |  |  |
| recipient | [User](#operations_ecosys-User) |  |  |
| acknowledged | [bool](#bool) |  |  |
| rejected | [bool](#bool) |  |  |
| last_replied | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| aifs_id | [int64](#int64) |  |  |






<a name="operations_ecosys-BroadcastResponse"></a>

### BroadcastResponse
Passing around multiple broadcasts in one message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| broadcast | [Broadcast](#operations_ecosys-Broadcast) |  |  |






<a name="operations_ecosys-BulkRosters"></a>

### BulkRosters



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rosters | [Roster](#operations_ecosys-Roster) | repeated |  |






<a name="operations_ecosys-Client"></a>

### Client
All clients who hire the company to protect their compounds. 
TODO: decide if we want to have an inactive field for clients
      so that we don&#39;t need to delete cleints completely.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| abbreviation | [string](#string) |  |  |
| email | [string](#string) |  |  |
| address | [string](#string) |  |  |
| postal_code | [int64](#int64) |  |  |
| phone_number | [string](#string) |  |  |






<a name="operations_ecosys-ClientFilter"></a>

### ClientFilter
Filter the types of clients to be returned.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [ClientFilter.Field](#operations_ecosys-ClientFilter-Field) |  |  |
| comparisons | [Filter](#operations_ecosys-Filter) |  |  |






<a name="operations_ecosys-ClientQuery"></a>

### ClientQuery
Get specific types clients as specified in the Filter. 
If one wants to get all objects, leave filters empty. 
A default limit of 10 will be used if the field is empty.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [ClientFilter](#operations_ecosys-ClientFilter) | repeated |  |
| limit | [int64](#int64) |  | Limit the number of objects being returned. If only 5 objects should be shown, limit = 5; |
| skip | [int64](#int64) |  | Skip n rows from the database |
| order_by | [OrderByClient](#operations_ecosys-OrderByClient) |  | Order the queries, by default the order is desc by creation date |






<a name="operations_ecosys-ClientResponse"></a>

### ClientResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| client | [Client](#operations_ecosys-Client) |  |  |






<a name="operations_ecosys-EmployeeEvaluation"></a>

### EmployeeEvaluation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| employee | [User](#operations_ecosys-User) |  |  |
| employee_score | [float](#float) |  |  |
| is_available | [bool](#bool) |  | This is used if it is needed to determine if the employee is available within a previously specified time. |






<a name="operations_ecosys-EmployeeEvaluationResponse"></a>

### EmployeeEvaluationResponse
ROSTERING EMPLOYEE EVALUATION MESSAGES                    *
**************************************************************************


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| employee | [EmployeeEvaluation](#operations_ecosys-EmployeeEvaluation) |  |  |






<a name="operations_ecosys-Filter"></a>

### Filter
This is used to indicate what kind of objects should be returned that
fit this critera. 
For example, if one wishes to get all broadcasts that have more than 
one recipient. They might put the comparison as GREATER, value = 1;
The field to be compared with is in the corresponding XXXFilter message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| comparison | [Filter.Comparisons](#operations_ecosys-Filter-Comparisons) |  |  |
| value | [string](#string) |  |  |






<a name="operations_ecosys-OrderByBroadcast"></a>

### OrderByBroadcast



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [BroadcastFilter.Field](#operations_ecosys-BroadcastFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-OrderByClient"></a>

### OrderByClient



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [ClientFilter.Field](#operations_ecosys-ClientFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-OrderByQuery"></a>

### OrderByQuery
TODO change name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [AvailabilityFilter.Field](#operations_ecosys-AvailabilityFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-OrderByRoster"></a>

### OrderByRoster



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [RosterFilter.Field](#operations_ecosys-RosterFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-OrderByUser"></a>

### OrderByUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [UserFilter.Field](#operations_ecosys-UserFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-Response"></a>

### Response
Generic reponses to add or update requests


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Response.Type](#operations_ecosys-Response-Type) |  |  |
| error_message | [string](#string) |  |  |
| primary_key | [int64](#int64) |  | Return any pk of the row that the query modified |






<a name="operations_ecosys-Roster"></a>

### Roster
The default fields of a roster


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rostering_id | [int64](#int64) |  | Roster IDs are only useful for the backend database. |
| aifs_id | [int64](#int64) |  |  |
| start_time | [string](#string) |  | Format must be in: YYYY-MM-DD HH:MM:SS |
| end_time | [string](#string) |  |  |
| clients | [AIFSClientRoster](#operations_ecosys-AIFSClientRoster) | repeated |  |
| guard_assigned | [RosterAssignement](#operations_ecosys-RosterAssignement) | repeated |  |
| status | [Roster.Status](#operations_ecosys-Roster-Status) |  | The roster is not actually in the database |






<a name="operations_ecosys-RosterAssignement"></a>

### RosterAssignement



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| roster_assignment_id | [int64](#int64) |  |  |
| guard_assigned | [EmployeeEvaluation](#operations_ecosys-EmployeeEvaluation) |  | The whole user fields does not need to be filled, as long as the user is identifiable. |
| custom_start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| custom_end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| confirmed | [bool](#bool) |  |  |
| attended | [bool](#bool) |  |  |
| attendance_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| is_assigned | [bool](#bool) |  | If the assignment is part of the current assignment or if it was previously assigned and is now removed |
| rejected | [bool](#bool) |  |  |






<a name="operations_ecosys-RosterAssignmentResponse"></a>

### RosterAssignmentResponse
Passing around multiple roster assignments in one message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| roster_assignment | [RosterAssignement](#operations_ecosys-RosterAssignement) |  |  |






<a name="operations_ecosys-RosterFilter"></a>

### RosterFilter
Filter the types of rosters to be returned.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [RosterFilter.Field](#operations_ecosys-RosterFilter-Field) |  |  |
| comparisons | [Filter](#operations_ecosys-Filter) |  |  |






<a name="operations_ecosys-RosterQuery"></a>

### RosterQuery
Get specific types rosters as specified in the Filter. 
If one wants to get all objects, leave filters empty. 
A default limit of 10 will be used if the field is empty.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [RosterFilter](#operations_ecosys-RosterFilter) | repeated |  |
| limit | [int64](#int64) |  | Limit the number of objects being returned. If only 5 objects should be shown, limit = 5; |
| skip | [int64](#int64) |  | Skip n rows from the database |
| order_by | [OrderByRoster](#operations_ecosys-OrderByRoster) |  | Order the queries, by default the order is desc by creation date |






<a name="operations_ecosys-RosterResponse"></a>

### RosterResponse
Passing around multiple roster in one message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| roster | [Roster](#operations_ecosys-Roster) |  |  |






<a name="operations_ecosys-User"></a>

### User
All users who use the operations ecosystem.
TODO: decide if we want to have an inactive field for users
      so that we don&#39;t need to delete users completely.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  |  |
| user_type | [User.UserType](#operations_ecosys-User-UserType) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| phone_number | [string](#string) |  |  |
| telegram_handle | [string](#string) |  |  |
| user_security_img | [string](#string) |  |  |
| is_part_timer | [bool](#bool) |  |  |
| tele_chat_id | [int64](#int64) |  |  |






<a name="operations_ecosys-UserFilter"></a>

### UserFilter
Filter the types of users to be returned.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [UserFilter.Field](#operations_ecosys-UserFilter-Field) |  |  |
| comparisons | [Filter](#operations_ecosys-Filter) |  |  |






<a name="operations_ecosys-UserQuery"></a>

### UserQuery
Get specific types users as specified in the Filter. 
If one wants to get all objects, leave filters empty. 
A default limit of 10 will be used if the field is empty.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [UserFilter](#operations_ecosys-UserFilter) | repeated |  |
| limit | [int64](#int64) |  | Limit the number of objects being returned. If only 5 objects should be shown, limit = 5; |
| skip | [int64](#int64) |  | Skip n rows from the database |
| order_by | [OrderByUser](#operations_ecosys-OrderByUser) |  | Order the queries, by default the order is desc by creation date |






<a name="operations_ecosys-UsersResponse"></a>

### UsersResponse
Passing around multiple users in one message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| user | [User](#operations_ecosys-User) |  |  |





 


<a name="operations_ecosys-AvailabilityFilter-Field"></a>

### AvailabilityFilter.Field
More fields can be added in the future.

| Name | Number | Description |
| ---- | ------ | ----------- |
| AVAILABILITY_ID | 0 |  |
| WEEK | 1 | Week of the year |
| YEAR | 2 |  |
| GUARD_ID | 3 | Note: single user ID |
| SUN | 4 | Day of the week to look at The values for these fields in Filter comparisons should be an empty string |
| MON | 5 |  |
| TUES | 6 |  |
| WED | 7 |  |
| THURS | 8 |  |
| FRI | 9 |  |
| SAT | 10 |  |
| NEXT_SUN | 11 |  |



<a name="operations_ecosys-Broadcast-BroadcastType"></a>

### Broadcast.BroadcastType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ANNOUNCEMENT | 0 |  |
| ASSIGNMENT | 1 |  |



<a name="operations_ecosys-Broadcast-UrgencyType"></a>

### Broadcast.UrgencyType


| Name | Number | Description |
| ---- | ------ | ----------- |
| LOW | 0 |  |
| MEDIUM | 1 |  |
| HIGH | 2 |  |



<a name="operations_ecosys-BroadcastFilter-Field"></a>

### BroadcastFilter.Field
More fields can be added in the future.

| Name | Number | Description |
| ---- | ------ | ----------- |
| BROADCAST_ID | 0 |  |
| TYPE | 1 |  |
| CONTENT | 2 |  |
| CREATION_DATE | 3 |  |
| DEADLINE | 4 |  |
| CREATOR_ID | 5 | Creator and recipient values should be the user id of these users |
| RECEIPEIENT_ID | 6 | Note: Single recipient |
| NUM_RECEIPIENTS | 7 |  |
| URGENCY | 8 |  |
| AIFS_ID | 9 |  |
| BROADCAST_RECIPIENT_TABLE_ID | 10 |  |



<a name="operations_ecosys-ClientFilter-Field"></a>

### ClientFilter.Field
More fields can be added in the future.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CLIENT_ID | 0 |  |



<a name="operations_ecosys-Filter-Comparisons"></a>

### Filter.Comparisons


| Name | Number | Description |
| ---- | ------ | ----------- |
| GREATER | 0 |  |
| GREATER_EQ | 1 |  |
| EQUAL | 2 |  |
| LESSER_EQ | 3 |  |
| LESSER | 4 |  |
| CONTAINS | 5 |  |
| IN | 6 |  |
| NOT_IN | 7 |  |



<a name="operations_ecosys-OrderBy"></a>

### OrderBy


| Name | Number | Description |
| ---- | ------ | ----------- |
| ASC | 0 |  |
| DESC | 1 |  |



<a name="operations_ecosys-Response-Type"></a>

### Response.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| ACK | 0 |  |
| ERROR | 1 |  |



<a name="operations_ecosys-Roster-Status"></a>

### Roster.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| IS_DEFAULT | 0 |  |
| PENDING | 1 |  |
| CONFIRMED | 2 |  |
| REJECTED | 3 |  |



<a name="operations_ecosys-RosterFilter-Field"></a>

### RosterFilter.Field
More fields can be added in the future.

| Name | Number | Description |
| ---- | ------ | ----------- |
| ROSTER_ID | 0 |  |
| ROSTER_ASSIGNMENT_ID | 1 |  |
| ROSTER_AIFS_CLIENT_ID | 2 |  |
| AIFS_ID | 3 |  |
| GUARD_ASSIGNED_ID | 4 | Note: single user ID |
| CLIENT_ID | 5 | Note: single client ID |
| GUARD_ASSIGNMENT_CONFIRMATION | 6 |  |
| GUARD_ASSIGNMENT_ATTENDED | 7 |  |
| START_TIME | 8 |  |
| END_TIME | 9 |  |
| IS_ASSIGNED | 10 |  |
| DEFAULT_ROSTERING_DAY_OF_WEEK | 11 |  |
| GUARD_ASSIGNMENT_REJECTION | 12 |  |



<a name="operations_ecosys-User-UserType"></a>

### User.UserType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ISPECIALIST | 0 |  |
| SECURITY_GUARD | 1 |  |
| CONTROLLER | 2 |  |
| MANAGER | 3 |  |



<a name="operations_ecosys-UserFilter-Field"></a>

### UserFilter.Field
More fields can be added in the future.

| Name | Number | Description |
| ---- | ------ | ----------- |
| USER_ID | 0 |  |
| TYPE | 1 |  |
| NAME | 2 |  |
| EMAIL | 3 |  |
| PHONE_NUMBER | 4 |  |
| TELEGRAM_HANDLE | 5 |  |
| IS_PART_TIMER | 6 |  |


 

 


<a name="operations_ecosys-AdminServices"></a>

### AdminServices


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddUser | [User](#operations_ecosys-User) | [Response](#operations_ecosys-Response) | User |
| UpdateUser | [User](#operations_ecosys-User) | [Response](#operations_ecosys-Response) |  |
| DeleteUser | [User](#operations_ecosys-User) | [Response](#operations_ecosys-Response) |  |
| FindUsers | [UserQuery](#operations_ecosys-UserQuery) | [UsersResponse](#operations_ecosys-UsersResponse) stream | TODO change user response to have user scoring and stuff |
| AddClient | [Client](#operations_ecosys-Client) | [Response](#operations_ecosys-Response) | Client |
| UpdateClient | [Client](#operations_ecosys-Client) | [Response](#operations_ecosys-Response) |  |
| DeleteClient | [Client](#operations_ecosys-Client) | [Response](#operations_ecosys-Response) |  |
| FindClients | [ClientQuery](#operations_ecosys-ClientQuery) | [ClientResponse](#operations_ecosys-ClientResponse) stream |  |


<a name="operations_ecosys-BroadcastServices"></a>

### BroadcastServices


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddBroadcast | [Broadcast](#operations_ecosys-Broadcast) | [Response](#operations_ecosys-Response) |  |
| UpdateBroadcast | [Broadcast](#operations_ecosys-Broadcast) | [Response](#operations_ecosys-Response) | Note that this update does not update the broadcast&#39;s recipient&#39;s inner status such as the acknowledgement or rejection status but only if the recipient is part of the broadcast. |
| DeleteBroadcast | [Broadcast](#operations_ecosys-Broadcast) | [Response](#operations_ecosys-Response) |  |
| FindBroadcasts | [BroadcastQuery](#operations_ecosys-BroadcastQuery) | [BroadcastResponse](#operations_ecosys-BroadcastResponse) stream |  |
| UpdateBroadcastRecipient | [BroadcastRecipient](#operations_ecosys-BroadcastRecipient) | [Response](#operations_ecosys-Response) | Updating of broadcast recipients |


<a name="operations_ecosys-RosterServices"></a>

### RosterServices


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddRoster | [BulkRosters](#operations_ecosys-BulkRosters) | [Response](#operations_ecosys-Response) | Add multiple rosters for different AIFS at the same time |
| UpdateRoster | [BulkRosters](#operations_ecosys-BulkRosters) | [Response](#operations_ecosys-Response) | Note that this update does not update the roster&#39;s guard&#39;s inner status such as the acknowledgement or attended status but only if the guard is part of the roster. |
| DeleteRoster | [Roster](#operations_ecosys-Roster) | [Response](#operations_ecosys-Response) |  |
| FindRosters | [RosterQuery](#operations_ecosys-RosterQuery) | [RosterResponse](#operations_ecosys-RosterResponse) stream |  |
| GetAvailableUsers | [AvailabilityQuery](#operations_ecosys-AvailabilityQuery) | [EmployeeEvaluationResponse](#operations_ecosys-EmployeeEvaluationResponse) stream |  |
| FindRosterAssignments | [RosterQuery](#operations_ecosys-RosterQuery) | [RosterAssignmentResponse](#operations_ecosys-RosterAssignmentResponse) stream | Specifically for the roster assignments |
| UpdateRosterAssignment | [RosterAssignement](#operations_ecosys-RosterAssignement) | [Response](#operations_ecosys-Response) | Updates the individual roster assignment |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
