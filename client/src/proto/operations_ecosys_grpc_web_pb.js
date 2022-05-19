/**
 * @fileoverview gRPC-Web generated client stub for operations_ecosys
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.operations_ecosys = require('./operations_ecosys_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.operations_ecosys.AdminServicesClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.operations_ecosys.AdminServicesPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.operations_ecosys.User,
 *   !proto.operations_ecosys.Response>}
 */
const methodDescriptor_AdminServices_AddUser = new grpc.web.MethodDescriptor(
  '/operations_ecosys.AdminServices/AddUser',
  grpc.web.MethodType.UNARY,
  proto.operations_ecosys.User,
  proto.operations_ecosys.Response,
  /**
   * @param {!proto.operations_ecosys.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.operations_ecosys.Response.deserializeBinary
);


/**
 * @param {!proto.operations_ecosys.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.operations_ecosys.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.operations_ecosys.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.operations_ecosys.AdminServicesClient.prototype.addUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/operations_ecosys.AdminServices/AddUser',
      request,
      metadata || {},
      methodDescriptor_AdminServices_AddUser,
      callback);
};


/**
 * @param {!proto.operations_ecosys.User} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.operations_ecosys.Response>}
 *     Promise that resolves to the response
 */
proto.operations_ecosys.AdminServicesPromiseClient.prototype.addUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/operations_ecosys.AdminServices/AddUser',
      request,
      metadata || {},
      methodDescriptor_AdminServices_AddUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.operations_ecosys.User,
 *   !proto.operations_ecosys.Response>}
 */
const methodDescriptor_AdminServices_UpdateUser = new grpc.web.MethodDescriptor(
  '/operations_ecosys.AdminServices/UpdateUser',
  grpc.web.MethodType.UNARY,
  proto.operations_ecosys.User,
  proto.operations_ecosys.Response,
  /**
   * @param {!proto.operations_ecosys.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.operations_ecosys.Response.deserializeBinary
);


/**
 * @param {!proto.operations_ecosys.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.operations_ecosys.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.operations_ecosys.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.operations_ecosys.AdminServicesClient.prototype.updateUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/operations_ecosys.AdminServices/UpdateUser',
      request,
      metadata || {},
      methodDescriptor_AdminServices_UpdateUser,
      callback);
};


/**
 * @param {!proto.operations_ecosys.User} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.operations_ecosys.Response>}
 *     Promise that resolves to the response
 */
proto.operations_ecosys.AdminServicesPromiseClient.prototype.updateUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/operations_ecosys.AdminServices/UpdateUser',
      request,
      metadata || {},
      methodDescriptor_AdminServices_UpdateUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.operations_ecosys.User,
 *   !proto.operations_ecosys.Response>}
 */
const methodDescriptor_AdminServices_DeleteUser = new grpc.web.MethodDescriptor(
  '/operations_ecosys.AdminServices/DeleteUser',
  grpc.web.MethodType.UNARY,
  proto.operations_ecosys.User,
  proto.operations_ecosys.Response,
  /**
   * @param {!proto.operations_ecosys.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.operations_ecosys.Response.deserializeBinary
);


/**
 * @param {!proto.operations_ecosys.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.operations_ecosys.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.operations_ecosys.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.operations_ecosys.AdminServicesClient.prototype.deleteUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/operations_ecosys.AdminServices/DeleteUser',
      request,
      metadata || {},
      methodDescriptor_AdminServices_DeleteUser,
      callback);
};


/**
 * @param {!proto.operations_ecosys.User} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.operations_ecosys.Response>}
 *     Promise that resolves to the response
 */
proto.operations_ecosys.AdminServicesPromiseClient.prototype.deleteUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/operations_ecosys.AdminServices/DeleteUser',
      request,
      metadata || {},
      methodDescriptor_AdminServices_DeleteUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.operations_ecosys.Query,
 *   !proto.operations_ecosys.BulkUsers>}
 */
const methodDescriptor_AdminServices_FindUsers = new grpc.web.MethodDescriptor(
  '/operations_ecosys.AdminServices/FindUsers',
  grpc.web.MethodType.UNARY,
  proto.operations_ecosys.Query,
  proto.operations_ecosys.BulkUsers,
  /**
   * @param {!proto.operations_ecosys.Query} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.operations_ecosys.BulkUsers.deserializeBinary
);


/**
 * @param {!proto.operations_ecosys.Query} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.operations_ecosys.BulkUsers)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.operations_ecosys.BulkUsers>|undefined}
 *     The XHR Node Readable Stream
 */
proto.operations_ecosys.AdminServicesClient.prototype.findUsers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/operations_ecosys.AdminServices/FindUsers',
      request,
      metadata || {},
      methodDescriptor_AdminServices_FindUsers,
      callback);
};


/**
 * @param {!proto.operations_ecosys.Query} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.operations_ecosys.BulkUsers>}
 *     Promise that resolves to the response
 */
proto.operations_ecosys.AdminServicesPromiseClient.prototype.findUsers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/operations_ecosys.AdminServices/FindUsers',
      request,
      metadata || {},
      methodDescriptor_AdminServices_FindUsers);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.operations_ecosys.BroadcastServicesClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.operations_ecosys.BroadcastServicesPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.operations_ecosys.Broadcast,
 *   !proto.operations_ecosys.Response>}
 */
const methodDescriptor_BroadcastServices_AddBroadcast = new grpc.web.MethodDescriptor(
  '/operations_ecosys.BroadcastServices/AddBroadcast',
  grpc.web.MethodType.UNARY,
  proto.operations_ecosys.Broadcast,
  proto.operations_ecosys.Response,
  /**
   * @param {!proto.operations_ecosys.Broadcast} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.operations_ecosys.Response.deserializeBinary
);


/**
 * @param {!proto.operations_ecosys.Broadcast} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.operations_ecosys.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.operations_ecosys.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.operations_ecosys.BroadcastServicesClient.prototype.addBroadcast =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/operations_ecosys.BroadcastServices/AddBroadcast',
      request,
      metadata || {},
      methodDescriptor_BroadcastServices_AddBroadcast,
      callback);
};


/**
 * @param {!proto.operations_ecosys.Broadcast} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.operations_ecosys.Response>}
 *     Promise that resolves to the response
 */
proto.operations_ecosys.BroadcastServicesPromiseClient.prototype.addBroadcast =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/operations_ecosys.BroadcastServices/AddBroadcast',
      request,
      metadata || {},
      methodDescriptor_BroadcastServices_AddBroadcast);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.operations_ecosys.Broadcast,
 *   !proto.operations_ecosys.Response>}
 */
const methodDescriptor_BroadcastServices_UpdateBroadcast = new grpc.web.MethodDescriptor(
  '/operations_ecosys.BroadcastServices/UpdateBroadcast',
  grpc.web.MethodType.UNARY,
  proto.operations_ecosys.Broadcast,
  proto.operations_ecosys.Response,
  /**
   * @param {!proto.operations_ecosys.Broadcast} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.operations_ecosys.Response.deserializeBinary
);


/**
 * @param {!proto.operations_ecosys.Broadcast} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.operations_ecosys.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.operations_ecosys.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.operations_ecosys.BroadcastServicesClient.prototype.updateBroadcast =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/operations_ecosys.BroadcastServices/UpdateBroadcast',
      request,
      metadata || {},
      methodDescriptor_BroadcastServices_UpdateBroadcast,
      callback);
};


/**
 * @param {!proto.operations_ecosys.Broadcast} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.operations_ecosys.Response>}
 *     Promise that resolves to the response
 */
proto.operations_ecosys.BroadcastServicesPromiseClient.prototype.updateBroadcast =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/operations_ecosys.BroadcastServices/UpdateBroadcast',
      request,
      metadata || {},
      methodDescriptor_BroadcastServices_UpdateBroadcast);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.operations_ecosys.Broadcast,
 *   !proto.operations_ecosys.Response>}
 */
const methodDescriptor_BroadcastServices_DeleteBroadcast = new grpc.web.MethodDescriptor(
  '/operations_ecosys.BroadcastServices/DeleteBroadcast',
  grpc.web.MethodType.UNARY,
  proto.operations_ecosys.Broadcast,
  proto.operations_ecosys.Response,
  /**
   * @param {!proto.operations_ecosys.Broadcast} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.operations_ecosys.Response.deserializeBinary
);


/**
 * @param {!proto.operations_ecosys.Broadcast} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.operations_ecosys.Response)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.operations_ecosys.Response>|undefined}
 *     The XHR Node Readable Stream
 */
proto.operations_ecosys.BroadcastServicesClient.prototype.deleteBroadcast =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/operations_ecosys.BroadcastServices/DeleteBroadcast',
      request,
      metadata || {},
      methodDescriptor_BroadcastServices_DeleteBroadcast,
      callback);
};


/**
 * @param {!proto.operations_ecosys.Broadcast} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.operations_ecosys.Response>}
 *     Promise that resolves to the response
 */
proto.operations_ecosys.BroadcastServicesPromiseClient.prototype.deleteBroadcast =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/operations_ecosys.BroadcastServices/DeleteBroadcast',
      request,
      metadata || {},
      methodDescriptor_BroadcastServices_DeleteBroadcast);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.operations_ecosys.Query,
 *   !proto.operations_ecosys.BulkBroadcasts>}
 */
const methodDescriptor_BroadcastServices_FindBroadcasts = new grpc.web.MethodDescriptor(
  '/operations_ecosys.BroadcastServices/FindBroadcasts',
  grpc.web.MethodType.UNARY,
  proto.operations_ecosys.Query,
  proto.operations_ecosys.BulkBroadcasts,
  /**
   * @param {!proto.operations_ecosys.Query} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.operations_ecosys.BulkBroadcasts.deserializeBinary
);


/**
 * @param {!proto.operations_ecosys.Query} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.operations_ecosys.BulkBroadcasts)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.operations_ecosys.BulkBroadcasts>|undefined}
 *     The XHR Node Readable Stream
 */
proto.operations_ecosys.BroadcastServicesClient.prototype.findBroadcasts =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/operations_ecosys.BroadcastServices/FindBroadcasts',
      request,
      metadata || {},
      methodDescriptor_BroadcastServices_FindBroadcasts,
      callback);
};


/**
 * @param {!proto.operations_ecosys.Query} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.operations_ecosys.BulkBroadcasts>}
 *     Promise that resolves to the response
 */
proto.operations_ecosys.BroadcastServicesPromiseClient.prototype.findBroadcasts =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/operations_ecosys.BroadcastServices/FindBroadcasts',
      request,
      metadata || {},
      methodDescriptor_BroadcastServices_FindBroadcasts);
};


module.exports = proto.operations_ecosys;

