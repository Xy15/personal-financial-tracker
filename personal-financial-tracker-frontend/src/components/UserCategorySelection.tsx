import { useQuery } from "react-query";
import { Response, UserCategory } from "../api/types";
import { getUserCategoriesByUserID } from "../api/user_category/user_category";
import { USER_ID } from "../constants/constants";
import { QUERY_KEYS, queryKeys } from "../api/queryKeys";
import CategoryItem from "./CategoryItem";
import { useContext } from "react";
import Button from "./common/Button";
import { CategoryContext } from "./provider/CategoryProvider";

const UserCategorySelection = () => {
  const { isError, isLoading, data } = useQuery<Response<UserCategory[]>>(
    queryKeys[QUERY_KEYS.user].getUserCategoriesByUserID(USER_ID).queryKey,
    () => getUserCategoriesByUserID(USER_ID)
  );

  const {
    categoryType,
    onChangeCategoryType,
  } = useContext(CategoryContext);

  if (isLoading) return <div>Loading...</div>;
  if (isError) return <div>Error loading categories</div>;

  return (
    <>
      <Button onClick={onChangeCategoryType}>{categoryType}</Button>
      <div className="flex justify-center items-center my-4">
        {data &&
          data?.data.map(
            (userCategory) =>
              userCategory.type === categoryType && (
                <CategoryItem
                  key={userCategory.id}
                  userCategory={userCategory}
                />
              )
          )}
      </div>
    </>
  );
};

export default UserCategorySelection;
