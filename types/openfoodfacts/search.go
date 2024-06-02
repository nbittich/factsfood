package openfoodfacts

import "github.com/nbittich/factsfood/services/db"

type OFFSearchCriteria struct {
	Code string         `json:"code" form:"code" param:"code" query:"code"`
	Name string         `json:"name" form:"name" param:"name" query:"name"`
	Page db.PageOptions `json:"page" form:"page"`
}
