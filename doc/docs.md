# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [http_webapp.proto](#http_webapp-proto)
    - [HTTPAssignmentResponse](#http_webapp-HTTPAssignmentResponse)
    - [HTTPAssignmentsGetRequest](#http_webapp-HTTPAssignmentsGetRequest)
    - [HTTPMessage](#http_webapp-HTTPMessage)
    - [HTTPRegistrationFormRequest](#http_webapp-HTTPRegistrationFormRequest)
    - [HTTPReportPostRequest](#http_webapp-HTTPReportPostRequest)
    - [HTTPRosterResponse](#http_webapp-HTTPRosterResponse)
  
    - [WebAppServices](#http_webapp-WebAppServices)
  
- [iot_prototype.proto](#iot_prototype-proto)
    - [CpuTemp](#gate_prototype-CpuTemp)
    - [CpuTempState](#gate_prototype-CpuTempState)
    - [FireAlarm](#gate_prototype-FireAlarm)
    - [FireAlarmState](#gate_prototype-FireAlarmState)
    - [Gate](#gate_prototype-Gate)
    - [GateState](#gate_prototype-GateState)
  
    - [FireAlarmState.AlarmState](#gate_prototype-FireAlarmState-AlarmState)
    - [GateState.GatePosition](#gate_prototype-GateState-GatePosition)
  
    - [IotControlPrototypeService](#gate_prototype-IotControlPrototypeService)
  
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
    - [Camera](#operations_ecosys-Camera)
    - [CameraIot](#operations_ecosys-CameraIot)
    - [CameraIotFilter](#operations_ecosys-CameraIotFilter)
    - [CameraIotQuery](#operations_ecosys-CameraIotQuery)
    - [CameraIotResponse](#operations_ecosys-CameraIotResponse)
    - [Client](#operations_ecosys-Client)
    - [ClientFilter](#operations_ecosys-ClientFilter)
    - [ClientQuery](#operations_ecosys-ClientQuery)
    - [ClientResponse](#operations_ecosys-ClientResponse)
    - [EmployeeEvaluation](#operations_ecosys-EmployeeEvaluation)
    - [EmployeeEvaluationResponse](#operations_ecosys-EmployeeEvaluationResponse)
    - [Filter](#operations_ecosys-Filter)
    - [FullUser](#operations_ecosys-FullUser)
    - [IncidentReport](#operations_ecosys-IncidentReport)
    - [IncidentReportContent](#operations_ecosys-IncidentReportContent)
    - [IncidentReportFilter](#operations_ecosys-IncidentReportFilter)
    - [IncidentReportQuery](#operations_ecosys-IncidentReportQuery)
    - [IncidentReportResponse](#operations_ecosys-IncidentReportResponse)
    - [LoginRequest](#operations_ecosys-LoginRequest)
    - [OrderByBroadcast](#operations_ecosys-OrderByBroadcast)
    - [OrderByCameraIot](#operations_ecosys-OrderByCameraIot)
    - [OrderByClient](#operations_ecosys-OrderByClient)
    - [OrderByIncidentReport](#operations_ecosys-OrderByIncidentReport)
    - [OrderByQuery](#operations_ecosys-OrderByQuery)
    - [OrderByRegistrationOTP](#operations_ecosys-OrderByRegistrationOTP)
    - [OrderByRoster](#operations_ecosys-OrderByRoster)
    - [OrderByUser](#operations_ecosys-OrderByUser)
    - [OrderByUserToken](#operations_ecosys-OrderByUserToken)
    - [RegistrationCode](#operations_ecosys-RegistrationCode)
    - [RegistrationCodeRequest](#operations_ecosys-RegistrationCodeRequest)
    - [RegistrationCodeResponse](#operations_ecosys-RegistrationCodeResponse)
    - [RegistrationOTP](#operations_ecosys-RegistrationOTP)
    - [RegistrationOTPFilter](#operations_ecosys-RegistrationOTPFilter)
    - [RegistrationOTPQuery](#operations_ecosys-RegistrationOTPQuery)
    - [Response](#operations_ecosys-Response)
    - [ResponseNonce](#operations_ecosys-ResponseNonce)
    - [Roster](#operations_ecosys-Roster)
    - [RosterAssignement](#operations_ecosys-RosterAssignement)
    - [RosterAssignmentResponse](#operations_ecosys-RosterAssignmentResponse)
    - [RosterFilter](#operations_ecosys-RosterFilter)
    - [RosterQuery](#operations_ecosys-RosterQuery)
    - [RosterResponse](#operations_ecosys-RosterResponse)
    - [SecurityStringResponse](#operations_ecosys-SecurityStringResponse)
    - [User](#operations_ecosys-User)
    - [UserFilter](#operations_ecosys-UserFilter)
    - [UserQuery](#operations_ecosys-UserQuery)
    - [UserToken](#operations_ecosys-UserToken)
    - [UserTokenFilter](#operations_ecosys-UserTokenFilter)
    - [UserTokenQuery](#operations_ecosys-UserTokenQuery)
    - [UserTokenResponse](#operations_ecosys-UserTokenResponse)
    - [UsersResponse](#operations_ecosys-UsersResponse)
  
    - [AvailabilityFilter.Field](#operations_ecosys-AvailabilityFilter-Field)
    - [Broadcast.BroadcastType](#operations_ecosys-Broadcast-BroadcastType)
    - [Broadcast.UrgencyType](#operations_ecosys-Broadcast-UrgencyType)
    - [BroadcastFilter.Field](#operations_ecosys-BroadcastFilter-Field)
    - [CameraIot.MessageType](#operations_ecosys-CameraIot-MessageType)
    - [CameraIotFilter.Field](#operations_ecosys-CameraIotFilter-Field)
    - [ClientFilter.Field](#operations_ecosys-ClientFilter-Field)
    - [Filter.Comparisons](#operations_ecosys-Filter-Comparisons)
    - [IncidentReport.ReportType](#operations_ecosys-IncidentReport-ReportType)
    - [IncidentReportFilter.Field](#operations_ecosys-IncidentReportFilter-Field)
    - [OrderBy](#operations_ecosys-OrderBy)
    - [RegistrationCodeRequest.CodeType](#operations_ecosys-RegistrationCodeRequest-CodeType)
    - [RegistrationCodeResponse.CodeType](#operations_ecosys-RegistrationCodeResponse-CodeType)
    - [RegistrationOTPFilter.Field](#operations_ecosys-RegistrationOTPFilter-Field)
    - [Response.Type](#operations_ecosys-Response-Type)
    - [Roster.Status](#operations_ecosys-Roster-Status)
    - [RosterFilter.Field](#operations_ecosys-RosterFilter-Field)
    - [User.UserType](#operations_ecosys-User-UserType)
    - [UserFilter.Field](#operations_ecosys-UserFilter-Field)
    - [UserTokenFilter.Field](#operations_ecosys-UserTokenFilter-Field)
  
    - [AdminServices](#operations_ecosys-AdminServices)
    - [BroadcastServices](#operations_ecosys-BroadcastServices)
    - [CameraIotServices](#operations_ecosys-CameraIotServices)
    - [IncidentReportServices](#operations_ecosys-IncidentReportServices)
    - [RosterServices](#operations_ecosys-RosterServices)
  
- [Scalar Value Types](#scalar-value-types)



<a name="http_webapp-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## http_webapp.proto



<a name="http_webapp-HTTPAssignmentResponse"></a>

### HTTPAssignmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [HTTPMessage](#http_webapp-HTTPMessage) |  |  |
| rosters | [HTTPRosterResponse](#http_webapp-HTTPRosterResponse) | repeated |  |






<a name="http_webapp-HTTPAssignmentsGetRequest"></a>

### HTTPAssignmentsGetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| twan | [string](#string) |  |  |
| tele_user_id | [int64](#int64) |  |  |
| startDate | [string](#string) |  |  |
| endDate | [string](#string) |  |  |






<a name="http_webapp-HTTPMessage"></a>

### HTTPMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [int64](#int64) |  | status==0 denotes no issues. status &gt; 0 denotes some issue. |
| value | [string](#string) |  |  |
| valueArr | [string](#string) | repeated |  |






<a name="http_webapp-HTTPRegistrationFormRequest"></a>

### HTTPRegistrationFormRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  |  |
| tele_user_id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| phoneNumber | [string](#string) |  |  |
| loginString | [string](#string) |  |  |
| hasedLoginString | [string](#string) |  |  |
| isPartTime | [bool](#bool) |  |  |
| tele_handle | [string](#string) |  |  |






<a name="http_webapp-HTTPReportPostRequest"></a>

### HTTPReportPostRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| twan | [string](#string) |  |  |
| tele_user_id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| report_type | [string](#string) |  |  |
| address | [string](#string) |  |  |
| date | [string](#string) |  |  |
| time | [string](#string) |  |  |
| details | [string](#string) |  |  |
| isPeopleInjured | [bool](#bool) |  |  |
| injuryDetails | [string](#string) |  |  |
| isPoliceNotified | [bool](#bool) |  |  |
| isPropertyStolen | [bool](#bool) |  |  |
| propertyStolenDetails | [string](#string) |  |  |
| isActionTaken | [bool](#bool) |  |  |
| actionDetails | [string](#string) |  |  |






<a name="http_webapp-HTTPRosterResponse"></a>

### HTTPRosterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aifs_id | [int64](#int64) |  |  |
| startDatetime | [string](#string) |  |  |
| endDatetime | [string](#string) |  |  |
| addresses | [string](#string) | repeated |  |





 

 

 


<a name="http_webapp-WebAppServices"></a>

### WebAppServices


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetRosterAssignmentsForWebApp | [HTTPAssignmentsGetRequest](#http_webapp-HTTPAssignmentsGetRequest) | [HTTPAssignmentResponse](#http_webapp-HTTPAssignmentResponse) | &lt;&gt;&lt;&gt; Telegram WebApp - HTTP &lt;&gt;&lt;&gt; |
| PostWReportFromWebApp | [HTTPReportPostRequest](#http_webapp-HTTPReportPostRequest) | [HTTPMessage](#http_webapp-HTTPMessage) |  |
| PostRegistrationFormFromWebApp | [HTTPRegistrationFormRequest](#http_webapp-HTTPRegistrationFormRequest) | [HTTPMessage](#http_webapp-HTTPMessage) |  |

 



<a name="iot_prototype-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## iot_prototype.proto



<a name="gate_prototype-CpuTemp"></a>

### CpuTemp
The request message containing the fire alarm&#39;s id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="gate_prototype-CpuTempState"></a>

### CpuTempState
The response message containing the state of the stepper


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| temp | [double](#double) |  |  |






<a name="gate_prototype-FireAlarm"></a>

### FireAlarm
The request message containing the fire alarm&#39;s id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="gate_prototype-FireAlarmState"></a>

### FireAlarmState
The response message containing the state of the stepper


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| state | [FireAlarmState.AlarmState](#gate_prototype-FireAlarmState-AlarmState) |  |  |






<a name="gate_prototype-Gate"></a>

### Gate
The request message containing the stepper&#39;s name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="gate_prototype-GateState"></a>

### GateState
The response message containing the state of the stepper


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| state | [GateState.GatePosition](#gate_prototype-GateState-GatePosition) |  |  |





 


<a name="gate_prototype-FireAlarmState-AlarmState"></a>

### FireAlarmState.AlarmState


| Name | Number | Description |
| ---- | ------ | ----------- |
| OFF | 0 |  |
| ON | 1 |  |
| ERROR | 2 |  |



<a name="gate_prototype-GateState-GatePosition"></a>

### GateState.GatePosition


| Name | Number | Description |
| ---- | ------ | ----------- |
| CLOSED | 0 |  |
| OPEN | 1 |  |
| ERROR | 2 |  |
| INITIAL | 3 |  |


 

 


<a name="gate_prototype-IotControlPrototypeService"></a>

### IotControlPrototypeService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetGateState | [Gate](#gate_prototype-Gate) | [GateState](#gate_prototype-GateState) stream | A server-to-client streaming RPC.

Obtains the State of the Stepper. Results are streamed rather than returned at once (e.g. in a response message with a repeated field), as the Stepper&#39;s state will change after every action made by the Controller. |
| SetGateState | [GateState](#gate_prototype-GateState) stream | [GateState](#gate_prototype-GateState) stream |  |
| GetFireAlarmState | [FireAlarm](#gate_prototype-FireAlarm) | [FireAlarmState](#gate_prototype-FireAlarmState) stream |  |
| GetCpuTemp | [CpuTemp](#gate_prototype-CpuTemp) | [CpuTempState](#gate_prototype-CpuTempState) stream |  |

 



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






<a name="operations_ecosys-Camera"></a>

### Camera
camera identifier


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="operations_ecosys-CameraIot"></a>

### CameraIot
Camera and Iot Monitoring and Controls


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| camera_iot_id | [int64](#int64) |  | The ID ties the request to a paricular location |
| name | [string](#string) |  |  |
| camera | [Camera](#operations_ecosys-Camera) |  | Devices |
| gate | [gate_prototype.GateState](#gate_prototype-GateState) |  |  |
| fire_alarm | [gate_prototype.FireAlarmState](#gate_prototype-FireAlarmState) |  |  |
| cpu_temperature | [gate_prototype.CpuTempState](#gate_prototype-CpuTempState) |  |  |
| type | [CameraIot.MessageType](#operations_ecosys-CameraIot-MessageType) |  |  |






<a name="operations_ecosys-CameraIotFilter"></a>

### CameraIotFilter
Filter the types of camera iot attributes to be returned.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [CameraIotFilter.Field](#operations_ecosys-CameraIotFilter-Field) |  |  |
| comparisons | [Filter](#operations_ecosys-Filter) |  |  |






<a name="operations_ecosys-CameraIotQuery"></a>

### CameraIotQuery
Get specific types camera Iot attributes as specified in the Filter. 
If one wants to get all objects, leave filters empty. 
A default limit of 10 will be used if the field is empty.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [CameraIotFilter](#operations_ecosys-CameraIotFilter) | repeated |  |
| limit | [int64](#int64) |  | Limit the number of objects being returned. If only 5 objects should be shown, limit = 5; |
| skip | [int64](#int64) |  | Skip n rows from the database |
| order_by | [OrderByCameraIot](#operations_ecosys-OrderByCameraIot) |  | Order the queries, by default the order is desc id |






<a name="operations_ecosys-CameraIotResponse"></a>

### CameraIotResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| camera_iot | [CameraIot](#operations_ecosys-CameraIot) |  |  |






<a name="operations_ecosys-Client"></a>

### Client
All clients who hire the company to protect their compounds.


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






<a name="operations_ecosys-FullUser"></a>

### FullUser
This message is mostly used by the internal server
and for the creation of new users only. 
There should be no other reasons to pass around these 
confidential messages.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#operations_ecosys-User) |  |  |
| nonce | [string](#string) |  |  |
| security_string | [string](#string) |  |  |
| hashed_password | [string](#string) |  |  |






<a name="operations_ecosys-IncidentReport"></a>

### IncidentReport
The default fields of an incident report


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| incident_report_id | [int64](#int64) |  | Report IDs are only useful for the backend database. |
| type | [IncidentReport.ReportType](#operations_ecosys-IncidentReport-ReportType) |  |  |
| creator | [User](#operations_ecosys-User) |  |  |
| creation_date | [string](#string) |  | Format must be in: YYYY-MM-DD HH:MM:SS |
| last_modified_date | [string](#string) |  |  |
| last_modifed_user | [User](#operations_ecosys-User) |  |  |
| is_original | [bool](#bool) |  |  |
| is_approved | [bool](#bool) |  |  |
| signature | [User](#operations_ecosys-User) |  |  |
| approval_date | [string](#string) |  |  |
| incident_report_content | [IncidentReportContent](#operations_ecosys-IncidentReportContent) |  |  |
| aifs_id | [int64](#int64) |  | Which AIFS was the creator from |






<a name="operations_ecosys-IncidentReportContent"></a>

### IncidentReportContent
The actual content of the report


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| report_content_id | [int64](#int64) |  |  |
| last_modified_date | [string](#string) |  | Format must be in: YYYY-MM-DD HH:MM:SS |
| last_modifed_user | [User](#operations_ecosys-User) |  |  |
| address | [string](#string) |  |  |
| incident_time | [string](#string) |  |  |
| is_police_notified | [bool](#bool) |  |  |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| has_action_taken | [bool](#bool) |  |  |
| action_taken | [string](#string) |  |  |
| has_injury | [bool](#bool) |  |  |
| injury_description | [string](#string) |  |  |
| has_stolen_item | [bool](#bool) |  |  |
| stolen_item_description | [string](#string) |  |  |
| report_image_link | [string](#string) |  | Should there be a string here? |






<a name="operations_ecosys-IncidentReportFilter"></a>

### IncidentReportFilter
Filter the types of incident reports to be returned.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [IncidentReportFilter.Field](#operations_ecosys-IncidentReportFilter-Field) |  |  |
| comparisons | [Filter](#operations_ecosys-Filter) |  |  |






<a name="operations_ecosys-IncidentReportQuery"></a>

### IncidentReportQuery
Get specific types rosters as specified in the Filter. 
If one wants to get all objects, leave filters empty. 
A default limit of 10 will be used if the field is empty.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [IncidentReportFilter](#operations_ecosys-IncidentReportFilter) | repeated |  |
| limit | [int64](#int64) |  | Limit the number of objects being returned. If only 5 objects should be shown, limit = 5; |
| skip | [int64](#int64) |  | Skip n rows from the database |
| order_by | [OrderByIncidentReport](#operations_ecosys-OrderByIncidentReport) |  | Order the queries, by default the order is desc by creation date |






<a name="operations_ecosys-IncidentReportResponse"></a>

### IncidentReportResponse
Passing around multiple reports in one message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| incident_report | [IncidentReport](#operations_ecosys-IncidentReport) |  |  |






<a name="operations_ecosys-LoginRequest"></a>

### LoginRequest
Any client that wishes to login send their hashed password to the backend  
for authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_email | [string](#string) |  | The user to login |
| hashed_password | [string](#string) |  |  |






<a name="operations_ecosys-OrderByBroadcast"></a>

### OrderByBroadcast



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [BroadcastFilter.Field](#operations_ecosys-BroadcastFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-OrderByCameraIot"></a>

### OrderByCameraIot



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [CameraIotFilter.Field](#operations_ecosys-CameraIotFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-OrderByClient"></a>

### OrderByClient



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [ClientFilter.Field](#operations_ecosys-ClientFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-OrderByIncidentReport"></a>

### OrderByIncidentReport



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [IncidentReportFilter.Field](#operations_ecosys-IncidentReportFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-OrderByQuery"></a>

### OrderByQuery



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [AvailabilityFilter.Field](#operations_ecosys-AvailabilityFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-OrderByRegistrationOTP"></a>

### OrderByRegistrationOTP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [RegistrationOTPFilter.Field](#operations_ecosys-RegistrationOTPFilter-Field) |  |  |
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






<a name="operations_ecosys-OrderByUserToken"></a>

### OrderByUserToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [UserTokenFilter.Field](#operations_ecosys-UserTokenFilter-Field) |  |  |
| order_by | [OrderBy](#operations_ecosys-OrderBy) |  |  |






<a name="operations_ecosys-RegistrationCode"></a>

### RegistrationCode



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  |  |






<a name="operations_ecosys-RegistrationCodeRequest"></a>

### RegistrationCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#operations_ecosys-User) |  |  |
| type | [RegistrationCodeRequest.CodeType](#operations_ecosys-RegistrationCodeRequest-CodeType) |  |  |






<a name="operations_ecosys-RegistrationCodeResponse"></a>

### RegistrationCodeResponse
Returns a registration code that can be used to create a new user


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| type | [RegistrationCodeResponse.CodeType](#operations_ecosys-RegistrationCodeResponse-CodeType) |  |  |
| code | [string](#string) |  |  |






<a name="operations_ecosys-RegistrationOTP"></a>

### RegistrationOTP
Message used iternally to get the registration OTP from the DB


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| registration_otp_id | [int64](#int64) |  |  |
| token | [string](#string) |  |  |
| user_type | [User.UserType](#operations_ecosys-User-UserType) |  |  |
| creation_datetime | [string](#string) |  | Format must be in: YYYY-MM-DD HH:MM:SS |
| creator | [User](#operations_ecosys-User) |  |  |
| is_used | [bool](#bool) |  |  |






<a name="operations_ecosys-RegistrationOTPFilter"></a>

### RegistrationOTPFilter
Filter the types of registration otp to be returned.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [RegistrationOTPFilter.Field](#operations_ecosys-RegistrationOTPFilter-Field) |  |  |
| comparisons | [Filter](#operations_ecosys-Filter) |  |  |






<a name="operations_ecosys-RegistrationOTPQuery"></a>

### RegistrationOTPQuery
Get specific types of Registration OTP as specified in the Filter. 
If one wants to get all objects, leave filters empty. 
A default limit of 10 will be used if the field is empty.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [RegistrationOTPFilter](#operations_ecosys-RegistrationOTPFilter) | repeated |  |
| limit | [int64](#int64) |  | Limit the number of objects being returned. If only 5 objects should be shown, limit = 5; |
| skip | [int64](#int64) |  | Skip n rows from the database |
| order_by | [OrderByRegistrationOTP](#operations_ecosys-OrderByRegistrationOTP) |  | Order the queries, by default the order is desc by id |






<a name="operations_ecosys-Response"></a>

### Response
Generic reponses to add or update requests


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Response.Type](#operations_ecosys-Response-Type) |  |  |
| error_message | [string](#string) |  |  |
| primary_key | [int64](#int64) |  | Return any pk of the row that the query modified |






<a name="operations_ecosys-ResponseNonce"></a>

### ResponseNonce
Used to provide nonces


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| nonce | [string](#string) |  |  |






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






<a name="operations_ecosys-SecurityStringResponse"></a>

### SecurityStringResponse
Returns the security string that is tied to the user


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| security_string | [string](#string) |  |  |






<a name="operations_ecosys-User"></a>

### User
All users who use the operations ecosystem.


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
| tele_user_id | [int64](#int64) |  |  |






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






<a name="operations_ecosys-UserToken"></a>

### UserToken
Message used to provide a user token to the client


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_token_id | [int64](#int64) |  |  |
| user | [User](#operations_ecosys-User) |  |  |
| token | [string](#string) |  |  |
| creation_datetime | [string](#string) |  | Format must be in: YYYY-MM-DD HH:MM:SS |
| expiry_datetime | [string](#string) |  | Format must be in: YYYY-MM-DD HH:MM:SS |






<a name="operations_ecosys-UserTokenFilter"></a>

### UserTokenFilter
Filter the types of clients to be returned.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field | [UserTokenFilter.Field](#operations_ecosys-UserTokenFilter-Field) |  |  |
| comparisons | [Filter](#operations_ecosys-Filter) |  |  |






<a name="operations_ecosys-UserTokenQuery"></a>

### UserTokenQuery
Get specific types of user tokens as specified in the Filter. 
If one wants to get all objects, leave filters empty. 
A default limit of 10 will be used if the field is empty.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filters | [UserTokenFilter](#operations_ecosys-UserTokenFilter) | repeated |  |
| limit | [int64](#int64) |  | Limit the number of objects being returned. If only 5 objects should be shown, limit = 5; |
| skip | [int64](#int64) |  | Skip n rows from the database |
| order_by | [OrderByUserToken](#operations_ecosys-OrderByUserToken) |  | Order the queries, by default the order is desc by creation date |






<a name="operations_ecosys-UserTokenResponse"></a>

### UserTokenResponse
Replies with a user token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#operations_ecosys-Response) |  |  |
| userToken | [UserToken](#operations_ecosys-UserToken) |  |  |






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



<a name="operations_ecosys-CameraIot-MessageType"></a>

### CameraIot.MessageType


| Name | Number | Description |
| ---- | ------ | ----------- |
| INITIAL | 0 |  |
| CHANGE_GATE | 1 |  |
| CHANGE_FIRE_ALARM | 2 |  |
| CHANGE_CPU_TEMP | 3 |  |



<a name="operations_ecosys-CameraIotFilter-Field"></a>

### CameraIotFilter.Field
More fields can be added in the future.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CAMERA_IOT_ID | 0 |  |



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



<a name="operations_ecosys-IncidentReport-ReportType"></a>

### IncidentReport.ReportType


| Name | Number | Description |
| ---- | ------ | ----------- |
| FIRE_ALARM | 0 |  |
| INTRUDER | 1 |  |
| OTHERS | 2 |  |



<a name="operations_ecosys-IncidentReportFilter-Field"></a>

### IncidentReportFilter.Field
More fields can be added in the future.

| Name | Number | Description |
| ---- | ------ | ----------- |
| REPORT_ID | 0 |  |
| REPORT_CONTENT_ID | 1 |  |
| REPORT_TYPE | 2 |  |
| MODIFIER | 3 |  |
| LAST_MODIFIED_DATE | 5 |  |
| GET_ORIGINAL | 6 |  |
| IS_APPROVED | 7 |  |
| SIGNATURE | 8 |  |
| APPROVAL_DATE | 9 |  |



<a name="operations_ecosys-OrderBy"></a>

### OrderBy


| Name | Number | Description |
| ---- | ------ | ----------- |
| ASC | 0 |  |
| DESC | 1 |  |



<a name="operations_ecosys-RegistrationCodeRequest-CodeType"></a>

### RegistrationCodeRequest.CodeType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ISPECIALIST | 0 |  |
| SECURITYGUARD | 1 |  |
| CONTROLLER | 2 |  |
| MANAGER | 3 |  |



<a name="operations_ecosys-RegistrationCodeResponse-CodeType"></a>

### RegistrationCodeResponse.CodeType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ISPECIALIST | 0 |  |
| CONTROLLER | 1 |  |
| MANAGER | 2 |  |



<a name="operations_ecosys-RegistrationOTPFilter-Field"></a>

### RegistrationOTPFilter.Field
More fields can be added in the future.

| Name | Number | Description |
| ---- | ------ | ----------- |
| REG_OTP_ID | 0 |  |
| TOKEN | 1 |  |
| USER_TYPE | 2 |  |
| CREATION_DATE | 3 |  |
| CREATOR_ID | 4 |  |
| IS_USED | 5 |  |



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
| TELEGRAM_USER_ID | 7 |  |



<a name="operations_ecosys-UserTokenFilter-Field"></a>

### UserTokenFilter.Field
More fields can be added in the future.

| Name | Number | Description |
| ---- | ------ | ----------- |
| USER_TOKEN_ID | 0 |  |
| USER_ID | 1 |  |
| EXPIRY | 2 |  |
| TOKEN | 3 |  |


 

 


<a name="operations_ecosys-AdminServices"></a>

### AdminServices


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddUser | [FullUser](#operations_ecosys-FullUser) | [Response](#operations_ecosys-Response) | User |
| UpdateUser | [User](#operations_ecosys-User) | [Response](#operations_ecosys-Response) |  |
| DeleteUser | [User](#operations_ecosys-User) | [Response](#operations_ecosys-Response) |  |
| FindUsers | [UserQuery](#operations_ecosys-UserQuery) | [UsersResponse](#operations_ecosys-UsersResponse) stream |  |
| AddClient | [Client](#operations_ecosys-Client) | [Response](#operations_ecosys-Response) | Client |
| UpdateClient | [Client](#operations_ecosys-Client) | [Response](#operations_ecosys-Response) |  |
| DeleteClient | [Client](#operations_ecosys-Client) | [Response](#operations_ecosys-Response) |  |
| FindClients | [ClientQuery](#operations_ecosys-ClientQuery) | [ClientResponse](#operations_ecosys-ClientResponse) stream |  |
| GetWANonce | [User](#operations_ecosys-User) | [ResponseNonce](#operations_ecosys-ResponseNonce) | Security Related |
| GetSecurityString | [User](#operations_ecosys-User) | [SecurityStringResponse](#operations_ecosys-SecurityStringResponse) |  |
| AuthenticateUser | [LoginRequest](#operations_ecosys-LoginRequest) | [UserTokenResponse](#operations_ecosys-UserTokenResponse) |  |
| GetRegistrationCode | [RegistrationCodeRequest](#operations_ecosys-RegistrationCodeRequest) | [RegistrationCodeResponse](#operations_ecosys-RegistrationCodeResponse) | Is this user or client? |
| CheckRegistrationCode | [RegistrationCode](#operations_ecosys-RegistrationCode) | [SecurityStringResponse](#operations_ecosys-SecurityStringResponse) |  |


<a name="operations_ecosys-BroadcastServices"></a>

### BroadcastServices


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddBroadcast | [Broadcast](#operations_ecosys-Broadcast) | [Response](#operations_ecosys-Response) |  |
| UpdateBroadcast | [Broadcast](#operations_ecosys-Broadcast) | [Response](#operations_ecosys-Response) | Note that this update does not update the broadcast&#39;s recipient&#39;s inner status such as the acknowledgement or rejection status but only if the recipient is part of the broadcast. |
| DeleteBroadcast | [Broadcast](#operations_ecosys-Broadcast) | [Response](#operations_ecosys-Response) |  |
| FindBroadcasts | [BroadcastQuery](#operations_ecosys-BroadcastQuery) | [BroadcastResponse](#operations_ecosys-BroadcastResponse) stream |  |
| UpdateBroadcastRecipient | [BroadcastRecipient](#operations_ecosys-BroadcastRecipient) | [Response](#operations_ecosys-Response) | Updating of broadcast recipients |


<a name="operations_ecosys-CameraIotServices"></a>

### CameraIotServices


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetGateState | [.gate_prototype.GateState](#gate_prototype-GateState) | [Response](#operations_ecosys-Response) |  |
| GetIotState | [.google.protobuf.Empty](#google-protobuf-Empty) | [CameraIotResponse](#operations_ecosys-CameraIotResponse) stream | Continuously provides the states of the gates, fire alarms and cpu temperature as well as the camera endpoints. Responses are sent only when there is a change in state Upon connection, all states are sent for all locations are sent. |


<a name="operations_ecosys-IncidentReportServices"></a>

### IncidentReportServices


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddIncidentReport | [IncidentReport](#operations_ecosys-IncidentReport) | [Response](#operations_ecosys-Response) |  |
| UpdateIncidentReport | [IncidentReport](#operations_ecosys-IncidentReport) | [Response](#operations_ecosys-Response) |  |
| DeleteIncidentReport | [IncidentReport](#operations_ecosys-IncidentReport) | [Response](#operations_ecosys-Response) |  |
| FindIncidentReports | [IncidentReportQuery](#operations_ecosys-IncidentReportQuery) | [IncidentReportResponse](#operations_ecosys-IncidentReportResponse) stream |  |


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

