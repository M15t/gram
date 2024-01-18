package memo

import (
	"himin-runar/internal/types"
	httputil "himin-runar/pkg/util/http"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	contextutil "himin-runar/internal/api/context"
)

// HTTP represents memo http service
type HTTP struct {
	contextutil.Context
	svc Service
}

// Service represents memo application interface
type Service interface {
	Create(contextutil.Context, CreateMemoReq) (*types.Memo, error)
	Read(contextutil.Context, string) (*types.Memo, error)
	List(contextutil.Context, ListMemoReq) (*ListMemosResp, error)
	Update(contextutil.Context, string, UpdateMemoReq) (*types.Memo, error)
	Delete(contextutil.Context, string) error
}

// NewHTTP attaches handlers to Echo routers under given group
func NewHTTP(svc Service, eg *echo.Group) {
	h := HTTP{svc: svc}

	// swagger:operation POST /admin/memos admin-memos memosCreate
	// ---
	// summary: Creates new memo
	// parameters:
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/CreateMemoReq"
	// responses:
	//   "200":
	//     description: The new memo
	//     schema:
	//       "$ref": "#/definitions/Memo"
	//   default:
	//     description: 'Possible errors: 400, 401, 403, 500'
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"
	eg.POST("", h.create)

	// swagger:operation GET /admin/memos/{id} admin-memos memosRead
	// ---
	// summary: Returns a single memo
	// parameters:
	// - name: id
	//   in: path
	//   description: id of memo
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     description: The memo
	//     schema:
	//       "$ref": "#/definitions/Memo"
	//   default:
	//     description: 'Possible errors: 400, 401, 403, 404, 500'
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"
	eg.GET("/:id", h.read)

	// swagger:operation GET /admin/memos admin-memos memosList
	// ---
	// summary: Returns list of memos
	// responses:
	//   "200":
	//     description: List of memos
	//     schema:
	//       "$ref": "#/definitions/ListMemosResp"
	//   default:
	//     description: 'Possible errors: 400, 401, 403, 500'
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"
	eg.GET("", h.list)

	// swagger:operation PATCH /admin/memos/{id} admin-memos memosUpdate
	// ---
	// summary: Updates memo information
	// parameters:
	// - name: id
	//   in: path
	//   description: id of memo
	//   type: string
	//   required: true
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/UpdateMemoReq"
	// responses:
	//   "200":
	//     description: The updated memo
	//     schema:
	//       "$ref": "#/definitions/Memo"
	//   default:
	//     description: 'Possible errors: 400, 401, 403, 404, 500'
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"
	eg.PATCH("/:id", h.update)

	// swagger:operation DELETE /admin/memos/{id} admin-memos memosDelete
	// ---
	// summary: Deletes an memo
	// parameters:
	// - name: id
	//   in: path
	//   description: id of memo
	//   type: string
	//   required: true
	// responses:
	//   "204":
	//     "$ref": "#/responses/ok"
	//   default:
	//     description: 'Possible errors: 401, 403, 404, 500'
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"
	eg.DELETE("/:id", h.delete)
}

func (h *HTTP) create(c echo.Context) error {
	r := CreateMemoReq{}
	if err := c.Bind(&r); err != nil {
		return err
	}
	r.Content = strings.TrimSpace(r.Content)

	resp, err := h.svc.Create(contextutil.NewContext(c), r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) read(c echo.Context) error {
	id, err := httputil.ReqID(c)
	if err != nil {
		return err
	}
	resp, err := h.svc.Read(contextutil.NewContext(c), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) list(c echo.Context) error {
	req := ListMemoReq{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	resp, err := h.svc.List(contextutil.NewContext(c), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) update(c echo.Context) error {
	id, err := httputil.ReqID(c)
	if err != nil {
		return err
	}
	r := UpdateMemoReq{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	r.Content = httputil.TrimSpacePointer(r.Content)

	resp, err := h.svc.Update(contextutil.NewContext(c), id, r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) delete(c echo.Context) error {
	id, err := httputil.ReqID(c)
	if err != nil {
		return err
	}
	if err := h.svc.Delete(contextutil.NewContext(c), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
