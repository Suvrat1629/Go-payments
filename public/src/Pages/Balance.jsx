import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { HiArrowLeft } from 'react-icons/hi'; // Import the left arrow icon

const Balance = () => {
  // Dummy data for balance, username, and transactions with usernames
  const balance = 5000; // Current balance
  const username = 'John Doe'; // User's name
  const transactions = [
    { id: 1, username: 'Alice', amount: -200 },
    { id: 2, username: 'Bob', amount: 1000 },
    { id: 3, username: 'Charlie', amount: -150 },
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

        {/* Balance Section */}
        <div className="bg-gray-800 rounded-2xl shadow-lg p-8 mb-10">
          <div className="text-left">
            <h2 className="text-4xl font-semibold text-white">Your Balance</h2>
            <p className="text-6xl font-bold text-blue-400 mt-4">${balance}</p>
            <p className="text-xl text-gray-400 mt-4">{username}</p>
          </div>
        </div>

        {/* Transactions Section */}
        <div className="bg-gray-800 rounded-2xl shadow-lg p-8 mb-10">
          <h3 className="text-3xl font-semibold text-white mb-6">Recent Transactions</h3>
          <ul className="space-y-4">
            {transactions.map((transaction) => (
              <li key={transaction.id} className="flex justify-between items-center">
                <span className="text-xl text-gray-400">{transaction.username}</span>
                <span
                  className={`text-xl ${transaction.amount > 0 ? 'text-green-500' : 'text-red-500'} font-semibold`}
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
