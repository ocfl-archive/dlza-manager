package service

import (
	"context"
	"github.com/pkg/errors"
	pb "gitlab.switch.ch/ub-unibas/dlza/dlza-manager/dlzamanagerproto"
	"time"
)

type DispatcherHandlerService struct {
	ClientDispatcherHandler        pb.DispatcherHandlerServiceClient
	ClientDispatcherStorageHandler pb.DispatcherStorageHandlerServiceClient
}

func NewDispatcherHandlerService(clientDispatcherHandler pb.DispatcherHandlerServiceClient, clientDispatcherStorageHandler pb.DispatcherStorageHandlerServiceClient) *DispatcherHandlerService {
	return &DispatcherHandlerService{ClientDispatcherHandler: clientDispatcherHandler, ClientDispatcherStorageHandler: clientDispatcherStorageHandler}
}

func (d *DispatcherHandlerService) GetLowQualityCollectionsAndAct() error {
	c := context.Background()
	cont, cancel := context.WithTimeout(c, 10000*time.Second)
	defer cancel()

	collectionAliases, err := d.ClientDispatcherHandler.GetLowQualityCollections(cont, &pb.NoParam{})

	if err != nil {
		return errors.Wrapf(err, "cannot get LowQualityCollections")
	}

	if len(collectionAliases.CollectionAliases) != 0 {
		_, err := d.ClientDispatcherStorageHandler.ChangeQualityForCollections(cont, collectionAliases)
		if err != nil {
			return errors.Wrapf(err, "cannot change quality of collections with low quality")
		}
	}
	return nil
}
