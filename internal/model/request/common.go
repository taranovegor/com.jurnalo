package request

type Paginator struct {
	Offset int64 `schema:"offset,default:0"`
	Limit  int64 `schema:"limit,default:10"`
}
