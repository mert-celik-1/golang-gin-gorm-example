package response

type ProductsResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ProductGetResponse struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type ProductUpdateResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
