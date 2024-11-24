import React, { useState } from "react";

const MobileRecharge = () => {
  const [phone, setPhone] = useState("");
  const [amount, setAmount] = useState("");

  const handleRecharge = () => {
    alert(`Recharging ${phone} with ${amount}`);
  };

  return (
    <div className="bg-black min-h-screen text-white px-6 py-8">
      <h1 className="text-3xl font-bold mb-6">Mobile Recharge</h1>
      <div className="space-y-4">
        <input
          type="text"
          placeholder="Phone Number"
          value={phone}
          onChange={(e) => setPhone(e.target.value)}
          className="w-full bg-gray-800 text-white py-2 px-4 rounded-md focus:outline-none focus:ring focus:ring-gray-600"
        />
        <input
          type="text"
          placeholder="Amount"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          className="w-full bg-gray-800 text-white py-2 px-4 rounded-md focus:outline-none focus:ring focus:ring-gray-600"
        />
        <button
          onClick={handleRecharge}
          className="bg-blue-500 text-white px-6 py-2 rounded-lg hover:bg-blue-600"
        >
          Recharge
        </button>
      </div>
    </div>
  );
};

export default MobileRecharge;
