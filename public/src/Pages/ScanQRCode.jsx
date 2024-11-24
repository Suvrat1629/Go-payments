import React from "react";

const ScanQRCode = () => {
  return (
    <div className="bg-black min-h-screen text-white px-6 py-8">
      <h1 className="text-3xl font-bold mb-6">Scan QR Code</h1>
      <div className="flex flex-col items-center">
        <div className="w-64 h-64 bg-gray-800 flex items-center justify-center rounded-lg shadow-lg">
          <span className="text-gray-500 text-xl">[QR Scanner Placeholder]</span>
        </div>
        <button className="mt-6 bg-blue-500 text-white px-6 py-2 rounded-lg hover:bg-blue-600">
          Open Scanner
        </button>
      </div>
    </div>
  );
};

export default ScanQRCode;
