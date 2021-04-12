package main

import "fmt"

//This file houses the models of the project. V2 should look to maybe split this into individualized model files due to the
//relation direction this project took
// Last Modified: Ren Thao
// Date: 04/12/2021

const (
	TYPECOCKTAIL           = "COCKTAIL"
	TYPEINGREDIENT         = "ingredient"
	TYPECOCKTAILINGREDIENT = "COCKTAILINGREDIENT"
	TYPEINSTRUCTION        = "INSTRUCTION"
	TYPEINSTRUCTIONTYPE    = "INSTRUCTIONLANGUAGE"
)

//APICocktail is the object structure of the thecocktaildb.com api's cocktail data
type APICocktail struct {
	IDDrink                     int    `json:"idDrink"`
	StrDrink                    string `json:"strDrink"`
	StrDrinkAlternate           string `json:"strDrinkAlternate"`
	StrTags                     string `json:"strTags"`
	StrVideo                    string `json:"strVideo"`
	StrCategory                 string `json:"strCategory"`
	StrIBA                      string `json:"strIBA"`
	StrAlcoholic                string `json:"strAlcoholic"`
	StrGlass                    string `json:"strGlass"`
	StrInstructions             string `json:"strInstructions"`
	StrInstructionsES           string `json:"strInstructionsES"`
	StrInstructionsDE           string `json:"strInstructionsDE"`
	StrInstructionsFR           string `json:"strInstructionsFR"`
	StrInstructionsIT           string `json:"strInstructionsIT"`
	StrInstructionsZHHANS       string `json:"strInstructionsZH-HANS"`
	StrInstructionsZHHANT       string `json:"strInstructionsZH-HANT"`
	StrDrinkThumb               string `json:"strDrinkThumb"`
	StrIngredient1              string `json:"strIngredient1"`
	StrIngredient2              string `json:"strIngredient2"`
	StrIngredient3              string `json:"strIngredient3"`
	StrIngredient4              string `json:"strIngredient4"`
	StrIngredient5              string `json:"strIngredient5"`
	StrIngredient6              string `json:"strIngredient6"`
	StrIngredient7              string `json:"strIngredient7"`
	StrIngredient8              string `json:"strIngredient8"`
	StrIngredient9              string `json:"strIngredient9"`
	StrIngredient10             string `json:"strIngredient10"`
	StrIngredient11             string `json:"strIngredient11"`
	StrIngredient12             string `json:"strIngredient12"`
	StrIngredient13             string `json:"strIngredient13"`
	StrIngredient14             string `json:"strIngredient14"`
	StrIngredient15             string `json:"strIngredient15"`
	StrMeasure1                 string `json:"strMeasure1"`
	StrMeasure2                 string `json:"strMeasure2"`
	StrMeasure3                 string `json:"strMeasure3"`
	StrMeasure4                 string `json:"strMeasure4"`
	StrMeasure5                 string `json:"strMeasure5"`
	StrMeasure6                 string `json:"strMeasure6"`
	StrMeasure7                 string `json:"strMeasure7"`
	StrMeasure8                 string `json:"strMeasure8"`
	StrMeasure9                 string `json:"strMeasure9"`
	StrMeasure10                string `json:"strMeasure10"`
	StrMeasure11                string `json:"strMeasure11"`
	StrMeasure12                string `json:"strMeasure12"`
	StrMeasure13                string `json:"strMeasure13"`
	StrMeasure14                string `json:"strMeasure14"`
	StrMeasure15                string `json:"strMeasure15"`
	StrImageSourcestring        string `json:"strImageSource"`
	StrImageAttributuion        string `json:"strImageAttribution"`
	StrCreativeCommonsConfirmed string `json:"strCreativeCommonsConfirmed"`
	DateModified                string `json:"dateModified"`
}

//simple oracle schema models. if I were to do a nosql variation, I would most likely create documents similarly
//however indexed by its unique constraint (likely name) and rather than a join table I'd use recipe: []values where
//values contains ingredient information + measurments

//Ingredient is the cocktail ingredient table model
type Ingredient struct {
	IngredientID   uint64 `json:"ingredient_id"`
	IngredientName string `json:"ingredient_name"`
}

//Cocktail is the cocktail table model
type Cocktail struct {
	CocktailID               int    `json:"cocktail_id"`
	CocktailName             string `json:"cocktail_name"`
	CocktailNameAlternate    string `json:"cocktail_name_alternate"`
	Tags                     string `json:"tags"`
	Video                    string `json:"video"`
	Category                 string `json:"category"`
	IBA                      string `json:"iba"`
	Alcoholic                string `json:"alcoholic"`
	Glass                    string `json:"glass"`
	DrinkThumb               string `json:"drinkthumb"`
	ImageSourcestring        string `json:"image_source"`
	ImageAttributuion        string `json:"image_attribution"`
	CreativeCommonsConfirmed string `json:"creative_commons_confirmed"`
	DateModified             string `json:"date_modified"`
}

//CocktailIngredient is the join table model for cocktails and ingredients. this will be the go to for lists for traversing between tables
type CocktailIngredient struct {
	CocktailIngredientID uint64 `json:"cocktail_ingredient_id"`
	CocktailID           uint64 `json:"cocktail_id"`
	IngredientID         uint64 `json:"ingredient_id"`
	Measurements         string `json:"measurements"`
}

type Instruction struct {
	InstrctionID      uint64 `json:"instruction_id"`
	CocktailID        uint64 `json:"cocktail_id"`
	InstructionTypeID uint64 `json:"instruction_type_id"`
	Instructions      string `json:"instructions"`
}

type InstructionType struct {
	InstructionTypeID       uint64 `json:"instruction_type_id"`
	InstructionTypeLanguage string `json:"instruction_type_language"`
}

//GetCocktails returns cocktail data based on filter parameters
func GetCocktails(filters []Filter) ([]Cocktail, error) {
	var cocktails []Cocktail
	return cocktails, nil
}

//GetIngredients returns ingredient data based on filter parameters
func GetIngredients(filters []Filter) ([]Ingredient, error) {
	var ingredients []Ingredient
	return ingredients, nil

}

//GetCocktailIngredients returns cocktail ingredient data based on filter parameters
func GetCocktailIngredients(filters []Filter) ([]CocktailIngredient, error) {
	var cocktailIngredients []CocktailIngredient
	return cocktailIngredients, nil

}

//GetInstructions returns instruction data based on filter parameters
func GetInstructions(filters []Filter) ([]Instruction, error) {
	var instruction []Instruction
	return instruction, nil

}

//GetInstructionTypes returns instruction type data based on filter parameters
func GetInstructionTypes(filters []Filter) ([]InstructionType, error) {
	var instructionType []InstructionType
	return instructionType, nil

}
func InsertCocktail() (uint64, error) {
	var insertedID uint64
	return insertedID, nil
}

func InsertIngredient(ingredient Ingredient) (uint64, error) {
	var insertedID uint64
	return insertedID, nil
}

func InsertCocktailIngredient(cocktailIngredient CocktailIngredient) (uint64, error) {
	var insertedID uint64
	return insertedID, nil
}

func InsertInstruction(instruction Instruction) (uint64, error) {
	var insertedID uint64
	return insertedID, nil
}

func InsertInstructionType(instructionType InstructionType) (uint64, error) {
	var insertedID uint64
	return insertedID, nil
}

//Delete is a generic delete function that allows deletion on filter. allows for large scale deletes, if needed.
//
//-note- a lot of power.
func Delete(filters []Filter, entity string) error {

	return nil
}

//InsertCocktailRecipe inserts an entire cocktail into the database.
//must pass along a map of ingredients and thier measurements. if nil assume null/non-applicable
//-note- this requires measurements to have been converted to its empty string form already to avoid type casting issues
func InsertCocktailRecipe(cocktail Cocktail, ingredientList map[uint64]string, instructions map[string]string) error {

	//insert cocktail
	insertedCocktailID, err := InsertCocktail()
	if err != nil {
		fmt.Println(err)
		return err
	}

	//not the most efficient 1.0 build however it allows for more flexibilty and can be fine-tuned for performance in the next versioning
	// we would likely need to cut down on the search queries with IN keywords due to the nature of having many ingredients and
	//likely many instructions
	for ingredientID, measurements := range ingredientList {
		//verify ingredient id exists
		filters := []Filter{}
		//we'll want to eventually phase out hard strings for pulling the actual field tags
		filters = append(filters, Filter{Entity: TYPEINGREDIENT, Field: "Ingredient_ID", Value: ingredientID})
		ingredients, err := GetIngredients(filters)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if len(ingredients) < 1 {
			err := fmt.Errorf("failed to find ingredient")
			fmt.Println(err)
			return err
		}
		//insert ingredientcocktail entry
		_, err = InsertCocktailIngredient(CocktailIngredient{IngredientID: ingredientID, CocktailID: insertedCocktailID, Measurements: measurements})

	}

	for instructionType, instruction := range instructions {
		//verify instruction type exists
		filters := []Filter{}
		filters = append(filters, Filter{Entity: TYPEINSTRUCTIONTYPE, Field: "Instruction_Type_Language", Value: instructionType})

		//
		instructionTypes, err := GetInstructionTypes(filters)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if len(instructionTypes) < 1 {
			err := fmt.Errorf("failed to find instruction type")
			fmt.Println(err)
			return err
		}
		toInsert := Instruction{CocktailID: insertedCocktailID, Instructions: instruction, InstructionTypeID: instructionTypes[0].InstructionTypeID}
		_, err = InsertInstruction(toInsert)
	}

	return nil
}
