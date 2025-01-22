package tests

import (
	pb "github.com/ocfl-archive/dlza-manager/dlzamanagerproto"
	"testing"

	"github.com/ocfl-archive/dlza-manager/models"
	"github.com/ocfl-archive/dlza-manager/service"
)

func TestGetCheapestStorageLocationsForQuality(t *testing.T) {
	minQuality := 66
	storageLocation1 := models.StorageLocation{Quality: 34, Price: 99}
	storageLocation2 := models.StorageLocation{Quality: 22, Price: 67}
	storageLocation3 := models.StorageLocation{Quality: 10, Price: 40}
	storageLocation4 := models.StorageLocation{Quality: 27, Price: 56}
	storageLocation5 := models.StorageLocation{Quality: 68, Price: 200}

	storageLocations := make([]models.StorageLocation, 0)
	storageLocations = append(storageLocations, storageLocation1, storageLocation2, storageLocation3,
		storageLocation4, storageLocation5)

	storageLocationsFiltered := service.GetCheapestStorageLocationsForQuality(storageLocations, minQuality)

	qualitySum := 0
	for _, storageLocation := range storageLocationsFiltered {
		qualitySum += storageLocation.Quality
	}

	if len(storageLocationsFiltered) != 3 || qualitySum < minQuality {
		panic("TestGetCheapestStorageLocationsForQuality failed")
	}
}

func TestGetCheapestStorageLocationsForQuality2(t *testing.T) {
	minQuality := 1000
	storageLocation1 := models.StorageLocation{Quality: 34, Price: 99}
	storageLocation2 := models.StorageLocation{Quality: 22, Price: 67}
	storageLocation3 := models.StorageLocation{Quality: 10, Price: 40}
	storageLocation4 := models.StorageLocation{Quality: 27, Price: 56}
	storageLocation5 := models.StorageLocation{Quality: 68, Price: 200}

	storageLocations := make([]models.StorageLocation, 0)
	storageLocations = append(storageLocations, storageLocation1, storageLocation2, storageLocation3,
		storageLocation4, storageLocation5)

	storageLocationsFiltered := service.GetCheapestStorageLocationsForQuality(storageLocations, minQuality)

	if len(storageLocationsFiltered) != 0 {
		panic("TestGetCheapestStorageLocationsForQuality2 failed")
	}
}

func TestGetStorageLocationsToCopyTo(t *testing.T) {

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "sftp", Id: "sftp"}, {Type: "Amazon S3", Id: "Amazon S3"}, {Type: "Switch S3", Id: "Switch S3"}, {Type: "local", Id: "local"}}}
	currentStorageLocations := []*pb.StorageLocation{{Type: "sftp", Id: "sftp"}}

	storageLocationsToCopyIn := service.GetStorageLocationsToCopyTo(storageLocationsNeeded, currentStorageLocations)
	if len(storageLocationsToCopyIn) != 3 {
		panic("TestGetStorageLocationsToCopyTo has failed")
	}
	for _, storageLocation := range storageLocationsToCopyIn {
		if storageLocation.Type == "sftp" {
			panic("TestGetStorageLocationsToCopyTo has failed")
		}
	}
}

func TestGetStorageLocationsToCopyTo2(t *testing.T) {

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "sftp", Id: "sftp"}, {Type: "Amazon S3", Id: "Amazon S3"}, {Type: "Switch S3", Id: "Switch S3"}, {Type: "local", Id: "local"}}}
	currentStorageLocations := []*pb.StorageLocation{{Type: "sftp", Id: "sftp"}, {Type: "Amazon S3", Id: "Amazon S3"}}

	storageLocationsToCopyIn := service.GetStorageLocationsToCopyTo(storageLocationsNeeded, currentStorageLocations)
	if len(storageLocationsToCopyIn) != 2 {
		panic("TestGetStorageLocationsToCopyTo2 has failed")
	}
	for _, storageLocation := range storageLocationsToCopyIn {
		if storageLocation.Type == "sftp" || storageLocation.Type == "Amazon S3" {
			panic("TestGetStorageLocationsToCopyTo2 has failed")
		}
	}
}

func TestGetStorageLocationsToCopyTo3(t *testing.T) {

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "sftp", Id: "sftp"}}}
	currentStorageLocations := []*pb.StorageLocation{{Type: "Amazon S3", Id: "Amazon S3"}, {Type: "Switch S3", Id: "Switch S3"}, {Type: "local", Id: "local"}}

	storageLocationsToCopyIn := service.GetStorageLocationsToCopyTo(storageLocationsNeeded, currentStorageLocations)
	if len(storageLocationsToCopyIn) != 1 {
		panic("TestGetStorageLocationsToCopyTo3 has failed")
	}
	for _, storageLocation := range storageLocationsToCopyIn {
		if storageLocation.Type != "sftp" {
			panic("TestGetStorageLocationsToCopyTo3 has failed")
		}
	}
}

func TestGetStorageLocationsToDeleteFrom(t *testing.T) {

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "sftp", Id: "sftp"}, {Type: "Amazon S3", Id: "Amazon S3"}, {Type: "Switch S3", Id: "Switch S3"}, {Type: "local", Id: "local"}}}
	currentStorageLocations := map[*pb.ObjectInstance]*pb.StorageLocation{{Id: "1"}: {Type: "sftp", Id: "sftp"}}

	storageLocationsToDeleteFrom := service.GetStorageLocationsToDeleteFrom(storageLocationsNeeded, currentStorageLocations)
	if len(storageLocationsToDeleteFrom) != 0 {
		panic("TestGetStorageLocationsToDeleteFrom has failed")
	}
}

func TestGetStorageLocationsToDeleteFrom2(t *testing.T) {

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "Switch S3", Id: "Switch S3"}, {Type: "local", Id: "local"}}}
	currentStorageLocations := map[*pb.ObjectInstance]*pb.StorageLocation{{Id: "1"}: {Type: "sftp", Id: "sftp"}, {Id: "2"}: {Type: "Amazon S3", Id: "Amazon S3"}}

	storageLocationsToDeleteFrom := service.GetStorageLocationsToDeleteFrom(storageLocationsNeeded, currentStorageLocations)
	if len(storageLocationsToDeleteFrom) != 2 {
		panic("TestGetStorageLocationsToDeleteFrom2 has failed")
	}
	for _, storageLocation := range storageLocationsToDeleteFrom {
		if storageLocation.Type == "Switch S3" || storageLocation.Type == "local" {
			panic("TestGetStorageLocationsToDeleteFrom2 has failed")
		}
	}
}

func TestGetStorageLocationsToDeleteFrom3(t *testing.T) {

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "sftp", Id: "sftp"}}}
	currentStorageLocations := map[*pb.ObjectInstance]*pb.StorageLocation{{Id: "1"}: {Type: "Amazon S3", Id: "Amazon S3"}, {Id: "2"}: {Type: "Switch S3", Id: "Switch S3"}, {Id: "3"}: {Type: "local", Id: "local"}}

	storageLocationsToDeleteFrom := service.GetStorageLocationsToDeleteFrom(storageLocationsNeeded, currentStorageLocations)
	if len(storageLocationsToDeleteFrom) != 3 {
		panic("TestGetStorageLocationsToDeleteFrom3 has failed")
	}
	for oi, storageLocation := range storageLocationsToDeleteFrom {
		if storageLocation.Type == "sftp" {
			panic("TestGetStorageLocationsToDeleteFrom3 has failed")
		}
		switch {
		case oi.Id == "1":
			if storageLocation.Id != "Amazon S3" {
				panic("TestGetStorageLocationsToDeleteFrom3 has failed")
			}
		case oi.Id == "2":
			if storageLocation.Id != "Switch S3" {
				panic("TestGetStorageLocationsToDeleteFrom3 has failed")
			}
		}
	}
}
