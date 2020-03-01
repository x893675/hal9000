package schema


type PageableResponse struct {
	Items      interface{} `json:"items" description:"paging data"`
	TotalCount int         `json:"total_count" description:"total count"`
}
