package datamodels

type EdgeCamera struct {
	Name       string
	Resolution string
}
type EdgeStats struct {
	Camera           []EdgeCamera
	FolderUpdateTime string
}
type Stats struct {
	Stats []map[string]EdgeStats
}
