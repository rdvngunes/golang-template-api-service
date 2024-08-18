package utils

const (
	FROM_FT_TO_MT              = 0.3048
	FROM_FL_TO_FT              = 100
	FROM_FL_TO_MT              = FROM_FL_TO_FT * FROM_FT_TO_MT
	FROM_MT_TO_DEGREES_EQUATOR = 1e-5 / 1.11
	FROM_MT_TO_DEGREES_STRICT  = 1e-5
)

func ConvertToFt(alt string) int {
	var alt_msl_ft int

	// REFERENCE CAN BE -> MSL, GND, STD
	switch alt {
	// Reference: MSL => Units: F (or FT), M (or MT)

	// Reference: STD => Units: FL
	case "FL":
		alt_value := 1 // TODO find FL value
		alt_msl_ft = alt_value * FROM_FL_TO_FT

	// Reference: GND (or AGL) => Units: F (or FT), M (or MT)
	case "GND", "AGL":
	}

	//log.Printf("Converted %f%s %s to %fFT MSL.", alt.Value, alt.Units, alt.Reference, alt_msl_ft)
	return alt_msl_ft
}
