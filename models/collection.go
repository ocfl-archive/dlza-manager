package models

type Collection struct {
	Id          string `json:"id"`
	Alias       string `json:"alias"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
	OwnerMail   string `json:"ownerMail"`
	Name        string `json:"name"`
	Quality     int    `json:"quality"`
	TenantId    string `json:"tenantId"`
	//virtual columns
	TotalFileSize    int64 `json:"totalFileSize"`
	TotalFileCount   int64 `json:"totalFileCount"`
	TotalObjectCount int64 `json:"totalObjectCount"`
}
