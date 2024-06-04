package openfoodfacts

import "github.com/nbittich/factsfood/services/db"

const (
	OpenFoodFactsImgCollection = "openfoodfacts_img"
	OpenFoodFactsCollection    = "openfoodfacts"
	InitialCapLogs             = 10  // Logs slice initial capacity
	sleepBetweenBatchesMs      = 100 // Sleep 100ms to allow mongodb between each batch to rest a lil bit.
	//                               //   If you change this, make sure you  also change the job config in db.
	//                               //   See BatchSize100Ms below
)

var offCollection = db.GetCollection(OpenFoodFactsCollection)
