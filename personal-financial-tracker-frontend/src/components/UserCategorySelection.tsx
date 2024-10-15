import { useQuery } from "react-query";
import { Response, UserCategory } from "../api/types";
import { getUserCategoriesByUserID } from "../api/user_category/user_category";
import { USER_ID } from "../constants/constants";
import { QUERY_KEYS, queryKeys } from "../api/queryKeys";

const UserCategorySelection = () => {
  const { isError, isLoading, data } = useQuery<Response<UserCategory[]>>(
    queryKeys[QUERY_KEYS.user].getUserCategoriesByUserID(USER_ID).queryKey,
    () => getUserCategoriesByUserID(USER_ID),
  );
  
  return (
    <>
      <div className="flex">

      </div>
    </>
  )
}

export default UserCategorySelection;