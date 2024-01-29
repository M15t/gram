package memo

import (
	"gram/internal/types"
	requestutil "gram/pkg/util/request"
)

// CreateMemoReq contains request data to create new memo
// swagger:model
type CreateMemoReq struct {
	Content string `json:"content" validate:"required"`
}

// UpdateMemoReq contains request data to update existing memo
// swagger:model
type UpdateMemoReq struct {
	Content *string `json:"content,omitempty"`
}

// ChangePasswordReq contains request data to change memo password
// swagger:model
type ChangePasswordReq struct {
	OldPassword        string `json:"old_password" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required,min=8"`
	NewPasswordConfirm string `json:"new_password_confirm" validate:"required,eqfield=NewPassword"`
}

// ListMemoReq contains request data to get list of countries
// swagger:parameters memosList
type ListMemoReq struct {
	requestutil.ListQueryRequest
}

// ListMemosResp contains list of paginated memos and total numbers after filtered
// swagger:model
type ListMemosResp struct {
	Data       []*types.Memo `json:"data"`
	TotalCount int64         `json:"total_count"`
}

// ToListCond transforms the service request to repo conditions
// func (lq *ListMemoReq) ToListCond() *requestutil.ListCondition[repo.MemosFilter] {
// 	return &requestutil.ListCondition[repo.MemosFilter]{
// 		Page:    lq.Page,
// 		PerPage: lq.PerPage,
// 		Sort:    lq.Sort,
// 		Count:   true,
// 		Filter: repo.MemosFilter{
// 			Search: lq.Search,
// 		},
// 	}
// }
