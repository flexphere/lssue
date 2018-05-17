export default class JSON_Client {
  base_url: string = "http://127.0.0.1";

  options: RequestInit = {
    method: "get",
    mode: "cors",
    credentials: "include"
  };

  constructor(base_url: string, options?: RequestInit) {
    this.base_url = base_url;
    if (options) this.options = options;
  }

  async post(path: string, data?: any) {
    this.options.method = "post";
    if (data) this.options.body = JSON.stringify(data);
    return this.request(path);
  }

  async request(path: string) {
    let response = await fetch(this.base_url + path, this.options);
    if (response.status > 299 && response.status != 404) {
      throw new Error(response.statusText);
    }

    if (response.status == 204) {
      return true;
    }

    return await response.json();
  }
}
