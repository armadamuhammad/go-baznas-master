package model

// Page pagination model
type Page struct {
	Items      interface{} `json:"items" swaggertype:"object"` // items of data
	Page       int64       `json:"page" example:"0"`           // current page, start from zero
	Size       int64       `json:"size" example:"10"`          // size per page, default `10`
	MaxPage    int64       `json:"max_page" example:"9"`       // maximum pages for current schema
	TotalPages int64       `json:"total_pages" example:"10"`   // total pages
	Total      int64       `json:"total" example:"100"`        // total data without filters
	First      bool        `json:"first" example:"true"`       // indicate first value
	Last       bool        `json:"last" example:"false"`       // indicate last value
	Visible    int64       `json:"visible" example:"10"`       // current length
}
