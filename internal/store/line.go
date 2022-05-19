package store

type Line struct {
	User   string `json:"user,omitempty"`
	URL    string `json:"original_url"`
	Short  string `json:"short_url"`
	Status int    `json:"status"`
}
