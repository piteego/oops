package oops

var levelString = map[Level]string{
	Critical:      "Critical",
	High:          "High",
	Medium:        "Medium",
	Low:           "Low",
	Informational: "Informational",
}

func (l Level) String() string {
	if str, exists := levelString[l]; exists {
		return str
	}
	return "Unknown"
}
