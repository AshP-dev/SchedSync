import React, { useEffect, useState } from "react";
import Card from "./Card";

const CARD_URL = "http://localhost:3010";

const CardList = () => {
  const [cards, setCards] = useState([]);

  useEffect(() => {
    fetch(`${CARD_URL}/api/cards`)
      .then((response) => response.json())
      .then((data) => setCards(data))
      .catch((error) => console.error("Error fetching cards:", error));
  }, []);

  return (
    <div className='card-list'>
      {cards.map((card) => (
        <Card key={card.id} card={card} />
      ))}
    </div>
  );
};

export default CardList;
