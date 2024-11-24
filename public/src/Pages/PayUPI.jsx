import React, { useState } from "react";

const PayUPI = () => {
  const [upiId, setUpiId] = useState("");

  const handlePayment = () => {
    alert(`Paying to ${upiId}`);
  };

  return (
    <div className="bg-black min-h-screen text-white px-6 py-8">
      <h1 className="text-3xl font-bold mb-6">Pay via UPI ID or Number</h1>
      <div className="space-y-4">
        <input
          type="text"
          placeholder="Enter UPI ID or Number"
          value={upiId}
          onChange={(e) => setUpiId(e.target.value)}
          className="w-full bg-gray-800 text-white py-2 px-4 rounded-md focus:outline-none focus:ring focus:ring-gray-600"
        />
        <button
          onClick={handlePayment}
          className="bg-green-500 text-white px-6 py-2 rounded-lg hover:bg-green-600"
        >
          Pay Now
        </button>
      </div>
    </div>
  );
};

export default PayUPI;
