import { CommonParam } from ".";

export interface Category {
  id: number;
  title: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

export interface CategoryListParam extends CommonParam {}

export interface CategoryCreateParam extends CommonParam {
  title: string;
}

export interface CategoryDeleteParam extends CommonParam {
  id: number;
}
