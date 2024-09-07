import Table from "./Table";
import { getTransactionsByUserID } from "../api/transaction/transaction";
import { useQuery } from "react-query";
import { Response, Transaction } from "../api/types";
import { QUERY_KEYS, queryKeys } from "../api/queryKeys";
import { USER_ID } from "../constants/constants";

const TransactionTable = () => {
  const { isError, isLoading, data } = useQuery<Response<Transaction[]>>(
    queryKeys[QUERY_KEYS.transaction].getTransactionsByUserID(USER_ID).queryKey,
    () => getTransactionsByUserID(USER_ID),
  );

  const columns = [
    {
      key: "category.name",
      header: "Category",
    },
    {
      key: "description",
      header: "Description",
    },
    {
      key: "amount",
      header: "Amount",
    },
    {
      key: "transaction_date",
      header: "Date",
    },
  ];

  return (
    <>
      {isLoading && <div>Loading...</div>}
      {isError ? (
        <div>Error loading transaction...</div>
      ) : (
        data && <Table tableCols={columns} tableData={data.data} />
      )}
    </>
  );
};

export default TransactionTable;
