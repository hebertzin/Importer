package domains

type HttpResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Body    interface{} `json:"body,omitempty"`
}
