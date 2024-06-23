package infrastructure

import (
	"context"
	"os"

	"github.com/aperezgdev/trello-clone-api/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatusBoardMongoRepository struct {
	collection *mongo.Collection
}

func NewStatusBoardMongoRepository(client *mongo.Client) *StatusBoardMongoRepository {
	return &StatusBoardMongoRepository{collection: client.Database(os.Getenv("MONGO_DB_NAME")).Collection("status_boards")}
}

func (sm *StatusBoardMongoRepository) Create(sb domain.StatusBoard) {
	sm.collection.InsertOne(context.TODO(), sb)
}

func (sm *StatusBoardMongoRepository) Delete(id uint32) {
	sm.collection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
}

func (sm *StatusBoardMongoRepository) Update(sb domain.StatusBoard) {
	sm.collection.UpdateOne(context.TODO(), bson.D{{Key: "id", Value: sb.Id}}, sb)
}

func (sm *StatusBoardMongoRepository) FindByUser(idUser uint32) []domain.StatusBoard {
	cursor, err := sm.collection.Find(context.TODO(), bson.D{{Key: "userId", Value: idUser}})

	if err != nil {
		panic(err)
	}

	var statusBoards []domain.StatusBoard = make([]domain.StatusBoard, 0)
	if err = cursor.All(context.TODO(), &statusBoards); err != nil {
		panic(err)
	}

	return statusBoards
}
