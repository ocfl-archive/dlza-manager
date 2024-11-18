package models

type File struct {
	Checksum string   `json:"checksum"`
	Name     []string `json:"name"`
	Size     int      `json:"size"`
	MimeType string   `json:"mime_type"`
	Pronom   string   `json:"pronom"`
	Width    int64    `json:"width"`
	Height   int64    `json:"height"`
	Duration int64    `json:"duration"`
	Id       string   `json:"id"`
	ObjectId string   `json:"object_id"`
}
