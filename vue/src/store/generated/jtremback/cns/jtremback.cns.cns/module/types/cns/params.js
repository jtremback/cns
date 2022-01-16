/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "jtremback.cns.cns";
const baseParams = { writer: "" };
export const Params = {
    encode(message, writer = Writer.create()) {
        if (message.writer !== "") {
            writer.uint32(10).string(message.writer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseParams };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.writer = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseParams };
        if (object.writer !== undefined && object.writer !== null) {
            message.writer = String(object.writer);
        }
        else {
            message.writer = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.writer !== undefined && (obj.writer = message.writer);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseParams };
        if (object.writer !== undefined && object.writer !== null) {
            message.writer = object.writer;
        }
        else {
            message.writer = "";
        }
        return message;
    },
};
