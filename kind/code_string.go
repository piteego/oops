package kind

var stringOf = map[Code]string{
	Internal:      "Internal Server Error",
	Unprocessable: "Unprocessable",
	Validation:    "Validation Error",
	Unauthorized:  "Unauthorized Error",
	Forbidden:     "Forbidden Error",
	NotFound:      "Not Found Error",
	Duplication:   "Duplication Error",
}

func (c Code) String() string {
	if str, exists := stringOf[c]; exists {
		return str
	}
	return "Unknown Error"
}
