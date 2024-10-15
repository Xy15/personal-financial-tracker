import { getTransactionsByUserID } from "../api/transaction/transaction";
import { useQuery } from "react-query";
import { GroupedTransactions, Response, Transaction } from "../api/types";
import { QUERY_KEYS, queryKeys } from "../api/queryKeys";
import { USER_ID } from "../constants/constants";
import ListItem from "./common/ListItem";

const TransactionList = () => {
  const { isError, isLoading, data } = useQuery<Response<GroupedTransactions>>(
    queryKeys[QUERY_KEYS.transaction].getTransactionsByUserID(USER_ID).queryKey,
    () => getTransactionsByUserID(USER_ID)
  );

  return (
    <>
      {isLoading && <>Loading...</>}
      {isError ? (
        <>Error loading transaction...</>
      ) : (
        data &&
        Object.keys(data?.data).map((date: string) => {
          const transactions: Transaction[] = data.data[date];
          return (
            <div key={date} className="">
              <h3>{date}</h3>
              {transactions.map((tx) => {
                const formattedTime = new Date(
                  tx.transaction_date
                ).toLocaleTimeString("en-GB", {
                  hour: "2-digit",
                  minute: "2-digit",
                  hour12: false,
                  timeZone: "UTC",
                });

                return (
                  <ListItem
                    key={tx.id}
                    icon={tx.category_image_url}
                    title={tx.category_name}
                    description={tx.description}
                    time={formattedTime}
                    amount={tx.amount}
                  />
                );
              })}
            </div>
          );
        })
      )}
    </>
  );
};

export default TransactionList;
