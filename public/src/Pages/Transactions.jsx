import React from "react";
import { FaArrowUp, FaArrowDown } from "react-icons/fa";
import { HiArrowLeft } from "react-icons/hi";
import { useNavigate } from "react-router-dom";

export default function Transactions() {
  const navigate = useNavigate();

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
    {
      id: 4,
      type: "Received",
      name: "Charlie Johnson",
      publicId: "0x4567ijkl",
      status: "completed",
      date: "2024-11-21",
      time: "12:00 PM",
      method: "Crypto",
    },
    {
      id: 5,
      type: "Sent",
      name: "Emma Wilson",
      publicId: "0x7890mnop",
      status: "pending",
      date: "2024-11-20",
      time: "9:30 AM",
      method: "UPI",
    },
    {
      id: 6,
      type: "Received",
      name: "Daniel Lee",
      publicId: "0x1357qrst",
      status: "completed",
      date: "2024-11-19",
      time: "5:00 PM",
      method: "ForEX",
    },
    {
      id: 7,
      type: "Sent",
      name: "Sophia Taylor",
      publicId: "0x2468uvwx",
      status: "failed",
      date: "2024-11-18",
      time: "8:00 AM",
      method: "Crypto",
    },
    {
      id: 8,
      type: "Received",
      name: "Mia Anderson",
      publicId: "0x3690yzab",
      status: "completed",
      date: "2024-11-17",
      time: "11:45 AM",
      method: "UPI",
    },
    {
      id: 9,
      type: "Sent",
      name: "Ethan Thomas",
      publicId: "0x4829cdef",
      status: "pending",
      date: "2024-11-16",
      time: "2:20 PM",
      method: "ForEX",
    },
    {
      id: 10,
      type: "Received",
      name: "Olivia Martinez",
      publicId: "0x5731ghij",
      status: "failed",
      date: "2024-11-15",
      time: "6:15 PM",
      method: "Crypto",
    },
  ];
  
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

  return (
    <div className="p-6 space-y-6 bg-black text-white h-full">
      {/* Header with Back Button */}
      <div className="flex items-center space-x-4 border-b border-gray-700 pb-4">
        <button
          onClick={handleGoBack}
          className="text-white text-3xl p-2 bg-gray-800 rounded-full hover:bg-gray-700"
        >
          <HiArrowLeft />
        </button>
        <h1 className="text-3xl font-bold">Your Transactions</h1>
      </div>

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
