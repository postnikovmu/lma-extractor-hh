package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

type Areas []struct {
	Name     string      `json:"name"`
	ID       string      `json:"id"`
	ParentID interface{} `json:"parent_id"`
	Areas    []struct {
		Name     string        `json:"name"`
		ID       string        `json:"id"`
		ParentID string        `json:"parent_id"`
		Areas    []interface{} `json:"areas"`
	} `json:"areas"`
}

type Areas2 []struct {
	Name     string      `json:"name"`
	ID       string      `json:"id"`
	ParentID interface{} `json:"parent_id"`
	Areas2   Areas2      `json:"areas"`
}

type pages struct {
	Items []struct {
		ID                     string      `json:"id"`
		Premium                bool        `json:"premium"`
		Name                   string      `json:"name"`
		Department             interface{} `json:"department"`
		HasTest                bool        `json:"has_test"`
		ResponseLetterRequired bool        `json:"response_letter_required"`
		Area                   struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"area"`
		Salary interface{} `json:"salary"`
		Type   struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		Address struct {
			City        string      `json:"city"`
			Street      string      `json:"street"`
			Building    string      `json:"building"`
			Description interface{} `json:"description"`
			Lat         float64     `json:"lat"`
			Lng         float64     `json:"lng"`
			Raw         string      `json:"raw"`
			Metro       struct {
				StationName string  `json:"station_name"`
				LineName    string  `json:"line_name"`
				StationID   string  `json:"station_id"`
				LineID      string  `json:"line_id"`
				Lat         float64 `json:"lat"`
				Lng         float64 `json:"lng"`
			} `json:"metro"`
			MetroStations []struct {
				StationName string  `json:"station_name"`
				LineName    string  `json:"line_name"`
				StationID   string  `json:"station_id"`
				LineID      string  `json:"line_id"`
				Lat         float64 `json:"lat"`
				Lng         float64 `json:"lng"`
			} `json:"metro_stations"`
			ID string `json:"id"`
		} `json:"address"`
		ResponseURL       interface{}   `json:"response_url"`
		SortPointDistance interface{}   `json:"sort_point_distance"`
		PublishedAt       string        `json:"published_at"`
		CreatedAt         string        `json:"created_at"`
		Archived          bool          `json:"archived"`
		ApplyAlternateURL string        `json:"apply_alternate_url"`
		InsiderInterview  interface{}   `json:"insider_interview"`
		URL               string        `json:"url"`
		AdvResponseURL    string        `json:"adv_response_url"`
		AlternateURL      string        `json:"alternate_url"`
		Relations         []interface{} `json:"relations"`
		Employer          struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			URL          string `json:"url"`
			AlternateURL string `json:"alternate_url"`
			LogoUrls     struct {
				Num90    string `json:"90"`
				Num240   string `json:"240"`
				Original string `json:"original"`
			} `json:"logo_urls"`
			VacanciesURL string `json:"vacancies_url"`
			Trusted      bool   `json:"trusted"`
		} `json:"employer"`
		Snippet struct {
			Requirement    string `json:"requirement"`
			Responsibility string `json:"responsibility"`
		} `json:"snippet"`
		Contacts interface{} `json:"contacts"`
		Schedule struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"schedule"`
		WorkingDays          []interface{} `json:"working_days"`
		WorkingTimeIntervals []interface{} `json:"working_time_intervals"`
		WorkingTimeModes     []interface{} `json:"working_time_modes"`
		AcceptTemporary      bool          `json:"accept_temporary"`
	} `json:"items"`
	Found        int         `json:"found"`
	Pages        int         `json:"pages"`
	PerPage      int         `json:"per_page"`
	Page         int         `json:"page"`
	Clusters     interface{} `json:"clusters"`
	Arguments    interface{} `json:"arguments"`
	AlternateURL string      `json:"alternate_url"`
}

//var gDictHhareas Areas
var gDictHhareas Areas2

type hhVac struct {
	ID          string `json:"id"`
	Premium     bool   `json:"premium"`
	BillingType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"billing_type"`
	Relations              []interface{} `json:"relations"`
	Name                   string        `json:"name"`
	InsiderInterview       interface{}   `json:"insider_interview"`
	ResponseLetterRequired bool          `json:"response_letter_required"`
	Area                   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"area"`
	Salary interface{} `json:"salary"`
	Type   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Address       interface{} `json:"address"`
	AllowMessages bool        `json:"allow_messages"`
	Experience    struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"experience"`
	Schedule struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"schedule"`
	Employment struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"employment"`
	Department                 interface{} `json:"department"`
	Contacts                   interface{} `json:"contacts"`
	Description                string      `json:"description"`
	BrandedDescription         interface{} `json:"branded_description"`
	VacancyConstructorTemplate interface{} `json:"vacancy_constructor_template"`
	KeySkills                  []struct {
		Name string `json:"name"`
	} `json:"key_skills"`
	AcceptHandicapped bool        `json:"accept_handicapped"`
	AcceptKids        bool        `json:"accept_kids"`
	Archived          bool        `json:"archived"`
	ResponseURL       interface{} `json:"response_url"`
	Specializations   []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		ProfareaID   string `json:"profarea_id"`
		ProfareaName string `json:"profarea_name"`
	} `json:"specializations"`
	ProfessionalRoles []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"professional_roles"`
	Code                    interface{}   `json:"code"`
	Hidden                  bool          `json:"hidden"`
	QuickResponsesAllowed   bool          `json:"quick_responses_allowed"`
	DriverLicenseTypes      []interface{} `json:"driver_license_types"`
	AcceptIncompleteResumes bool          `json:"accept_incomplete_resumes"`
	Employer                struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		URL          string `json:"url"`
		AlternateURL string `json:"alternate_url"`
		LogoUrls     struct {
			Num90    string `json:"90"`
			Num240   string `json:"240"`
			Original string `json:"original"`
		} `json:"logo_urls"`
		VacanciesURL string `json:"vacancies_url"`
		Trusted      bool   `json:"trusted"`
	} `json:"employer"`
	PublishedAt          string        `json:"published_at"`
	CreatedAt            string        `json:"created_at"`
	NegotiationsURL      interface{}   `json:"negotiations_url"`
	SuitableResumesURL   interface{}   `json:"suitable_resumes_url"`
	ApplyAlternateURL    string        `json:"apply_alternate_url"`
	HasTest              bool          `json:"has_test"`
	Test                 interface{}   `json:"test"`
	AlternateURL         string        `json:"alternate_url"`
	WorkingDays          []interface{} `json:"working_days"`
	WorkingTimeIntervals []interface{} `json:"working_time_intervals"`
	WorkingTimeModes     []interface{} `json:"working_time_modes"`
	AcceptTemporary      bool          `json:"accept_temporary"`
	Languages            []interface{} `json:"languages"`
}

type vacancie struct {
	URL          string `json:"strUrl"`
	Name         string `json:"strJobTitle"`
	AreaName     string `json:"strArea"`
	EmployerName string `json:"strCompany"`
	Description  string `json:"strBodyFull"`
	KeySkills    []struct {
		Name string `json:"name"`
	} `json:"strArrKeySkills"`
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
		log.Fatal("Sorry, an error occurred, please try again")
	}
	defer resp.Body.Close()
	//Create a variable of the same type as our model
	var cResp hhresponse
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("Sorry, an error occurred, please try again")
	}
	fmt.Fprint(res, cResp.Found)
}

func hh3(res http.ResponseWriter, req *http.Request) {
	fmt.Println(0, req.URL)
	//Build The URL string
	URL := "https://api.hh.ru/vacancies?text=Golang&area=66"
	fmt.Println(1)
	//We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("Sorry, an error occurred, please try again")
	}
	defer resp.Body.Close()
	fmt.Println(2)

	//Create a variable of the same type as our model
	var cResp_full hhresponse_full
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp_full); err != nil {
		log.Fatal("Sorry, an error occurred, please try again")
	}

	fmt.Println(3)
	//fmt.Fprint(res, cResp.Found)
	fmt.Fprint(res, cResp_full)
	fmt.Println(4)
}

func hh4(res http.ResponseWriter, req *http.Request) {
	//reqUrlQuery := strings.Split(req.URL.RawQuery, "&")
	reqUrlQuery, _ := url.ParseQuery(req.URL.RawQuery)

	var lvText = reqUrlQuery["text"][0]          //variable with the text of search
	var lvAreaText = reqUrlQuery["area"][0]      //variable with the area of search as text
	lvArea := findNode(gDictHhareas, lvAreaText) //variable with the area ID from api

	var ltReturnedVac []vacancie
	//ltReturnedVac = make([]vacancie, 0)

	const lcPerPage = 100
	lvPageNumber := 0
	for i := 0; i < 20; i++ {
		lvPageNumber = i
		lvPagesFromAPI := getPageOfVacancies(lcPerPage, lvPageNumber, lvText, lvArea, &ltReturnedVac)
		if lvPagesFromAPI <= (i + 1) {
			break
		}
	}
	//j, _ := json.Marshal(ltReturnedVac)
	j, _ := JSONMarshal(ltReturnedVac, true)

	res.Header().Set("Content-Type", "application/json")
	//res.WriteHeader(http.StatusCreated)
	res.Write(j)
	//fmt.Fprint(res, string(j))
}

func getPageOfVacancies(ivPerPage int, ivPageNumber int, ivText string, ivArea string, ctReturnedVac *[]vacancie) int {
	lvPageNumberStr := strconv.Itoa(ivPageNumber)
	lvPerPage := strconv.Itoa(ivPerPage)

	lvText20 := strings.Join(strings.Split(ivText, " "), "%20")
	//Build The URL string
	URL := "https://api.hh.ru/vacancies?" + "text=" + lvText20 + "&" + "area=" + ivArea + "&" + "per_page=" + lvPerPage + "&" + "page=" + lvPageNumberStr

	//We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("Sorry, an error occurred, please try again")
	}

	//Create a variable of the same type as our model
	var lsPage pages
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&lsPage); err != nil {
		log.Fatal("Sorry, an error occurred, please try again")
	}

	for _, value := range lsPage.Items {
		var lsVac = vacancie{}
		lsVac.Name = value.Name
		lsVac.URL = value.URL
		lsVacDetailed := getVacancie(lsVac.URL)
		lsVac.Description = lsVacDetailed.Description
		lsVac.KeySkills = lsVacDetailed.KeySkills
		lsVac.AreaName = lsVacDetailed.Area.Name
		lsVac.EmployerName = lsVacDetailed.Employer.Name
		*ctReturnedVac = append(*ctReturnedVac, lsVac)
	}

	//fmt.Println(lsPage.Pages, lvPageNumberStr, lvPerPage)
	return lsPage.Pages
}

func getVacancie(ivURL string) hhVac {
	//We make HTTP request using the Get function
	resp, err := http.Get(ivURL)
	if err != nil {
		log.Fatal("Sorry, an error occurred, please try again")
	}

	//Create a variable of the same type as our model
	var lsVac hhVac
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&lsVac); err != nil {
		log.Fatal("Sorry, an error occurred, please try again")
	}
	return lsVac

}

func getHhAreas() {
	//Build The URL string
	URL := "https://api.hh.ru/areas"
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("Sorry, an error occurred, please try again")
	}
	defer resp.Body.Close()
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&gDictHhareas); err != nil {
		log.Fatal("Sorry, an error occurred, please try again")
	}
}

func findNode(n Areas2, s string) string {
	for _, value := range n {

		if value.Name == s {
			return value.ID
		} else {
			var nChild Areas2
			nChild = value.Areas2
			result := findNode(nChild, s)
			if result != "" {
				return result
			}
		}

	}
	return ""
}

func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func main() {

	getHhAreas()

	http.HandleFunc("/", hh)
	//http.HandleFunc("/hh", hh1)
	//http.HandleFunc("/hh2", hh2)
	//http.HandleFunc("/hh3", hh3)
	http.HandleFunc("/hh4", hh4)
	//http.ListenAndServe("localhost:8080", nil) //locally
	http.ListenAndServe(":8080", nil) //SAP BTP CloudFounry
}
