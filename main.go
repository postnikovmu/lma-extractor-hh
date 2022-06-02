package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type hhresponse struct {
	//NumberVacFound string `json:"found"`
	Found int `json:"found"`
}

type hhresponse_full struct {
	PerPage int `json:"per_page"`
	Items   []struct {
		Salary struct {
			To       interface{} `json:"to"`
			From     int         `json:"from"`
			Currency string      `json:"currency"`
			Gross    bool        `json:"gross"`
		} `json:"salary"`
		Name             string `json:"name"`
		InsiderInterview struct {
			ID  string `json:"id"`
			URL string `json:"url"`
		} `json:"insider_interview"`
		Area struct {
			URL  string `json:"url"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		URL         string        `json:"url"`
		PublishedAt string        `json:"published_at"`
		Relations   []interface{} `json:"relations"`
		Employer    struct {
			LogoUrls struct {
				Num90    string `json:"90"`
				Num240   string `json:"240"`
				Original string `json:"original"`
			} `json:"logo_urls"`
			Name         string `json:"name"`
			URL          string `json:"url"`
			AlternateURL string `json:"alternate_url"`
			ID           string `json:"id"`
			Trusted      bool   `json:"trusted"`
		} `json:"employer"`
		Contacts struct {
			Name   string `json:"name"`
			Email  string `json:"email"`
			Phones []struct {
				Country string      `json:"country"`
				City    string      `json:"city"`
				Number  string      `json:"number"`
				Comment interface{} `json:"comment"`
			} `json:"phones"`
		} `json:"contacts"`
		ResponseLetterRequired bool `json:"response_letter_required"`
		Address                struct {
			City          string  `json:"city"`
			Street        string  `json:"street"`
			Building      string  `json:"building"`
			Description   string  `json:"description"`
			Lat           float64 `json:"lat"`
			Lng           float64 `json:"lng"`
			MetroStations []struct {
				StationID   string  `json:"station_id"`
				StationName string  `json:"station_name"`
				LineID      string  `json:"line_id"`
				LineName    string  `json:"line_name"`
				Lat         float64 `json:"lat"`
				Lng         float64 `json:"lng"`
			} `json:"metro_stations"`
		} `json:"address"`
		SortPointDistance float64 `json:"sort_point_distance"`
		AlternateURL      string  `json:"alternate_url"`
		ApplyAlternateURL string  `json:"apply_alternate_url"`
		Department        struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"department"`
		Type struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		ID          string      `json:"id"`
		HasTest     bool        `json:"has_test"`
		ResponseURL interface{} `json:"response_url"`
		Snippet     struct {
			Requirement    string `json:"requirement"`
			Responsibility string `json:"responsibility"`
		} `json:"snippet"`
		Schedule struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"schedule"`
		Counters struct {
			Responses int `json:"responses"`
		} `json:"counters"`
	} `json:"items"`
	Page      int         `json:"page"`
	Pages     int         `json:"pages"`
	Found     int         `json:"found"`
	Clusters  interface{} `json:"clusters"`
	Arguments interface{} `json:"arguments"`
}

func hh(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello from Go")
}

func hh1(res http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("https://api.hh.ru/vacancies?text=Golang&area=66")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)

	fmt.Fprint(res, sb)

}

func hh2(res http.ResponseWriter, req *http.Request) {

	//Build The URL string
	URL := "https://api.hh.ru/vacancies?text=Golang&area=66"
	//We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()
	//Create a variable of the same type as our model
	var cResp hhresponse
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	fmt.Fprint(res, cResp.Found)
}

func hh3(res http.ResponseWriter, req *http.Request) {
	fmt.Println(0)
	//Build The URL string
	URL := "https://api.hh.ru/vacancies?text=Golang&area=66"
	fmt.Println(1)
	//We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()
	fmt.Println(2)

	//Create a variable of the same type as our model
	var cResp_full hhresponse_full
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp_full); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	fmt.Println(3)
	//fmt.Fprint(res, cResp.Found)
	fmt.Fprint(res, cResp_full)
	fmt.Println(4)
}

func main() {
	http.HandleFunc("/", hh)
	http.HandleFunc("/hh", hh1)
	http.HandleFunc("/hh2", hh2)
	http.HandleFunc("/hh3", hh3)
	http.ListenAndServe("localhost:4000", nil)
}
