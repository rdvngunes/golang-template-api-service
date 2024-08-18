package common

type UploadRequest struct {
	UnsignedPath string `json:"unsigned_path" binding:"required"`
}
