// package: lnks_homecloud.clh021.webshell
// file: shell.proto

import * as jspb from "google-protobuf";

export class Msg extends jspb.Message {
  getS(): string;
  setS(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Msg.AsObject;
  static toObject(includeInstance: boolean, msg: Msg): Msg.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Msg, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Msg;
  static deserializeBinaryFromReader(message: Msg, reader: jspb.BinaryReader): Msg;
}

export namespace Msg {
  export type AsObject = {
    s: string,
  }
}

export class CmdRequest extends jspb.Message {
  getProg(): string;
  setProg(value: string): void;

  clearArgsList(): void;
  getArgsList(): Array<string>;
  setArgsList(value: Array<string>): void;
  addArgs(value: string, index?: number): string;

  getStdin(): string;
  setStdin(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CmdRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CmdRequest): CmdRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CmdRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CmdRequest;
  static deserializeBinaryFromReader(message: CmdRequest, reader: jspb.BinaryReader): CmdRequest;
}

export namespace CmdRequest {
  export type AsObject = {
    prog: string,
    argsList: Array<string>,
    stdin: string,
  }
}

export class Output extends jspb.Message {
  getStdout(): Uint8Array | string;
  getStdout_asU8(): Uint8Array;
  getStdout_asB64(): string;
  setStdout(value: Uint8Array | string): void;

  getStderr(): Uint8Array | string;
  getStderr_asU8(): Uint8Array;
  getStderr_asB64(): string;
  setStderr(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Output.AsObject;
  static toObject(includeInstance: boolean, msg: Output): Output.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Output, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Output;
  static deserializeBinaryFromReader(message: Output, reader: jspb.BinaryReader): Output;
}

export namespace Output {
  export type AsObject = {
    stdout: Uint8Array | string,
    stderr: Uint8Array | string,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

