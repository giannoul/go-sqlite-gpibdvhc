package beer

// Beer is the model
type Beer struct {
	Id      int    `json:"id"`
	Uid     string `json:"uid"`
	Brand   string `json:"brand"`
	Name    string `json:"name"`
	Style   string `json:"style"`
	Hop     string `json:"hop"`
	Yeast   string `json:"yeast"`
	Malts   string `json:"malts"`
	Ibu     string `json:"ibu"`
	Alcohol string `json:"alcohol"`
	Blg     string `json:"blg"`
}
