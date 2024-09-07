export interface Response<T> {
  data: T;
  message: string;
  status: string;
}

export interface CategoryImage {
  created_at: string;
  updated_at: string;
  image_id: string;
  url: string;
}

export interface Category {
  created_at: string;
  updated_at: string;
  category_id: string;
  user_id: string;
  image_id: string;
  name: string;
  type: string;
  category_image: CategoryImage;
}

export interface Transaction {
  created_at: string;
  updated_at: string;
  category_id: string;
  description?: string;
  amount: number;
  transaction_date: string;
  user_id: string;
  category: Category;
}

export interface CreateTransactionReq {
  category_id: string;
  description?: string;
  amount: number;
  user_id: string;
}