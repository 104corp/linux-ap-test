package issue

type Issue struct {
	HtmlURL string  `json:"html_url"`
	Body    string  `json:"body"`
	Labels  []Label `json:"labels"`
}
