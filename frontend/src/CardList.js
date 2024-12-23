import React, { useEffect, useState } from "react";
import Card from "./Card";

const CardList = () => {
  const [cards, setCards] = useState([]);

  useEffect(() => {
    fetch("http://localhost:3010/api/cards")
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
