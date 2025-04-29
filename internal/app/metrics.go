package app

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// WriteMetrics writes ElectricUsage data to a CSV file instead of InfluxDB.
func WriteMetrics(records []ElectricUsage, _ any) error {
	// Open the file
	file, err := os.Create("output.csv")
	if err != nil {
		return fmt.Errorf("failed to create output.csv: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"StartTime", "EndTime", "WattHours", "CostInCents", "MeterName"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write each record
	for _, record := range records {
		row := []string{
			record.StartTime.Format(time.RFC3339),
			record.EndTime.Format(time.RFC3339),
			strconv.FormatInt(record.WattHours, 10),
			"",
			"",
		}
		if record.CostInCents != nil {
			row[3] = strconv.FormatInt(*record.CostInCents, 10)
		}
		if record.MeterName != nil {
			row[4] = *record.MeterName
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write row: %w", err)
		}
	}

	return nil
}
