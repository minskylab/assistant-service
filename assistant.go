package assistservice

func GetReport(payload DiseasesPayload, weight DiseasesWeight) float64 {
	fiebre := payload.Fiebre * weight.Fiebre
	congestionNasal := payload.CongestionNasal * weight.CongestionNasal
	nauseas := payload.Nauseas * weight.Nauseas
	vomitos := payload.Vomitos * weight.Vomitos
	escalofrios := payload.Escalofrios * weight.Escalofrios
	dolorDeCabeza := payload.DolorDeCabeza * weight.DolorDeCabeza
	dolorMuscular := payload.DolorMuscular * weight.DolorMuscular
	dolorDeHuesos := payload.DolorDeHuesos * weight.DolorDeHuesos
	dolorDeGarganta := payload.DolorDeGarganta * weight.DolorDeGarganta
	cansancio := payload.Cansancio * weight.Cansancio
	tosSeca := payload.TosSeca * weight.TosSeca
	tosProductiva := payload.TosProductiva * weight.TosProductiva
	faltaDeAireAlRespirar := payload.FaltaDeAireAlRespirar * weight.FaltaDeAireAlRespirar

	result := fiebre + congestionNasal + nauseas + vomitos + escalofrios + dolorDeCabeza + dolorMuscular + dolorDeHuesos + dolorDeGarganta + cansancio + tosSeca + tosProductiva + faltaDeAireAlRespirar

	return result
}
