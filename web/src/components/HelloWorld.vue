<template>
  <div class="hello">After docker-compose up to test:
    <button @click="grpcWeb">grpcWeb</button>
    <button @click="grpcWebProxy">grpcWebProxy</button>
    <button @click="nginxProxy">nginxProxy</button>
    <hr />
    地址: <input type="text" v-model="grpcHost">
    命令:<input class="sinput" type="text" v-model="cmd">
    参数:<input class="sinput" type="text" v-model="arg">
    <button @click="getCmdResult">发送测试</button><br>
    <textarea v-model="cmdResult" cols="80" rows="40"></textarea>
  </div>
</template>

<script>
import {grpc} from "@improbable-eng/grpc-web";
import {ShellService} from '../lib/shell_pb_service'
import {ShellService as ShellService2} from './shell_pb_service2'
import {CmdRequest} from '../lib/shell_pb'
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
    grpcWeb() {
      this.grpcHost="http://127.0.0.1:7900";
      this.getDiskResult();
    },
    grpcWebProxy(){
      this.grpcHost="http://127.0.0.1:7900";
      this.newGrpcClient2();
      this.arg="-lh";
      this.getCmdResult();
    },
    nginxProxy(){
      this.grpcHost="http://127.0.0.1:7903";
      this.arg="-l";
      this.getCmdResult();
    },
    newGrpcClient2() {
      this.client = grpc.client(ShellService2.Run, {
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
    newGrpcClient() {
      this.client = grpc.client(ShellService.Run, {
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
      this.newGrpcClient();
      this.client.send(cmdRequest);
      // this.client.finishSend();
      // this.client.close();
    },
    getDiskResult() {
      const cmdRequest = new CmdRequest();
      cmdRequest.setProg("df")
      cmdRequest.addArgs("-lh")
      this.newGrpcClient();
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
