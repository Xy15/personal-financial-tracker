import { Transaction } from "../../api/types";

export interface TableCols {
  key: string;
  header: string;
  type?: "date" | "amount";
}

type TableData = {
  [key: string]: Transaction[]; // Dynamic keys (e.g., additionalProp1, additionalProp2), each containing an array of transactions
};

export interface TableProps<T> {
  tableCols: TableCols[];
  tableData: TableData;
}

const Table = <T,>({ tableCols, tableData }: TableProps<T>) => {
  // Helper function to get the nested value
  const getNestedValue = (obj: any, keys: string[]): any => {
    return keys.reduce((acc, key) => acc && acc[key], obj);
  };

  return (
    <table className="bg-background rounded-lg table-auto overflow-hidden">
      <thead className="bg-secondary">
        <tr>
          {tableCols &&
            tableCols.length > 0 &&
            tableCols.map((col, index) => (
              <th key={index} className="px-4 py-2">
                {col.header}
              </th>
            ))}
        </tr>
      </thead>
      <tbody className="">
      {tableData ? (
          Object.keys(tableData).map((row, index) => (
            <tr key={index} className="border-b even:bg-tertiary/70">
              {tableCols.map((col) => {
                let keys = col.key.split(".");

                let value = getNestedValue(row, keys);

                if (col.type) {
                  switch (col.type) {
                    case "amount":
                      value = new Intl.NumberFormat("en-US", {
                        style: "decimal",
                        minimumFractionDigits: 2,
                        maximumFractionDigits: 2,
                      }).format(Number(value));
                      break;
                    case "date":
                      value = new Intl.DateTimeFormat("en-US", {
                        day: "2-digit",
                        month: "short",
                        year: "numeric",
                      }).format(new Date(value));
                      break;
                  }
                }

                return (
                  <td key={col.key} className="px-4 py-2">
                    {value}
                  </td>
                );
              })}
            </tr>
          ))
        ) : (
          <tr>
            <td colSpan={4}>No Data Found</td>
          </tr>
        )}
        {/* {tableData && tableData.length > 0 ? (
          tableData.map((row, index) => (
            <tr key={index} className="border-b even:bg-tertiary/70">
              {tableCols.map((col) => {
                let keys = col.key.split(".");

                let value = getNestedValue(row, keys);

                if (col.type) {
                  switch (col.type) {
                    case "amount":
                      value = new Intl.NumberFormat("en-US", {
                        style: "decimal",
                        minimumFractionDigits: 2,
                        maximumFractionDigits: 2,
                      }).format(Number(value));
                      break;
                    case "date":
                      value = new Intl.DateTimeFormat("en-US", {
                        day: "2-digit",
                        month: "short",
                        year: "numeric",
                      }).format(new Date(value));
                      break;
                  }
                }

                return (
                  <td key={col.key} className="px-4 py-2">
                    {value}
                  </td>
                );
              })}
            </tr>
          ))
        ) : (
          <tr>
            <td colSpan={4}>No Data Found</td>
          </tr>
        )} */}
      </tbody>
    </table>
  );
};

export default Table;
