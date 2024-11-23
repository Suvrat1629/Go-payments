import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

// Import your components
import Login from './Pages/Login.jsx';
import Signup from './Pages/Signup.jsx';
import Home from './Pages/Home.jsx';
import Pay from './Pages/Pay.jsx';
import Receive from './Pages/Receive.jsx';
import Transactions from './Pages/Transactions.jsx';
import Balance from './Pages/Balance.jsx';
import OnRamp from './Pages/OnRamp.jsx';
import Navbar from './components/ui/Navbar.jsx';
import UserDashboard from './Pages/UserDashboard.jsx';

export default function App() {
  return (
    <Router>
      {/* Container for sidebar and main content */}
      <div className="flex h-screen">
        {/* Sidebar (Navbar) */}
        <Navbar />

        {/* Main Content Area */}
        <main className="flex-1  overflow-y-auto">
          {/* Define routes here */}
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/signup" element={<Signup />} />
            <Route path="/home" element={<Home />} />
            <Route path="/pay" element={<Pay />} />
            <Route path="/receive" element={<Receive />} />
            <Route path="/transactions" element={<Transactions />} />
            <Route path="/balance" element={<Balance />} />
            <Route path="/onramp" element={<OnRamp />} />
            <Route path="/user-dashboard" element={<UserDashboard />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}
