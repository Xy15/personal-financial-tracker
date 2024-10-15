import axios from "axios"
import { API_BASE_URL } from "../../constants/constants";

export const getUserCategoriesByUserID = async (userID: string) => {
  const response = await axios.get(`${API_BASE_URL}/user/${userID}/category`)
  const userCategories = response.data;

  console.log("User Categories: ", userCategories)
  return userCategories
}