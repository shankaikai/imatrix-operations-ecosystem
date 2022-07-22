/**
 * @fileoverview gRPC-Web generated client stub for operations_ecosys
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as iot_prototype_pb from './iot_prototype_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as operations_ecosys_pb from './operations_ecosys_pb';


export class AdminServicesClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorAddUser = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/AddUser',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.FullUser,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.FullUser) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  addUser(
    request: operations_ecosys_pb.FullUser,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  addUser(
    request: operations_ecosys_pb.FullUser,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  addUser(
    request: operations_ecosys_pb.FullUser,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.AdminServices/AddUser',
        request,
        metadata || {},
        this.methodDescriptorAddUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.AdminServices/AddUser',
    request,
    metadata || {},
    this.methodDescriptorAddUser);
  }

  methodDescriptorUpdateUser = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/UpdateUser',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.User,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.User) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  updateUser(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  updateUser(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  updateUser(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.AdminServices/UpdateUser',
        request,
        metadata || {},
        this.methodDescriptorUpdateUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.AdminServices/UpdateUser',
    request,
    metadata || {},
    this.methodDescriptorUpdateUser);
  }

  methodDescriptorDeleteUser = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/DeleteUser',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.User,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.User) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  deleteUser(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  deleteUser(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  deleteUser(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.AdminServices/DeleteUser',
        request,
        metadata || {},
        this.methodDescriptorDeleteUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.AdminServices/DeleteUser',
    request,
    metadata || {},
    this.methodDescriptorDeleteUser);
  }

  methodDescriptorFindUsers = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/FindUsers',
    grpcWeb.MethodType.SERVER_STREAMING,
    operations_ecosys_pb.UserQuery,
    operations_ecosys_pb.UsersResponse,
    (request: operations_ecosys_pb.UserQuery) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.UsersResponse.deserializeBinary
  );

  findUsers(
    request: operations_ecosys_pb.UserQuery,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<operations_ecosys_pb.UsersResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/operations_ecosys.AdminServices/FindUsers',
      request,
      metadata || {},
      this.methodDescriptorFindUsers);
  }

  methodDescriptorAddClient = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/AddClient',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.Client,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.Client) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  addClient(
    request: operations_ecosys_pb.Client,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  addClient(
    request: operations_ecosys_pb.Client,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  addClient(
    request: operations_ecosys_pb.Client,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.AdminServices/AddClient',
        request,
        metadata || {},
        this.methodDescriptorAddClient,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.AdminServices/AddClient',
    request,
    metadata || {},
    this.methodDescriptorAddClient);
  }

  methodDescriptorUpdateClient = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/UpdateClient',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.Client,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.Client) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  updateClient(
    request: operations_ecosys_pb.Client,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  updateClient(
    request: operations_ecosys_pb.Client,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  updateClient(
    request: operations_ecosys_pb.Client,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.AdminServices/UpdateClient',
        request,
        metadata || {},
        this.methodDescriptorUpdateClient,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.AdminServices/UpdateClient',
    request,
    metadata || {},
    this.methodDescriptorUpdateClient);
  }

  methodDescriptorDeleteClient = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/DeleteClient',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.Client,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.Client) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  deleteClient(
    request: operations_ecosys_pb.Client,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  deleteClient(
    request: operations_ecosys_pb.Client,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  deleteClient(
    request: operations_ecosys_pb.Client,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.AdminServices/DeleteClient',
        request,
        metadata || {},
        this.methodDescriptorDeleteClient,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.AdminServices/DeleteClient',
    request,
    metadata || {},
    this.methodDescriptorDeleteClient);
  }

  methodDescriptorFindClients = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/FindClients',
    grpcWeb.MethodType.SERVER_STREAMING,
    operations_ecosys_pb.ClientQuery,
    operations_ecosys_pb.ClientResponse,
    (request: operations_ecosys_pb.ClientQuery) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.ClientResponse.deserializeBinary
  );

  findClients(
    request: operations_ecosys_pb.ClientQuery,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<operations_ecosys_pb.ClientResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/operations_ecosys.AdminServices/FindClients',
      request,
      metadata || {},
      this.methodDescriptorFindClients);
  }

  methodDescriptorGetWANonce = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/GetWANonce',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.User,
    operations_ecosys_pb.ResponseNonce,
    (request: operations_ecosys_pb.User) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.ResponseNonce.deserializeBinary
  );

  getWANonce(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.ResponseNonce>;

  getWANonce(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.ResponseNonce) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.ResponseNonce>;

  getWANonce(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.ResponseNonce) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.AdminServices/GetWANonce',
        request,
        metadata || {},
        this.methodDescriptorGetWANonce,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.AdminServices/GetWANonce',
    request,
    metadata || {},
    this.methodDescriptorGetWANonce);
  }

  methodDescriptorGetSecurityString = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/GetSecurityString',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.User,
    operations_ecosys_pb.SecurityStringResponse,
    (request: operations_ecosys_pb.User) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.SecurityStringResponse.deserializeBinary
  );

  getSecurityString(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.SecurityStringResponse>;

  getSecurityString(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.SecurityStringResponse) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.SecurityStringResponse>;

  getSecurityString(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.SecurityStringResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.AdminServices/GetSecurityString',
        request,
        metadata || {},
        this.methodDescriptorGetSecurityString,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.AdminServices/GetSecurityString',
    request,
    metadata || {},
    this.methodDescriptorGetSecurityString);
  }

  methodDescriptorAuthenticateUser = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/AuthenticateUser',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.LoginRequest,
    operations_ecosys_pb.UserTokenResponse,
    (request: operations_ecosys_pb.LoginRequest) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.UserTokenResponse.deserializeBinary
  );

  authenticateUser(
    request: operations_ecosys_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.UserTokenResponse>;

  authenticateUser(
    request: operations_ecosys_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.UserTokenResponse) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.UserTokenResponse>;

  authenticateUser(
    request: operations_ecosys_pb.LoginRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.UserTokenResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.AdminServices/AuthenticateUser',
        request,
        metadata || {},
        this.methodDescriptorAuthenticateUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.AdminServices/AuthenticateUser',
    request,
    metadata || {},
    this.methodDescriptorAuthenticateUser);
  }

}

export class BroadcastServicesClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorAddBroadcast = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.BroadcastServices/AddBroadcast',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.Broadcast,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.Broadcast) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  addBroadcast(
    request: operations_ecosys_pb.Broadcast,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  addBroadcast(
    request: operations_ecosys_pb.Broadcast,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  addBroadcast(
    request: operations_ecosys_pb.Broadcast,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.BroadcastServices/AddBroadcast',
        request,
        metadata || {},
        this.methodDescriptorAddBroadcast,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.BroadcastServices/AddBroadcast',
    request,
    metadata || {},
    this.methodDescriptorAddBroadcast);
  }

  methodDescriptorUpdateBroadcast = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.BroadcastServices/UpdateBroadcast',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.Broadcast,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.Broadcast) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  updateBroadcast(
    request: operations_ecosys_pb.Broadcast,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  updateBroadcast(
    request: operations_ecosys_pb.Broadcast,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  updateBroadcast(
    request: operations_ecosys_pb.Broadcast,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.BroadcastServices/UpdateBroadcast',
        request,
        metadata || {},
        this.methodDescriptorUpdateBroadcast,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.BroadcastServices/UpdateBroadcast',
    request,
    metadata || {},
    this.methodDescriptorUpdateBroadcast);
  }

  methodDescriptorDeleteBroadcast = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.BroadcastServices/DeleteBroadcast',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.Broadcast,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.Broadcast) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  deleteBroadcast(
    request: operations_ecosys_pb.Broadcast,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  deleteBroadcast(
    request: operations_ecosys_pb.Broadcast,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  deleteBroadcast(
    request: operations_ecosys_pb.Broadcast,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.BroadcastServices/DeleteBroadcast',
        request,
        metadata || {},
        this.methodDescriptorDeleteBroadcast,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.BroadcastServices/DeleteBroadcast',
    request,
    metadata || {},
    this.methodDescriptorDeleteBroadcast);
  }

  methodDescriptorFindBroadcasts = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.BroadcastServices/FindBroadcasts',
    grpcWeb.MethodType.SERVER_STREAMING,
    operations_ecosys_pb.BroadcastQuery,
    operations_ecosys_pb.BroadcastResponse,
    (request: operations_ecosys_pb.BroadcastQuery) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.BroadcastResponse.deserializeBinary
  );

  findBroadcasts(
    request: operations_ecosys_pb.BroadcastQuery,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<operations_ecosys_pb.BroadcastResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/operations_ecosys.BroadcastServices/FindBroadcasts',
      request,
      metadata || {},
      this.methodDescriptorFindBroadcasts);
  }

  methodDescriptorUpdateBroadcastRecipient = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.BroadcastServices/UpdateBroadcastRecipient',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.BroadcastRecipient,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.BroadcastRecipient) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  updateBroadcastRecipient(
    request: operations_ecosys_pb.BroadcastRecipient,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  updateBroadcastRecipient(
    request: operations_ecosys_pb.BroadcastRecipient,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  updateBroadcastRecipient(
    request: operations_ecosys_pb.BroadcastRecipient,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.BroadcastServices/UpdateBroadcastRecipient',
        request,
        metadata || {},
        this.methodDescriptorUpdateBroadcastRecipient,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.BroadcastServices/UpdateBroadcastRecipient',
    request,
    metadata || {},
    this.methodDescriptorUpdateBroadcastRecipient);
  }

}

export class RosterServicesClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorAddRoster = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.RosterServices/AddRoster',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.BulkRosters,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.BulkRosters) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  addRoster(
    request: operations_ecosys_pb.BulkRosters,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  addRoster(
    request: operations_ecosys_pb.BulkRosters,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  addRoster(
    request: operations_ecosys_pb.BulkRosters,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.RosterServices/AddRoster',
        request,
        metadata || {},
        this.methodDescriptorAddRoster,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.RosterServices/AddRoster',
    request,
    metadata || {},
    this.methodDescriptorAddRoster);
  }

  methodDescriptorUpdateRoster = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.RosterServices/UpdateRoster',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.BulkRosters,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.BulkRosters) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  updateRoster(
    request: operations_ecosys_pb.BulkRosters,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  updateRoster(
    request: operations_ecosys_pb.BulkRosters,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  updateRoster(
    request: operations_ecosys_pb.BulkRosters,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.RosterServices/UpdateRoster',
        request,
        metadata || {},
        this.methodDescriptorUpdateRoster,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.RosterServices/UpdateRoster',
    request,
    metadata || {},
    this.methodDescriptorUpdateRoster);
  }

  methodDescriptorDeleteRoster = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.RosterServices/DeleteRoster',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.Roster,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.Roster) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  deleteRoster(
    request: operations_ecosys_pb.Roster,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  deleteRoster(
    request: operations_ecosys_pb.Roster,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  deleteRoster(
    request: operations_ecosys_pb.Roster,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.RosterServices/DeleteRoster',
        request,
        metadata || {},
        this.methodDescriptorDeleteRoster,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.RosterServices/DeleteRoster',
    request,
    metadata || {},
    this.methodDescriptorDeleteRoster);
  }

  methodDescriptorFindRosters = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.RosterServices/FindRosters',
    grpcWeb.MethodType.SERVER_STREAMING,
    operations_ecosys_pb.RosterQuery,
    operations_ecosys_pb.RosterResponse,
    (request: operations_ecosys_pb.RosterQuery) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.RosterResponse.deserializeBinary
  );

  findRosters(
    request: operations_ecosys_pb.RosterQuery,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<operations_ecosys_pb.RosterResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/operations_ecosys.RosterServices/FindRosters',
      request,
      metadata || {},
      this.methodDescriptorFindRosters);
  }

  methodDescriptorGetAvailableUsers = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.RosterServices/GetAvailableUsers',
    grpcWeb.MethodType.SERVER_STREAMING,
    operations_ecosys_pb.AvailabilityQuery,
    operations_ecosys_pb.EmployeeEvaluationResponse,
    (request: operations_ecosys_pb.AvailabilityQuery) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.EmployeeEvaluationResponse.deserializeBinary
  );

  getAvailableUsers(
    request: operations_ecosys_pb.AvailabilityQuery,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<operations_ecosys_pb.EmployeeEvaluationResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/operations_ecosys.RosterServices/GetAvailableUsers',
      request,
      metadata || {},
      this.methodDescriptorGetAvailableUsers);
  }

  methodDescriptorFindRosterAssignments = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.RosterServices/FindRosterAssignments',
    grpcWeb.MethodType.SERVER_STREAMING,
    operations_ecosys_pb.RosterQuery,
    operations_ecosys_pb.RosterAssignmentResponse,
    (request: operations_ecosys_pb.RosterQuery) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.RosterAssignmentResponse.deserializeBinary
  );

  findRosterAssignments(
    request: operations_ecosys_pb.RosterQuery,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<operations_ecosys_pb.RosterAssignmentResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/operations_ecosys.RosterServices/FindRosterAssignments',
      request,
      metadata || {},
      this.methodDescriptorFindRosterAssignments);
  }

  methodDescriptorUpdateRosterAssignment = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.RosterServices/UpdateRosterAssignment',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.RosterAssignement,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.RosterAssignement) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  updateRosterAssignment(
    request: operations_ecosys_pb.RosterAssignement,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  updateRosterAssignment(
    request: operations_ecosys_pb.RosterAssignement,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  updateRosterAssignment(
    request: operations_ecosys_pb.RosterAssignement,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.RosterServices/UpdateRosterAssignment',
        request,
        metadata || {},
        this.methodDescriptorUpdateRosterAssignment,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.RosterServices/UpdateRosterAssignment',
    request,
    metadata || {},
    this.methodDescriptorUpdateRosterAssignment);
  }

}

export class IncidentReportServicesClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorAddIncidentReport = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.IncidentReportServices/AddIncidentReport',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.IncidentReport,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.IncidentReport) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  addIncidentReport(
    request: operations_ecosys_pb.IncidentReport,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  addIncidentReport(
    request: operations_ecosys_pb.IncidentReport,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  addIncidentReport(
    request: operations_ecosys_pb.IncidentReport,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.IncidentReportServices/AddIncidentReport',
        request,
        metadata || {},
        this.methodDescriptorAddIncidentReport,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.IncidentReportServices/AddIncidentReport',
    request,
    metadata || {},
    this.methodDescriptorAddIncidentReport);
  }

  methodDescriptorUpdateIncidentReport = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.IncidentReportServices/UpdateIncidentReport',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.IncidentReport,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.IncidentReport) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  updateIncidentReport(
    request: operations_ecosys_pb.IncidentReport,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  updateIncidentReport(
    request: operations_ecosys_pb.IncidentReport,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  updateIncidentReport(
    request: operations_ecosys_pb.IncidentReport,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.IncidentReportServices/UpdateIncidentReport',
        request,
        metadata || {},
        this.methodDescriptorUpdateIncidentReport,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.IncidentReportServices/UpdateIncidentReport',
    request,
    metadata || {},
    this.methodDescriptorUpdateIncidentReport);
  }

  methodDescriptorDeleteIncidentReport = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.IncidentReportServices/DeleteIncidentReport',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.IncidentReport,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.IncidentReport) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  deleteIncidentReport(
    request: operations_ecosys_pb.IncidentReport,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  deleteIncidentReport(
    request: operations_ecosys_pb.IncidentReport,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  deleteIncidentReport(
    request: operations_ecosys_pb.IncidentReport,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.IncidentReportServices/DeleteIncidentReport',
        request,
        metadata || {},
        this.methodDescriptorDeleteIncidentReport,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.IncidentReportServices/DeleteIncidentReport',
    request,
    metadata || {},
    this.methodDescriptorDeleteIncidentReport);
  }

  methodDescriptorFindIncidentReports = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.IncidentReportServices/FindIncidentReports',
    grpcWeb.MethodType.SERVER_STREAMING,
    operations_ecosys_pb.IncidentReportQuery,
    operations_ecosys_pb.IncidentReportResponse,
    (request: operations_ecosys_pb.IncidentReportQuery) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.IncidentReportResponse.deserializeBinary
  );

  findIncidentReports(
    request: operations_ecosys_pb.IncidentReportQuery,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<operations_ecosys_pb.IncidentReportResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/operations_ecosys.IncidentReportServices/FindIncidentReports',
      request,
      metadata || {},
      this.methodDescriptorFindIncidentReports);
  }

}

export class CameraIotServicesClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorSetGateState = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.CameraIotServices/SetGateState',
    grpcWeb.MethodType.UNARY,
    iot_prototype_pb.GateState,
    operations_ecosys_pb.Response,
    (request: iot_prototype_pb.GateState) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  setGateState(
    request: iot_prototype_pb.GateState,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  setGateState(
    request: iot_prototype_pb.GateState,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  setGateState(
    request: iot_prototype_pb.GateState,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/operations_ecosys.CameraIotServices/SetGateState',
        request,
        metadata || {},
        this.methodDescriptorSetGateState,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/operations_ecosys.CameraIotServices/SetGateState',
    request,
    metadata || {},
    this.methodDescriptorSetGateState);
  }

  methodDescriptorGetIotState = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.CameraIotServices/GetIotState',
    grpcWeb.MethodType.SERVER_STREAMING,
    google_protobuf_empty_pb.Empty,
    operations_ecosys_pb.CameraIotResponse,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.CameraIotResponse.deserializeBinary
  );

  getIotState(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<operations_ecosys_pb.CameraIotResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/operations_ecosys.CameraIotServices/GetIotState',
      request,
      metadata || {},
      this.methodDescriptorGetIotState);
  }

}

