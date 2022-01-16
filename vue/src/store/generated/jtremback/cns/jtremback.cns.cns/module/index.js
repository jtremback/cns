// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateGRecord } from "./types/cns/tx";
import { MsgCreateGRecord } from "./types/cns/tx";
import { MsgDeleteGRecord } from "./types/cns/tx";
const types = [
    ["/jtremback.cns.cns.MsgUpdateGRecord", MsgUpdateGRecord],
    ["/jtremback.cns.cns.MsgCreateGRecord", MsgCreateGRecord],
    ["/jtremback.cns.cns.MsgDeleteGRecord", MsgDeleteGRecord],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgUpdateGRecord: (data) => ({ typeUrl: "/jtremback.cns.cns.MsgUpdateGRecord", value: MsgUpdateGRecord.fromPartial(data) }),
        msgCreateGRecord: (data) => ({ typeUrl: "/jtremback.cns.cns.MsgCreateGRecord", value: MsgCreateGRecord.fromPartial(data) }),
        msgDeleteGRecord: (data) => ({ typeUrl: "/jtremback.cns.cns.MsgDeleteGRecord", value: MsgDeleteGRecord.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
