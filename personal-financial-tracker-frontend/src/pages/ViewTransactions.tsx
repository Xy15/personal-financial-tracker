import { getTransactionsByUserID } from "../api/transaction/transaction";
import { useQuery } from "react-query";
import { Response, Transaction } from "../api/types";
import { QUERY_KEYS, queryKeys } from "../api/queryKeys";
import { USER_ID } from "../constants/constants";
import TransactionList from "../components/TransactionList";

type GroupedTransactions = {
  [date: string]: Transaction[];
};

const ViewTransactions = () => {
  const { isError, isLoading, data } = useQuery<Response<GroupedTransactions[]>>(
    queryKeys[QUERY_KEYS.transaction].getTransactionsByUserID(USER_ID).queryKey,
    () => getTransactionsByUserID(USER_ID)
  );

  // const groupByDate = (transactions: Transaction[]) => {
  //   return transactions.reduce((acc, tx) => {
  //     const date = new Date(tx.transaction_date).toLocaleDateString("en-GB", {
  //       day: "2-digit",
  //       month: "short",
  //       year: "numeric",
  //     });

  //     if (!acc[date]) {
  //       acc[date] = [];
  //     }
  //     acc[date].push(tx);
  //     return acc;
  //   }, {} as GroupedTransactions);
  // };

  // const transactions = data?.data ?? [];
  // const groupedTransactions = groupByDate(transactions);

  const groupedTransactions = data?.data ?? [];
  return (
    <>
      {/* {Object.keys(groupedTransactions).map((date) => (
        <div key={date} className="transaction-card">
          <h2>{date}</h2>
          {groupedTransactions[date].map((tx) => (
            <div key={tx.id} className="transaction-item">
              <span>
                {new Date(tx.transaction_date).toLocaleTimeString("en-GB", {
                  hour: "2-digit",
                  minute: "2-digit",
                  hour12: false,
                })}
              </span>
              <span>{tx.description}</span>
            </div>
          ))}
        </div>
      ))} */}
    </>
  );
};

export default ViewTransactions;
