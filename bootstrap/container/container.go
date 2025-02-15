package container

import (
	"net/http"
	"time"

	"github.com/guillospy92/clientHttp/gohttp"
	"github.com/guillospy92/crabi/internal/adapters/api"
	"github.com/guillospy92/crabi/internal/adapters/db"
	"github.com/guillospy92/crabi/internal/adapters/http/handlers"
	"github.com/guillospy92/crabi/internal/app/usecases"
	"github.com/guillospy92/crabi/pkg/mongo"
	"github.com/guillospy92/crabi/resources"
)

// CreateUserHandler global var create user handler
var CreateUserHandler *handlers.CreateUserHandler

// AuthUserHandler  global var auth user handler
var AuthUserHandler *handlers.AuthUserHandler

// GetUserInfoHandler global var get user info handler
var GetUserInfoHandler *handlers.GetUserInfoHandler

// InitializeServiceContainer initializes the container to instantiate all properties within it
func InitializeServiceContainer() {
	clientHTTP := gohttp.NewClient().SetHeaders(http.Header{"Content-Type": []string{"application/json"}}).Build()
	userBlackList := adapterapi.NewUserBlackListAdapter(clientHTTP)
	mongoClient := pkgmongo.NewInstanceMongoDBClient(pkgmongo.MongoDBConfig{
		Cluster:        resources.ConfigurationEnv().DBHostMongo,
		UserName:       resources.ConfigurationEnv().DBUserMongo,
		Password:       resources.ConfigurationEnv().DBPassMongo,
		Port:           resources.ConfigurationEnv().DBPortMongo,
		DBName:         resources.ConfigurationEnv().DBNameDataBaseMongo,
		ConnectTimeOut: time.Duration(resources.ConfigurationEnv().DBTimeOutMongo) * time.Second,
	})

	userRepository := adapterdb.NewMongoUserRepository(mongoClient)

	createUserUseCase := usecases.NewUserCreateUseCase(userBlackList, userRepository)
	authUserUseCase := usecases.NewUserAuthUseCase(userRepository)
	getInfoUserUseCase := usecases.NewUserGetUseCaseInfoCase(userRepository)

	CreateUserHandler = handlers.NewCreateUserHandler(createUserUseCase)
	AuthUserHandler = handlers.NewAuthUserHandler(authUserUseCase)
	GetUserInfoHandler = handlers.NewGetUserInfoHandler(getInfoUserUseCase)
}
