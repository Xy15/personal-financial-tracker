import Form from "./common/Form";
import { InputProps, SelectOptionsProps } from "./common/Input";
import { useMutation, useQueryClient } from "react-query";
import { createTransaction } from "../api/transaction/transaction";
import { CreateTransactionReq } from "../api/types";
import { QUERY_KEYS, queryKeys } from "../api/queryKeys";
import { USER_ID } from "../constants/constants";
import { AxiosError } from "axios";

interface TransactionFormProps {
  user_category_id: string;
  transaction_date: string;
  description: string;
  type: "Income" | "Expense";
  amount: number;
}

const TransactionForm = () => {
  const queryClient = useQueryClient();

  const { mutate } = useMutation(createTransaction, {
    onSuccess: (data) => {
      console.log("data", data);
      queryClient.invalidateQueries(queryKeys[QUERY_KEYS.transaction].getTransactionsByUserID(USER_ID).queryKey);
    },
    onError: (error: AxiosError) => {
      console.log("error", error.response?.data);
    },
  });
  //test
//test
  const onSubmit = (formData: TransactionFormProps) => {
    console.log("Form data submitted:", formData);

    const transaction: CreateTransactionReq = {
      ...formData,
      amount: Number(formData.amount),
      user_id: USER_ID,
    };

    mutate(transaction);
  };

  const categoryList: SelectOptionsProps[] = [
    //Get from DB
    {
      value: "3b008112-353b-4156-8e78-ddb25d0ee5e1",
      name: "Food",
    },
    {
      value: "59236e85-dee0-4000-b42f-bfcbd0432d46",
      name: "Drink",
    },
    {
      value: "7aaec20b-49f2-4525-bc40-f66a7420659f",
      name: "Jeans",
    },
    {
      value: "976a664a-f3ca-41ef-96c8-b85525ba56d8",
      name: "Snack",
    },
  ];

  const transactionFormInputs: InputProps[] = [
    {
      name: "user_category_id",
      label: "Category",
      type: "select",
      options: categoryList,
    },
    {
      name: "transaction_date",
      label: "Date",
    },
    {
      name: "description",
      label: "Description",
    },
    {
      name: "type",
      label: "Type",
      type: "select",
      options: [
        { value: 'income', name: 'Income' },
        { value: 'expense', name: 'Expense' },
      ],
      // rules: {
      //   validate: (value: string) =>
      //     ["in", "out"].includes(value) || 'Type must be "in" or "out"',
      // },
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
