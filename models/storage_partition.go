package models

type StoragePartition struct {
	Alias             string
	Name              string
	MaxSize           int64
	MaxObjects        int
	CurrentSize       int64
	CurrentObjects    int
	Id                string
	StorageLocationId string
}
