package postgres

// структура таблицы
type GoodResponse struct {
	Message  string `json:"message"        example:"message"`
	Code     int    `json:"code"           example:"status"`
	Affected int    `json:"affected_rows"  example:"0"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
