package data

// Data structure representing https://releases.aks.azure.com/data.json

type AKSReleaseStatusRemote struct {
	LastUpdateTime *string   `json:"LastUpdateTime,omitempty"`
	Regions        []*Region `json:"Regions,omitempty"`
}

// AKSReleaseStatusStorage omits LastUpdateTime, as it just adds noise
type AKSReleaseStatusStorage []*Region

type Region struct {
	Name       *string    `json:"Name,omitempty"`
	Alias      *string    `json:"Alias,omitempty"`
	Continent  *string    `json:"Continent,omitempty"`
	BatchIndex *int       `json:"BatchIndex,omitempty"`
	InProgress *bool      `json:"InProgress,omitempty"`
	Current    *Release   `json:"Current,omitempty"`
	Recent     []*Release `json:"Recent,omitempty"`
}

type Release struct {
	Version     *string `json:"Version,omitempty"`
	ReleaseNote *string `json:"ReleaseNote,omitempty"`
}
