package services

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

func StructureSpreadsheetData(rows [][]string) []map[string]interface{} {

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
