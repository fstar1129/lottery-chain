/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "lotterychainnel.lottery";

export interface MsgClaimOwner {
  creator: string;
}

export interface MsgClaimOwnerResponse {
  success: string;
}

const baseMsgClaimOwner: object = { creator: "" };

export const MsgClaimOwner = {
  encode(message: MsgClaimOwner, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClaimOwner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgClaimOwner } as MsgClaimOwner;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimOwner {
    const message = { ...baseMsgClaimOwner } as MsgClaimOwner;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgClaimOwner): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgClaimOwner>): MsgClaimOwner {
    const message = { ...baseMsgClaimOwner } as MsgClaimOwner;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgClaimOwnerResponse: object = { success: "" };

export const MsgClaimOwnerResponse = {
  encode(
    message: MsgClaimOwnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success !== "") {
      writer.uint32(10).string(message.success);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClaimOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgClaimOwnerResponse } as MsgClaimOwnerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimOwnerResponse {
    const message = { ...baseMsgClaimOwnerResponse } as MsgClaimOwnerResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = String(object.success);
    } else {
      message.success = "";
    }
    return message;
  },

  toJSON(message: MsgClaimOwnerResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgClaimOwnerResponse>
  ): MsgClaimOwnerResponse {
    const message = { ...baseMsgClaimOwnerResponse } as MsgClaimOwnerResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ClaimOwner(request: MsgClaimOwner): Promise<MsgClaimOwnerResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  ClaimOwner(request: MsgClaimOwner): Promise<MsgClaimOwnerResponse> {
    const data = MsgClaimOwner.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Msg",
      "ClaimOwner",
      data
    );
    return promise.then((data) =>
      MsgClaimOwnerResponse.decode(new Reader(data))
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
