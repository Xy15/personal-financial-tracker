import Table from "./Table";
import { getTransactionsByUserID } from "../api/transaction/transaction";
import { useQuery } from "react-query";
import { GetTransactionRes } from "../api/types";
import { QUERY_KEYS, queryKeys } from "../api/queryKeys";
import { userID } from "../constants/constants";

export interface TransactionTableProps {
  category: string;
  description: string;
  type: "in" | "out";
  amount: number;
}

const TransactionTable = () => {
  const { isError, isLoading, data } = useQuery<GetTransactionRes[]>(
    queryKeys[QUERY_KEYS.transaction].getTransactionsByUserID(userID).queryKey,
    () => getTransactionsByUserID(userID),
  );

  const columns = [
    {
      key: "column1",
      name: "Header 1",
    },
    {
      key: "column2",
      name: "Header 2",
    },
    {
      key: "column3",
      name: "Header 3",
    },
  ];

  // const tableData = [
  //   {
  //     column1: "row1 data1",
  //     column2: "row1 data2",
  //     column3: "row1 data3",
  //   },
  //   {
  //     column1: "row2 data1",
  //     column2: "row2 data2",
  //     column3: "row2 data3",
  //   },
  //   {
  //     column1: "row3 data1",
  //     column2: "row3 data2",
  //     column3: "row3 data3",
  //   },
  // ];

  return (
    <>
      {isLoading && <div>Loading...</div>}
      {isError ? (
        <div>Error loading transaction...</div>
      ) : (
        data && <Table tableCols={columns} tableData={data} />
      )}
    </>
  );
};

export default TransactionTable;
