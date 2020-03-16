package assistservice

import "time"

type DiseasesPayload struct {
	MusclePain      float64 `json:"muscle_pain"`
	Breath          float64 `json:"breath"`
	NasalCongestion float64 `json:"nasal_congestion"`
	Headache        float64 `json:"head_ache"`
	BoneAche        float64 `json:"bone_ache"`
	SoreThroat      float64 `json:"sore_throat"`
	Tiredness       float64 `json:"tiredness"`
	DryCaught       float64 `json:"dry_cought"`
}

type DiseasesWeight struct {
	MusclePain      float64 `json:"muscle_pain"`
	Breath          float64 `json:"breath"`
	NasalCongestion float64 `json:"nasal_congestion"`
	Headache        float64 `json:"head_ache"`
	BoneAche        float64 `json:"bone_ache"`
	SoreThroat      float64 `json:"sore_throat"`
	Tiredness       float64 `json:"tiredness"`
	DryCaught       float64 `json:"dry_cought"`
}

type PatientRecord struct {
	ID               string          `json:"id"`
	CreatedAt        time.Time       `json:"created_at"`
	PatientID        string          `json:"patient_id"`
	Input            DiseasesPayload `json:"input"`
	EvaluatedWeight  DiseasesWeight  `json:"evaluated_weight"`
	EvaluationResult float64         `json:"evaluation_result"`
}

type PatientProfile struct {
	ID      string          `json:"id"`
	DNI     string          `json:"dni"`
	Phone   string          `json:"phone"`
	Email   string          `json:"email"`
	Records []PatientRecord `json:"records"`
}
