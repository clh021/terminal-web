// package: lnks_homecloud.clh021.webshell
// file: shell.proto

var shell_pb = require("./shell_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ShellService = (function () {
  function ShellService() {}
  ShellService.serviceName = "lnks_homecloud.clh021.webshell.ShellService";
  return ShellService;
}());

ShellService.CMsg = {
  methodName: "CMsg",
  service: ShellService,
  requestStream: true,
  responseStream: false,
  requestType: shell_pb.Msg,
  responseType: shell_pb.Empty
};

ShellService.SMsg = {
  methodName: "SMsg",
  service: ShellService,
  requestStream: false,
  responseStream: true,
  requestType: shell_pb.CmdRequest,
  responseType: shell_pb.Msg
};

ShellService.Run = {
  methodName: "Run2",
  service: ShellService,
  requestStream: true,
  responseStream: true,
  requestType: shell_pb.CmdRequest,
  responseType: shell_pb.Output
};

exports.ShellService = ShellService;

function ShellServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ShellServiceClient.prototype.cMsg = function cMsg(metadata) {
  var listeners = {
    end: [],
    status: []
  };
  var client = grpc.client(ShellService.CMsg, {
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport
  });
  client.onEnd(function (status, statusMessage, trailers) {
    listeners.status.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners.end.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners = null;
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    write: function (requestMessage) {
      if (!client.started) {
        client.start(metadata);
      }
      client.send(requestMessage);
      return this;
    },
    end: function () {
      client.finishSend();
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

ShellServiceClient.prototype.sMsg = function sMsg(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(ShellService.SMsg, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

ShellServiceClient.prototype.run = function run(metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.client(ShellService.Run, {
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport
  });
  client.onEnd(function (status, statusMessage, trailers) {
    listeners.status.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners.end.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners = null;
  });
  client.onMessage(function (message) {
    listeners.data.forEach(function (handler) {
      handler(message);
    })
  });
  client.start(metadata);
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    write: function (requestMessage) {
      client.send(requestMessage);
      return this;
    },
    end: function () {
      client.finishSend();
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.ShellServiceClient = ShellServiceClient;

