package kakaorestapi

import (
	"encoding/json"
	"errors"
)

// SearchAddressResponse response struct.
type SearchAddressResponse struct {
	Meta struct {
		TotalCount    int  `json:"total_count"`
		PageableCount int  `json:"pageable_count"`
		IsEnd         bool `json:"is_end"`
	} `json:"meta"`
	Documents []struct {
		AddressName string `json:"address_name"`
		Y           string `json:"y"`
		X           string `json:"x"`
		AddressType string `json:"address_type"`
		Address     struct {
			AddressName       string `json:"address_name"`
			Region1DepthName  string `json:"region_1depth_name"`
			Region2DepthName  string `json:"region_2depth_name"`
			Region3DepthName  string `json:"region_3depth_name"`
			Region3DepthHName string `json:"region_3depth_h_name"`
			HCode             string `json:"h_code"`
			BCode             string `json:"b_code"`
			MountainYn        string `json:"mountain_yn"`
			MainAddressNo     string `json:"main_address_no"`
			SubAddressNo      string `json:"sub_address_no"`
			ZipCode           string `json:"zip_code"`
			X                 string `json:"x"`
			Y                 string `json:"y"`
		} `json:"address"`
		RoadAddress struct {
			AddressName      string `json:"address_name"`
			Region1DepthName string `json:"region_1depth_name"`
			Region2DepthName string `json:"region_2depth_name"`
			Region3DepthName string `json:"region_3depth_name"`
			RoadName         string `json:"road_name"`
			UndergroundYn    string `json:"underground_yn"`
			MainBuildingNo   string `json:"main_building_no"`
			SubBuildingNo    string `json:"sub_building_no"`
			BuildingName     string `json:"building_name"`
			ZoneNo           string `json:"zone_no"`
			Y                string `json:"y"`
			X                string `json:"x"`
		} `json:"road_address"`
	} `json:"documents"`
}

// SearchKeywordResponse reponse struct.
type SearchKeywordResponse struct {
	Meta struct {
		SameName struct {
			Region         []interface{} `json:"region"`
			Keyword        string        `json:"keyword"`
			SelectedRegion string        `json:"selected_region"`
		} `json:"same_name"`
		PageableCount int  `json:"pageable_count"`
		TotalCount    int  `json:"total_count"`
		IsEnd         bool `json:"is_end"`
	} `json:"meta"`
	Documents []struct {
		PlaceName         string `json:"place_name"`
		Distance          string `json:"distance"`
		PlaceURL          string `json:"place_url"`
		CategoryName      string `json:"category_name"`
		AddressName       string `json:"address_name"`
		RoadAddressName   string `json:"road_address_name"`
		ID                string `json:"id"`
		Phone             string `json:"phone"`
		CategoryGroupCode string `json:"category_group_code"`
		CategoryGroupName string `json:"category_group_name"`
		X                 string `json:"x"`
		Y                 string `json:"y"`
	} `json:"documents"`
}

// Coordinates2RegionCodeResponse reponse struct.
type Coordinates2RegionCodeResponse struct {
	Meta struct {
		TotalCount int `json:"total_count"`
	} `json:"meta"`
	Documents []struct {
		RegionType       string  `json:"region_type"`
		AddressName      string  `json:"address_name"`
		Region1DepthName string  `json:"region_1depth_name"`
		Region2DepthName string  `json:"region_2depth_name"`
		Region3DepthName string  `json:"region_3depth_name"`
		Region4DepthName string  `json:"region_4depth_name"`
		Code             string  `json:"code"`
		X                float64 `json:"x"`
		Y                float64 `json:"y"`
	} `json:"documents"`
}

// SearchAddressParams used for passing params values to search address function.
type SearchAddressParams struct {
	Query string // query string you want to search (requred)
	Page  int    // page number (optional)
	Size  int    // number of documents to show on one page (optional)
}

// SearchAddress returns coordinates information for passed address.
func (c *Client) SearchAddress(params SearchAddressParams) (*SearchAddressResponse, error) {
	values := map[string]interface{}{
		"query": params.Query,
		"page":  params.Page,
		"size":  params.Size,
	}

	response := &SearchAddressResponse{}

	bytes, err := c.get("/v2/local/search/address.json", values)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(bytes, response)

	return response, err
}

// SearchKeywordParams used for passing params values to search keywords function.
type SearchKeywordParams struct {
	Query             string
	CategoryGroupCode string
	X                 string
	Y                 string
	Radius            int
	Rect              string
	Page              int
	Size              int
	Sort              string
}

// SearchKeyword returns search result that match the query terms based on passed sorting criteria.
func (c *Client) SearchKeyword(params SearchKeywordParams) (*SearchKeywordResponse, error) {
	response := &SearchKeywordResponse{}

	values := make(map[string]interface{})

	if params.Query == "" {
		return response, errors.New("search query is required")
	}

	values["query"] = params.Query

	if params.CategoryGroupCode != "" {
		values["category_group_code"] = params.CategoryGroupCode
	}

	if params.X != "" {
		values["x"] = params.X
	}

	if params.Y != "" {
		values["y"] = params.Y
	}

	if params.Radius > 0 {
		values["radius"] = params.Radius
	}

	if params.Page > 0 {
		values["page"] = params.Page
	}

	if params.Size > 0 {
		values["size"] = params.Size
	}

	if params.Sort != "" {
		values["sort"] = params.Sort
	}

	bytes, err := c.get("/v2/local/search/keyword.json", values)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(bytes, response)

	return response, err
}

// Coordinates2RegionCodeParams used for passing params values to convert coordinates to region code.
type Coordinates2RegionCodeParams struct {
	X           string
	Y           string
	InputCoord  string
	OutputCoord string
	Lang        string
}

// Coordinates2RegionCode converts passed coordinates to approximate local information.
func (c *Client) Coordinates2RegionCode(params Coordinates2RegionCodeParams) (*Coordinates2RegionCodeResponse, error) {
	response := &Coordinates2RegionCodeResponse{}

	values := make(map[string]interface{})

	if params.X == "" || params.X == "0" {
		return response, errors.New("x params is required")
	}

	if params.Y == "" || params.Y == "0" {
		return response, errors.New("y params is required")
	}

	values["x"] = params.X
	values["y"] = params.Y

	if params.InputCoord != "" {
		values["input_coord"] = params.InputCoord
	}

	if params.OutputCoord != "" {
		values["output_coord"] = params.OutputCoord
	}

	if params.Lang != "" {
		values["lang"] = params.Lang
	}

	bytes, err := c.get("/v2/local/geo/coord2regioncode.json", values)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(bytes, response)

	return response, err
}
