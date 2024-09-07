import React from 'react';
import TransactionForm from '../components/TransactionForm';
import TransactionTable from '../components/TransactionTable';

export const Home = () => {
  return (
    <>
      <TransactionTable />
      <TransactionForm />
    </>
  )
}
