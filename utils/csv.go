package utils

import (
	"bytes"
	"capstone/model/payload"
	"encoding/csv"
	"fmt"
)

func ExportCSV(complaints []*payload.GetAllComplaintsResponse) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	csvWriter := csv.NewWriter(buffer)

	headers := []string{"No", "ID", "Name", "Type", "Category", "Description", "Status", "Is Public", "Created At"}
	if err := csvWriter.Write(headers); err != nil {
		return nil, err
	}

	no := 1
	for _, complaint := range complaints {
		record := []string{
			fmt.Sprintf("%d", no),
			fmt.Sprintf("%d", complaint.ID),
			complaint.Name,
			complaint.Type,
			complaint.Category,
			complaint.Description,
			complaint.Status,
			fmt.Sprintf("%t", complaint.IsPublic),
			complaint.CreatedAt,
		}
		if err := csvWriter.Write(record); err != nil {
			return nil, err
		}
		no++
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return nil, err
	}

	return buffer, nil
}
