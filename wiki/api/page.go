package api

//Page struct is for creating the wiki page
type Page struct {
	Title     string `json:"title"`
	Type      string `json:"type"`
	Ancestors `json:"ancestors"`
	Space     `json:"space"`
}

type Ancestors []Ancestor

type Ancestor struct {
	AncestorID string `json:"id"`
}

//Space Key of wiki
type Space struct {
	SpaceID   int    `json:"id"`
	SpaceKey  string `json:"key"`
	SpaceName string `json:"name"`
}

//Body of wiki
type Body struct {
	Editor struct {
		Representation string `json:"representation"`
		Value          string `json:"value"`
	} `json:"editor"`
	ExportView struct {
		Representation string `json:"representation"`
		Value          string `json:"value"`
	} `json:"export_view"`
	Storage struct {
		Representation string `json:"representation"`
		Value          string `json:"value"`
	} `json:"storage"`
}

//Container of wiki
type Container struct {
	ContainerID   int    `json:"id"`
	ContainerKey  string `json:"key"`
	ContainerName string `json:"name"`
}
