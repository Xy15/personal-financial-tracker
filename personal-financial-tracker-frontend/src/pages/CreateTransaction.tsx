import CategoryProvider from "../components/provider/CategoryProvider";
import TransactionForm from "../components/TransactionForm";
import UserCategorySelection from "../components/UserCategorySelection";

const CreateTransaction = () => {
  return (
    <>
      <CategoryProvider>
        <div className="text-center">
          <UserCategorySelection />
        </div>
        <div className="flex justify-center items-center">
          <TransactionForm />
        </div>
      </CategoryProvider>
    </>
  );
};

export default CreateTransaction;
