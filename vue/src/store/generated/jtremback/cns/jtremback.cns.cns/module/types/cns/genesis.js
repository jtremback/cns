/* eslint-disable */
import { Params } from "../cns/params";
import { GRecord } from "../cns/g_record";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "jtremback.cns.cns";
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.gRecordList) {
            GRecord.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.gRecordList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.gRecordList.push(GRecord.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.gRecordList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.gRecordList !== undefined && object.gRecordList !== null) {
            for (const e of object.gRecordList) {
                message.gRecordList.push(GRecord.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.gRecordList) {
            obj.gRecordList = message.gRecordList.map((e) => e ? GRecord.toJSON(e) : undefined);
        }
        else {
            obj.gRecordList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.gRecordList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.gRecordList !== undefined && object.gRecordList !== null) {
            for (const e of object.gRecordList) {
                message.gRecordList.push(GRecord.fromPartial(e));
            }
        }
        return message;
    },
};
