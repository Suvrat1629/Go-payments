import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

// Import your components
import Login from './Pages/Login.jsx';
import Signup from './Pages/Signup.jsx';
import Home from './Pages/Home.jsx';
import PayContacts from './Pages/PayContacts.jsx';
import ScanQRCode from './Pages/ScanQRCode.jsx';
import PayUPI from './Pages/PayUPI.jsx';
import PayPhoneNumber from './Pages/PayPhoneNumber.jsx';
import SelfTransfer from './Pages/SelfTransfer.jsx';
import PayBills from './Pages/PayBills.jsx';
import BankTransfer from './Pages/BankTransfer.jsx';
import MobileRecharge from './Pages/MobileRecharge.jsx';
import Navbar from './components/ui/Navbar.jsx';
import Transactions from './Pages/Transactions.jsx';
import Balance from './Pages/Balance.jsx'

export default function App() {
  return (
    <Router>
      {/* Container for sidebar and main content */}
      <div className="flex h-screen">
        {/* Sidebar (Navbar) */}
        <Navbar />

        {/* Main Content Area */}
        <main className="flex-1 overflow-y-auto">
        <main className="flex-1 overflow-y-auto">
          {/* Define routes here */}
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/signup" element={<Signup />} />
            <Route path="/" element={<Home />} />
            <Route path="/pay-contacts" element={<PayContacts />} />
            <Route path="/show-qr" element={<ScanQRCode />} />
            <Route path="/pay-upi" element={<PayUPI />} />
            <Route path="/pay-phone" element={<PayPhoneNumber />} />
            <Route path="/self-transfer" element={<SelfTransfer />} />
            <Route path="/pay-bills" element={<PayBills />} />
            <Route path="/bank-transfer" element={<BankTransfer />} />
            <Route path="/mobile-recharge" element={<MobileRecharge />} />
            <Route path="/transactions" element={<Transactions />} />
            <Route path="/balance" element={<Balance />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}