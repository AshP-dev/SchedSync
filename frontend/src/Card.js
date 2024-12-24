import React, { useState } from "react";
//import "./Card.css";

const Card = ({ card }) => {
  const [flipped, setFlipped] = useState(false);

  const handleFlip = () => {
    setFlipped(!flipped);
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
        </div>
      </div>
    </div>
  );
};

export default Card;
