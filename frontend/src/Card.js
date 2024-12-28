import React, { useState } from "react";
//import "./Card.css";

const Card = ({ card }) => {
  const [flipped, setFlipped] = useState(false);

  const handleFlip = () => {
    setFlipped(!flipped);
  };

  const handleReview = (rating) => {
    fetch(`http://localhost:3010/api/cards/${card.id}/review`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ rating }),
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
            <p>How likely are you to make this again?</p>
            <button onClick={() => handleReview(1)}>1</button>
            <button onClick={() => handleReview(2)}>2</button>
            <button onClick={() => handleReview(3)}>3</button>
            <button onClick={() => handleReview(4)}>4</button>
            <button onClick={() => handleReview(5)}>5</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Card;
