package openfoodfacts

type OpenFoodFactCSVEntry struct {
	Code                                            string `json:"code" bson:"code"`
	URL                                             string `json:"url" bson:"url"`
	Creator                                         string `json:"creator" bson:"creator"`
	CreatedT                                        string `json:"created_t" bson:"created_t"`
	CreatedDatetime                                 string `json:"created_datetime" bson:"created_datetime"`
	LastModifiedT                                   string `json:"last_modified_t" bson:"last_modified_t"`
	LastModifiedDatetime                            string `json:"last_modified_datetime" bson:"last_modified_datetime"`
	LastModifiedBy                                  string `json:"last_modified_by" bson:"last_modified_by"`
	LastUpdatedT                                    string `json:"last_updated_t" bson:"last_updated_t"`
	LastUpdatedDatetime                             string `json:"last_updated_datetime" bson:"last_updated_datetime"`
	ProductName                                     string `json:"product_name" bson:"product_name"`
	AbbreviatedProductName                          string `json:"abbreviated_product_name" bson:"abbreviated_product_name"`
	GenericName                                     string `json:"generic_name" bson:"generic_name"`
	Quantity                                        string `json:"quantity" bson:"quantity"`
	Packaging                                       string `json:"packaging" bson:"packaging"`
	PackagingTags                                   string `json:"packaging_tags" bson:"packaging_tags"`
	PackagingEn                                     string `json:"packaging_en" bson:"packaging_en"`
	PackagingText                                   string `json:"packaging_text" bson:"packaging_text"`
	Brands                                          string `json:"brands" bson:"brands"`
	BrandsTags                                      string `json:"brands_tags" bson:"brands_tags"`
	Categories                                      string `json:"categories" bson:"categories"`
	CategoriesTags                                  string `json:"categories_tags" bson:"categories_tags"`
	CategoriesEn                                    string `json:"categories_en" bson:"categories_en"`
	Origins                                         string `json:"origins" bson:"origins"`
	OriginsTags                                     string `json:"origins_tags" bson:"origins_tags"`
	OriginsEn                                       string `json:"origins_en" bson:"origins_en"`
	ManufacturingPlaces                             string `json:"manufacturing_places" bson:"manufacturing_places"`
	ManufacturingPlacesTags                         string `json:"manufacturing_places_tags" bson:"manufacturing_places_tags"`
	Labels                                          string `json:"labels" bson:"labels"`
	LabelsTags                                      string `json:"labels_tags" bson:"labels_tags"`
	LabelsEn                                        string `json:"labels_en" bson:"labels_en"`
	EmbCodes                                        string `json:"emb_codes" bson:"emb_codes"`
	EmbCodesTags                                    string `json:"emb_codes_tags" bson:"emb_codes_tags"`
	FirstPackagingCodeGeo                           string `json:"first_packaging_code_geo" bson:"first_packaging_code_geo"`
	Cities                                          string `json:"cities" bson:"cities"`
	CitiesTags                                      string `json:"cities_tags" bson:"cities_tags"`
	PurchasePlaces                                  string `json:"purchase_places" bson:"purchase_places"`
	Stores                                          string `json:"stores" bson:"stores"`
	Countries                                       string `json:"countries" bson:"countries"`
	CountriesTags                                   string `json:"countries_tags" bson:"countries_tags"`
	CountriesEn                                     string `json:"countries_en" bson:"countries_en"`
	IngredientsText                                 string `json:"ingredients_text" bson:"ingredients_text"`
	IngredientsTags                                 string `json:"ingredients_tags" bson:"ingredients_tags"`
	IngredientsAnalysisTags                         string `json:"ingredients_analysis_tags" bson:"ingredients_analysis_tags"`
	Allergens                                       string `json:"allergens" bson:"allergens"`
	AllergensEn                                     string `json:"allergens_en" bson:"allergens_en"`
	Traces                                          string `json:"traces" bson:"traces"`
	TracesTags                                      string `json:"traces_tags" bson:"traces_tags"`
	TracesEn                                        string `json:"traces_en" bson:"traces_en"`
	ServingSize                                     string `json:"serving_size" bson:"serving_size"`
	ServingQuantity                                 string `json:"serving_quantity" bson:"serving_quantity"`
	NoNutritionData                                 string `json:"no_nutrition_data" bson:"no_nutrition_data"`
	AdditivesN                                      string `json:"additives_n" bson:"additives_n"`
	Additives                                       string `json:"additives" bson:"additives"`
	AdditivesTags                                   string `json:"additives_tags" bson:"additives_tags"`
	AdditivesEn                                     string `json:"additives_en" bson:"additives_en"`
	NutriscoreScore                                 string `json:"nutriscore_score" bson:"nutriscore_score"`
	NutriscoreGrade                                 string `json:"nutriscore_grade" bson:"nutriscore_grade"`
	NovaGroup                                       string `json:"nova_group" bson:"nova_group"`
	PnnsGroups1                                     string `json:"pnns_groups_1" bson:"pnns_groups_1"`
	PnnsGroups2                                     string `json:"pnns_groups_2" bson:"pnns_groups_2"`
	FoodGroups                                      string `json:"food_groups" bson:"food_groups"`
	FoodGroupsTags                                  string `json:"food_groups_tags" bson:"food_groups_tags"`
	FoodGroupsEn                                    string `json:"food_groups_en" bson:"food_groups_en"`
	States                                          string `json:"states" bson:"states"`
	StatesTags                                      string `json:"states_tags" bson:"states_tags"`
	StatesEn                                        string `json:"states_en" bson:"states_en"`
	BrandOwner                                      string `json:"brand_owner" bson:"brand_owner"`
	EcoscoreScore                                   string `json:"ecoscore_score" bson:"ecoscore_score"`
	EcoscoreGrade                                   string `json:"ecoscore_grade" bson:"ecoscore_grade"`
	NutrientLevelsTags                              string `json:"nutrient_levels_tags" bson:"nutrient_levels_tags"`
	ProductQuantity                                 string `json:"product_quantity" bson:"product_quantity"`
	Owner                                           string `json:"owner" bson:"owner"`
	DataQualityErrorsTags                           string `json:"data_quality_errors_tags" bson:"data_quality_errors_tags"`
	UniqueScansN                                    string `json:"unique_scans_n" bson:"unique_scans_n"`
	PopularityTags                                  string `json:"popularity_tags" bson:"popularity_tags"`
	Completeness                                    string `json:"completeness" bson:"completeness"`
	LastImageT                                      string `json:"last_image_t" bson:"last_image_t"`
	LastImageDatetime                               string `json:"last_image_datetime" bson:"last_image_datetime"`
	MainCategory                                    string `json:"main_category" bson:"main_category"`
	MainCategoryEn                                  string `json:"main_category_en" bson:"main_category_en"`
	ImageURL                                        string `json:"image_url" bson:"image_url"`
	ImageSmallURL                                   string `json:"image_small_url" bson:"image_small_url"`
	ImageIngredientsURL                             string `json:"image_ingredients_url" bson:"image_ingredients_url"`
	ImageIngredientsSmallURL                        string `json:"image_ingredients_small_url" bson:"image_ingredients_small_url"`
	ImageNutritionURL                               string `json:"image_nutrition_url" bson:"image_nutrition_url"`
	ImageNutritionSmallURL                          string `json:"image_nutrition_small_url" bson:"image_nutrition_small_url"`
	Energykj100g                                    string `json:"energykj_100g" bson:"energykj_100g"`
	Energykcal100g                                  string `json:"energykcal_100g" bson:"energykcal_100g"`
	Energy100g                                      string `json:"energy_100g" bson:"energy_100g"`
	Energyfromfat100g                               string `json:"energyfromfat_100g" bson:"energyfromfat_100g"`
	Fat100g                                         string `json:"fat_100g" bson:"fat_100g"`
	Saturatedfat100g                                string `json:"saturatedfat_100g" bson:"saturatedfat_100g"`
	Butyricacid100g                                 string `json:"butyricacid_100g" bson:"butyricacid_100g"`
	Caproicacid100g                                 string `json:"caproicacid_100g" bson:"caproicacid_100g"`
	Caprylicacid100g                                string `json:"caprylicacid_100g" bson:"caprylicacid_100g"`
	Capricacid100g                                  string `json:"capricacid_100g" bson:"capricacid_100g"`
	Lauricacid100g                                  string `json:"lauricacid_100g" bson:"lauricacid_100g"`
	Myristicacid100g                                string `json:"myristicacid_100g" bson:"myristicacid_100g"`
	Palmiticacid100g                                string `json:"palmiticacid_100g" bson:"palmiticacid_100g"`
	Stearicacid100g                                 string `json:"stearicacid_100g" bson:"stearicacid_100g"`
	Arachidicacid100g                               string `json:"arachidicacid_100g" bson:"arachidicacid_100g"`
	Behenicacid100g                                 string `json:"behenicacid_100g" bson:"behenicacid_100g"`
	Lignocericacid100g                              string `json:"lignocericacid_100g" bson:"lignocericacid_100g"`
	Ceroticacid100g                                 string `json:"ceroticacid_100g" bson:"ceroticacid_100g"`
	Montanicacid100g                                string `json:"montanicacid_100g" bson:"montanicacid_100g"`
	Melissicacid100g                                string `json:"melissicacid_100g" bson:"melissicacid_100g"`
	Unsaturatedfat100g                              string `json:"unsaturatedfat_100g" bson:"unsaturatedfat_100g"`
	Monounsaturatedfat100g                          string `json:"monounsaturatedfat_100g" bson:"monounsaturatedfat_100g"`
	Omega9fat100g                                   string `json:"omega9fat_100g" bson:"omega9fat_100g"`
	Polyunsaturatedfat100g                          string `json:"polyunsaturatedfat_100g" bson:"polyunsaturatedfat_100g"`
	Omega3fat100g                                   string `json:"omega3fat_100g" bson:"omega3fat_100g"`
	Omega6fat100g                                   string `json:"omega6fat_100g" bson:"omega6fat_100g"`
	Alphalinolenicacid100g                          string `json:"alphalinolenicacid_100g" bson:"alphalinolenicacid_100g"`
	Eicosapentaenoicacid100g                        string `json:"eicosapentaenoicacid_100g" bson:"eicosapentaenoicacid_100g"`
	Docosahexaenoicacid100g                         string `json:"docosahexaenoicacid_100g" bson:"docosahexaenoicacid_100g"`
	Linoleicacid100g                                string `json:"linoleicacid_100g" bson:"linoleicacid_100g"`
	Arachidonicacid100g                             string `json:"arachidonicacid_100g" bson:"arachidonicacid_100g"`
	Gammalinolenicacid100g                          string `json:"gammalinolenicacid_100g" bson:"gammalinolenicacid_100g"`
	Dihomogammalinolenicacid100g                    string `json:"dihomogammalinolenicacid_100g" bson:"dihomogammalinolenicacid_100g"`
	Oleicacid100g                                   string `json:"oleicacid_100g" bson:"oleicacid_100g"`
	Elaidicacid100g                                 string `json:"elaidicacid_100g" bson:"elaidicacid_100g"`
	Gondoicacid100g                                 string `json:"gondoicacid_100g" bson:"gondoicacid_100g"`
	Meadacid100g                                    string `json:"meadacid_100g" bson:"meadacid_100g"`
	Erucicacid100g                                  string `json:"erucicacid_100g" bson:"erucicacid_100g"`
	Nervonicacid100g                                string `json:"nervonicacid_100g" bson:"nervonicacid_100g"`
	Transfat100g                                    string `json:"transfat_100g" bson:"transfat_100g"`
	Cholesterol100g                                 string `json:"cholesterol_100g" bson:"cholesterol_100g"`
	Carbohydrates100g                               string `json:"carbohydrates_100g" bson:"carbohydrates_100g"`
	Sugars100g                                      string `json:"sugars_100g" bson:"sugars_100g"`
	Addedsugars100g                                 string `json:"addedsugars_100g" bson:"addedsugars_100g"`
	Sucrose100g                                     string `json:"sucrose_100g" bson:"sucrose_100g"`
	Glucose100g                                     string `json:"glucose_100g" bson:"glucose_100g"`
	Fructose100g                                    string `json:"fructose_100g" bson:"fructose_100g"`
	Lactose100g                                     string `json:"lactose_100g" bson:"lactose_100g"`
	Maltose100g                                     string `json:"maltose_100g" bson:"maltose_100g"`
	Maltodextrins100g                               string `json:"maltodextrins_100g" bson:"maltodextrins_100g"`
	Starch100g                                      string `json:"starch_100g" bson:"starch_100g"`
	Polyols100g                                     string `json:"polyols_100g" bson:"polyols_100g"`
	Erythritol100g                                  string `json:"erythritol_100g" bson:"erythritol_100g"`
	Fiber100g                                       string `json:"fiber_100g" bson:"fiber_100g"`
	Solublefiber100g                                string `json:"solublefiber_100g" bson:"solublefiber_100g"`
	Insolublefiber100g                              string `json:"insolublefiber_100g" bson:"insolublefiber_100g"`
	Proteins100g                                    string `json:"proteins_100g" bson:"proteins_100g"`
	Casein100g                                      string `json:"casein_100g" bson:"casein_100g"`
	Serumproteins100g                               string `json:"serumproteins_100g" bson:"serumproteins_100g"`
	Nucleotides100g                                 string `json:"nucleotides_100g" bson:"nucleotides_100g"`
	Salt100g                                        string `json:"salt_100g" bson:"salt_100g"`
	Addedsalt100g                                   string `json:"addedsalt_100g" bson:"addedsalt_100g"`
	Sodium100g                                      string `json:"sodium_100g" bson:"sodium_100g"`
	Alcohol100g                                     string `json:"alcohol_100g" bson:"alcohol_100g"`
	Vitamina100g                                    string `json:"vitamina_100g" bson:"vitamina_100g"`
	Betacarotene100g                                string `json:"betacarotene_100g" bson:"betacarotene_100g"`
	Vitamind100g                                    string `json:"vitamind_100g" bson:"vitamind_100g"`
	Vitamine100g                                    string `json:"vitamine_100g" bson:"vitamine_100g"`
	Vitamink100g                                    string `json:"vitamink_100g" bson:"vitamink_100g"`
	Vitaminc100g                                    string `json:"vitaminc_100g" bson:"vitaminc_100g"`
	Vitaminb1100g                                   string `json:"vitaminb1_100g" bson:"vitaminb1_100g"`
	Vitaminb2100g                                   string `json:"vitaminb2_100g" bson:"vitaminb2_100g"`
	Vitaminpp100g                                   string `json:"vitaminpp_100g" bson:"vitaminpp_100g"`
	Vitaminb6100g                                   string `json:"vitaminb6_100g" bson:"vitaminb6_100g"`
	Vitaminb9100g                                   string `json:"vitaminb9_100g" bson:"vitaminb9_100g"`
	Folates100g                                     string `json:"folates_100g" bson:"folates_100g"`
	Vitaminb12100g                                  string `json:"vitaminb12_100g" bson:"vitaminb12_100g"`
	Biotin100g                                      string `json:"biotin_100g" bson:"biotin_100g"`
	Pantothenicacid100g                             string `json:"pantothenicacid_100g" bson:"pantothenicacid_100g"`
	Silica100g                                      string `json:"silica_100g" bson:"silica_100g"`
	Bicarbonate100g                                 string `json:"bicarbonate_100g" bson:"bicarbonate_100g"`
	Potassium100g                                   string `json:"potassium_100g" bson:"potassium_100g"`
	Chloride100g                                    string `json:"chloride_100g" bson:"chloride_100g"`
	Calcium100g                                     string `json:"calcium_100g" bson:"calcium_100g"`
	Phosphorus100g                                  string `json:"phosphorus_100g" bson:"phosphorus_100g"`
	Iron100g                                        string `json:"iron_100g" bson:"iron_100g"`
	Magnesium100g                                   string `json:"magnesium_100g" bson:"magnesium_100g"`
	Zinc100g                                        string `json:"zinc_100g" bson:"zinc_100g"`
	Copper100g                                      string `json:"copper_100g" bson:"copper_100g"`
	Manganese100g                                   string `json:"manganese_100g" bson:"manganese_100g"`
	Fluoride100g                                    string `json:"fluoride_100g" bson:"fluoride_100g"`
	Selenium100g                                    string `json:"selenium_100g" bson:"selenium_100g"`
	Chromium100g                                    string `json:"chromium_100g" bson:"chromium_100g"`
	Molybdenum100g                                  string `json:"molybdenum_100g" bson:"molybdenum_100g"`
	Iodine100g                                      string `json:"iodine_100g" bson:"iodine_100g"`
	Caffeine100g                                    string `json:"caffeine_100g" bson:"caffeine_100g"`
	Taurine100g                                     string `json:"taurine_100g" bson:"taurine_100g"`
	Ph100g                                          string `json:"ph_100g" bson:"ph_100g"`
	Fruitsvegetablesnuts100g                        string `json:"fruitsvegetablesnuts_100g" bson:"fruitsvegetablesnuts_100g"`
	Fruitsvegetablesnutsdried100g                   string `json:"fruitsvegetablesnutsdried_100g" bson:"fruitsvegetablesnutsdried_100g"`
	Fruitsvegetablesnutsestimate100g                string `json:"fruitsvegetablesnutsestimate_100g" bson:"fruitsvegetablesnutsestimate_100g"`
	Fruitsvegetablesnutsestimatefromingredients100g string `json:"fruitsvegetablesnutsestimatefromingredients_100g" bson:"fruitsvegetablesnutsestimatefromingredients_100g"`
	Collagenmeatproteinratio100g                    string `json:"collagenmeatproteinratio_100g" bson:"collagenmeatproteinratio_100g"`
	Cocoa100g                                       string `json:"cocoa_100g" bson:"cocoa_100g"`
	Chlorophyl100g                                  string `json:"chlorophyl_100g" bson:"chlorophyl_100g"`
	Carbonfootprint100g                             string `json:"carbonfootprint_100g" bson:"carbonfootprint_100g"`
	Carbonfootprintfrommeatorfish100g               string `json:"carbonfootprintfrommeatorfish_100g" bson:"carbonfootprintfrommeatorfish_100g"`
	Nutritionscorefr100g                            string `json:"nutritionscorefr_100g" bson:"nutritionscorefr_100g"`
	Nutritionscoreuk100g                            string `json:"nutritionscoreuk_100g" bson:"nutritionscoreuk_100g"`
	Glycemicindex100g                               string `json:"glycemicindex_100g" bson:"glycemicindex_100g"`
	Waterhardness100g                               string `json:"waterhardness_100g" bson:"waterhardness_100g"`
	Choline100g                                     string `json:"choline_100g" bson:"choline_100g"`
	Phylloquinone100g                               string `json:"phylloquinone_100g" bson:"phylloquinone_100g"`
	Betaglucan100g                                  string `json:"betaglucan_100g" bson:"betaglucan_100g"`
	Inositol100g                                    string `json:"inositol_100g" bson:"inositol_100g"`
	Carnitine100g                                   string `json:"carnitine_100g" bson:"carnitine_100g"`
	Sulphate100g                                    string `json:"sulphate_100g" bson:"sulphate_100g"`
	Nitrate100g                                     string `json:"nitrate_100g" bson:"nitrate_100g"`
	Acidity100g                                     string `json:"acidity_100g" bson:"acidity_100g"`

	// keep stuff that are not in the csv at the end
	ID string `bson:"_id" json:"_id"`
}

func (off OpenFoodFactCSVEntry) GetID() string {
	return off.ID
}

func (off *OpenFoodFactCSVEntry) SetID(id string) {
	off.ID = id
}
