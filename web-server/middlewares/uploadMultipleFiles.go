package middlewares

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadMultipleFilesMiddleware handles the uploading of multiple media files
func UploadMultipleFilesMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(141414)

		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve multipart form: " + err.Error()})
			return
		}

		files := form.File["media_files"] // Assuming "media_files" is the name of the form-data field

		for _, file := range files {
			// Check file type by extension
			ext := filepath.Ext(file.Filename)
			switch ext {
			case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", // Images
				".mp4", ".avi", ".mov", ".wmv", // Videos
				".mp3", ".wav", ".ogg", ".m4a", // Audio
				".pdf", ".doc", ".docx", ".ppt", ".pptx", ".xls", ".xlsx": // Documents
				// File type is supported, continue processing
			default:
				ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("File type '%s' not supported", ext)})
				return
			}

			// Save file to a specified directory
			dst := filepath.Join("uploads", filepath.Base(file.Filename))
			if err := ctx.SaveUploadedFile(file, dst); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save file: " + err.Error()})
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully!"})
	}
}
