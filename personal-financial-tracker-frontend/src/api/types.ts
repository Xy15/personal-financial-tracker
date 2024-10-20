import { EnumType } from "typescript";

export interface Response<T> {
  data: T;
  message: string;
  status: string;
}

export type CategoryType = "Income" | "Expense"

export interface CategoryImage {
  created_at: string;
  updated_at: string;
  image_id: string;
  image_url: string;
}

export interface UserCategory {
  id: string;
  created_at: string;
  updated_at: string;
  user_id: string;
  category_image_id: string;
  name: string;
  type: string;
  category_image: CategoryImage;
}

export type GroupedTransactions = {
  [date: string]: Transaction[];
};

export interface Transaction {
  amount: number;
  category_image_url: string;
  category_name: string;
  created_at: string;
  description: string;
  id: string;
  transaction_date: string;
  updated_at: string;
  user_id: string;
};

export interface CreateTransactionReq {
  user_category_id: string;
  description?: string;
  amount: number;
  transaction_date: string;
  user_id: string;
}