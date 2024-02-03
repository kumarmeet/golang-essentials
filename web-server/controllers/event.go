package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/learning-webserver/models"
	"github.com/xuri/excelize/v2"
)

func extractEmergencyContacts(emergencyContact []string, emergencyData *map[string][]map[string]string) {
	for i := 0; i < len(emergencyContact); i += 5 {
		if emergencyContact[i] == "Person Type" &&
			emergencyContact[i+1] == "Name" &&
			emergencyContact[i+2] == "Relationship" &&
			emergencyContact[i+3] == "Phone" &&
			emergencyContact[i+4] == "Home Contact" {
			continue // Skip the header row
		}

		details := map[string]string{
			"Person Type":  emergencyContact[i],
			"Name":         emergencyContact[i+1],
			"Relationship": emergencyContact[i+2],
			"Phone":        emergencyContact[i+3],
			"Home Contact": emergencyContact[i+4],
		}

		(*emergencyData)["Emergency Contact"] = append((*emergencyData)["Emergency Contact"], details)
	}
}

func extractProposers(proposers []string, proposerData *map[string][]map[string]string) {
	for i := 0; i < len(proposers); i += 5 {
		if proposers[i] == "Type" &&
			proposers[i+1] == "Name" &&
			proposers[i+2] == "Membership No" &&
			proposers[i+3] == "Phone" &&
			proposers[i+4] == "Home Contact" {
			continue // Skip the header row
		}

		details := map[string]string{
			"Type":          proposers[i],
			"Name":          proposers[i+1],
			"Membership No": proposers[i+2],
			"Phone":         proposers[i+3],
			"Home Contact":  proposers[i+4],
		}
		(*proposerData)["Proposers"] = append((*proposerData)["Proposers"], details)
	}
}

func mergeEmergencyAndProposers(excelData []map[string]interface{}, emergencyData, proposerData *map[string][]map[string]string) []map[string]interface{} {
	var k = 0

	// Creating a dynamic "Emergency Contact" array within excelData
	for i, data := range excelData {
		if i < len((*emergencyData)["Emergency Contact"]) {
			data["Emergency Contact"] = (*emergencyData)["Emergency Contact"][k : k+2] // Get 2 values from emergencyData
			data["Proposers"] = (*proposerData)["Proposers"][k : k+2]                  // Get 2 values from emergencyData
			k += 2
		} else {
			break // If there are no more emergency contacts, stop the loop
		}
		excelData[i] = data // Update the excelData with the new "Emergency Contact" data
	}

	return excelData
}

func structureSpreadsheetData(rows [][]string) []map[string]interface{} {

	var excelData = make([]map[string]interface{}, 0)
	var headers []string
	var emergencyData = make(map[string][]map[string]string)
	var proposerData = make(map[string][]map[string]string)

	if len(rows) > 0 {
		// Assume rows represents the data extracted from the Excel sheet

		// Iterate over rows to populate excelData
		for rowIndex, row := range rows {
			// If it's the first row, assume it contains headers
			if rowIndex == 0 {
				headers = row // Store the headers
				continue      // Skip to the next iteration to avoid adding headers to data
			}

			rowData := make(map[string]interface{})

			// Iterate over each cell in the row
			for colIndex, colCell := range row {
				// Ensure the column index is within the range of headers
				if colIndex < len(headers) {
					// Assign data to the respective header key
					if (colCell != "") &&
						(headers[colIndex] != "") &&
						(headers[colIndex] != "Emergency contact") &&
						(headers[colIndex] != "Proposers") {
						rowData[headers[colIndex]] = colCell
					} else if headers[colIndex] == "Emergency contact" {
						emergencyContact := row[27:32]
						extractEmergencyContacts(emergencyContact, &emergencyData)
					} else if headers[colIndex] == "Proposers" {
						proposers := row[32:37]
						extractProposers(proposers, &proposerData)
					}
				} else {
					// Handle cases where the number of headers doesn't match the data
					// You may want to log an error or handle it in another appropriate way
				}
			}

			// Append the row data to excelData
			excelData = append(excelData, rowData)
		}
	}

	for i := 0; i < len(excelData); {
		if len(excelData[i]) == 0 {
			excelData = append(excelData[:i], excelData[i+1:]...)
		} else {
			i++
		}
	}

	return mergeEmergencyAndProposers(excelData, &emergencyData, &proposerData)
}

func ImportCsvXlsx(ctx *gin.Context) {
	file, _ := ctx.Get("uploadedFiles")

	file_url, ok := file.([]string)

	if !ok || len(file_url) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid image data"})
		return
	}

	f, err := excelize.OpenFile(file_url[0])

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!", "err": err.Error()})
		return
	}

	// cols, err := f.GetCols("Sheet1")

	var excelData = structureSpreadsheetData(rows)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!", "err": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Sucessfully uploaded data!", "data": excelData})
}

func UploadFile(ctx *gin.Context) {
	event_id, err := strconv.ParseInt(ctx.Param("event_id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	image, _ := ctx.Get("uploadedFiles")

	image_url, ok := image.([]string)

	if !ok || len(image_url) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid image data"})
		return
	}

	var eventImage models.EventImage

	eventImageDetails, err := eventImage.Save(image_url[0], event_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not save image, try again later!", "err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "status": true, "data": eventImageDetails})
}

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch data, try again later"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": events})
}

func GetEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	var event models.Event

	events, err := event.GetEvent(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": events})
}

func DeleteEventById(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")

	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	var event models.Event

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parsed request data!"})
		return
	}

	events, err := event.DeleteEvent(eventId, userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": events})
}
func UpdateEventById(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event, try again later!"})
		return
	}

	var event models.Event

	err = ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parsed request data!"})
		return
	}

	events, err := event.UpdateEvent(eventId, userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": events})
}

func InsertEvent(ctx *gin.Context) {
	id := ctx.GetInt64("userId")

	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event.UserId = id

	eventId, err := event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	eventData, _ := event.GetEvent(eventId)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Data successfully saved.", "event": eventData})
}
