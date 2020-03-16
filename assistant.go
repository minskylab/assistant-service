package assistservice

func NewDiseasesWeight() *DiseasesWeight {
	return &DiseasesWeight{
		MusclePain:      0.1,
		Breath:          0.1,
		NasalCongestion: 0.1,
		Headache:        0.1,
		BoneAche:        0.1,
		SoreThroat:      0.1,
		Tiredness:       0.1,
		DryCaught:       0.1,
	}
}

func GetReport(payload DiseasesPayload, weight DiseasesWeight) float64 {
	boneAche := payload.BoneAche * weight.BoneAche
	breath := payload.Breath * weight.Breath
	drycaught := payload.DryCaught * weight.DryCaught
	headache := payload.Headache * weight.Headache
	musclePain := payload.MusclePain * weight.MusclePain
	nasal := payload.NasalCongestion * weight.NasalCongestion
	soreThroat := payload.SoreThroat * weight.SoreThroat
	tired := payload.Tiredness * weight.Tiredness

	result := boneAche + breath + drycaught + headache + musclePain + nasal + soreThroat + tired

	return result
}
