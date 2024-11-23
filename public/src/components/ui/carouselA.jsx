import React from 'react';

export const Carousel = ({ children }) => {
  return <div className="carousel-container">{children}</div>;
};

export const CarouselContent = ({ children }) => {
  return <div className="carousel-content">{children}</div>;
};

export const CarouselItem = ({ children }) => {
  return <div className="carousel-item">{children}</div>;
};

export const CarouselNext = () => {
  return <button className="carousel-next">Next</button>;
};

export const CarouselPrevious = () => {
  return <button className="carousel-prev">Previous</button>;
};
