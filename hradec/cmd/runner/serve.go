package runner

import (
	"hradec/internal/config"
	"hradec/internal/handlers"
	"hradec/internal/ports/database"
	"hradec/internal/usecases"

	"encoding/json"
	"hradec/internal/setup"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nextap-solutions/goNextService"
	"github.com/nextap-solutions/goNextService/components"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

type HradecServerComponents struct {
	httpServer goNextService.Component
	cleanup    goNextService.Component
}

func Serve() error {
	configuration := config.LoadConfig()
	hradecApi, err := setupService(configuration)
	if err != nil {
		return err
	}
	app := goNextService.NewApplications(hradecApi.httpServer)
	return app.Run()
}

func setupService(configuration *config.Config) (*HradecServerComponents, error) {
	logger := setup.InitLogger(*configuration)
	s, err := json.MarshalIndent(configuration, "", "\t")
	if err != nil {
		logger.Error("Failed to marshal configuration", zap.Error(err))
		return nil, err
	}

	logger.Info("Logger initialized successfully")
	logger.Info(string(s))
	dbconn, err := setup.SetupDb(configuration)
	if err != nil {
		return nil, err
	}
	placeStore := database.NewDatabasePlaceStore(dbconn)
	placeUsecase := usecases.NewPlaceUsecase(placeStore)

	placeHandler := handlers.NewPlaceHandler(placeUsecase)

	router := mux.NewRouter()
	router.Handle("/", placeHandler.Ping()).Methods("GET")
	router.Handle("/places", placeHandler.GetPlacesByViewport()).Methods("GET")
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsMiddleware.Handler(router)
	api := http.Server{
		Addr:         "0.0.0.0:" + configuration.Server.Port,
		ReadTimeout:  configuration.Server.ReadTimeout,
		WriteTimeout: configuration.Server.WriteTimeout,
		Handler:      handler,
	}
	httpComponent := components.NewHttpComponent(handler, components.WithHttpServer(&api))
	var lifecycleRun components.LifeCycleFunc

	return &HradecServerComponents{
		httpServer: httpComponent,
		cleanup: components.NewLifecycleComponent([]components.LifeCycleFunc{},
			lifecycleRun, nil),
	}, nil
}
