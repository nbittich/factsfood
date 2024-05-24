package openfoodfacts

type OpenFoodFactCSVEntry struct {
	Code                                            string `json:"_id" bson:"_id"`
	URL                                             string `json:"url,omitempty" bson:"url,omitempty"`
	Creator                                         string `json:"creator,omitempty" bson:"creator,omitempty"`
	CreatedT                                        string `json:"created_t,omitempty" bson:"created_t,omitempty"`
	CreatedDatetime                                 string `json:"created_datetime,omitempty" bson:"created_datetime,omitempty"`
	LastModifiedT                                   string `json:"last_modified_t,omitempty" bson:"last_modified_t,omitempty"`
	LastModifiedDatetime                            string `json:"last_modified_datetime,omitempty" bson:"last_modified_datetime,omitempty"`
	LastModifiedBy                                  string `json:"last_modified_by,omitempty" bson:"last_modified_by,omitempty"`
	LastUpdatedT                                    string `json:"last_updated_t,omitempty" bson:"last_updated_t,omitempty"`
	LastUpdatedDatetime                             string `json:"last_updated_datetime,omitempty" bson:"last_updated_datetime,omitempty"`
	ProductName                                     string `json:"product_name,omitempty" bson:"product_name,omitempty"`
	AbbreviatedProductName                          string `json:"abbreviated_product_name,omitempty" bson:"abbreviated_product_name,omitempty"`
	GenericName                                     string `json:"generic_name,omitempty" bson:"generic_name,omitempty"`
	Quantity                                        string `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Packaging                                       string `json:"packaging,omitempty" bson:"packaging,omitempty"`
	PackagingTags                                   string `json:"packaging_tags,omitempty" bson:"packaging_tags,omitempty"`
	PackagingEn                                     string `json:"packaging_en,omitempty" bson:"packaging_en,omitempty"`
	PackagingText                                   string `json:"packaging_text,omitempty" bson:"packaging_text,omitempty"`
	Brands                                          string `json:"brands,omitempty" bson:"brands,omitempty"`
	BrandsTags                                      string `json:"brands_tags,omitempty" bson:"brands_tags,omitempty"`
	Categories                                      string `json:"categories,omitempty" bson:"categories,omitempty"`
	CategoriesTags                                  string `json:"categories_tags,omitempty" bson:"categories_tags,omitempty"`
	CategoriesEn                                    string `json:"categories_en,omitempty" bson:"categories_en,omitempty"`
	Origins                                         string `json:"origins,omitempty" bson:"origins,omitempty"`
	OriginsTags                                     string `json:"origins_tags,omitempty" bson:"origins_tags,omitempty"`
	OriginsEn                                       string `json:"origins_en,omitempty" bson:"origins_en,omitempty"`
	ManufacturingPlaces                             string `json:"manufacturing_places,omitempty" bson:"manufacturing_places,omitempty"`
	ManufacturingPlacesTags                         string `json:"manufacturing_places_tags,omitempty" bson:"manufacturing_places_tags,omitempty"`
	Labels                                          string `json:"labels,omitempty" bson:"labels,omitempty"`
	LabelsTags                                      string `json:"labels_tags,omitempty" bson:"labels_tags,omitempty"`
	LabelsEn                                        string `json:"labels_en,omitempty" bson:"labels_en,omitempty"`
	EmbCodes                                        string `json:"emb_codes,omitempty" bson:"emb_codes,omitempty"`
	EmbCodesTags                                    string `json:"emb_codes_tags,omitempty" bson:"emb_codes_tags,omitempty"`
	FirstPackagingCodeGeo                           string `json:"first_packaging_code_geo,omitempty" bson:"first_packaging_code_geo,omitempty"`
	Cities                                          string `json:"cities,omitempty" bson:"cities,omitempty"`
	CitiesTags                                      string `json:"cities_tags,omitempty" bson:"cities_tags,omitempty"`
	PurchasePlaces                                  string `json:"purchase_places,omitempty" bson:"purchase_places,omitempty"`
	Stores                                          string `json:"stores,omitempty" bson:"stores,omitempty"`
	Countries                                       string `json:"countries,omitempty" bson:"countries,omitempty"`
	CountriesTags                                   string `json:"countries_tags,omitempty" bson:"countries_tags,omitempty"`
	CountriesEn                                     string `json:"countries_en,omitempty" bson:"countries_en,omitempty"`
	IngredientsText                                 string `json:"ingredients_text,omitempty" bson:"ingredients_text,omitempty"`
	IngredientsTags                                 string `json:"ingredients_tags,omitempty" bson:"ingredients_tags,omitempty"`
	IngredientsAnalysisTags                         string `json:"ingredients_analysis_tags,omitempty" bson:"ingredients_analysis_tags,omitempty"`
	Allergens                                       string `json:"allergens,omitempty" bson:"allergens,omitempty"`
	AllergensEn                                     string `json:"allergens_en,omitempty" bson:"allergens_en,omitempty"`
	Traces                                          string `json:"traces,omitempty" bson:"traces,omitempty"`
	TracesTags                                      string `json:"traces_tags,omitempty" bson:"traces_tags,omitempty"`
	TracesEn                                        string `json:"traces_en,omitempty" bson:"traces_en,omitempty"`
	ServingSize                                     string `json:"serving_size,omitempty" bson:"serving_size,omitempty"`
	ServingQuantity                                 string `json:"serving_quantity,omitempty" bson:"serving_quantity,omitempty"`
	NoNutritionData                                 string `json:"no_nutrition_data,omitempty" bson:"no_nutrition_data,omitempty"`
	AdditivesN                                      string `json:"additives_n,omitempty" bson:"additives_n,omitempty"`
	Additives                                       string `json:"additives,omitempty" bson:"additives,omitempty"`
	AdditivesTags                                   string `json:"additives_tags,omitempty" bson:"additives_tags,omitempty"`
	AdditivesEn                                     string `json:"additives_en,omitempty" bson:"additives_en,omitempty"`
	NutriscoreScore                                 string `json:"nutriscore_score,omitempty" bson:"nutriscore_score,omitempty"`
	NutriscoreGrade                                 string `json:"nutriscore_grade,omitempty" bson:"nutriscore_grade,omitempty"`
	NovaGroup                                       string `json:"nova_group,omitempty" bson:"nova_group,omitempty"`
	PnnsGroups1                                     string `json:"pnns_groups_1,omitempty" bson:"pnns_groups_1,omitempty"`
	PnnsGroups2                                     string `json:"pnns_groups_2,omitempty" bson:"pnns_groups_2,omitempty"`
	FoodGroups                                      string `json:"food_groups,omitempty" bson:"food_groups,omitempty"`
	FoodGroupsTags                                  string `json:"food_groups_tags,omitempty" bson:"food_groups_tags,omitempty"`
	FoodGroupsEn                                    string `json:"food_groups_en,omitempty" bson:"food_groups_en,omitempty"`
	States                                          string `json:"states,omitempty" bson:"states,omitempty"`
	StatesTags                                      string `json:"states_tags,omitempty" bson:"states_tags,omitempty"`
	StatesEn                                        string `json:"states_en,omitempty" bson:"states_en,omitempty"`
	BrandOwner                                      string `json:"brand_owner,omitempty" bson:"brand_owner,omitempty"`
	EcoscoreScore                                   string `json:"ecoscore_score,omitempty" bson:"ecoscore_score,omitempty"`
	EcoscoreGrade                                   string `json:"ecoscore_grade,omitempty" bson:"ecoscore_grade,omitempty"`
	NutrientLevelsTags                              string `json:"nutrient_levels_tags,omitempty" bson:"nutrient_levels_tags,omitempty"`
	ProductQuantity                                 string `json:"product_quantity,omitempty" bson:"product_quantity,omitempty"`
	Owner                                           string `json:"owner,omitempty" bson:"owner,omitempty"`
	DataQualityErrorsTags                           string `json:"data_quality_errors_tags,omitempty" bson:"data_quality_errors_tags,omitempty"`
	UniqueScansN                                    string `json:"unique_scans_n,omitempty" bson:"unique_scans_n,omitempty"`
	PopularityTags                                  string `json:"popularity_tags,omitempty" bson:"popularity_tags,omitempty"`
	Completeness                                    string `json:"completeness,omitempty" bson:"completeness,omitempty"`
	LastImageT                                      string `json:"last_image_t,omitempty" bson:"last_image_t,omitempty"`
	LastImageDatetime                               string `json:"last_image_datetime,omitempty" bson:"last_image_datetime,omitempty"`
	MainCategory                                    string `json:"main_category,omitempty" bson:"main_category,omitempty"`
	MainCategoryEn                                  string `json:"main_category_en,omitempty" bson:"main_category_en,omitempty"`
	ImageURL                                        string `json:"image_url,omitempty" bson:"image_url,omitempty"`
	ImageSmallURL                                   string `json:"image_small_url,omitempty" bson:"image_small_url,omitempty"`
	ImageIngredientsURL                             string `json:"image_ingredients_url,omitempty" bson:"image_ingredients_url,omitempty"`
	ImageIngredientsSmallURL                        string `json:"image_ingredients_small_url,omitempty" bson:"image_ingredients_small_url,omitempty"`
	ImageNutritionUrl                               string `json:"image_nutrition_url,omitempty" bson:"image_nutrition_url,omitempty"`
	ImageNutritionSmallURL                          string `json:"image_nutrition_small_url,omitempty" bson:"image_nutrition_small_url,omitempty"`
	Energykj100g                                    string `json:"energykj_100g,omitempty" bson:"energykj_100g,omitempty"`
	Energykcal100g                                  string `json:"energykcal_100g,omitempty" bson:"energykcal_100g,omitempty"`
	Energy100g                                      string `json:"energy_100g,omitempty" bson:"energy_100g,omitempty"`
	Energyfromfat100g                               string `json:"energyfromfat_100g,omitempty" bson:"energyfromfat_100g,omitempty"`
	Fat100g                                         string `json:"fat_100g,omitempty" bson:"fat_100g,omitempty"`
	Saturatedfat100g                                string `json:"saturatedfat_100g,omitempty" bson:"saturatedfat_100g,omitempty"`
	Butyricacid100g                                 string `json:"butyricacid_100g,omitempty" bson:"butyricacid_100g,omitempty"`
	Caproicacid100g                                 string `json:"caproicacid_100g,omitempty" bson:"caproicacid_100g,omitempty"`
	Caprylicacid100g                                string `json:"caprylicacid_100g,omitempty" bson:"caprylicacid_100g,omitempty"`
	Capricacid100g                                  string `json:"capricacid_100g,omitempty" bson:"capricacid_100g,omitempty"`
	Lauricacid100g                                  string `json:"lauricacid_100g,omitempty" bson:"lauricacid_100g,omitempty"`
	Myristicacid100g                                string `json:"myristicacid_100g,omitempty" bson:"myristicacid_100g,omitempty"`
	Palmiticacid100g                                string `json:"palmiticacid_100g,omitempty" bson:"palmiticacid_100g,omitempty"`
	Stearicacid100g                                 string `json:"stearicacid_100g,omitempty" bson:"stearicacid_100g,omitempty"`
	Arachidicacid100g                               string `json:"arachidicacid_100g,omitempty" bson:"arachidicacid_100g,omitempty"`
	Behenicacid100g                                 string `json:"behenicacid_100g,omitempty" bson:"behenicacid_100g,omitempty"`
	Lignocericacid100g                              string `json:"lignocericacid_100g,omitempty" bson:"lignocericacid_100g,omitempty"`
	Ceroticacid100g                                 string `json:"ceroticacid_100g,omitempty" bson:"ceroticacid_100g,omitempty"`
	Montanicacid100g                                string `json:"montanicacid_100g,omitempty" bson:"montanicacid_100g,omitempty"`
	Melissicacid100g                                string `json:"melissicacid_100g,omitempty" bson:"melissicacid_100g,omitempty"`
	Unsaturatedfat100g                              string `json:"unsaturatedfat_100g,omitempty" bson:"unsaturatedfat_100g,omitempty"`
	Monounsaturatedfat100g                          string `json:"monounsaturatedfat_100g,omitempty" bson:"monounsaturatedfat_100g,omitempty"`
	Omega9fat100g                                   string `json:"omega9fat_100g,omitempty" bson:"omega9fat_100g,omitempty"`
	Polyunsaturatedfat100g                          string `json:"polyunsaturatedfat_100g,omitempty" bson:"polyunsaturatedfat_100g,omitempty"`
	Omega3fat100g                                   string `json:"omega3fat_100g,omitempty" bson:"omega3fat_100g,omitempty"`
	Omega6fat100g                                   string `json:"omega6fat_100g,omitempty" bson:"omega6fat_100g,omitempty"`
	Alphalinolenicacid100g                          string `json:"alphalinolenicacid_100g,omitempty" bson:"alphalinolenicacid_100g,omitempty"`
	Eicosapentaenoicacid100g                        string `json:"eicosapentaenoicacid_100g,omitempty" bson:"eicosapentaenoicacid_100g,omitempty"`
	Docosahexaenoicacid100g                         string `json:"docosahexaenoicacid_100g,omitempty" bson:"docosahexaenoicacid_100g,omitempty"`
	Linoleicacid100g                                string `json:"linoleicacid_100g,omitempty" bson:"linoleicacid_100g,omitempty"`
	Arachidonicacid100g                             string `json:"arachidonicacid_100g,omitempty" bson:"arachidonicacid_100g,omitempty"`
	Gammalinolenicacid100g                          string `json:"gammalinolenicacid_100g,omitempty" bson:"gammalinolenicacid_100g,omitempty"`
	Dihomogammalinolenicacid100g                    string `json:"dihomogammalinolenicacid_100g,omitempty" bson:"dihomogammalinolenicacid_100g,omitempty"`
	Oleicacid100g                                   string `json:"oleicacid_100g,omitempty" bson:"oleicacid_100g,omitempty"`
	Elaidicacid100g                                 string `json:"elaidicacid_100g,omitempty" bson:"elaidicacid_100g,omitempty"`
	Gondoicacid100g                                 string `json:"gondoicacid_100g,omitempty" bson:"gondoicacid_100g,omitempty"`
	Meadacid100g                                    string `json:"meadacid_100g,omitempty" bson:"meadacid_100g,omitempty"`
	Erucicacid100g                                  string `json:"erucicacid_100g,omitempty" bson:"erucicacid_100g,omitempty"`
	Nervonicacid100g                                string `json:"nervonicacid_100g,omitempty" bson:"nervonicacid_100g,omitempty"`
	Transfat100g                                    string `json:"transfat_100g,omitempty" bson:"transfat_100g,omitempty"`
	Cholesterol100g                                 string `json:"cholesterol_100g,omitempty" bson:"cholesterol_100g,omitempty"`
	Carbohydrates100g                               string `json:"carbohydrates_100g,omitempty" bson:"carbohydrates_100g,omitempty"`
	Sugars100g                                      string `json:"sugars_100g,omitempty" bson:"sugars_100g,omitempty"`
	Addedsugars100g                                 string `json:"addedsugars_100g,omitempty" bson:"addedsugars_100g,omitempty"`
	Sucrose100g                                     string `json:"sucrose_100g,omitempty" bson:"sucrose_100g,omitempty"`
	Glucose100g                                     string `json:"glucose_100g,omitempty" bson:"glucose_100g,omitempty"`
	Fructose100g                                    string `json:"fructose_100g,omitempty" bson:"fructose_100g,omitempty"`
	Lactose100g                                     string `json:"lactose_100g,omitempty" bson:"lactose_100g,omitempty"`
	Maltose100g                                     string `json:"maltose_100g,omitempty" bson:"maltose_100g,omitempty"`
	Maltodextrins100g                               string `json:"maltodextrins_100g,omitempty" bson:"maltodextrins_100g,omitempty"`
	Starch100g                                      string `json:"starch_100g,omitempty" bson:"starch_100g,omitempty"`
	Polyols100g                                     string `json:"polyols_100g,omitempty" bson:"polyols_100g,omitempty"`
	Erythritol100g                                  string `json:"erythritol_100g,omitempty" bson:"erythritol_100g,omitempty"`
	Fiber100g                                       string `json:"fiber_100g,omitempty" bson:"fiber_100g,omitempty"`
	Solublefiber100g                                string `json:"solublefiber_100g,omitempty" bson:"solublefiber_100g,omitempty"`
	Insolublefiber100g                              string `json:"insolublefiber_100g,omitempty" bson:"insolublefiber_100g,omitempty"`
	Proteins100g                                    string `json:"proteins_100g,omitempty" bson:"proteins_100g,omitempty"`
	Casein100g                                      string `json:"casein_100g,omitempty" bson:"casein_100g,omitempty"`
	Serumproteins100g                               string `json:"serumproteins_100g,omitempty" bson:"serumproteins_100g,omitempty"`
	Nucleotides100g                                 string `json:"nucleotides_100g,omitempty" bson:"nucleotides_100g,omitempty"`
	Salt100g                                        string `json:"salt_100g,omitempty" bson:"salt_100g,omitempty"`
	Addedsalt100g                                   string `json:"addedsalt_100g,omitempty" bson:"addedsalt_100g,omitempty"`
	Sodium100g                                      string `json:"sodium_100g,omitempty" bson:"sodium_100g,omitempty"`
	Alcohol100g                                     string `json:"alcohol_100g,omitempty" bson:"alcohol_100g,omitempty"`
	Vitamina100g                                    string `json:"vitamina_100g,omitempty" bson:"vitamina_100g,omitempty"`
	Betacarotene100g                                string `json:"betacarotene_100g,omitempty" bson:"betacarotene_100g,omitempty"`
	Vitamind100g                                    string `json:"vitamind_100g,omitempty" bson:"vitamind_100g,omitempty"`
	Vitamine100g                                    string `json:"vitamine_100g,omitempty" bson:"vitamine_100g,omitempty"`
	Vitamink100g                                    string `json:"vitamink_100g,omitempty" bson:"vitamink_100g,omitempty"`
	Vitaminc100g                                    string `json:"vitaminc_100g,omitempty" bson:"vitaminc_100g,omitempty"`
	Vitaminb1100g                                   string `json:"vitaminb1_100g,omitempty" bson:"vitaminb1_100g,omitempty"`
	Vitaminb2100g                                   string `json:"vitaminb2_100g,omitempty" bson:"vitaminb2_100g,omitempty"`
	Vitaminpp100g                                   string `json:"vitaminpp_100g,omitempty" bson:"vitaminpp_100g,omitempty"`
	Vitaminb6100g                                   string `json:"vitaminb6_100g,omitempty" bson:"vitaminb6_100g,omitempty"`
	Vitaminb9100g                                   string `json:"vitaminb9_100g,omitempty" bson:"vitaminb9_100g,omitempty"`
	Folates100g                                     string `json:"folates_100g,omitempty" bson:"folates_100g,omitempty"`
	Vitaminb12100g                                  string `json:"vitaminb12_100g,omitempty" bson:"vitaminb12_100g,omitempty"`
	Biotin100g                                      string `json:"biotin_100g,omitempty" bson:"biotin_100g,omitempty"`
	Pantothenicacid100g                             string `json:"pantothenicacid_100g,omitempty" bson:"pantothenicacid_100g,omitempty"`
	Silica100g                                      string `json:"silica_100g,omitempty" bson:"silica_100g,omitempty"`
	Bicarbonate100g                                 string `json:"bicarbonate_100g,omitempty" bson:"bicarbonate_100g,omitempty"`
	Potassium100g                                   string `json:"potassium_100g,omitempty" bson:"potassium_100g,omitempty"`
	Chloride100g                                    string `json:"chloride_100g,omitempty" bson:"chloride_100g,omitempty"`
	Calcium100g                                     string `json:"calcium_100g,omitempty" bson:"calcium_100g,omitempty"`
	Phosphorus100g                                  string `json:"phosphorus_100g,omitempty" bson:"phosphorus_100g,omitempty"`
	Iron100g                                        string `json:"iron_100g,omitempty" bson:"iron_100g,omitempty"`
	Magnesium100g                                   string `json:"magnesium_100g,omitempty" bson:"magnesium_100g,omitempty"`
	Zinc100g                                        string `json:"zinc_100g,omitempty" bson:"zinc_100g,omitempty"`
	Copper100g                                      string `json:"copper_100g,omitempty" bson:"copper_100g,omitempty"`
	Manganese100g                                   string `json:"manganese_100g,omitempty" bson:"manganese_100g,omitempty"`
	Fluoride100g                                    string `json:"fluoride_100g,omitempty" bson:"fluoride_100g,omitempty"`
	Selenium100g                                    string `json:"selenium_100g,omitempty" bson:"selenium_100g,omitempty"`
	Chromium100g                                    string `json:"chromium_100g,omitempty" bson:"chromium_100g,omitempty"`
	Molybdenum100g                                  string `json:"molybdenum_100g,omitempty" bson:"molybdenum_100g,omitempty"`
	Iodine100g                                      string `json:"iodine_100g,omitempty" bson:"iodine_100g,omitempty"`
	Caffeine100g                                    string `json:"caffeine_100g,omitempty" bson:"caffeine_100g,omitempty"`
	Taurine100g                                     string `json:"taurine_100g,omitempty" bson:"taurine_100g,omitempty"`
	Ph100g                                          string `json:"ph_100g,omitempty" bson:"ph_100g,omitempty"`
	Fruitsvegetablesnuts100g                        string `json:"fruitsvegetablesnuts_100g,omitempty" bson:"fruitsvegetablesnuts_100g,omitempty"`
	Fruitsvegetablesnutsdried100g                   string `json:"fruitsvegetablesnutsdried_100g,omitempty" bson:"fruitsvegetablesnutsdried_100g,omitempty"`
	Fruitsvegetablesnutsestimate100g                string `json:"fruitsvegetablesnutsestimate_100g,omitempty" bson:"fruitsvegetablesnutsestimate_100g,omitempty"`
	Fruitsvegetablesnutsestimatefromingredients100g string `json:"fruitsvegetablesnutsestimatefromingredients_100g,omitempty" bson:"fruitsvegetablesnutsestimatefromingredients_100g,omitempty"`
	Collagenmeatproteinratio100g                    string `json:"collagenmeatproteinratio_100g,omitempty" bson:"collagenmeatproteinratio_100g,omitempty"`
	Cocoa100g                                       string `json:"cocoa_100g,omitempty" bson:"cocoa_100g,omitempty"`
	Chlorophyl100g                                  string `json:"chlorophyl_100g,omitempty" bson:"chlorophyl_100g,omitempty"`
	Carbonfootprint100g                             string `json:"carbonfootprint_100g,omitempty" bson:"carbonfootprint_100g,omitempty"`
	Carbonfootprintfrommeatorfish100g               string `json:"carbonfootprintfrommeatorfish_100g,omitempty" bson:"carbonfootprintfrommeatorfish_100g,omitempty"`
	Nutritionscorefr100g                            string `json:"nutritionscorefr_100g,omitempty" bson:"nutritionscorefr_100g,omitempty"`
	Nutritionscoreuk100g                            string `json:"nutritionscoreuk_100g,omitempty" bson:"nutritionscoreuk_100g,omitempty"`
	Glycemicindex100g                               string `json:"glycemicindex_100g,omitempty" bson:"glycemicindex_100g,omitempty"`
	Waterhardness100g                               string `json:"waterhardness_100g,omitempty" bson:"waterhardness_100g,omitempty"`
	Choline100g                                     string `json:"choline_100g,omitempty" bson:"choline_100g,omitempty"`
	Phylloquinone100g                               string `json:"phylloquinone_100g,omitempty" bson:"phylloquinone_100g,omitempty"`
	Betaglucan100g                                  string `json:"betaglucan_100g,omitempty" bson:"betaglucan_100g,omitempty"`
	Inositol100g                                    string `json:"inositol_100g,omitempty" bson:"inositol_100g,omitempty"`
	Carnitine100g                                   string `json:"carnitine_100g,omitempty" bson:"carnitine_100g,omitempty"`
	Sulphate100g                                    string `json:"sulphate_100g,omitempty" bson:"sulphate_100g,omitempty"`
	Nitrate100g                                     string `json:"nitrate_100g,omitempty" bson:"nitrate_100g,omitempty"`
	Acidity100g                                     string `json:"acidity_100g,omitempty" bson:"acidity_100g,omitempty"`
}

func (off OpenFoodFactCSVEntry) GetID() string {
	return off.Code
}

func (off *OpenFoodFactCSVEntry) SetID(id string) {
	off.Code = id
}
