package store

type Line struct {
	User          string `json:"user,omitempty"`
	URL           string `json:"original_url"`
	Short         string `json:"short_url"`
	CorrelationID string `json:"correlation_id"`
	Status        int    `json:"status"`
}

type (
	// Line = store.Line
	Urls  int // количество сокращённых URL в сервисе
	Users int // количество пользователей в сервисе
)
