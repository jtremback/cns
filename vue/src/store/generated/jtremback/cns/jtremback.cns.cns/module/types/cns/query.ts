/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cns/params";
import { GRecord } from "../cns/g_record";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "jtremback.cns.cns";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetGRecordRequest {
  key: string;
}

export interface QueryGetGRecordResponse {
  gRecord: GRecord | undefined;
}

export interface QueryAllGRecordRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllGRecordResponse {
  gRecord: GRecord[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetGRecordRequest: object = { key: "" };

export const QueryGetGRecordRequest = {
  encode(
    message: QueryGetGRecordRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGRecordRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetGRecordRequest } as QueryGetGRecordRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGRecordRequest {
    const message = { ...baseQueryGetGRecordRequest } as QueryGetGRecordRequest;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    return message;
  },

  toJSON(message: QueryGetGRecordRequest): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetGRecordRequest>
  ): QueryGetGRecordRequest {
    const message = { ...baseQueryGetGRecordRequest } as QueryGetGRecordRequest;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    return message;
  },
};

const baseQueryGetGRecordResponse: object = {};

export const QueryGetGRecordResponse = {
  encode(
    message: QueryGetGRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.gRecord !== undefined) {
      GRecord.encode(message.gRecord, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetGRecordResponse,
    } as QueryGetGRecordResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gRecord = GRecord.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGRecordResponse {
    const message = {
      ...baseQueryGetGRecordResponse,
    } as QueryGetGRecordResponse;
    if (object.gRecord !== undefined && object.gRecord !== null) {
      message.gRecord = GRecord.fromJSON(object.gRecord);
    } else {
      message.gRecord = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetGRecordResponse): unknown {
    const obj: any = {};
    message.gRecord !== undefined &&
      (obj.gRecord = message.gRecord
        ? GRecord.toJSON(message.gRecord)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetGRecordResponse>
  ): QueryGetGRecordResponse {
    const message = {
      ...baseQueryGetGRecordResponse,
    } as QueryGetGRecordResponse;
    if (object.gRecord !== undefined && object.gRecord !== null) {
      message.gRecord = GRecord.fromPartial(object.gRecord);
    } else {
      message.gRecord = undefined;
    }
    return message;
  },
};

const baseQueryAllGRecordRequest: object = {};

export const QueryAllGRecordRequest = {
  encode(
    message: QueryAllGRecordRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllGRecordRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllGRecordRequest } as QueryAllGRecordRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllGRecordRequest {
    const message = { ...baseQueryAllGRecordRequest } as QueryAllGRecordRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllGRecordRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllGRecordRequest>
  ): QueryAllGRecordRequest {
    const message = { ...baseQueryAllGRecordRequest } as QueryAllGRecordRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllGRecordResponse: object = {};

export const QueryAllGRecordResponse = {
  encode(
    message: QueryAllGRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.gRecord) {
      GRecord.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllGRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllGRecordResponse,
    } as QueryAllGRecordResponse;
    message.gRecord = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gRecord.push(GRecord.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllGRecordResponse {
    const message = {
      ...baseQueryAllGRecordResponse,
    } as QueryAllGRecordResponse;
    message.gRecord = [];
    if (object.gRecord !== undefined && object.gRecord !== null) {
      for (const e of object.gRecord) {
        message.gRecord.push(GRecord.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllGRecordResponse): unknown {
    const obj: any = {};
    if (message.gRecord) {
      obj.gRecord = message.gRecord.map((e) =>
        e ? GRecord.toJSON(e) : undefined
      );
    } else {
      obj.gRecord = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllGRecordResponse>
  ): QueryAllGRecordResponse {
    const message = {
      ...baseQueryAllGRecordResponse,
    } as QueryAllGRecordResponse;
    message.gRecord = [];
    if (object.gRecord !== undefined && object.gRecord !== null) {
      for (const e of object.gRecord) {
        message.gRecord.push(GRecord.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a GRecord by index. */
  GRecord(request: QueryGetGRecordRequest): Promise<QueryGetGRecordResponse>;
  /** Queries a list of GRecord items. */
  GRecordAll(request: QueryAllGRecordRequest): Promise<QueryAllGRecordResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("jtremback.cns.cns.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  GRecord(request: QueryGetGRecordRequest): Promise<QueryGetGRecordResponse> {
    const data = QueryGetGRecordRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jtremback.cns.cns.Query",
      "GRecord",
      data
    );
    return promise.then((data) =>
      QueryGetGRecordResponse.decode(new Reader(data))
    );
  }

  GRecordAll(
    request: QueryAllGRecordRequest
  ): Promise<QueryAllGRecordResponse> {
    const data = QueryAllGRecordRequest.encode(request).finish();
    const promise = this.rpc.request(
      "jtremback.cns.cns.Query",
      "GRecordAll",
      data
    );
    return promise.then((data) =>
      QueryAllGRecordResponse.decode(new Reader(data))
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
