import { CommonParam } from ".";

export interface Label {
  id: number;
  title: string;
  bgcolor: string;
  color: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

export interface LabelListParam extends CommonParam {}

export interface LabelCreateParam extends CommonParam {
  title: string;
}

export interface LabelDeleteParam extends CommonParam {
  id: number;
}
