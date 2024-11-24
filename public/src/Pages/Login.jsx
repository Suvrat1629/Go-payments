import React, { useState } from "react";

function Login() {
  return (
    <div className="min-h-screen bg-black text-white flex justify-center items-center">
      <div className="bg-gray-800 p-8 rounded-lg shadow-md w-96">
        <h2 className="text-2xl font-bold text-center mb-4">
          {isSignUp ? "Sign Up" : "Login"}
        </h2>

        {error && <div className="text-red-500 text-center mb-4">{error}</div>}
        {successMessage && <div className="text-green-500 text-center mb-4">{successMessage}</div>}

        <form onSubmit={handleSubmit} className="space-y-4">
          {isSignUp && (
            <>
              <div>
                <label className="block text-sm font-semibold" htmlFor="email">
                  Email
                </label>
                <input
                  type="email"
                  id="email"
                  className="w-full p-2 mt-1 bg-gray-700 text-white rounded-lg"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                />
              </div>

              <div>
                <label className="block text-sm font-semibold" htmlFor="accountNumber">
                  Account Number
                </label>
                <input
                  type="text"
                  id="accountNumber"
                  className="w-full p-2 mt-1 bg-gray-700 text-white rounded-lg"
                  value={accountNumber}
                  onChange={(e) => setAccountNumber(e.target.value)}
                  required
                />
              </div>
            </>
          )}

          <div>
            <label className="block text-sm font-semibold" htmlFor="username">
              Username
            </label>
            <input
              type="text"
              id="username"
              className="w-full p-2 mt-1 bg-gray-700 text-white rounded-lg"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>

          <div>
            <label className="block text-sm font-semibold" htmlFor="password">
              Password
            </label>
            <input
              type="password"
              id="password"
              className="w-full p-2 mt-1 bg-gray-700 text-white rounded-lg"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>

          <div className="flex justify-center mt-4">
            <button
              type="submit"
              className="bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700"
            >
              {isSignUp ? "Sign Up" : "Login"}
            </button>
          </div>
        </form>

        <div className="text-center mt-4">
          <button
            onClick={() => setIsSignUp(!isSignUp)}
            className="text-blue-400 hover:text-blue-500"
          >
            {isSignUp
              ? "Already have an account? Login"
              : "New user? Sign up here"}
          </button>
        </div>
      </div>
    </div>
  );
}
