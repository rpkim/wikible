package api

//PageContent is for wiki page content structure with json
type PageContent struct {
	ID     string `json:"id"`
	Space		  `json:"space"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	Version struct {
		Number int `json:"number"`
	} `json:"version"`
}