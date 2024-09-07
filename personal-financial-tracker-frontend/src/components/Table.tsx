import { Transaction } from "../api/types";

export interface TableProps<T> {
  tableCols: {
    key: string;
    header: string;
  }[];
  tableData: T[];
}

const Table = <T,>({ tableCols, tableData }: TableProps<T>) => {
  console.log("tableData", tableData);

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
        {tableData && tableData.length > 0 ? (
          tableData.map((row, index) => (
            <tr key={index} className="border-b even:bg-tertiary/70">
              {tableCols.map((col) => {
                let keys = col.key.split(".");

                // Helper function to get the nested value
                const getNestedValue = (obj: any, keys: string[]): any => {
                  return keys.reduce((acc, key) => acc && acc[key], obj);
                };

                const value = getNestedValue(row, keys)

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
      </tbody>
    </table>
  );
};

export default Table;
