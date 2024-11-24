import React, { useState } from "react";

const PayPhoneNumber = () => {
  const [phone, setPhone] = useState("");

  const handlePayment = () => {
    alert(`Paying to phone number: ${phone}`);
  };

  return (
    <div className="bg-black min-h-screen text-white px-6 py-8">
      <h1 className="text-3xl font-bold mb-6">Pay via Phone Number</h1>
      <div className="space-y-4">
        <input
          type="text"
          placeholder="Enter Phone Number"
          value={phone}
          onChange={(e) => setPhone(e.target.value)}
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

export default PayPhoneNumber;
