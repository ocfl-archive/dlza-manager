package mapper

import (
	pb "github.com/ocfl-archive/dlza-manager/dlzamanagerproto"
	"github.com/ocfl-archive/dlza-manager/models"
)

func ConvertToTenant(tenantPb *pb.Tenant) models.Tenant {
	var tenant models.Tenant
	tenant.Id = tenantPb.Id
	tenant.Alias = tenantPb.Alias
	tenant.Person = tenantPb.Person
	tenant.Email = tenantPb.Email
	tenant.Name = tenantPb.Name
	tenant.ApiKeyId = tenantPb.ApiKeyId
	return tenant
}

func ConvertToTenantPb(tenant models.Tenant) *pb.Tenant {
	tenantPb := &pb.Tenant{}
	tenantPb.Name = tenant.Name
	tenantPb.Id = tenant.Id
	tenantPb.Person = tenant.Person
	tenantPb.Alias = tenant.Alias
	tenantPb.Email = tenant.Email
	tenantPb.ApiKeyId = tenant.ApiKeyId
	return tenantPb
}

func ConvertToCollection(collectionPb *pb.Collection) models.Collection {
	var collection models.Collection
	collection.Id = collectionPb.Id
	collection.Alias = collectionPb.Alias
	collection.Description = collectionPb.Description
	collection.Owner = collectionPb.Owner
	collection.Name = collectionPb.Name
	collection.OwnerMail = collectionPb.OwnerMail
	collection.Quality = int(collectionPb.Quality)
	collection.TenantId = collectionPb.TenantId
	collection.TotalObjectCount = collectionPb.TotalObjectCount
	collection.TotalFileCount = collectionPb.TotalFileCount
	collection.TotalFileSize = collectionPb.TotalFileSize
	return collection
}

func ConvertToCollectionPb(collection models.Collection) *pb.Collection {
	collectionPb := &pb.Collection{}
	collectionPb.Alias = collection.Alias
	collectionPb.Description = collection.Description
	collectionPb.Owner = collection.Owner
	collectionPb.Name = collection.Name
	collectionPb.OwnerMail = collection.OwnerMail
	collectionPb.Quality = int32(collection.Quality)
	collectionPb.TenantId = collection.TenantId
	collectionPb.Id = collection.Id
	collectionPb.TotalObjectCount = collection.TotalObjectCount
	collectionPb.TotalFileCount = collection.TotalFileCount
	collectionPb.TotalFileSize = collection.TotalFileSize
	return collectionPb
}

func ConvertToStorageLocation(storageLocationPb *pb.StorageLocation) models.StorageLocation {
	var storageLocation models.StorageLocation
	storageLocation.Id = storageLocationPb.Id
	storageLocation.Alias = storageLocationPb.Alias
	storageLocation.Type = storageLocationPb.Type
	storageLocation.Vault = storageLocationPb.Vault
	storageLocation.Connection = storageLocationPb.Connection
	storageLocation.Quality = int(storageLocationPb.Quality)
	storageLocation.Price = int(storageLocationPb.Price)
	storageLocation.SecurityCompliency = storageLocationPb.SecurityCompliency
	storageLocation.FillFirst = storageLocationPb.FillFirst
	storageLocation.OcflType = storageLocationPb.OcflType
	storageLocation.NumberOfThreads = int(storageLocationPb.NumberOfThreads)
	storageLocation.Group = storageLocationPb.Group
	storageLocation.TenantId = storageLocationPb.TenantId
	storageLocation.TotalFilesSize = storageLocationPb.TotalFilesSize
	storageLocation.TotalExistingVolume = storageLocationPb.TotalExistingVolume
	return storageLocation
}

func ConvertToStorageLocationPb(storageLocation models.StorageLocation) *pb.StorageLocation {
	storageLocationPb := &pb.StorageLocation{}
	storageLocationPb.Id = storageLocation.Id
	storageLocationPb.Alias = storageLocation.Alias
	storageLocationPb.Type = storageLocation.Type
	storageLocationPb.Vault = storageLocation.Vault
	storageLocationPb.Connection = storageLocation.Connection
	storageLocationPb.Quality = int32(storageLocation.Quality)
	storageLocationPb.Price = int32(storageLocation.Price)
	storageLocationPb.SecurityCompliency = storageLocation.SecurityCompliency
	storageLocationPb.FillFirst = storageLocation.FillFirst
	storageLocationPb.OcflType = storageLocation.OcflType
	storageLocationPb.NumberOfThreads = int32(storageLocation.NumberOfThreads)
	storageLocationPb.Group = storageLocation.Group
	storageLocationPb.TenantId = storageLocation.TenantId
	storageLocationPb.TotalFilesSize = storageLocation.TotalFilesSize
	storageLocationPb.TotalExistingVolume = storageLocation.TotalExistingVolume
	return storageLocationPb
}

func ConvertToObject(objectPb *pb.Object) models.Object {
	var object models.Object
	object.Signature = objectPb.Signature
	object.Sets = objectPb.Sets
	object.Identifiers = objectPb.Identifiers
	object.Title = objectPb.Title
	object.AlternativeTitles = objectPb.AlternativeTitles
	object.Description = objectPb.Description
	object.Keywords = objectPb.Keywords
	object.References = objectPb.References
	object.IngestWorkflow = objectPb.IngestWorkflow
	object.User = objectPb.User
	object.Address = objectPb.Address
	object.Created = objectPb.Created
	object.LastChanged = objectPb.LastChanged
	object.Size = objectPb.Size
	object.Id = objectPb.Id
	object.CollectionId = objectPb.CollectionId
	object.Checksum = objectPb.Checksum
	object.Expiration = objectPb.Expiration
	object.Authors = objectPb.Authors
	object.Holding = objectPb.Holding
	object.Versions = objectPb.Versions
	object.Head = objectPb.Head
	object.TotalFileCount = objectPb.TotalFileCount
	object.TotalFileSize = objectPb.TotalFileSize
	object.Binary = objectPb.Binary
	return object
}

func ConvertToObjectPb(object models.Object) *pb.Object {
	objectPb := &pb.Object{}
	objectPb.Signature = object.Signature
	objectPb.Sets = object.Sets
	objectPb.Identifiers = object.Identifiers
	objectPb.Title = object.Title
	objectPb.AlternativeTitles = object.AlternativeTitles
	objectPb.Description = object.Description
	objectPb.Keywords = object.Keywords
	objectPb.References = object.References
	objectPb.IngestWorkflow = object.IngestWorkflow
	objectPb.User = object.User
	objectPb.Address = object.Address
	objectPb.Created = object.Created
	objectPb.LastChanged = object.LastChanged
	objectPb.Size = object.Size
	objectPb.Id = object.Id
	objectPb.Expiration = object.Expiration
	objectPb.Authors = object.Authors
	objectPb.Holding = object.Holding
	objectPb.CollectionId = object.CollectionId
	objectPb.Checksum = object.Checksum
	objectPb.Versions = object.Versions
	objectPb.Head = object.Head
	objectPb.TotalFileCount = object.TotalFileCount
	objectPb.TotalFileSize = object.TotalFileSize
	objectPb.Binary = object.Binary
	return objectPb
}

func ConvertToObjectInstance(objectInstancePb *pb.ObjectInstance) models.ObjectInstance {
	var objectInstance models.ObjectInstance
	objectInstance.Path = objectInstancePb.Path
	objectInstance.Size = int(objectInstancePb.Size)
	objectInstance.Created = objectInstancePb.Created
	objectInstance.Status = objectInstancePb.Status
	objectInstance.Id = objectInstancePb.Id
	objectInstance.StoragePartitionId = objectInstancePb.StoragePartitionId
	objectInstance.ObjectId = objectInstancePb.ObjectId

	return objectInstance
}

func ConvertToObjectInstancePb(objectInstance models.ObjectInstance) *pb.ObjectInstance {
	objectInstancePb := &pb.ObjectInstance{}
	objectInstancePb.Path = objectInstance.Path
	objectInstancePb.Size = int64(objectInstance.Size)
	objectInstancePb.Created = objectInstance.Created
	objectInstancePb.Status = objectInstance.Status
	objectInstancePb.Id = objectInstance.Id
	objectInstancePb.StoragePartitionId = objectInstance.StoragePartitionId
	objectInstancePb.ObjectId = objectInstance.ObjectId
	return objectInstancePb
}

func ConvertToStoragePartition(storagePartitionPb *pb.StoragePartition) models.StoragePartition {
	var storagePartition models.StoragePartition
	storagePartition.Alias = storagePartitionPb.Alias
	storagePartition.Name = storagePartitionPb.Name
	storagePartition.MaxSize = storagePartitionPb.MaxSize
	storagePartition.MaxObjects = int(storagePartitionPb.MaxObjects)
	storagePartition.CurrentSize = storagePartitionPb.CurrentSize
	storagePartition.CurrentObjects = int(storagePartitionPb.CurrentObjects)
	storagePartition.Id = storagePartitionPb.Id
	storagePartition.StorageLocationId = storagePartitionPb.StorageLocationId

	return storagePartition
}

func ConvertToStoragePartitionPb(storagePartition models.StoragePartition) *pb.StoragePartition {
	storagePartitionPb := &pb.StoragePartition{}
	storagePartitionPb.Alias = storagePartition.Alias
	storagePartitionPb.Name = storagePartition.Name
	storagePartitionPb.MaxSize = int64(storagePartition.MaxSize)
	storagePartitionPb.MaxObjects = int64(storagePartition.MaxObjects)
	storagePartitionPb.CurrentSize = int64(storagePartition.CurrentSize)
	storagePartitionPb.CurrentObjects = int64(storagePartition.CurrentObjects)
	storagePartitionPb.Id = storagePartition.Id
	storagePartitionPb.StorageLocationId = storagePartition.StorageLocationId
	return storagePartitionPb
}

func ConvertToFile(filePb *pb.File) models.File {
	var file models.File
	file.Checksum = filePb.Checksum
	file.Name = filePb.Name
	file.Size = int(filePb.Size)
	file.MimeType = filePb.MimeType
	file.Pronom = filePb.Pronom
	file.Width = filePb.Width
	file.Height = filePb.Height
	file.Duration = filePb.Duration
	file.Id = filePb.Id
	file.ObjectId = filePb.ObjectId
	return file
}

func ConvertToFilePb(file models.File) *pb.File {
	filePb := &pb.File{}
	filePb.Checksum = file.Checksum
	filePb.Name = file.Name
	filePb.Size = int64(file.Size)
	filePb.MimeType = file.MimeType
	filePb.Pronom = file.Pronom
	filePb.Width = file.Width
	filePb.Height = file.Height
	filePb.Duration = file.Duration
	filePb.Id = file.Id
	filePb.ObjectId = file.ObjectId
	return filePb
}

func ConvertToObjectInstanceCheckPb(objectInstanceCheck models.ObjectInstanceCheck) *pb.ObjectInstanceCheck {
	var objectInstanceCheckPb pb.ObjectInstanceCheck
	objectInstanceCheckPb.CheckTime = objectInstanceCheck.CheckTime
	objectInstanceCheckPb.Error = objectInstanceCheck.Error
	objectInstanceCheckPb.Message = objectInstanceCheck.Message
	objectInstanceCheckPb.Id = objectInstanceCheck.Id
	objectInstanceCheckPb.ObjectInstanceId = objectInstanceCheck.ObjectInstanceId
	objectInstanceCheckPb.CheckType = objectInstanceCheck.CheckType
	return &objectInstanceCheckPb
}

func ConvertToObjectInstanceCheck(objectInstanceCheckPb *pb.ObjectInstanceCheck) models.ObjectInstanceCheck {
	var objectInstanceCheck models.ObjectInstanceCheck
	objectInstanceCheck.CheckTime = objectInstanceCheckPb.CheckTime
	objectInstanceCheck.Error = objectInstanceCheckPb.Error
	objectInstanceCheck.Message = objectInstanceCheckPb.Message
	objectInstanceCheck.Id = objectInstanceCheckPb.Id
	objectInstanceCheck.ObjectInstanceId = objectInstanceCheckPb.ObjectInstanceId
	objectInstanceCheck.CheckType = objectInstanceCheckPb.CheckType
	return objectInstanceCheck
}

func ConvertToPagination(paginationPb *pb.Pagination) models.Pagination {
	var pagination models.Pagination
	pagination.Id = paginationPb.Id
	pagination.SecondId = paginationPb.SecondId
	pagination.SearchField = paginationPb.SearchField
	pagination.SortKey = paginationPb.SortKey
	pagination.SortDirection = paginationPb.SortDirection
	pagination.Skip = int(paginationPb.Skip)
	pagination.Take = int(paginationPb.Take)
	pagination.AllowedTenants = paginationPb.AllowedTenants
	return pagination
}

func ConvertStorageLocationsCombinationsForCollection(object models.CollectionWithExistingStorageLocationsCombinations) *pb.StorageLocationsCombinationsForCollection {
	var objectPb pb.StorageLocationsCombinationsForCollection
	objectPb.Id = object.Id
	objectPb.Alias = object.Alias
	objectPb.LocationsIds = object.LocationsIds
	objectPb.Quality = object.Quality
	objectPb.Price = object.Price
	objectPb.Size = object.Size
	return &objectPb
}
