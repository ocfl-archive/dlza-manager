package tests

import (
	"testing"

	pb "github.com/ocfl-archive/dlza-manager/dlzamanagerproto"

	"github.com/ocfl-archive/dlza-manager/service"
)

func TestGetCheapestStorageLocationsForQuality(t *testing.T) {
	minQuality := 66
	storageLocation1 := &pb.StorageLocation{Quality: 34, Price: 99}
	storageLocation2 := &pb.StorageLocation{Quality: 22, Price: 67}
	storageLocation3 := &pb.StorageLocation{Quality: 10, Price: 40}
	storageLocation4 := &pb.StorageLocation{Quality: 27, Price: 56}
	storageLocation5 := &pb.StorageLocation{Quality: 68, Price: 200}

	storageLocations := make([]*pb.StorageLocation, 0)
	storageLocations = append(storageLocations, storageLocation1, storageLocation2, storageLocation3,
		storageLocation4, storageLocation5)

	storageLocationsFiltered := service.GetCheapestStorageLocationsForQuality(&pb.StorageLocations{StorageLocations: storageLocations}, minQuality)

	qualitySum := 0
	for _, storageLocation := range storageLocationsFiltered {
		qualitySum += int(storageLocation.Quality)
	}

	if len(storageLocationsFiltered) != 3 || qualitySum < minQuality {
		panic("TestGetCheapestStorageLocationsForQuality failed")
	}
}

func TestGetCheapestStorageLocationsForQuality2(t *testing.T) {
	minQuality := 1000
	storageLocation1 := &pb.StorageLocation{Quality: 34, Price: 99}
	storageLocation2 := &pb.StorageLocation{Quality: 22, Price: 67}
	storageLocation3 := &pb.StorageLocation{Quality: 10, Price: 40}
	storageLocation4 := &pb.StorageLocation{Quality: 27, Price: 56}
	storageLocation5 := &pb.StorageLocation{Quality: 68, Price: 200}
	storageLocations := make([]*pb.StorageLocation, 0)
	storageLocations = append(storageLocations, storageLocation1, storageLocation2, storageLocation3,
		storageLocation4, storageLocation5)

	storageLocationsFiltered := service.GetCheapestStorageLocationsForQuality(&pb.StorageLocations{StorageLocations: storageLocations}, minQuality)

	if len(storageLocationsFiltered) != 0 {
		panic("TestGetCheapestStorageLocationsForQuality2 failed")
	}
}

func TestGetCheapestStorageLocationsForQuality3(t *testing.T) {
	minQuality := 60
	storageLocation1 := &pb.StorageLocation{Quality: 30, Price: 99}
	storageLocation2 := &pb.StorageLocation{Quality: 30, Price: 67}
	storageLocation3 := &pb.StorageLocation{Quality: 30, Price: 40}
	storageLocation4 := &pb.StorageLocation{Quality: 30, Price: 56}
	storageLocation5 := &pb.StorageLocation{Quality: 30, Price: 200}

	storageLocations := make([]*pb.StorageLocation, 0)
	storageLocations = append(storageLocations, storageLocation1, storageLocation2, storageLocation3,
		storageLocation4, storageLocation5)

	storageLocationsFiltered := service.GetCheapestStorageLocationsForQuality(&pb.StorageLocations{StorageLocations: storageLocations}, minQuality)

	qualitySum := 0
	for _, storageLocation := range storageLocationsFiltered {
		qualitySum += int(storageLocation.Quality)
	}

	if len(storageLocationsFiltered) != 2 || qualitySum < minQuality {
		panic("TestGetCheapestStorageLocationsForQualit3 failed")
	}
	for _, storageLocationsFilteredItem := range storageLocationsFiltered {
		if storageLocationsFilteredItem.Price != 40 && storageLocationsFilteredItem.Quality != 56 {
			panic("TestGetCheapestStorageLocationsForQuality3 failed")
		}
	}
}

func TestGetCheapestStorageLocationsForQuality4(t *testing.T) {
	minQuality := 70
	storageLocation1 := &pb.StorageLocation{Quality: 10, Price: 999}
	storageLocation2 := &pb.StorageLocation{Quality: 20, Price: 6079}
	storageLocation3 := &pb.StorageLocation{Quality: 30, Price: 900}
	storageLocation4 := &pb.StorageLocation{Quality: 35, Price: 1000}
	storageLocation5 := &pb.StorageLocation{Quality: 40, Price: 1000}

	storageLocations := make([]*pb.StorageLocation, 0)
	storageLocations = append(storageLocations, storageLocation1, storageLocation2, storageLocation3,
		storageLocation4, storageLocation5)

	storageLocationsFiltered := service.GetCheapestStorageLocationsForQuality(&pb.StorageLocations{StorageLocations: storageLocations}, minQuality)

	qualitySum := 0
	for _, storageLocation := range storageLocationsFiltered {
		qualitySum += int(storageLocation.Quality)
	}

	if len(storageLocationsFiltered) != 2 || qualitySum < minQuality {
		panic("TestGetCheapestStorageLocationsForQualit4 failed")
	}
	for _, storageLocationsFilteredItem := range storageLocationsFiltered {
		if storageLocationsFilteredItem.Price != 900 && storageLocationsFilteredItem.Quality != 1000 {
			panic("TestGetCheapestStorageLocationsForQuality4 failed")
		}
	}
}

func TestGetCheapestStorageLocationsForQuality5(t *testing.T) {
	minQuality := 70
	storageLocation1 := &pb.StorageLocation{Quality: 10, Price: 999}
	storageLocation2 := &pb.StorageLocation{Quality: 20, Price: 6079}
	storageLocation3 := &pb.StorageLocation{Quality: 30, Price: 1100}
	storageLocation4 := &pb.StorageLocation{Quality: 35, Price: 1000}
	storageLocation5 := &pb.StorageLocation{Quality: 40, Price: 1000}

	storageLocations := make([]*pb.StorageLocation, 0)
	storageLocations = append(storageLocations, storageLocation1, storageLocation2, storageLocation3,
		storageLocation4, storageLocation5)

	storageLocationsFiltered := service.GetCheapestStorageLocationsForQuality(&pb.StorageLocations{StorageLocations: storageLocations}, minQuality)

	qualitySum := 0
	for _, storageLocation := range storageLocationsFiltered {
		qualitySum += int(storageLocation.Quality)
	}

	if len(storageLocationsFiltered) != 2 || qualitySum < minQuality {
		panic("TestGetCheapestStorageLocationsForQualit5 failed")
	}
	for _, storageLocationsFilteredItem := range storageLocationsFiltered {
		if storageLocationsFilteredItem.Price != 1000 && storageLocationsFilteredItem.Quality != 1000 {
			panic("TestGetCheapestStorageLocationsForQuality5 failed")
		}
	}
}

func TestGetCheapestStorageLocationsForQuality6(t *testing.T) {
	minQuality := 70
	storageLocation5 := &pb.StorageLocation{Quality: 70, Price: 1000}

	storageLocations := make([]*pb.StorageLocation, 0)
	storageLocations = append(storageLocations, storageLocation5)

	storageLocationsFiltered := service.GetCheapestStorageLocationsForQuality(&pb.StorageLocations{StorageLocations: storageLocations}, minQuality)

	qualitySum := 0
	for _, storageLocation := range storageLocationsFiltered {
		qualitySum += int(storageLocation.Quality)
	}

	if len(storageLocationsFiltered) == 0 || qualitySum < minQuality {
		panic("TestGetCheapestStorageLocationsForQualit6 failed")
	}
	for _, storageLocationsFilteredItem := range storageLocationsFiltered {
		if storageLocationsFilteredItem.Quality != 70 {
			panic("TestGetCheapestStorageLocationsForQuality6 failed")
		}
	}
}

func TestGetCheapestStorageLocationsForQuality7(t *testing.T) {
	minQuality := 70
	storageLocation1 := &pb.StorageLocation{Quality: 10, Price: 999, Group: "group1"}
	storageLocation2 := &pb.StorageLocation{Quality: 30, Price: 6079, Group: "group2"}
	storageLocation3 := &pb.StorageLocation{Quality: 20, Price: 1100, Group: "group3"}
	storageLocation5 := &pb.StorageLocation{Quality: 40, Price: 1000, Group: "group5"}
	storageLocation6 := &pb.StorageLocation{Quality: 40, Price: 1000, Group: "group5"}

	storageLocations := make([]*pb.StorageLocation, 0)
	storageLocations = append(storageLocations, storageLocation1, storageLocation2, storageLocation3, storageLocation5, storageLocation6)

	storageLocationsFiltered := service.GetCheapestStorageLocationsForQuality(&pb.StorageLocations{StorageLocations: storageLocations}, minQuality)

	qualitySum := 0
	for _, storageLocation := range storageLocationsFiltered {
		qualitySum += int(storageLocation.Quality)
	}

	if len(storageLocationsFiltered) != 3 || qualitySum != 70 {
		panic("TestGetCheapestStorageLocationsForQualit7 failed")
	}
}

func TestGetStorageLocationsToCopyTo(t *testing.T) {

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "sftp", Group: "sftp"}, {Type: "Amazon S3", Group: "Amazon S3"}, {Type: "Switch S3", Group: "Switch S3"}, {Type: "local", Group: "local"}}}
	currentStorageLocations := []*pb.StorageLocation{{Type: "sftp", Group: "sftp"}}

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

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "sftp", Group: "sftp"}, {Type: "Amazon S3", Group: "Amazon S3"}, {Type: "Switch S3", Group: "Switch S3"}, {Type: "local", Group: "local"}}}
	currentStorageLocations := []*pb.StorageLocation{{Type: "sftp", Group: "sftp"}, {Type: "Amazon S3", Group: "Amazon S3"}}

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

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "sftp", Group: "sftp"}}}
	currentStorageLocations := []*pb.StorageLocation{{Type: "Amazon S3", Group: "Amazon S3"}, {Type: "Switch S3", Group: "Switch S3"}, {Type: "local", Group: "local"}}

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

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "sftp", Group: "sftp"}, {Type: "Amazon S3", Group: "Amazon S3"}, {Type: "Switch S3", Group: "Switch S3"}, {Type: "local", Group: "local"}}}
	currentStorageLocations := map[*pb.ObjectInstance]*pb.StorageLocation{{Id: "1"}: {Type: "sftp", Group: "sftp"}}

	storageLocationsToDeleteFrom := service.GetStorageLocationsToDeleteFrom(storageLocationsNeeded, currentStorageLocations)
	if len(storageLocationsToDeleteFrom) != 0 {
		panic("TestGetStorageLocationsToDeleteFrom has failed")
	}
}

func TestGetStorageLocationsToDeleteFrom2(t *testing.T) {

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "Switch S3", Group: "Switch S3"}, {Type: "local", Group: "local"}}}
	currentStorageLocations := map[*pb.ObjectInstance]*pb.StorageLocation{{Id: "1"}: {Type: "sftp", Group: "sftp"}, {Id: "2"}: {Type: "Amazon S3", Group: "Amazon S3"}}

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

	storageLocationsNeeded := &pb.StorageLocations{StorageLocations: []*pb.StorageLocation{{Type: "sftp", Group: "sftp"}}}
	currentStorageLocations := map[*pb.ObjectInstance]*pb.StorageLocation{{Id: "1"}: {Type: "Amazon S3", Group: "Amazon S3"}, {Id: "2"}: {Type: "Switch S3", Group: "Switch S3"}, {Id: "3"}: {Type: "local", Group: "local"}}

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
			if storageLocation.Group != "Amazon S3" {
				panic("TestGetStorageLocationsToDeleteFrom3 has failed")
			}
		case oi.Id == "2":
			if storageLocation.Group != "Switch S3" {
				panic("TestGetStorageLocationsToDeleteFrom3 has failed")
			}
		}
	}
}
