package repositories

import (
	"context"
	"time"

	"schedsync/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCardRepository struct {
	collection *mongo.Collection
}

func NewMongoCardRepository(collection *mongo.Collection) *MongoCardRepository {
	return &MongoCardRepository{collection: collection}
}

func (r *MongoCardRepository) CreateCard(card models.Card) (string, error) {
	card.CreatedAt = time.Now()
	card.UpdatedAt = time.Now()
	result, err := r.collection.InsertOne(context.Background(), card)
	if err != nil {
		return "", err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", err
	}
	return oid.Hex(), nil
}

func (r *MongoCardRepository) DeleteCard(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *MongoCardRepository) GetCardByID(cardID string) (models.Card, error) {
	var card models.Card
	objID, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		return card, err
	}
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&card)
	return card, err
}

func (r *MongoCardRepository) GetCards(deckID string, tag string, dueDate string) ([]models.Card, error) {
	filter := bson.M{}
	if deckID != "" {
		filter["deck_id"] = deckID
	}
	if tag != "" {
		filter["tags"] = bson.M{"$regex": tag}
	}
	if dueDate != "" {
		filter["due_date"] = dueDate
	}

	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var cards []models.Card
	for cursor.Next(context.Background()) {
		var card models.Card
		if err := cursor.Decode(&card); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (r *MongoCardRepository) ReviewCard(cardID string, rating int) (models.Card, error) {
	update := bson.M{
		"$set": bson.M{
			"reviewed_at": time.Now(),
		},
	}
	objID, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		return models.Card{}, err
	}
	_, err = r.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return models.Card{}, err
	}
	return r.GetCardByID(cardID)
}

func (r *MongoCardRepository) UpdateCard(cardID string, card models.Card) (models.Card, error) {
	card.UpdatedAt = time.Now()
	update := bson.M{
		"$set": card,
	}
	objID, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		return models.Card{}, err
	}
	_, err = r.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return models.Card{}, err
	}
	return r.GetCardByID(cardID)
}
