/**
 * @fileoverview gRPC-Web generated client stub for gate_prototype
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as iot_prototype_pb from './iot_prototype_pb';


export class IotControlPrototypeServiceClient {
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

  methodDescriptorGetGateState = new grpcWeb.MethodDescriptor(
    '/gate_prototype.IotControlPrototypeService/GetGateState',
    grpcWeb.MethodType.SERVER_STREAMING,
    iot_prototype_pb.Gate,
    iot_prototype_pb.GateState,
    (request: iot_prototype_pb.Gate) => {
      return request.serializeBinary();
    },
    iot_prototype_pb.GateState.deserializeBinary
  );

  getGateState(
    request: iot_prototype_pb.Gate,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<iot_prototype_pb.GateState> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/gate_prototype.IotControlPrototypeService/GetGateState',
      request,
      metadata || {},
      this.methodDescriptorGetGateState);
  }

  methodDescriptorGetFireAlarmState = new grpcWeb.MethodDescriptor(
    '/gate_prototype.IotControlPrototypeService/GetFireAlarmState',
    grpcWeb.MethodType.SERVER_STREAMING,
    iot_prototype_pb.FireAlarm,
    iot_prototype_pb.FireAlarmState,
    (request: iot_prototype_pb.FireAlarm) => {
      return request.serializeBinary();
    },
    iot_prototype_pb.FireAlarmState.deserializeBinary
  );

  getFireAlarmState(
    request: iot_prototype_pb.FireAlarm,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<iot_prototype_pb.FireAlarmState> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/gate_prototype.IotControlPrototypeService/GetFireAlarmState',
      request,
      metadata || {},
      this.methodDescriptorGetFireAlarmState);
  }

  methodDescriptorGetCpuTemp = new grpcWeb.MethodDescriptor(
    '/gate_prototype.IotControlPrototypeService/GetCpuTemp',
    grpcWeb.MethodType.SERVER_STREAMING,
    iot_prototype_pb.CpuTemp,
    iot_prototype_pb.CpuTempState,
    (request: iot_prototype_pb.CpuTemp) => {
      return request.serializeBinary();
    },
    iot_prototype_pb.CpuTempState.deserializeBinary
  );

  getCpuTemp(
    request: iot_prototype_pb.CpuTemp,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<iot_prototype_pb.CpuTempState> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/gate_prototype.IotControlPrototypeService/GetCpuTemp',
      request,
      metadata || {},
      this.methodDescriptorGetCpuTemp);
  }

}

