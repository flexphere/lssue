import { CommonParam } from ".";

export interface Issue {
  id: number;
  issue_id: number;
  repo: string;
  title: string;
  state: string;
  url: string;
  assignees: string;
  original: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

export interface IssueListParam extends CommonParam {}

export interface IssueBindParam extends CommonParam {
  id: number;
  ticket_id: number;
}

export interface IssueUnbindParam extends CommonParam {
  id: number;
  ticket_id: number;
}
