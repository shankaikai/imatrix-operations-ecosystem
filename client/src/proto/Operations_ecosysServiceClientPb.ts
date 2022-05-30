/**
 * @fileoverview gRPC-Web generated client stub for operations_ecosys
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

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
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorAddUser = new grpcWeb.MethodDescriptor(
    '/operations_ecosys.AdminServices/AddUser',
    grpcWeb.MethodType.UNARY,
    operations_ecosys_pb.User,
    operations_ecosys_pb.Response,
    (request: operations_ecosys_pb.User) => {
      return request.serializeBinary();
    },
    operations_ecosys_pb.Response.deserializeBinary
  );

  addUser(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null): Promise<operations_ecosys_pb.Response>;

  addUser(
    request: operations_ecosys_pb.User,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: operations_ecosys_pb.Response) => void): grpcWeb.ClientReadableStream<operations_ecosys_pb.Response>;

  addUser(
    request: operations_ecosys_pb.User,
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
    options['format'] = 'binary';

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

}

