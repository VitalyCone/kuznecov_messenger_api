package apiserver

import (
	"log"

	"github.com/VitalyCone/kuznecov_messenger_api/docs"
	//"github.com/VitalyCone/kuznecov_messenger_api/internal/app/apicontroller"
	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/apiserver/endpoints"
	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/store"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	mainPath string = "/messenger"
)

type APIServer struct {
	config *Config
	router *gin.Engine
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: gin.Default(),
	}
}

func (s *APIServer) Start() error {

	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureEndpoints()

	log.Println("starting api server on ")

	return s.router.Run()
}

func (s *APIServer) configureStore() error {
	st := store.New(store.NewConfig())
	if err := st.Open(); err != nil {
		return err
	}

	log.Print("Db is running")

	s.store = st

	return nil
}

func (s *APIServer) configureEndpoints() {
	endpoint := endpoints.NewEndpoints(s.store)

	docs.SwaggerInfo.BasePath = mainPath
	path := s.router.Group(mainPath)
	{
		//path.GET("/helloworld",endpoint.Helloworld)
		//path.GET("/", apicontroller.NewAPIController().Main)
		userPath := path.Group("user")
		{
			userPath.POST("/", endpoint.CreateUser)
			userPath.DELETE("/", endpoint.DeleteUser)
			userPath.GET("/", endpoint.GetUsers)
			userPath.PUT("/", endpoint.ModifyUser)
		}
		chatPath := path.Group("chat")
		{
			chatPath.GET("/all", endpoint.GetAllChats)
			chatPath.GET("/", endpoint.GetChats)
			chatPath.POST("/", endpoint.CreateChat)
			chatPath.DELETE("/", endpoint.DeleteChat)
		}
	}

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
