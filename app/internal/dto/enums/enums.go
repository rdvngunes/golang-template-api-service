package enums

// AltitudeMeasurementErrorEnum represents an enum for altitude measurement errors
var AltitudeMeasurementError = map[string]string{
	"BAROMETRIC": "BAROMETRIC",
	"GPS_BASED":  "GPS_BASED",
}

// ContingencyManoeuvreLateralEnum represents an enum for lateral contingency manoeuvres
var ContingencyManoeuvreLateral = map[string]string{
	"STOPPING":              "STOPPING",
	"TURN_180_DEGREES":      "TURN_180_DEGREES",
	"PARACHUTE_TERMINATION": "PARACHUTE_TERMINATION",
}

// AwsFileType
var AwsFileType = map[string]string{
	"operation":        "operation",
	"drone":            "drone",
	"commandcontrol":   "commandcontrol",
	"waypointskml":     "waypointskml",
	"adjacentairspace": "adjacentairspace",
	"adjacentarea":     "adjacentarea",
	"containment":      "containment",
}

// ContingencyManoeuvreVerticalEnum represents an enum for vertical contingency manoeuvres
var ContingencyManoeuvreVertical = map[string]string{
	"KINETIC_TO_POTENTIAL":           "KINETIC_TO_POTENTIAL",
	"CIRCULAR_PATH_45_DEGREES_PITCH": "CIRCULAR_PATH_45_DEGREES_PITCH",
	"PARACHUTE_TERMINATION":          "PARACHUTE_TERMINATION",
}

// TerminationMethodEnum represents an enum for termination methods
var TerminationMethod = map[string]string{
	"SIMPLIFIED_APPROACH":   "SIMPLIFIED_APPROACH",
	"BALLISTIC_APPROACH":    "BALLISTIC_APPROACH",
	"PARACHUTE_TERMINATION": "PARACHUTE_TERMINATION",
	"SWITCH_OFF_GLIDING":    "SWITCH_OFF_GLIDING",
	"SWITCH_OFF_NO_GLIDING": "SWITCH_OFF_NO_GLIDING",
}

// MultiRotorLateralEnum represents an enum for multi-rotor lateral manoeuvres
var MultiRotorLateral = map[string]string{
	"STOPPING":              "STOPPING",
	"PARACHUTE_TERMINATION": "PARACHUTE_TERMINATION",
}

// MultiRotorVerticalEnum represents an enum for multi-rotor vertical manoeuvres
var MultiRotorVertical = map[string]string{
	"KINETIC_TO_POTENTIAL":  "KINETIC_TO_POTENTIAL",
	"PARACHUTE_TERMINATION": "PARACHUTE_TERMINATION",
}

// MultiRotorTerminationEnum represents an enum for multi-rotor termination methods
var MultiRotorTermination = map[string]string{
	"SIMPLIFIED_APPROACH":   "SIMPLIFIED_APPROACH",
	"BALLISTIC_APPROACH":    "BALLISTIC_APPROACH",
	"PARACHUTE_TERMINATION": "PARACHUTE_TERMINATION",
}

// FixedWingsLateralEnum represents an enum for fixed-wing lateral manoeuvres
var FixedWingsLateral = map[string]string{
	"TURN_180_DEGREES":      "TURN_180_DEGREES",
	"PARACHUTE_TERMINATION": "PARACHUTE_TERMINATION",
}

// FixedWingsVerticalEnum represents an enum for fixed-wing vertical manoeuvres
var FixedWingsVertical = map[string]string{
	"CIRCULAR_PATH_45_DEGREES_PITCH": "CIRCULAR_PATH_45_DEGREES_PITCH",
	"PARACHUTE_TERMINATION":          "PARACHUTE_TERMINATION",
}

// FixedWingsTerminationEnum represents an enum for fixed-wing termination methods
var FixedWingsTermination = map[string]string{
	"SWITCH_OFF_GLIDING":    "SWITCH_OFF_GLIDING",
	"SWITCH_OFF_NO_GLIDING": "SWITCH_OFF_NO_GLIDING",
	"PARACHUTE_TERMINATION": "PARACHUTE_TERMINATION",
}

// ConopsActions
var ActionType = map[string]string{
	"Conops":   "Conops",
	"Grc":      "Grc",
	"Arc":      "Arc",
	"SailOso":  "SailOso",
	"Adjacent": "Adjacent",
	"Final":    "Final",
}

// Application Steps
var ApplicationSteps = map[string]string{
	"Draft":                        "Draft",
	"Reviewer_Assigned":            "Reviewer_Assigned",
	"In_Progress":                  "In_Progress",
	"Requested_Info_Submitted":     "Requested_Info_Submitted",
	"More_Info_Requested":          "More_Info_Requested",
	"Submit_to_Authority":          "Submit_to_Authority",
	"Submitted_To_Senior_Reviewer": "Submitted_To_Senior_Reviewer",
	"Approved":                     "Approved",
	"Declined":                     "Declined",
	"Shared":                       "Shared",
}

// Application Steps
var AssessmentStatus = map[string]string{
	"Draft": "Draft",
	"Saved": "Saved",
}

// Application Type
var ApplicationType = map[string]string{
	"Renewal":         "Renewal",
	"Extended_Area":   "Extended_Area",
	"New_Application": "New_Application",
}

// Assessment Category Status
var AssessmentCategoryStatus = map[string]string{
	"Awaiting":            "Awaiting",
	"More_Info_Requested": "More_Info_Requested",
	"Approved":            "Approved",
	"Declined":            "Declined",
}

// ArcResults
var ArcResults = map[string]string{
	"Arc-a": "Arc-a",
	"Arc-b": "Arc-b",
	"Arc-c": "Arc-c",
	"Arc-d": "Arc-d",
}

// YesNoCondition
var YesNoCondition = map[string]string{
	"Yes": "yes",
	"No":  "no",
}

// DataSourceType
var DataSourceType = map[string]string{
	"conops":      "conops",
	"grc":         "grc",
	"arc":         "arc",
	"oso":         "oso",
	"containment": "containment",
}

// ArcRpasEnvironments
var ArcRpasEnvironments = map[string]string{
	"CLASS_BCD":                         "CLASS_BCD",
	"CLASS_EFG":                         "CLASS_EFG",
	"URBAN_AREA_BELOW_500":              "URBAN_AREA_BELOW_500",
	"VEIL_TMZ_BELOW_500":                "VEIL_TMZ_BELOW_500",
	"UNCONTROLLED_OVER_RURAL_BELOW_500": "UNCONTROLLED_OVER_RURAL_BELOW_500",
	"ATYPICAL_AIRSPACE":                 "ATYPICAL_AIRSPACE",
	"VEIL_TMZ":                          "VEIL_TMZ",
	"URBAN_AREA":                        "URBAN_AREA",
	"UNCONTROLLED_OVER_RURAL":           "UNCONTROLLED_OVER_RURAL",
	"IN_FL_600":                         "IN_FL_600",
	"CONTROLLED_AIRSPACE":               "CONTROLLED_AIRSPACE",
	"CONTROLLED_OUTSIDE_AIRPORT":        "CONTROLLED_OUTSIDE_AIRPORT",
	"CONTROLLED_AIRSPACE_BELOW_500":     "CONTROLLED_AIRSPACE_BELOW_500",
}

// ArcRpasEnvironments
var ApplicationListingStatus = map[string]string{
	"in_review":       "in_review",
	"new_application": "new_application",
	"approved":        "approved",
	"declined":        "declined",
}

// HeightUnit
var HeightUnit = map[string]string{
	"FT": "FT",
	"M":  "M",
}

// ArcRpasEnvironments
var FootPoundEnergyUnit = map[string]string{
	"j":  "j",
	"kj": "kj",
}

// SpeedUnit
var SpeedUnit = map[string]string{
	"mph":  "mph",
	"km/h": "km/h",
}

// AltitudeUnit
var AltitudeUnit = map[string]string{
	"m": "m",
}

// HeightUnit
var GrcMitigation = map[string]string{
	"N":       "N",
	"L":       "L",
	"M":       "M",
	"H":       "H",
	"H_MIN_2": "H_MIN_2",
	"H_MIN_3": "H_MIN_3",
}

// DroneType
var DroneType = map[string]string{
	"FIXED_WINGS": "FIXED_WINGS",
	"MULTI_ROTOR": "MULTI_ROTOR",
	"HELICOPTER":  "HELICOPTER",
}

// AdjAreaRequestTerminationMethod
var AdjAreaRequestTerminationMethod = map[string]string{
	"SIMPLIFIED_APPROACH":   "SIMPLIFIED_APPROACH",
	"BALLISTIC_APPROACH":    "BALLISTIC_APPROACH",
	"SWITCH_OFF_GLIDING":    "SWITCH_OFF_GLIDING",
	"SWITCH_OFF_NO_GLIDING": "SWITCH_OFF_NO_GLIDING",
	"PARACHUTE_TERMINATION": "PARACHUTE_TERMINATION",
}

// Automation
var ArcAutomation = map[string]string{
	"airspace_classes":        "airspace_classes",
	"operational_environment": "operational_environment",
}

// Sort Order
var SortOrder = map[string]string{
	"asc":  "asc",
	"desc": "desc",
}

// Order
var ApplicationSorting = map[string]string{
	"application_no":   "application_no",
	"application_type": "application_type",
	"organization":     "organization",
	"name":             "name",
	"status":           "status",
	"grc":              "grc",
	"arc":              "arc",
	"contact_user":     "contact_user",
	"submission_time":  "submission_time",
	"decision_date":    "decision_date",
	"sail":             "sail",
}

// Order
var AssessmentSorting = map[string]string{
	"created_date":    "created_date",
	"updated_date":    "updated_date",
	"assessment_name": "assessment_name",
	"status":          "status",
	"final_grc":       "final_grc",
	"residual_arc":    "residual_arc",
	"sail":            "sail",
}

// Order
var UserSorting = map[string]string{
	"created_date":   "created_date",
	"first_name":     "first_name",
	"user_sora_role": "user_sora_role",
	"is_active":      "is_active",
	"email":          "email",
	"phone":          "phone",
}

// User Roles
var UserRoles = map[string]string{
	"SORAUSER":             "SORAUSER",
	"SORA_REVIEWER":        "SORA_REVIEWER",
	"SORA_SENIOR_REVIEWER": "SORA_SENIOR_REVIEWER",
	"SORA_REVIEWER_ADMIN":  "SORA_REVIEWER_ADMIN",
	"PILOT":                "PILOT",
}

// Object Keys
var ObjectKeys = map[string]string{
	"Conops":   "ConopsObjectKey",
	"Kml":      "KmlObjectKey",
	"Pdf":      "PdfObjectKey",
	"User":     "UserObjectKey",
	"Adjacent": "AdjacentObjectKey",
}
