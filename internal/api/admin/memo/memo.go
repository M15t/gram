package memo

import (
	"himin-runar/internal/rbac"
	"himin-runar/internal/types"
	"himin-runar/pkg/server"

	contextutil "himin-runar/internal/api/context"

	structutil "himin-runar/pkg/util/struct"
)

// Create creates new memo
func (s *Memo) Create(c contextutil.Context, data CreateMemoReq) (*types.Memo, error) {
	if err := s.enforce(c, rbac.ActionCreateAll); err != nil {
		return nil, err
	}

	// ! add custom logic here

	rec := &types.Memo{
		UserID:  c.AuthUser().ID,
		Content: data.Content,
	}

	if err := s.repo.Memo.Create(c.GetContext(), rec); err != nil {
		return nil, server.NewHTTPInternalError("error creating memo").SetInternal(err)
	}

	return rec, nil
}

// Read returns single memo by id
func (s *Memo) Read(c contextutil.Context, id string) (*types.Memo, error) {
	if err := s.enforce(c, rbac.ActionReadAll); err != nil {
		return nil, err
	}

	rec := &types.Memo{}
	if err := s.repo.Memo.ReadByID(c.GetContext(), rec, id); err != nil {
		return nil, server.NewHTTPInternalError("error reading memo").SetInternal(err)
	}

	return rec, nil
}

// List returns the list of memos
func (s *Memo) List(c contextutil.Context, req ListMemoReq) (*ListMemosResp, error) {
	if err := s.enforce(c, rbac.ActionReadAll); err != nil {
		return nil, err
	}

	// ! there are 3 ways to initialize filter and maybe more to be explored
	// * 1. using default
	// * initialize filter

	var count int64 = 0
	data := []*types.Memo{}
	filter := map[string]any{}
	lqc := req.ToListQueryCond([]any{filter})
	if err := s.repo.Memo.ReadAllByCondition(c.GetContext(), &data, &count, lqc); err != nil {
		return nil, server.NewHTTPInternalError("Error listing memo").SetInternal(err)
	}

	// * 2. add filter directly from request

	// ! this will be translated to "first_name LIKE %req.Name%"
	// ! any other filter that use gowhere must be added before mapping to ListQueryCondition

	// var count int64 = 0
	// data := []*types.Memo{}
	// filter := map[string]any{}
	// filter["something__contains"] = req.Something
	// filter["role"] = "admin"
	// lqc := req.ToListQueryCond([]any{filter})
	// if err := s.repo.Memo.ReadAllByCondition(c.GetContext(), &data, &count, lqc); err != nil {
	// 	return nil, server.NewHTTPInternalError("Error listing memo").SetInternal(err)
	// }

	// * 3. using custom filter
	// * that defines in type.go
	// * the logic will be processed in repo

	// var count int64 = 0
	// data := []*types.Memo{}
	// if err := s.repo.Memo.List(c.GetContext(), &data, &count, req.ToListCond()); err != nil {
	// 	return nil, server.NewHTTPInternalError("Error listing memo").SetInternal(err)
	// }

	return &ListMemosResp{
		Data:       data,
		TotalCount: count,
	}, nil
}

// Update updates memo information
func (s *Memo) Update(c contextutil.Context, id string, data UpdateMemoReq) (*types.Memo, error) {
	if err := s.enforce(c, rbac.ActionUpdateAll); err != nil {
		return nil, err
	}

	if err := s.repo.Memo.Update(c.GetContext(), structutil.ToMap(data), id); err != nil {
		return nil, server.NewHTTPInternalError("error reading memo").SetInternal(err)
	}

	return s.Read(c, id)
}

// Delete deletes memo by id
func (s *Memo) Delete(c contextutil.Context, id string) error {
	if err := s.enforce(c, rbac.ActionDeleteAll); err != nil {
		return err
	}

	if existed, err := s.repo.Memo.Existed(c.GetContext(), id); err != nil || !existed {
		return ErrMemoNotFound.SetInternal(err)
	}

	return s.repo.Memo.Delete(c.GetContext(), id)
}

// enforce checks memo permission to perform the action
func (s *Memo) enforce(c contextutil.Context, action string) error {
	au := c.AuthUser()
	if au == nil || !s.rbac.Enforce(au.Role, rbac.ObjectMemo, action) {
		return rbac.ErrForbiddenAction
	}
	return nil
}
