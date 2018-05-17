import * as model from "./model";

const dev_config: model.Config = {
  API_URL: "http://127.0.0.1:8080/api",
  WS_URL: "ws://127.0.0.1:8080/ws"
};

const prod_config: model.Config = {
  API_URL: "https://lssue.io/api",
  WS_URL: "ws://lssue.io/ws"
};

export default (IS_PROD ? prod_config : dev_config);
