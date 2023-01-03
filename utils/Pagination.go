package utils

import "math"

type Pagination struct {
	Page int
	Vertial int     // 세로
	Total int

	Horizontal int  // 가로
}


func (p *Pagination) Pages() int {

	pages := int( math.Ceil( float64(p.Total)/float64(p.Vertial) ) )

	return pages
}


