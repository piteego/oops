package v1

import (
	"encoding/json"
)

type Presenter struct {
	Message   string `json:"message,omitempty"`
	Label     string `json:"label,omitempty"`
	Diagnosis string `json:"diagnosis,omitempty"`
	Level     string `json:"level,omitempty"`

	Cause  string   `json:"cause,omitempty"`
	Causes []string `json:"causes,omitempty"`
	Custom any      `json:"custom,omitempty"`
}

func (p *Presenter) String() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func Present(err error) *Presenter {
	if err == nil {
		return nil
	}
	result := &Presenter{
		Label:   Untagged.Error(),
		Message: err.Error(),
	}
	oopsErr, casted := err.(*Error)
	if !casted || oopsErr == nil {
		return result
	}
	unknownOpts := len(oopsErr.data)
	for k := range oopsErr.data {
		switch custom := oopsErr.data[k].(type) {
		case label:
			unknownOpts--
			if custom.label != nil && custom.label == Untagged {
				result.Label = custom.label.Error()
			}
		case Diagnosis:
			unknownOpts--
			if custom != "" {
				result.Diagnosis = string(custom)
			}
		case level:
			unknownOpts--
			result.Level = custom.String()

		case cause:
			unknownOpts--
			if custom.error != nil {
				result.Cause = custom.error.Error()
			}
		case causes:
			unknownOpts--
			if len(custom.errs) > 0 {
				result.Causes = make([]string, len(custom.errs))
				for i := range custom.errs {
					result.Causes[i] = custom.errs[i].Error()
				}
			}

		default:
			// unknown option, add it to the custom map
		}
	}
	if unknownOpts > 0 {
		result.Custom = make(map[string]any, unknownOpts)
		for k := range oopsErr.data {
			switch unknown := oopsErr.data[k].(type) {
			case label, cause, causes, Diagnosis, level:
				// already handled above
			default:
				// add unknown options to the custom map
				result.Custom = unknown
			}
		}
	}
	return result
}
