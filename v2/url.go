package twitter

type Url struct {
	Url         string `json:"url"`
	ExpandedUrl string `json:"expanded_url"`
	Status      int64  `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UnwoundUrl  string `json:"unwound_url"`
}
