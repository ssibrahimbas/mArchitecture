package dtos

type CreateExampleRequest struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type ListExampleRequest struct {
	Offset *int `query:"offset"  validate:"omitempty,gt=0"`
	Limit  *int `query:"limit"  validate:"omitempty,gt=0"`
}

type GetExampleRequest struct {
	Key string `param:"key" validate:"required"`
}

type UpdateExampleRequest struct {
	Key   string `param:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

func (r *ListExampleRequest) Default() {
	if r.Offset == nil {
		r.Offset = new(int)
		*r.Offset = 0
	}
	if r.Limit == nil {
		r.Limit = new(int)
		*r.Limit = 10
	}
}
