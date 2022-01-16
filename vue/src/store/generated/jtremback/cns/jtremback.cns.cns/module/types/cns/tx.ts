/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "jtremback.cns.cns";

export interface MsgCreateGRecord {
  creator: string;
  key: string;
  value: string;
}

export interface MsgCreateGRecordResponse {}

export interface MsgUpdateGRecord {
  creator: string;
  key: string;
  value: string;
}

export interface MsgUpdateGRecordResponse {}

export interface MsgDeleteGRecord {
  creator: string;
  key: string;
}

export interface MsgDeleteGRecordResponse {}

const baseMsgCreateGRecord: object = { creator: "", key: "", value: "" };

export const MsgCreateGRecord = {
  encode(message: MsgCreateGRecord, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateGRecord {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateGRecord } as MsgCreateGRecord;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.key = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGRecord {
    const message = { ...baseMsgCreateGRecord } as MsgCreateGRecord;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: MsgCreateGRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateGRecord>): MsgCreateGRecord {
    const message = { ...baseMsgCreateGRecord } as MsgCreateGRecord;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseMsgCreateGRecordResponse: object = {};

export const MsgCreateGRecordResponse = {
  encode(
    _: MsgCreateGRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateGRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateGRecordResponse,
    } as MsgCreateGRecordResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateGRecordResponse {
    const message = {
      ...baseMsgCreateGRecordResponse,
    } as MsgCreateGRecordResponse;
    return message;
  },

  toJSON(_: MsgCreateGRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateGRecordResponse>
  ): MsgCreateGRecordResponse {
    const message = {
      ...baseMsgCreateGRecordResponse,
    } as MsgCreateGRecordResponse;
    return message;
  },
};

const baseMsgUpdateGRecord: object = { creator: "", key: "", value: "" };

export const MsgUpdateGRecord = {
  encode(message: MsgUpdateGRecord, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateGRecord {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateGRecord } as MsgUpdateGRecord;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.key = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateGRecord {
    const message = { ...baseMsgUpdateGRecord } as MsgUpdateGRecord;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateGRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateGRecord>): MsgUpdateGRecord {
    const message = { ...baseMsgUpdateGRecord } as MsgUpdateGRecord;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseMsgUpdateGRecordResponse: object = {};

export const MsgUpdateGRecordResponse = {
  encode(
    _: MsgUpdateGRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateGRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateGRecordResponse,
    } as MsgUpdateGRecordResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateGRecordResponse {
    const message = {
      ...baseMsgUpdateGRecordResponse,
    } as MsgUpdateGRecordResponse;
    return message;
  },

  toJSON(_: MsgUpdateGRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateGRecordResponse>
  ): MsgUpdateGRecordResponse {
    const message = {
      ...baseMsgUpdateGRecordResponse,
    } as MsgUpdateGRecordResponse;
    return message;
  },
};

const baseMsgDeleteGRecord: object = { creator: "", key: "" };

export const MsgDeleteGRecord = {
  encode(message: MsgDeleteGRecord, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteGRecord {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteGRecord } as MsgDeleteGRecord;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.key = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteGRecord {
    const message = { ...baseMsgDeleteGRecord } as MsgDeleteGRecord;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteGRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteGRecord>): MsgDeleteGRecord {
    const message = { ...baseMsgDeleteGRecord } as MsgDeleteGRecord;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    return message;
  },
};

const baseMsgDeleteGRecordResponse: object = {};

export const MsgDeleteGRecordResponse = {
  encode(
    _: MsgDeleteGRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteGRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteGRecordResponse,
    } as MsgDeleteGRecordResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteGRecordResponse {
    const message = {
      ...baseMsgDeleteGRecordResponse,
    } as MsgDeleteGRecordResponse;
    return message;
  },

  toJSON(_: MsgDeleteGRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteGRecordResponse>
  ): MsgDeleteGRecordResponse {
    const message = {
      ...baseMsgDeleteGRecordResponse,
    } as MsgDeleteGRecordResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateGRecord(request: MsgCreateGRecord): Promise<MsgCreateGRecordResponse>;
  UpdateGRecord(request: MsgUpdateGRecord): Promise<MsgUpdateGRecordResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteGRecord(request: MsgDeleteGRecord): Promise<MsgDeleteGRecordResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateGRecord(request: MsgCreateGRecord): Promise<MsgCreateGRecordResponse> {
    const data = MsgCreateGRecord.encode(request).finish();
    const promise = this.rpc.request(
      "jtremback.cns.cns.Msg",
      "CreateGRecord",
      data
    );
    return promise.then((data) =>
      MsgCreateGRecordResponse.decode(new Reader(data))
    );
  }

  UpdateGRecord(request: MsgUpdateGRecord): Promise<MsgUpdateGRecordResponse> {
    const data = MsgUpdateGRecord.encode(request).finish();
    const promise = this.rpc.request(
      "jtremback.cns.cns.Msg",
      "UpdateGRecord",
      data
    );
    return promise.then((data) =>
      MsgUpdateGRecordResponse.decode(new Reader(data))
    );
  }

  DeleteGRecord(request: MsgDeleteGRecord): Promise<MsgDeleteGRecordResponse> {
    const data = MsgDeleteGRecord.encode(request).finish();
    const promise = this.rpc.request(
      "jtremback.cns.cns.Msg",
      "DeleteGRecord",
      data
    );
    return promise.then((data) =>
      MsgDeleteGRecordResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
