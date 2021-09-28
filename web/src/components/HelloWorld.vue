<template>
  <div class="hello">After docker-compose up:
    <button @click="handleGrpcWeb">grpcWeb</button>
    <button @click="handleGrpcWebProxy">grpcWebProxy</button>
    <button @click="handleNginxProxy('http://grpc21.localhost:7564','-l','ls')">nginxProxyS1</button>
    <button @click="handleNginxProxy('http://grpc22.localhost:7564','-lah','ls')">nginxProxyS2</button>
    <hr />
    地址: <input type="text" v-model="grpcHost">
    命令:<input class="sinput" type="text" v-model="cmd">
    参数:<input class="sinput" type="text" v-model="arg">
    <button @click="handleDoCmd">发送测试</button><br>
    <textarea v-model="cmdResult" cols="80" rows="40"></textarea>
  </div>
</template>

<script>
import {grpc} from "@improbable-eng/grpc-web";
import {ShellService} from '../lib/shell_pb_service'
import {CmdRequest} from '../lib/shell_pb'
import {ShellService as ShellService2} from '../lib2/shell_pb_service'
// import {CmdRequest as CmdRequest2} from '../lib2/shell_pb'
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data() {
    return {
      cmd: "ls",
      arg: "-lah",
      cmdResult: "",
      client: null,
      grpcHost: "http://127.0.0.1:7900"
    }
  },
  mounted() {
      // grpc.setDefaultTransport(grpc.WebsocketTransport()); // 目前仅 nginx proxy 方式还不支持 websocket 访问
  },
  methods: {
    handleGrpcWeb() {
      this.grpcHost="http://127.0.0.1:7900";
      this.newGrpcClient();
      this.getDiskResult();
    },
    handleGrpcWebProxy(){
      this.grpcHost="http://127.0.0.1:7900";
      this.newGrpcClient(ShellService2.Run2);
      this.arg="-lh";
      this.getCmdResult();
    },
    handleNginxProxy(grpcHost="http://127.0.0.1:7564", arg="-l", cmd="ls"){
      this.grpcHost=grpcHost;
      this.arg=arg;
      this.cmd=cmd;
      this.newGrpcClient(ShellService2.Run2);
      this.getCmdResult();
    },
    handleDoCmd() {
      this.newGrpcClient();
      this.getCmdResult();
    },
    newGrpcClient(methods=ShellService.Run) {
      this.client = grpc.client(methods, {
        host: this.grpcHost,
        debug: true,
      });
      this.client.onMessage((message) => {
        let m = message.getStdout();
        if (m.length > 0) {
          this.cmdResult = (new TextDecoder()).decode(m);
          console.log(this.cmdResult);
          console.log("message:", message.toObject());
        } else {
          console.log("message:", message.toObject());
        }
      });
      this.client.onHeaders((headers) => {
        console.log("onHeaders", headers);
      });
      this.client.onEnd((code, msg, trailers) => {
        console.log("onEnd", code, msg, trailers);
      });
      this.client.start();
    },
    getCmdResult() {
      const cmdRequest = new CmdRequest();
      cmdRequest.setProg(this.cmd)
      let args = this.arg.split(" ")
      args.forEach(arg => {
        console.log("arg:", arg);
        cmdRequest.addArgs(arg)
      });
      this.client.send(cmdRequest);
      // this.client.finishSend();
      // this.client.close();
    },
    getDiskResult() {
      const cmdRequest = new CmdRequest();
      cmdRequest.setProg("df")
      cmdRequest.addArgs("-lh")
      this.client.send(cmdRequest);
      // this.client.finishSend();
      // this.client.close();
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.sinput{
  width:50px;
}
</style>
