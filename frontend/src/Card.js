// frontend/src/Card.js
import React from "react";

const Card = ({ card }) => {
  return (
    <div className='card'>
      <h2>{card.front}</h2>
      <p>{card.back}</p>
      <p>Deck: {card.deck_id}</p>
      <p>Tags: {card.tags}</p>
      <p>Due Date: {new Date(card.due_date).toLocaleDateString()}</p>
    </div>
  );
};

export default Card;
