package vkapi

// ===========
//  StatusGet
// ===========

// StatusGetParams параметры метода StatusGet.
type StatusGetParams struct {
	UserID  int
	GroupID uint
}

// StatusGetResp структура, возвращаемая методом StatusGet.
type StatusGetResp struct {
	Text string `json:"text"`
	Audio *Audio  `json:"audio"`
}

// StatusGet получает текст статуса пользователя или сообщества.
func (api *API) StatusGet(p StatusGetParams) (*StatusGetResp, error) {
	resp, err := api.Request("status.get", p, new(StatusGetResp))
	if err != nil {
		return nil, err
	}
	return resp.(*StatusGetResp), nil
}

// ===========
//  StatusSet
// ===========

// StatusSetParams параметры метода StatusSet.
type StatusSetParams struct {
	Text    string
	GroupID uint
}

// StatusSet устанавливает новый статус текущему пользователю или сообществу.
func (api *API) StatusSet(p StatusSetParams) (bool, error) {
	resp, err := api.Request("status.set", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}
