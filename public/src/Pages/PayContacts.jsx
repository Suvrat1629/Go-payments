import React from "react";
import { HiOutlineSearch } from "react-icons/hi"; // Icon for the search bar
import { Link } from "react-router-dom";

const PayContacts = () => {
  const contacts = [
    { id: 1, initials: "AB", name: "Alice Brown" },
    { id: 2, initials: "CD", name: "Charlie Doe" },
    { id: 3, initials: "EF", name: "Ethan Foster" },
    { id: 4, initials: "GH", name: "Grace Hall" },
    { id: 5, initials: "IJ", name: "Isla James" },
    { id: 6, initials: "KL", name: "Kevin Lee" },
    { id: 7, initials: "MN", name: "Mia Nolan" },
    { id: 8, initials: "OP", name: "Oscar Price" },
  ];

  return (
    <div className="bg-black min-h-screen text-white px-6 py-8">
      {/* Search Bar */}
      <div className="flex items-center space-x-3 mb-6">
        <HiOutlineSearch className="text-gray-400 text-2xl" />
        <input
          type="text"
          placeholder="Search contacts..."
          className="w-full bg-gray-800 text-white py-2 px-4 rounded-md focus:outline-none focus:ring focus:ring-gray-600"
        />
      </div>

      {/* Contacts Grid */}
      <div className="grid grid-cols-4 gap-6">
        {contacts.slice(0, 8).map((contact) => (
          <div
            key={contact.id}
            className="flex flex-col items-center bg-gray-800 p-4 rounded-lg shadow-md hover:bg-gray-700"
          >
            <div className="w-12 h-12 flex items-center justify-center bg-gray-600 text-white font-bold rounded-full text-xl">
              {contact.initials}
            </div>
            <span className="text-sm text-gray-300 mt-2">{contact.name}</span>
          </div>
        ))}
      </div>

      {/* View More Link */}
      <div className="text-center mt-6">
        <Link
          to="/view-more"
          className="text-blue-400 hover:underline text-lg"
        >
          View More
        </Link>
      </div>
    </div>
  );
};

export default PayContacts;
