import React from "react";
import { Link } from "react-router-dom";
import {
  AiOutlineQrcode,
  AiOutlineContacts,
  AiOutlinePhone,
} from "react-icons/ai";
import {
  MdAccountBalance,
  MdPayment,
  MdTransform,
  MdOutlineReceipt,
  MdPhoneAndroid,
} from "react-icons/md";

export default function Home() {
  // Button data for rendering
  const actions = [
    { label: "Scan any QR code", icon: <AiOutlineQrcode />, link: "/show-qr" },
    { label: "Pay contacts", icon: <AiOutlineContacts />, link: "/pay-contacts" },
    { label: "Pay phone number", icon: <AiOutlinePhone />, link: "/pay-phone" },
    { label: "Bank transfer", icon: <MdAccountBalance />, link: "/bank-transfer" },
    { label: "Pay UPI ID or number", icon: <MdPayment />, link: "/pay-upi" },
    { label: "Self transfer", icon: <MdTransform />, link: "/self-transfer" },
    { label: "Pay bills", icon: <MdOutlineReceipt />, link: "/pay-bills" },
    { label: "Mobile recharge", icon: <MdPhoneAndroid />, link: "/mobile-recharge" },
  ];

  return (
    <div className="flex flex-col h-screen bg-black text-white">
      {/* Header */}
      <header className="p-4 bg-black flex items-center justify-between h-[81px] border-b border-gray-500">
        <h1 className="text-2xl font-bold">Pay Friends and Merchants</h1>
        <Link to="/user-dashboard">
          <img
            src="/path/to/profile.jpg" // Replace with dynamic profile image
            alt="Profile"
            className="w-10 h-10 rounded-full cursor-pointer"
          />
        </Link>
      </header>

      {/* Search Bar */}
      <div className="p-4">
        <input
          type="text"
          placeholder="Search..."
          className="w-full p-3 rounded-md bg-gray-800 text-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-600"
        />
      </div>

      {/* Actions Section */}
      <main className="flex-1 p-4 grid grid-cols-2 sm:grid-cols-4 gap-4">
        {actions.map((action, index) => (
          <Link
            key={index}
            to={action.link}
            className="flex flex-col items-center justify-center bg-gray-700 text-blue-400 p-4 rounded-lg shadow-md hover:bg-gray-600 transition transform hover:scale-105"
          >
            <div className="text-4xl mb-2">{action.icon}</div>
            <span className="text-sm text-gray-300 text-center">{action.label}</span>
          </Link>
        ))}
      </main>

      {/* Footer */}
      <footer className="p-4 bg-black text-center text-sm">
        <p>UPI ID: priya@upi</p>
        <p>Balance: ₹0</p>
      </footer>
    </div>
  );
}
