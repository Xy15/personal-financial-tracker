import { createContext, Dispatch, SetStateAction, useEffect, useState } from "react";
import { CategoryType } from "../../api/types";
import UserCategorySelection from "../UserCategorySelection";

interface CategoryContextType {
  selectedCategory: string;
  // setSelectedCategory: Dispatch<SetStateAction<string>>;
  onSelectCategory: (categoryID: string) => void;  // Update this line
  categoryType: CategoryType;
  // setCategoryType: Dispatch<SetStateAction<CategoryType>>;
  onChangeCategoryType: () => void;
}

export const CategoryContext = createContext<CategoryContextType>({
  selectedCategory: "",
  onSelectCategory: () => {},
  categoryType: "Income",
  onChangeCategoryType: () => {},
});

const CategoryProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [selectedCategory, setSelectedCategory] = useState<string>("");
  const [categoryType, setCategoryType] = useState<CategoryType>("Income");

  const onSelectCategory = (categoryID: string) => {
    setSelectedCategory(categoryID);
  }

  const onChangeCategoryType = () => {
    setCategoryType(prevType => (prevType === "Income" ? "Expense" : "Income"));
  };

  useEffect(() => {
    setSelectedCategory("");
  }, [categoryType])

  return (
    <CategoryContext.Provider value={{ selectedCategory, onSelectCategory, categoryType, onChangeCategoryType }}>
      {children}
    </CategoryContext.Provider>
  );
};

export default CategoryProvider