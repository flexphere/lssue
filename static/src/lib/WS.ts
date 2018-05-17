import Vue from "vue";
import config from "./config";
import * as model from "./model";

const bus = new Vue({ data: model.WS_Message });
let client: WebSocket;

const connect = (board_name: string) => {
  if (client) client.close();
  client = new WebSocket(`${config.WS_URL}/${board_name}`);
  client.onmessage = recvMessage;
};

const recvMessage = (e: MessageEvent) => {
  bus.$emit("receive", e.data);
};

const sendMessage = (message: string) => {
  client.send(message);
};

bus.$on("send", sendMessage);
bus.$on("connect", connect);

export default bus;
