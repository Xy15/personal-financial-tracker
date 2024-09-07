import React from 'react';
import Table from '../components/Table';
import Form from '../components/Form';
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
