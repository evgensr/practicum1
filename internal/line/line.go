package line

type Line struct {
	User   string `json:"user,omitempty"`
	Url    string `json:"original_url"`
	Short  string `json:"short_url"`
	Status int    `json:"status"`
}
