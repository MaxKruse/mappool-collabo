package exportservice

import (
	"backend/models/entities"
	"encoding/csv"
	"fmt"
	"io"
)

func ExportCSV(writeTo io.Writer, maps []entities.Map) error {
	// write the csv header like follows:
	// Banner, Mod, ID, Artist - Title [Difficulty], SR, BPM, Length, CS, AR, OD, Mapper

	headers := []string{"Banner", "Mod", "ID", "Artist - Title [Difficulty]", "SR", "BPM", "Length", "CS", "AR", "OD", "Mapper"}

	csvWriter := csv.NewWriter(writeTo)
	defer csvWriter.Flush()

	// write the header to the io.Writer
	if err := csvWriter.Write(headers); err != nil {
		return err
	}

	// then loop through the maps and check if the fields exist. If they do not exist, return an error
	// if they do exist, write them to the writeTo io.Writer
	for _, m := range maps {
		// check the fields

		banner := fmt.Sprintf("https://assets.ppy.sh/beatmaps/%d/covers/fullsize.jpg", m.BeatmapId)
		mod := m.SlotName()
		id := m.BeatmapId
		title := m.Name
		sr := m.Difficulty.Stars
		bpm := m.Difficulty.BPM
		length := m.Difficulty.Length
		cs := m.Difficulty.CS
		ar := m.Difficulty.AR
		od := m.Difficulty.OD
		mapper := m.Creator

		// check all these fields for nil or empty values
		if banner == "" {
			return fmt.Errorf("banner is empty")
		}

		if mod == "" {
			return fmt.Errorf("mod is empty")
		}

		if id == 0 {
			return fmt.Errorf("id is empty")
		}

		if title == "" {
			return fmt.Errorf("title is empty")
		}

		if sr == 0 {
			return fmt.Errorf("sr is empty")
		}

		if bpm == 0 {
			return fmt.Errorf("bpm is empty")
		}

		if length == 0 {
			return fmt.Errorf("length is empty")
		}

		if mapper == "" {
			return fmt.Errorf("mapper is empty")
		}

		writeable := []string{banner, mod, fmt.Sprintf("%d", id), title, fmt.Sprintf("%f", sr), fmt.Sprintf("%d", bpm), fmt.Sprintf("%f", length), fmt.Sprintf("%f", cs), fmt.Sprintf("%f", ar), fmt.Sprintf("%f", od), mapper}

		if err := csvWriter.Write(writeable); err != nil {
			return err
		}

		csvWriter.Flush()
	}

	return nil
}
