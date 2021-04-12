package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
)

const port string = "8080"

var client *http.Client

func main() {
	client = &http.Client{Timeout: time.Second * 10}
	buildData()
	router := setupRouter()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("------- project is now listening ---------")
	log.Fatal(srv.ListenAndServe())
}

//setupRouter is a basic router function
func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/search", search)
	router.HandleFunc("/insert", insert)
	router.HandleFunc("/update", update)
	router.HandleFunc("/delete", delete)
	return router
}

//builds the data. this function is specifically added here assuming this is a "test application". this api would not normally
//set its own data if planning to run as a service
func buildData() error {
	cocktailData := callAPI()
	if len(cocktailData) == 0 {
		err := fmt.Errorf("failed to retrieve dataset")
		fmt.Println(err)
		return err
	}

	//if time allowed, I would have converted this data slice into usable entries per object via maps, one for
	//instruction type, one for ingredients, one for cocktails while duplicationg data. probably these would be indexed
	//like "name":"other values" as the map pairing. we would then loop to insert into x relation database (probably oracle 18x in my case)
	//since this is generally test data, I'd likely have to write a teardown script otherwise if this was a real application; this
	//would be moved to a "run" script that runs a DDL, needing only to run once

	return nil
}

func callAPI() []APICocktail {
	var cocktailList []APICocktail
	var respVal map[string]interface{}
	call := "https://www.thecocktaildb.com/api/json/v1/1/search.php?s"

	//more in depth call used rather than http.Get. can reuse or build into more complex api queries as needed
	req, err := http.NewRequest(http.MethodGet, call, nil)
	if err != nil {
		fmt.Println(err)
		return []APICocktail{}
	}
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []APICocktail{}
	}

	if resp == nil {
		err := fmt.Errorf("got back a nil response")
		fmt.Println(err)
		return []APICocktail{}
	}

	//decode response
	err = jsoniter.NewDecoder(resp.Body).Decode(&respVal)
	if err != nil {
		fmt.Println(err)
		return []APICocktail{}
	}

	//verify payload isnt nil
	if respVal["drinks"] == nil {
		return []APICocktail{}
	}

	//safe unboxing of interface
	payload, ok := respVal["drinks"].([]interface{})
	if !ok {
		err = fmt.Errorf("payload an unexpected type")
		fmt.Println(err)
		return []APICocktail{}
	}
	for _, v := range payload {
		cocktail := APICocktail{}

		//safe unboxing of interface
		cocktailMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("failed to convert into map type")
			fmt.Println(err)
			return []APICocktail{}
		}

		//verifies correct field typing before conversion, should only ever be string, however if changed to int type it will allow for it
		if reflect.TypeOf(cocktailMap["idDrink"]).String() == "string" {
			convertedValue, err := strconv.Atoi(cocktailMap["idDrink"].(string))
			if err != nil {
				fmt.Println(err)
				return []APICocktail{}
			}
			v.(map[string]interface{})["idDrink"] = convertedValue
		}

		//conversion into object
		dta, err := jsoniter.Marshal(v)
		if err != nil {
			fmt.Println(err)
			return []APICocktail{}
		}

		err = jsoniter.Unmarshal(dta, &cocktail)
		if err != nil {
			fmt.Println(err)
			return []APICocktail{}
		}

		cocktailList = append(cocktailList, cocktail)

	}

	return cocktailList
}
