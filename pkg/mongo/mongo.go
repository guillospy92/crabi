package pkgmongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connectionStringTemplate connection string mongo DB params
const connectionStringTemplate = "mongodb://%s:%s@%s:%s/"

// MongoSingleResult contains a single search result
type MongoSingleResult interface {
	Decode(v any) error
	DecodeBytes() (bson.Raw, error)
	Err() error
}

// MongoDBAPI contain all methods operate
type MongoDBAPI interface {
	FindOne(ctx context.Context, nameCollection string, filter any) MongoSingleResult
	InsertOne(ctx context.Context, nameCollection string, document any) (*mongo.InsertOneResult, error)
}

// MongoDBConfig all parameter necessary to connection mongo DB
type MongoDBConfig struct {
	Cluster        string
	UserName       string
	Password       string
	DBName         string
	Port           string
	ConnectTimeOut time.Duration
}

// MongoDBClient struct que implements la interface MongoDBAPI
type MongoDBClient struct {
	client   *mongo.Client
	database *mongo.Database
	context  context.Context
}

// NewInstanceMongoDBClient new instance of MongoDBAPI
func NewInstanceMongoDBClient(mongoDBConfig MongoDBConfig) MongoDBAPI {
	connectURI := fmt.Sprintf(connectionStringTemplate,
		mongoDBConfig.UserName,
		mongoDBConfig.Password,
		mongoDBConfig.Cluster,
		mongoDBConfig.Port,
	)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDBConfig.ConnectTimeOut)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Errorf("error connecting to MongoDB: %w", err))
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		panic(fmt.Errorf("error ping to MongoDB: %w", err))
	}

	return &MongoDBClient{
		client:   client,
		context:  ctx,
		database: client.Database(mongoDBConfig.DBName),
	}
}

// FindOne search unique register mongo DB
func (d *MongoDBClient) FindOne(ctx context.Context, nameCollection string, filter any) MongoSingleResult {
	return d.database.Collection(nameCollection).FindOne(ctx, filter)
}

// InsertOne insert unique register mongo DB
func (d *MongoDBClient) InsertOne(
	ctx context.Context,
	nameCollection string,
	document any,
) (*mongo.InsertOneResult, error) {
	return d.database.Collection(nameCollection).InsertOne(ctx, document)
}
