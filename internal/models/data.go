package models

type DataMetadata struct {
	UserID   string `json:"user_id"`
	DataSize int    `json:"data_size"`
	Preview  string `json:"preview"`
}
