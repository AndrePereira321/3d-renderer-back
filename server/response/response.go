package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}

type ResponseWriter struct {
	http.ResponseWriter
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		w,
	}
}

func NewResponse() *Response {
	return &Response{}
}

func (w *ResponseWriter) WriteJSON(data *Response) error {
	w.Header().Set("Content-Type", "application/json")
	return w.WriteData(data)
}

func (w *ResponseWriter) WriteData(data interface{}) error {
	buff, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(buff)
	if err != nil {
		return err
	}
	return nil
}

func (w *ResponseWriter) WriteError(status int, code string) error {
	r := NewResponse()
	r.ErrorCode = code

	w.Header().Set("Content-Type", "application/json")
	buff, err := json.Marshal(r)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	_, err = w.Write(buff)
	if err != nil {
		return err
	}
	return nil
}

func (w *ResponseWriter) WriteErrorMessage(status int, code string, errorMsg string) error {
	r := NewResponse()
	r.ErrorCode = code
	r.ErrorMessage = errorMsg

	w.Header().Set("Content-Type", "application/json")
	buff, err := json.Marshal(r)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	_, err = w.Write(buff)
	if err != nil {
		return err
	}
	return nil
}
