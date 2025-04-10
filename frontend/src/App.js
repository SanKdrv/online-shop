import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import Home from './components/Home';
import About from './components/About';
import Admin from './components/Admin';
import './App.css';
import AdminBrand from "./components/repos/AdminBrand";
import AdminCartContent from "./components/repos/AdminCartContent";
import AdminCategory from "./components/repos/AdminCategory";
import AdminOrder from "./components/repos/AdminOrder";
import AdminOrderContent from "./components/repos/AdminOrderContent";
import AdminProduct from "./components/repos/AdminProduct";
import AdminProductImage from "./components/repos/AdminProductImage";
import AdminRefreshSession from "./components/repos/AdminRefreshSession";
import AdminUser from "./components/repos/AdminUser";

const App = () => {
  return (
      <Router>
        <nav>
          <ul>
            <li>
              <Link to="/">Домашняя страница</Link>
            </li>
            <li>
              <Link to="/about">О нас</Link>
            </li>
          </ul>
        </nav>

        <Routes>
          <Route path="/" element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path="/admin" element={<Admin />} />
            <Route path="/brand" element={<AdminBrand />} />
            <Route path="/cart-content" element={<AdminCartContent />} />
            <Route path="/category" element={<AdminCategory />} />
            <Route path="/order" element={<AdminOrder />} />
            <Route path="/order-content" element={<AdminOrderContent />} />
            <Route path="/product" element={<AdminProduct />} />
            <Route path="/product-image" element={<AdminProductImage />} />
            <Route path="/refresh-session" element={<AdminRefreshSession />} />
            <Route path="/user" element={<AdminUser />} />
        </Routes>
      </Router>
  );
};

export default App;