package assistservice

import "time"

type DiseasesPayload struct {
	Fiebre float64 `json:"fiebre"`
	CongestionNasal float64 `json:"congestion_nasal"`
	Nauseas float64 `json:"nauseas"`
	Vomitos float64 `json:"vomitos"`
	Escalofrios float64 `json:"escalofrios"`
	DolorDeCabeza float64 `json:"dolor_de_cabeza"`
	DolorMuscular float64 `json:"dolor_muscular"`
	DolorDeHuesos float64 `json:"dolor_de_huesos"`
	DolorDeGarganta float64 `json:"dolor_de_garganta"`
	Cansancio float64 `json:"cansancio"`
	TosSeca float64 `json:"tos_seca"`
	TosProductiva float64 `json:"tos_productiva"`
	FaltaDeAireAlRespirar float64 `json:"falta_de_aire_al_respirar"`
}

type DiseasesWeight struct {
	Fiebre float64 `json:"fiebre"`
	CongestionNasal float64 `json:"congestion_nasal"`
	Nauseas float64 `json:"nauseas"`
	Vomitos float64 `json:"vomitos"`
	Escalofrios float64 `json:"escalofrios"`
	DolorDeCabeza float64 `json:"dolor_de_cabeza"`
	DolorMuscular float64 `json:"dolor_muscular"`
	DolorDeHuesos float64 `json:"dolor_de_huesos"`
	DolorDeGarganta float64 `json:"dolor_de_garganta"`
	Cansancio float64 `json:"cansancio"`
	TosSeca float64 `json:"tos_seca"`
	TosProductiva float64 `json:"tos_productiva"`
	FaltaDeAireAlRespirar float64 `json:"falta_de_aire_al_respirar"`
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
