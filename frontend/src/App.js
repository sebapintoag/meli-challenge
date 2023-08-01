import React from 'react';
import './App.css';
import './assets/styles.css';
import Navbar from './components/Navbar.js';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/home';
import About from './pages/about';
import Create from './pages/create';
import Delete from './pages/delete';
import Find from './pages/find';

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route exact path="/" element={<Home />} />
        <Route path="/create" element={<Create />} />
        <Route path="/find" element={<Find />} />
        <Route path="/delete" element={<Delete />} />
        <Route path="/about" element={<About />} />
      </Routes>
    </Router>
  );
}

export default App;
