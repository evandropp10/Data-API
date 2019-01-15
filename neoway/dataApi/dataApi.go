package dataApi

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
	"google.golang.org/appengine"
)

type Company struct {
	Name    string `json:"name,omitempty"`
	Zip     string `json:"zip,omitempty"`
	Website string `json:"website,omitempty"`
}

type CompanyID struct {
	Cid     string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Zip     string `json:"zip,omitempty"`
	Website string `json:"website,omitempty"`
}

type Log struct {
	Msg      string
	Function string
	Desc     string
}

type ReturnBody struct {
	Type    string
	Message string
	Extra   string
	Error   string
}

func GetCompany(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	name := strings.ToUpper(params["name"])
	zip := params["zip"]

	ctx := appengine.NewContext(req)

	// Search in database
	comp, err := GetCompanyDB(ctx, name, zip)
	if err != nil {
		w.Write([]byte("Not found."))
		return
	}

	json.NewEncoder(w).Encode(comp)
}

func GetCompanyDB(ctx context.Context, name string, zip string) (CompanyID, error) {
	IDprojeto := os.Getenv("PROJECT")

	var comp CompanyID
	var returnError error

	datastoreClient, err := datastore.NewClient(ctx, IDprojeto)
	if err != nil {
		returnError = errors.New("Cannot create datastore client")
		return comp, returnError
	}

	// Sets the kind for the new entity.
	tipo := "companies"

	q := datastore.NewQuery(tipo).
		Filter("Zip = ", zip)

	ents := datastoreClient.Run(ctx, q)
	if ents == nil {
		returnError = errors.New("Not found")
		return comp, returnError
	}

	for {
		var c CompanyID
		key, err := ents.Next(&c)
		if err == iterator.Done {
			if comp.Name == "" {
				returnError = errors.New("Not found")
			}
			break
		}

		if strings.Contains(c.Name, name) {
			comp = c
			comp.Cid = strconv.FormatInt(key.ID, 10)
		}
	}

	return comp, returnError
}

func LoadCompanies(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	reader := csv.NewReader(req.Body)

	var comps [][]string
	var erro string

	// Data treatment
	for {
		// read rows from csv
		record, err := reader.Read()
		if err != nil {
			erro = err.Error()
			break
		}

		arrCSV := strings.Split(record[0], ";")

		if len(arrCSV) != 2 {
			erro = "Bad format. The CSV must have two columns separated by ;"
			break
		}

		// discart header
		if arrCSV[0] == "name" {
			continue
		}

		// Upper Name
		arrCSV[0] = strings.ToUpper(arrCSV[0])

		comps = append(comps, arrCSV)
	}

	// Load in database
	errdb := LoadCompaniesDB(ctx, comps)
	if errdb != nil {
		erro = errdb.Error()
	}

	if erro != "" && erro != "EOF" {
		w.WriteHeader(500)
		w.Write([]byte(erro))
	} else {
		w.Write([]byte("Data was recorded."))
	}
}

func LoadCompaniesDB(ctx context.Context, comps [][]string) error {

	IDprojeto := os.Getenv("PROJECT")

	var returnError error

	datastoreClient, err := datastore.NewClient(ctx, IDprojeto)
	if err != nil {
		returnError = errors.New("Cannot create datastore client")
		return returnError
	}

	// Sets the kind for the new entity.
	tipo := "companies"

	for i := 0; i < len(comps); i++ {
		comp := &Company{
			Name:    comps[i][0],
			Zip:     comps[i][1],
			Website: "",
		}

		// Creates a Key instance.
		compKey := datastore.IncompleteKey(tipo, nil)

		// Saves the new entity.
		_, erro := datastoreClient.Put(ctx, compKey, comp)
		if erro != nil {
			returnError = errors.New("Cannot Put entity")
			return returnError
		}
	}
	return returnError
}

func IntegrateCompanies(w http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)

	reader := csv.NewReader(req.Body)

	var comps [][]string
	var erro string

	// Data treatment
	for {
		// read rows from csv
		record, err := reader.Read()
		if err != nil {
			erro = err.Error()
			break
		}

		arrCSV := strings.Split(record[0], ";")

		if len(arrCSV) != 3 {
			erro = "Bad format. The CSV must have three columns separated by ;"
			break
		}

		// discart header
		if arrCSV[0] == "name" {
			continue
		}

		// Upper Name
		arrCSV[0] = strings.ToUpper(arrCSV[0])
		// lower Website
		arrCSV[2] = strings.ToLower(arrCSV[2])

		comps = append(comps, arrCSV)

	}

	// Integrate with database
	errdb := IntegrateCompaniesDB(ctx, comps)
	if errdb != nil {
		erro = errdb.Error()
	}

	if erro != "" && erro != "EOF" {
		w.WriteHeader(500)
		w.Write([]byte(erro))
	} else {
		w.Write([]byte("Data was integrated."))
	}
}

func IntegrateCompaniesDB(ctx context.Context, comps [][]string) error {

	IDprojeto := os.Getenv("PROJECT")

	var companies []*Company
	var returnError error

	datastoreClient, err := datastore.NewClient(ctx, IDprojeto)
	if err != nil {
		returnError = errors.New("Cannot create datastore client")
		return returnError
	}

	// Sets the kind for the new entity.
	tipo := "companies"

	for i := 0; i < len(comps); i++ {
		comp := &Company{
			Name:    comps[i][0],
			Zip:     comps[i][1],
			Website: comps[i][2],
		}

		q := datastore.NewQuery(tipo).
			Filter("Name = ", comp.Name).KeysOnly()

		keys, err := datastoreClient.GetAll(ctx, q, &companies)
		if err != nil {
			returnError = errors.New("Cannot Get Entity")
			return returnError
		}

		// If found, then modify entity, else create new entity
		if len(keys) > 0 {
			for j := 0; j < len(keys); j++ {
				_, err1 := datastoreClient.Put(ctx, keys[j], comp)
				if err1 != nil {
					returnError = errors.New("Cannot Put Entity")
					return returnError
				}
			}
		} else {
			compKey := datastore.IncompleteKey(tipo, nil)

			// Saves the new entity.
			_, erro := datastoreClient.Put(ctx, compKey, comp)
			if erro != nil {
				returnError = errors.New("Cannot Put Entity")
				return returnError
			}
		}
	}
	return returnError
}

func ErrorLog(msg string, function string, desc string, ctx context.Context) {

	IDprojeto := os.Getenv("PROJECT")

	comp := &Log{
		Msg:      msg,
		Function: function,
		Desc:     desc,
	}

	datastoreClient, err := datastore.NewClient(ctx, IDprojeto)
	if err != nil {
		//fmt.Errorf(err.Error())
	}

	// Sets the kind for the new entity.
	tipo := "log"

	// Creates a Key instance.
	compKey := datastore.IncompleteKey(tipo, nil)

	// Saves the new entity.
	_, erro := datastoreClient.Put(ctx, compKey, comp)
	if erro != nil {
		//fmt.Printf("Failed to save task: %v", logKey, erro)
	}
}
