package portal

type RestModel struct {
	Id          string `json:"-"`
	Name        string `json:"name"`
	Target      string `json:"target"`
	Type        uint8  `json:"type"`
	X           int16  `json:"x"`
	Y           int16  `json:"y"`
	TargetMapId uint32 `json:"targetMapId"`
	ScriptName  string `json:"scriptName"`
}

func (r RestModel) GetName() string {
	return "portals"
}

func (r RestModel) GetID() string {
	return r.Id
}

func (r *RestModel) SetID(id string) error {
	r.Id = id
	return nil
}
