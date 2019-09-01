package vkapi

import "strconv"

// ===================
//  DatabaseGetChairs
// ===================

// DatabaseGetChairsParams параметры метода DatabaseGetChairs.
type DatabaseGetChairsParams struct {
	FacultyID uint
	Offset    uint
	Count     uint
}

// DatabaseGetChairsResp структура, возвращаемая методом DatabaseGetChairs.
type DatabaseGetChairsResp struct {
	Count int `json:"count"`
	Items []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"items"`
}

// DatabaseGetChairs возвращает список кафедр университета по указанному факультету.
func (api *API) DatabaseGetChairs(p DatabaseGetChairsParams) (*DatabaseGetChairsResp, error) {
	resp, err := api.Request("database.getChairs", p, new(DatabaseGetChairsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetChairsResp), nil
}

// ===================
//  DatabaseGetCities
// ===================

// DatabaseGetCitiesParams параметры метода DatabaseGetCities.
type DatabaseGetCitiesParams struct {
	CountryID uint
	RegionID  uint
	Q         string
	NeedAll   bool
	Offset    uint
	Count     uint
}

// DatabaseGetCitiesResp структура, возвращаемая методом DatabaseGetCities.
type DatabaseGetCitiesResp struct {
	Count int `json:"count"`
	Items []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Area      string `json:"area"`
		Region    string `json:"region"`
		Important int    `json:"important"`
	} `json:"items"`
}

// DatabaseGetCities возвращает список городов.
func (api *API) DatabaseGetCities(p DatabaseGetCitiesParams) (*DatabaseGetCitiesResp, error) {
	resp, err := api.Request("database.getCities", p, new(DatabaseGetCitiesResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetCitiesResp), nil
}

// =======================
//  DatabaseGetCitiesByID
// =======================

// DatabaseGetCitiesByIDParams параметры метода DatabaseGetCitiesByID.
type DatabaseGetCitiesByIDParams struct {
	CityIDS []int
}

// DatabaseGetCitiesByIDResp структура, возвращаемая методом DatabaseGetCitiesByID.
type DatabaseGetCitiesByIDResp []struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// DatabaseGetCitiesByID возвращает информацию о городах и регионах по их идентификаторам. Идентификаторы (id) могут быть получены с помощью методов users.get, places.getbyid, places.search, places.getcheckins.
func (api *API) DatabaseGetCitiesByID(p DatabaseGetCitiesByIDParams) (*DatabaseGetCitiesByIDResp, error) {
	resp, err := api.Request("database.getCitiesById", p, new(DatabaseGetCitiesByIDResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetCitiesByIDResp), nil
}

// ======================
//  DatabaseGetCountries
// ======================

// DatabaseGetCountriesParams параметры метода DatabaseGetCountries.
type DatabaseGetCountriesParams struct {
	NeedAll bool
	Code    string
	Offset  uint
	Count   uint
}

// DatabaseGetCountriesResp структура, возвращаемая методом DatabaseGetCountries.
type DatabaseGetCountriesResp struct {
	Count int `json:"count"`
	Items []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"items"`
}

// DatabaseGetCountries возвращает список стран.
func (api *API) DatabaseGetCountries(p DatabaseGetCountriesParams) (*DatabaseGetCountriesResp, error) {
	resp, err := api.Request("database.getCountries", p, new(DatabaseGetCountriesResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetCountriesResp), nil
}

// ==========================
//  DatabaseGetCountriesByID
// ==========================

// DatabaseGetCountriesByIDParams параметры метода DatabaseGetCountriesByID.
type DatabaseGetCountriesByIDParams struct {
	CountryIDS []int
}

// DatabaseGetCountriesByIDResp структура, возвращаемая методом DatabaseGetCountriesByID.
type DatabaseGetCountriesByIDResp []struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// DatabaseGetCountriesByID возвращает информацию о странах по их идентификаторам. Идентификаторы (id) могут быть получены с помощью методов users.get, places.getbyid, places.search, places.getcheckins.
func (api *API) DatabaseGetCountriesByID(p DatabaseGetCountriesByIDParams) (*DatabaseGetCountriesByIDResp, error) {
	resp, err := api.Request("database.getCountriesById", p, new(DatabaseGetCountriesByIDResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetCountriesByIDResp), nil
}

// ======================
//  DatabaseGetFaculties
// ======================

// DatabaseGetFacultiesParams параметры метода DatabaseGetFaculties.
type DatabaseGetFacultiesParams struct {
	UniversityID uint
	Offset       uint
	Count        uint
}

// DatabaseGetFacultiesResp структура, возвращаемая методом DatabaseGetFaculties.
type DatabaseGetFacultiesResp struct {
	Count int `json:"count"`
	Items []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"items"`
}

// DatabaseGetFaculties возвращает список факультетов.
func (api *API) DatabaseGetFaculties(p DatabaseGetFacultiesParams) (*DatabaseGetFacultiesResp, error) {
	resp, err := api.Request("database.getFaculties", p, new(DatabaseGetFacultiesResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetFacultiesResp), nil
}

// ==========================
//  DatabaseGetMetroStations
// ==========================

// DatabaseGetMetroStationsParams параметры метода DatabaseGetMetroStations.
type DatabaseGetMetroStationsParams struct {
	CityID   uint
	Offset   uint
	Count    uint
	Extended bool
}

// DatabaseGetMetroStationsResp структура, возвращаемая методом DatabaseGetMetroStations.
type DatabaseGetMetroStationsResp struct {
	Count int `json:"count"`
	Items []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"items"`
}

// DatabaseGetMetroStations возвращает список станций метро.
func (api *API) DatabaseGetMetroStations(p DatabaseGetMetroStationsParams) (*DatabaseGetMetroStationsResp, error) {
	resp, err := api.Request("database.getMetroStations", p, new(DatabaseGetMetroStationsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetMetroStationsResp), nil
}

// ==============================
//  DatabaseGetMetroStationsByID
// ==============================

// DatabaseGetMetroStationsByIDParams параметры метода DatabaseGetMetroStationsByID.
type DatabaseGetMetroStationsByIDParams struct {
	StationIDS []int
}

// DatabaseGetMetroStationsByIDResp структура, возвращаемая методом DatabaseGetMetroStationsByID.
type DatabaseGetMetroStationsByIDResp []struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	CityID int    `json:"city_id"`
}

// DatabaseGetMetroStationsByID возвращает информацию об одной или нескольких станциях метро по их идентификаторам.
func (api *API) DatabaseGetMetroStationsByID(p DatabaseGetMetroStationsByIDParams) (*DatabaseGetMetroStationsByIDResp, error) {
	resp, err := api.Request("database.getMetroStationsById", p, new(DatabaseGetMetroStationsByIDResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetMetroStationsByIDResp), nil
}

// ====================
//  DatabaseGetRegions
// ====================

// DatabaseGetRegionsParams параметры метода DatabaseGetRegions.
type DatabaseGetRegionsParams struct {
	CountryID uint
	Q         string
	Offset    uint
	Count     uint
}

// DatabaseGetRegionsResp структура, возвращаемая методом DatabaseGetRegions.
type DatabaseGetRegionsResp struct {
	Count int `json:"count"`
	Items []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"items"`
}

// DatabaseGetRegions возвращает список регионов.
func (api *API) DatabaseGetRegions(p DatabaseGetRegionsParams) (*DatabaseGetRegionsResp, error) {
	resp, err := api.Request("database.getRegions", p, new(DatabaseGetRegionsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetRegionsResp), nil
}

// ==========================
//  DatabaseGetSchoolClasses
// ==========================

// DatabaseGetSchoolClassesParams параметры метода DatabaseGetSchoolClasses.
type DatabaseGetSchoolClassesParams struct {
	CountryID uint
}

// DatabaseGetSchoolClasses возвращает список классов, характерных для школ определенной страны.
func (api *API) DatabaseGetSchoolClasses(p DatabaseGetSchoolClassesParams) (map[int]string, error) {
	resp, err := api.Request("database.getSchoolClasses", p, new(interface{}))
	if err != nil {
		return nil, err
	}

	temp := resp.([]interface{})
	m := make(map[int]string, len(temp))
	for _, v := range temp {
		key := int((v.([]interface{})[0]).(float64))
		t := v.([]interface{})[1]
		var value string
		switch r := t.(type) {
		case string:
			value = r
		case float64:
			value = strconv.FormatFloat(r, 'f', -1, 64)
		}
		m[key] = value
	}

	return m, nil
}

// ====================
//  DatabaseGetSchools
// ====================

// DatabaseGetSchoolsParams параметры метода DatabaseGetSchools.
type DatabaseGetSchoolsParams struct {
	Q      string
	CityID uint
	Offset uint
	Count  uint
}

// DatabaseGetSchoolsResp структура, возвращаемая методом DatabaseGetSchools.
type DatabaseGetSchoolsResp struct {
	Count int `json:"count"`
	Items []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"items"`
}

// DatabaseGetSchools возвращает список школ.
func (api *API) DatabaseGetSchools(p DatabaseGetSchoolsParams) (*DatabaseGetSchoolsResp, error) {
	resp, err := api.Request("database.getSchools", p, new(DatabaseGetSchoolsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetSchoolsResp), nil
}

// =========================
//  DatabaseGetUniversities
// =========================

// DatabaseGetUniversitiesParams параметры метода DatabaseGetUniversities.
type DatabaseGetUniversitiesParams struct {
	Q         string
	CountryID uint
	CityID    uint
	Offset    uint
	Count     uint
}

// DatabaseGetUniversitiesResp структура, возвращаемая методом DatabaseGetUniversities.
type DatabaseGetUniversitiesResp struct {
	Count int `json:"count"`
	Items []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"items"`
}

// DatabaseGetUniversities возвращает список высших учебных заведений.
func (api *API) DatabaseGetUniversities(p DatabaseGetUniversitiesParams) (*DatabaseGetUniversitiesResp, error) {
	resp, err := api.Request("database.getUniversities", p, new(DatabaseGetUniversitiesResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DatabaseGetUniversitiesResp), nil
}
