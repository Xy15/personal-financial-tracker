export interface TableProps {
  tableCols: {
    key: string
    name: string
  }[];
  tableData: Record<string, any>[];
  // tableData: {
  //   [key: string]: string;
  // }[];
}

const Table = ({
  tableCols,
  tableData,
}: TableProps) => {

  return (
    <table className="bg-background rounded-lg table-auto overflow-hidden">
      <thead className="bg-secondary">
        <tr>
          {tableCols.map((col, index) => (
            <th key={index} className="px-4 py-2">{col.name}</th>  
          ))}
        </tr>
      </thead>
      <tbody className="">
        {tableData.map((row, index) => (
          <tr key={index} className="border-b even:bg-tertiary/70">
            {tableCols.map((col) => (
              <td key={col.key} className="px-4 py-2">{row[col.key]}</td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  )
}

export default Table