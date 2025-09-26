package service

import (
	"math"
	"slices"

	"github.com/ocfl-archive/dlza-manager/dlzamanagerproto"
)

type minCostWithIndices struct {
	MinCost int
	Indices []int
}

func GetCheapestStorageLocationsForQuality(storageLocations *dlzamanagerproto.StorageLocations, minQuality int) []*dlzamanagerproto.StorageLocation {

	storageLocationsGroups := dlzamanagerproto.StorageLocations{}

	for _, storageLocation := range storageLocations.StorageLocations {
		var storageLocationWithSameGroup *dlzamanagerproto.StorageLocation
		for _, storageLocationU := range storageLocations.StorageLocations {
			if storageLocation.Group == storageLocationU.Group && storageLocationWithSameGroup == nil {
				storageLocationWithSameGroup = storageLocationU
			} else if storageLocation.Alias == storageLocationU.Alias && storageLocationWithSameGroup != nil {
				storageLocationsGroups.StorageLocations = append(storageLocationsGroups.StorageLocations, storageLocationWithSameGroup)
				break
			}
			storageLocationsGroups.StorageLocations = append(storageLocationsGroups.StorageLocations, storageLocationWithSameGroup)
		}
	}

	maxQuality := 0

	for _, storageLocation := range storageLocationsGroups.StorageLocations {
		maxQuality += int(storageLocation.Quality)
	}

	minCosts := make([]minCostWithIndices, maxQuality+1)

	for i := 1; i <= maxQuality; i++ {
		minCosts[i].MinCost = math.MaxInt64
		if minCosts[i].Indices == nil {
			minCosts[i].Indices = make([]int, 0)
		}
	}

	for i := 0; i < len(storageLocationsGroups.StorageLocations); i++ {
		for j := maxQuality; j >= int(storageLocationsGroups.StorageLocations[i].Quality); j-- {

			if minCosts[j-int(storageLocationsGroups.StorageLocations[i].Quality)].MinCost != math.MaxInt {

				appendBool := minCosts[j].MinCost > minCosts[j-int(storageLocationsGroups.StorageLocations[i].Quality)].MinCost+int(storageLocationsGroups.StorageLocations[i].Price)
				minCosts[j].MinCost = int(math.Min(float64(minCosts[j].MinCost), float64(minCosts[j-int(storageLocationsGroups.StorageLocations[i].Quality)].MinCost+int(storageLocationsGroups.StorageLocations[i].Price))))

				if appendBool {
					minCosts[j].Indices = make([]int, len(minCosts[j-int(storageLocationsGroups.StorageLocations[i].Quality)].Indices))
					copy(minCosts[j].Indices, minCosts[j-int(storageLocationsGroups.StorageLocations[i].Quality)].Indices)
					minCosts[j].Indices = append(minCosts[j].Indices, i)
				}
			}
		}
	}
	answer := math.MaxInt64
	indices := make([]int, 0)
	for i := minQuality; i <= maxQuality; i++ {
		if minCosts[i].MinCost != math.MaxInt64 {
			if answer > minCosts[i].MinCost {
				answer = minCosts[i].MinCost
				indices = minCosts[i].Indices
			}
		}
	}
	return getStorageLocationsByIndices(storageLocations, indices)
}

func GetStorageLocationsToCopyTo(relevantStorageLocations *dlzamanagerproto.StorageLocations, storageLocationsInUse []*dlzamanagerproto.StorageLocation) []*dlzamanagerproto.StorageLocation {
	storageLocationsToCopyTo := make([]*dlzamanagerproto.StorageLocation, 0)

	for _, relevantStorageLocation := range relevantStorageLocations.StorageLocations {
		for index, currentStorageLocation := range storageLocationsInUse {
			if relevantStorageLocation.Id == currentStorageLocation.Id {
				break
			}
			if index == len(storageLocationsInUse)-1 {
				storageLocationsToCopyTo = append(storageLocationsToCopyTo, relevantStorageLocation)
			}
		}
	}
	return storageLocationsToCopyTo
}

func GetStorageLocationsToDeleteFrom(relevantStorageLocations *dlzamanagerproto.StorageLocations, storageLocationsInUse map[*dlzamanagerproto.ObjectInstance]*dlzamanagerproto.StorageLocation) map[*dlzamanagerproto.ObjectInstance]*dlzamanagerproto.StorageLocation {
	storageLocationsToDeleteFrom := make(map[*dlzamanagerproto.ObjectInstance]*dlzamanagerproto.StorageLocation)

	for oiInUse, storageLocationInUse := range storageLocationsInUse {
		for index, relevantStorageLocation := range relevantStorageLocations.StorageLocations {
			if storageLocationInUse.Id == relevantStorageLocation.Id {
				break
			}
			if index == len(relevantStorageLocations.StorageLocations)-1 {
				storageLocationsToDeleteFrom[oiInUse] = storageLocationInUse
			}
		}
	}
	return storageLocationsToDeleteFrom
}

func getStorageLocationsByIndices(storageLocations *dlzamanagerproto.StorageLocations, indices []int) []*dlzamanagerproto.StorageLocation {
	filteredStorageLocations := make([]*dlzamanagerproto.StorageLocation, 0)
	for index, storageLocation := range storageLocations.StorageLocations {
		if slices.Contains(indices, index) {
			filteredStorageLocations = append(filteredStorageLocations, storageLocation)
		}
	}
	return filteredStorageLocations
}
