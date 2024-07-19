package main

import (
	"app/cmd/config"
	"app/cmd/middleware"
	"app/cmd/storage"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	// env
	// ...

	// application
	// - config
	config.LoadConfig()
	cfg := &ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
	app := NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: "",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	db := storage.NewLoaderTicketCSV(a.dbFile)
	tickets, err := db.Load()
	if err != nil {
		return
	}
	rp := repository.NewRepositoryTicketMap(a.dbFile, 0)
	rp.SetTickets(tickets)

	// service ...
	sv := service.NewServiceTicketDefault(rp)
	// handler ...
	hd := handler.NewHandlerTicketDefault(sv)

	// routes

	app := a.rt

	app.Route("/api/v1", func(outerRouter chi.Router) {

		outerRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("OK"))
		})

		outerRouter.Route("/tickets", func(ticketRouter chi.Router) {

			ticketRouter.Use(middleware.LogMiddleware)

			ticketRouter.Get("/", hd.GetTotalTickets)

			ticketRouter.Get("/getByCountry/{destination}", hd.GetTicketsAmountByDestinationCountry)

			ticketRouter.Get("/getAverage/{destination}", hd.GetPercentageTicketsByDestinationCountry)
		})
	})

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	log.Printf("server running on %s", a.serverAddr)
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
