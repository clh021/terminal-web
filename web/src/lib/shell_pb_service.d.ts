// package: lnks_homecloud.clh021.webshell
// file: shell.proto

import * as shell_pb from "./shell_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ShellServiceCMsg = {
  readonly methodName: string;
  readonly service: typeof ShellService;
  readonly requestStream: true;
  readonly responseStream: false;
  readonly requestType: typeof shell_pb.Msg;
  readonly responseType: typeof shell_pb.Empty;
};

type ShellServiceSMsg = {
  readonly methodName: string;
  readonly service: typeof ShellService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof shell_pb.CmdRequest;
  readonly responseType: typeof shell_pb.Msg;
};

type ShellServiceRun = {
  readonly methodName: string;
  readonly service: typeof ShellService;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof shell_pb.CmdRequest;
  readonly responseType: typeof shell_pb.Output;
};

export class ShellService {
  static readonly serviceName: string;
  static readonly CMsg: ShellServiceCMsg;
  static readonly SMsg: ShellServiceSMsg;
  static readonly Run: ShellServiceRun;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ShellServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  cMsg(metadata?: grpc.Metadata): RequestStream<shell_pb.Msg>;
  sMsg(requestMessage: shell_pb.CmdRequest, metadata?: grpc.Metadata): ResponseStream<shell_pb.Msg>;
  run(metadata?: grpc.Metadata): BidirectionalStream<shell_pb.CmdRequest, shell_pb.Output>;
}

