import React from "react";
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

export default function Home() {
  // Button data for rendering
  const actions = [
    { label: "Scan any QR code", icon: <AiOutlineQrcode />, link: "/scan-qr" },
    { label: "Pay contacts", icon: <AiOutlineContacts />, link: "/pay-contacts" },
    { label: "Pay phone number", icon: <AiOutlinePhone />, link: "/pay-phone" },
    { label: "Bank transfer", icon: <MdAccountBalance />, link: "/bank-transfer" },
    { label: "Pay UPI ID or number", icon: <MdPayment />, link: "/pay-upi" },
    { label: "Self transfer", icon: <MdTransform />, link: "/self-transfer" },
    { label: "Pay bills", icon: <MdOutlineReceipt />, link: "/pay-bills" },
    { label: "Mobile recharge", icon: <MdPhoneAndroid />, link: "/mobile-recharge" },
  ];

  // Demo image data
  const demoImages = [
    "/path/to/image1.jpg",
    "/path/to/image2.jpg",
    "/path/to/image3.jpg",
    "/path/to/image4.jpg",
    "/path/to/image5.jpg",
  ];

  return (
    <div className="flex flex-col h-screen bg-gray-900 text-white">
      {/* Header */}
      <header className="p-4 bg-black flex items-center justify-between h-[81px] border-b border-gray-500">
        <h1 className="text-2xl font-bold">Pay Friends and Merchants</h1>
        <img
          src="/path/to/profile.jpg" // Replace with dynamic profile image
          alt="Profile"
          className="w-10 h-10 rounded-full"
        />
      </header>

      {/* Image Carousel */}
      <div className="p-4">
        <Carousel className="w-full max-w-5xl mx-auto">
          <CarouselContent className="-ml-1">
            {demoImages.map((image, index) => (
              <CarouselItem
                key={index}
                className="pl-1 basis-full sm:basis-1/2 md:basis-1/3 lg:basis-1/4"
              >
                <div className="p-1">
                  <Card>
                    <CardContent className="flex aspect-[16/9] items-center justify-center p-0 overflow-hidden">
                      <img
                        src={image}
                        alt={`Slide ${index + 1}`}
                        className="w-full h-full object-cover rounded-md"
                      />
                    </CardContent>
                  </Card>
                </div>
              </CarouselItem>
            ))}
          </CarouselContent>
          <CarouselPrevious />
          <CarouselNext />
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
        <p>UPI ID: priya@upi</p>
        <p>Balance: ₹0</p>
      </footer>
    </div>
  );
}
