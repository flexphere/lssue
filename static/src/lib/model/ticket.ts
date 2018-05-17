import { CommonParam, Category, Label, Issue } from ".";

export interface Ticket {
  id: number;
  title: string;
  due: string;
  position: number;
  memo: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
  pipe_id: number;
  category_id: number;
  label_ids: Array<number>;
  issue_ids: Array<number>;
  category: Category;
  labels: Array<Label>;
  issues: Array<Issue>;
}

export interface TicketListParam extends CommonParam {}

export interface TicketCreateParam extends CommonParam {
  title: string;
  due: string;
  memo: string;
  pipe_id: number;
  category_id: number;
  label_ids: Array<number>;
}

export interface TicketUpdateParam extends TicketCreateParam {
  id: number;
}

export interface TicketDeleteParam extends CommonParam {
  id: number;
}

export interface TicketSortParam extends CommonParam {
  pipe_id: number;
  tickets: Array<number>;
}
