package vkapi

// ================
//  SearchGetHints
// ================

// SearchGetHintsParams параметры метода SearchGetHints.
type SearchGetHintsParams struct {
	Q            string
	Offset       uint
	Limit        uint
	Filters      string
	Fields       string
	SearchGlobal bool
}

// SearchGetHintsResp структура, возвращаемая методом SearchGetHints.
type SearchGetHintsResp struct {
	Count int `json:"count"`
	Items []struct {
		Type  string
		Group *struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			ScreenName  string `json:"screen_name"`
			IsClosed    int    `json:"is_closed"`
			IsAdmin     int    `json:"is_admin"`
			IsMember    int    `json:"is_member"`
			Type        string `json:"type"`
			Photo       string `json:"photo"`
			PhotoMedium string `json:"photo_medium"`
			PhotoBig    string `json:"photo_big"`
		} `json:"group"`
		Profile *struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"profile"`
		Section     string `json:"section"`
		Description string `json:"description"`
		Global      int    `json:"global"`
	} `json:"items"`
}

// SearchGetHints метод позволяет получить результаты быстрого поиска по произвольной подстроке.
func (api *API) SearchGetHints(p SearchGetHintsParams) (*SearchGetHintsResp, error) {
	resp, err := api.Request("search.getHints", p, new(SearchGetHintsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*SearchGetHintsResp), nil
}
