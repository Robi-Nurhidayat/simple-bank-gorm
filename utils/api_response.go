package utils

type response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ApiResponse(msg string, code int, data interface{}) response {

	var res response
	res.Meta.Message = msg
	res.Meta.Code = code
	res.Data = data

	return res
}
