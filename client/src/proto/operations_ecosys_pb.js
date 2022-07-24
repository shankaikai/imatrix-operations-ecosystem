// source: operations_ecosys.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() {
  if (this) { return this; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  if (typeof self !== 'undefined') { return self; }
  return Function('return this')();
}.call(null));

var iot_prototype_pb = require('./iot_prototype_pb.js');
goog.object.extend(proto, iot_prototype_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
goog.object.extend(proto, google_protobuf_empty_pb);
var google_api_annotations_pb = require('./google/api/annotations_pb.js');
goog.object.extend(proto, google_api_annotations_pb);
goog.exportSymbol('proto.operations_ecosys.AIFSBroadcastRecipient', null, global);
goog.exportSymbol('proto.operations_ecosys.AIFSClientRoster', null, global);
goog.exportSymbol('proto.operations_ecosys.AvailabilityFilter', null, global);
goog.exportSymbol('proto.operations_ecosys.AvailabilityFilter.Field', null, global);
goog.exportSymbol('proto.operations_ecosys.AvailabilityQuery', null, global);
goog.exportSymbol('proto.operations_ecosys.Broadcast', null, global);
goog.exportSymbol('proto.operations_ecosys.Broadcast.BroadcastType', null, global);
goog.exportSymbol('proto.operations_ecosys.Broadcast.UrgencyType', null, global);
goog.exportSymbol('proto.operations_ecosys.BroadcastFilter', null, global);
goog.exportSymbol('proto.operations_ecosys.BroadcastFilter.Field', null, global);
goog.exportSymbol('proto.operations_ecosys.BroadcastQuery', null, global);
goog.exportSymbol('proto.operations_ecosys.BroadcastRecipient', null, global);
goog.exportSymbol('proto.operations_ecosys.BroadcastResponse', null, global);
goog.exportSymbol('proto.operations_ecosys.BulkRosters', null, global);
goog.exportSymbol('proto.operations_ecosys.Camera', null, global);
goog.exportSymbol('proto.operations_ecosys.CameraIot', null, global);
goog.exportSymbol('proto.operations_ecosys.CameraIot.MessageType', null, global);
goog.exportSymbol('proto.operations_ecosys.CameraIotFilter', null, global);
goog.exportSymbol('proto.operations_ecosys.CameraIotFilter.Field', null, global);
goog.exportSymbol('proto.operations_ecosys.CameraIotQuery', null, global);
goog.exportSymbol('proto.operations_ecosys.CameraIotResponse', null, global);
goog.exportSymbol('proto.operations_ecosys.Client', null, global);
goog.exportSymbol('proto.operations_ecosys.ClientFilter', null, global);
goog.exportSymbol('proto.operations_ecosys.ClientFilter.Field', null, global);
goog.exportSymbol('proto.operations_ecosys.ClientQuery', null, global);
goog.exportSymbol('proto.operations_ecosys.ClientResponse', null, global);
goog.exportSymbol('proto.operations_ecosys.EmployeeEvaluation', null, global);
goog.exportSymbol('proto.operations_ecosys.EmployeeEvaluationResponse', null, global);
goog.exportSymbol('proto.operations_ecosys.Filter', null, global);
goog.exportSymbol('proto.operations_ecosys.Filter.Comparisons', null, global);
goog.exportSymbol('proto.operations_ecosys.IncidentReport', null, global);
goog.exportSymbol('proto.operations_ecosys.IncidentReport.ReportType', null, global);
goog.exportSymbol('proto.operations_ecosys.IncidentReportContent', null, global);
goog.exportSymbol('proto.operations_ecosys.IncidentReportFilter', null, global);
goog.exportSymbol('proto.operations_ecosys.IncidentReportFilter.Field', null, global);
goog.exportSymbol('proto.operations_ecosys.IncidentReportQuery', null, global);
goog.exportSymbol('proto.operations_ecosys.IncidentReportResponse', null, global);
goog.exportSymbol('proto.operations_ecosys.OrderBy', null, global);
goog.exportSymbol('proto.operations_ecosys.OrderByBroadcast', null, global);
goog.exportSymbol('proto.operations_ecosys.OrderByCameraIot', null, global);
goog.exportSymbol('proto.operations_ecosys.OrderByClient', null, global);
goog.exportSymbol('proto.operations_ecosys.OrderByIncidentReport', null, global);
goog.exportSymbol('proto.operations_ecosys.OrderByQuery', null, global);
goog.exportSymbol('proto.operations_ecosys.OrderByRoster', null, global);
goog.exportSymbol('proto.operations_ecosys.OrderByUser', null, global);
goog.exportSymbol('proto.operations_ecosys.Response', null, global);
goog.exportSymbol('proto.operations_ecosys.Response.Type', null, global);
goog.exportSymbol('proto.operations_ecosys.ResponseNonce', null, global);
goog.exportSymbol('proto.operations_ecosys.Roster', null, global);
goog.exportSymbol('proto.operations_ecosys.Roster.Status', null, global);
goog.exportSymbol('proto.operations_ecosys.RosterAssignement', null, global);
goog.exportSymbol('proto.operations_ecosys.RosterAssignmentResponse', null, global);
goog.exportSymbol('proto.operations_ecosys.RosterFilter', null, global);
goog.exportSymbol('proto.operations_ecosys.RosterFilter.Field', null, global);
goog.exportSymbol('proto.operations_ecosys.RosterQuery', null, global);
goog.exportSymbol('proto.operations_ecosys.RosterResponse', null, global);
goog.exportSymbol('proto.operations_ecosys.User', null, global);
goog.exportSymbol('proto.operations_ecosys.User.UserType', null, global);
goog.exportSymbol('proto.operations_ecosys.UserFilter', null, global);
goog.exportSymbol('proto.operations_ecosys.UserFilter.Field', null, global);
goog.exportSymbol('proto.operations_ecosys.UserQuery', null, global);
goog.exportSymbol('proto.operations_ecosys.UsersResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.User = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.User, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.User.displayName = 'proto.operations_ecosys.User';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.UsersResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.UsersResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.UsersResponse.displayName = 'proto.operations_ecosys.UsersResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.UserFilter = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.UserFilter, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.UserFilter.displayName = 'proto.operations_ecosys.UserFilter';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.UserQuery = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.UserQuery.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.UserQuery, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.UserQuery.displayName = 'proto.operations_ecosys.UserQuery';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.OrderByUser = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.OrderByUser, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.OrderByUser.displayName = 'proto.operations_ecosys.OrderByUser';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.Client = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.Client, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.Client.displayName = 'proto.operations_ecosys.Client';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.ClientResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.ClientResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.ClientResponse.displayName = 'proto.operations_ecosys.ClientResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.ClientFilter = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.ClientFilter, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.ClientFilter.displayName = 'proto.operations_ecosys.ClientFilter';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.ClientQuery = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.ClientQuery.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.ClientQuery, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.ClientQuery.displayName = 'proto.operations_ecosys.ClientQuery';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.OrderByClient = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.OrderByClient, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.OrderByClient.displayName = 'proto.operations_ecosys.OrderByClient';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.ResponseNonce = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.ResponseNonce, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.ResponseNonce.displayName = 'proto.operations_ecosys.ResponseNonce';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.Broadcast = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.Broadcast.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.Broadcast, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.Broadcast.displayName = 'proto.operations_ecosys.Broadcast';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.AIFSBroadcastRecipient = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.AIFSBroadcastRecipient.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.AIFSBroadcastRecipient, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.AIFSBroadcastRecipient.displayName = 'proto.operations_ecosys.AIFSBroadcastRecipient';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.BroadcastRecipient = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.BroadcastRecipient, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.BroadcastRecipient.displayName = 'proto.operations_ecosys.BroadcastRecipient';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.BroadcastResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.BroadcastResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.BroadcastResponse.displayName = 'proto.operations_ecosys.BroadcastResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.BroadcastFilter = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.BroadcastFilter, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.BroadcastFilter.displayName = 'proto.operations_ecosys.BroadcastFilter';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.BroadcastQuery = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.BroadcastQuery.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.BroadcastQuery, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.BroadcastQuery.displayName = 'proto.operations_ecosys.BroadcastQuery';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.OrderByBroadcast = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.OrderByBroadcast, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.OrderByBroadcast.displayName = 'proto.operations_ecosys.OrderByBroadcast';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.Roster = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.Roster.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.Roster, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.Roster.displayName = 'proto.operations_ecosys.Roster';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.AIFSClientRoster = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.AIFSClientRoster, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.AIFSClientRoster.displayName = 'proto.operations_ecosys.AIFSClientRoster';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.RosterAssignement = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.RosterAssignement, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.RosterAssignement.displayName = 'proto.operations_ecosys.RosterAssignement';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.BulkRosters = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.BulkRosters.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.BulkRosters, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.BulkRosters.displayName = 'proto.operations_ecosys.BulkRosters';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.RosterResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.RosterResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.RosterResponse.displayName = 'proto.operations_ecosys.RosterResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.RosterAssignmentResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.RosterAssignmentResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.RosterAssignmentResponse.displayName = 'proto.operations_ecosys.RosterAssignmentResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.RosterFilter = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.RosterFilter, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.RosterFilter.displayName = 'proto.operations_ecosys.RosterFilter';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.RosterQuery = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.RosterQuery.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.RosterQuery, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.RosterQuery.displayName = 'proto.operations_ecosys.RosterQuery';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.OrderByRoster = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.OrderByRoster, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.OrderByRoster.displayName = 'proto.operations_ecosys.OrderByRoster';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.AvailabilityQuery = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.AvailabilityQuery.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.AvailabilityQuery, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.AvailabilityQuery.displayName = 'proto.operations_ecosys.AvailabilityQuery';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.OrderByQuery = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.OrderByQuery, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.OrderByQuery.displayName = 'proto.operations_ecosys.OrderByQuery';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.AvailabilityFilter = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.AvailabilityFilter, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.AvailabilityFilter.displayName = 'proto.operations_ecosys.AvailabilityFilter';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.EmployeeEvaluationResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.EmployeeEvaluationResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.EmployeeEvaluationResponse.displayName = 'proto.operations_ecosys.EmployeeEvaluationResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.EmployeeEvaluation = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.EmployeeEvaluation, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.EmployeeEvaluation.displayName = 'proto.operations_ecosys.EmployeeEvaluation';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.IncidentReport = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.IncidentReport, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.IncidentReport.displayName = 'proto.operations_ecosys.IncidentReport';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.IncidentReportContent = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.IncidentReportContent, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.IncidentReportContent.displayName = 'proto.operations_ecosys.IncidentReportContent';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.IncidentReportResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.IncidentReportResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.IncidentReportResponse.displayName = 'proto.operations_ecosys.IncidentReportResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.IncidentReportFilter = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.IncidentReportFilter, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.IncidentReportFilter.displayName = 'proto.operations_ecosys.IncidentReportFilter';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.IncidentReportQuery = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.IncidentReportQuery.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.IncidentReportQuery, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.IncidentReportQuery.displayName = 'proto.operations_ecosys.IncidentReportQuery';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.OrderByIncidentReport = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.OrderByIncidentReport, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.OrderByIncidentReport.displayName = 'proto.operations_ecosys.OrderByIncidentReport';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.CameraIot = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.CameraIot, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.CameraIot.displayName = 'proto.operations_ecosys.CameraIot';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.Camera = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.Camera, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.Camera.displayName = 'proto.operations_ecosys.Camera';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.CameraIotResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.CameraIotResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.CameraIotResponse.displayName = 'proto.operations_ecosys.CameraIotResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.CameraIotFilter = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.CameraIotFilter, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.CameraIotFilter.displayName = 'proto.operations_ecosys.CameraIotFilter';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.CameraIotQuery = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.operations_ecosys.CameraIotQuery.repeatedFields_, null);
};
goog.inherits(proto.operations_ecosys.CameraIotQuery, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.CameraIotQuery.displayName = 'proto.operations_ecosys.CameraIotQuery';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.OrderByCameraIot = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.OrderByCameraIot, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.OrderByCameraIot.displayName = 'proto.operations_ecosys.OrderByCameraIot';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.Response = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.Response, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.Response.displayName = 'proto.operations_ecosys.Response';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.operations_ecosys.Filter = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.operations_ecosys.Filter, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.operations_ecosys.Filter.displayName = 'proto.operations_ecosys.Filter';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.User.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.User.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.User} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.User.toObject = function(includeInstance, msg) {
  var f, obj = {
    userId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    userType: jspb.Message.getFieldWithDefault(msg, 2, 0),
    name: jspb.Message.getFieldWithDefault(msg, 3, ""),
    email: jspb.Message.getFieldWithDefault(msg, 4, ""),
    phoneNumber: jspb.Message.getFieldWithDefault(msg, 5, ""),
    telegramHandle: jspb.Message.getFieldWithDefault(msg, 6, ""),
    userSecurityImg: jspb.Message.getFieldWithDefault(msg, 7, ""),
    isPartTimer: jspb.Message.getBooleanFieldWithDefault(msg, 8, false),
    teleUserId: jspb.Message.getFieldWithDefault(msg, 9, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.User}
 */
proto.operations_ecosys.User.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.User;
  return proto.operations_ecosys.User.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.User} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.User}
 */
proto.operations_ecosys.User.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserId(value);
      break;
    case 2:
      var value = /** @type {!proto.operations_ecosys.User.UserType} */ (reader.readEnum());
      msg.setUserType(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setPhoneNumber(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setTelegramHandle(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserSecurityImg(value);
      break;
    case 8:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsPartTimer(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTeleUserId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.User.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.User.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.User} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.User.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getUserType();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getPhoneNumber();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getTelegramHandle();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getUserSecurityImg();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getIsPartTimer();
  if (f) {
    writer.writeBool(
      8,
      f
    );
  }
  f = message.getTeleUserId();
  if (f !== 0) {
    writer.writeInt64(
      9,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.User.UserType = {
  ISPECIALIST: 0,
  SECURITY_GUARD: 1,
  CONTROLLER: 2,
  MANAGER: 3
};

/**
 * optional int64 user_id = 1;
 * @return {number}
 */
proto.operations_ecosys.User.prototype.getUserId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.User} returns this
 */
proto.operations_ecosys.User.prototype.setUserId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional UserType user_type = 2;
 * @return {!proto.operations_ecosys.User.UserType}
 */
proto.operations_ecosys.User.prototype.getUserType = function() {
  return /** @type {!proto.operations_ecosys.User.UserType} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.operations_ecosys.User.UserType} value
 * @return {!proto.operations_ecosys.User} returns this
 */
proto.operations_ecosys.User.prototype.setUserType = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional string name = 3;
 * @return {string}
 */
proto.operations_ecosys.User.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.User} returns this
 */
proto.operations_ecosys.User.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string email = 4;
 * @return {string}
 */
proto.operations_ecosys.User.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.User} returns this
 */
proto.operations_ecosys.User.prototype.setEmail = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string phone_number = 5;
 * @return {string}
 */
proto.operations_ecosys.User.prototype.getPhoneNumber = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.User} returns this
 */
proto.operations_ecosys.User.prototype.setPhoneNumber = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string telegram_handle = 6;
 * @return {string}
 */
proto.operations_ecosys.User.prototype.getTelegramHandle = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.User} returns this
 */
proto.operations_ecosys.User.prototype.setTelegramHandle = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string user_security_img = 7;
 * @return {string}
 */
proto.operations_ecosys.User.prototype.getUserSecurityImg = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.User} returns this
 */
proto.operations_ecosys.User.prototype.setUserSecurityImg = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional bool is_part_timer = 8;
 * @return {boolean}
 */
proto.operations_ecosys.User.prototype.getIsPartTimer = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 8, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.User} returns this
 */
proto.operations_ecosys.User.prototype.setIsPartTimer = function(value) {
  return jspb.Message.setProto3BooleanField(this, 8, value);
};


/**
 * optional int64 tele_chat_id = 9;
 * @return {number}
 */
proto.operations_ecosys.User.prototype.getTeleUserId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.User} returns this
 */
proto.operations_ecosys.User.prototype.setTeleUserId = function(value) {
  return jspb.Message.setProto3IntField(this, 9, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.UsersResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.UsersResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.UsersResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.UsersResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    response: (f = msg.getResponse()) && proto.operations_ecosys.Response.toObject(includeInstance, f),
    user: (f = msg.getUser()) && proto.operations_ecosys.User.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.UsersResponse}
 */
proto.operations_ecosys.UsersResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.UsersResponse;
  return proto.operations_ecosys.UsersResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.UsersResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.UsersResponse}
 */
proto.operations_ecosys.UsersResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.Response;
      reader.readMessage(value,proto.operations_ecosys.Response.deserializeBinaryFromReader);
      msg.setResponse(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.User;
      reader.readMessage(value,proto.operations_ecosys.User.deserializeBinaryFromReader);
      msg.setUser(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.UsersResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.UsersResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.UsersResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.UsersResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResponse();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.operations_ecosys.Response.serializeBinaryToWriter
    );
  }
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.User.serializeBinaryToWriter
    );
  }
};


/**
 * optional Response response = 1;
 * @return {?proto.operations_ecosys.Response}
 */
proto.operations_ecosys.UsersResponse.prototype.getResponse = function() {
  return /** @type{?proto.operations_ecosys.Response} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Response, 1));
};


/**
 * @param {?proto.operations_ecosys.Response|undefined} value
 * @return {!proto.operations_ecosys.UsersResponse} returns this
*/
proto.operations_ecosys.UsersResponse.prototype.setResponse = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.UsersResponse} returns this
 */
proto.operations_ecosys.UsersResponse.prototype.clearResponse = function() {
  return this.setResponse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.UsersResponse.prototype.hasResponse = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional User user = 2;
 * @return {?proto.operations_ecosys.User}
 */
proto.operations_ecosys.UsersResponse.prototype.getUser = function() {
  return /** @type{?proto.operations_ecosys.User} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.User, 2));
};


/**
 * @param {?proto.operations_ecosys.User|undefined} value
 * @return {!proto.operations_ecosys.UsersResponse} returns this
*/
proto.operations_ecosys.UsersResponse.prototype.setUser = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.UsersResponse} returns this
 */
proto.operations_ecosys.UsersResponse.prototype.clearUser = function() {
  return this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.UsersResponse.prototype.hasUser = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.UserFilter.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.UserFilter.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.UserFilter} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.UserFilter.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    comparisons: (f = msg.getComparisons()) && proto.operations_ecosys.Filter.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.UserFilter}
 */
proto.operations_ecosys.UserFilter.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.UserFilter;
  return proto.operations_ecosys.UserFilter.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.UserFilter} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.UserFilter}
 */
proto.operations_ecosys.UserFilter.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.UserFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Filter;
      reader.readMessage(value,proto.operations_ecosys.Filter.deserializeBinaryFromReader);
      msg.setComparisons(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.UserFilter.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.UserFilter.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.UserFilter} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.UserFilter.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getComparisons();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Filter.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.UserFilter.Field = {
  USER_ID: 0,
  TYPE: 1,
  NAME: 2,
  EMAIL: 3,
  PHONE_NUMBER: 4,
  TELEGRAM_HANDLE: 5,
  IS_PART_TIMER: 6
};

/**
 * optional Field field = 1;
 * @return {!proto.operations_ecosys.UserFilter.Field}
 */
proto.operations_ecosys.UserFilter.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.UserFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.UserFilter.Field} value
 * @return {!proto.operations_ecosys.UserFilter} returns this
 */
proto.operations_ecosys.UserFilter.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional Filter comparisons = 2;
 * @return {?proto.operations_ecosys.Filter}
 */
proto.operations_ecosys.UserFilter.prototype.getComparisons = function() {
  return /** @type{?proto.operations_ecosys.Filter} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Filter, 2));
};


/**
 * @param {?proto.operations_ecosys.Filter|undefined} value
 * @return {!proto.operations_ecosys.UserFilter} returns this
*/
proto.operations_ecosys.UserFilter.prototype.setComparisons = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.UserFilter} returns this
 */
proto.operations_ecosys.UserFilter.prototype.clearComparisons = function() {
  return this.setComparisons(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.UserFilter.prototype.hasComparisons = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.UserQuery.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.UserQuery.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.UserQuery.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.UserQuery} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.UserQuery.toObject = function(includeInstance, msg) {
  var f, obj = {
    filtersList: jspb.Message.toObjectList(msg.getFiltersList(),
    proto.operations_ecosys.UserFilter.toObject, includeInstance),
    limit: jspb.Message.getFieldWithDefault(msg, 2, 0),
    skip: jspb.Message.getFieldWithDefault(msg, 3, 0),
    orderBy: (f = msg.getOrderBy()) && proto.operations_ecosys.OrderByUser.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.UserQuery}
 */
proto.operations_ecosys.UserQuery.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.UserQuery;
  return proto.operations_ecosys.UserQuery.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.UserQuery} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.UserQuery}
 */
proto.operations_ecosys.UserQuery.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.UserFilter;
      reader.readMessage(value,proto.operations_ecosys.UserFilter.deserializeBinaryFromReader);
      msg.addFilters(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLimit(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSkip(value);
      break;
    case 4:
      var value = new proto.operations_ecosys.OrderByUser;
      reader.readMessage(value,proto.operations_ecosys.OrderByUser.deserializeBinaryFromReader);
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.UserQuery.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.UserQuery.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.UserQuery} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.UserQuery.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFiltersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.operations_ecosys.UserFilter.serializeBinaryToWriter
    );
  }
  f = message.getLimit();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getSkip();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getOrderBy();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.operations_ecosys.OrderByUser.serializeBinaryToWriter
    );
  }
};


/**
 * repeated UserFilter filters = 1;
 * @return {!Array<!proto.operations_ecosys.UserFilter>}
 */
proto.operations_ecosys.UserQuery.prototype.getFiltersList = function() {
  return /** @type{!Array<!proto.operations_ecosys.UserFilter>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.UserFilter, 1));
};


/**
 * @param {!Array<!proto.operations_ecosys.UserFilter>} value
 * @return {!proto.operations_ecosys.UserQuery} returns this
*/
proto.operations_ecosys.UserQuery.prototype.setFiltersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.operations_ecosys.UserFilter=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.UserFilter}
 */
proto.operations_ecosys.UserQuery.prototype.addFilters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.operations_ecosys.UserFilter, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.UserQuery} returns this
 */
proto.operations_ecosys.UserQuery.prototype.clearFiltersList = function() {
  return this.setFiltersList([]);
};


/**
 * optional int64 limit = 2;
 * @return {number}
 */
proto.operations_ecosys.UserQuery.prototype.getLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.UserQuery} returns this
 */
proto.operations_ecosys.UserQuery.prototype.setLimit = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 skip = 3;
 * @return {number}
 */
proto.operations_ecosys.UserQuery.prototype.getSkip = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.UserQuery} returns this
 */
proto.operations_ecosys.UserQuery.prototype.setSkip = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional OrderByUser order_by = 4;
 * @return {?proto.operations_ecosys.OrderByUser}
 */
proto.operations_ecosys.UserQuery.prototype.getOrderBy = function() {
  return /** @type{?proto.operations_ecosys.OrderByUser} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.OrderByUser, 4));
};


/**
 * @param {?proto.operations_ecosys.OrderByUser|undefined} value
 * @return {!proto.operations_ecosys.UserQuery} returns this
*/
proto.operations_ecosys.UserQuery.prototype.setOrderBy = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.UserQuery} returns this
 */
proto.operations_ecosys.UserQuery.prototype.clearOrderBy = function() {
  return this.setOrderBy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.UserQuery.prototype.hasOrderBy = function() {
  return jspb.Message.getField(this, 4) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.OrderByUser.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.OrderByUser.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.OrderByUser} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByUser.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    orderBy: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.OrderByUser}
 */
proto.operations_ecosys.OrderByUser.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.OrderByUser;
  return proto.operations_ecosys.OrderByUser.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.OrderByUser} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.OrderByUser}
 */
proto.operations_ecosys.OrderByUser.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.UserFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = /** @type {!proto.operations_ecosys.OrderBy} */ (reader.readEnum());
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.OrderByUser.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.OrderByUser.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.OrderByUser} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByUser.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getOrderBy();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * optional UserFilter.Field field = 1;
 * @return {!proto.operations_ecosys.UserFilter.Field}
 */
proto.operations_ecosys.OrderByUser.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.UserFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.UserFilter.Field} value
 * @return {!proto.operations_ecosys.OrderByUser} returns this
 */
proto.operations_ecosys.OrderByUser.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional OrderBy order_by = 2;
 * @return {!proto.operations_ecosys.OrderBy}
 */
proto.operations_ecosys.OrderByUser.prototype.getOrderBy = function() {
  return /** @type {!proto.operations_ecosys.OrderBy} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.operations_ecosys.OrderBy} value
 * @return {!proto.operations_ecosys.OrderByUser} returns this
 */
proto.operations_ecosys.OrderByUser.prototype.setOrderBy = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.Client.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.Client.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.Client} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Client.toObject = function(includeInstance, msg) {
  var f, obj = {
    clientId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    abbreviation: jspb.Message.getFieldWithDefault(msg, 3, ""),
    email: jspb.Message.getFieldWithDefault(msg, 4, ""),
    address: jspb.Message.getFieldWithDefault(msg, 5, ""),
    postalCode: jspb.Message.getFieldWithDefault(msg, 6, 0),
    phoneNumber: jspb.Message.getFieldWithDefault(msg, 7, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.Client}
 */
proto.operations_ecosys.Client.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.Client;
  return proto.operations_ecosys.Client.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.Client} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.Client}
 */
proto.operations_ecosys.Client.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setClientId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setAbbreviation(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPostalCode(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setPhoneNumber(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.Client.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.Client.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.Client} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Client.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClientId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getAbbreviation();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getPostalCode();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getPhoneNumber();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
};


/**
 * optional int64 client_id = 1;
 * @return {number}
 */
proto.operations_ecosys.Client.prototype.getClientId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.Client} returns this
 */
proto.operations_ecosys.Client.prototype.setClientId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.operations_ecosys.Client.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Client} returns this
 */
proto.operations_ecosys.Client.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string abbreviation = 3;
 * @return {string}
 */
proto.operations_ecosys.Client.prototype.getAbbreviation = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Client} returns this
 */
proto.operations_ecosys.Client.prototype.setAbbreviation = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string email = 4;
 * @return {string}
 */
proto.operations_ecosys.Client.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Client} returns this
 */
proto.operations_ecosys.Client.prototype.setEmail = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string address = 5;
 * @return {string}
 */
proto.operations_ecosys.Client.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Client} returns this
 */
proto.operations_ecosys.Client.prototype.setAddress = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional int64 postal_code = 6;
 * @return {number}
 */
proto.operations_ecosys.Client.prototype.getPostalCode = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.Client} returns this
 */
proto.operations_ecosys.Client.prototype.setPostalCode = function(value) {
  return jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional string phone_number = 7;
 * @return {string}
 */
proto.operations_ecosys.Client.prototype.getPhoneNumber = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Client} returns this
 */
proto.operations_ecosys.Client.prototype.setPhoneNumber = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.ClientResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.ClientResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.ClientResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.ClientResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    response: (f = msg.getResponse()) && proto.operations_ecosys.Response.toObject(includeInstance, f),
    client: (f = msg.getClient()) && proto.operations_ecosys.Client.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.ClientResponse}
 */
proto.operations_ecosys.ClientResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.ClientResponse;
  return proto.operations_ecosys.ClientResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.ClientResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.ClientResponse}
 */
proto.operations_ecosys.ClientResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.Response;
      reader.readMessage(value,proto.operations_ecosys.Response.deserializeBinaryFromReader);
      msg.setResponse(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Client;
      reader.readMessage(value,proto.operations_ecosys.Client.deserializeBinaryFromReader);
      msg.setClient(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.ClientResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.ClientResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.ClientResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.ClientResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResponse();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.operations_ecosys.Response.serializeBinaryToWriter
    );
  }
  f = message.getClient();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Client.serializeBinaryToWriter
    );
  }
};


/**
 * optional Response response = 1;
 * @return {?proto.operations_ecosys.Response}
 */
proto.operations_ecosys.ClientResponse.prototype.getResponse = function() {
  return /** @type{?proto.operations_ecosys.Response} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Response, 1));
};


/**
 * @param {?proto.operations_ecosys.Response|undefined} value
 * @return {!proto.operations_ecosys.ClientResponse} returns this
*/
proto.operations_ecosys.ClientResponse.prototype.setResponse = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.ClientResponse} returns this
 */
proto.operations_ecosys.ClientResponse.prototype.clearResponse = function() {
  return this.setResponse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.ClientResponse.prototype.hasResponse = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Client client = 2;
 * @return {?proto.operations_ecosys.Client}
 */
proto.operations_ecosys.ClientResponse.prototype.getClient = function() {
  return /** @type{?proto.operations_ecosys.Client} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Client, 2));
};


/**
 * @param {?proto.operations_ecosys.Client|undefined} value
 * @return {!proto.operations_ecosys.ClientResponse} returns this
*/
proto.operations_ecosys.ClientResponse.prototype.setClient = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.ClientResponse} returns this
 */
proto.operations_ecosys.ClientResponse.prototype.clearClient = function() {
  return this.setClient(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.ClientResponse.prototype.hasClient = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.ClientFilter.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.ClientFilter.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.ClientFilter} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.ClientFilter.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    comparisons: (f = msg.getComparisons()) && proto.operations_ecosys.Filter.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.ClientFilter}
 */
proto.operations_ecosys.ClientFilter.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.ClientFilter;
  return proto.operations_ecosys.ClientFilter.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.ClientFilter} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.ClientFilter}
 */
proto.operations_ecosys.ClientFilter.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.ClientFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Filter;
      reader.readMessage(value,proto.operations_ecosys.Filter.deserializeBinaryFromReader);
      msg.setComparisons(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.ClientFilter.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.ClientFilter.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.ClientFilter} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.ClientFilter.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getComparisons();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Filter.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.ClientFilter.Field = {
  CLIENT_ID: 0
};

/**
 * optional Field field = 1;
 * @return {!proto.operations_ecosys.ClientFilter.Field}
 */
proto.operations_ecosys.ClientFilter.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.ClientFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.ClientFilter.Field} value
 * @return {!proto.operations_ecosys.ClientFilter} returns this
 */
proto.operations_ecosys.ClientFilter.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional Filter comparisons = 2;
 * @return {?proto.operations_ecosys.Filter}
 */
proto.operations_ecosys.ClientFilter.prototype.getComparisons = function() {
  return /** @type{?proto.operations_ecosys.Filter} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Filter, 2));
};


/**
 * @param {?proto.operations_ecosys.Filter|undefined} value
 * @return {!proto.operations_ecosys.ClientFilter} returns this
*/
proto.operations_ecosys.ClientFilter.prototype.setComparisons = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.ClientFilter} returns this
 */
proto.operations_ecosys.ClientFilter.prototype.clearComparisons = function() {
  return this.setComparisons(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.ClientFilter.prototype.hasComparisons = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.ClientQuery.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.ClientQuery.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.ClientQuery.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.ClientQuery} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.ClientQuery.toObject = function(includeInstance, msg) {
  var f, obj = {
    filtersList: jspb.Message.toObjectList(msg.getFiltersList(),
    proto.operations_ecosys.ClientFilter.toObject, includeInstance),
    limit: jspb.Message.getFieldWithDefault(msg, 2, 0),
    skip: jspb.Message.getFieldWithDefault(msg, 3, 0),
    orderBy: (f = msg.getOrderBy()) && proto.operations_ecosys.OrderByClient.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.ClientQuery}
 */
proto.operations_ecosys.ClientQuery.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.ClientQuery;
  return proto.operations_ecosys.ClientQuery.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.ClientQuery} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.ClientQuery}
 */
proto.operations_ecosys.ClientQuery.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.ClientFilter;
      reader.readMessage(value,proto.operations_ecosys.ClientFilter.deserializeBinaryFromReader);
      msg.addFilters(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLimit(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSkip(value);
      break;
    case 4:
      var value = new proto.operations_ecosys.OrderByClient;
      reader.readMessage(value,proto.operations_ecosys.OrderByClient.deserializeBinaryFromReader);
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.ClientQuery.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.ClientQuery.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.ClientQuery} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.ClientQuery.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFiltersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.operations_ecosys.ClientFilter.serializeBinaryToWriter
    );
  }
  f = message.getLimit();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getSkip();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getOrderBy();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.operations_ecosys.OrderByClient.serializeBinaryToWriter
    );
  }
};


/**
 * repeated ClientFilter filters = 1;
 * @return {!Array<!proto.operations_ecosys.ClientFilter>}
 */
proto.operations_ecosys.ClientQuery.prototype.getFiltersList = function() {
  return /** @type{!Array<!proto.operations_ecosys.ClientFilter>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.ClientFilter, 1));
};


/**
 * @param {!Array<!proto.operations_ecosys.ClientFilter>} value
 * @return {!proto.operations_ecosys.ClientQuery} returns this
*/
proto.operations_ecosys.ClientQuery.prototype.setFiltersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.operations_ecosys.ClientFilter=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.ClientFilter}
 */
proto.operations_ecosys.ClientQuery.prototype.addFilters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.operations_ecosys.ClientFilter, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.ClientQuery} returns this
 */
proto.operations_ecosys.ClientQuery.prototype.clearFiltersList = function() {
  return this.setFiltersList([]);
};


/**
 * optional int64 limit = 2;
 * @return {number}
 */
proto.operations_ecosys.ClientQuery.prototype.getLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.ClientQuery} returns this
 */
proto.operations_ecosys.ClientQuery.prototype.setLimit = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 skip = 3;
 * @return {number}
 */
proto.operations_ecosys.ClientQuery.prototype.getSkip = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.ClientQuery} returns this
 */
proto.operations_ecosys.ClientQuery.prototype.setSkip = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional OrderByClient order_by = 4;
 * @return {?proto.operations_ecosys.OrderByClient}
 */
proto.operations_ecosys.ClientQuery.prototype.getOrderBy = function() {
  return /** @type{?proto.operations_ecosys.OrderByClient} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.OrderByClient, 4));
};


/**
 * @param {?proto.operations_ecosys.OrderByClient|undefined} value
 * @return {!proto.operations_ecosys.ClientQuery} returns this
*/
proto.operations_ecosys.ClientQuery.prototype.setOrderBy = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.ClientQuery} returns this
 */
proto.operations_ecosys.ClientQuery.prototype.clearOrderBy = function() {
  return this.setOrderBy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.ClientQuery.prototype.hasOrderBy = function() {
  return jspb.Message.getField(this, 4) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.OrderByClient.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.OrderByClient.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.OrderByClient} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByClient.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    orderBy: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.OrderByClient}
 */
proto.operations_ecosys.OrderByClient.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.OrderByClient;
  return proto.operations_ecosys.OrderByClient.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.OrderByClient} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.OrderByClient}
 */
proto.operations_ecosys.OrderByClient.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.ClientFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = /** @type {!proto.operations_ecosys.OrderBy} */ (reader.readEnum());
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.OrderByClient.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.OrderByClient.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.OrderByClient} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByClient.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getOrderBy();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * optional ClientFilter.Field field = 1;
 * @return {!proto.operations_ecosys.ClientFilter.Field}
 */
proto.operations_ecosys.OrderByClient.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.ClientFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.ClientFilter.Field} value
 * @return {!proto.operations_ecosys.OrderByClient} returns this
 */
proto.operations_ecosys.OrderByClient.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional OrderBy order_by = 2;
 * @return {!proto.operations_ecosys.OrderBy}
 */
proto.operations_ecosys.OrderByClient.prototype.getOrderBy = function() {
  return /** @type {!proto.operations_ecosys.OrderBy} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.operations_ecosys.OrderBy} value
 * @return {!proto.operations_ecosys.OrderByClient} returns this
 */
proto.operations_ecosys.OrderByClient.prototype.setOrderBy = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.ResponseNonce.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.ResponseNonce.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.ResponseNonce} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.ResponseNonce.toObject = function(includeInstance, msg) {
  var f, obj = {
    response: (f = msg.getResponse()) && proto.operations_ecosys.Response.toObject(includeInstance, f),
    nonce: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.ResponseNonce}
 */
proto.operations_ecosys.ResponseNonce.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.ResponseNonce;
  return proto.operations_ecosys.ResponseNonce.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.ResponseNonce} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.ResponseNonce}
 */
proto.operations_ecosys.ResponseNonce.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.Response;
      reader.readMessage(value,proto.operations_ecosys.Response.deserializeBinaryFromReader);
      msg.setResponse(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setNonce(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.ResponseNonce.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.ResponseNonce.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.ResponseNonce} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.ResponseNonce.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResponse();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.operations_ecosys.Response.serializeBinaryToWriter
    );
  }
  f = message.getNonce();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional Response response = 1;
 * @return {?proto.operations_ecosys.Response}
 */
proto.operations_ecosys.ResponseNonce.prototype.getResponse = function() {
  return /** @type{?proto.operations_ecosys.Response} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Response, 1));
};


/**
 * @param {?proto.operations_ecosys.Response|undefined} value
 * @return {!proto.operations_ecosys.ResponseNonce} returns this
*/
proto.operations_ecosys.ResponseNonce.prototype.setResponse = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.ResponseNonce} returns this
 */
proto.operations_ecosys.ResponseNonce.prototype.clearResponse = function() {
  return this.setResponse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.ResponseNonce.prototype.hasResponse = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string nonce = 2;
 * @return {string}
 */
proto.operations_ecosys.ResponseNonce.prototype.getNonce = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.ResponseNonce} returns this
 */
proto.operations_ecosys.ResponseNonce.prototype.setNonce = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.Broadcast.repeatedFields_ = [7];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.Broadcast.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.Broadcast.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.Broadcast} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Broadcast.toObject = function(includeInstance, msg) {
  var f, obj = {
    broadcastId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    type: jspb.Message.getFieldWithDefault(msg, 2, 0),
    content: jspb.Message.getFieldWithDefault(msg, 3, ""),
    creationDate: (f = msg.getCreationDate()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    deadline: (f = msg.getDeadline()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    creator: (f = msg.getCreator()) && proto.operations_ecosys.User.toObject(includeInstance, f),
    recipientsList: jspb.Message.toObjectList(msg.getRecipientsList(),
    proto.operations_ecosys.AIFSBroadcastRecipient.toObject, includeInstance),
    urgency: jspb.Message.getFieldWithDefault(msg, 8, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.Broadcast}
 */
proto.operations_ecosys.Broadcast.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.Broadcast;
  return proto.operations_ecosys.Broadcast.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.Broadcast} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.Broadcast}
 */
proto.operations_ecosys.Broadcast.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setBroadcastId(value);
      break;
    case 2:
      var value = /** @type {!proto.operations_ecosys.Broadcast.BroadcastType} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setContent(value);
      break;
    case 4:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreationDate(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setDeadline(value);
      break;
    case 6:
      var value = new proto.operations_ecosys.User;
      reader.readMessage(value,proto.operations_ecosys.User.deserializeBinaryFromReader);
      msg.setCreator(value);
      break;
    case 7:
      var value = new proto.operations_ecosys.AIFSBroadcastRecipient;
      reader.readMessage(value,proto.operations_ecosys.AIFSBroadcastRecipient.deserializeBinaryFromReader);
      msg.addRecipients(value);
      break;
    case 8:
      var value = /** @type {!proto.operations_ecosys.Broadcast.UrgencyType} */ (reader.readEnum());
      msg.setUrgency(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.Broadcast.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.Broadcast.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.Broadcast} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Broadcast.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getBroadcastId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getContent();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getCreationDate();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getDeadline();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getCreator();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.operations_ecosys.User.serializeBinaryToWriter
    );
  }
  f = message.getRecipientsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      7,
      f,
      proto.operations_ecosys.AIFSBroadcastRecipient.serializeBinaryToWriter
    );
  }
  f = message.getUrgency();
  if (f !== 0.0) {
    writer.writeEnum(
      8,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.Broadcast.BroadcastType = {
  ANNOUNCEMENT: 0,
  ASSIGNMENT: 1
};

/**
 * @enum {number}
 */
proto.operations_ecosys.Broadcast.UrgencyType = {
  LOW: 0,
  MEDIUM: 1,
  HIGH: 2
};

/**
 * optional int64 broadcast_id = 1;
 * @return {number}
 */
proto.operations_ecosys.Broadcast.prototype.getBroadcastId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.Broadcast} returns this
 */
proto.operations_ecosys.Broadcast.prototype.setBroadcastId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional BroadcastType type = 2;
 * @return {!proto.operations_ecosys.Broadcast.BroadcastType}
 */
proto.operations_ecosys.Broadcast.prototype.getType = function() {
  return /** @type {!proto.operations_ecosys.Broadcast.BroadcastType} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.operations_ecosys.Broadcast.BroadcastType} value
 * @return {!proto.operations_ecosys.Broadcast} returns this
 */
proto.operations_ecosys.Broadcast.prototype.setType = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional string content = 3;
 * @return {string}
 */
proto.operations_ecosys.Broadcast.prototype.getContent = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Broadcast} returns this
 */
proto.operations_ecosys.Broadcast.prototype.setContent = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional google.protobuf.Timestamp creation_date = 4;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.operations_ecosys.Broadcast.prototype.getCreationDate = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 4));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.operations_ecosys.Broadcast} returns this
*/
proto.operations_ecosys.Broadcast.prototype.setCreationDate = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.Broadcast} returns this
 */
proto.operations_ecosys.Broadcast.prototype.clearCreationDate = function() {
  return this.setCreationDate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.Broadcast.prototype.hasCreationDate = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional google.protobuf.Timestamp deadline = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.operations_ecosys.Broadcast.prototype.getDeadline = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.operations_ecosys.Broadcast} returns this
*/
proto.operations_ecosys.Broadcast.prototype.setDeadline = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.Broadcast} returns this
 */
proto.operations_ecosys.Broadcast.prototype.clearDeadline = function() {
  return this.setDeadline(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.Broadcast.prototype.hasDeadline = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional User creator = 6;
 * @return {?proto.operations_ecosys.User}
 */
proto.operations_ecosys.Broadcast.prototype.getCreator = function() {
  return /** @type{?proto.operations_ecosys.User} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.User, 6));
};


/**
 * @param {?proto.operations_ecosys.User|undefined} value
 * @return {!proto.operations_ecosys.Broadcast} returns this
*/
proto.operations_ecosys.Broadcast.prototype.setCreator = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.Broadcast} returns this
 */
proto.operations_ecosys.Broadcast.prototype.clearCreator = function() {
  return this.setCreator(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.Broadcast.prototype.hasCreator = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * repeated AIFSBroadcastRecipient recipients = 7;
 * @return {!Array<!proto.operations_ecosys.AIFSBroadcastRecipient>}
 */
proto.operations_ecosys.Broadcast.prototype.getRecipientsList = function() {
  return /** @type{!Array<!proto.operations_ecosys.AIFSBroadcastRecipient>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.AIFSBroadcastRecipient, 7));
};


/**
 * @param {!Array<!proto.operations_ecosys.AIFSBroadcastRecipient>} value
 * @return {!proto.operations_ecosys.Broadcast} returns this
*/
proto.operations_ecosys.Broadcast.prototype.setRecipientsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 7, value);
};


/**
 * @param {!proto.operations_ecosys.AIFSBroadcastRecipient=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.AIFSBroadcastRecipient}
 */
proto.operations_ecosys.Broadcast.prototype.addRecipients = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 7, opt_value, proto.operations_ecosys.AIFSBroadcastRecipient, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.Broadcast} returns this
 */
proto.operations_ecosys.Broadcast.prototype.clearRecipientsList = function() {
  return this.setRecipientsList([]);
};


/**
 * optional UrgencyType urgency = 8;
 * @return {!proto.operations_ecosys.Broadcast.UrgencyType}
 */
proto.operations_ecosys.Broadcast.prototype.getUrgency = function() {
  return /** @type {!proto.operations_ecosys.Broadcast.UrgencyType} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/**
 * @param {!proto.operations_ecosys.Broadcast.UrgencyType} value
 * @return {!proto.operations_ecosys.Broadcast} returns this
 */
proto.operations_ecosys.Broadcast.prototype.setUrgency = function(value) {
  return jspb.Message.setProto3EnumField(this, 8, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.AIFSBroadcastRecipient.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.AIFSBroadcastRecipient.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.AIFSBroadcastRecipient.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.AIFSBroadcastRecipient} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.AIFSBroadcastRecipient.toObject = function(includeInstance, msg) {
  var f, obj = {
    recipientList: jspb.Message.toObjectList(msg.getRecipientList(),
    proto.operations_ecosys.BroadcastRecipient.toObject, includeInstance),
    aifsId: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.AIFSBroadcastRecipient}
 */
proto.operations_ecosys.AIFSBroadcastRecipient.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.AIFSBroadcastRecipient;
  return proto.operations_ecosys.AIFSBroadcastRecipient.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.AIFSBroadcastRecipient} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.AIFSBroadcastRecipient}
 */
proto.operations_ecosys.AIFSBroadcastRecipient.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.BroadcastRecipient;
      reader.readMessage(value,proto.operations_ecosys.BroadcastRecipient.deserializeBinaryFromReader);
      msg.addRecipient(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAifsId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.AIFSBroadcastRecipient.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.AIFSBroadcastRecipient.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.AIFSBroadcastRecipient} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.AIFSBroadcastRecipient.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRecipientList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.operations_ecosys.BroadcastRecipient.serializeBinaryToWriter
    );
  }
  f = message.getAifsId();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * repeated BroadcastRecipient recipient = 1;
 * @return {!Array<!proto.operations_ecosys.BroadcastRecipient>}
 */
proto.operations_ecosys.AIFSBroadcastRecipient.prototype.getRecipientList = function() {
  return /** @type{!Array<!proto.operations_ecosys.BroadcastRecipient>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.BroadcastRecipient, 1));
};


/**
 * @param {!Array<!proto.operations_ecosys.BroadcastRecipient>} value
 * @return {!proto.operations_ecosys.AIFSBroadcastRecipient} returns this
*/
proto.operations_ecosys.AIFSBroadcastRecipient.prototype.setRecipientList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.operations_ecosys.BroadcastRecipient=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.BroadcastRecipient}
 */
proto.operations_ecosys.AIFSBroadcastRecipient.prototype.addRecipient = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.operations_ecosys.BroadcastRecipient, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.AIFSBroadcastRecipient} returns this
 */
proto.operations_ecosys.AIFSBroadcastRecipient.prototype.clearRecipientList = function() {
  return this.setRecipientList([]);
};


/**
 * optional int64 aifs_id = 2;
 * @return {number}
 */
proto.operations_ecosys.AIFSBroadcastRecipient.prototype.getAifsId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.AIFSBroadcastRecipient} returns this
 */
proto.operations_ecosys.AIFSBroadcastRecipient.prototype.setAifsId = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.BroadcastRecipient.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.BroadcastRecipient.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.BroadcastRecipient} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.BroadcastRecipient.toObject = function(includeInstance, msg) {
  var f, obj = {
    broadcastRecipientsId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    recipient: (f = msg.getRecipient()) && proto.operations_ecosys.User.toObject(includeInstance, f),
    acknowledged: jspb.Message.getBooleanFieldWithDefault(msg, 3, false),
    rejected: jspb.Message.getBooleanFieldWithDefault(msg, 4, false),
    lastReplied: (f = msg.getLastReplied()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    aifsId: jspb.Message.getFieldWithDefault(msg, 6, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.BroadcastRecipient}
 */
proto.operations_ecosys.BroadcastRecipient.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.BroadcastRecipient;
  return proto.operations_ecosys.BroadcastRecipient.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.BroadcastRecipient} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.BroadcastRecipient}
 */
proto.operations_ecosys.BroadcastRecipient.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setBroadcastRecipientsId(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.User;
      reader.readMessage(value,proto.operations_ecosys.User.deserializeBinaryFromReader);
      msg.setRecipient(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setAcknowledged(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setRejected(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLastReplied(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAifsId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.BroadcastRecipient.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.BroadcastRecipient.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.BroadcastRecipient} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.BroadcastRecipient.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getBroadcastRecipientsId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getRecipient();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.User.serializeBinaryToWriter
    );
  }
  f = message.getAcknowledged();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getRejected();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getLastReplied();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getAifsId();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
};


/**
 * optional int64 broadcast_recipients_id = 1;
 * @return {number}
 */
proto.operations_ecosys.BroadcastRecipient.prototype.getBroadcastRecipientsId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.BroadcastRecipient} returns this
 */
proto.operations_ecosys.BroadcastRecipient.prototype.setBroadcastRecipientsId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional User recipient = 2;
 * @return {?proto.operations_ecosys.User}
 */
proto.operations_ecosys.BroadcastRecipient.prototype.getRecipient = function() {
  return /** @type{?proto.operations_ecosys.User} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.User, 2));
};


/**
 * @param {?proto.operations_ecosys.User|undefined} value
 * @return {!proto.operations_ecosys.BroadcastRecipient} returns this
*/
proto.operations_ecosys.BroadcastRecipient.prototype.setRecipient = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.BroadcastRecipient} returns this
 */
proto.operations_ecosys.BroadcastRecipient.prototype.clearRecipient = function() {
  return this.setRecipient(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.BroadcastRecipient.prototype.hasRecipient = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional bool acknowledged = 3;
 * @return {boolean}
 */
proto.operations_ecosys.BroadcastRecipient.prototype.getAcknowledged = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.BroadcastRecipient} returns this
 */
proto.operations_ecosys.BroadcastRecipient.prototype.setAcknowledged = function(value) {
  return jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * optional bool rejected = 4;
 * @return {boolean}
 */
proto.operations_ecosys.BroadcastRecipient.prototype.getRejected = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 4, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.BroadcastRecipient} returns this
 */
proto.operations_ecosys.BroadcastRecipient.prototype.setRejected = function(value) {
  return jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional google.protobuf.Timestamp last_replied = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.operations_ecosys.BroadcastRecipient.prototype.getLastReplied = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.operations_ecosys.BroadcastRecipient} returns this
*/
proto.operations_ecosys.BroadcastRecipient.prototype.setLastReplied = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.BroadcastRecipient} returns this
 */
proto.operations_ecosys.BroadcastRecipient.prototype.clearLastReplied = function() {
  return this.setLastReplied(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.BroadcastRecipient.prototype.hasLastReplied = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional int64 aifs_id = 6;
 * @return {number}
 */
proto.operations_ecosys.BroadcastRecipient.prototype.getAifsId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.BroadcastRecipient} returns this
 */
proto.operations_ecosys.BroadcastRecipient.prototype.setAifsId = function(value) {
  return jspb.Message.setProto3IntField(this, 6, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.BroadcastResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.BroadcastResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.BroadcastResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.BroadcastResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    response: (f = msg.getResponse()) && proto.operations_ecosys.Response.toObject(includeInstance, f),
    broadcast: (f = msg.getBroadcast()) && proto.operations_ecosys.Broadcast.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.BroadcastResponse}
 */
proto.operations_ecosys.BroadcastResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.BroadcastResponse;
  return proto.operations_ecosys.BroadcastResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.BroadcastResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.BroadcastResponse}
 */
proto.operations_ecosys.BroadcastResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.Response;
      reader.readMessage(value,proto.operations_ecosys.Response.deserializeBinaryFromReader);
      msg.setResponse(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Broadcast;
      reader.readMessage(value,proto.operations_ecosys.Broadcast.deserializeBinaryFromReader);
      msg.setBroadcast(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.BroadcastResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.BroadcastResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.BroadcastResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.BroadcastResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResponse();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.operations_ecosys.Response.serializeBinaryToWriter
    );
  }
  f = message.getBroadcast();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Broadcast.serializeBinaryToWriter
    );
  }
};


/**
 * optional Response response = 1;
 * @return {?proto.operations_ecosys.Response}
 */
proto.operations_ecosys.BroadcastResponse.prototype.getResponse = function() {
  return /** @type{?proto.operations_ecosys.Response} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Response, 1));
};


/**
 * @param {?proto.operations_ecosys.Response|undefined} value
 * @return {!proto.operations_ecosys.BroadcastResponse} returns this
*/
proto.operations_ecosys.BroadcastResponse.prototype.setResponse = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.BroadcastResponse} returns this
 */
proto.operations_ecosys.BroadcastResponse.prototype.clearResponse = function() {
  return this.setResponse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.BroadcastResponse.prototype.hasResponse = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Broadcast broadcast = 2;
 * @return {?proto.operations_ecosys.Broadcast}
 */
proto.operations_ecosys.BroadcastResponse.prototype.getBroadcast = function() {
  return /** @type{?proto.operations_ecosys.Broadcast} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Broadcast, 2));
};


/**
 * @param {?proto.operations_ecosys.Broadcast|undefined} value
 * @return {!proto.operations_ecosys.BroadcastResponse} returns this
*/
proto.operations_ecosys.BroadcastResponse.prototype.setBroadcast = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.BroadcastResponse} returns this
 */
proto.operations_ecosys.BroadcastResponse.prototype.clearBroadcast = function() {
  return this.setBroadcast(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.BroadcastResponse.prototype.hasBroadcast = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.BroadcastFilter.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.BroadcastFilter.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.BroadcastFilter} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.BroadcastFilter.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    comparisons: (f = msg.getComparisons()) && proto.operations_ecosys.Filter.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.BroadcastFilter}
 */
proto.operations_ecosys.BroadcastFilter.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.BroadcastFilter;
  return proto.operations_ecosys.BroadcastFilter.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.BroadcastFilter} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.BroadcastFilter}
 */
proto.operations_ecosys.BroadcastFilter.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.BroadcastFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Filter;
      reader.readMessage(value,proto.operations_ecosys.Filter.deserializeBinaryFromReader);
      msg.setComparisons(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.BroadcastFilter.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.BroadcastFilter.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.BroadcastFilter} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.BroadcastFilter.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getComparisons();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Filter.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.BroadcastFilter.Field = {
  BROADCAST_ID: 0,
  TYPE: 1,
  CONTENT: 2,
  CREATION_DATE: 3,
  DEADLINE: 4,
  CREATOR_ID: 5,
  RECEIPEIENT_ID: 6,
  NUM_RECEIPIENTS: 7,
  URGENCY: 8,
  AIFS_ID: 9,
  BROADCAST_RECIPIENT_TABLE_ID: 10
};

/**
 * optional Field field = 1;
 * @return {!proto.operations_ecosys.BroadcastFilter.Field}
 */
proto.operations_ecosys.BroadcastFilter.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.BroadcastFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.BroadcastFilter.Field} value
 * @return {!proto.operations_ecosys.BroadcastFilter} returns this
 */
proto.operations_ecosys.BroadcastFilter.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional Filter comparisons = 2;
 * @return {?proto.operations_ecosys.Filter}
 */
proto.operations_ecosys.BroadcastFilter.prototype.getComparisons = function() {
  return /** @type{?proto.operations_ecosys.Filter} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Filter, 2));
};


/**
 * @param {?proto.operations_ecosys.Filter|undefined} value
 * @return {!proto.operations_ecosys.BroadcastFilter} returns this
*/
proto.operations_ecosys.BroadcastFilter.prototype.setComparisons = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.BroadcastFilter} returns this
 */
proto.operations_ecosys.BroadcastFilter.prototype.clearComparisons = function() {
  return this.setComparisons(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.BroadcastFilter.prototype.hasComparisons = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.BroadcastQuery.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.BroadcastQuery.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.BroadcastQuery.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.BroadcastQuery} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.BroadcastQuery.toObject = function(includeInstance, msg) {
  var f, obj = {
    filtersList: jspb.Message.toObjectList(msg.getFiltersList(),
    proto.operations_ecosys.BroadcastFilter.toObject, includeInstance),
    limit: jspb.Message.getFieldWithDefault(msg, 2, 0),
    skip: jspb.Message.getFieldWithDefault(msg, 3, 0),
    orderBy: (f = msg.getOrderBy()) && proto.operations_ecosys.OrderByBroadcast.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.BroadcastQuery}
 */
proto.operations_ecosys.BroadcastQuery.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.BroadcastQuery;
  return proto.operations_ecosys.BroadcastQuery.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.BroadcastQuery} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.BroadcastQuery}
 */
proto.operations_ecosys.BroadcastQuery.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.BroadcastFilter;
      reader.readMessage(value,proto.operations_ecosys.BroadcastFilter.deserializeBinaryFromReader);
      msg.addFilters(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLimit(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSkip(value);
      break;
    case 4:
      var value = new proto.operations_ecosys.OrderByBroadcast;
      reader.readMessage(value,proto.operations_ecosys.OrderByBroadcast.deserializeBinaryFromReader);
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.BroadcastQuery.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.BroadcastQuery.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.BroadcastQuery} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.BroadcastQuery.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFiltersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.operations_ecosys.BroadcastFilter.serializeBinaryToWriter
    );
  }
  f = message.getLimit();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getSkip();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getOrderBy();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.operations_ecosys.OrderByBroadcast.serializeBinaryToWriter
    );
  }
};


/**
 * repeated BroadcastFilter filters = 1;
 * @return {!Array<!proto.operations_ecosys.BroadcastFilter>}
 */
proto.operations_ecosys.BroadcastQuery.prototype.getFiltersList = function() {
  return /** @type{!Array<!proto.operations_ecosys.BroadcastFilter>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.BroadcastFilter, 1));
};


/**
 * @param {!Array<!proto.operations_ecosys.BroadcastFilter>} value
 * @return {!proto.operations_ecosys.BroadcastQuery} returns this
*/
proto.operations_ecosys.BroadcastQuery.prototype.setFiltersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.operations_ecosys.BroadcastFilter=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.BroadcastFilter}
 */
proto.operations_ecosys.BroadcastQuery.prototype.addFilters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.operations_ecosys.BroadcastFilter, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.BroadcastQuery} returns this
 */
proto.operations_ecosys.BroadcastQuery.prototype.clearFiltersList = function() {
  return this.setFiltersList([]);
};


/**
 * optional int64 limit = 2;
 * @return {number}
 */
proto.operations_ecosys.BroadcastQuery.prototype.getLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.BroadcastQuery} returns this
 */
proto.operations_ecosys.BroadcastQuery.prototype.setLimit = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 skip = 3;
 * @return {number}
 */
proto.operations_ecosys.BroadcastQuery.prototype.getSkip = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.BroadcastQuery} returns this
 */
proto.operations_ecosys.BroadcastQuery.prototype.setSkip = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional OrderByBroadcast order_by = 4;
 * @return {?proto.operations_ecosys.OrderByBroadcast}
 */
proto.operations_ecosys.BroadcastQuery.prototype.getOrderBy = function() {
  return /** @type{?proto.operations_ecosys.OrderByBroadcast} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.OrderByBroadcast, 4));
};


/**
 * @param {?proto.operations_ecosys.OrderByBroadcast|undefined} value
 * @return {!proto.operations_ecosys.BroadcastQuery} returns this
*/
proto.operations_ecosys.BroadcastQuery.prototype.setOrderBy = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.BroadcastQuery} returns this
 */
proto.operations_ecosys.BroadcastQuery.prototype.clearOrderBy = function() {
  return this.setOrderBy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.BroadcastQuery.prototype.hasOrderBy = function() {
  return jspb.Message.getField(this, 4) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.OrderByBroadcast.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.OrderByBroadcast.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.OrderByBroadcast} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByBroadcast.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    orderBy: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.OrderByBroadcast}
 */
proto.operations_ecosys.OrderByBroadcast.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.OrderByBroadcast;
  return proto.operations_ecosys.OrderByBroadcast.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.OrderByBroadcast} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.OrderByBroadcast}
 */
proto.operations_ecosys.OrderByBroadcast.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.BroadcastFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = /** @type {!proto.operations_ecosys.OrderBy} */ (reader.readEnum());
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.OrderByBroadcast.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.OrderByBroadcast.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.OrderByBroadcast} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByBroadcast.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getOrderBy();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * optional BroadcastFilter.Field field = 1;
 * @return {!proto.operations_ecosys.BroadcastFilter.Field}
 */
proto.operations_ecosys.OrderByBroadcast.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.BroadcastFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.BroadcastFilter.Field} value
 * @return {!proto.operations_ecosys.OrderByBroadcast} returns this
 */
proto.operations_ecosys.OrderByBroadcast.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional OrderBy order_by = 2;
 * @return {!proto.operations_ecosys.OrderBy}
 */
proto.operations_ecosys.OrderByBroadcast.prototype.getOrderBy = function() {
  return /** @type {!proto.operations_ecosys.OrderBy} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.operations_ecosys.OrderBy} value
 * @return {!proto.operations_ecosys.OrderByBroadcast} returns this
 */
proto.operations_ecosys.OrderByBroadcast.prototype.setOrderBy = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.Roster.repeatedFields_ = [5,6];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.Roster.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.Roster.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.Roster} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Roster.toObject = function(includeInstance, msg) {
  var f, obj = {
    rosteringId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    aifsId: jspb.Message.getFieldWithDefault(msg, 2, 0),
    startTime: jspb.Message.getFieldWithDefault(msg, 3, ""),
    endTime: jspb.Message.getFieldWithDefault(msg, 4, ""),
    clientsList: jspb.Message.toObjectList(msg.getClientsList(),
    proto.operations_ecosys.AIFSClientRoster.toObject, includeInstance),
    guardAssignedList: jspb.Message.toObjectList(msg.getGuardAssignedList(),
    proto.operations_ecosys.RosterAssignement.toObject, includeInstance),
    status: jspb.Message.getFieldWithDefault(msg, 7, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.Roster}
 */
proto.operations_ecosys.Roster.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.Roster;
  return proto.operations_ecosys.Roster.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.Roster} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.Roster}
 */
proto.operations_ecosys.Roster.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setRosteringId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAifsId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setStartTime(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setEndTime(value);
      break;
    case 5:
      var value = new proto.operations_ecosys.AIFSClientRoster;
      reader.readMessage(value,proto.operations_ecosys.AIFSClientRoster.deserializeBinaryFromReader);
      msg.addClients(value);
      break;
    case 6:
      var value = new proto.operations_ecosys.RosterAssignement;
      reader.readMessage(value,proto.operations_ecosys.RosterAssignement.deserializeBinaryFromReader);
      msg.addGuardAssigned(value);
      break;
    case 7:
      var value = /** @type {!proto.operations_ecosys.Roster.Status} */ (reader.readEnum());
      msg.setStatus(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.Roster.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.Roster.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.Roster} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Roster.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRosteringId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getAifsId();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getStartTime();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getEndTime();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getClientsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      5,
      f,
      proto.operations_ecosys.AIFSClientRoster.serializeBinaryToWriter
    );
  }
  f = message.getGuardAssignedList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      6,
      f,
      proto.operations_ecosys.RosterAssignement.serializeBinaryToWriter
    );
  }
  f = message.getStatus();
  if (f !== 0.0) {
    writer.writeEnum(
      7,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.Roster.Status = {
  IS_DEFAULT: 0,
  PENDING: 1,
  CONFIRMED: 2,
  REJECTED: 3
};

/**
 * optional int64 rostering_id = 1;
 * @return {number}
 */
proto.operations_ecosys.Roster.prototype.getRosteringId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.Roster} returns this
 */
proto.operations_ecosys.Roster.prototype.setRosteringId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 aifs_id = 2;
 * @return {number}
 */
proto.operations_ecosys.Roster.prototype.getAifsId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.Roster} returns this
 */
proto.operations_ecosys.Roster.prototype.setAifsId = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional string start_time = 3;
 * @return {string}
 */
proto.operations_ecosys.Roster.prototype.getStartTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Roster} returns this
 */
proto.operations_ecosys.Roster.prototype.setStartTime = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string end_time = 4;
 * @return {string}
 */
proto.operations_ecosys.Roster.prototype.getEndTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Roster} returns this
 */
proto.operations_ecosys.Roster.prototype.setEndTime = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * repeated AIFSClientRoster clients = 5;
 * @return {!Array<!proto.operations_ecosys.AIFSClientRoster>}
 */
proto.operations_ecosys.Roster.prototype.getClientsList = function() {
  return /** @type{!Array<!proto.operations_ecosys.AIFSClientRoster>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.AIFSClientRoster, 5));
};


/**
 * @param {!Array<!proto.operations_ecosys.AIFSClientRoster>} value
 * @return {!proto.operations_ecosys.Roster} returns this
*/
proto.operations_ecosys.Roster.prototype.setClientsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 5, value);
};


/**
 * @param {!proto.operations_ecosys.AIFSClientRoster=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.AIFSClientRoster}
 */
proto.operations_ecosys.Roster.prototype.addClients = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 5, opt_value, proto.operations_ecosys.AIFSClientRoster, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.Roster} returns this
 */
proto.operations_ecosys.Roster.prototype.clearClientsList = function() {
  return this.setClientsList([]);
};


/**
 * repeated RosterAssignement guard_assigned = 6;
 * @return {!Array<!proto.operations_ecosys.RosterAssignement>}
 */
proto.operations_ecosys.Roster.prototype.getGuardAssignedList = function() {
  return /** @type{!Array<!proto.operations_ecosys.RosterAssignement>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.RosterAssignement, 6));
};


/**
 * @param {!Array<!proto.operations_ecosys.RosterAssignement>} value
 * @return {!proto.operations_ecosys.Roster} returns this
*/
proto.operations_ecosys.Roster.prototype.setGuardAssignedList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 6, value);
};


/**
 * @param {!proto.operations_ecosys.RosterAssignement=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.RosterAssignement}
 */
proto.operations_ecosys.Roster.prototype.addGuardAssigned = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 6, opt_value, proto.operations_ecosys.RosterAssignement, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.Roster} returns this
 */
proto.operations_ecosys.Roster.prototype.clearGuardAssignedList = function() {
  return this.setGuardAssignedList([]);
};


/**
 * optional Status status = 7;
 * @return {!proto.operations_ecosys.Roster.Status}
 */
proto.operations_ecosys.Roster.prototype.getStatus = function() {
  return /** @type {!proto.operations_ecosys.Roster.Status} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {!proto.operations_ecosys.Roster.Status} value
 * @return {!proto.operations_ecosys.Roster} returns this
 */
proto.operations_ecosys.Roster.prototype.setStatus = function(value) {
  return jspb.Message.setProto3EnumField(this, 7, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.AIFSClientRoster.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.AIFSClientRoster.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.AIFSClientRoster} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.AIFSClientRoster.toObject = function(includeInstance, msg) {
  var f, obj = {
    aifsClientRosterId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    client: (f = msg.getClient()) && proto.operations_ecosys.Client.toObject(includeInstance, f),
    patrolOrder: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.AIFSClientRoster}
 */
proto.operations_ecosys.AIFSClientRoster.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.AIFSClientRoster;
  return proto.operations_ecosys.AIFSClientRoster.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.AIFSClientRoster} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.AIFSClientRoster}
 */
proto.operations_ecosys.AIFSClientRoster.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAifsClientRosterId(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Client;
      reader.readMessage(value,proto.operations_ecosys.Client.deserializeBinaryFromReader);
      msg.setClient(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPatrolOrder(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.AIFSClientRoster.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.AIFSClientRoster.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.AIFSClientRoster} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.AIFSClientRoster.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAifsClientRosterId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getClient();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Client.serializeBinaryToWriter
    );
  }
  f = message.getPatrolOrder();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
};


/**
 * optional int64 aifs_client_roster_id = 1;
 * @return {number}
 */
proto.operations_ecosys.AIFSClientRoster.prototype.getAifsClientRosterId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.AIFSClientRoster} returns this
 */
proto.operations_ecosys.AIFSClientRoster.prototype.setAifsClientRosterId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional Client client = 2;
 * @return {?proto.operations_ecosys.Client}
 */
proto.operations_ecosys.AIFSClientRoster.prototype.getClient = function() {
  return /** @type{?proto.operations_ecosys.Client} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Client, 2));
};


/**
 * @param {?proto.operations_ecosys.Client|undefined} value
 * @return {!proto.operations_ecosys.AIFSClientRoster} returns this
*/
proto.operations_ecosys.AIFSClientRoster.prototype.setClient = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.AIFSClientRoster} returns this
 */
proto.operations_ecosys.AIFSClientRoster.prototype.clearClient = function() {
  return this.setClient(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.AIFSClientRoster.prototype.hasClient = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional int64 patrol_order = 3;
 * @return {number}
 */
proto.operations_ecosys.AIFSClientRoster.prototype.getPatrolOrder = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.AIFSClientRoster} returns this
 */
proto.operations_ecosys.AIFSClientRoster.prototype.setPatrolOrder = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.RosterAssignement.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.RosterAssignement.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.RosterAssignement} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.RosterAssignement.toObject = function(includeInstance, msg) {
  var f, obj = {
    rosterAssignmentId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    guardAssigned: (f = msg.getGuardAssigned()) && proto.operations_ecosys.EmployeeEvaluation.toObject(includeInstance, f),
    customStartTime: (f = msg.getCustomStartTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    customEndTime: (f = msg.getCustomEndTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    confirmed: jspb.Message.getBooleanFieldWithDefault(msg, 5, false),
    attended: jspb.Message.getBooleanFieldWithDefault(msg, 6, false),
    attendanceTime: (f = msg.getAttendanceTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    isAssigned: jspb.Message.getBooleanFieldWithDefault(msg, 8, false),
    rejected: jspb.Message.getBooleanFieldWithDefault(msg, 9, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.RosterAssignement}
 */
proto.operations_ecosys.RosterAssignement.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.RosterAssignement;
  return proto.operations_ecosys.RosterAssignement.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.RosterAssignement} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.RosterAssignement}
 */
proto.operations_ecosys.RosterAssignement.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setRosterAssignmentId(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.EmployeeEvaluation;
      reader.readMessage(value,proto.operations_ecosys.EmployeeEvaluation.deserializeBinaryFromReader);
      msg.setGuardAssigned(value);
      break;
    case 3:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCustomStartTime(value);
      break;
    case 4:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCustomEndTime(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setConfirmed(value);
      break;
    case 6:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setAttended(value);
      break;
    case 7:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setAttendanceTime(value);
      break;
    case 8:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsAssigned(value);
      break;
    case 9:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setRejected(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.RosterAssignement.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.RosterAssignement.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.RosterAssignement} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.RosterAssignement.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRosterAssignmentId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getGuardAssigned();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.EmployeeEvaluation.serializeBinaryToWriter
    );
  }
  f = message.getCustomStartTime();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getCustomEndTime();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getConfirmed();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
  f = message.getAttended();
  if (f) {
    writer.writeBool(
      6,
      f
    );
  }
  f = message.getAttendanceTime();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getIsAssigned();
  if (f) {
    writer.writeBool(
      8,
      f
    );
  }
  f = message.getRejected();
  if (f) {
    writer.writeBool(
      9,
      f
    );
  }
};


/**
 * optional int64 roster_assignment_id = 1;
 * @return {number}
 */
proto.operations_ecosys.RosterAssignement.prototype.getRosterAssignmentId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
 */
proto.operations_ecosys.RosterAssignement.prototype.setRosterAssignmentId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional EmployeeEvaluation guard_assigned = 2;
 * @return {?proto.operations_ecosys.EmployeeEvaluation}
 */
proto.operations_ecosys.RosterAssignement.prototype.getGuardAssigned = function() {
  return /** @type{?proto.operations_ecosys.EmployeeEvaluation} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.EmployeeEvaluation, 2));
};


/**
 * @param {?proto.operations_ecosys.EmployeeEvaluation|undefined} value
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
*/
proto.operations_ecosys.RosterAssignement.prototype.setGuardAssigned = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
 */
proto.operations_ecosys.RosterAssignement.prototype.clearGuardAssigned = function() {
  return this.setGuardAssigned(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.RosterAssignement.prototype.hasGuardAssigned = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.Timestamp custom_start_time = 3;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.operations_ecosys.RosterAssignement.prototype.getCustomStartTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
*/
proto.operations_ecosys.RosterAssignement.prototype.setCustomStartTime = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
 */
proto.operations_ecosys.RosterAssignement.prototype.clearCustomStartTime = function() {
  return this.setCustomStartTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.RosterAssignement.prototype.hasCustomStartTime = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.Timestamp custom_end_time = 4;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.operations_ecosys.RosterAssignement.prototype.getCustomEndTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 4));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
*/
proto.operations_ecosys.RosterAssignement.prototype.setCustomEndTime = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
 */
proto.operations_ecosys.RosterAssignement.prototype.clearCustomEndTime = function() {
  return this.setCustomEndTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.RosterAssignement.prototype.hasCustomEndTime = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional bool confirmed = 5;
 * @return {boolean}
 */
proto.operations_ecosys.RosterAssignement.prototype.getConfirmed = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 5, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
 */
proto.operations_ecosys.RosterAssignement.prototype.setConfirmed = function(value) {
  return jspb.Message.setProto3BooleanField(this, 5, value);
};


/**
 * optional bool attended = 6;
 * @return {boolean}
 */
proto.operations_ecosys.RosterAssignement.prototype.getAttended = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 6, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
 */
proto.operations_ecosys.RosterAssignement.prototype.setAttended = function(value) {
  return jspb.Message.setProto3BooleanField(this, 6, value);
};


/**
 * optional google.protobuf.Timestamp attendance_time = 7;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.operations_ecosys.RosterAssignement.prototype.getAttendanceTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 7));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
*/
proto.operations_ecosys.RosterAssignement.prototype.setAttendanceTime = function(value) {
  return jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
 */
proto.operations_ecosys.RosterAssignement.prototype.clearAttendanceTime = function() {
  return this.setAttendanceTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.RosterAssignement.prototype.hasAttendanceTime = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional bool is_assigned = 8;
 * @return {boolean}
 */
proto.operations_ecosys.RosterAssignement.prototype.getIsAssigned = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 8, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
 */
proto.operations_ecosys.RosterAssignement.prototype.setIsAssigned = function(value) {
  return jspb.Message.setProto3BooleanField(this, 8, value);
};


/**
 * optional bool rejected = 9;
 * @return {boolean}
 */
proto.operations_ecosys.RosterAssignement.prototype.getRejected = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 9, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.RosterAssignement} returns this
 */
proto.operations_ecosys.RosterAssignement.prototype.setRejected = function(value) {
  return jspb.Message.setProto3BooleanField(this, 9, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.BulkRosters.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.BulkRosters.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.BulkRosters.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.BulkRosters} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.BulkRosters.toObject = function(includeInstance, msg) {
  var f, obj = {
    rostersList: jspb.Message.toObjectList(msg.getRostersList(),
    proto.operations_ecosys.Roster.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.BulkRosters}
 */
proto.operations_ecosys.BulkRosters.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.BulkRosters;
  return proto.operations_ecosys.BulkRosters.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.BulkRosters} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.BulkRosters}
 */
proto.operations_ecosys.BulkRosters.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.Roster;
      reader.readMessage(value,proto.operations_ecosys.Roster.deserializeBinaryFromReader);
      msg.addRosters(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.BulkRosters.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.BulkRosters.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.BulkRosters} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.BulkRosters.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRostersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.operations_ecosys.Roster.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Roster rosters = 1;
 * @return {!Array<!proto.operations_ecosys.Roster>}
 */
proto.operations_ecosys.BulkRosters.prototype.getRostersList = function() {
  return /** @type{!Array<!proto.operations_ecosys.Roster>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.Roster, 1));
};


/**
 * @param {!Array<!proto.operations_ecosys.Roster>} value
 * @return {!proto.operations_ecosys.BulkRosters} returns this
*/
proto.operations_ecosys.BulkRosters.prototype.setRostersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.operations_ecosys.Roster=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.Roster}
 */
proto.operations_ecosys.BulkRosters.prototype.addRosters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.operations_ecosys.Roster, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.BulkRosters} returns this
 */
proto.operations_ecosys.BulkRosters.prototype.clearRostersList = function() {
  return this.setRostersList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.RosterResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.RosterResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.RosterResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.RosterResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    response: (f = msg.getResponse()) && proto.operations_ecosys.Response.toObject(includeInstance, f),
    roster: (f = msg.getRoster()) && proto.operations_ecosys.Roster.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.RosterResponse}
 */
proto.operations_ecosys.RosterResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.RosterResponse;
  return proto.operations_ecosys.RosterResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.RosterResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.RosterResponse}
 */
proto.operations_ecosys.RosterResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.Response;
      reader.readMessage(value,proto.operations_ecosys.Response.deserializeBinaryFromReader);
      msg.setResponse(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Roster;
      reader.readMessage(value,proto.operations_ecosys.Roster.deserializeBinaryFromReader);
      msg.setRoster(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.RosterResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.RosterResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.RosterResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.RosterResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResponse();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.operations_ecosys.Response.serializeBinaryToWriter
    );
  }
  f = message.getRoster();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Roster.serializeBinaryToWriter
    );
  }
};


/**
 * optional Response response = 1;
 * @return {?proto.operations_ecosys.Response}
 */
proto.operations_ecosys.RosterResponse.prototype.getResponse = function() {
  return /** @type{?proto.operations_ecosys.Response} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Response, 1));
};


/**
 * @param {?proto.operations_ecosys.Response|undefined} value
 * @return {!proto.operations_ecosys.RosterResponse} returns this
*/
proto.operations_ecosys.RosterResponse.prototype.setResponse = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.RosterResponse} returns this
 */
proto.operations_ecosys.RosterResponse.prototype.clearResponse = function() {
  return this.setResponse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.RosterResponse.prototype.hasResponse = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Roster roster = 2;
 * @return {?proto.operations_ecosys.Roster}
 */
proto.operations_ecosys.RosterResponse.prototype.getRoster = function() {
  return /** @type{?proto.operations_ecosys.Roster} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Roster, 2));
};


/**
 * @param {?proto.operations_ecosys.Roster|undefined} value
 * @return {!proto.operations_ecosys.RosterResponse} returns this
*/
proto.operations_ecosys.RosterResponse.prototype.setRoster = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.RosterResponse} returns this
 */
proto.operations_ecosys.RosterResponse.prototype.clearRoster = function() {
  return this.setRoster(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.RosterResponse.prototype.hasRoster = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.RosterAssignmentResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.RosterAssignmentResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.RosterAssignmentResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.RosterAssignmentResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    response: (f = msg.getResponse()) && proto.operations_ecosys.Response.toObject(includeInstance, f),
    rosterAssignment: (f = msg.getRosterAssignment()) && proto.operations_ecosys.RosterAssignement.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.RosterAssignmentResponse}
 */
proto.operations_ecosys.RosterAssignmentResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.RosterAssignmentResponse;
  return proto.operations_ecosys.RosterAssignmentResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.RosterAssignmentResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.RosterAssignmentResponse}
 */
proto.operations_ecosys.RosterAssignmentResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.Response;
      reader.readMessage(value,proto.operations_ecosys.Response.deserializeBinaryFromReader);
      msg.setResponse(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.RosterAssignement;
      reader.readMessage(value,proto.operations_ecosys.RosterAssignement.deserializeBinaryFromReader);
      msg.setRosterAssignment(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.RosterAssignmentResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.RosterAssignmentResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.RosterAssignmentResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.RosterAssignmentResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResponse();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.operations_ecosys.Response.serializeBinaryToWriter
    );
  }
  f = message.getRosterAssignment();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.RosterAssignement.serializeBinaryToWriter
    );
  }
};


/**
 * optional Response response = 1;
 * @return {?proto.operations_ecosys.Response}
 */
proto.operations_ecosys.RosterAssignmentResponse.prototype.getResponse = function() {
  return /** @type{?proto.operations_ecosys.Response} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Response, 1));
};


/**
 * @param {?proto.operations_ecosys.Response|undefined} value
 * @return {!proto.operations_ecosys.RosterAssignmentResponse} returns this
*/
proto.operations_ecosys.RosterAssignmentResponse.prototype.setResponse = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.RosterAssignmentResponse} returns this
 */
proto.operations_ecosys.RosterAssignmentResponse.prototype.clearResponse = function() {
  return this.setResponse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.RosterAssignmentResponse.prototype.hasResponse = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional RosterAssignement roster_assignment = 2;
 * @return {?proto.operations_ecosys.RosterAssignement}
 */
proto.operations_ecosys.RosterAssignmentResponse.prototype.getRosterAssignment = function() {
  return /** @type{?proto.operations_ecosys.RosterAssignement} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.RosterAssignement, 2));
};


/**
 * @param {?proto.operations_ecosys.RosterAssignement|undefined} value
 * @return {!proto.operations_ecosys.RosterAssignmentResponse} returns this
*/
proto.operations_ecosys.RosterAssignmentResponse.prototype.setRosterAssignment = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.RosterAssignmentResponse} returns this
 */
proto.operations_ecosys.RosterAssignmentResponse.prototype.clearRosterAssignment = function() {
  return this.setRosterAssignment(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.RosterAssignmentResponse.prototype.hasRosterAssignment = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.RosterFilter.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.RosterFilter.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.RosterFilter} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.RosterFilter.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    comparisons: (f = msg.getComparisons()) && proto.operations_ecosys.Filter.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.RosterFilter}
 */
proto.operations_ecosys.RosterFilter.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.RosterFilter;
  return proto.operations_ecosys.RosterFilter.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.RosterFilter} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.RosterFilter}
 */
proto.operations_ecosys.RosterFilter.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.RosterFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Filter;
      reader.readMessage(value,proto.operations_ecosys.Filter.deserializeBinaryFromReader);
      msg.setComparisons(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.RosterFilter.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.RosterFilter.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.RosterFilter} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.RosterFilter.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getComparisons();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Filter.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.RosterFilter.Field = {
  ROSTER_ID: 0,
  ROSTER_ASSIGNMENT_ID: 1,
  ROSTER_AIFS_CLIENT_ID: 2,
  AIFS_ID: 3,
  GUARD_ASSIGNED_ID: 4,
  CLIENT_ID: 5,
  GUARD_ASSIGNMENT_CONFIRMATION: 6,
  GUARD_ASSIGNMENT_ATTENDED: 7,
  START_TIME: 8,
  END_TIME: 9,
  IS_ASSIGNED: 10,
  DEFAULT_ROSTERING_DAY_OF_WEEK: 11,
  GUARD_ASSIGNMENT_REJECTION: 12
};

/**
 * optional Field field = 1;
 * @return {!proto.operations_ecosys.RosterFilter.Field}
 */
proto.operations_ecosys.RosterFilter.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.RosterFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.RosterFilter.Field} value
 * @return {!proto.operations_ecosys.RosterFilter} returns this
 */
proto.operations_ecosys.RosterFilter.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional Filter comparisons = 2;
 * @return {?proto.operations_ecosys.Filter}
 */
proto.operations_ecosys.RosterFilter.prototype.getComparisons = function() {
  return /** @type{?proto.operations_ecosys.Filter} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Filter, 2));
};


/**
 * @param {?proto.operations_ecosys.Filter|undefined} value
 * @return {!proto.operations_ecosys.RosterFilter} returns this
*/
proto.operations_ecosys.RosterFilter.prototype.setComparisons = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.RosterFilter} returns this
 */
proto.operations_ecosys.RosterFilter.prototype.clearComparisons = function() {
  return this.setComparisons(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.RosterFilter.prototype.hasComparisons = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.RosterQuery.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.RosterQuery.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.RosterQuery.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.RosterQuery} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.RosterQuery.toObject = function(includeInstance, msg) {
  var f, obj = {
    filtersList: jspb.Message.toObjectList(msg.getFiltersList(),
    proto.operations_ecosys.RosterFilter.toObject, includeInstance),
    limit: jspb.Message.getFieldWithDefault(msg, 2, 0),
    skip: jspb.Message.getFieldWithDefault(msg, 3, 0),
    orderBy: (f = msg.getOrderBy()) && proto.operations_ecosys.OrderByRoster.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.RosterQuery}
 */
proto.operations_ecosys.RosterQuery.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.RosterQuery;
  return proto.operations_ecosys.RosterQuery.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.RosterQuery} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.RosterQuery}
 */
proto.operations_ecosys.RosterQuery.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.RosterFilter;
      reader.readMessage(value,proto.operations_ecosys.RosterFilter.deserializeBinaryFromReader);
      msg.addFilters(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLimit(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSkip(value);
      break;
    case 4:
      var value = new proto.operations_ecosys.OrderByRoster;
      reader.readMessage(value,proto.operations_ecosys.OrderByRoster.deserializeBinaryFromReader);
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.RosterQuery.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.RosterQuery.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.RosterQuery} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.RosterQuery.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFiltersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.operations_ecosys.RosterFilter.serializeBinaryToWriter
    );
  }
  f = message.getLimit();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getSkip();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getOrderBy();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.operations_ecosys.OrderByRoster.serializeBinaryToWriter
    );
  }
};


/**
 * repeated RosterFilter filters = 1;
 * @return {!Array<!proto.operations_ecosys.RosterFilter>}
 */
proto.operations_ecosys.RosterQuery.prototype.getFiltersList = function() {
  return /** @type{!Array<!proto.operations_ecosys.RosterFilter>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.RosterFilter, 1));
};


/**
 * @param {!Array<!proto.operations_ecosys.RosterFilter>} value
 * @return {!proto.operations_ecosys.RosterQuery} returns this
*/
proto.operations_ecosys.RosterQuery.prototype.setFiltersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.operations_ecosys.RosterFilter=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.RosterFilter}
 */
proto.operations_ecosys.RosterQuery.prototype.addFilters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.operations_ecosys.RosterFilter, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.RosterQuery} returns this
 */
proto.operations_ecosys.RosterQuery.prototype.clearFiltersList = function() {
  return this.setFiltersList([]);
};


/**
 * optional int64 limit = 2;
 * @return {number}
 */
proto.operations_ecosys.RosterQuery.prototype.getLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.RosterQuery} returns this
 */
proto.operations_ecosys.RosterQuery.prototype.setLimit = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 skip = 3;
 * @return {number}
 */
proto.operations_ecosys.RosterQuery.prototype.getSkip = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.RosterQuery} returns this
 */
proto.operations_ecosys.RosterQuery.prototype.setSkip = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional OrderByRoster order_by = 4;
 * @return {?proto.operations_ecosys.OrderByRoster}
 */
proto.operations_ecosys.RosterQuery.prototype.getOrderBy = function() {
  return /** @type{?proto.operations_ecosys.OrderByRoster} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.OrderByRoster, 4));
};


/**
 * @param {?proto.operations_ecosys.OrderByRoster|undefined} value
 * @return {!proto.operations_ecosys.RosterQuery} returns this
*/
proto.operations_ecosys.RosterQuery.prototype.setOrderBy = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.RosterQuery} returns this
 */
proto.operations_ecosys.RosterQuery.prototype.clearOrderBy = function() {
  return this.setOrderBy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.RosterQuery.prototype.hasOrderBy = function() {
  return jspb.Message.getField(this, 4) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.OrderByRoster.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.OrderByRoster.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.OrderByRoster} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByRoster.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    orderBy: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.OrderByRoster}
 */
proto.operations_ecosys.OrderByRoster.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.OrderByRoster;
  return proto.operations_ecosys.OrderByRoster.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.OrderByRoster} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.OrderByRoster}
 */
proto.operations_ecosys.OrderByRoster.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.RosterFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = /** @type {!proto.operations_ecosys.OrderBy} */ (reader.readEnum());
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.OrderByRoster.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.OrderByRoster.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.OrderByRoster} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByRoster.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getOrderBy();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * optional RosterFilter.Field field = 1;
 * @return {!proto.operations_ecosys.RosterFilter.Field}
 */
proto.operations_ecosys.OrderByRoster.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.RosterFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.RosterFilter.Field} value
 * @return {!proto.operations_ecosys.OrderByRoster} returns this
 */
proto.operations_ecosys.OrderByRoster.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional OrderBy order_by = 2;
 * @return {!proto.operations_ecosys.OrderBy}
 */
proto.operations_ecosys.OrderByRoster.prototype.getOrderBy = function() {
  return /** @type {!proto.operations_ecosys.OrderBy} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.operations_ecosys.OrderBy} value
 * @return {!proto.operations_ecosys.OrderByRoster} returns this
 */
proto.operations_ecosys.OrderByRoster.prototype.setOrderBy = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.AvailabilityQuery.repeatedFields_ = [5];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.AvailabilityQuery.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.AvailabilityQuery.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.AvailabilityQuery} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.AvailabilityQuery.toObject = function(includeInstance, msg) {
  var f, obj = {
    startTime: jspb.Message.getFieldWithDefault(msg, 1, ""),
    endTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    limit: jspb.Message.getFieldWithDefault(msg, 3, 0),
    skip: jspb.Message.getFieldWithDefault(msg, 4, 0),
    filtersList: jspb.Message.toObjectList(msg.getFiltersList(),
    proto.operations_ecosys.AvailabilityFilter.toObject, includeInstance),
    orderBy: (f = msg.getOrderBy()) && proto.operations_ecosys.OrderByQuery.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.AvailabilityQuery}
 */
proto.operations_ecosys.AvailabilityQuery.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.AvailabilityQuery;
  return proto.operations_ecosys.AvailabilityQuery.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.AvailabilityQuery} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.AvailabilityQuery}
 */
proto.operations_ecosys.AvailabilityQuery.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setStartTime(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setEndTime(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLimit(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSkip(value);
      break;
    case 5:
      var value = new proto.operations_ecosys.AvailabilityFilter;
      reader.readMessage(value,proto.operations_ecosys.AvailabilityFilter.deserializeBinaryFromReader);
      msg.addFilters(value);
      break;
    case 6:
      var value = new proto.operations_ecosys.OrderByQuery;
      reader.readMessage(value,proto.operations_ecosys.OrderByQuery.deserializeBinaryFromReader);
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.AvailabilityQuery.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.AvailabilityQuery.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.AvailabilityQuery} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.AvailabilityQuery.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getStartTime();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getEndTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getLimit();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getSkip();
  if (f !== 0) {
    writer.writeInt64(
      4,
      f
    );
  }
  f = message.getFiltersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      5,
      f,
      proto.operations_ecosys.AvailabilityFilter.serializeBinaryToWriter
    );
  }
  f = message.getOrderBy();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.operations_ecosys.OrderByQuery.serializeBinaryToWriter
    );
  }
};


/**
 * optional string start_time = 1;
 * @return {string}
 */
proto.operations_ecosys.AvailabilityQuery.prototype.getStartTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.AvailabilityQuery} returns this
 */
proto.operations_ecosys.AvailabilityQuery.prototype.setStartTime = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string end_time = 2;
 * @return {string}
 */
proto.operations_ecosys.AvailabilityQuery.prototype.getEndTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.AvailabilityQuery} returns this
 */
proto.operations_ecosys.AvailabilityQuery.prototype.setEndTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int64 limit = 3;
 * @return {number}
 */
proto.operations_ecosys.AvailabilityQuery.prototype.getLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.AvailabilityQuery} returns this
 */
proto.operations_ecosys.AvailabilityQuery.prototype.setLimit = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int64 skip = 4;
 * @return {number}
 */
proto.operations_ecosys.AvailabilityQuery.prototype.getSkip = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.AvailabilityQuery} returns this
 */
proto.operations_ecosys.AvailabilityQuery.prototype.setSkip = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * repeated AvailabilityFilter filters = 5;
 * @return {!Array<!proto.operations_ecosys.AvailabilityFilter>}
 */
proto.operations_ecosys.AvailabilityQuery.prototype.getFiltersList = function() {
  return /** @type{!Array<!proto.operations_ecosys.AvailabilityFilter>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.AvailabilityFilter, 5));
};


/**
 * @param {!Array<!proto.operations_ecosys.AvailabilityFilter>} value
 * @return {!proto.operations_ecosys.AvailabilityQuery} returns this
*/
proto.operations_ecosys.AvailabilityQuery.prototype.setFiltersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 5, value);
};


/**
 * @param {!proto.operations_ecosys.AvailabilityFilter=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.AvailabilityFilter}
 */
proto.operations_ecosys.AvailabilityQuery.prototype.addFilters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 5, opt_value, proto.operations_ecosys.AvailabilityFilter, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.AvailabilityQuery} returns this
 */
proto.operations_ecosys.AvailabilityQuery.prototype.clearFiltersList = function() {
  return this.setFiltersList([]);
};


/**
 * optional OrderByQuery order_by = 6;
 * @return {?proto.operations_ecosys.OrderByQuery}
 */
proto.operations_ecosys.AvailabilityQuery.prototype.getOrderBy = function() {
  return /** @type{?proto.operations_ecosys.OrderByQuery} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.OrderByQuery, 6));
};


/**
 * @param {?proto.operations_ecosys.OrderByQuery|undefined} value
 * @return {!proto.operations_ecosys.AvailabilityQuery} returns this
*/
proto.operations_ecosys.AvailabilityQuery.prototype.setOrderBy = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.AvailabilityQuery} returns this
 */
proto.operations_ecosys.AvailabilityQuery.prototype.clearOrderBy = function() {
  return this.setOrderBy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.AvailabilityQuery.prototype.hasOrderBy = function() {
  return jspb.Message.getField(this, 6) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.OrderByQuery.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.OrderByQuery.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.OrderByQuery} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByQuery.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    orderBy: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.OrderByQuery}
 */
proto.operations_ecosys.OrderByQuery.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.OrderByQuery;
  return proto.operations_ecosys.OrderByQuery.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.OrderByQuery} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.OrderByQuery}
 */
proto.operations_ecosys.OrderByQuery.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.AvailabilityFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = /** @type {!proto.operations_ecosys.OrderBy} */ (reader.readEnum());
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.OrderByQuery.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.OrderByQuery.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.OrderByQuery} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByQuery.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getOrderBy();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * optional AvailabilityFilter.Field field = 1;
 * @return {!proto.operations_ecosys.AvailabilityFilter.Field}
 */
proto.operations_ecosys.OrderByQuery.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.AvailabilityFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.AvailabilityFilter.Field} value
 * @return {!proto.operations_ecosys.OrderByQuery} returns this
 */
proto.operations_ecosys.OrderByQuery.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional OrderBy order_by = 2;
 * @return {!proto.operations_ecosys.OrderBy}
 */
proto.operations_ecosys.OrderByQuery.prototype.getOrderBy = function() {
  return /** @type {!proto.operations_ecosys.OrderBy} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.operations_ecosys.OrderBy} value
 * @return {!proto.operations_ecosys.OrderByQuery} returns this
 */
proto.operations_ecosys.OrderByQuery.prototype.setOrderBy = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.AvailabilityFilter.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.AvailabilityFilter.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.AvailabilityFilter} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.AvailabilityFilter.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    comparisons: (f = msg.getComparisons()) && proto.operations_ecosys.Filter.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.AvailabilityFilter}
 */
proto.operations_ecosys.AvailabilityFilter.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.AvailabilityFilter;
  return proto.operations_ecosys.AvailabilityFilter.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.AvailabilityFilter} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.AvailabilityFilter}
 */
proto.operations_ecosys.AvailabilityFilter.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.AvailabilityFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Filter;
      reader.readMessage(value,proto.operations_ecosys.Filter.deserializeBinaryFromReader);
      msg.setComparisons(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.AvailabilityFilter.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.AvailabilityFilter.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.AvailabilityFilter} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.AvailabilityFilter.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getComparisons();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Filter.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.AvailabilityFilter.Field = {
  AVAILABILITY_ID: 0,
  WEEK: 1,
  YEAR: 2,
  GUARD_ID: 3,
  SUN: 4,
  MON: 5,
  TUES: 6,
  WED: 7,
  THURS: 8,
  FRI: 9,
  SAT: 10,
  NEXT_SUN: 11
};

/**
 * optional Field field = 1;
 * @return {!proto.operations_ecosys.AvailabilityFilter.Field}
 */
proto.operations_ecosys.AvailabilityFilter.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.AvailabilityFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.AvailabilityFilter.Field} value
 * @return {!proto.operations_ecosys.AvailabilityFilter} returns this
 */
proto.operations_ecosys.AvailabilityFilter.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional Filter comparisons = 2;
 * @return {?proto.operations_ecosys.Filter}
 */
proto.operations_ecosys.AvailabilityFilter.prototype.getComparisons = function() {
  return /** @type{?proto.operations_ecosys.Filter} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Filter, 2));
};


/**
 * @param {?proto.operations_ecosys.Filter|undefined} value
 * @return {!proto.operations_ecosys.AvailabilityFilter} returns this
*/
proto.operations_ecosys.AvailabilityFilter.prototype.setComparisons = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.AvailabilityFilter} returns this
 */
proto.operations_ecosys.AvailabilityFilter.prototype.clearComparisons = function() {
  return this.setComparisons(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.AvailabilityFilter.prototype.hasComparisons = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.EmployeeEvaluationResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.EmployeeEvaluationResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.EmployeeEvaluationResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.EmployeeEvaluationResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    response: (f = msg.getResponse()) && proto.operations_ecosys.Response.toObject(includeInstance, f),
    employee: (f = msg.getEmployee()) && proto.operations_ecosys.EmployeeEvaluation.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.EmployeeEvaluationResponse}
 */
proto.operations_ecosys.EmployeeEvaluationResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.EmployeeEvaluationResponse;
  return proto.operations_ecosys.EmployeeEvaluationResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.EmployeeEvaluationResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.EmployeeEvaluationResponse}
 */
proto.operations_ecosys.EmployeeEvaluationResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.Response;
      reader.readMessage(value,proto.operations_ecosys.Response.deserializeBinaryFromReader);
      msg.setResponse(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.EmployeeEvaluation;
      reader.readMessage(value,proto.operations_ecosys.EmployeeEvaluation.deserializeBinaryFromReader);
      msg.setEmployee(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.EmployeeEvaluationResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.EmployeeEvaluationResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.EmployeeEvaluationResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.EmployeeEvaluationResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResponse();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.operations_ecosys.Response.serializeBinaryToWriter
    );
  }
  f = message.getEmployee();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.EmployeeEvaluation.serializeBinaryToWriter
    );
  }
};


/**
 * optional Response response = 1;
 * @return {?proto.operations_ecosys.Response}
 */
proto.operations_ecosys.EmployeeEvaluationResponse.prototype.getResponse = function() {
  return /** @type{?proto.operations_ecosys.Response} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Response, 1));
};


/**
 * @param {?proto.operations_ecosys.Response|undefined} value
 * @return {!proto.operations_ecosys.EmployeeEvaluationResponse} returns this
*/
proto.operations_ecosys.EmployeeEvaluationResponse.prototype.setResponse = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.EmployeeEvaluationResponse} returns this
 */
proto.operations_ecosys.EmployeeEvaluationResponse.prototype.clearResponse = function() {
  return this.setResponse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.EmployeeEvaluationResponse.prototype.hasResponse = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional EmployeeEvaluation employee = 2;
 * @return {?proto.operations_ecosys.EmployeeEvaluation}
 */
proto.operations_ecosys.EmployeeEvaluationResponse.prototype.getEmployee = function() {
  return /** @type{?proto.operations_ecosys.EmployeeEvaluation} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.EmployeeEvaluation, 2));
};


/**
 * @param {?proto.operations_ecosys.EmployeeEvaluation|undefined} value
 * @return {!proto.operations_ecosys.EmployeeEvaluationResponse} returns this
*/
proto.operations_ecosys.EmployeeEvaluationResponse.prototype.setEmployee = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.EmployeeEvaluationResponse} returns this
 */
proto.operations_ecosys.EmployeeEvaluationResponse.prototype.clearEmployee = function() {
  return this.setEmployee(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.EmployeeEvaluationResponse.prototype.hasEmployee = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.EmployeeEvaluation.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.EmployeeEvaluation.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.EmployeeEvaluation} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.EmployeeEvaluation.toObject = function(includeInstance, msg) {
  var f, obj = {
    employee: (f = msg.getEmployee()) && proto.operations_ecosys.User.toObject(includeInstance, f),
    employeeScore: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0),
    isAvailable: jspb.Message.getBooleanFieldWithDefault(msg, 3, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.EmployeeEvaluation}
 */
proto.operations_ecosys.EmployeeEvaluation.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.EmployeeEvaluation;
  return proto.operations_ecosys.EmployeeEvaluation.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.EmployeeEvaluation} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.EmployeeEvaluation}
 */
proto.operations_ecosys.EmployeeEvaluation.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.User;
      reader.readMessage(value,proto.operations_ecosys.User.deserializeBinaryFromReader);
      msg.setEmployee(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setEmployeeScore(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsAvailable(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.EmployeeEvaluation.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.EmployeeEvaluation.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.EmployeeEvaluation} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.EmployeeEvaluation.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEmployee();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.operations_ecosys.User.serializeBinaryToWriter
    );
  }
  f = message.getEmployeeScore();
  if (f !== 0.0) {
    writer.writeFloat(
      2,
      f
    );
  }
  f = message.getIsAvailable();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
};


/**
 * optional User employee = 1;
 * @return {?proto.operations_ecosys.User}
 */
proto.operations_ecosys.EmployeeEvaluation.prototype.getEmployee = function() {
  return /** @type{?proto.operations_ecosys.User} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.User, 1));
};


/**
 * @param {?proto.operations_ecosys.User|undefined} value
 * @return {!proto.operations_ecosys.EmployeeEvaluation} returns this
*/
proto.operations_ecosys.EmployeeEvaluation.prototype.setEmployee = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.EmployeeEvaluation} returns this
 */
proto.operations_ecosys.EmployeeEvaluation.prototype.clearEmployee = function() {
  return this.setEmployee(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.EmployeeEvaluation.prototype.hasEmployee = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional float employee_score = 2;
 * @return {number}
 */
proto.operations_ecosys.EmployeeEvaluation.prototype.getEmployeeScore = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.EmployeeEvaluation} returns this
 */
proto.operations_ecosys.EmployeeEvaluation.prototype.setEmployeeScore = function(value) {
  return jspb.Message.setProto3FloatField(this, 2, value);
};


/**
 * optional bool is_available = 3;
 * @return {boolean}
 */
proto.operations_ecosys.EmployeeEvaluation.prototype.getIsAvailable = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.EmployeeEvaluation} returns this
 */
proto.operations_ecosys.EmployeeEvaluation.prototype.setIsAvailable = function(value) {
  return jspb.Message.setProto3BooleanField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.IncidentReport.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.IncidentReport.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.IncidentReport} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.IncidentReport.toObject = function(includeInstance, msg) {
  var f, obj = {
    incidentReportId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    type: jspb.Message.getFieldWithDefault(msg, 2, 0),
    creator: (f = msg.getCreator()) && proto.operations_ecosys.User.toObject(includeInstance, f),
    creationDate: jspb.Message.getFieldWithDefault(msg, 4, ""),
    lastModifiedDate: jspb.Message.getFieldWithDefault(msg, 5, ""),
    lastModifedUser: (f = msg.getLastModifedUser()) && proto.operations_ecosys.User.toObject(includeInstance, f),
    isOriginal: jspb.Message.getBooleanFieldWithDefault(msg, 7, false),
    isApproved: jspb.Message.getBooleanFieldWithDefault(msg, 8, false),
    signature: (f = msg.getSignature()) && proto.operations_ecosys.User.toObject(includeInstance, f),
    approvalDate: jspb.Message.getFieldWithDefault(msg, 10, ""),
    incidentReportContent: (f = msg.getIncidentReportContent()) && proto.operations_ecosys.IncidentReportContent.toObject(includeInstance, f),
    aifsId: jspb.Message.getFieldWithDefault(msg, 12, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.IncidentReport}
 */
proto.operations_ecosys.IncidentReport.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.IncidentReport;
  return proto.operations_ecosys.IncidentReport.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.IncidentReport} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.IncidentReport}
 */
proto.operations_ecosys.IncidentReport.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setIncidentReportId(value);
      break;
    case 2:
      var value = /** @type {!proto.operations_ecosys.IncidentReport.ReportType} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 3:
      var value = new proto.operations_ecosys.User;
      reader.readMessage(value,proto.operations_ecosys.User.deserializeBinaryFromReader);
      msg.setCreator(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setCreationDate(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastModifiedDate(value);
      break;
    case 6:
      var value = new proto.operations_ecosys.User;
      reader.readMessage(value,proto.operations_ecosys.User.deserializeBinaryFromReader);
      msg.setLastModifedUser(value);
      break;
    case 7:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsOriginal(value);
      break;
    case 8:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsApproved(value);
      break;
    case 9:
      var value = new proto.operations_ecosys.User;
      reader.readMessage(value,proto.operations_ecosys.User.deserializeBinaryFromReader);
      msg.setSignature(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setApprovalDate(value);
      break;
    case 11:
      var value = new proto.operations_ecosys.IncidentReportContent;
      reader.readMessage(value,proto.operations_ecosys.IncidentReportContent.deserializeBinaryFromReader);
      msg.setIncidentReportContent(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setAifsId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.IncidentReport.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.IncidentReport.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.IncidentReport} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.IncidentReport.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getIncidentReportId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getCreator();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.operations_ecosys.User.serializeBinaryToWriter
    );
  }
  f = message.getCreationDate();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLastModifiedDate();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getLastModifedUser();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.operations_ecosys.User.serializeBinaryToWriter
    );
  }
  f = message.getIsOriginal();
  if (f) {
    writer.writeBool(
      7,
      f
    );
  }
  f = message.getIsApproved();
  if (f) {
    writer.writeBool(
      8,
      f
    );
  }
  f = message.getSignature();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      proto.operations_ecosys.User.serializeBinaryToWriter
    );
  }
  f = message.getApprovalDate();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getIncidentReportContent();
  if (f != null) {
    writer.writeMessage(
      11,
      f,
      proto.operations_ecosys.IncidentReportContent.serializeBinaryToWriter
    );
  }
  f = message.getAifsId();
  if (f !== 0) {
    writer.writeInt64(
      12,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.IncidentReport.ReportType = {
  FIRE_ALARM: 0,
  INTRUDER: 1,
  OTHERS: 2
};

/**
 * optional int64 incident_report_id = 1;
 * @return {number}
 */
proto.operations_ecosys.IncidentReport.prototype.getIncidentReportId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.setIncidentReportId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional ReportType type = 2;
 * @return {!proto.operations_ecosys.IncidentReport.ReportType}
 */
proto.operations_ecosys.IncidentReport.prototype.getType = function() {
  return /** @type {!proto.operations_ecosys.IncidentReport.ReportType} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.operations_ecosys.IncidentReport.ReportType} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.setType = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional User creator = 3;
 * @return {?proto.operations_ecosys.User}
 */
proto.operations_ecosys.IncidentReport.prototype.getCreator = function() {
  return /** @type{?proto.operations_ecosys.User} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.User, 3));
};


/**
 * @param {?proto.operations_ecosys.User|undefined} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
*/
proto.operations_ecosys.IncidentReport.prototype.setCreator = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.clearCreator = function() {
  return this.setCreator(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReport.prototype.hasCreator = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional string creation_date = 4;
 * @return {string}
 */
proto.operations_ecosys.IncidentReport.prototype.getCreationDate = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.setCreationDate = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string last_modified_date = 5;
 * @return {string}
 */
proto.operations_ecosys.IncidentReport.prototype.getLastModifiedDate = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.setLastModifiedDate = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional User last_modifed_user = 6;
 * @return {?proto.operations_ecosys.User}
 */
proto.operations_ecosys.IncidentReport.prototype.getLastModifedUser = function() {
  return /** @type{?proto.operations_ecosys.User} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.User, 6));
};


/**
 * @param {?proto.operations_ecosys.User|undefined} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
*/
proto.operations_ecosys.IncidentReport.prototype.setLastModifedUser = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.clearLastModifedUser = function() {
  return this.setLastModifedUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReport.prototype.hasLastModifedUser = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional bool is_original = 7;
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReport.prototype.getIsOriginal = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 7, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.setIsOriginal = function(value) {
  return jspb.Message.setProto3BooleanField(this, 7, value);
};


/**
 * optional bool is_approved = 8;
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReport.prototype.getIsApproved = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 8, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.setIsApproved = function(value) {
  return jspb.Message.setProto3BooleanField(this, 8, value);
};


/**
 * optional User signature = 9;
 * @return {?proto.operations_ecosys.User}
 */
proto.operations_ecosys.IncidentReport.prototype.getSignature = function() {
  return /** @type{?proto.operations_ecosys.User} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.User, 9));
};


/**
 * @param {?proto.operations_ecosys.User|undefined} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
*/
proto.operations_ecosys.IncidentReport.prototype.setSignature = function(value) {
  return jspb.Message.setWrapperField(this, 9, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.clearSignature = function() {
  return this.setSignature(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReport.prototype.hasSignature = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional string approval_date = 10;
 * @return {string}
 */
proto.operations_ecosys.IncidentReport.prototype.getApprovalDate = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.setApprovalDate = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional IncidentReportContent incident_report_content = 11;
 * @return {?proto.operations_ecosys.IncidentReportContent}
 */
proto.operations_ecosys.IncidentReport.prototype.getIncidentReportContent = function() {
  return /** @type{?proto.operations_ecosys.IncidentReportContent} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.IncidentReportContent, 11));
};


/**
 * @param {?proto.operations_ecosys.IncidentReportContent|undefined} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
*/
proto.operations_ecosys.IncidentReport.prototype.setIncidentReportContent = function(value) {
  return jspb.Message.setWrapperField(this, 11, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.clearIncidentReportContent = function() {
  return this.setIncidentReportContent(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReport.prototype.hasIncidentReportContent = function() {
  return jspb.Message.getField(this, 11) != null;
};


/**
 * optional int64 aifs_id = 12;
 * @return {number}
 */
proto.operations_ecosys.IncidentReport.prototype.getAifsId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.IncidentReport} returns this
 */
proto.operations_ecosys.IncidentReport.prototype.setAifsId = function(value) {
  return jspb.Message.setProto3IntField(this, 12, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.IncidentReportContent.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.IncidentReportContent.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.IncidentReportContent} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.IncidentReportContent.toObject = function(includeInstance, msg) {
  var f, obj = {
    reportContentId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    lastModifiedDate: jspb.Message.getFieldWithDefault(msg, 2, ""),
    lastModifedUser: (f = msg.getLastModifedUser()) && proto.operations_ecosys.User.toObject(includeInstance, f),
    address: jspb.Message.getFieldWithDefault(msg, 4, ""),
    incidentTime: jspb.Message.getFieldWithDefault(msg, 5, ""),
    isPoliceNotified: jspb.Message.getBooleanFieldWithDefault(msg, 6, false),
    title: jspb.Message.getFieldWithDefault(msg, 7, ""),
    description: jspb.Message.getFieldWithDefault(msg, 8, ""),
    hasActionTaken: jspb.Message.getBooleanFieldWithDefault(msg, 9, false),
    actionTaken: jspb.Message.getFieldWithDefault(msg, 10, ""),
    hasInjury: jspb.Message.getBooleanFieldWithDefault(msg, 11, false),
    injuryDescription: jspb.Message.getFieldWithDefault(msg, 12, ""),
    hasStolenItem: jspb.Message.getBooleanFieldWithDefault(msg, 13, false),
    stolenItemDescription: jspb.Message.getFieldWithDefault(msg, 14, ""),
    reportImageLink: jspb.Message.getFieldWithDefault(msg, 15, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.IncidentReportContent}
 */
proto.operations_ecosys.IncidentReportContent.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.IncidentReportContent;
  return proto.operations_ecosys.IncidentReportContent.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.IncidentReportContent} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.IncidentReportContent}
 */
proto.operations_ecosys.IncidentReportContent.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setReportContentId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastModifiedDate(value);
      break;
    case 3:
      var value = new proto.operations_ecosys.User;
      reader.readMessage(value,proto.operations_ecosys.User.deserializeBinaryFromReader);
      msg.setLastModifedUser(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setAddress(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setIncidentTime(value);
      break;
    case 6:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsPoliceNotified(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setTitle(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    case 9:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setHasActionTaken(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setActionTaken(value);
      break;
    case 11:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setHasInjury(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.setInjuryDescription(value);
      break;
    case 13:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setHasStolenItem(value);
      break;
    case 14:
      var value = /** @type {string} */ (reader.readString());
      msg.setStolenItemDescription(value);
      break;
    case 15:
      var value = /** @type {string} */ (reader.readString());
      msg.setReportImageLink(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.IncidentReportContent.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.IncidentReportContent.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.IncidentReportContent} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.IncidentReportContent.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getReportContentId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getLastModifiedDate();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getLastModifedUser();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.operations_ecosys.User.serializeBinaryToWriter
    );
  }
  f = message.getAddress();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getIncidentTime();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getIsPoliceNotified();
  if (f) {
    writer.writeBool(
      6,
      f
    );
  }
  f = message.getTitle();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getDescription();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getHasActionTaken();
  if (f) {
    writer.writeBool(
      9,
      f
    );
  }
  f = message.getActionTaken();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getHasInjury();
  if (f) {
    writer.writeBool(
      11,
      f
    );
  }
  f = message.getInjuryDescription();
  if (f.length > 0) {
    writer.writeString(
      12,
      f
    );
  }
  f = message.getHasStolenItem();
  if (f) {
    writer.writeBool(
      13,
      f
    );
  }
  f = message.getStolenItemDescription();
  if (f.length > 0) {
    writer.writeString(
      14,
      f
    );
  }
  f = message.getReportImageLink();
  if (f.length > 0) {
    writer.writeString(
      15,
      f
    );
  }
};


/**
 * optional int64 report_content_id = 1;
 * @return {number}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getReportContentId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setReportContentId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string last_modified_date = 2;
 * @return {string}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getLastModifiedDate = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setLastModifiedDate = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional User last_modifed_user = 3;
 * @return {?proto.operations_ecosys.User}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getLastModifedUser = function() {
  return /** @type{?proto.operations_ecosys.User} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.User, 3));
};


/**
 * @param {?proto.operations_ecosys.User|undefined} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
*/
proto.operations_ecosys.IncidentReportContent.prototype.setLastModifedUser = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.clearLastModifedUser = function() {
  return this.setLastModifedUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReportContent.prototype.hasLastModifedUser = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional string address = 4;
 * @return {string}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getAddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setAddress = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string incident_time = 5;
 * @return {string}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getIncidentTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setIncidentTime = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional bool is_police_notified = 6;
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getIsPoliceNotified = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 6, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setIsPoliceNotified = function(value) {
  return jspb.Message.setProto3BooleanField(this, 6, value);
};


/**
 * optional string title = 7;
 * @return {string}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getTitle = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setTitle = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string description = 8;
 * @return {string}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setDescription = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional bool has_action_taken = 9;
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getHasActionTaken = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 9, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setHasActionTaken = function(value) {
  return jspb.Message.setProto3BooleanField(this, 9, value);
};


/**
 * optional string action_taken = 10;
 * @return {string}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getActionTaken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setActionTaken = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional bool has_injury = 11;
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getHasInjury = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 11, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setHasInjury = function(value) {
  return jspb.Message.setProto3BooleanField(this, 11, value);
};


/**
 * optional string injury_description = 12;
 * @return {string}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getInjuryDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 12, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setInjuryDescription = function(value) {
  return jspb.Message.setProto3StringField(this, 12, value);
};


/**
 * optional bool has_stolen_item = 13;
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getHasStolenItem = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 13, false));
};


/**
 * @param {boolean} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setHasStolenItem = function(value) {
  return jspb.Message.setProto3BooleanField(this, 13, value);
};


/**
 * optional string stolen_item_description = 14;
 * @return {string}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getStolenItemDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 14, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setStolenItemDescription = function(value) {
  return jspb.Message.setProto3StringField(this, 14, value);
};


/**
 * optional string report_image_link = 15;
 * @return {string}
 */
proto.operations_ecosys.IncidentReportContent.prototype.getReportImageLink = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 15, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.IncidentReportContent} returns this
 */
proto.operations_ecosys.IncidentReportContent.prototype.setReportImageLink = function(value) {
  return jspb.Message.setProto3StringField(this, 15, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.IncidentReportResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.IncidentReportResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.IncidentReportResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.IncidentReportResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    response: (f = msg.getResponse()) && proto.operations_ecosys.Response.toObject(includeInstance, f),
    incidentReport: (f = msg.getIncidentReport()) && proto.operations_ecosys.IncidentReport.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.IncidentReportResponse}
 */
proto.operations_ecosys.IncidentReportResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.IncidentReportResponse;
  return proto.operations_ecosys.IncidentReportResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.IncidentReportResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.IncidentReportResponse}
 */
proto.operations_ecosys.IncidentReportResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.Response;
      reader.readMessage(value,proto.operations_ecosys.Response.deserializeBinaryFromReader);
      msg.setResponse(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.IncidentReport;
      reader.readMessage(value,proto.operations_ecosys.IncidentReport.deserializeBinaryFromReader);
      msg.setIncidentReport(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.IncidentReportResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.IncidentReportResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.IncidentReportResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.IncidentReportResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResponse();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.operations_ecosys.Response.serializeBinaryToWriter
    );
  }
  f = message.getIncidentReport();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.IncidentReport.serializeBinaryToWriter
    );
  }
};


/**
 * optional Response response = 1;
 * @return {?proto.operations_ecosys.Response}
 */
proto.operations_ecosys.IncidentReportResponse.prototype.getResponse = function() {
  return /** @type{?proto.operations_ecosys.Response} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Response, 1));
};


/**
 * @param {?proto.operations_ecosys.Response|undefined} value
 * @return {!proto.operations_ecosys.IncidentReportResponse} returns this
*/
proto.operations_ecosys.IncidentReportResponse.prototype.setResponse = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.IncidentReportResponse} returns this
 */
proto.operations_ecosys.IncidentReportResponse.prototype.clearResponse = function() {
  return this.setResponse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReportResponse.prototype.hasResponse = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional IncidentReport incident_report = 2;
 * @return {?proto.operations_ecosys.IncidentReport}
 */
proto.operations_ecosys.IncidentReportResponse.prototype.getIncidentReport = function() {
  return /** @type{?proto.operations_ecosys.IncidentReport} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.IncidentReport, 2));
};


/**
 * @param {?proto.operations_ecosys.IncidentReport|undefined} value
 * @return {!proto.operations_ecosys.IncidentReportResponse} returns this
*/
proto.operations_ecosys.IncidentReportResponse.prototype.setIncidentReport = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.IncidentReportResponse} returns this
 */
proto.operations_ecosys.IncidentReportResponse.prototype.clearIncidentReport = function() {
  return this.setIncidentReport(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReportResponse.prototype.hasIncidentReport = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.IncidentReportFilter.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.IncidentReportFilter.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.IncidentReportFilter} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.IncidentReportFilter.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    comparisons: (f = msg.getComparisons()) && proto.operations_ecosys.Filter.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.IncidentReportFilter}
 */
proto.operations_ecosys.IncidentReportFilter.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.IncidentReportFilter;
  return proto.operations_ecosys.IncidentReportFilter.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.IncidentReportFilter} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.IncidentReportFilter}
 */
proto.operations_ecosys.IncidentReportFilter.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.IncidentReportFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Filter;
      reader.readMessage(value,proto.operations_ecosys.Filter.deserializeBinaryFromReader);
      msg.setComparisons(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.IncidentReportFilter.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.IncidentReportFilter.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.IncidentReportFilter} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.IncidentReportFilter.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getComparisons();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Filter.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.IncidentReportFilter.Field = {
  REPORT_ID: 0,
  REPORT_CONTENT_ID: 1,
  REPORT_TYPE: 2,
  MODIFIER: 3,
  LAST_MODIFIED_DATE: 5,
  GET_ORIGINAL: 6,
  IS_APPROVED: 7,
  SIGNATURE: 8,
  APPROVAL_DATE: 9
};

/**
 * optional Field field = 1;
 * @return {!proto.operations_ecosys.IncidentReportFilter.Field}
 */
proto.operations_ecosys.IncidentReportFilter.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.IncidentReportFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.IncidentReportFilter.Field} value
 * @return {!proto.operations_ecosys.IncidentReportFilter} returns this
 */
proto.operations_ecosys.IncidentReportFilter.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional Filter comparisons = 2;
 * @return {?proto.operations_ecosys.Filter}
 */
proto.operations_ecosys.IncidentReportFilter.prototype.getComparisons = function() {
  return /** @type{?proto.operations_ecosys.Filter} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Filter, 2));
};


/**
 * @param {?proto.operations_ecosys.Filter|undefined} value
 * @return {!proto.operations_ecosys.IncidentReportFilter} returns this
*/
proto.operations_ecosys.IncidentReportFilter.prototype.setComparisons = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.IncidentReportFilter} returns this
 */
proto.operations_ecosys.IncidentReportFilter.prototype.clearComparisons = function() {
  return this.setComparisons(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReportFilter.prototype.hasComparisons = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.IncidentReportQuery.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.IncidentReportQuery.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.IncidentReportQuery.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.IncidentReportQuery} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.IncidentReportQuery.toObject = function(includeInstance, msg) {
  var f, obj = {
    filtersList: jspb.Message.toObjectList(msg.getFiltersList(),
    proto.operations_ecosys.IncidentReportFilter.toObject, includeInstance),
    limit: jspb.Message.getFieldWithDefault(msg, 2, 0),
    skip: jspb.Message.getFieldWithDefault(msg, 3, 0),
    orderBy: (f = msg.getOrderBy()) && proto.operations_ecosys.OrderByIncidentReport.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.IncidentReportQuery}
 */
proto.operations_ecosys.IncidentReportQuery.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.IncidentReportQuery;
  return proto.operations_ecosys.IncidentReportQuery.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.IncidentReportQuery} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.IncidentReportQuery}
 */
proto.operations_ecosys.IncidentReportQuery.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.IncidentReportFilter;
      reader.readMessage(value,proto.operations_ecosys.IncidentReportFilter.deserializeBinaryFromReader);
      msg.addFilters(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLimit(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSkip(value);
      break;
    case 4:
      var value = new proto.operations_ecosys.OrderByIncidentReport;
      reader.readMessage(value,proto.operations_ecosys.OrderByIncidentReport.deserializeBinaryFromReader);
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.IncidentReportQuery.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.IncidentReportQuery.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.IncidentReportQuery} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.IncidentReportQuery.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFiltersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.operations_ecosys.IncidentReportFilter.serializeBinaryToWriter
    );
  }
  f = message.getLimit();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getSkip();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getOrderBy();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.operations_ecosys.OrderByIncidentReport.serializeBinaryToWriter
    );
  }
};


/**
 * repeated IncidentReportFilter filters = 1;
 * @return {!Array<!proto.operations_ecosys.IncidentReportFilter>}
 */
proto.operations_ecosys.IncidentReportQuery.prototype.getFiltersList = function() {
  return /** @type{!Array<!proto.operations_ecosys.IncidentReportFilter>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.IncidentReportFilter, 1));
};


/**
 * @param {!Array<!proto.operations_ecosys.IncidentReportFilter>} value
 * @return {!proto.operations_ecosys.IncidentReportQuery} returns this
*/
proto.operations_ecosys.IncidentReportQuery.prototype.setFiltersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.operations_ecosys.IncidentReportFilter=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.IncidentReportFilter}
 */
proto.operations_ecosys.IncidentReportQuery.prototype.addFilters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.operations_ecosys.IncidentReportFilter, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.IncidentReportQuery} returns this
 */
proto.operations_ecosys.IncidentReportQuery.prototype.clearFiltersList = function() {
  return this.setFiltersList([]);
};


/**
 * optional int64 limit = 2;
 * @return {number}
 */
proto.operations_ecosys.IncidentReportQuery.prototype.getLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.IncidentReportQuery} returns this
 */
proto.operations_ecosys.IncidentReportQuery.prototype.setLimit = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 skip = 3;
 * @return {number}
 */
proto.operations_ecosys.IncidentReportQuery.prototype.getSkip = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.IncidentReportQuery} returns this
 */
proto.operations_ecosys.IncidentReportQuery.prototype.setSkip = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional OrderByIncidentReport order_by = 4;
 * @return {?proto.operations_ecosys.OrderByIncidentReport}
 */
proto.operations_ecosys.IncidentReportQuery.prototype.getOrderBy = function() {
  return /** @type{?proto.operations_ecosys.OrderByIncidentReport} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.OrderByIncidentReport, 4));
};


/**
 * @param {?proto.operations_ecosys.OrderByIncidentReport|undefined} value
 * @return {!proto.operations_ecosys.IncidentReportQuery} returns this
*/
proto.operations_ecosys.IncidentReportQuery.prototype.setOrderBy = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.IncidentReportQuery} returns this
 */
proto.operations_ecosys.IncidentReportQuery.prototype.clearOrderBy = function() {
  return this.setOrderBy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.IncidentReportQuery.prototype.hasOrderBy = function() {
  return jspb.Message.getField(this, 4) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.OrderByIncidentReport.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.OrderByIncidentReport.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.OrderByIncidentReport} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByIncidentReport.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    orderBy: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.OrderByIncidentReport}
 */
proto.operations_ecosys.OrderByIncidentReport.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.OrderByIncidentReport;
  return proto.operations_ecosys.OrderByIncidentReport.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.OrderByIncidentReport} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.OrderByIncidentReport}
 */
proto.operations_ecosys.OrderByIncidentReport.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.IncidentReportFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = /** @type {!proto.operations_ecosys.OrderBy} */ (reader.readEnum());
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.OrderByIncidentReport.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.OrderByIncidentReport.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.OrderByIncidentReport} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByIncidentReport.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getOrderBy();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * optional IncidentReportFilter.Field field = 1;
 * @return {!proto.operations_ecosys.IncidentReportFilter.Field}
 */
proto.operations_ecosys.OrderByIncidentReport.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.IncidentReportFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.IncidentReportFilter.Field} value
 * @return {!proto.operations_ecosys.OrderByIncidentReport} returns this
 */
proto.operations_ecosys.OrderByIncidentReport.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional OrderBy order_by = 2;
 * @return {!proto.operations_ecosys.OrderBy}
 */
proto.operations_ecosys.OrderByIncidentReport.prototype.getOrderBy = function() {
  return /** @type {!proto.operations_ecosys.OrderBy} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.operations_ecosys.OrderBy} value
 * @return {!proto.operations_ecosys.OrderByIncidentReport} returns this
 */
proto.operations_ecosys.OrderByIncidentReport.prototype.setOrderBy = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.CameraIot.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.CameraIot.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.CameraIot} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.CameraIot.toObject = function(includeInstance, msg) {
  var f, obj = {
    cameraIotId: jspb.Message.getFieldWithDefault(msg, 1, 0),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    camera: (f = msg.getCamera()) && proto.operations_ecosys.Camera.toObject(includeInstance, f),
    gate: (f = msg.getGate()) && iot_prototype_pb.GateState.toObject(includeInstance, f),
    fireAlarm: (f = msg.getFireAlarm()) && iot_prototype_pb.FireAlarmState.toObject(includeInstance, f),
    cpuTemperature: (f = msg.getCpuTemperature()) && iot_prototype_pb.CpuTempState.toObject(includeInstance, f),
    type: jspb.Message.getFieldWithDefault(msg, 7, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.CameraIot}
 */
proto.operations_ecosys.CameraIot.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.CameraIot;
  return proto.operations_ecosys.CameraIot.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.CameraIot} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.CameraIot}
 */
proto.operations_ecosys.CameraIot.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCameraIotId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = new proto.operations_ecosys.Camera;
      reader.readMessage(value,proto.operations_ecosys.Camera.deserializeBinaryFromReader);
      msg.setCamera(value);
      break;
    case 4:
      var value = new iot_prototype_pb.GateState;
      reader.readMessage(value,iot_prototype_pb.GateState.deserializeBinaryFromReader);
      msg.setGate(value);
      break;
    case 5:
      var value = new iot_prototype_pb.FireAlarmState;
      reader.readMessage(value,iot_prototype_pb.FireAlarmState.deserializeBinaryFromReader);
      msg.setFireAlarm(value);
      break;
    case 6:
      var value = new iot_prototype_pb.CpuTempState;
      reader.readMessage(value,iot_prototype_pb.CpuTempState.deserializeBinaryFromReader);
      msg.setCpuTemperature(value);
      break;
    case 7:
      var value = /** @type {!proto.operations_ecosys.CameraIot.MessageType} */ (reader.readEnum());
      msg.setType(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.CameraIot.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.CameraIot.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.CameraIot} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.CameraIot.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCameraIotId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getCamera();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.operations_ecosys.Camera.serializeBinaryToWriter
    );
  }
  f = message.getGate();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      iot_prototype_pb.GateState.serializeBinaryToWriter
    );
  }
  f = message.getFireAlarm();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      iot_prototype_pb.FireAlarmState.serializeBinaryToWriter
    );
  }
  f = message.getCpuTemperature();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      iot_prototype_pb.CpuTempState.serializeBinaryToWriter
    );
  }
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      7,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.CameraIot.MessageType = {
  INITIAL: 0,
  CHANGE_GATE: 1,
  CHANGE_FIRE_ALARM: 2,
  CHANGE_CPU_TEMP: 3
};

/**
 * optional int64 camera_iot_id = 1;
 * @return {number}
 */
proto.operations_ecosys.CameraIot.prototype.getCameraIotId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.CameraIot} returns this
 */
proto.operations_ecosys.CameraIot.prototype.setCameraIotId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.operations_ecosys.CameraIot.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.CameraIot} returns this
 */
proto.operations_ecosys.CameraIot.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Camera camera = 3;
 * @return {?proto.operations_ecosys.Camera}
 */
proto.operations_ecosys.CameraIot.prototype.getCamera = function() {
  return /** @type{?proto.operations_ecosys.Camera} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Camera, 3));
};


/**
 * @param {?proto.operations_ecosys.Camera|undefined} value
 * @return {!proto.operations_ecosys.CameraIot} returns this
*/
proto.operations_ecosys.CameraIot.prototype.setCamera = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.CameraIot} returns this
 */
proto.operations_ecosys.CameraIot.prototype.clearCamera = function() {
  return this.setCamera(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.CameraIot.prototype.hasCamera = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional gate_prototype.GateState gate = 4;
 * @return {?proto.gate_prototype.GateState}
 */
proto.operations_ecosys.CameraIot.prototype.getGate = function() {
  return /** @type{?proto.gate_prototype.GateState} */ (
    jspb.Message.getWrapperField(this, iot_prototype_pb.GateState, 4));
};


/**
 * @param {?proto.gate_prototype.GateState|undefined} value
 * @return {!proto.operations_ecosys.CameraIot} returns this
*/
proto.operations_ecosys.CameraIot.prototype.setGate = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.CameraIot} returns this
 */
proto.operations_ecosys.CameraIot.prototype.clearGate = function() {
  return this.setGate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.CameraIot.prototype.hasGate = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional gate_prototype.FireAlarmState fire_alarm = 5;
 * @return {?proto.gate_prototype.FireAlarmState}
 */
proto.operations_ecosys.CameraIot.prototype.getFireAlarm = function() {
  return /** @type{?proto.gate_prototype.FireAlarmState} */ (
    jspb.Message.getWrapperField(this, iot_prototype_pb.FireAlarmState, 5));
};


/**
 * @param {?proto.gate_prototype.FireAlarmState|undefined} value
 * @return {!proto.operations_ecosys.CameraIot} returns this
*/
proto.operations_ecosys.CameraIot.prototype.setFireAlarm = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.CameraIot} returns this
 */
proto.operations_ecosys.CameraIot.prototype.clearFireAlarm = function() {
  return this.setFireAlarm(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.CameraIot.prototype.hasFireAlarm = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional gate_prototype.CpuTempState cpu_temperature = 6;
 * @return {?proto.gate_prototype.CpuTempState}
 */
proto.operations_ecosys.CameraIot.prototype.getCpuTemperature = function() {
  return /** @type{?proto.gate_prototype.CpuTempState} */ (
    jspb.Message.getWrapperField(this, iot_prototype_pb.CpuTempState, 6));
};


/**
 * @param {?proto.gate_prototype.CpuTempState|undefined} value
 * @return {!proto.operations_ecosys.CameraIot} returns this
*/
proto.operations_ecosys.CameraIot.prototype.setCpuTemperature = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.CameraIot} returns this
 */
proto.operations_ecosys.CameraIot.prototype.clearCpuTemperature = function() {
  return this.setCpuTemperature(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.CameraIot.prototype.hasCpuTemperature = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional MessageType type = 7;
 * @return {!proto.operations_ecosys.CameraIot.MessageType}
 */
proto.operations_ecosys.CameraIot.prototype.getType = function() {
  return /** @type {!proto.operations_ecosys.CameraIot.MessageType} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {!proto.operations_ecosys.CameraIot.MessageType} value
 * @return {!proto.operations_ecosys.CameraIot} returns this
 */
proto.operations_ecosys.CameraIot.prototype.setType = function(value) {
  return jspb.Message.setProto3EnumField(this, 7, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.Camera.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.Camera.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.Camera} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Camera.toObject = function(includeInstance, msg) {
  var f, obj = {
    url: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.Camera}
 */
proto.operations_ecosys.Camera.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.Camera;
  return proto.operations_ecosys.Camera.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.Camera} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.Camera}
 */
proto.operations_ecosys.Camera.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUrl(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.Camera.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.Camera.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.Camera} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Camera.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUrl();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string url = 1;
 * @return {string}
 */
proto.operations_ecosys.Camera.prototype.getUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Camera} returns this
 */
proto.operations_ecosys.Camera.prototype.setUrl = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.CameraIotResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.CameraIotResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.CameraIotResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.CameraIotResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    response: (f = msg.getResponse()) && proto.operations_ecosys.Response.toObject(includeInstance, f),
    cameraIot: (f = msg.getCameraIot()) && proto.operations_ecosys.CameraIot.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.CameraIotResponse}
 */
proto.operations_ecosys.CameraIotResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.CameraIotResponse;
  return proto.operations_ecosys.CameraIotResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.CameraIotResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.CameraIotResponse}
 */
proto.operations_ecosys.CameraIotResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.Response;
      reader.readMessage(value,proto.operations_ecosys.Response.deserializeBinaryFromReader);
      msg.setResponse(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.CameraIot;
      reader.readMessage(value,proto.operations_ecosys.CameraIot.deserializeBinaryFromReader);
      msg.setCameraIot(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.CameraIotResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.CameraIotResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.CameraIotResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.CameraIotResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResponse();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.operations_ecosys.Response.serializeBinaryToWriter
    );
  }
  f = message.getCameraIot();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.CameraIot.serializeBinaryToWriter
    );
  }
};


/**
 * optional Response response = 1;
 * @return {?proto.operations_ecosys.Response}
 */
proto.operations_ecosys.CameraIotResponse.prototype.getResponse = function() {
  return /** @type{?proto.operations_ecosys.Response} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Response, 1));
};


/**
 * @param {?proto.operations_ecosys.Response|undefined} value
 * @return {!proto.operations_ecosys.CameraIotResponse} returns this
*/
proto.operations_ecosys.CameraIotResponse.prototype.setResponse = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.CameraIotResponse} returns this
 */
proto.operations_ecosys.CameraIotResponse.prototype.clearResponse = function() {
  return this.setResponse(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.CameraIotResponse.prototype.hasResponse = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional CameraIot camera_iot = 2;
 * @return {?proto.operations_ecosys.CameraIot}
 */
proto.operations_ecosys.CameraIotResponse.prototype.getCameraIot = function() {
  return /** @type{?proto.operations_ecosys.CameraIot} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.CameraIot, 2));
};


/**
 * @param {?proto.operations_ecosys.CameraIot|undefined} value
 * @return {!proto.operations_ecosys.CameraIotResponse} returns this
*/
proto.operations_ecosys.CameraIotResponse.prototype.setCameraIot = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.CameraIotResponse} returns this
 */
proto.operations_ecosys.CameraIotResponse.prototype.clearCameraIot = function() {
  return this.setCameraIot(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.CameraIotResponse.prototype.hasCameraIot = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.CameraIotFilter.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.CameraIotFilter.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.CameraIotFilter} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.CameraIotFilter.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    comparisons: (f = msg.getComparisons()) && proto.operations_ecosys.Filter.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.CameraIotFilter}
 */
proto.operations_ecosys.CameraIotFilter.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.CameraIotFilter;
  return proto.operations_ecosys.CameraIotFilter.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.CameraIotFilter} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.CameraIotFilter}
 */
proto.operations_ecosys.CameraIotFilter.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.CameraIotFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = new proto.operations_ecosys.Filter;
      reader.readMessage(value,proto.operations_ecosys.Filter.deserializeBinaryFromReader);
      msg.setComparisons(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.CameraIotFilter.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.CameraIotFilter.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.CameraIotFilter} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.CameraIotFilter.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getComparisons();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.operations_ecosys.Filter.serializeBinaryToWriter
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.CameraIotFilter.Field = {
  CAMERA_IOT_ID: 0
};

/**
 * optional Field field = 1;
 * @return {!proto.operations_ecosys.CameraIotFilter.Field}
 */
proto.operations_ecosys.CameraIotFilter.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.CameraIotFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.CameraIotFilter.Field} value
 * @return {!proto.operations_ecosys.CameraIotFilter} returns this
 */
proto.operations_ecosys.CameraIotFilter.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional Filter comparisons = 2;
 * @return {?proto.operations_ecosys.Filter}
 */
proto.operations_ecosys.CameraIotFilter.prototype.getComparisons = function() {
  return /** @type{?proto.operations_ecosys.Filter} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.Filter, 2));
};


/**
 * @param {?proto.operations_ecosys.Filter|undefined} value
 * @return {!proto.operations_ecosys.CameraIotFilter} returns this
*/
proto.operations_ecosys.CameraIotFilter.prototype.setComparisons = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.CameraIotFilter} returns this
 */
proto.operations_ecosys.CameraIotFilter.prototype.clearComparisons = function() {
  return this.setComparisons(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.CameraIotFilter.prototype.hasComparisons = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.operations_ecosys.CameraIotQuery.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.CameraIotQuery.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.CameraIotQuery.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.CameraIotQuery} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.CameraIotQuery.toObject = function(includeInstance, msg) {
  var f, obj = {
    filtersList: jspb.Message.toObjectList(msg.getFiltersList(),
    proto.operations_ecosys.CameraIotFilter.toObject, includeInstance),
    limit: jspb.Message.getFieldWithDefault(msg, 2, 0),
    skip: jspb.Message.getFieldWithDefault(msg, 3, 0),
    orderBy: (f = msg.getOrderBy()) && proto.operations_ecosys.OrderByCameraIot.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.CameraIotQuery}
 */
proto.operations_ecosys.CameraIotQuery.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.CameraIotQuery;
  return proto.operations_ecosys.CameraIotQuery.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.CameraIotQuery} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.CameraIotQuery}
 */
proto.operations_ecosys.CameraIotQuery.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.operations_ecosys.CameraIotFilter;
      reader.readMessage(value,proto.operations_ecosys.CameraIotFilter.deserializeBinaryFromReader);
      msg.addFilters(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLimit(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSkip(value);
      break;
    case 4:
      var value = new proto.operations_ecosys.OrderByCameraIot;
      reader.readMessage(value,proto.operations_ecosys.OrderByCameraIot.deserializeBinaryFromReader);
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.CameraIotQuery.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.CameraIotQuery.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.CameraIotQuery} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.CameraIotQuery.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFiltersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.operations_ecosys.CameraIotFilter.serializeBinaryToWriter
    );
  }
  f = message.getLimit();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getSkip();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getOrderBy();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.operations_ecosys.OrderByCameraIot.serializeBinaryToWriter
    );
  }
};


/**
 * repeated CameraIotFilter filters = 1;
 * @return {!Array<!proto.operations_ecosys.CameraIotFilter>}
 */
proto.operations_ecosys.CameraIotQuery.prototype.getFiltersList = function() {
  return /** @type{!Array<!proto.operations_ecosys.CameraIotFilter>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.operations_ecosys.CameraIotFilter, 1));
};


/**
 * @param {!Array<!proto.operations_ecosys.CameraIotFilter>} value
 * @return {!proto.operations_ecosys.CameraIotQuery} returns this
*/
proto.operations_ecosys.CameraIotQuery.prototype.setFiltersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.operations_ecosys.CameraIotFilter=} opt_value
 * @param {number=} opt_index
 * @return {!proto.operations_ecosys.CameraIotFilter}
 */
proto.operations_ecosys.CameraIotQuery.prototype.addFilters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.operations_ecosys.CameraIotFilter, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.operations_ecosys.CameraIotQuery} returns this
 */
proto.operations_ecosys.CameraIotQuery.prototype.clearFiltersList = function() {
  return this.setFiltersList([]);
};


/**
 * optional int64 limit = 2;
 * @return {number}
 */
proto.operations_ecosys.CameraIotQuery.prototype.getLimit = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.CameraIotQuery} returns this
 */
proto.operations_ecosys.CameraIotQuery.prototype.setLimit = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 skip = 3;
 * @return {number}
 */
proto.operations_ecosys.CameraIotQuery.prototype.getSkip = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.CameraIotQuery} returns this
 */
proto.operations_ecosys.CameraIotQuery.prototype.setSkip = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional OrderByCameraIot order_by = 4;
 * @return {?proto.operations_ecosys.OrderByCameraIot}
 */
proto.operations_ecosys.CameraIotQuery.prototype.getOrderBy = function() {
  return /** @type{?proto.operations_ecosys.OrderByCameraIot} */ (
    jspb.Message.getWrapperField(this, proto.operations_ecosys.OrderByCameraIot, 4));
};


/**
 * @param {?proto.operations_ecosys.OrderByCameraIot|undefined} value
 * @return {!proto.operations_ecosys.CameraIotQuery} returns this
*/
proto.operations_ecosys.CameraIotQuery.prototype.setOrderBy = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.operations_ecosys.CameraIotQuery} returns this
 */
proto.operations_ecosys.CameraIotQuery.prototype.clearOrderBy = function() {
  return this.setOrderBy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.operations_ecosys.CameraIotQuery.prototype.hasOrderBy = function() {
  return jspb.Message.getField(this, 4) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.OrderByCameraIot.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.OrderByCameraIot.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.OrderByCameraIot} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByCameraIot.toObject = function(includeInstance, msg) {
  var f, obj = {
    field: jspb.Message.getFieldWithDefault(msg, 1, 0),
    orderBy: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.OrderByCameraIot}
 */
proto.operations_ecosys.OrderByCameraIot.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.OrderByCameraIot;
  return proto.operations_ecosys.OrderByCameraIot.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.OrderByCameraIot} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.OrderByCameraIot}
 */
proto.operations_ecosys.OrderByCameraIot.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.CameraIotFilter.Field} */ (reader.readEnum());
      msg.setField(value);
      break;
    case 2:
      var value = /** @type {!proto.operations_ecosys.OrderBy} */ (reader.readEnum());
      msg.setOrderBy(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.OrderByCameraIot.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.OrderByCameraIot.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.OrderByCameraIot} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.OrderByCameraIot.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getOrderBy();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * optional CameraIotFilter.Field field = 1;
 * @return {!proto.operations_ecosys.CameraIotFilter.Field}
 */
proto.operations_ecosys.OrderByCameraIot.prototype.getField = function() {
  return /** @type {!proto.operations_ecosys.CameraIotFilter.Field} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.CameraIotFilter.Field} value
 * @return {!proto.operations_ecosys.OrderByCameraIot} returns this
 */
proto.operations_ecosys.OrderByCameraIot.prototype.setField = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional OrderBy order_by = 2;
 * @return {!proto.operations_ecosys.OrderBy}
 */
proto.operations_ecosys.OrderByCameraIot.prototype.getOrderBy = function() {
  return /** @type {!proto.operations_ecosys.OrderBy} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.operations_ecosys.OrderBy} value
 * @return {!proto.operations_ecosys.OrderByCameraIot} returns this
 */
proto.operations_ecosys.OrderByCameraIot.prototype.setOrderBy = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.Response.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.Response.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.Response} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Response.toObject = function(includeInstance, msg) {
  var f, obj = {
    type: jspb.Message.getFieldWithDefault(msg, 1, 0),
    errorMessage: jspb.Message.getFieldWithDefault(msg, 2, ""),
    primaryKey: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.Response}
 */
proto.operations_ecosys.Response.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.Response;
  return proto.operations_ecosys.Response.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.Response} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.Response}
 */
proto.operations_ecosys.Response.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.Response.Type} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setErrorMessage(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPrimaryKey(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.Response.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.Response.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.Response} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Response.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getErrorMessage();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getPrimaryKey();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.Response.Type = {
  ACK: 0,
  ERROR: 1
};

/**
 * optional Type type = 1;
 * @return {!proto.operations_ecosys.Response.Type}
 */
proto.operations_ecosys.Response.prototype.getType = function() {
  return /** @type {!proto.operations_ecosys.Response.Type} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.Response.Type} value
 * @return {!proto.operations_ecosys.Response} returns this
 */
proto.operations_ecosys.Response.prototype.setType = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string error_message = 2;
 * @return {string}
 */
proto.operations_ecosys.Response.prototype.getErrorMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Response} returns this
 */
proto.operations_ecosys.Response.prototype.setErrorMessage = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int64 primary_key = 3;
 * @return {number}
 */
proto.operations_ecosys.Response.prototype.getPrimaryKey = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.operations_ecosys.Response} returns this
 */
proto.operations_ecosys.Response.prototype.setPrimaryKey = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.operations_ecosys.Filter.prototype.toObject = function(opt_includeInstance) {
  return proto.operations_ecosys.Filter.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.operations_ecosys.Filter} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Filter.toObject = function(includeInstance, msg) {
  var f, obj = {
    comparison: jspb.Message.getFieldWithDefault(msg, 1, 0),
    value: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.operations_ecosys.Filter}
 */
proto.operations_ecosys.Filter.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.operations_ecosys.Filter;
  return proto.operations_ecosys.Filter.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.operations_ecosys.Filter} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.operations_ecosys.Filter}
 */
proto.operations_ecosys.Filter.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.operations_ecosys.Filter.Comparisons} */ (reader.readEnum());
      msg.setComparison(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setValue(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.operations_ecosys.Filter.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.operations_ecosys.Filter.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.operations_ecosys.Filter} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.operations_ecosys.Filter.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getComparison();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getValue();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.operations_ecosys.Filter.Comparisons = {
  GREATER: 0,
  GREATER_EQ: 1,
  EQUAL: 2,
  LESSER_EQ: 3,
  LESSER: 4,
  CONTAINS: 5,
  IN: 6,
  NOT_IN: 7
};

/**
 * optional Comparisons comparison = 1;
 * @return {!proto.operations_ecosys.Filter.Comparisons}
 */
proto.operations_ecosys.Filter.prototype.getComparison = function() {
  return /** @type {!proto.operations_ecosys.Filter.Comparisons} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.operations_ecosys.Filter.Comparisons} value
 * @return {!proto.operations_ecosys.Filter} returns this
 */
proto.operations_ecosys.Filter.prototype.setComparison = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string value = 2;
 * @return {string}
 */
proto.operations_ecosys.Filter.prototype.getValue = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.operations_ecosys.Filter} returns this
 */
proto.operations_ecosys.Filter.prototype.setValue = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * @enum {number}
 */
proto.operations_ecosys.OrderBy = {
  ASC: 0,
  DESC: 1
};

goog.object.extend(exports, proto.operations_ecosys);
