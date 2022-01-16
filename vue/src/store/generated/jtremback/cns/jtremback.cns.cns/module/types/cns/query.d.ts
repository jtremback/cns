import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cns/params";
import { GRecord } from "../cns/g_record";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "jtremback.cns.cns";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
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
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryGetGRecordRequest: {
    encode(message: QueryGetGRecordRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetGRecordRequest;
    fromJSON(object: any): QueryGetGRecordRequest;
    toJSON(message: QueryGetGRecordRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetGRecordRequest>): QueryGetGRecordRequest;
};
export declare const QueryGetGRecordResponse: {
    encode(message: QueryGetGRecordResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetGRecordResponse;
    fromJSON(object: any): QueryGetGRecordResponse;
    toJSON(message: QueryGetGRecordResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetGRecordResponse>): QueryGetGRecordResponse;
};
export declare const QueryAllGRecordRequest: {
    encode(message: QueryAllGRecordRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllGRecordRequest;
    fromJSON(object: any): QueryAllGRecordRequest;
    toJSON(message: QueryAllGRecordRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllGRecordRequest>): QueryAllGRecordRequest;
};
export declare const QueryAllGRecordResponse: {
    encode(message: QueryAllGRecordResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllGRecordResponse;
    fromJSON(object: any): QueryAllGRecordResponse;
    toJSON(message: QueryAllGRecordResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllGRecordResponse>): QueryAllGRecordResponse;
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
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    GRecord(request: QueryGetGRecordRequest): Promise<QueryGetGRecordResponse>;
    GRecordAll(request: QueryAllGRecordRequest): Promise<QueryAllGRecordResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
