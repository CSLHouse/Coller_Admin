package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/example"

type ExaFilesRequest struct {
	Files []example.ExaFileUploadAndDownload `json:"files"`
}
