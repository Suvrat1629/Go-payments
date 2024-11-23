import React from "react";
import { FaArrowUp, FaArrowDown } from "react-icons/fa";

export default function Transactions() {
  // Sample transactions data
  const transactions = [
    {
      id: 1,
      type: "Sent",
      name: "John Doe",
      publicId: "0x1234abcd",
      status: "completed",
      date: "2024-11-24",
      time: "10:30 AM",
      method: "UPI",
    },
    {
      id: 2,
      type: "Received",
      name: "Alice Smith",
      publicId: "0x5678efgh",
      status: "pending",
      date: "2024-11-23",
      time: "3:15 PM",
      method: "Crypto",
    },
    {
      id: 3,
      type: "Sent",
      name: "Bob Brown",
      publicId: "0x9abc1234",
      status: "failed",
      date: "2024-11-22",
      time: "8:45 PM",
      method: "ForEX",
    },
  ];

  // Status color mapping
  const statusColors = {
    completed: "bg-green-600",
    pending: "bg-yellow-600",
    failed: "bg-red-600",
  };

  return (
    <div className="p-6 space-y-6 bg-black text-white h-full">
      {/* Header */}
      <h1 className="text-3xl font-bold border-b border-gray-700 pb-4">
        Your Transactions
      </h1>

      {/* Status Legend */}
      <div className="flex items-center space-x-6">
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

      {/* Transactions Cards */}
      <div className="grid gap-4">
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
