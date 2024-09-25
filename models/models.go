package models

type Choice struct {
	Value   string `json:"value"`
	Correct bool   `json:"correct"`
}

func (c Choice) FilterValue() string {
	return c.Value
}

func (c Choice) Title() string {
	return c.Value
}
