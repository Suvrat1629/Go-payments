import React, { useState } from "react";

const SelfTransfer = () => {
  const [amount, setAmount] = useState("");

  const handleTransfer = () => {
    alert(`Transferring ${amount} to your linked account`);
  };

  return (
    <div className="bg-black min-h-screen text-white px-6 py-8">
      <h1 className="text-3xl font-bold mb-6">Self Transfer</h1>
      <div className="space-y-4">
        <input
          type="text"
          placeholder="Enter Amount"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          className="w-full bg-gray-800 text-white py-2 px-4 rounded-md focus:outline-none focus:ring focus:ring-gray-600"
        />
        <button
          onClick={handleTransfer}
          className="bg-blue-500 text-white px-6 py-2 rounded-lg hover:bg-blue-600"
        >
          Transfer
        </button>
      </div>
    </div>
  );
};

export default SelfTransfer;
