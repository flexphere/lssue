import JSON_Client from "./JSONClient";
import config from "./config";
import * as model from "./model";

const client = new JSON_Client(config.API_URL);

const API = {
  board: {
    list: async (): Promise<model.Board[]> => await client.post("/board")
  },

  pipe: {
    list: async (data: model.PipeListParam) =>
      <model.Pipe[]>await client.post("/pipe", data)
  },

  ticket: {
    list: async (data: model.TicketListParam) =>
      <Array<model.Ticket>>await client.post("/ticket", data),
    create: async (data: model.TicketCreateParam) =>
      await client.post("/ticket/create", data),
    update: async (data: model.TicketUpdateParam) =>
      await client.post("/ticket/update", data),
    delete: async (data: model.TicketDeleteParam) =>
      await client.post("/ticket/delete", data),
    sort: async (data: model.TicketSortParam) =>
      await client.post("/ticket/sort", data)
  },

  category: {
    list: async (data: model.CategoryListParam) =>
      <model.Category[]>await client.post("/category", data),
    create: async (data: model.CategoryCreateParam) =>
      await client.post("/category/create", data),
    delete: async (data: model.CategoryDeleteParam) =>
      await client.post("/category/delete", data)
  },

  label: {
    list: async (data: model.LabelListParam) =>
      <model.Label[]>await client.post("/label", data),
    create: async (data: model.LabelCreateParam) =>
      await client.post("/label/create", data),
    delete: async (data: model.LabelDeleteParam) =>
      await client.post("/label/delete", data)
  },

  issue: {
    list: async (data: model.IssueListParam) =>
      <model.Issue[]>await client.post("/issue", data),
    bind: async (data: model.IssueBindParam) =>
      await client.post("/issue/bind", data),
    unbind: async (data: model.IssueUnbindParam) =>
      await client.post("/issue/unbind", data)
  }
};

export default API;
