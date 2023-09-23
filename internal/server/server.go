package server

import (
	"custom-blockchain-iot-protocol/config"
	"custom-blockchain-iot-protocol/internal/core/controller/http"
	"custom-blockchain-iot-protocol/internal/core/platform"
	"custom-blockchain-iot-protocol/internal/core/repository"
	"custom-blockchain-iot-protocol/pkg"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{Config: config}
}

func (s *Server) Start() error {
	engine := gin.New()
	engine.Use(gin.Recovery())

	main := engine.Group("/api")

	ethClient, err := platform.NewClient(s.Config.EthHost)
	if err != nil {
		return err
	}

	log.Println("Connected to Ethereum node")

	defer ethClient.Close()

	log.Printf("Deploying contract from %s\n", s.Config.PrivateKey)

	auth, err := pkg.GetAccountAuth(ethClient, s.Config.PrivateKey)
	if err != nil {
		return err
	}

	log.Println("Deploying contract...")

	deployedContract, _, _, err := pkg.DeployPkg(auth, ethClient)
	if err != nil {

	}

	log.Printf("Deployed contract at %s\n", deployedContract.Hex())

	encryption := &pkg.Encryption{
		Key: []byte(s.Config.EncryptionKey),
	}

	ethRepo, err := repository.NewEthRepo(ethClient, deployedContract, encryption)
	if err != nil {
		return err
	}

	handler := http.NewMainHandler(ethRepo)
	err = http.Map(main, handler)

	if err != nil {
		return err
	}

	go func() {
		err = engine.Run()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}()

	exitChan := make(chan os.Signal, 1) // buffered channel to avoid missing signals

	signal.Notify(exitChan, syscall.SIGTERM)

	<-exitChan // block main routine until a signal is received

	return nil
}
