package main

import (
	"emperror.dev/emperror"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	pb "gitlab.switch.ch/ub-unibas/dlza/dlza-manager/dlzamanagerproto"
	"gitlab.switch.ch/ub-unibas/dlza/dlza-manager/pkg/dispatcher/configuration"
	"gitlab.switch.ch/ub-unibas/dlza/dlza-manager/pkg/dispatcher/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/fs"
	"log"
	"os"
	"time"
)

var configParam = flag.String("config", "", "config file in toml format")

func main() {
	flag.Parse()
	var cfgData []byte
	if *configParam == "" {
		cfgDataS, err := fs.ReadFile(configuration.ConfigFS, "config.toml")
		if err != nil {
			emperror.Panic(errors.Wrapf(err, "cannot read configFile %v", "config.toml"))
		}
		cfgData = cfgDataS
	} else {
		cfgDataS, err := os.ReadFile(*configParam)
		if err != nil {
			emperror.Panic(errors.Wrapf(err, "cannot read %s", *configParam))
		}
		cfgData = cfgDataS
	}
	configObj := &configuration.Config{}
	if err := toml.Unmarshal(cfgData, configObj); err != nil {
		panic(errors.Wrap(err, "cannot unmarshal toml"))
	}
	connectionDispatcherHandler, err := grpc.Dial("localhost"+configObj.PortHandler, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	connectionDispatcherStorageHandler, err := grpc.Dial("localhost"+configObj.PortStorageHandler, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	dispatcherHandlerServiceClient := pb.NewDispatcherHandlerServiceClient(connectionDispatcherHandler)
	dispatcherStorageHandlerServiceClient := pb.NewDispatcherStorageHandlerServiceClient(connectionDispatcherStorageHandler)

	dispatcherHandlerService := service.NewDispatcherHandlerService(dispatcherHandlerServiceClient, dispatcherStorageHandlerServiceClient)

	for {
		err = dispatcherHandlerService.GetLowQualityCollectionsAndAct()
		if err != nil {
			log.Printf("error in GetLowQualityCollectionsAndAct methos: %v", err)
		}
		time.Sleep(configObj.CycleLength * time.Second)
		fmt.Print("Check\n")
	}

}
