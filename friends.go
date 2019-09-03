package vkapi

// ============
//  FriendsAdd
// ============

// FriendsAddParams параметры метода FriendsAdd.
type FriendsAddParams struct {
	UserID uint
	Text   string
	Follow bool
}

// FriendsAdd одобряет или создает заявку на добавление в друзья. Если идентификатор выбранного пользователя присутствует в списке заявок на добавление в друзья, полученном методом friends.getrequests, то одобряет заявку на добавление и добавляет выбранного пользователя в друзья к текущему пользователю. В противном случае создает заявку на добавление в друзья текущего пользователя к выбранному пользователю.
func (api *API) FriendsAdd(p FriendsAddParams) (int, error) {
	resp, err := api.Request("friends.add", p, new(int))
	if err != nil {
		return 0, err
	}
	return resp.(int), nil
}

// ================
//  FriendsAddList
// ================

// FriendsAddListParams параметры метода FriendsAddList.
type FriendsAddListParams struct {
	Name    string
	UserIDS []int
}

// FriendsAddListResp структура, возвращаемая методом FriendsAddList.
type FriendsAddListResp struct {
	ListID int `json:"list_id"`
}

// FriendsAddList создает новый список друзей у текущего пользователя.
func (api *API) FriendsAddList(p FriendsAddListParams) (*FriendsAddListResp, error) {
	resp, err := api.Request("friends.addList", p, new(FriendsAddListResp))
	if err != nil {
		return nil, err
	}
	return resp.(*FriendsAddListResp), nil
}

// ===================
//  FriendsAreFriends
// ===================

// FriendsAreFriendsParams параметры метода FriendsAreFriends.
type FriendsAreFriendsParams struct {
	UserIDS  []int
	NeedSign bool
}

// FriendsAreFriendsResp структура, возвращаемая методом FriendsAreFriends.
type FriendsAreFriendsResp []struct {
	UserID         int    `json:"user_id"`
	FriendStatus   int    `json:"friend_status"`
	RequestMessage string `json:"request_message"`
	ReadState      int    `json:"read_state"`
	Sign           string `json:"sign"`
}

// FriendsAreFriends возвращает информацию о том, добавлен ли текущий пользователь в друзья у указанных пользователей. Также возвращает информацию о наличии исходящей или входящей заявки в друзья (подписки).
func (api *API) FriendsAreFriends(p FriendsAreFriendsParams) (*FriendsAreFriendsResp, error) {
	resp, err := api.Request("friends.areFriends", p, new(FriendsAreFriendsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*FriendsAreFriendsResp), nil
}

// ===============
//  FriendsDelete
// ===============

// FriendsDeleteParams параметры метода FriendsDelete.
type FriendsDeleteParams struct {
	UserID uint
}

// FriendsDeleteResp структура, возвращаемая методом FriendsDelete.
type FriendsDeleteResp struct {
	Success           int `json:"success"`
	FriendDeleted     int `json:"friend_deleted"`
	OutRequestDeleted int `json:"out_request_deleted"`
	InRequestDeleted  int `json:"in_request_deleted"`
	SuggestionDeleted int `json:"suggestion_deleted"`
}

// FriendsDelete удаляет пользователя из списка друзей или отклоняет заявку в друзья. Если идентификатор выбранного пользователя присутствует в списке заявок на добавление в друзья, полученном методом friends.getrequests, то отклоняет заявку на добавление в друзья к текущему пользователю. В противном случае удаляет выбранного пользователя из списка друзей текущего пользователя, который может быть получен методом friends.get.
func (api *API) FriendsDelete(p FriendsDeleteParams) (*FriendsDeleteResp, error) {
	resp, err := api.Request("friends.delete", p, new(FriendsDeleteResp))
	if err != nil {
		return nil, err
	}
	return resp.(*FriendsDeleteResp), nil
}

// ==========================
//  FriendsDeleteAllRequests
// ==========================

// FriendsDeleteAllRequests отмечает все входящие заявки на добавление в друзья как просмотренные.
func (api *API) FriendsDeleteAllRequests() (bool, error) {
	resp, err := api.Request("friends.deleteAllRequests", struct{}{}, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===================
//  FriendsDeleteList
// ===================

// FriendsDeleteListParams параметры метода FriendsDeleteList.
type FriendsDeleteListParams struct {
	ListID uint
}

// FriendsDeleteList удаляет существующий список друзей текущего пользователя.
func (api *API) FriendsDeleteList(p FriendsDeleteListParams) (bool, error) {
	resp, err := api.Request("friends.deleteList", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =============
//  FriendsEdit
// =============

// FriendsEditParams параметры метода FriendsEdit.
type FriendsEditParams struct {
	UserID  uint
	ListIDS []int
}

// FriendsEdit редактирует списки друзей для выбранного друга.
func (api *API) FriendsEdit(p FriendsEditParams) (bool, error) {
	resp, err := api.Request("friends.edit", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =================
//  FriendsEditList
// =================

// FriendsEditListParams параметры метода FriendsEditList.
type FriendsEditListParams struct {
	Name          string
	ListID        uint
	UserIDS       []int
	AddUserIDS    []int
	DeleteUserIDS []int
}

// FriendsEditList редактирует существующий список друзей текущего пользователя.
func (api *API) FriendsEditList(p FriendsEditListParams) (bool, error) {
	resp, err := api.Request("friends.editList", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ============
//  FriendsGet
// ============

// FriendsGetParams параметры метода FriendsGet.
type FriendsGetParams struct {
	UserID   int
	Order    string
	ListID   uint
	Count    uint
	Offset   uint
	Fields   string
	NameCase string
	Ref      string
}

// FriendsGetIDSResp структура, возвращаемая методом FriendsGet, содержит идентификаторы друзей.
type FriendsGetIDSResp struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// FriendsGetUsersResp структура, возвращаемая методом FriendsGet, содержит объекты-профили.
type FriendsGetUsersResp struct {
	Count int    `json:"count"`
	Items []User `json:"items"`
}

// FriendsGet возвращает список идентификаторов друзей пользователя или расширенную информацию о друзьях пользователя (при использовании параметра fields). Если вы используете социальный граф пользователя вконтакте в своем приложении, обратите внимание на п. 4.4. правил платформы. https://vk.com/dev/rules Возвращает *FriendsGetUsersResp, если задан параметр Fields или *FriendsGetIDSResp в остальных случаях.
func (api *API) FriendsGet(p FriendsGetParams) (interface{}, error) {
	var holder interface{}
	switch len(p.Fields) > 0 {
	case true:
		holder = new(FriendsGetUsersResp)
	default:
		holder = new(FriendsGetIDSResp)
	}

	resp, err := api.Request("friends.get", p, holder)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ====================
//  FriendsGetAppUsers
// ====================

// FriendsGetAppUsers возвращает список идентификаторов друзей текущего пользователя, которые установили данное приложение.
func (api *API) FriendsGetAppUsers() ([]int, error) {
	resp, err := api.Request("friends.getAppUsers", struct{}{}, new([]int))
	if err != nil {
		return nil, err
	}
	return resp.([]int), nil
}

// ====================
//  FriendsGetByPhones
// ====================

// FriendsGetByPhonesParams параметры метода FriendsGetByPhones.
type FriendsGetByPhonesParams struct {
	Phones string
	Fields string
}

// FriendsGetByPhones возвращает список друзей пользователя, у которых завалидированные или указанные в профиле телефонные номера входят в заданный список. Использование данного метода возможно только если у текущего пользователя завалидирован номер мобильного телефона. Для проверки этого условия можно использовать метод users.get c параметрами user_ids=api_user и fields=has_mobile, где api_user равен идентификатору текущего пользователя. Для доступа к этому методу приложение должно быть доверенным.
func (api *API) FriendsGetByPhones(p FriendsGetByPhonesParams) ([]User, error) {
	resp, err := api.Request("friends.getByPhones", p, new([]User))
	if err != nil {
		return nil, err
	}
	return resp.([]User), nil
}

// =================
//  FriendsGetLists
// =================

// FriendsGetListsParams параметры метода FriendsGetLists.
type FriendsGetListsParams struct {
	UserID       uint
	ReturnSystem bool
}

// FriendsGetListsResp структура, возвращаемая методом FriendsGetLists.
type FriendsGetListsResp struct {
	Count int `json:"count"`
	Items []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
}

// FriendsGetLists возвращает список меток друзей пользователя.
func (api *API) FriendsGetLists(p FriendsGetListsParams) (*FriendsGetListsResp, error) {
	resp, err := api.Request("friends.getLists", p, new(FriendsGetListsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*FriendsGetListsResp), nil
}

// ==================
//  FriendsGetMutual
// ==================

// FriendsGetMutualParams параметры метода FriendsGetMutual.
type FriendsGetMutualParams struct {
	SourceUID  uint
	TargetUID  uint
	TargetUIDS []int
	Order      string
	Count      uint
	Offset     uint
}

// FriendsGetMutual возвращает список идентификаторов общих друзей между парой пользователей.
func (api *API) FriendsGetMutual(p FriendsGetMutualParams) ([]int, error) {
	resp, err := api.Request("friends.getMutual", p, new([]int))
	if err != nil {
		return nil, err
	}
	return resp.([]int), nil
}

// ==================
//  FriendsGetOnline
// ==================

// FriendsGetOnlineParams параметры метода FriendsGetOnline.
type FriendsGetOnlineParams struct {
	UserID       uint
	ListID       uint
	OnlineMobile bool
	Order        string
	Count        uint
	Offset       uint
}

// FriendsGetOnlineResp структура, возвращаемая методом FriendsGetOnline.
type FriendsGetOnlineResp struct {
	Online       []int `json:"online,omitempty"`
	OnlineMobile []int `json:"online_mobile,omitempty"`
}

// FriendsGetOnline возвращает список идентификаторов друзей пользователя, находящихся на сайте. Возвращает []int, если OnlineMobile=false или *FriendsGetOnlineResp, если OnlineMobile=true.
func (api *API) FriendsGetOnline(p FriendsGetOnlineParams) (interface{}, error) {
	var holder interface{}
	switch p.OnlineMobile {
	case true:
		holder = new(FriendsGetOnlineResp)
	default:
		holder = new([]int)
	}

	resp, err := api.Request("friends.getOnline", p, holder)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ==================
//  FriendsGetRecent
// ==================

// FriendsGetRecentParams параметры метода FriendsGetRecent.
type FriendsGetRecentParams struct {
	Count uint
}

// FriendsGetRecent возвращает список идентификаторов недавно добавленных друзей текущего пользователя.
func (api *API) FriendsGetRecent(p FriendsGetRecentParams) ([]int, error) {
	resp, err := api.Request("friends.getRecent", p, new([]int))
	if err != nil {
		return nil, err
	}
	return resp.([]int), nil
}

// ====================
//  FriendsGetRequests
// ====================

// FriendsGetRequestsParams параметры метода FriendsGetRequests.
type FriendsGetRequestsParams struct {
	Offset     uint
	Count      uint
	Extended   bool
	NeedMutual bool
	Out        bool
	Sort       uint
	NeedViewed bool
	Suggested  bool
	Ref        string
	Fields     string
}

// FriendsGetRequestsResp структура, возвращаемая методом FriendsGetRequests.
type FriendsGetRequestsResp struct {
	Count int `json:"count"`
	Items []struct {
		UserID          int    `json:"user_id"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		IsClosed        bool   `json:"is_closed"`
		CanAccessClosed bool   `json:"can_access_closed"`
		Mutual          *struct {
			Count int   `json:"count"`
			Users []int `json:"users"`
		} `json:"mutual"`
		TrackCode string `json:"track_code"`
	} `json:"items"`
}

// FriendsGetRequests возвращает информацию о полученных или отправленных заявках на добавление в друзья для текущего пользователя.
func (api *API) FriendsGetRequests(p FriendsGetRequestsParams) (*FriendsGetRequestsResp, error) {
	resp, err := api.Request("friends.getRequests", p, new(FriendsGetRequestsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*FriendsGetRequestsResp), nil
}

// =======================
//  FriendsGetSuggestions
// =======================

// FriendsGetSuggestionsParams параметры метода FriendsGetSuggestions.
type FriendsGetSuggestionsParams struct {
	Filter   string
	Count    uint
	Offset   uint
	Fields   string
	NameCase string
}

// FriendsGetSuggestionsResp структура, возвращаемая методом FriendsGetSuggestions.
type FriendsGetSuggestionsResp struct {
	Count int    `json:"count"`
	Items []User `json:"items"`
}

// FriendsGetSuggestions возвращает список профилей пользователей, которые могут быть друзьями текущего пользователя.
func (api *API) FriendsGetSuggestions(p FriendsGetSuggestionsParams) (*FriendsGetSuggestionsResp, error) {
	resp, err := api.Request("friends.getSuggestions", p, new(FriendsGetSuggestionsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*FriendsGetSuggestionsResp), nil
}

// ===============
//  FriendsSearch
// ===============

// FriendsSearchParams параметры метода FriendsSearch.
type FriendsSearchParams struct {
	UserID   uint
	Q        string
	Fields   string
	NameCase string
	Offset   uint
	Count    uint
}

// FriendsSearchResp структура, возвращаемая методом FriendsSearch.
type FriendsSearchResp struct {
	Count int    `json:"count"`
	Items []User `json:"items"`
}

// FriendsSearch позволяет искать по списку друзей пользователей. Для расширенного поиска по списку друзей можно использовать метод users.search с параметром from_list=friends.
func (api *API) FriendsSearch(p FriendsSearchParams) (*FriendsSearchResp, error) {
	resp, err := api.Request("friends.search", p, new(FriendsSearchResp))
	if err != nil {
		return nil, err
	}
	return resp.(*FriendsSearchResp), nil
}
