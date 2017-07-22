package uhandlers

// IResponse : required behaviour for Response objects
type IResponse interface {
	Json() string
}

// Response : Base response struct
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// JSON : marshal base response to JSON format
func (res *Response) JSON() string {
	return ToJSON(res)
}

// UploadResponse : response struct for upload urls
type UploadResponse struct {
	Response
	Count int `json:"count"`
}

// JSON : overriding base imeplementation of JSON with in inherited struct UploadResponse
func (res *UploadResponse) JSON() string {
	return ToJSON(res)
}

// CheckURLResponse : response struct for url verification service
type CheckURLResponse struct {
	Response
}
