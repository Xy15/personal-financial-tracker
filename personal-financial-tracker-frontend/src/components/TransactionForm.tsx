import Form from "./Form";
import { InputProps, SelectOptionsProps } from "./Input";
import { useMutation, useQueryClient } from "react-query";
import { createTransaction } from "../api/transaction/transaction";
import { CreateTransactionReq } from "../api/types";
import { QUERY_KEYS, queryKeys } from "../api/queryKeys";
import { USER_ID } from "../constants/constants";

interface TransactionFormProps {
  category_id: string;
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
    onError: (error) => {
      console.log("error", error);
    },
  });

  const onSubmit = (formData: TransactionFormProps) => {
    console.log("Form data submitted:", formData);

    const transaction: CreateTransactionReq = {
      ...formData,
      user_id: USER_ID,
    };

    mutate(transaction);
  };

  const categoryList: SelectOptionsProps[] = [
    //Get from DB
    {
      value: "9590d800-2775-492d-b222-f961b783197b",
      name: "Food",
    },
    {
      value: "drink",
      name: "Drink",
    },
    {
      value: "bag",
      name: "Bag",
    },
    {
      value: "phone",
      name: "Phone",
    },
  ];

  const transactionFormInputs: InputProps[] = [
    {
      name: "category",
      label: "Category",
      type: "select",
      options: categoryList,
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
