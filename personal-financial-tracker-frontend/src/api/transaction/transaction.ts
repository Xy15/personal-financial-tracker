import axios from "axios"
import { CreateTransactionReq, Transaction } from "../types"
import { API_BASE_URL } from "../../constants/constants";

export const getTransactionsByUserID = async (userID: string) => {
  const response = await axios.get(`${API_BASE_URL}/transaction/user/${userID}`)
  const transactions = response.data;

  console.log("Transactions: ", transactions)
  return transactions
}

export const createTransaction = async (data: CreateTransactionReq) => {
  const response = await axios.post(`${API_BASE_URL}/transaction`, data)

  console.log("Post Transaction: ", response.data)
  return response.data
}

