package openfoodfacts

import "github.com/nbittich/factsfood/types"

type FactsFood struct {
	OpenFoodFact    `bson:",inline"`
	OpenFoodFactImg *OpenFoodFactImg `json:"openfoodfact_img,omitempty"`
}
type OpenFoodFactImg struct {
	ID                       string `json:"_id" bson:"_id"`
	OpenFoodFactID           string `json:"openfoodfacts_id"`
	LastImageT               int    `json:"lastImageT,omitempty" bson:"lastImageT,omitempty"`
	ImageURL                 string `json:"imageURL,omitempty" bson:"imageURL,omitempty"`
	ImageSmallURL            string `json:"imageSmallURL,omitempty" bson:"imageSmallURL,omitempty"`
	ImageIngredientsURL      string `json:"imageIngredientsURL,omitempty" bson:"imageIngredientsURL,omitempty"`
	ImageIngredientsSmallURL string `json:"imageIngredientsSmallURL,omitempty" bson:"imageIngredientsSmallURL,omitempty"`
	ImageNutritionURL        string `json:"imageNutritionURL,omitempty" bson:"imageNutritionURL,omitempty"`
	ImageNutritionSmallURL   string `json:"imageNutritionSmallURL,omitempty" bson:"imageNutritionSmallURL,omitempty"`
}

type OpenFoodFact struct {
	Code                                            string            `json:"_id" bson:"_id"`
	URL                                             string            `json:"url,omitempty" bson:"url,omitempty"`
	Creator                                         string            `json:"-" bson:"-"`
	CreatedT                                        int               `json:"createdT,omitempty" bson:"createdT,omitempty"`
	CreatedDatetime                                 types.TimeISO8601 `json:"createdDatetime,omitempty" bson:"createdDatetime,omitempty"`
	LastModifiedT                                   int               `json:"lastModifiedT,omitempty" bson:"lastModifiedT,omitempty"`
	LastModifiedDatetime                            types.TimeISO8601 `json:"lastModifiedDatetime,omitempty" bson:"lastModifiedDatetime,omitempty"`
	LastModifiedBy                                  string            `json:"-" bson:"-"`
	LastUpdatedT                                    int               `json:"lastUpdatedT,omitempty" bson:"lastUpdatedT,omitempty"`
	LastUpdatedDatetime                             types.TimeISO8601 `json:"lastUpdatedDatetime,omitempty" bson:"lastUpdatedDatetime,omitempty"`
	ProductName                                     string            `json:"productName,omitempty" bson:"productName,omitempty"`
	AbbreviatedProductName                          string            `json:"abbreviatedProductName,omitempty" bson:"abbreviatedProductName,omitempty"`
	GenericName                                     string            `json:"genericName,omitempty" bson:"genericName,omitempty"`
	Quantity                                        string            `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Packaging                                       string            `json:"-" bson:"-"`
	PackagingTags                                   string            `json:"-" bson:"-"`
	PackagingEn                                     string            `json:"-" bson:"-"`
	PackagingText                                   string            `json:"-" bson:"-"`
	Brands                                          string            `json:"brands,omitempty" bson:"brands,omitempty"`
	BrandsTags                                      string            `json:"brandsTags,omitempty" bson:"brandsTags,omitempty"`
	Categories                                      string            `json:"categories,omitempty" bson:"categories,omitempty"`
	CategoriesTags                                  string            `json:"categoriesTags,omitempty" bson:"categoriesTags,omitempty"`
	CategoriesEn                                    string            `json:"categoriesEn,omitempty" bson:"categoriesEn,omitempty"`
	Origins                                         string            `json:"origins,omitempty" bson:"origins,omitempty"`
	OriginsTags                                     string            `json:"originsTags,omitempty" bson:"originsTags,omitempty"`
	OriginsEn                                       string            `json:"originsEn,omitempty" bson:"originsEn,omitempty"`
	ManufacturingPlaces                             string            `json:"manufacturingPlaces,omitempty" bson:"manufacturingPlaces,omitempty"`
	ManufacturingPlacesTags                         string            `json:"manufacturingPlacesTags,omitempty" bson:"manufacturingPlacesTags,omitempty"`
	Labels                                          string            `json:"labels,omitempty" bson:"labels,omitempty"`
	LabelsTags                                      string            `json:"labelsTags,omitempty" bson:"labelsTags,omitempty"`
	LabelsEn                                        string            `json:"labelsEn,omitempty" bson:"labelsEn,omitempty"`
	EmbCodes                                        string            `json:"embCodes,omitempty" bson:"embCodes,omitempty"`
	EmbCodesTags                                    string            `json:"embCodesTags,omitempty" bson:"embCodesTags,omitempty"`
	FirstPackagingCodeGeo                           string            `json:"firstPackagingCodeGeo,omitempty" bson:"firstPackagingCodeGeo,omitempty"`
	Cities                                          string            `json:"cities,omitempty" bson:"cities,omitempty"`
	CitiesTags                                      string            `json:"citiesTags,omitempty" bson:"citiesTags,omitempty"`
	PurchasePlaces                                  string            `json:"purchasePlaces,omitempty" bson:"purchasePlaces,omitempty"`
	Stores                                          string            `json:"stores,omitempty" bson:"stores,omitempty"`
	Countries                                       string            `json:"countries,omitempty" bson:"countries,omitempty"`
	CountriesTags                                   string            `json:"countriesTags,omitempty" bson:"countriesTags,omitempty"`
	CountriesEn                                     string            `json:"countriesEn,omitempty" bson:"countriesEn,omitempty"`
	IngredientsText                                 string            `json:"ingredientsText,omitempty" bson:"ingredientsText,omitempty"`
	IngredientsTags                                 string            `json:"ingredientsTags,omitempty" bson:"ingredientsTags,omitempty"`
	IngredientsAnalysisTags                         string            `json:"ingredientsAnalysisTags,omitempty" bson:"ingredientsAnalysisTags,omitempty"`
	Allergens                                       string            `json:"allergens,omitempty" bson:"allergens,omitempty"`
	AllergensEn                                     string            `json:"allergensEn,omitempty" bson:"allergensEn,omitempty"`
	Traces                                          string            `json:"traces,omitempty" bson:"traces,omitempty"`
	TracesTags                                      string            `json:"tracesTags,omitempty" bson:"tracesTags,omitempty"`
	TracesEn                                        string            `json:"tracesEn,omitempty" bson:"tracesEn,omitempty"`
	ServingSize                                     string            `json:"servingSize,omitempty" bson:"servingSize,omitempty"`
	ServingQuantity                                 int               `json:"servingQuantity,omitempty" bson:"servingQuantity,omitempty"`
	NoNutritionData                                 string            `json:"-" bson:"-"`
	AdditivesN                                      int               `json:"additivesN,omitempty" bson:"additivesN,omitempty"`
	Additives                                       string            `json:"additives,omitempty" bson:"additives,omitempty"`
	AdditivesTags                                   string            `json:"additivesTags,omitempty" bson:"additivesTags,omitempty"`
	AdditivesEn                                     string            `json:"additivesEn,omitempty" bson:"additivesEn,omitempty"`
	NutriscoreScore                                 string            `json:"nutriscoreScore,omitempty" bson:"nutriscoreScore,omitempty"`
	NutriscoreGrade                                 string            `json:"nutriscoreGrade,omitempty" bson:"nutriscoreGrade,omitempty"`
	NovaGroup                                       string            `json:"novaGroup,omitempty" bson:"novaGroup,omitempty"`
	PnnsGroups1                                     string            `json:"pnnsGroups1,omitempty" bson:"pnnsGroups1,omitempty"`
	PnnsGroups2                                     string            `json:"pnnsGroups2,omitempty" bson:"pnnsGroups2,omitempty"`
	FoodGroups                                      string            `json:"foodGroups,omitempty" bson:"foodGroups,omitempty"`
	FoodGroupsTags                                  string            `json:"foodGroupsTags,omitempty" bson:"foodGroupsTags,omitempty"`
	FoodGroupsEn                                    string            `json:"foodGroupsEn,omitempty" bson:"foodGroupsEn,omitempty"`
	States                                          string            `json:"-" bson:"-"`
	StatesTags                                      string            `json:"statesTags,omitempty" bson:"statesTags,omitempty"`
	StatesEn                                        string            `json:"statesEn,omitempty" bson:"statesEn,omitempty"`
	BrandOwner                                      string            `json:"brandOwner,omitempty" bson:"brandOwner,omitempty"`
	EcoscoreScore                                   string            `json:"ecoscoreScore,omitempty" bson:"ecoscoreScore,omitempty"`
	EcoscoreGrade                                   string            `json:"ecoscoreGrade,omitempty" bson:"ecoscoreGrade,omitempty"`
	NutrientLevelsTags                              string            `json:"nutrientLevelsTags,omitempty" bson:"nutrientLevelsTags,omitempty"`
	ProductQuantity                                 int               `json:"productQuantity,omitempty" bson:"productQuantity,omitempty"`
	Owner                                           string            `json:"owner,omitempty" bson:"owner,omitempty"`
	DataQualityErrorsTags                           string            `json:"dataQualityErrorsTags,omitempty" bson:"dataQualityErrorsTags,omitempty"`
	UniqueScansN                                    int               `json:"uniqueScansN,omitempty" bson:"uniqueScansN,omitempty"`
	PopularityTags                                  string            `json:"-" bson:"-"`
	Completeness                                    float64           `json:"completeness,omitempty" bson:"completeness,omitempty"`
	LastImageT                                      int               `json:"lastImageT,omitempty" bson:"lastImageT,omitempty"`
	LastImageDatetime                               types.TimeISO8601 `json:"lastImageDatetime,omitempty" bson:"lastImageDatetime,omitempty"`
	MainCategory                                    string            `json:"mainCategory,omitempty" bson:"mainCategory,omitempty"`
	MainCategoryEn                                  string            `json:"mainCategoryEn,omitempty" bson:"mainCategoryEn,omitempty"`
	ImageURL                                        string            `json:"imageURL,omitempty" bson:"imageURL,omitempty"`
	ImageSmallURL                                   string            `json:"imageSmallURL,omitempty" bson:"imageSmallURL,omitempty"`
	ImageIngredientsURL                             string            `json:"imageIngredientsURL,omitempty" bson:"imageIngredientsURL,omitempty"`
	ImageIngredientsSmallURL                        string            `json:"imageIngredientsSmallURL,omitempty" bson:"imageIngredientsSmallURL,omitempty"`
	ImageNutritionURL                               string            `json:"imageNutritionURL,omitempty" bson:"imageNutritionURL,omitempty"`
	ImageNutritionSmallURL                          string            `json:"imageNutritionSmallURL,omitempty" bson:"imageNutritionSmallURL,omitempty"`
	EnergyKj100G                                    int               `json:"energyKj100G,omitempty" bson:"energyKj100G,omitempty"`
	EnergyKcal100G                                  int               `json:"energyKcal100G,omitempty" bson:"energyKcal100G,omitempty"`
	Energy100G                                      int               `json:"energy100G,omitempty" bson:"energy100G,omitempty"`
	EnergyFromFat100G                               int               `json:"energyFromFat100G,omitempty" bson:"energyFromFat100G,omitempty"`
	Fat100G                                         int               `json:"fat100G,omitempty" bson:"fat100G,omitempty"`
	SaturatedFat100G                                int               `json:"saturatedFat100G,omitempty" bson:"saturatedFat100G,omitempty"`
	ButyricAcid100G                                 int               `json:"butyricAcid100G,omitempty" bson:"butyricAcid100G,omitempty"`
	CaproicAcid100G                                 int               `json:"caproicAcid100G,omitempty" bson:"caproicAcid100G,omitempty"`
	CaprylicAcid100G                                int               `json:"caprylicAcid100G,omitempty" bson:"caprylicAcid100G,omitempty"`
	CapricAcid100G                                  int               `json:"capricAcid100G,omitempty" bson:"capricAcid100G,omitempty"`
	LauricAcid100G                                  int               `json:"lauricAcid100G,omitempty" bson:"lauricAcid100G,omitempty"`
	MyristicAcid100G                                int               `json:"myristicAcid100G,omitempty" bson:"myristicAcid100G,omitempty"`
	PalmiticAcid100G                                int               `json:"palmiticAcid100G,omitempty" bson:"palmiticAcid100G,omitempty"`
	StearicAcid100G                                 int               `json:"stearicAcid100G,omitempty" bson:"stearicAcid100G,omitempty"`
	ArachidicAcid100G                               int               `json:"arachidicAcid100G,omitempty" bson:"arachidicAcid100G,omitempty"`
	BehenicAcid100G                                 int               `json:"behenicAcid100G,omitempty" bson:"behenicAcid100G,omitempty"`
	LignocericAcid100G                              int               `json:"lignocericAcid100G,omitempty" bson:"lignocericAcid100G,omitempty"`
	CeroticAcid100G                                 int               `json:"ceroticAcid100G,omitempty" bson:"ceroticAcid100G,omitempty"`
	MontanicAcid100G                                int               `json:"montanicAcid100G,omitempty" bson:"montanicAcid100G,omitempty"`
	MelissicAcid100G                                int               `json:"melissicAcid100G,omitempty" bson:"melissicAcid100G,omitempty"`
	UnsaturatedFat100G                              int               `json:"unsaturatedFat100G,omitempty" bson:"unsaturatedFat100G,omitempty"`
	MonounsaturatedFat100G                          int               `json:"monounsaturatedFat100G,omitempty" bson:"monounsaturatedFat100G,omitempty"`
	Omega9Fat100G                                   int               `json:"omega9Fat100G,omitempty" bson:"omega9Fat100G,omitempty"`
	PolyunsaturatedFat100G                          int               `json:"polyunsaturatedFat100G,omitempty" bson:"polyunsaturatedFat100G,omitempty"`
	Omega3Fat100G                                   int               `json:"omega3Fat100G,omitempty" bson:"omega3Fat100G,omitempty"`
	Omega6Fat100G                                   int               `json:"omega6Fat100G,omitempty" bson:"omega6Fat100G,omitempty"`
	AlphaLinolenicAcid100G                          int               `json:"alphaLinolenicAcid100G,omitempty" bson:"alphaLinolenicAcid100G,omitempty"`
	EicosapentaenoicAcid100G                        int               `json:"eicosapentaenoicAcid100G,omitempty" bson:"eicosapentaenoicAcid100G,omitempty"`
	DocosahexaenoicAcid100G                         int               `json:"docosahexaenoicAcid100G,omitempty" bson:"docosahexaenoicAcid100G,omitempty"`
	LinoleicAcid100G                                int               `json:"linoleicAcid100G,omitempty" bson:"linoleicAcid100G,omitempty"`
	ArachidonicAcid100G                             int               `json:"arachidonicAcid100G,omitempty" bson:"arachidonicAcid100G,omitempty"`
	GammaLinolenicAcid100G                          int               `json:"gammaLinolenicAcid100G,omitempty" bson:"gammaLinolenicAcid100G,omitempty"`
	DihomoGammaLinolenicAcid100G                    int               `json:"dihomoGammaLinolenicAcid100G,omitempty" bson:"dihomoGammaLinolenicAcid100G,omitempty"`
	OleicAcid100G                                   int               `json:"oleicAcid100G,omitempty" bson:"oleicAcid100G,omitempty"`
	ElaidicAcid100G                                 int               `json:"elaidicAcid100G,omitempty" bson:"elaidicAcid100G,omitempty"`
	GondoicAcid100G                                 int               `json:"gondoicAcid100G,omitempty" bson:"gondoicAcid100G,omitempty"`
	MeadAcid100G                                    int               `json:"meadAcid100G,omitempty" bson:"meadAcid100G,omitempty"`
	ErucicAcid100G                                  int               `json:"erucicAcid100G,omitempty" bson:"erucicAcid100G,omitempty"`
	NervonicAcid100G                                int               `json:"nervonicAcid100G,omitempty" bson:"nervonicAcid100G,omitempty"`
	TransFat100G                                    int               `json:"transFat100G,omitempty" bson:"transFat100G,omitempty"`
	Cholesterol100G                                 int               `json:"cholesterol100G,omitempty" bson:"cholesterol100G,omitempty"`
	Carbohydrates100G                               int               `json:"carbohydrates100G,omitempty" bson:"carbohydrates100G,omitempty"`
	Sugars100G                                      int               `json:"sugars100G,omitempty" bson:"sugars100G,omitempty"`
	AddedSugars100G                                 int               `json:"addedSugars100G,omitempty" bson:"addedSugars100G,omitempty"`
	Sucrose100G                                     int               `json:"sucrose100G,omitempty" bson:"sucrose100G,omitempty"`
	Glucose100G                                     int               `json:"glucose100G,omitempty" bson:"glucose100G,omitempty"`
	Fructose100G                                    int               `json:"fructose100G,omitempty" bson:"fructose100G,omitempty"`
	Lactose100G                                     int               `json:"lactose100G,omitempty" bson:"lactose100G,omitempty"`
	Maltose100G                                     int               `json:"maltose100G,omitempty" bson:"maltose100G,omitempty"`
	Maltodextrins100G                               int               `json:"maltodextrins100G,omitempty" bson:"maltodextrins100G,omitempty"`
	Starch100G                                      int               `json:"starch100G,omitempty" bson:"starch100G,omitempty"`
	Polyols100G                                     int               `json:"polyols100G,omitempty" bson:"polyols100G,omitempty"`
	Erythritol100G                                  int               `json:"erythritol100G,omitempty" bson:"erythritol100G,omitempty"`
	Fiber100G                                       int               `json:"fiber100G,omitempty" bson:"fiber100G,omitempty"`
	SolubleFiber100G                                int               `json:"solubleFiber100G,omitempty" bson:"solubleFiber100G,omitempty"`
	InsolubleFiber100G                              int               `json:"insolubleFiber100G,omitempty" bson:"insolubleFiber100G,omitempty"`
	Proteins100G                                    int               `json:"proteins100G,omitempty" bson:"proteins100G,omitempty"`
	Casein100G                                      int               `json:"casein100G,omitempty" bson:"casein100G,omitempty"`
	SerumProteins100G                               int               `json:"serumProteins100G,omitempty" bson:"serumProteins100G,omitempty"`
	Nucleotides100G                                 int               `json:"nucleotides100G,omitempty" bson:"nucleotides100G,omitempty"`
	Salt100G                                        int               `json:"salt100G,omitempty" bson:"salt100G,omitempty"`
	AddedSalt100G                                   int               `json:"addedSalt100G,omitempty" bson:"addedSalt100G,omitempty"`
	Sodium100G                                      int               `json:"sodium100G,omitempty" bson:"sodium100G,omitempty"`
	Alcohol100G                                     int               `json:"alcohol100G,omitempty" bson:"alcohol100G,omitempty"`
	VitaminA100G                                    int               `json:"vitaminA100G,omitempty" bson:"vitaminA100G,omitempty"`
	BetaCarotene100G                                int               `json:"betaCarotene100G,omitempty" bson:"betaCarotene100G,omitempty"`
	VitaminD100G                                    int               `json:"vitaminD100G,omitempty" bson:"vitaminD100G,omitempty"`
	VitaminE100G                                    int               `json:"vitaminE100G,omitempty" bson:"vitaminE100G,omitempty"`
	VitaminK100G                                    int               `json:"vitaminK100G,omitempty" bson:"vitaminK100G,omitempty"`
	VitaminC100G                                    int               `json:"vitaminC100G,omitempty" bson:"vitaminC100G,omitempty"`
	VitaminB1100G                                   int               `json:"vitaminB1100G,omitempty" bson:"vitaminB1100G,omitempty"`
	VitaminB2100G                                   int               `json:"vitaminB2100G,omitempty" bson:"vitaminB2100G,omitempty"`
	VitaminPp100G                                   int               `json:"vitaminPp100G,omitempty" bson:"vitaminPp100G,omitempty"`
	VitaminB6100G                                   int               `json:"vitaminB6100G,omitempty" bson:"vitaminB6100G,omitempty"`
	VitaminB9100G                                   int               `json:"vitaminB9100G,omitempty" bson:"vitaminB9100G,omitempty"`
	Folates100G                                     int               `json:"folates100G,omitempty" bson:"folates100G,omitempty"`
	VitaminB12100G                                  int               `json:"vitaminB12100G,omitempty" bson:"vitaminB12100G,omitempty"`
	Biotin100G                                      int               `json:"biotin100G,omitempty" bson:"biotin100G,omitempty"`
	PantothenicAcid100G                             int               `json:"pantothenicAcid100G,omitempty" bson:"pantothenicAcid100G,omitempty"`
	Silica100G                                      int               `json:"silica100G,omitempty" bson:"silica100G,omitempty"`
	Bicarbonate100G                                 int               `json:"bicarbonate100G,omitempty" bson:"bicarbonate100G,omitempty"`
	Potassium100G                                   int               `json:"potassium100G,omitempty" bson:"potassium100G,omitempty"`
	Chloride100G                                    int               `json:"chloride100G,omitempty" bson:"chloride100G,omitempty"`
	Calcium100G                                     int               `json:"calcium100G,omitempty" bson:"calcium100G,omitempty"`
	Phosphorus100G                                  int               `json:"phosphorus100G,omitempty" bson:"phosphorus100G,omitempty"`
	Iron100G                                        int               `json:"iron100G,omitempty" bson:"iron100G,omitempty"`
	Magnesium100G                                   int               `json:"magnesium100G,omitempty" bson:"magnesium100G,omitempty"`
	Zinc100G                                        int               `json:"zinc100G,omitempty" bson:"zinc100G,omitempty"`
	Copper100G                                      int               `json:"copper100G,omitempty" bson:"copper100G,omitempty"`
	Manganese100G                                   int               `json:"manganese100G,omitempty" bson:"manganese100G,omitempty"`
	Fluoride100G                                    int               `json:"fluoride100G,omitempty" bson:"fluoride100G,omitempty"`
	Selenium100G                                    int               `json:"selenium100G,omitempty" bson:"selenium100G,omitempty"`
	Chromium100G                                    int               `json:"chromium100G,omitempty" bson:"chromium100G,omitempty"`
	Molybdenum100G                                  int               `json:"molybdenum100G,omitempty" bson:"molybdenum100G,omitempty"`
	Iodine100G                                      int               `json:"iodine100G,omitempty" bson:"iodine100G,omitempty"`
	Caffeine100G                                    int               `json:"caffeine100G,omitempty" bson:"caffeine100G,omitempty"`
	Taurine100G                                     int               `json:"taurine100G,omitempty" bson:"taurine100G,omitempty"`
	Ph100G                                          int               `json:"ph100G,omitempty" bson:"ph100G,omitempty"`
	FruitsVegetablesNuts100G                        int               `json:"fruitsVegetablesNuts100G,omitempty" bson:"fruitsVegetablesNuts100G,omitempty"`
	FruitsVegetablesNutsDried100G                   int               `json:"fruitsVegetablesNutsDried100G,omitempty" bson:"fruitsVegetablesNutsDried100G,omitempty"`
	FruitsVegetablesNutsEstimate100G                int               `json:"fruitsVegetablesNutsEstimate100G,omitempty" bson:"fruitsVegetablesNutsEstimate100G,omitempty"`
	FruitsVegetablesNutsEstimateFromIngredients100G int               `json:"fruitsVegetablesNutsEstimateFromIngredients100G,omitempty" bson:"fruitsVegetablesNutsEstimateFromIngredients100G,omitempty"`
	CollagenMeatProteinRatio100G                    int               `json:"collagenMeatProteinRatio100G,omitempty" bson:"collagenMeatProteinRatio100G,omitempty"`
	Cocoa100G                                       int               `json:"cocoa100G,omitempty" bson:"cocoa100G,omitempty"`
	Chlorophyl100G                                  int               `json:"chlorophyl100G,omitempty" bson:"chlorophyl100G,omitempty"`
	CarbonFootprint100G                             int               `json:"carbonFootprint100G,omitempty" bson:"carbonFootprint100G,omitempty"`
	CarbonFootprintFromMeatOrFish100G               int               `json:"carbonFootprintFromMeatOrFish100G,omitempty" bson:"carbonFootprintFromMeatOrFish100G,omitempty"`
	NutritionScoreFr100G                            int               `json:"nutritionScoreFr100G,omitempty" bson:"nutritionScoreFr100G,omitempty"`
	NutritionScoreUk100G                            int               `json:"nutritionScoreUk100G,omitempty" bson:"nutritionScoreUk100G,omitempty"`
	GlycemicIndex100G                               int               `json:"glycemicIndex100G,omitempty" bson:"glycemicIndex100G,omitempty"`
	WaterHardness100G                               int               `json:"waterHardness100G,omitempty" bson:"waterHardness100G,omitempty"`
	Choline100G                                     int               `json:"choline100G,omitempty" bson:"choline100G,omitempty"`
	Phylloquinone100G                               int               `json:"phylloquinone100G,omitempty" bson:"phylloquinone100G,omitempty"`
	BetaGlucan100G                                  int               `json:"betaGlucan100G,omitempty" bson:"betaGlucan100G,omitempty"`
	Inositol100G                                    int               `json:"inositol100G,omitempty" bson:"inositol100G,omitempty"`
	Carnitine100G                                   int               `json:"carnitine100G,omitempty" bson:"carnitine100G,omitempty"`
	Sulphate100G                                    int               `json:"sulphate100G,omitempty" bson:"sulphate100G,omitempty"`
	Nitrate100G                                     int               `json:"nitrate100G,omitempty" bson:"nitrate100G,omitempty"`
	Acidity100G                                     int               `json:"acidity100G,omitempty" bson:"acidity100G,omitempty"`
}

func (off OpenFoodFact) GetID() string {
	return off.Code
}

func (off *OpenFoodFact) SetID(id string) {
	off.Code = id
}

func (off OpenFoodFactImg) GetID() string {
	return off.ID
}

func (off *OpenFoodFactImg) SetID(id string) {
	off.ID = id
}
