import React from "react";
import { Link } from "react-router-dom";

export default function UserDashboard() {
  return (
    <div className="flex flex-col h-screen bg-black text-white p-4">
      {/* Header */}
      <header className="flex items-center justify-between p-4 border-b border-gray-500">
        <h1 className="text-2xl font-bold">User Dashboard</h1>
        <Link to="/home">
          <button className="text-blue-400 hover:text-blue-600">Go Back</button>
        </Link>
      </header>

      {/* User Info Card */}
      <div className="flex justify-center items-center flex-col flex-grow mt-12">
        <div className="bg-gray-700 p-8 rounded-lg shadow-lg h-[450px] w-[400px]">
          <div className="flex justify-center mb-6">
            <img
              src="/path/to/profile.jpg" // Replace with dynamic profile image
              alt="User Profile"
              className="w-36 h-36 mt-8 rounded-full border-4 border-gray-500"
            />
          </div>
          <div className="text-center text-white">
            <h2 className="text-3xl font-semibold mb-2">John Doe</h2>
            <p className="text-lg text-gray-300">johndoe@example.com</p>
            {/* Show QR Code Button wrapped in Link */}
            <Link to="/show-qr">
              <button className="mt-12 px-6 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-400">
                Show QR Code
              </button>
            </Link>
          </div>
        </div>
      </div>

      {/* Footer */}
      <footer className="mt-auto p-4 bg-black text-center text-sm">
        <p>&copy; 2024 Lavi Rabbit</p>
      </footer>
    </div>
  );
}
