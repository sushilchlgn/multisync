package sessions

type Session struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Device   string `json:"device"`
	URL      string `json:"url"`
	Status   string `json:"status"`
}