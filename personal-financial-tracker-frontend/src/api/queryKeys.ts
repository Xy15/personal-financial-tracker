import {
  createQueryKeys,
  mergeQueryKeys,
} from "@lukemorales/query-key-factory";

export const QUERY_KEYS = {
  transaction: "transaction",
  user: "user",
  category: "category",
  category_image: "category_image",
};

const transactionQueryKeys = createQueryKeys(QUERY_KEYS.transaction, {
  getTransactionsByUserID: (user_id: string) => ["user_id", user_id],
});

const userQueryKeys = createQueryKeys(QUERY_KEYS.user, {
  getUserCategoriesByUserID: (user_id: string) => ["user_id", user_id],
});


export const queryKeys = mergeQueryKeys(
  transactionQueryKeys,
  userQueryKeys,
);
