export interface Transaction {
  category_id: string;       // CategoryID in Go, but snake_case in JSON
  description?: string;      // Optional field (as a pointer in Go)
  type: string;              // Required
  amount: number;            // float64 in Go translates to number in TypeScript
  user_id: string;           // UserID in Go, but snake_case in JSON
}

export interface GetTransactionRes {
  category_id: string;       // CategoryID in Go, but snake_case in JSON
  description?: string;      // Optional field (as a pointer in Go)
  type: string;              // Required
  amount: string;            // float64 in Go translates to number in TypeScript
}