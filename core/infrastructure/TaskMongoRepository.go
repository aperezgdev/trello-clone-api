package infrastructure

import (
	"context"
	"os"

	"github.com/aperezgdev/trello-clone-api/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskMongoRepository struct {
	collection *mongo.Collection
}

func NewTaskMongoRepository(client *mongo.Client) *TaskMongoRepository {
	return &TaskMongoRepository{
		collection: client.Database(os.Getenv("MONGO_DB_NAME")).Collection("tasks"),
	}
}

func (tr *TaskMongoRepository) Create(t domain.Task) {

	tr.collection.InsertOne(context.TODO(), t)
}

func (tr *TaskMongoRepository) FindAll() []domain.Task {
	cursor, err := tr.collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		panic(err)
	}

	var tasks []domain.Task = make([]domain.Task, 0)
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		panic(err)
	}

	return tasks
}

func (tr *TaskMongoRepository) Delete(id uint32) {
	tr.collection.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
}

func (tr *TaskMongoRepository) Update(t domain.Task) {
	tr.collection.UpdateOne(context.TODO(), bson.D{{Key: "id", Value: t.Id}}, t)
}

func (tr *TaskMongoRepository) FindByBoard(idBoard uint32) []domain.Task {
	cursor, err := tr.collection.Find(context.TODO(), bson.D{{Key: "boardStatusId", Value: idBoard}})

	if err != nil {
		panic(err)
	}

	var tasks []domain.Task = make([]domain.Task, 0)
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		panic(err)
	}

	return tasks
}
