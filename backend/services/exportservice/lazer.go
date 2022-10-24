package exportservice

import (
	"backend/models/entities"
	"encoding/json"
	"io"
	"time"
)

type LazerRoundExport struct {
	Rounds []ExportRound `json:"Rounds"`
}

type ExportRound struct {
	Name        string          `json:"Name"`
	Description string          `json:"Description"`
	BestOf      int             `json:"BestOf"`
	Beatmaps    []ExportBeatmap `json:"Beatmaps"`
	StartDate   time.Time       `json:"StartDate"`
	Matches     []int           `json:"Matches"`
}

type ExportBeatmap struct {
	ID   int    `json:"ID"`
	Mods string `json:"Mods"`
}

const EXPORT_LAZER = "lazer"

func ExportLazer(writeTo io.Writer, round entities.Round) error {
	jsonWriter := json.NewEncoder(writeTo)
	jsonWriter.SetIndent("", "  ")

	var exportRes LazerRoundExport

	exportRes.Rounds = append(exportRes.Rounds, ExportRound{
		Name:        round.Name,
		Description: "Description Here",
		BestOf:      727,
		StartDate:   time.Now().Add(time.Hour * 24 * 7),
		Matches:     []int{},
		Beatmaps:    []ExportBeatmap{},
	})

	for _, m := range round.Mappool {
		exportRes.Rounds[0].Beatmaps = append(exportRes.Rounds[0].Beatmaps, ExportBeatmap{
			ID:   int(m.BeatmapId),
			Mods: m.SlotName(),
		})
	}

	return jsonWriter.Encode(exportRes)
}
