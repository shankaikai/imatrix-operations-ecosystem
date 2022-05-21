// package: operations_ecosys
// file: operations_ecosys.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class User extends jspb.Message {
  getUserId(): number;
  setUserId(value: number): void;

  getUserType(): User.UserTypeMap[keyof User.UserTypeMap];
  setUserType(value: User.UserTypeMap[keyof User.UserTypeMap]): void;

  getName(): string;
  setName(value: string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getPhoneNumber(): string;
  setPhoneNumber(value: string): void;

  getTelegramHandle(): string;
  setTelegramHandle(value: string): void;

  getUserSecurityImg(): string;
  setUserSecurityImg(value: string): void;

  getIsPartTimer(): boolean;
  setIsPartTimer(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    userId: number,
    userType: User.UserTypeMap[keyof User.UserTypeMap],
    name: string,
    email: string,
    phoneNumber: string,
    telegramHandle: string,
    userSecurityImg: string,
    isPartTimer: boolean,
  }

  export interface UserTypeMap {
    ISPECIALIST: 0;
    SECURITY_GUARD: 1;
    CONTROLLER: 2;
    MANAGER: 3;
  }

  export const UserType: UserTypeMap;
}

export class BulkUsers extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): Response | undefined;
  setResponse(value?: Response): void;

  clearUsersList(): void;
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): void;
  addUsers(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BulkUsers.AsObject;
  static toObject(includeInstance: boolean, msg: BulkUsers): BulkUsers.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BulkUsers, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BulkUsers;
  static deserializeBinaryFromReader(message: BulkUsers, reader: jspb.BinaryReader): BulkUsers;
}

export namespace BulkUsers {
  export type AsObject = {
    response?: Response.AsObject,
    usersList: Array<User.AsObject>,
  }
}

export class UserFilter extends jspb.Message {
  getField(): UserFilter.FieldMap[keyof UserFilter.FieldMap];
  setField(value: UserFilter.FieldMap[keyof UserFilter.FieldMap]): void;

  hasComparisons(): boolean;
  clearComparisons(): void;
  getComparisons(): Filter | undefined;
  setComparisons(value?: Filter): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserFilter.AsObject;
  static toObject(includeInstance: boolean, msg: UserFilter): UserFilter.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UserFilter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserFilter;
  static deserializeBinaryFromReader(message: UserFilter, reader: jspb.BinaryReader): UserFilter;
}

export namespace UserFilter {
  export type AsObject = {
    field: UserFilter.FieldMap[keyof UserFilter.FieldMap],
    comparisons?: Filter.AsObject,
  }

  export interface FieldMap {
    USER_ID: 0;
    TYPE: 1;
    NAME: 2;
    EMAIL: 3;
    PHONE_NUMBER: 4;
    TELEGRAM_HANDLE: 5;
    IS_PART_TIMER: 6;
  }

  export const Field: FieldMap;
}

export class UserQuery extends jspb.Message {
  clearFiltersList(): void;
  getFiltersList(): Array<UserFilter>;
  setFiltersList(value: Array<UserFilter>): void;
  addFilters(value?: UserFilter, index?: number): UserFilter;

  getLimit(): number;
  setLimit(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserQuery.AsObject;
  static toObject(includeInstance: boolean, msg: UserQuery): UserQuery.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UserQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserQuery;
  static deserializeBinaryFromReader(message: UserQuery, reader: jspb.BinaryReader): UserQuery;
}

export namespace UserQuery {
  export type AsObject = {
    filtersList: Array<UserFilter.AsObject>,
    limit: number,
  }
}

export class Broadcast extends jspb.Message {
  getBroadcastId(): number;
  setBroadcastId(value: number): void;

  getType(): Broadcast.BroadcastTypeMap[keyof Broadcast.BroadcastTypeMap];
  setType(value: Broadcast.BroadcastTypeMap[keyof Broadcast.BroadcastTypeMap]): void;

  getTitle(): string;
  setTitle(value: string): void;

  getContent(): string;
  setContent(value: string): void;

  hasCreationDate(): boolean;
  clearCreationDate(): void;
  getCreationDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreationDate(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeadline(): boolean;
  clearDeadline(): void;
  getDeadline(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeadline(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasCreator(): boolean;
  clearCreator(): void;
  getCreator(): User | undefined;
  setCreator(value?: User): void;

  clearReceipientsList(): void;
  getReceipientsList(): Array<User>;
  setReceipientsList(value: Array<User>): void;
  addReceipients(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Broadcast.AsObject;
  static toObject(includeInstance: boolean, msg: Broadcast): Broadcast.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Broadcast, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Broadcast;
  static deserializeBinaryFromReader(message: Broadcast, reader: jspb.BinaryReader): Broadcast;
}

export namespace Broadcast {
  export type AsObject = {
    broadcastId: number,
    type: Broadcast.BroadcastTypeMap[keyof Broadcast.BroadcastTypeMap],
    title: string,
    content: string,
    creationDate?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deadline?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    creator?: User.AsObject,
    receipientsList: Array<User.AsObject>,
  }

  export interface BroadcastTypeMap {
    ANNOUNCEMENT: 0;
    ASSIGNMENT: 1;
  }

  export const BroadcastType: BroadcastTypeMap;
}

export class BulkBroadcasts extends jspb.Message {
  hasResponse(): boolean;
  clearResponse(): void;
  getResponse(): Response | undefined;
  setResponse(value?: Response): void;

  clearBroadcastsList(): void;
  getBroadcastsList(): Array<Broadcast>;
  setBroadcastsList(value: Array<Broadcast>): void;
  addBroadcasts(value?: Broadcast, index?: number): Broadcast;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BulkBroadcasts.AsObject;
  static toObject(includeInstance: boolean, msg: BulkBroadcasts): BulkBroadcasts.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BulkBroadcasts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BulkBroadcasts;
  static deserializeBinaryFromReader(message: BulkBroadcasts, reader: jspb.BinaryReader): BulkBroadcasts;
}

export namespace BulkBroadcasts {
  export type AsObject = {
    response?: Response.AsObject,
    broadcastsList: Array<Broadcast.AsObject>,
  }
}

export class BroadcastFilter extends jspb.Message {
  getField(): BroadcastFilter.FieldMap[keyof BroadcastFilter.FieldMap];
  setField(value: BroadcastFilter.FieldMap[keyof BroadcastFilter.FieldMap]): void;

  hasComparisons(): boolean;
  clearComparisons(): void;
  getComparisons(): Filter | undefined;
  setComparisons(value?: Filter): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BroadcastFilter.AsObject;
  static toObject(includeInstance: boolean, msg: BroadcastFilter): BroadcastFilter.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BroadcastFilter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BroadcastFilter;
  static deserializeBinaryFromReader(message: BroadcastFilter, reader: jspb.BinaryReader): BroadcastFilter;
}

export namespace BroadcastFilter {
  export type AsObject = {
    field: BroadcastFilter.FieldMap[keyof BroadcastFilter.FieldMap],
    comparisons?: Filter.AsObject,
  }

  export interface FieldMap {
    BROADCAST_ID: 0;
    TYPE: 1;
    TITLE: 2;
    CONTENT: 3;
    CREATION_DATE: 4;
    DEADLINE: 5;
    CREATOR: 6;
    RECEIPEIENTS: 7;
    NUM_RECEIPIENTS: 8;
  }

  export const Field: FieldMap;
}

export class BroadcastQuery extends jspb.Message {
  clearFiltersList(): void;
  getFiltersList(): Array<BroadcastFilter>;
  setFiltersList(value: Array<BroadcastFilter>): void;
  addFilters(value?: BroadcastFilter, index?: number): BroadcastFilter;

  getLimit(): number;
  setLimit(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BroadcastQuery.AsObject;
  static toObject(includeInstance: boolean, msg: BroadcastQuery): BroadcastQuery.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BroadcastQuery, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BroadcastQuery;
  static deserializeBinaryFromReader(message: BroadcastQuery, reader: jspb.BinaryReader): BroadcastQuery;
}

export namespace BroadcastQuery {
  export type AsObject = {
    filtersList: Array<BroadcastFilter.AsObject>,
    limit: number,
  }
}

export class Response extends jspb.Message {
  getType(): Response.TypeMap[keyof Response.TypeMap];
  setType(value: Response.TypeMap[keyof Response.TypeMap]): void;

  getErrorMessage(): string;
  setErrorMessage(value: string): void;

  getPrimaryKey(): number;
  setPrimaryKey(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Response.AsObject;
  static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Response;
  static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
}

export namespace Response {
  export type AsObject = {
    type: Response.TypeMap[keyof Response.TypeMap],
    errorMessage: string,
    primaryKey: number,
  }

  export interface TypeMap {
    ACK: 0;
    ERROR: 1;
  }

  export const Type: TypeMap;
}

export class Filter extends jspb.Message {
  getComparison(): Filter.ComparisonsMap[keyof Filter.ComparisonsMap];
  setComparison(value: Filter.ComparisonsMap[keyof Filter.ComparisonsMap]): void;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Filter.AsObject;
  static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Filter;
  static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
}

export namespace Filter {
  export type AsObject = {
    comparison: Filter.ComparisonsMap[keyof Filter.ComparisonsMap],
    value: string,
  }

  export interface ComparisonsMap {
    GREATER: 0;
    GREATER_EQ: 1;
    EQUAL: 2;
    LESSER_EQ: 3;
    LESSER: 4;
    CONTAINS: 5;
  }

  export const Comparisons: ComparisonsMap;
}

