package response

import "github.com/taranovegor/jurnalo/internal/model/request"

type Paginated struct {
	TotalCount int64       `json:"totalCount"`
	Limit      int64       `json:"limit"`
	Offset     int64       `json:"offset"`
	Items      interface{} `json:"items"`
}

func NewPaginated(totalCount int64, limit int64, offset int64, items interface{}) Paginated {
	return Paginated{
		TotalCount: totalCount,
		Limit:      limit,
		Offset:     offset,
		Items:      items,
	}
}

func NewPaginatedFromPaginator(p request.Paginator, totalCount int64, items interface{}) Paginated {
	return NewPaginated(totalCount, p.Limit, p.Offset, items)
}
