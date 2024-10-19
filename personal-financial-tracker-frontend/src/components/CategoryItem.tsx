import { UserCategory } from "../api/types";
import Button from "./common/Button";
import { useContext } from "react";
import { CategoryContext } from "./provider/CategoryProvider";

interface CategoryItemProps {
  userCategory: UserCategory;
}

const CategoryItem = ({
  userCategory,
}: CategoryItemProps) => {
  const { selectedCategory, onSelectCategory } = useContext(CategoryContext);

  return (
    <Button 
      className={`${selectedCategory != userCategory.id && "!bg-background"} rounded-lg grid grid-cols-3 items-center mx-2 w-1/6 h-[4rem]`} 
      onClick={() => onSelectCategory(userCategory.id)}
    >
      <img src={userCategory.category_image.image_url} alt={userCategory.name} className="max-h-[2.5rem]"/>
      <p className="col-span-2">{userCategory.name}</p>
    </Button>
  )
}

export default CategoryItem;