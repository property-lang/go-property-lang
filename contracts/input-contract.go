package contracts

type OptionPropertyContract struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type PropertyContract struct {
	Key          string                   `json:"key"`
	Type         string                   `json:"type"`
	Min          int                      `json:"min"`
	Max          int                      `json:"max"`
	Required     bool                     `json:"required"`
	Placeholder  *string                  `json:"placeholder"`
	Label        string                   `json:"label"`
	Description  string                   `json:"description"`
	Prefix       *string                  `json:"prefix"`
	Postfix      *string                  `json:"postfix"`
	BeforeIcon   *string                  `json:"beforeIcon"`
	AfterIcon    *string                  `json:"afterIcon"`
	DefaultValue interface{}              `json:"defaultValue"`
	Tags         []string                 `json:"tags"`
	Options      []OptionPropertyContract `json:"options"` // Используем interface{} для объектов с любым типом
}

func (u *PropertyContract) GetFullName() string {
	return u.Key + " " + u.Type
}
