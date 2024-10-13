import TransactionForm from "../components/TransactionForm";
import TransactionList from "../components/TransactionList";

export const Home = () => {
  return (
    <div className="p-10">
      <div className="flex justify-center items-center">
        <TransactionForm />
      </div>
      
      <TransactionList />
    </div>
  );
};
