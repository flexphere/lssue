import * as model from "./model";

const dev_config: model.Config = {
  API_URL: "http://127.0.0.1:8080/api",
  WS_URL: "ws://127.0.0.1:8080/ws"
};

const prod_config: model.Config = {
  API_URL: "https://lssue.io/api",
  WS_URL: "wss://lssue.io/ws"
};

export default (process.env.NODE_ENV == "production"
  ? prod_config
  : dev_config);
