package adapterdb

import (
	"context"
	"errors"

	"github.com/guillospy92/crabi/internal/core/domain"
	pkgmongo "github.com/guillospy92/crabi/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const nameCollection = "users"

type userMongo struct {
	ID        string `bson:"_id,omitempty"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Password  string `bson:"password"`
	Email     string `bson:"email"`
}

// MongoUserRepository implements interface UserRepositoryInterface
type MongoUserRepository struct {
	mongoClient pkgmongo.MongoDBAPI
}

// Save register user in mongo DB
func (m *MongoUserRepository) Save(ctx context.Context, user domain.UserEntity) error {
	userMongoRegister := userMongo{
		FirstName: user.FirstName,
		Email:     user.Email,
		Password:  user.Password,
		LastName:  user.LastName,
	}
	_, err := m.mongoClient.InsertOne(ctx, nameCollection, userMongoRegister)

	return err
}

// FindByEmail search unique register mongo DB
func (m *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*domain.UserEntity, error) {
	response := m.mongoClient.FindOne(ctx, nameCollection, bson.M{"email": email})

	var useMongoResponse userMongo

	if err := response.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &domain.UserEntity{}, nil
		}

		return nil, err
	}

	err := response.Decode(&useMongoResponse)
	if err != nil {
		return nil, err
	}

	return &domain.UserEntity{
		FirstName: useMongoResponse.FirstName,
		LastName:  useMongoResponse.LastName,
		Email:     useMongoResponse.Email,
		Password:  useMongoResponse.Password,
	}, nil
}

// NewMongoUserRepository MongoUserRepository
func NewMongoUserRepository(mongoClient pkgmongo.MongoDBAPI) *MongoUserRepository {
	return &MongoUserRepository{
		mongoClient: mongoClient,
	}
}
