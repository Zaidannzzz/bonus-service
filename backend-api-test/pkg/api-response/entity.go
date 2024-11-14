package api_response

/*
   |--------------------------------------------------------------------------
   | Default Struct Response API
   |--------------------------------------------------------------------------
   |
   | You can change this every moment when do you want    |
*/

type Response struct {
	Meta    interface{} `json:"meta"`
	Message interface{} `json:"message"`
	Status  uint        `json:"status"`
	Code    uint        `json:"code"`
	Data    interface{} `json:"data"`
}

/*
   |--------------------------------------------------------------------------
   | Default Struct Pagination data
   |--------------------------------------------------------------------------
   |
*/

type PaginationQuery struct {
	Limit  *int `form:"limit,default=10" binding:"required,number"`
	Page   *int `form:"page,default=1" binding:"required,number"`
	Offset *int `form:"offset,default=0" binding:"omitempty,number"`
}

/*
   |--------------------------------------------------------------------------
   | Default Struct Time
   |--------------------------------------------------------------------------
   |
*/

type TimeResp struct {
	CreatedAt *uint `json:"created_at"`
	UpdatedAt *uint `json:"updated_at"`
}
