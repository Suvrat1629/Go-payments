import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import {
  AiOutlineQrcode,
  AiOutlineContacts,
  AiOutlinePhone,
} from "react-icons/ai";
import {
  MdAccountBalance,
  MdPayment,
  MdTransform,
  MdOutlineReceipt,
  MdPhoneAndroid,
} from "react-icons/md";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel"; // Ensure you have these components in your project
import { Card, CardContent } from "@/components/ui/card"; // Ensure this as well
import { Line } from "react-chartjs-2";
import { Chart as ChartJS, LineElement, CategoryScale, LinearScale, PointElement } from "chart.js";

ChartJS.register(LineElement, CategoryScale, LinearScale, PointElement);

export default function Home() {
  const [cryptoData, setCryptoData] = useState([]);

  // Coins to display
  const coins = ["bitcoin", "ethereum", "cardano"]; // Replace with your chosen coins

  // Fetch data for each coin
  useEffect(() => {
    const fetchData = async () => {
      const promises = coins.map((coin) =>
        fetch(
          `https://api.coingecko.com/api/v3/coins/${coin}/market_chart?vs_currency=usd&days=1`
        )
          .then((res) => res.json())
          .then((data) => ({
            name: coin,
            prices: data.prices.map(([timestamp, price]) => ({
              time: new Date(timestamp).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
              price,
            })),
          }))
      );
      const result = await Promise.all(promises);
      setCryptoData(result);
    };

    fetchData();
  }, []);

  // Prepare chart data
  const getChartData = (data) => ({
    labels: data.map((point) => point.time),
    datasets: [
      {
        label: "Price (USD)",
        data: data.map((point) => point.price),
        borderColor: "rgba(75,192,192,1)",
        backgroundColor: "rgba(75,192,192,0.2)",
        tension: 0.3, // Smooth lines
      },
    ],
  });

  // Button data for rendering
  const actions = [
    { label: "Scan any QR code", icon: <AiOutlineQrcode />, link: "/show-qr" },
    { label: "Pay contacts", icon: <AiOutlineContacts />, link: "/pay-contacts" },
    { label: "Pay phone number", icon: <AiOutlinePhone />, link: "/pay-phone" },
    { label: "Bank transfer", icon: <MdAccountBalance />, link: "/bank-transfer" },
    { label: "Pay UPI ID or number", icon: <MdPayment />, link: "/pay-upi" },
    { label: "Self transfer", icon: <MdTransform />, link: "/self-transfer" },
    { label: "Pay bills", icon: <MdOutlineReceipt />, link: "/pay-bills" },
    { label: "Mobile recharge", icon: <MdPhoneAndroid />, link: "/mobile-recharge" },
  ];

  // Demo image data (for carousel)
  const demoImages = [
    "/path/to/image1.jpg",
    "/path/to/image2.jpg",
    "/path/to/image3.jpg",
    "/path/to/image4.jpg",
    "/path/to/image5.jpg",
  ];

  return (
    <div className="flex flex-col h-screen bg-black text-white">
      {/* Header */}
      <header className="p-4 bg-black flex items-center justify-between h-[81px] border-b border-gray-500">
        <h1 className="text-2xl font-bold">Crypto Dashboard</h1>
        <Link to="/user-dashboard">
          <img
            src="/path/to/profile.jpg" // Replace with dynamic profile image
            alt="Profile"
            className="w-10 h-10 rounded-full cursor-pointer"
          />
        </Link>
      </header>

      {/* Crypto Carousel */}
      <div className="p-4 relative">
        <Carousel className="w-full mx-auto h-[400px] overflow-hidden relative">
          <CarouselContent className="flex h-full">
            {cryptoData.map((crypto, index) => (
              <CarouselItem
                key={index}
                className="flex-shrink-0 w-full h-full"
              >
                <div className="flex items-center justify-center bg-gray-800 h-full rounded-md p-4 shadow-lg">
                  <div className="w-full h-full">
                    <h2 className="text-lg font-semibold mb-2 text-center capitalize">
                      {crypto.name}
                    </h2>
                    <div className="h-[350px] w-full">
                      <Line
                        data={getChartData(crypto.prices)}
                        options={{
                          responsive: true,
                          maintainAspectRatio: false,
                          scales: {
                            x: {
                              ticks: { color: "#FFF" },
                              grid: { color: "rgba(255, 255, 255, 0.1)" },
                            },
                            y: {
                              ticks: { color: "#FFF" },
                              grid: { color: "rgba(255, 255, 255, 0.1)" },
                            },
                          },
                          plugins: {
                            legend: {
                              display: false,
                            },
                          },
                        }}
                      />
                    </div>
                  </div>
                </div>
              </CarouselItem>
            ))}
          </CarouselContent>
          
          <CarouselPrevious
            className="absolute left-4 top-1/2 transform -translate-y-1/2 z-10 bg-gray-700 p-2 rounded-full cursor-pointer hover:bg-gray-600"
          />
          <CarouselNext
            className="absolute right-4 top-1/2 transform -translate-y-1/2 z-10 bg-gray-700 p-2 rounded-full cursor-pointer hover:bg-gray-600"
          />
        </Carousel>
      </div>

      {/* Actions Section */}
      <main className="flex-1 p-4 grid grid-cols-2 sm:grid-cols-4 gap-4">
        {actions.map((action, index) => (
          <Link
            key={index}
            to={action.link}
            className="flex flex-col items-center justify-center bg-gray-700 text-blue-400 p-4 rounded-lg shadow-md hover:bg-gray-600 transition transform hover:scale-105"
          >
            <div className="text-4xl mb-2">{action.icon}</div>
            <span className="text-sm text-gray-300 text-center">{action.label}</span>
          </Link>
        ))}
      </main>

      {/* Footer */}
      <footer className="p-4 bg-black text-center text-sm">
        <p>&copy; 2024 Crypto Dashboard</p>
      </footer>
    </div>
  );
}
