import React, { useState } from "react";

const PayBills = () => {
  const [bill, setBill] = useState("");

  const handleBillPayment = () => {
    alert(`Paying ${bill} bill`);
  };

  return (
    <div className="bg-black min-h-screen text-white px-6 py-8">
      <h1 className="text-3xl font-bold mb-6">Pay Bills</h1>
      <div className="space-y-4">
        <select
          value={bill}
          onChange={(e) => setBill(e.target.value)}
          className="w-full bg-gray-800 text-white py-2 px-4 rounded-md focus:outline-none focus:ring focus:ring-gray-600"
        >
          <option value="" disabled>
            Select Bill Type
          </option>
          <option value="Electricity">Electricity</option>
          <option value="Water">Water</option>
          <option value="Internet">Internet</option>
        </select>
        <button
          onClick={handleBillPayment}
          className="bg-blue-500 text-white px-6 py-2 rounded-lg hover:bg-blue-600"
        >
          Pay Now
        </button>
      </div>
    </div>
  );
};

export default PayBills;
