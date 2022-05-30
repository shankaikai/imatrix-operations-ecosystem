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

export class Broadcast extends jspb.Message {
  getBroadcastId(): number;
  setBroadcastId(value: number): Broadcast;

  getType(): Broadcast.BroadcastType;
  setType(value: Broadcast.BroadcastType): Broadcast;

  getTitle(): string;
  setTitle(value: string): Broadcast;

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

  getRecipientsList(): Array<BroadcastRecipient>;
  setRecipientsList(value: Array<BroadcastRecipient>): Broadcast;
  clearRecipientsList(): Broadcast;
  addRecipients(value?: BroadcastRecipient, index?: number): BroadcastRecipient;

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
    title: string,
    content: string,
    creationDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deadline?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    creator?: User.AsObject,
    recipientsList: Array<BroadcastRecipient.AsObject>,
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
    TITLE = 2,
    CONTENT = 3,
    CREATION_DATE = 4,
    DEADLINE = 5,
    CREATOR_ID = 6,
    RECEIPEIENT_ID = 7,
    NUM_RECEIPIENTS = 8,
    URGENCY = 9,
    AIFS_ID = 10,
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
  }
}

export enum OrderBy { 
  ASC = 0,
  DESC = 1,
}
