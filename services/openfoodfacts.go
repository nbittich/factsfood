package services

import (
	"context"

	"github.com/nbittich/factsfood/services/db"
	"github.com/nbittich/factsfood/types/openfoodfacts"
	"go.mongodb.org/mongo-driver/bson"
)

type OFFService struct{}

const (
	OpenFoodFactsCollection = "openfoodfacts"
)

func (service *OFFService) Search(ctx context.Context, criteria *openfoodfacts.OFFSearchCriteria) ([]openfoodfacts.OpenFoodFactCSVEntry, error) {
	col := db.GetCollection(OpenFoodFactsCollection)
	filters := make([]bson.M, 0, 2)
	if criteria.Code != "" {
		filters = append(filters, db.FilterByID(criteria.Code))
	}
	if criteria.Name != "" {
		filters = append(filters, bson.M{
			"$or": []bson.M{
				{"productName": criteria.Name},
				{"genericName": criteria.Name},
			},
		})
	}
	// find all
	if len(filters) == 0 {
		return db.FindAll[openfoodfacts.OpenFoodFactCSVEntry](ctx, col, &criteria.Page)
	} else {
		filter := bson.M{
			"$and": filters,
		}
		return db.Find[openfoodfacts.OpenFoodFactCSVEntry](ctx, &filter, col, &criteria.Page)
	}
}
