package main

import (
	"fmt"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type request struct {
	EntityType string `json:"entity_type"`
	//may need to update this naming convention. can be usable for multiple things outside of filters
	Filters []Filter `json:"filter"`
}

type insertRequest struct {
	EntityType string `json:"entity_type"`
	//may need to update this naming convention. can be usable for multiple things outside of filters
	InsertPayload map[string]interface{} `json:"insert_payload"`
}

type insertCocktail struct {
	CocktailValue Cocktail          `json:"cocktail_value"`
	Ingredients   map[uint64]string `json:"ingredients"`
	Instructions  map[string]string `json:"instructions"`
}

//Filter struct is the nested json that will allow for search via name, id with is not functionality
//if both id and name are populated, ID takes precedent. verify this functionality with project lead
//to determine if this is correct course of action
//
//-Note- filter works well for this simple schema however with more complex schemas, it could cause an issue
// due to non-relation filters being passed depending on how front end handles its pagination.
//
//-Note- would like to make this more flexible, however due to time consraints not fesible to implement. we could
//clean this up by allowing IN values into the framework, giving list searches more flexiblity without multiple
//database calls
type Filter struct {
	Entity string      `json:"entity"`
	Field  string      `json:"field"`
	Value  interface{} `json:"value"`
	Is     bool        `json:"is"`
}

//errorStruct is the return response for error'd handlers
type errorStruct struct {
	Err error `json:"error"`
}

//we can use this endpoint to run a wide search on any table, provided which page a web-app is on
func search(w http.ResponseWriter, r *http.Request) {
	var req request
	w.Header().Add("Content-Type", "application/json")
	if !strings.EqualFold(r.Method, "Get") {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := fmt.Errorf("invalid request type")
		fmt.Println(err)

		response := errorStruct{Err: err}
		jsoniter.NewEncoder(w).Encode(response)
		return
	}

	err := jsoniter.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := errorStruct{Err: err}
		jsoniter.NewEncoder(w).Encode(response)
		return
	}

	//this will return relevant db information however it will miss
	//relavent string information needed for visualization. we'll likely
	//need a view object for these in model and will need to add joins
	//for relevent infromation depending on table being searched
	switch req.EntityType {
	case TYPECOCKTAIL:
		GetCocktails(req.Filters)
	case TYPEINGREDIENT:
		GetIngredients(req.Filters)
	case TYPECOCKTAILINGREDIENT:
		GetCocktailIngredients(req.Filters)
	case TYPEINSTRUCTION:
		GetInstructions(req.Filters)
	default:
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	//validate that item is stored and delete, return non-existant if 404
}

//payload should be a slice of id's if possible. there is a possibility of deleting by field/search parameters as well. however I cannot see
// see the functional value in it via the given controls. work to ask if this is something that should be looked into for a v2.
func delete(w http.ResponseWriter, r *http.Request) {
	var req request
	w.Header().Add("Content-Type", "application/json")
	if !strings.EqualFold(r.Method, "Get") {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := fmt.Errorf("invalid request type")
		fmt.Println(err)

		response := errorStruct{Err: err}
		jsoniter.NewEncoder(w).Encode(response)
		return
	}

	err := jsoniter.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := errorStruct{Err: err}
		jsoniter.NewEncoder(w).Encode(response)
		return
	}

	err = Delete(req.Filters, req.EntityType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := errorStruct{Err: err}
		jsoniter.NewEncoder(w).Encode(response)
		return
	}

}

func insert(w http.ResponseWriter, r *http.Request) {
	var req insertRequest
	w.Header().Add("Content-Type", "application/json")
	if !strings.EqualFold(r.Method, "Get") {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := fmt.Errorf("invalid request type")
		fmt.Println(err)

		response := errorStruct{Err: err}
		jsoniter.NewEncoder(w).Encode(response)
		return
	}

	err := jsoniter.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := errorStruct{Err: err}
		jsoniter.NewEncoder(w).Encode(response)
		return
	}

	//switch case to insert a new entity into the database. this allows for new language types, updating instructions, and updating recipe lists
	//this covers a good chunk of update when it comes to the recipies themselves, however it wont be practical to rely on this entry point as an update
	//especially when it comes to recipe entry updates
	switch req.EntityType {
	case TYPECOCKTAIL:
		//special payload. will need the frontend side to likely list relevant information already on a rendered cocktail form
		var cocktailPayload insertCocktail
		data, err := jsoniter.Marshal(req.InsertPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
		err = jsoniter.Unmarshal(data, &cocktailPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
		err = InsertCocktailRecipe(cocktailPayload.CocktailValue, cocktailPayload.Ingredients, cocktailPayload.Instructions)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
	case TYPEINGREDIENT:
		var ingredient Ingredient
		data, err := jsoniter.Marshal(req.InsertPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
		err = jsoniter.Unmarshal(data, &ingredient)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
		_, err = InsertIngredient(ingredient)
	case TYPECOCKTAILINGREDIENT:
		var cocktailIngredient CocktailIngredient
		data, err := jsoniter.Marshal(req.InsertPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
		err = jsoniter.Unmarshal(data, &cocktailIngredient)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
		_, err = InsertCocktailIngredient(cocktailIngredient)
	case TYPEINSTRUCTION:
		var instruction Instruction
		data, err := jsoniter.Marshal(req.InsertPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
		err = jsoniter.Unmarshal(data, &instruction)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
		_, err = InsertInstruction(instruction)
	case TYPEINSTRUCTIONTYPE:
		var instructionType InstructionType
		data, err := jsoniter.Marshal(req.InsertPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
		err = jsoniter.Unmarshal(data, &instructionType)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := errorStruct{Err: err}
			jsoniter.NewEncoder(w).Encode(response)
			return
		}
		_, err = InsertInstructionType(instructionType)
	}
}
