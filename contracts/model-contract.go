package contracts

type ModelContract struct {
	Model      string             `json:"model"`
	Properties []PropertyContract `json:"properties"`
}
