import * as jspb from 'google-protobuf'



export class Gate extends jspb.Message {
  getId(): number;
  setId(value: number): Gate;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Gate.AsObject;
  static toObject(includeInstance: boolean, msg: Gate): Gate.AsObject;
  static serializeBinaryToWriter(message: Gate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Gate;
  static deserializeBinaryFromReader(message: Gate, reader: jspb.BinaryReader): Gate;
}

export namespace Gate {
  export type AsObject = {
    id: number,
  }
}

export class GateState extends jspb.Message {
  getId(): number;
  setId(value: number): GateState;

  getState(): GateState.GatePosition;
  setState(value: GateState.GatePosition): GateState;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GateState.AsObject;
  static toObject(includeInstance: boolean, msg: GateState): GateState.AsObject;
  static serializeBinaryToWriter(message: GateState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GateState;
  static deserializeBinaryFromReader(message: GateState, reader: jspb.BinaryReader): GateState;
}

export namespace GateState {
  export type AsObject = {
    id: number,
    state: GateState.GatePosition,
  }

  export enum GatePosition { 
    CLOSED = 0,
    OPEN = 1,
    ERROR = 2,
    INITIAL = 3,
  }
}

export class FireAlarm extends jspb.Message {
  getId(): number;
  setId(value: number): FireAlarm;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FireAlarm.AsObject;
  static toObject(includeInstance: boolean, msg: FireAlarm): FireAlarm.AsObject;
  static serializeBinaryToWriter(message: FireAlarm, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FireAlarm;
  static deserializeBinaryFromReader(message: FireAlarm, reader: jspb.BinaryReader): FireAlarm;
}

export namespace FireAlarm {
  export type AsObject = {
    id: number,
  }
}

export class FireAlarmState extends jspb.Message {
  getId(): number;
  setId(value: number): FireAlarmState;

  getState(): FireAlarmState.AlarmState;
  setState(value: FireAlarmState.AlarmState): FireAlarmState;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FireAlarmState.AsObject;
  static toObject(includeInstance: boolean, msg: FireAlarmState): FireAlarmState.AsObject;
  static serializeBinaryToWriter(message: FireAlarmState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FireAlarmState;
  static deserializeBinaryFromReader(message: FireAlarmState, reader: jspb.BinaryReader): FireAlarmState;
}

export namespace FireAlarmState {
  export type AsObject = {
    id: number,
    state: FireAlarmState.AlarmState,
  }

  export enum AlarmState { 
    OFF = 0,
    ON = 1,
    ERROR = 2,
  }
}

export class CpuTemp extends jspb.Message {
  getId(): number;
  setId(value: number): CpuTemp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CpuTemp.AsObject;
  static toObject(includeInstance: boolean, msg: CpuTemp): CpuTemp.AsObject;
  static serializeBinaryToWriter(message: CpuTemp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CpuTemp;
  static deserializeBinaryFromReader(message: CpuTemp, reader: jspb.BinaryReader): CpuTemp;
}

export namespace CpuTemp {
  export type AsObject = {
    id: number,
  }
}

export class CpuTempState extends jspb.Message {
  getId(): number;
  setId(value: number): CpuTempState;

  getTemp(): number;
  setTemp(value: number): CpuTempState;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CpuTempState.AsObject;
  static toObject(includeInstance: boolean, msg: CpuTempState): CpuTempState.AsObject;
  static serializeBinaryToWriter(message: CpuTempState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CpuTempState;
  static deserializeBinaryFromReader(message: CpuTempState, reader: jspb.BinaryReader): CpuTempState;
}

export namespace CpuTempState {
  export type AsObject = {
    id: number,
    temp: number,
  }
}

