import React, { useState } from "react";

const CARD_URL = "http://localhost:3010";

function CreateCard() {
  const [front, setFront] = useState("");
  const [back, setBack] = useState("");

  async function createNewCard() {
    await fetch(`${CARD_URL}/api/cards`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ front, back, deck_id: "Deck 1" }),
    });
  }

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        gap: "10px",
        maxWidth: "300px",
        margin: "0 auto",
      }}
    >
      <input
        type='text'
        placeholder='Front'
        value={front}
        onChange={(e) => setFront(e.target.value)}
        style={{
          padding: "10px",
          fontSize: "16px",
          borderRadius: "5px",
          border: "1px solid #ccc",
        }}
      />
      <input
        type='text'
        placeholder='Back'
        value={back}
        onChange={(e) => setBack(e.target.value)}
        style={{
          padding: "10px",
          fontSize: "16px",
          borderRadius: "5px",
          border: "1px solid #ccc",
        }}
      />
      <button
        onClick={createNewCard}
        style={{
          padding: "10px",
          fontSize: "16px",
          borderRadius: "5px",
          border: "none",
          backgroundColor: "#1DA1F2",
          color: "#fff",
          cursor: "pointer",
        }}
      >
        Create Card
      </button>
    </div>
  );
}

export default CreateCard;
