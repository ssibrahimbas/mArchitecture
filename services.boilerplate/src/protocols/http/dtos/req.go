package dtos

type CreateExampleRequest struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type ListExampleRequest struct {
	Offset *int `query:"offset" validate:"required"`
	Limit  *int `query:"limit" validate:"required"`
}

type GetExampleRequest struct {
	Key string `param:"key" validate:"required"`
}

type UpdateExampleRequest struct {
	Key   string `param:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}
