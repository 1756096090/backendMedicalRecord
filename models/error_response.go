package models

type ErrorResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    HTTPStatus int  `json:"http_status"`
}