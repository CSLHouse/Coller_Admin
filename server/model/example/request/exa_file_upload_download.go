package response

import "cooller/server/model/example"

type ExaFilesRequest struct {
	Files []example.ExaFileUploadAndDownload `json:"files"`
}
