import Form from "./common/Form";
import { InputProps, SelectOptionsProps } from "./common/Input";
import { useMutation, useQueryClient } from "react-query";
import { createTransaction } from "../api/transaction/transaction";
import { CreateTransactionReq } from "../api/types";
import { QUERY_KEYS, queryKeys } from "../api/queryKeys";
import { USER_ID } from "../constants/constants";
import { AxiosError } from "axios";
import { useContext } from "react";
import { CategoryContext } from "./provider/CategoryProvider";

interface TransactionFormProps {
  transaction_date: string;
  description: string;
  type: "Income" | "Expense";
  amount: number;
}

const TransactionForm = () => {
  const queryClient = useQueryClient();
  const { selectedCategory } = useContext(CategoryContext);

  const { mutate } = useMutation(createTransaction, {
    onSuccess: (data) => {
      console.log("data", data);
      queryClient.invalidateQueries(queryKeys[QUERY_KEYS.transaction].getTransactionsByUserID(USER_ID).queryKey);
    },
    onError: (error: AxiosError) => {
      console.log("error", error.response?.data);
    },
  });

  const onSubmit = (formData: TransactionFormProps) => {
    console.log("Form data submitted:", formData);

    const transaction: CreateTransactionReq = {
      ...formData,
      amount: Number(formData.amount),
      user_id: USER_ID,
      user_category_id: selectedCategory,
    };

    mutate(transaction);
  };

  const transactionFormInputs: InputProps[] = [
    {
      name: "transaction_date",
      label: "Date",
    },
    {
      name: "description",
      label: "Description",
    },
    {
      name: "amount",
      label: "Amount",
      type: "number",
    },
  ];

  return <Form formInputs={transactionFormInputs} onSubmit={onSubmit} />;
};

export default TransactionForm;
