import React from "react";
import { Link, useNavigate } from "react-router-dom";
import { HiArrowLeft } from "react-icons/hi"; // Import the left arrow icon
import { FaUniversity, FaEthereum, FaBitcoin } from "react-icons/fa"; // Import relevant icons

const Balance = () => {
  // Dummy data for balances
  const balances = [
    {
      id: 1,
      bankName: "Ayushi Bank",
      currency: "₹",
      amount: 5000,
      account: "1234567890",
      icon: <FaUniversity />,
    },
    {
      id: 2,
      bankName: "Aneesha Bank",
      currency: "$",
      amount: 300,
      account: "0987654321",
      icon: <FaUniversity />,
    },
    {
      id: 3,
      bankName: "Crypto (Ethereum)",
      currency: "ETH",
      amount: 1.5,
      account: "0xABCDEF123456",
      icon: <FaEthereum />,
    },
    {
      id: 4,
      bankName: "Crypto (Bitcoin)",
      currency: "BTC",
      amount: 0.02,
      account: "0x123456ABCDEF",
      icon: <FaBitcoin />,
    },
  ];

  // Dummy data for transactions
  const transactions = [
    { id: 1, username: "Alice", amount: -200 },
    { id: 2, username: "Bob", amount: 1000 },
    { id: 3, username: "Charlie", amount: -150 },
  ];

  // useNavigate hook to handle redirection
  const navigate = useNavigate();

  const handleGoBack = () => {
    navigate(-1); // Navigate back to the previous page
  };

  return (
    <div className="bg-black min-h-screen py-6 px-4 text-white flex justify-start">
      <div className="w-full max-w-5xl">
        {/* Go Back Icon */}
        <button
          onClick={handleGoBack}
          className="text-white text-3xl mb-4 p-2 bg-gray-800 rounded-full hover:bg-gray-700"
        >
          <HiArrowLeft />
        </button>

        {/* Balances Section */}
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-6 mb-10">
          {balances.map((balance) => (
            <div
              key={balance.id}
              className="bg-gray-800 rounded-2xl shadow-lg p-6 flex items-center"
            >
              <div className="text-4xl text-blue-400 mr-4">{balance.icon}</div>
              <div>
                <h2 className="text-2xl font-semibold text-white">
                  {balance.bankName}
                </h2>
                <p className="text-4xl font-bold text-blue-400 mt-2">
                  {balance.currency}
                  {balance.amount}
                </p>
                <p className="text-lg text-gray-400 mt-1">Account: {balance.account}</p>
              </div>
            </div>
          ))}
        </div>

        {/* Transactions Section */}
        <div className="bg-gray-800 rounded-2xl shadow-lg p-8">
          <h3 className="text-3xl font-semibold text-white mb-6">
            Recent Transactions
          </h3>
          <ul className="space-y-4">
            {transactions.map((transaction) => (
              <li key={transaction.id} className="flex justify-between items-center">
                <span className="text-xl text-gray-400">{transaction.username}</span>
                <span
                  className={`text-xl ${
                    transaction.amount > 0 ? "text-green-500" : "text-red-500"
                  } font-semibold`}
                >
                  {transaction.amount > 0 ? `+${transaction.amount}` : transaction.amount}
                </span>
              </li>
            ))}
          </ul>

          {/* View all transactions link */}
          <Link
            to="/transactions"
            className="text-blue-400 text-xl hover:underline mt-6 inline-block"
          >
            View all transactions
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Balance;
