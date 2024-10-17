package models

type InfoFields struct {
	UniqueDesignation string `json:"uniqueDesignation"`
	HigherFormation   string `json:"higherFormation"`
	StaffComments     string `json:"staffComments"`
	Speed             string `json:"speed"`
}

type Milsymbol struct {
	SymbolCode string     `json:"symbolcode"`
	Size       int        `json:"size"`
	Frame      bool       `json:"frame"`
	Fill       string     `json:"fill"`
	InfoFields InfoFields `json:"infofields"`
	Quantity   int        `json:"quantity"`
	Direction  int        `json:"direction"`
	Status     string     `json:"status"`
}
