import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "jtremback.cns.cns";
export interface GRecord {
    key: string;
    value: string;
    creator: string;
}
export declare const GRecord: {
    encode(message: GRecord, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GRecord;
    fromJSON(object: any): GRecord;
    toJSON(message: GRecord): unknown;
    fromPartial(object: DeepPartial<GRecord>): GRecord;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
