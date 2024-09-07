import React from 'react';
import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Layout from './layouts/Layout';
import { Home } from './pages/Home';
import { QueryClientProvider } from 'react-query';
import { queryClient } from './api/queryClient';

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          {/* <Route path="menu" element={<Menu />} /> */}
          {/* <Route path="blogs" element={<Blogs />} /> */}
          {/* <Route path="product/:id" element={<ProductDetail />} /> */}
          {/* <Route path="cart" element={<Cart />} /> */}
          {/* <Route path="*" element={<PageNotFound />} /> */}
        </Route>
      </Routes>
    </BrowserRouter>
  </QueryClientProvider>
    // <div className="App">
    //   <header className="App-header">
    //     <img src={logo} className="App-logo" alt="logo" />
    //     <p>
    //       Edit <code>src/App.tsx</code> and save to reload.
    //     </p>
    //     <a
    //       className="App-link"
    //       href="https://reactjs.org"
    //       target="_blank"
    //       rel="noopener noreferrer"
    //     >
    //       Learn React
    //     </a>
    //   </header>
    // </div>
  );
}

export default App;
