import React from "react";
import { Link } from "react-router-dom";

export default function Navbar() {
  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <aside className="w-64 bg-gray-800 text-white flex flex-col">
        {/* Header */}
        <div className="p-4 text-2xl font-bold border-b border-gray-700">
          Lavi Rabbit
        </div>

        {/* Menu Items */}
        <nav className="flex-1 p-4 space-y-4">
          <MenuButton label="Home" to="/home" />
          <MenuButton label="Transactions" to="/transactions" />
          <MenuButton label="Balance" to="/balance" />
        </nav>

        {/* Footer */}
        <div className="p-4 border-t border-gray-700 text-center text-sm">
          &copy; {new Date().getFullYear()} Lavi Rabbit
        </div>
      </aside>
    </div>
  );
}

// Menu Button Component
const MenuButton = ({ label, to }) => (
  <Link
    to={to}
    className="block w-full text-left px-4 py-2 rounded-md bg-gray-700 hover:bg-gray-600 focus:outline-none focus:ring focus:ring-gray-500"
  >
    {label}
  </Link>
);
