package oops

var kindString = map[Kind]string{
	Internal:      "Internal Server Error",
	Unprocessable: "Unprocessable",
	Validation:    "Validation Error",
	Unauthorized:  "Unauthorized Error",
	Forbidden:     "Forbidden Error",
	NotFound:      "Not Found Error",
	Duplication:   "Duplication Error",
}

func (k Kind) String() string {
	if str, exists := kindString[k]; exists {
		return str
	}
	return "Unknown Error"
}
