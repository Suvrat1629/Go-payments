import React, { useState } from "react";

const BankTransfer = () => {
  const [bankDetails, setBankDetails] = useState({
    accountNumber: "",
    amount: "",
  });

  const handleTransfer = () => {
    alert(`Transferring ${bankDetails.amount} to account ${bankDetails.accountNumber}`);
  };

  return (
    <div className="bg-black min-h-screen text-white px-6 py-8">
      <h1 className="text-3xl font-bold mb-6">Bank Transfer</h1>
      <div className="space-y-4">
        <input
          type="text"
          placeholder="Account Number"
          value={bankDetails.accountNumber}
          onChange={(e) =>
            setBankDetails({ ...bankDetails, accountNumber: e.target.value })
          }
          className="w-full bg-gray-800 text-white py-2 px-4 rounded-md focus:outline-none focus:ring focus:ring-gray-600"
        />
        <input
          type="text"
          placeholder="Amount"
          value={bankDetails.amount}
          onChange={(e) =>
            setBankDetails({ ...bankDetails, amount: e.target.value })
          }
          className="w-full bg-gray-800 text-white py-2 px-4 rounded-md focus:outline-none focus:ring focus:ring-gray-600"
        />
        <button
          onClick={handleTransfer}
          className="bg-green-500 text-white px-6 py-2 rounded-lg hover:bg-green-600"
        >
          Transfer
        </button>
      </div>
    </div>
  );
};

export default BankTransfer;
