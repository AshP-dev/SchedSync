import React, { useState } from "react";
//import "./Card.css";

const Card = ({ card }) => {
  const [flipped, setFlipped] = useState(false);

  const handleFlip = () => {
    setFlipped(!flipped);
  };

  const handleReview = (rating) => {
    const payload = JSON.stringify({ rating });
    console.log("Payload:", payload); // Log the payload

    fetch(`http://localhost:3010/api/cards/${card.id}/review`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: payload,
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Review submitted:", data);
      })
      .catch((error) => {
        console.error("Error submitting review:", error);
      });
  };
  return (
    <div className={`card ${flipped ? "flipped" : ""}`} onClick={handleFlip}>
      <div className='card-inner'>
        <div className='card-front'>
          <h2>{card.front}</h2>
        </div>
        <div className='card-back'>
          <h2>{card.front}</h2>
          <p>{card.back}</p>
          {card.image && <img src={card.image} alt={card.front} />}
          {card.link && (
            <a href={card.link} target='_blank' rel='noopener noreferrer'>
              More Info
            </a>
          )}
          <div className='review-buttons'>
            {[1, 2, 3, 4, 5].map((rating) => (
              <button key={rating} onClick={() => handleReview(rating)}>
                {rating}
              </button>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Card;
