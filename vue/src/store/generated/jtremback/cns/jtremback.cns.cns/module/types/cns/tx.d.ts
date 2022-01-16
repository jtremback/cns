import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "jtremback.cns.cns";
export interface MsgCreateGRecord {
    creator: string;
    key: string;
    value: string;
}
export interface MsgCreateGRecordResponse {
}
export interface MsgUpdateGRecord {
    creator: string;
    key: string;
    value: string;
}
export interface MsgUpdateGRecordResponse {
}
export interface MsgDeleteGRecord {
    creator: string;
    key: string;
}
export interface MsgDeleteGRecordResponse {
}
export declare const MsgCreateGRecord: {
    encode(message: MsgCreateGRecord, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateGRecord;
    fromJSON(object: any): MsgCreateGRecord;
    toJSON(message: MsgCreateGRecord): unknown;
    fromPartial(object: DeepPartial<MsgCreateGRecord>): MsgCreateGRecord;
};
export declare const MsgCreateGRecordResponse: {
    encode(_: MsgCreateGRecordResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateGRecordResponse;
    fromJSON(_: any): MsgCreateGRecordResponse;
    toJSON(_: MsgCreateGRecordResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateGRecordResponse>): MsgCreateGRecordResponse;
};
export declare const MsgUpdateGRecord: {
    encode(message: MsgUpdateGRecord, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateGRecord;
    fromJSON(object: any): MsgUpdateGRecord;
    toJSON(message: MsgUpdateGRecord): unknown;
    fromPartial(object: DeepPartial<MsgUpdateGRecord>): MsgUpdateGRecord;
};
export declare const MsgUpdateGRecordResponse: {
    encode(_: MsgUpdateGRecordResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateGRecordResponse;
    fromJSON(_: any): MsgUpdateGRecordResponse;
    toJSON(_: MsgUpdateGRecordResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateGRecordResponse>): MsgUpdateGRecordResponse;
};
export declare const MsgDeleteGRecord: {
    encode(message: MsgDeleteGRecord, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteGRecord;
    fromJSON(object: any): MsgDeleteGRecord;
    toJSON(message: MsgDeleteGRecord): unknown;
    fromPartial(object: DeepPartial<MsgDeleteGRecord>): MsgDeleteGRecord;
};
export declare const MsgDeleteGRecordResponse: {
    encode(_: MsgDeleteGRecordResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteGRecordResponse;
    fromJSON(_: any): MsgDeleteGRecordResponse;
    toJSON(_: MsgDeleteGRecordResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteGRecordResponse>): MsgDeleteGRecordResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateGRecord(request: MsgCreateGRecord): Promise<MsgCreateGRecordResponse>;
    UpdateGRecord(request: MsgUpdateGRecord): Promise<MsgUpdateGRecordResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteGRecord(request: MsgDeleteGRecord): Promise<MsgDeleteGRecordResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateGRecord(request: MsgCreateGRecord): Promise<MsgCreateGRecordResponse>;
    UpdateGRecord(request: MsgUpdateGRecord): Promise<MsgUpdateGRecordResponse>;
    DeleteGRecord(request: MsgDeleteGRecord): Promise<MsgDeleteGRecordResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
