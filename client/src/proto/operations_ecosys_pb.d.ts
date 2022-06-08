import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class User extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): User;

  getUserType(): User.UserType;
  setUserType(value: User.UserType): User;

  getName(): string;
  setName(value: string): User;

  getEmail(): string;
  setEmail(value: string): User;

  getPhoneNumber(): string;
  setPhoneNumber(value: string): User;

  getTelegramHandle(): string;
  setTelegramHandle(value: string): User;

  getUserSecurityImg(): string;
  setUserSecurityImg(value: string): User;

  getIsPartTimer(): boolean;
  setIsPartTimer(value: boolean): User;

  getTeleChatId(): number;
  setTeleChatId(value: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    userId: number,
    userType: User.UserType,
    name: string,
    email: string,
    phoneNumber: string,
    telegramHandle: string,
    userSecurityImg: string,
    isPartTimer: boolean,
    teleChatId: number,
  }

  export enum UserType { 
    ISPECIALIST = 0,
    SECURITY_GUARD = 1,
    CONTROLLER = 2,
    MANAGER = 3,
  }
}

export class UsersResponse extends jspb.Message {
  getResponse(): Response | undefined;
  setResponse(value?: Response): UsersResponse;
  hasResponse(): boolean;
  clearResponse(): UsersResponse;

  getUser(): User | undefined;
  setUser(value?: User): UsersResponse;
  hasUser(): boolean;
  clearUser(): UsersResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UsersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UsersResponse): UsersResponse.AsObject;
  static serializeBinaryToWriter(message: UsersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UsersResponse;
  static deserializeBinaryFromReader(message: UsersResponse, reader: jspb.BinaryReader): UsersResponse;
}

export namespace UsersResponse {
  export type AsObject = {
    response?: Response.AsObject,
    user?: User.AsObject,
  }
}

export class UserFilter extends jspb.Message {
  getField(): UserFilter.Field;
  setField(value: UserFilter.Field): UserFilter;

  getComparisons(): Filter | undefined;
  setComparisons(value?: Filter): UserFilter;
  hasComparisons(): boolean;
  clearComparisons(): UserFilter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserFilter.AsObject;
  static toObject(includeInstance: boolean, msg: UserFilter): UserFilter.AsObject;
  static serializeBinaryToWriter(message: UserFilter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserFilter;
  static deserializeBinaryFromReader(message: UserFilter, reader: jspb.BinaryReader): UserFilter;
}

export namespace UserFilter {
  export type AsObject = {
    field: UserFilter.Field,
    comparisons?: Filter.AsObject,
  }

  export enum Field { 
    USER_ID = 0,
    TYPE = 1,
    NAME = 2,
    EMAIL = 3,
    PHONE_NUMBER = 4,
    TELEGRAM_HANDLE = 5,
    IS_PART_TIMER = 6,
  }
}

export class UserQuery extends jspb.Message {
  getFiltersList(): Array<UserFilter>;
  setFiltersList(value: Array<UserFilter>): UserQuery;
  clearFiltersList(): UserQuery;
  addFilters(value?: UserFilter, index?: number): UserFilter;

  getLimit(): number;
  setLimit(value: number): UserQuery;

  getSkip(): number;
  setSkip(value: number): UserQuery;

  getOrderBy(): OrderByUser | undefined;
  setOrderBy(value?: OrderByUser): UserQuery;
  hasOrderBy(): boolean;
  clearOrderBy(): UserQuery;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserQuery.AsObject;
  static toObject(includeInstance: boolean, msg: UserQuery): UserQuery.AsObject;
  static serializeBinaryToWriter(message: UserQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserQuery;
  static deserializeBinaryFromReader(message: UserQuery, reader: jspb.BinaryReader): UserQuery;
}

export namespace UserQuery {
  export type AsObject = {
    filtersList: Array<UserFilter.AsObject>,
    limit: number,
    skip: number,
    orderBy?: OrderByUser.AsObject,
  }
}

export class OrderByUser extends jspb.Message {
  getField(): UserFilter.Field;
  setField(value: UserFilter.Field): OrderByUser;

  getOrderBy(): OrderBy;
  setOrderBy(value: OrderBy): OrderByUser;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OrderByUser.AsObject;
  static toObject(includeInstance: boolean, msg: OrderByUser): OrderByUser.AsObject;
  static serializeBinaryToWriter(message: OrderByUser, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OrderByUser;
  static deserializeBinaryFromReader(message: OrderByUser, reader: jspb.BinaryReader): OrderByUser;
}

export namespace OrderByUser {
  export type AsObject = {
    field: UserFilter.Field,
    orderBy: OrderBy,
  }
}

export class Client extends jspb.Message {
  getClientId(): number;
  setClientId(value: number): Client;

  getName(): string;
  setName(value: string): Client;

  getAbbreviation(): string;
  setAbbreviation(value: string): Client;

  getEmail(): string;
  setEmail(value: string): Client;

  getAddress(): string;
  setAddress(value: string): Client;

  getPostalCode(): number;
  setPostalCode(value: number): Client;

  getPhoneNumber(): string;
  setPhoneNumber(value: string): Client;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Client.AsObject;
  static toObject(includeInstance: boolean, msg: Client): Client.AsObject;
  static serializeBinaryToWriter(message: Client, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Client;
  static deserializeBinaryFromReader(message: Client, reader: jspb.BinaryReader): Client;
}

export namespace Client {
  export type AsObject = {
    clientId: number,
    name: string,
    abbreviation: string,
    email: string,
    address: string,
    postalCode: number,
    phoneNumber: string,
  }
}

export class ClientResponse extends jspb.Message {
  getResponse(): Response | undefined;
  setResponse(value?: Response): ClientResponse;
  hasResponse(): boolean;
  clearResponse(): ClientResponse;

  getClient(): Client | undefined;
  setClient(value?: Client): ClientResponse;
  hasClient(): boolean;
  clearClient(): ClientResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ClientResponse): ClientResponse.AsObject;
  static serializeBinaryToWriter(message: ClientResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientResponse;
  static deserializeBinaryFromReader(message: ClientResponse, reader: jspb.BinaryReader): ClientResponse;
}

export namespace ClientResponse {
  export type AsObject = {
    response?: Response.AsObject,
    client?: Client.AsObject,
  }
}

export class ClientFilter extends jspb.Message {
  getField(): ClientFilter.Field;
  setField(value: ClientFilter.Field): ClientFilter;

  getComparisons(): Filter | undefined;
  setComparisons(value?: Filter): ClientFilter;
  hasComparisons(): boolean;
  clearComparisons(): ClientFilter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientFilter.AsObject;
  static toObject(includeInstance: boolean, msg: ClientFilter): ClientFilter.AsObject;
  static serializeBinaryToWriter(message: ClientFilter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientFilter;
  static deserializeBinaryFromReader(message: ClientFilter, reader: jspb.BinaryReader): ClientFilter;
}

export namespace ClientFilter {
  export type AsObject = {
    field: ClientFilter.Field,
    comparisons?: Filter.AsObject,
  }

  export enum Field { 
    CLIENT_ID = 0,
  }
}

export class ClientQuery extends jspb.Message {
  getFiltersList(): Array<ClientFilter>;
  setFiltersList(value: Array<ClientFilter>): ClientQuery;
  clearFiltersList(): ClientQuery;
  addFilters(value?: ClientFilter, index?: number): ClientFilter;

  getLimit(): number;
  setLimit(value: number): ClientQuery;

  getSkip(): number;
  setSkip(value: number): ClientQuery;

  getOrderBy(): OrderByClient | undefined;
  setOrderBy(value?: OrderByClient): ClientQuery;
  hasOrderBy(): boolean;
  clearOrderBy(): ClientQuery;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientQuery.AsObject;
  static toObject(includeInstance: boolean, msg: ClientQuery): ClientQuery.AsObject;
  static serializeBinaryToWriter(message: ClientQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientQuery;
  static deserializeBinaryFromReader(message: ClientQuery, reader: jspb.BinaryReader): ClientQuery;
}

export namespace ClientQuery {
  export type AsObject = {
    filtersList: Array<ClientFilter.AsObject>,
    limit: number,
    skip: number,
    orderBy?: OrderByClient.AsObject,
  }
}

export class OrderByClient extends jspb.Message {
  getField(): ClientFilter.Field;
  setField(value: ClientFilter.Field): OrderByClient;

  getOrderBy(): OrderBy;
  setOrderBy(value: OrderBy): OrderByClient;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OrderByClient.AsObject;
  static toObject(includeInstance: boolean, msg: OrderByClient): OrderByClient.AsObject;
  static serializeBinaryToWriter(message: OrderByClient, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OrderByClient;
  static deserializeBinaryFromReader(message: OrderByClient, reader: jspb.BinaryReader): OrderByClient;
}

export namespace OrderByClient {
  export type AsObject = {
    field: ClientFilter.Field,
    orderBy: OrderBy,
  }
}

export class Broadcast extends jspb.Message {
  getBroadcastId(): number;
  setBroadcastId(value: number): Broadcast;

  getType(): Broadcast.BroadcastType;
  setType(value: Broadcast.BroadcastType): Broadcast;

  getContent(): string;
  setContent(value: string): Broadcast;

  getCreationDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreationDate(value?: google_protobuf_timestamp_pb.Timestamp): Broadcast;
  hasCreationDate(): boolean;
  clearCreationDate(): Broadcast;

  getDeadline(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeadline(value?: google_protobuf_timestamp_pb.Timestamp): Broadcast;
  hasDeadline(): boolean;
  clearDeadline(): Broadcast;

  getCreator(): User | undefined;
  setCreator(value?: User): Broadcast;
  hasCreator(): boolean;
  clearCreator(): Broadcast;

  getRecipientsList(): Array<AIFSBroadcastRecipient>;
  setRecipientsList(value: Array<AIFSBroadcastRecipient>): Broadcast;
  clearRecipientsList(): Broadcast;
  addRecipients(value?: AIFSBroadcastRecipient, index?: number): AIFSBroadcastRecipient;

  getUrgency(): Broadcast.UrgencyType;
  setUrgency(value: Broadcast.UrgencyType): Broadcast;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Broadcast.AsObject;
  static toObject(includeInstance: boolean, msg: Broadcast): Broadcast.AsObject;
  static serializeBinaryToWriter(message: Broadcast, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Broadcast;
  static deserializeBinaryFromReader(message: Broadcast, reader: jspb.BinaryReader): Broadcast;
}

export namespace Broadcast {
  export type AsObject = {
    broadcastId: number,
    type: Broadcast.BroadcastType,
    content: string,
    creationDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deadline?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    creator?: User.AsObject,
    recipientsList: Array<AIFSBroadcastRecipient.AsObject>,
    urgency: Broadcast.UrgencyType,
  }

  export enum BroadcastType { 
    ANNOUNCEMENT = 0,
    ASSIGNMENT = 1,
  }

  export enum UrgencyType { 
    LOW = 0,
    MEDIUM = 1,
    HIGH = 2,
  }
}

export class AIFSBroadcastRecipient extends jspb.Message {
  getRecipientList(): Array<BroadcastRecipient>;
  setRecipientList(value: Array<BroadcastRecipient>): AIFSBroadcastRecipient;
  clearRecipientList(): AIFSBroadcastRecipient;
  addRecipient(value?: BroadcastRecipient, index?: number): BroadcastRecipient;

  getAifsId(): number;
  setAifsId(value: number): AIFSBroadcastRecipient;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AIFSBroadcastRecipient.AsObject;
  static toObject(includeInstance: boolean, msg: AIFSBroadcastRecipient): AIFSBroadcastRecipient.AsObject;
  static serializeBinaryToWriter(message: AIFSBroadcastRecipient, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AIFSBroadcastRecipient;
  static deserializeBinaryFromReader(message: AIFSBroadcastRecipient, reader: jspb.BinaryReader): AIFSBroadcastRecipient;
}

export namespace AIFSBroadcastRecipient {
  export type AsObject = {
    recipientList: Array<BroadcastRecipient.AsObject>,
    aifsId: number,
  }
}

export class BroadcastRecipient extends jspb.Message {
  getBroadcastRecipientsId(): number;
  setBroadcastRecipientsId(value: number): BroadcastRecipient;

  getRecipient(): User | undefined;
  setRecipient(value?: User): BroadcastRecipient;
  hasRecipient(): boolean;
  clearRecipient(): BroadcastRecipient;

  getAcknowledged(): boolean;
  setAcknowledged(value: boolean): BroadcastRecipient;

  getRejected(): boolean;
  setRejected(value: boolean): BroadcastRecipient;

  getLastReplied(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastReplied(value?: google_protobuf_timestamp_pb.Timestamp): BroadcastRecipient;
  hasLastReplied(): boolean;
  clearLastReplied(): BroadcastRecipient;

  getAifsId(): number;
  setAifsId(value: number): BroadcastRecipient;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BroadcastRecipient.AsObject;
  static toObject(includeInstance: boolean, msg: BroadcastRecipient): BroadcastRecipient.AsObject;
  static serializeBinaryToWriter(message: BroadcastRecipient, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BroadcastRecipient;
  static deserializeBinaryFromReader(message: BroadcastRecipient, reader: jspb.BinaryReader): BroadcastRecipient;
}

export namespace BroadcastRecipient {
  export type AsObject = {
    broadcastRecipientsId: number,
    recipient?: User.AsObject,
    acknowledged: boolean,
    rejected: boolean,
    lastReplied?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    aifsId: number,
  }
}

export class BroadcastResponse extends jspb.Message {
  getResponse(): Response | undefined;
  setResponse(value?: Response): BroadcastResponse;
  hasResponse(): boolean;
  clearResponse(): BroadcastResponse;

  getBroadcast(): Broadcast | undefined;
  setBroadcast(value?: Broadcast): BroadcastResponse;
  hasBroadcast(): boolean;
  clearBroadcast(): BroadcastResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BroadcastResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BroadcastResponse): BroadcastResponse.AsObject;
  static serializeBinaryToWriter(message: BroadcastResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BroadcastResponse;
  static deserializeBinaryFromReader(message: BroadcastResponse, reader: jspb.BinaryReader): BroadcastResponse;
}

export namespace BroadcastResponse {
  export type AsObject = {
    response?: Response.AsObject,
    broadcast?: Broadcast.AsObject,
  }
}

export class BroadcastFilter extends jspb.Message {
  getField(): BroadcastFilter.Field;
  setField(value: BroadcastFilter.Field): BroadcastFilter;

  getComparisons(): Filter | undefined;
  setComparisons(value?: Filter): BroadcastFilter;
  hasComparisons(): boolean;
  clearComparisons(): BroadcastFilter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BroadcastFilter.AsObject;
  static toObject(includeInstance: boolean, msg: BroadcastFilter): BroadcastFilter.AsObject;
  static serializeBinaryToWriter(message: BroadcastFilter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BroadcastFilter;
  static deserializeBinaryFromReader(message: BroadcastFilter, reader: jspb.BinaryReader): BroadcastFilter;
}

export namespace BroadcastFilter {
  export type AsObject = {
    field: BroadcastFilter.Field,
    comparisons?: Filter.AsObject,
  }

  export enum Field { 
    BROADCAST_ID = 0,
    TYPE = 1,
    CONTENT = 2,
    CREATION_DATE = 3,
    DEADLINE = 4,
    CREATOR_ID = 5,
    RECEIPEIENT_ID = 6,
    NUM_RECEIPIENTS = 7,
    URGENCY = 8,
    AIFS_ID = 9,
    BROADCAST_RECIPIENT_TABLE_ID = 10,
  }
}

export class BroadcastQuery extends jspb.Message {
  getFiltersList(): Array<BroadcastFilter>;
  setFiltersList(value: Array<BroadcastFilter>): BroadcastQuery;
  clearFiltersList(): BroadcastQuery;
  addFilters(value?: BroadcastFilter, index?: number): BroadcastFilter;

  getLimit(): number;
  setLimit(value: number): BroadcastQuery;

  getSkip(): number;
  setSkip(value: number): BroadcastQuery;

  getOrderBy(): OrderByBroadcast | undefined;
  setOrderBy(value?: OrderByBroadcast): BroadcastQuery;
  hasOrderBy(): boolean;
  clearOrderBy(): BroadcastQuery;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BroadcastQuery.AsObject;
  static toObject(includeInstance: boolean, msg: BroadcastQuery): BroadcastQuery.AsObject;
  static serializeBinaryToWriter(message: BroadcastQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BroadcastQuery;
  static deserializeBinaryFromReader(message: BroadcastQuery, reader: jspb.BinaryReader): BroadcastQuery;
}

export namespace BroadcastQuery {
  export type AsObject = {
    filtersList: Array<BroadcastFilter.AsObject>,
    limit: number,
    skip: number,
    orderBy?: OrderByBroadcast.AsObject,
  }
}

export class OrderByBroadcast extends jspb.Message {
  getField(): BroadcastFilter.Field;
  setField(value: BroadcastFilter.Field): OrderByBroadcast;

  getOrderBy(): OrderBy;
  setOrderBy(value: OrderBy): OrderByBroadcast;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OrderByBroadcast.AsObject;
  static toObject(includeInstance: boolean, msg: OrderByBroadcast): OrderByBroadcast.AsObject;
  static serializeBinaryToWriter(message: OrderByBroadcast, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OrderByBroadcast;
  static deserializeBinaryFromReader(message: OrderByBroadcast, reader: jspb.BinaryReader): OrderByBroadcast;
}

export namespace OrderByBroadcast {
  export type AsObject = {
    field: BroadcastFilter.Field,
    orderBy: OrderBy,
  }
}

export class Roster extends jspb.Message {
  getRosteringId(): number;
  setRosteringId(value: number): Roster;

  getAifsId(): number;
  setAifsId(value: number): Roster;

  getStartTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStartTime(value?: google_protobuf_timestamp_pb.Timestamp): Roster;
  hasStartTime(): boolean;
  clearStartTime(): Roster;

  getEndTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setEndTime(value?: google_protobuf_timestamp_pb.Timestamp): Roster;
  hasEndTime(): boolean;
  clearEndTime(): Roster;

  getClientsList(): Array<AIFSClientRoster>;
  setClientsList(value: Array<AIFSClientRoster>): Roster;
  clearClientsList(): Roster;
  addClients(value?: AIFSClientRoster, index?: number): AIFSClientRoster;

  getGuardAssignedList(): Array<RosterAssignement>;
  setGuardAssignedList(value: Array<RosterAssignement>): Roster;
  clearGuardAssignedList(): Roster;
  addGuardAssigned(value?: RosterAssignement, index?: number): RosterAssignement;

  getIsDefault(): boolean;
  setIsDefault(value: boolean): Roster;

  getStatus(): Roster.Status;
  setStatus(value: Roster.Status): Roster;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Roster.AsObject;
  static toObject(includeInstance: boolean, msg: Roster): Roster.AsObject;
  static serializeBinaryToWriter(message: Roster, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Roster;
  static deserializeBinaryFromReader(message: Roster, reader: jspb.BinaryReader): Roster;
}

export namespace Roster {
  export type AsObject = {
    rosteringId: number,
    aifsId: number,
    startTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    endTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    clientsList: Array<AIFSClientRoster.AsObject>,
    guardAssignedList: Array<RosterAssignement.AsObject>,
    isDefault: boolean,
    status: Roster.Status,
  }

  export enum Status { 
    PENDING = 0,
    CONFIRMED = 1,
    REJECTED = 2,
  }
}

export class AIFSClientRoster extends jspb.Message {
  getAifsClientRosterId(): number;
  setAifsClientRosterId(value: number): AIFSClientRoster;

  getClient(): Client | undefined;
  setClient(value?: Client): AIFSClientRoster;
  hasClient(): boolean;
  clearClient(): AIFSClientRoster;

  getPatrolOrder(): number;
  setPatrolOrder(value: number): AIFSClientRoster;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AIFSClientRoster.AsObject;
  static toObject(includeInstance: boolean, msg: AIFSClientRoster): AIFSClientRoster.AsObject;
  static serializeBinaryToWriter(message: AIFSClientRoster, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AIFSClientRoster;
  static deserializeBinaryFromReader(message: AIFSClientRoster, reader: jspb.BinaryReader): AIFSClientRoster;
}

export namespace AIFSClientRoster {
  export type AsObject = {
    aifsClientRosterId: number,
    client?: Client.AsObject,
    patrolOrder: number,
  }
}

export class RosterAssignement extends jspb.Message {
  getRosterAssignmentId(): number;
  setRosterAssignmentId(value: number): RosterAssignement;

  getGuardAssigned(): EmployeeEvaluation | undefined;
  setGuardAssigned(value?: EmployeeEvaluation): RosterAssignement;
  hasGuardAssigned(): boolean;
  clearGuardAssigned(): RosterAssignement;

  getCustomStartTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCustomStartTime(value?: google_protobuf_timestamp_pb.Timestamp): RosterAssignement;
  hasCustomStartTime(): boolean;
  clearCustomStartTime(): RosterAssignement;

  getCustomEndTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCustomEndTime(value?: google_protobuf_timestamp_pb.Timestamp): RosterAssignement;
  hasCustomEndTime(): boolean;
  clearCustomEndTime(): RosterAssignement;

  getConfirmed(): boolean;
  setConfirmed(value: boolean): RosterAssignement;

  getAttended(): boolean;
  setAttended(value: boolean): RosterAssignement;

  getAttendanceTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setAttendanceTime(value?: google_protobuf_timestamp_pb.Timestamp): RosterAssignement;
  hasAttendanceTime(): boolean;
  clearAttendanceTime(): RosterAssignement;

  getIsAssigned(): boolean;
  setIsAssigned(value: boolean): RosterAssignement;

  getRejected(): boolean;
  setRejected(value: boolean): RosterAssignement;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RosterAssignement.AsObject;
  static toObject(includeInstance: boolean, msg: RosterAssignement): RosterAssignement.AsObject;
  static serializeBinaryToWriter(message: RosterAssignement, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RosterAssignement;
  static deserializeBinaryFromReader(message: RosterAssignement, reader: jspb.BinaryReader): RosterAssignement;
}

export namespace RosterAssignement {
  export type AsObject = {
    rosterAssignmentId: number,
    guardAssigned?: EmployeeEvaluation.AsObject,
    customStartTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    customEndTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    confirmed: boolean,
    attended: boolean,
    attendanceTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    isAssigned: boolean,
    rejected: boolean,
  }
}

export class BulkRosters extends jspb.Message {
  getRostersList(): Array<Roster>;
  setRostersList(value: Array<Roster>): BulkRosters;
  clearRostersList(): BulkRosters;
  addRosters(value?: Roster, index?: number): Roster;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BulkRosters.AsObject;
  static toObject(includeInstance: boolean, msg: BulkRosters): BulkRosters.AsObject;
  static serializeBinaryToWriter(message: BulkRosters, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BulkRosters;
  static deserializeBinaryFromReader(message: BulkRosters, reader: jspb.BinaryReader): BulkRosters;
}

export namespace BulkRosters {
  export type AsObject = {
    rostersList: Array<Roster.AsObject>,
  }
}

export class RosterResponse extends jspb.Message {
  getResponse(): Response | undefined;
  setResponse(value?: Response): RosterResponse;
  hasResponse(): boolean;
  clearResponse(): RosterResponse;

  getRoster(): Roster | undefined;
  setRoster(value?: Roster): RosterResponse;
  hasRoster(): boolean;
  clearRoster(): RosterResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RosterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RosterResponse): RosterResponse.AsObject;
  static serializeBinaryToWriter(message: RosterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RosterResponse;
  static deserializeBinaryFromReader(message: RosterResponse, reader: jspb.BinaryReader): RosterResponse;
}

export namespace RosterResponse {
  export type AsObject = {
    response?: Response.AsObject,
    roster?: Roster.AsObject,
  }
}

export class RosterFilter extends jspb.Message {
  getField(): RosterFilter.Field;
  setField(value: RosterFilter.Field): RosterFilter;

  getComparisons(): Filter | undefined;
  setComparisons(value?: Filter): RosterFilter;
  hasComparisons(): boolean;
  clearComparisons(): RosterFilter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RosterFilter.AsObject;
  static toObject(includeInstance: boolean, msg: RosterFilter): RosterFilter.AsObject;
  static serializeBinaryToWriter(message: RosterFilter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RosterFilter;
  static deserializeBinaryFromReader(message: RosterFilter, reader: jspb.BinaryReader): RosterFilter;
}

export namespace RosterFilter {
  export type AsObject = {
    field: RosterFilter.Field,
    comparisons?: Filter.AsObject,
  }

  export enum Field { 
    ROSTER_ID = 0,
    ROSTER_ASSIGNMENT_ID = 1,
    ROSTER_AIFS_CLIENT_ID = 2,
    AIFS_ID = 3,
    GUARD_ASSIGNED_ID = 4,
    CLIENT_ID = 5,
    GUARD_ASSIGNMENT_CONFIRMATION = 6,
    GUARD_ASSIGNMENT_ATTENDED = 7,
    START_TIME = 8,
    END_TIME = 9,
    IS_ASSIGNED = 10,
    DEFAULT_ROSTERING_DAY_OF_WEEK = 11,
    GUARD_ASSIGNMENT_REJECTION = 12,
  }
}

export class RosterQuery extends jspb.Message {
  getFiltersList(): Array<RosterFilter>;
  setFiltersList(value: Array<RosterFilter>): RosterQuery;
  clearFiltersList(): RosterQuery;
  addFilters(value?: RosterFilter, index?: number): RosterFilter;

  getLimit(): number;
  setLimit(value: number): RosterQuery;

  getSkip(): number;
  setSkip(value: number): RosterQuery;

  getOrderBy(): OrderByRoster | undefined;
  setOrderBy(value?: OrderByRoster): RosterQuery;
  hasOrderBy(): boolean;
  clearOrderBy(): RosterQuery;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RosterQuery.AsObject;
  static toObject(includeInstance: boolean, msg: RosterQuery): RosterQuery.AsObject;
  static serializeBinaryToWriter(message: RosterQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RosterQuery;
  static deserializeBinaryFromReader(message: RosterQuery, reader: jspb.BinaryReader): RosterQuery;
}

export namespace RosterQuery {
  export type AsObject = {
    filtersList: Array<RosterFilter.AsObject>,
    limit: number,
    skip: number,
    orderBy?: OrderByRoster.AsObject,
  }
}

export class OrderByRoster extends jspb.Message {
  getField(): RosterFilter.Field;
  setField(value: RosterFilter.Field): OrderByRoster;

  getOrderBy(): OrderBy;
  setOrderBy(value: OrderBy): OrderByRoster;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OrderByRoster.AsObject;
  static toObject(includeInstance: boolean, msg: OrderByRoster): OrderByRoster.AsObject;
  static serializeBinaryToWriter(message: OrderByRoster, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OrderByRoster;
  static deserializeBinaryFromReader(message: OrderByRoster, reader: jspb.BinaryReader): OrderByRoster;
}

export namespace OrderByRoster {
  export type AsObject = {
    field: RosterFilter.Field,
    orderBy: OrderBy,
  }
}

export class AvailabilityQuery extends jspb.Message {
  getStartTime(): string;
  setStartTime(value: string): AvailabilityQuery;

  getEndTime(): string;
  setEndTime(value: string): AvailabilityQuery;

  getLimit(): number;
  setLimit(value: number): AvailabilityQuery;

  getSkip(): number;
  setSkip(value: number): AvailabilityQuery;

  getFiltersList(): Array<AvailabilityFilter>;
  setFiltersList(value: Array<AvailabilityFilter>): AvailabilityQuery;
  clearFiltersList(): AvailabilityQuery;
  addFilters(value?: AvailabilityFilter, index?: number): AvailabilityFilter;

  getOrderBy(): OrderByQuery | undefined;
  setOrderBy(value?: OrderByQuery): AvailabilityQuery;
  hasOrderBy(): boolean;
  clearOrderBy(): AvailabilityQuery;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AvailabilityQuery.AsObject;
  static toObject(includeInstance: boolean, msg: AvailabilityQuery): AvailabilityQuery.AsObject;
  static serializeBinaryToWriter(message: AvailabilityQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AvailabilityQuery;
  static deserializeBinaryFromReader(message: AvailabilityQuery, reader: jspb.BinaryReader): AvailabilityQuery;
}

export namespace AvailabilityQuery {
  export type AsObject = {
    startTime: string,
    endTime: string,
    limit: number,
    skip: number,
    filtersList: Array<AvailabilityFilter.AsObject>,
    orderBy?: OrderByQuery.AsObject,
  }
}

export class OrderByQuery extends jspb.Message {
  getField(): AvailabilityFilter.Field;
  setField(value: AvailabilityFilter.Field): OrderByQuery;

  getOrderBy(): OrderBy;
  setOrderBy(value: OrderBy): OrderByQuery;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OrderByQuery.AsObject;
  static toObject(includeInstance: boolean, msg: OrderByQuery): OrderByQuery.AsObject;
  static serializeBinaryToWriter(message: OrderByQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OrderByQuery;
  static deserializeBinaryFromReader(message: OrderByQuery, reader: jspb.BinaryReader): OrderByQuery;
}

export namespace OrderByQuery {
  export type AsObject = {
    field: AvailabilityFilter.Field,
    orderBy: OrderBy,
  }
}

export class AvailabilityFilter extends jspb.Message {
  getField(): AvailabilityFilter.Field;
  setField(value: AvailabilityFilter.Field): AvailabilityFilter;

  getComparisons(): Filter | undefined;
  setComparisons(value?: Filter): AvailabilityFilter;
  hasComparisons(): boolean;
  clearComparisons(): AvailabilityFilter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AvailabilityFilter.AsObject;
  static toObject(includeInstance: boolean, msg: AvailabilityFilter): AvailabilityFilter.AsObject;
  static serializeBinaryToWriter(message: AvailabilityFilter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AvailabilityFilter;
  static deserializeBinaryFromReader(message: AvailabilityFilter, reader: jspb.BinaryReader): AvailabilityFilter;
}

export namespace AvailabilityFilter {
  export type AsObject = {
    field: AvailabilityFilter.Field,
    comparisons?: Filter.AsObject,
  }

  export enum Field { 
    AVAILABILITY_ID = 0,
    WEEK = 1,
    YEAR = 2,
    GUARD_ID = 3,
    SUN = 4,
    MON = 5,
    TUES = 6,
    WED = 7,
    THURS = 8,
    FRI = 9,
    SAT = 10,
    NEXT_SUN = 11,
  }
}

export class EmployeeEvaluationResponse extends jspb.Message {
  getResponse(): Response | undefined;
  setResponse(value?: Response): EmployeeEvaluationResponse;
  hasResponse(): boolean;
  clearResponse(): EmployeeEvaluationResponse;

  getEmployee(): EmployeeEvaluation | undefined;
  setEmployee(value?: EmployeeEvaluation): EmployeeEvaluationResponse;
  hasEmployee(): boolean;
  clearEmployee(): EmployeeEvaluationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EmployeeEvaluationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EmployeeEvaluationResponse): EmployeeEvaluationResponse.AsObject;
  static serializeBinaryToWriter(message: EmployeeEvaluationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EmployeeEvaluationResponse;
  static deserializeBinaryFromReader(message: EmployeeEvaluationResponse, reader: jspb.BinaryReader): EmployeeEvaluationResponse;
}

export namespace EmployeeEvaluationResponse {
  export type AsObject = {
    response?: Response.AsObject,
    employee?: EmployeeEvaluation.AsObject,
  }
}

export class EmployeeEvaluation extends jspb.Message {
  getEmployee(): User | undefined;
  setEmployee(value?: User): EmployeeEvaluation;
  hasEmployee(): boolean;
  clearEmployee(): EmployeeEvaluation;

  getEmployeeScore(): number;
  setEmployeeScore(value: number): EmployeeEvaluation;

  getIsAvailable(): boolean;
  setIsAvailable(value: boolean): EmployeeEvaluation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EmployeeEvaluation.AsObject;
  static toObject(includeInstance: boolean, msg: EmployeeEvaluation): EmployeeEvaluation.AsObject;
  static serializeBinaryToWriter(message: EmployeeEvaluation, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EmployeeEvaluation;
  static deserializeBinaryFromReader(message: EmployeeEvaluation, reader: jspb.BinaryReader): EmployeeEvaluation;
}

export namespace EmployeeEvaluation {
  export type AsObject = {
    employee?: User.AsObject,
    employeeScore: number,
    isAvailable: boolean,
  }
}

export class Response extends jspb.Message {
  getType(): Response.Type;
  setType(value: Response.Type): Response;

  getErrorMessage(): string;
  setErrorMessage(value: string): Response;

  getPrimaryKey(): number;
  setPrimaryKey(value: number): Response;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Response.AsObject;
  static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
  static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Response;
  static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
}

export namespace Response {
  export type AsObject = {
    type: Response.Type,
    errorMessage: string,
    primaryKey: number,
  }

  export enum Type { 
    ACK = 0,
    ERROR = 1,
  }
}

export class Filter extends jspb.Message {
  getComparison(): Filter.Comparisons;
  setComparison(value: Filter.Comparisons): Filter;

  getValue(): string;
  setValue(value: string): Filter;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Filter.AsObject;
  static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
  static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Filter;
  static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
}

export namespace Filter {
  export type AsObject = {
    comparison: Filter.Comparisons,
    value: string,
  }

  export enum Comparisons { 
    GREATER = 0,
    GREATER_EQ = 1,
    EQUAL = 2,
    LESSER_EQ = 3,
    LESSER = 4,
    CONTAINS = 5,
    IN = 6,
    NOT_IN = 7,
  }
}

export enum OrderBy { 
  ASC = 0,
  DESC = 1,
}
