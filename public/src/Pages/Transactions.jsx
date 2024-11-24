import React, { useEffect, useState } from "react";
import { FaArrowUp, FaArrowDown } from "react-icons/fa";
import { HiArrowLeft } from "react-icons/hi";
import { useNavigate } from "react-router-dom";
import axios from "axios";

export default function Transactions() {
  const navigate = useNavigate();
  const [transactions, setTransactions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Status color mapping
  const statusColors = {
    completed: "bg-green-600",
    pending: "bg-yellow-600",
    failed: "bg-red-600",
  };

  // Go back handler
  const handleGoBack = () => {
    navigate(-1); // Navigate back to the previous page
  };

  // Fetch transactions from the Go backend
  const fetchTransactions = async () => {
    try {
      const response = await axios.get("http://localhost:8080/payments");
      setTransactions(response.data); // Assuming the API returns an array of transactions
      setLoading(false);
    } catch (err) {
      setError("Failed to fetch transactions");
      setLoading(false);
    }
  };

  // Fetch transactions when the component mounts
  useEffect(() => {
    fetchTransactions();
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen bg-black text-white flex items-center justify-center">
        <div>Loading...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-black text-white flex items-center justify-center">
        <div>{error}</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-black text-white overflow-x-hidden flex flex-col">
      {/* Header with Back Button */}
      <div className="p-6 flex items-center space-x-4 border-b border-gray-700">
        <button
          onClick={handleGoBack}
          className="text-white text-3xl p-2 bg-gray-800 rounded-full hover:bg-gray-700"
        >
          <HiArrowLeft />
        </button>
        <h1 className="text-3xl font-bold">Your Transactions</h1>
      </div>

      {/* Status Legend */}
      <div className="p-6 flex items-center space-x-6">
        <div className="flex items-center space-x-2">
          <div className="w-3 h-3 rounded-full bg-green-600"></div>
          <span className="text-sm text-gray-400">Completed</span>
        </div>
        <div className="flex items-center space-x-2">
          <div className="w-3 h-3 rounded-full bg-yellow-600"></div>
          <span className="text-sm text-gray-400">Pending</span>
        </div>
        <div className="flex items-center space-x-2">
          <div className="w-3 h-3 rounded-full bg-red-600"></div>
          <span className="text-sm text-gray-400">Failed</span>
        </div>
      </div>

      {/* Transactions List */}
      <div className="p-6 overflow-y-auto flex-1 space-y-4">
        {transactions.map((transaction) => (
          <div
            key={transaction.id}
            className={`flex items-center p-4 rounded-lg shadow-md border border-gray-700 ${statusColors[transaction.status]} bg-opacity-20`}
          >
            {/* Icon */}
            <div className="text-2xl mr-4">
              {transaction.type === "Sent" ? (
                <FaArrowUp className="text-red-400" />
              ) : (
                <FaArrowDown className="text-green-400" />
              )}
            </div>

            {/* Transaction Details */}
            <div className="flex-1">
              <div className="text-lg font-bold">
                {transaction.name}
                <div className="text-sm text-gray-500">{transaction.publicId}</div>
              </div>
              <div className="text-sm text-gray-400 mt-1">
                {transaction.date} - {transaction.time}
              </div>
              <div className="text-sm text-gray-400">{transaction.method}</div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
