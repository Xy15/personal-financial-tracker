import React from 'react';
import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Layout from './layouts/Layout';
import { Home } from './pages/Home';
import { QueryClientProvider } from 'react-query';
import { queryClient } from './api/queryClient';
import CreateTransaction from './pages/CreateTransaction';

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="create_transaction" element={<CreateTransaction />} />
          {/* <Route path="blogs" element={<Blogs />} /> */}
          {/* <Route path="product/:id" element={<ProductDetail />} /> */}
          {/* <Route path="cart" element={<Cart />} /> */}
          {/* <Route path="*" element={<PageNotFound />} /> */}
        </Route>
      </Routes>
    </BrowserRouter>
  </QueryClientProvider>
  );
}

export default App;
