import Form from "./Form";
import { InputProps, SelectOptionsProps } from "./Input";
import { useMutation, useQueryClient } from "react-query";
import { postTransaction } from "../api/transaction/transaction";
import { Transaction } from "../api/types";
import { QUERY_KEYS, queryKeys } from "../api/queryKeys";
import { userID } from "../constants/constants";

interface TransactionFormProps {
  category_id: string;
  description: string;
  type: "in" | "out";
  amount: number;
}

const TransactionForm = () => {
  // const [formData, setFormData] = useState<TransactionFormProps>();
  const queryClient = useQueryClient();

  const { mutate } = useMutation(postTransaction, {
    // mutationFn: (transaction: Transaction) => {
    //   return fetchPostTransactions(transaction);
    // },
    onSuccess: (data) => {
      console.log("data", data);
      queryClient.invalidateQueries(queryKeys[QUERY_KEYS.transaction].getTransactionsByUserID(userID).queryKey);
    },
    onError: (error) => {
      console.log("error", error);
    },
  });

  const onSubmit = (formData: TransactionFormProps) => {
    console.log("Form data submitted:", formData);
    // setFormData(data);
    const transaction: Transaction = {
      ...formData,
      user_id: "a65d8bd1-608f-448d-aa39-d01603dcf5cd",
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
      // type: "select",
      // options: [
      //   { value: 'in', name: 'In' },
      //   { value: 'out', name: 'Out' },
      // ],
      rules: {
        validate: (value: string) =>
          ["in", "out"].includes(value) || 'Type must be "in" or "out"',
      },
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
