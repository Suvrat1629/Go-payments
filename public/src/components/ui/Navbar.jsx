import React from "react";
import { Link, useLocation } from "react-router-dom";

export default function Navbar() {
  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <aside className="w-80 bg-black text-white flex flex-col">
        {/* Header */}
        <div className="p-6 text-2xl font-bold border-b border-gray-700">
          Lavi Rabbit
        </div>

        {/* Menu Items */}
        <nav className="flex-1 p-6 space-y-4">
          <MenuButton label="Home" to="/" />
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
const MenuButton = ({ label, to }) => {
  const location = useLocation(); // Get the current location

  // Check if the current URL matches the 'to' prop to apply active class
  const isActive = location.pathname === to;

  return (
    <Link
      to={to}
      className={`block w-full text-left px-4 py-3 rounded-md 
        ${isActive ? "bg-gray-700 text-blue-400" : "bg-gray-800 text-white"}
        hover:bg-gray-700 focus:outline-none focus:ring focus:ring-gray-500`}
    >
      {label}
    </Link>
  );
};
