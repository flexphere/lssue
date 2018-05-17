import { CommonParam, Ticket } from ".";

export interface Pipe {
  id: number;
  title: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
  tickets: Array<Ticket>;
}

export interface PipeListParam extends CommonParam {}
