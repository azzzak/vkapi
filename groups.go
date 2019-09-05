package vkapi

// ==================
//  GroupsAddAddress
// ==================

// GroupsAddAddressParams параметры метода GroupsAddAddress.
type GroupsAddAddressParams struct {
	GroupID           uint
	Title             string
	Address           string
	AdditionalAddress string
	CountryID         uint
	CityID            uint
	MetroID           uint
	Latitude          float32
	Longitude         float32
	Phone             string
	WorkInfoStatus    string
	Timetable         *Timetable
	IsMainAddress     bool
}

// GroupsAddAddress позволяет добавить адрес в сообщество. Список адресов может быть получен методом groups.getaddresses. Для того, чтобы воспользоваться этим методом, вы должны быть администратором сообщества.
func (api *API) GroupsAddAddress(p GroupsAddAddressParams) (*Address, error) {
	if p.Timetable != nil {
		p.WorkInfoStatus = "timetable"
	}

	resp, err := api.Request("groups.addAddress", p, new(Address))
	if err != nil {
		return nil, err
	}
	return resp.(*Address), nil
}

// =========================
//  GroupsAddCallbackServer
// =========================

// GroupsAddCallbackServerParams параметры метода GroupsAddCallbackServer.
type GroupsAddCallbackServerParams struct {
	GroupID   uint
	URL       string
	Title     string
	SecretKey string
}

// GroupsAddCallbackServerResp структура, возвращаемая методом GroupsAddCallbackServer.
type GroupsAddCallbackServerResp struct {
	ServerID int `json:"server_id"`
}

// GroupsAddCallbackServer добавляет сервер для callback api в сообщество.
func (api *API) GroupsAddCallbackServer(p GroupsAddCallbackServerParams) (*GroupsAddCallbackServerResp, error) {
	resp, err := api.Request("groups.addCallbackServer", p, new(GroupsAddCallbackServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsAddCallbackServerResp), nil
}

// ===============
//  GroupsAddLink
// ===============

// GroupsAddLinkParams параметры метода GroupsAddLink.
type GroupsAddLinkParams struct {
	GroupID uint
	Link    string
	Text    string
}

// GroupsAddLinkResp структура, возвращаемая методом GroupsAddLink.
type GroupsAddLinkResp struct {
	ID              int    `json:"id"`
	URL             string `json:"url"`
	Name            string `json:"name"`
	EditTitle       int    `json:"edit_title"`
	Desc            string `json:"desc"`
	ImageProcessing int    `json:"image_processing"`
}

// GroupsAddLink позволяет добавлять ссылки в сообщество. Список ссылок может быть получен методом groups.getbyid с параметром fields=links. Для того, чтобы воспользоваться этим методом, вы должны быть администратором группы.
func (api *API) GroupsAddLink(p GroupsAddLinkParams) (*GroupsAddLinkResp, error) {
	resp, err := api.Request("groups.addLink", p, new(GroupsAddLinkResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsAddLinkResp), nil
}

// ======================
//  GroupsApproveRequest
// ======================

// GroupsApproveRequestParams параметры метода GroupsApproveRequest.
type GroupsApproveRequestParams struct {
	GroupID uint
	UserID  uint
}

// GroupsApproveRequest позволяет одобрить заявку в группу от пользователя. Для того, чтобы воспользоваться этим методом вы должны быть администратором группы.
func (api *API) GroupsApproveRequest(p GroupsApproveRequestParams) (bool, error) {
	resp, err := api.Request("groups.approveRequest", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===========
//  GroupsBan
// ===========

// GroupsBanParams параметры метода GroupsBan.
type GroupsBanParams struct {
	GroupID        uint
	OwnerID        int
	EndDate        int64
	Reason         uint
	Comment        string
	CommentVisible bool
}

// GroupsBan добавляет пользователя или группу в черный список сообщества.
func (api *API) GroupsBan(p GroupsBanParams) (bool, error) {
	resp, err := api.Request("groups.ban", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ==============
//  GroupsCreate
// ==============

// GroupsCreateParams параметры метода GroupsCreate.
type GroupsCreateParams struct {
	Title          string
	Description    string
	Type           string
	PublicCategory uint
	Subtype        uint
}

// GroupsCreate создает новое сообщество.
func (api *API) GroupsCreate(p GroupsCreateParams) (*Group, error) {
	resp, err := api.Request("groups.create", p, new(Group))
	if err != nil {
		return nil, err
	}
	return resp.(*Group), nil
}

// =====================
//  GroupsDeleteAddress
// =====================

// GroupsDeleteAddressParams параметры метода GroupsDeleteAddress.
type GroupsDeleteAddressParams struct {
	GroupID   uint
	AddressID uint
}

// GroupsDeleteAddress удаляет адрес сообщества.
func (api *API) GroupsDeleteAddress(p GroupsDeleteAddressParams) (bool, error) {
	resp, err := api.Request("groups.deleteAddress", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ============================
//  GroupsDeleteCallbackServer
// ============================

// GroupsDeleteCallbackServerParams параметры метода GroupsDeleteCallbackServer.
type GroupsDeleteCallbackServerParams struct {
	GroupID  uint
	ServerID uint
}

// GroupsDeleteCallbackServer удаляет сервер для callback api из сообщества.
func (api *API) GroupsDeleteCallbackServer(p GroupsDeleteCallbackServerParams) (bool, error) {
	resp, err := api.Request("groups.deleteCallbackServer", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ==================
//  GroupsDeleteLink
// ==================

// GroupsDeleteLinkParams параметры метода GroupsDeleteLink.
type GroupsDeleteLinkParams struct {
	GroupID uint
	LinkID  uint
}

// GroupsDeleteLink позволяет удалить ссылки из сообщества. Список ссылок может быть получен методом groups.getbyid, с параметром fields=links. Для того, чтобы воспользоваться этим методом вы должны быть администратором группы.
func (api *API) GroupsDeleteLink(p GroupsDeleteLinkParams) (bool, error) {
	resp, err := api.Request("groups.deleteLink", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =====================
//  GroupsDisableOnline
// =====================

// GroupsDisableOnlineParams параметры метода GroupsDisableOnline.
type GroupsDisableOnlineParams struct {
	GroupID uint
}

// GroupsDisableOnline выключает статус "онлайн" в сообществе.
func (api *API) GroupsDisableOnline(p GroupsDisableOnlineParams) (bool, error) {
	resp, err := api.Request("groups.disableOnline", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ============
//  GroupsEdit
// ============

// GroupsEditParams параметры метода GroupsEdit.
type GroupsEditParams struct {
	GroupID           uint
	Title             string
	Description       string
	ScreenName        string
	Access            uint
	Website           string
	Subject           uint
	Email             string
	Phone             string
	Rss               string
	EventStartDate    uint
	EventFinishDate   uint
	EventGroupID      uint
	PublicCategory    uint
	PublicSubcategory uint
	PublicDate        string
	Wall              uint
	Topics            uint
	Photos            uint
	Video             uint
	Audio             uint
	Links             bool
	Events            bool
	Places            bool
	Contacts          bool
	Docs              uint
	Wiki              uint
	Messages          bool
	Articles          bool
	Addresses         bool
	AgeLimits         uint
	Market            bool
	MarketComments    bool
	MarketCountry     string
	MarketCity        string
	MarketCurrency    uint
	MarketContact     uint
	MarketWiki        uint
	ObsceneFilter     bool
	ObsceneStopwords  bool
	ObsceneWords      string
	MainSection       uint
	SecondarySection  uint
	Country           uint
	City              uint
}

// GroupsEdit редактирует сообщество. для того, чтобы воспользоваться этим методом, вы должны быть администратором сообщества. Значения настроек можно получить методом groups.getsettings.
func (api *API) GroupsEdit(p GroupsEditParams) (bool, error) {
	resp, err := api.Request("groups.edit", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===================
//  GroupsEditAddress
// ===================

// GroupsEditAddressParams параметры метода GroupsEditAddress.
type GroupsEditAddressParams struct {
	GroupID           uint
	AddressID         uint
	Title             string
	Address           string
	AdditionalAddress string
	CountryID         uint
	CityID            uint
	MetroID           uint
	Latitude          float32
	Longitude         float32
	Phone             string
	WorkInfoStatus    string
	Timetable         *Timetable
	IsMainAddress     bool
}

// GroupsEditAddress позволяет отредактировать адрес в сообществе. Список адресов может быть получен методом groups.getaddresses. Для того, чтобы воспользоваться этим методом, вы должны быть администратором сообщества.
func (api *API) GroupsEditAddress(p GroupsEditAddressParams) (*Address, error) {
	if p.Timetable != nil {
		p.WorkInfoStatus = "timetable"
	}

	resp, err := api.Request("groups.editAddress", p, new(Address))
	if err != nil {
		return nil, err
	}
	return resp.(*Address), nil
}

// ==========================
//  GroupsEditCallbackServer
// ==========================

// GroupsEditCallbackServerParams параметры метода GroupsEditCallbackServer.
type GroupsEditCallbackServerParams struct {
	GroupID   uint
	ServerID  uint
	URL       string
	Title     string
	SecretKey string
}

// GroupsEditCallbackServer редактирует данные сервера для callback api в сообществе.
func (api *API) GroupsEditCallbackServer(p GroupsEditCallbackServerParams) (bool, error) {
	resp, err := api.Request("groups.editCallbackServer", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ================
//  GroupsEditLink
// ================

// GroupsEditLinkParams параметры метода GroupsEditLink.
type GroupsEditLinkParams struct {
	GroupID uint
	LinkID  uint
	Text    string
}

// GroupsEditLink позволяет редактировать ссылки в сообществе.
func (api *API) GroupsEditLink(p GroupsEditLinkParams) (bool, error) {
	resp, err := api.Request("groups.editLink", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===================
//  GroupsEditManager
// ===================

// GroupsEditManagerParams параметры метода GroupsEditManager.
type GroupsEditManagerParams struct {
	GroupID         uint
	UserID          uint
	Role            string
	IsContact       bool
	ContactPosition string
	ContactPhone    string
	ContactEmail    string
}

// GroupsEditManager позволяет назначить/разжаловать руководителя в сообществе или изменить уровень его полномочий. Для того, чтобы воспользоваться этим методом, вы должны быть администратором группы.
func (api *API) GroupsEditManager(p GroupsEditManagerParams) (bool, error) {
	resp, err := api.Request("groups.editManager", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ====================
//  GroupsEnableOnline
// ====================

// GroupsEnableOnlineParams параметры метода GroupsEnableOnline.
type GroupsEnableOnlineParams struct {
	GroupID uint
}

// GroupsEnableOnline включает статус "онлайн" в сообществе.
func (api *API) GroupsEnableOnline(p GroupsEnableOnlineParams) (bool, error) {
	resp, err := api.Request("groups.enableOnline", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===========
//  GroupsGet
// ===========

// GroupsGetParams параметры метода GroupsGet.
type GroupsGetParams struct {
	UserID   uint
	Extended bool
	Filter   string
	Fields   string
	Offset   uint
	Count    uint
}

// GroupsGetIDSResp структура, возвращаемая методом GroupsGet, содержит идентификаторы сообществ.
type GroupsGetIDSResp struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// GroupsGetGroupsResp структура, возвращаемая методом GroupsGet, содержит объекты-сообщества.
type GroupsGetGroupsResp struct {
	Count int     `json:"count"`
	Items []Group `json:"items"`
}

// GroupsGet возвращает список сообществ указанного пользователя. Возвращает *GroupsGetIDSResp, если Extended=false или *GroupsGetGroupsResp, если Extended=true.
func (api *API) GroupsGet(p GroupsGetParams) (interface{}, error) {
	var holder interface{}
	switch p.Extended {
	case true:
		holder = new(GroupsGetGroupsResp)
	default:
		holder = new(GroupsGetIDSResp)
	}

	resp, err := api.Request("groups.get", p, holder)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ====================
//  GroupsGetAddresses
// ====================

// GroupsGetAddressesParams параметры метода GroupsGetAddresses.
type GroupsGetAddressesParams struct {
	GroupID    uint
	AddressIDS []int
	Latitude   float32
	Longitude  float32
	Offset     uint
	Count      uint
	Fields     string
}

// GroupsGetAddressesResp структура, возвращаемая методом GroupsGetAddresses.
type GroupsGetAddressesResp struct {
	Count int       `json:"count"`
	Items []Address `json:"items"`
}

// GroupsGetAddresses возвращает список адресов сообщества.
func (api *API) GroupsGetAddresses(p GroupsGetAddressesParams) (*GroupsGetAddressesResp, error) {
	resp, err := api.Request("groups.getAddresses", p, new(GroupsGetAddressesResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetAddressesResp), nil
}

// =================
//  GroupsGetBanned
// =================

// GroupsGetBannedParams параметры метода GroupsGetBanned.
type GroupsGetBannedParams struct {
	GroupID uint
	Offset  uint
	Count   uint
	Fields  string
	OwnerID int
}

// GroupsGetBannedResp структура, возвращаемая методом GroupsGetBanned.
type GroupsGetBannedResp struct {
	Count int `json:"count"`
	Items []struct {
		Type    string `json:"type"`
		Group   *Group `json:"group"`
		Profile *User  `json:"profile"`
		BanInfo *struct {
			AdminID int    `json:"admin_id"`
			Date    int64  `json:"date"`
			Reason  int    `json:"reason"`
			Comment string `json:"comment"`
			EndDate int64  `json:"end_date"`
		} `json:"ban_info"`
	} `json:"items"`
}

// GroupsGetBanned возвращает список забаненных пользователей и сообществ в сообществе.
func (api *API) GroupsGetBanned(p GroupsGetBannedParams) (*GroupsGetBannedResp, error) {
	resp, err := api.Request("groups.getBanned", p, new(GroupsGetBannedResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetBannedResp), nil
}

// ===============
//  GroupsGetByID
// ===============

// GroupsGetByIDParams параметры метода GroupsGetByID.
type GroupsGetByIDParams struct {
	GroupIDS string
	GroupID  string
	Fields   string
}

// GroupsGetByID возвращает информацию о заданном сообществе или о нескольких сообществах.
func (api *API) GroupsGetByID(p GroupsGetByIDParams) ([]Group, error) {
	resp, err := api.Request("groups.getById", p, new([]Group))
	if err != nil {
		return nil, err
	}
	return resp.([]Group), nil
}

// ===================================
//  GroupsGetCallbackConfirmationCode
// ===================================

// GroupsGetCallbackConfirmationCodeParams параметры метода GroupsGetCallbackConfirmationCode.
type GroupsGetCallbackConfirmationCodeParams struct {
	GroupID uint
}

// GroupsGetCallbackConfirmationCode позволяет получить строку, необходимую для подтверждения адреса сервера в callback api. Для добавления или редактирования сервера используйте соответственно методы groups.addcallbackserver и groups.editcallbackserver.
func (api *API) GroupsGetCallbackConfirmationCode(p GroupsGetCallbackConfirmationCodeParams) (string, error) {
	resp, err := api.Request("groups.getCallbackConfirmationCode", p, new(interface{}))
	if err != nil {
		return "", err
	}

	m := resp.(map[string]interface{})
	if v, ok := m["code"]; ok {
		return v.(string), nil
	}

	return "", nil
}

// ==========================
//  GroupsGetCallbackServers
// ==========================

// GroupsGetCallbackServersParams параметры метода GroupsGetCallbackServers.
type GroupsGetCallbackServersParams struct {
	GroupID   uint
	ServerIDS []int
}

// GroupsGetCallbackServersResp структура, возвращаемая методом GroupsGetCallbackServers.
type GroupsGetCallbackServersResp struct {
	Count int `json:"count"`
	Items []struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		CreatorID int    `json:"creator_id"`
		URL       string `json:"url"`
		SecretKey string `json:"secret_key"`
		Status    string `json:"status"`
	} `json:"items"`
}

// GroupsGetCallbackServers получает информацию о серверах для callback api в сообществе.
func (api *API) GroupsGetCallbackServers(p GroupsGetCallbackServersParams) (*GroupsGetCallbackServersResp, error) {
	resp, err := api.Request("groups.getCallbackServers", p, new(GroupsGetCallbackServersResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetCallbackServersResp), nil
}

// ===========================
//  GroupsGetCallbackSettings
// ===========================

// GroupsGetCallbackSettingsParams параметры метода GroupsGetCallbackSettings.
type GroupsGetCallbackSettingsParams struct {
	GroupID  uint
	ServerID uint
}

// GroupsGetCallbackSettingsResp структура, возвращаемая методом GroupsGetCallbackSettings.
type GroupsGetCallbackSettingsResp struct {
	APIVersion string `json:"api_version"`
	Events     *struct {
		MessageNew           int `json:"message_new"`
		MessageReply         int `json:"message_reply"`
		MessageEdit          int `json:"message_edit"`
		MessageAllow         int `json:"message_allow"`
		MessageDeny          int `json:"message_deny"`
		MessageTypingState   int `json:"message_typing_state"`
		PhotoNew             int `json:"photo_new"`
		AudioNew             int `json:"audio_new"`
		VideoNew             int `json:"video_new"`
		WallReplyNew         int `json:"wall_reply_new"`
		WallReplyEdit        int `json:"wall_reply_edit"`
		WallReplyDelete      int `json:"wall_reply_delete"`
		WallReplyRestore     int `json:"wall_reply_restore"`
		WallPostNew          int `json:"wall_post_new"`
		WallRepost           int `json:"wall_repost"`
		BoardPostNew         int `json:"board_post_new"`
		BoardPostEdit        int `json:"board_post_edit"`
		BoardPostDelete      int `json:"board_post_delete"`
		BoardPostRestore     int `json:"board_post_restore"`
		PhotoCommentNew      int `json:"photo_comment_new"`
		PhotoCommentEdit     int `json:"photo_comment_edit"`
		PhotoCommentDelete   int `json:"photo_comment_delete"`
		PhotoCommentRestore  int `json:"photo_comment_restore"`
		VideoCommentNew      int `json:"video_comment_new"`
		VideoCommentEdit     int `json:"video_comment_edit"`
		VideoCommentDelete   int `json:"video_comment_delete"`
		VideoCommentRestore  int `json:"video_comment_restore"`
		MarketCommentNew     int `json:"market_comment_new"`
		MarketCommentEdit    int `json:"market_comment_edit"`
		MarketCommentDelete  int `json:"market_comment_delete"`
		MarketCommentRestore int `json:"market_comment_restore"`
		PollVoteNew          int `json:"poll_vote_new"`
		GroupJoin            int `json:"group_join"`
		GroupLeave           int `json:"group_leave"`
		UserBlock            int `json:"user_block"`
		UserUnblock          int `json:"user_unblock"`
		GroupChangeSettings  int `json:"group_change_settings"`
		GroupChangePhoto     int `json:"group_change_photo"`
		GroupOfficersEdit    int `json:"group_officers_edit"`
	} `json:"events"`
}

// GroupsGetCallbackSettings позволяет получить настройки уведомлений callback api для сообщества.
func (api *API) GroupsGetCallbackSettings(p GroupsGetCallbackSettingsParams) (*GroupsGetCallbackSettingsResp, error) {
	resp, err := api.Request("groups.getCallbackSettings", p, new(GroupsGetCallbackSettingsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetCallbackSettingsResp), nil
}

// ==================
//  GroupsGetCatalog
// ==================

// GroupsGetCatalogParams параметры метода GroupsGetCatalog.
type GroupsGetCatalogParams struct {
	CategoryID    uint
	SubcategoryID uint
}

// GroupsGetCatalogResp структура, возвращаемая методом GroupsGetCatalog.
type GroupsGetCatalogResp struct {
	Count int     `json:"count"`
	Items []Group `json:"items"`
}

// GroupsGetCatalog возвращает список сообществ выбранной категории каталога.
func (api *API) GroupsGetCatalog(p GroupsGetCatalogParams) (*GroupsGetCatalogResp, error) {
	resp, err := api.Request("groups.getCatalog", p, new(GroupsGetCatalogResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetCatalogResp), nil
}

// ======================
//  GroupsGetCatalogInfo
// ======================

// GroupsGetCatalogInfoParams параметры метода GroupsGetCatalogInfo.
type GroupsGetCatalogInfoParams struct {
	Extended      bool
	Subcategories bool
}

// GroupsGetCatalogInfoResp структура, возвращаемая методом GroupsGetCatalogInfo.
type GroupsGetCatalogInfoResp struct {
	Enabled    int `json:"enabled"`
	Categories []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Subcategories []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"subcategories"`
		PageCount    int `json:"page_count"`
		PagePreviews []struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			ScreenName   string `json:"screen_name"`
			IsClosed     int    `json:"is_closed"`
			Type         string `json:"type"`
			IsAdmin      int    `json:"is_admin"`
			IsMember     int    `json:"is_member"`
			IsAdvertiser int    `json:"is_advertiser"`
			Photo50      string `json:"photo_50"`
			Photo100     string `json:"photo_100"`
			Photo200     string `json:"photo_200"`
		} `json:"page_previews"`
	} `json:"categories"`
}

// GroupsGetCatalogInfo возвращает список категорий для каталога сообществ.
func (api *API) GroupsGetCatalogInfo(p GroupsGetCatalogInfoParams) (*GroupsGetCatalogInfoResp, error) {
	resp, err := api.Request("groups.getCatalogInfo", p, new(GroupsGetCatalogInfoResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetCatalogInfoResp), nil
}

// =======================
//  GroupsGetInvitedUsers
// =======================

// GroupsGetInvitedUsersParams параметры метода GroupsGetInvitedUsers.
type GroupsGetInvitedUsersParams struct {
	GroupID  uint
	Offset   uint
	Count    uint
	Fields   string
	NameCase string
}

// GroupsGetInvitedUsersResp структура, возвращаемая методом GroupsGetInvitedUsers.
type GroupsGetInvitedUsersResp struct {
	Count int    `json:"count"`
	Items []User `json:"items"`
}

// GroupsGetInvitedUsers возвращает список пользователей, которые были приглашены в группу.
func (api *API) GroupsGetInvitedUsers(p GroupsGetInvitedUsersParams) (*GroupsGetInvitedUsersResp, error) {
	resp, err := api.Request("groups.getInvitedUsers", p, new(GroupsGetInvitedUsersResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetInvitedUsersResp), nil
}

// ==================
//  GroupsGetInvites
// ==================

// GroupsGetInvitesParams параметры метода GroupsGetInvites.
type GroupsGetInvitesParams struct {
	Offset   uint
	Count    uint
	Extended bool
}

// GroupsGetInvitesResp структура, возвращаемая методом GroupsGetInvites.
type GroupsGetInvitesResp struct {
	Count    int     `json:"count"`
	Items    []Group `json:"items"`
	Profiles []struct {
		ID              int    `json:"id"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		IsClosed        bool   `json:"is_closed"`
		CanAccessClosed bool   `json:"can_access_closed"`
	} `json:"profiles"`
}

// GroupsGetInvites данный метод возвращает список приглашений в сообщества и встречи текущего пользователя.
func (api *API) GroupsGetInvites(p GroupsGetInvitesParams) (*GroupsGetInvitesResp, error) {
	resp, err := api.Request("groups.getInvites", p, new(GroupsGetInvitesResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetInvitesResp), nil
}

// =========================
//  GroupsGetLongPollServer
// =========================

// GroupsGetLongPollServerParams параметры метода GroupsGetLongPollServer.
type GroupsGetLongPollServerParams struct {
	GroupID uint
}

// GroupsGetLongPollServerResp структура, возвращаемая методом GroupsGetLongPollServer.
type GroupsGetLongPollServerResp struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     string `json:"ts"`
}

// GroupsGetLongPollServer возвращает данные для подключения к bots longpoll api.
func (api *API) GroupsGetLongPollServer(p GroupsGetLongPollServerParams) (*GroupsGetLongPollServerResp, error) {
	resp, err := api.Request("groups.getLongPollServer", p, new(GroupsGetLongPollServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetLongPollServerResp), nil
}

// ===========================
//  GroupsGetLongPollSettings
// ===========================

// GroupsGetLongPollSettingsParams параметры метода GroupsGetLongPollSettings.
type GroupsGetLongPollSettingsParams struct {
	GroupID uint
}

// GroupsGetLongPollSettingsResp структура, возвращаемая методом GroupsGetLongPollSettings.
type GroupsGetLongPollSettingsResp struct {
	IsEnabled  bool   `json:"is_enabled"`
	APIVersion string `json:"api_version"`
	Events     *struct {
		MessageNew           int `json:"message_new"`
		MessageReply         int `json:"message_reply"`
		MessageEdit          int `json:"message_edit"`
		MessageAllow         int `json:"message_allow"`
		MessageDeny          int `json:"message_deny"`
		MessageTypingState   int `json:"message_typing_state"`
		PhotoNew             int `json:"photo_new"`
		AudioNew             int `json:"audio_new"`
		VideoNew             int `json:"video_new"`
		WallReplyNew         int `json:"wall_reply_new"`
		WallReplyEdit        int `json:"wall_reply_edit"`
		WallReplyDelete      int `json:"wall_reply_delete"`
		WallReplyRestore     int `json:"wall_reply_restore"`
		WallPostNew          int `json:"wall_post_new"`
		WallRepost           int `json:"wall_repost"`
		BoardPostNew         int `json:"board_post_new"`
		BoardPostEdit        int `json:"board_post_edit"`
		BoardPostDelete      int `json:"board_post_delete"`
		BoardPostRestore     int `json:"board_post_restore"`
		PhotoCommentNew      int `json:"photo_comment_new"`
		PhotoCommentEdit     int `json:"photo_comment_edit"`
		PhotoCommentDelete   int `json:"photo_comment_delete"`
		PhotoCommentRestore  int `json:"photo_comment_restore"`
		VideoCommentNew      int `json:"video_comment_new"`
		VideoCommentEdit     int `json:"video_comment_edit"`
		VideoCommentDelete   int `json:"video_comment_delete"`
		VideoCommentRestore  int `json:"video_comment_restore"`
		MarketCommentNew     int `json:"market_comment_new"`
		MarketCommentEdit    int `json:"market_comment_edit"`
		MarketCommentDelete  int `json:"market_comment_delete"`
		MarketCommentRestore int `json:"market_comment_restore"`
		PollVoteNew          int `json:"poll_vote_new"`
		GroupJoin            int `json:"group_join"`
		GroupLeave           int `json:"group_leave"`
		UserBlock            int `json:"user_block"`
		UserUnblock          int `json:"user_unblock"`
		GroupChangeSettings  int `json:"group_change_settings"`
		GroupChangePhoto     int `json:"group_change_photo"`
		GroupOfficersEdit    int `json:"group_officers_edit"`
	} `json:"events"`
}

// GroupsGetLongPollSettings получает настройки bots longpoll api для сообщества.
func (api *API) GroupsGetLongPollSettings(p GroupsGetLongPollSettingsParams) (*GroupsGetLongPollSettingsResp, error) {
	resp, err := api.Request("groups.getLongPollSettings", p, new(GroupsGetLongPollSettingsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetLongPollSettingsResp), nil
}

// ==================
//  GroupsGetMembers
// ==================

// GroupsGetMembersParams параметры метода GroupsGetMembers.
type GroupsGetMembersParams struct {
	GroupID string
	Sort    string
	Offset  uint
	Count   uint
	Fields  string
	Filter  string
}

// GroupsGetMembersIDSResp структура, возвращаемая методом GroupsGetMembers, содержит идентификаторы участников сообщества.
type GroupsGetMembersIDSResp struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// GroupsGetMembersRolesResp структура, возвращаемая методом GroupsGetMembers, содержит роли руководителей сообщества.
type GroupsGetMembersRolesResp struct {
	Count int `json:"count"`
	Items []struct {
		ID          int      `json:"id"`
		Role        string   `json:"role"`
		Permissions []string `json:"permissions"`
	} `json:"items"`
}

// GroupsGetMembers возвращает список участников сообщества. Возвращает *GroupsGetMembersRolesResp, если Filter=managers или *GroupsGetMembersIDSResp в остальных случаях.
func (api *API) GroupsGetMembers(p GroupsGetMembersParams) (interface{}, error) {
	var holder interface{}
	switch p.Filter == "managers" {
	case true:
		holder = new(GroupsGetMembersRolesResp)
	default:
		holder = new(GroupsGetMembersIDSResp)
	}

	resp, err := api.Request("groups.getMembers", p, holder)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// =======================
//  GroupsGetOnlineStatus
// =======================

// GroupsGetOnlineStatusParams параметры метода GroupsGetOnlineStatus.
type GroupsGetOnlineStatusParams struct {
	GroupID uint
}

// GroupsGetOnlineStatusResp структура, возвращаемая методом GroupsGetOnlineStatus.
type GroupsGetOnlineStatusResp struct {
	Status  string `json:"status"`
	Minutes int    `json:"minutes"`
}

// GroupsGetOnlineStatus получает информацию о статусе "онлайн" в сообществе.
func (api *API) GroupsGetOnlineStatus(p GroupsGetOnlineStatusParams) (*GroupsGetOnlineStatusResp, error) {
	resp, err := api.Request("groups.getOnlineStatus", p, new(GroupsGetOnlineStatusResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetOnlineStatusResp), nil
}

// ===================
//  GroupsGetRequests
// ===================

// GroupsGetRequestsParams параметры метода GroupsGetRequests.
type GroupsGetRequestsParams struct {
	GroupID uint
	Offset  uint
	Count   uint
	Fields  string
}

// GroupsGetRequestsIDSResp структура, возвращаемая методом GroupsGetRequests, содержит идентификаторы пользователей.
type GroupsGetRequestsIDSResp struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

// GroupsGetRequestsUsersResp структура, возвращаемая методом GroupsGetRequests, содержит объекты-профили.
type GroupsGetRequestsUsersResp struct {
	Count int    `json:"count"`
	Items []User `json:"items"`
}

// GroupsGetRequests возвращает список заявок на вступление в сообщество. Чтобы использовать этот метод, вы должны быть руководителем сообщества. Возвращает *GroupsGetRequestsUsersResp, если задан параметр Fields или *GroupsGetRequestsIDSResp в остальных случаях.
func (api *API) GroupsGetRequests(p GroupsGetRequestsParams) (interface{}, error) {
	var holder interface{}
	switch len(p.Fields) > 0 {
	case true:
		holder = new(GroupsGetRequestsUsersResp)
	default:
		holder = new(GroupsGetRequestsIDSResp)
	}

	resp, err := api.Request("groups.getRequests", p, holder)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ===================
//  GroupsGetSettings
// ===================

// GroupsGetSettingsParams параметры метода GroupsGetSettings.
type GroupsGetSettingsParams struct {
	GroupID uint
}

// GroupsGetSettingsResp структура, возвращаемая методом GroupsGetSettings.
type GroupsGetSettingsResp struct {
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Address            string   `json:"address"`
	Place              *Place   `json:"place"`
	Wall               int      `json:"wall"`
	Photos             int      `json:"photos"`
	Video              int      `json:"video"`
	Audio              int      `json:"audio"`
	Docs               int      `json:"docs"`
	Topics             int      `json:"topics"`
	Wiki               int      `json:"wiki"`
	Messages           int      `json:"messages"`
	ObsceneFilter      int      `json:"obscene_filter"`
	ObsceneStopwords   int      `json:"obscene_stopwords"`
	ObsceneWords       []string `json:"obscene_words"`
	PublicCategory     int      `json:"public_category"`
	PublicSubcategory  int      `json:"public_subcategory"`
	PublicCategoryList []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		SubtypesList []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"subtypes_list"`
	} `json:"public_category_list"`
	Access      int `json:"access"`
	Subject     int `json:"subject"`
	SubjectList []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"subject_list"`
	RSS       string `json:"rss"`
	Website   string `json:"website"`
	AgeLimits int    `json:"age_limits"`
	Market    *struct {
		Enabled         int   `json:"enabled"`
		CommentsEnabled int   `json:"comments_enabled"`
		CountryIDS      []int `json:"country_ids"`
		CityIDS         []int `json:"city_ids"`
		ContactID       int   `json:"contact_id"`
		Currency        *struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"currency"`
	} `json:"market"`
}

// GroupsGetSettings позволяет получать данные, необходимые для отображения страницы редактирования данных сообщества. Чтобы использовать этот метод, вы должны быть администратором сообщества.
func (api *API) GroupsGetSettings(p GroupsGetSettingsParams) (*GroupsGetSettingsResp, error) {
	resp, err := api.Request("groups.getSettings", p, new(GroupsGetSettingsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetSettingsResp), nil
}

// ===========================
//  GroupsGetTokenPermissions
// ===========================

// GroupsGetTokenPermissionsResp структура, возвращаемая методом GroupsGetTokenPermissions.
type GroupsGetTokenPermissionsResp struct {
	Mask        int `json:"mask"`
	Permissions []struct {
		Setting int    `json:"setting"`
		Name    string `json:"name"`
	} `json:"permissions"`
}

// GroupsGetTokenPermissions возвращает настройки прав для ключа доступа сообщества.
func (api *API) GroupsGetTokenPermissions() (*GroupsGetTokenPermissionsResp, error) {
	resp, err := api.Request("groups.getTokenPermissions", struct{}{}, new(GroupsGetTokenPermissionsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsGetTokenPermissionsResp), nil
}

// ==============
//  GroupsInvite
// ==============

// GroupsInviteParams параметры метода GroupsInvite.
type GroupsInviteParams struct {
	GroupID uint
	UserID  uint
}

// GroupsInvite позволяет приглашать друзей в группу.
func (api *API) GroupsInvite(p GroupsInviteParams) (bool, error) {
	resp, err := api.Request("groups.invite", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ================
//  GroupsIsMember
// ================

// GroupsIsMemberParams параметры метода GroupsIsMember.
type GroupsIsMemberParams struct {
	GroupID  string
	UserID   uint
	UserIDS  []int
	Extended bool
}

// GroupsIsMemberResp структура, возвращаемая методом GroupsIsMember.
type GroupsIsMemberResp []struct {
	UserID     int `json:"user_id"`
	Member     int `json:"member"`
	Request    int `json:"request"`
	Invitation int `json:"invitation"`
	CanInvite  int `json:"can_invite"`
	CanRecall  int `json:"can_recall"`
}

// GroupsIsMember возвращает информацию о том, является ли пользователь участником сообщества.
func (api *API) GroupsIsMember(p GroupsIsMemberParams) (*GroupsIsMemberResp, error) {
	if p.UserID != 0 {
		p.UserIDS = append(p.UserIDS, int(p.UserID))
		p.UserID = 0
	}

	resp, err := api.Request("groups.isMember", p, new(GroupsIsMemberResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsIsMemberResp), nil
}

// ============
//  GroupsJoin
// ============

// GroupsJoinParams параметры метода GroupsJoin.
type GroupsJoinParams struct {
	GroupID uint
	NotSure uint
}

// GroupsJoin данный метод позволяет вступить в группу, публичную страницу, а также подтвердить участие во встрече.
func (api *API) GroupsJoin(p GroupsJoinParams) (bool, error) {
	resp, err := api.Request("groups.join", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =============
//  GroupsLeave
// =============

// GroupsLeaveParams параметры метода GroupsLeave.
type GroupsLeaveParams struct {
	GroupID uint
}

// GroupsLeave позволяет покинуть сообщество или отклонить приглашение в сообщество.
func (api *API) GroupsLeave(p GroupsLeaveParams) (bool, error) {
	resp, err := api.Request("groups.leave", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ==================
//  GroupsRemoveUser
// ==================

// GroupsRemoveUserParams параметры метода GroupsRemoveUser.
type GroupsRemoveUserParams struct {
	GroupID uint
	UserID  uint
}

// GroupsRemoveUser позволяет исключить пользователя из группы или отклонить заявку на вступление. Для того, чтобы воспользоваться этим методом, вы должны быть администратором группы.
func (api *API) GroupsRemoveUser(p GroupsRemoveUserParams) (bool, error) {
	resp, err := api.Request("groups.removeUser", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===================
//  GroupsReorderLink
// ===================

// GroupsReorderLinkParams параметры метода GroupsReorderLink.
type GroupsReorderLinkParams struct {
	GroupID uint
	LinkID  uint
	After   uint
}

// GroupsReorderLink позволяет менять местоположение ссылки в списке. Список ссылок может быть получен методом groups.getbyid, с параметром fields=links. Для того, чтобы воспользоваться этим методом вы должны быть администратором группы.
func (api *API) GroupsReorderLink(p GroupsReorderLinkParams) (bool, error) {
	resp, err := api.Request("groups.reorderLink", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ==============
//  GroupsSearch
// ==============

// GroupsSearchParams параметры метода GroupsSearch.
type GroupsSearchParams struct {
	Q         string
	Type      string
	CountryID uint
	CityID    uint
	Future    bool
	Market    bool
	Sort      int
	Offset    uint
	Count     uint
}

// GroupsSearchResp структура, возвращаемая методом GroupsSearch.
type GroupsSearchResp struct {
	Count int     `json:"count"`
	Items []Group `json:"items"`
}

// GroupsSearch осуществляет поиск сообществ по заданной подстроке.
func (api *API) GroupsSearch(p GroupsSearchParams) (*GroupsSearchResp, error) {
	resp, err := api.Request("groups.search", p, new(GroupsSearchResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupsSearchResp), nil
}

// ===========================
//  GroupsSetCallbackSettings
// ===========================

// GroupsSetCallbackSettingsParams параметры метода GroupsSetCallbackSettings.
type GroupsSetCallbackSettingsParams struct {
	GroupID              uint
	ServerID             uint
	APIVersion           string
	MessageNew           bool
	MessageReply         bool
	MessageAllow         bool
	MessageEdit          bool
	MessageDeny          bool
	MessageTypingState   bool
	PhotoNew             bool
	AudioNew             bool
	VideoNew             bool
	WallReplyNew         bool
	WallReplyEdit        bool
	WallReplyDelete      bool
	WallReplyRestore     bool
	WallPostNew          bool
	WallRepost           bool
	BoardPostNew         bool
	BoardPostEdit        bool
	BoardPostRestore     bool
	BoardPostDelete      bool
	PhotoCommentNew      bool
	PhotoCommentEdit     bool
	PhotoCommentDelete   bool
	PhotoCommentRestore  bool
	VideoCommentNew      bool
	VideoCommentEdit     bool
	VideoCommentDelete   bool
	VideoCommentRestore  bool
	MarketCommentNew     bool
	MarketCommentEdit    bool
	MarketCommentDelete  bool
	MarketCommentRestore bool
	PollVoteNew          bool
	GroupJoin            bool
	GroupLeave           bool
	GroupChangeSettings  bool
	GroupChangePhoto     bool
	GroupOfficersEdit    bool
	UserBlock            bool
	UserUnblock          bool
	LeadFormsNew         bool
}

// GroupsSetCallbackSettings позволяет задать настройки уведомлений о событиях в callback api.
func (api *API) GroupsSetCallbackSettings(p GroupsSetCallbackSettingsParams) (bool, error) {
	resp, err := api.Request("groups.setCallbackSettings", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===========================
//  GroupsSetLongPollSettings
// ===========================

// GroupsSetLongPollSettingsParams параметры метода GroupsSetLongPollSettings.
type GroupsSetLongPollSettingsParams struct {
	GroupID              uint
	Enabled              bool
	APIVersion           string
	MessageNew           bool
	MessageReply         bool
	MessageAllow         bool
	MessageDeny          bool
	MessageEdit          bool
	MessageTypingState   bool
	PhotoNew             bool
	AudioNew             bool
	VideoNew             bool
	WallReplyNew         bool
	WallReplyEdit        bool
	WallReplyDelete      bool
	WallReplyRestore     bool
	WallPostNew          bool
	WallRepost           bool
	BoardPostNew         bool
	BoardPostEdit        bool
	BoardPostRestore     bool
	BoardPostDelete      bool
	PhotoCommentNew      bool
	PhotoCommentEdit     bool
	PhotoCommentDelete   bool
	PhotoCommentRestore  bool
	VideoCommentNew      bool
	VideoCommentEdit     bool
	VideoCommentDelete   bool
	VideoCommentRestore  bool
	MarketCommentNew     bool
	MarketCommentEdit    bool
	MarketCommentDelete  bool
	MarketCommentRestore bool
	PollVoteNew          bool
	GroupJoin            bool
	GroupLeave           bool
	GroupChangeSettings  bool
	GroupChangePhoto     bool
	GroupOfficersEdit    bool
	UserBlock            bool
	UserUnblock          bool
}

// GroupsSetLongPollSettings задаёт настройки для bots long poll api в сообществе.
func (api *API) GroupsSetLongPollSettings(p GroupsSetLongPollSettingsParams) (bool, error) {
	resp, err := api.Request("groups.setLongPollSettings", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===================
//  GroupsSetSettings
// ===================

// GroupsSetSettingsParams параметры метода GroupsSetSettings.
type GroupsSetSettingsParams struct {
	GroupID          uint
	Messages         bool
	BotsCapabilities bool
	BotsStartButton  bool
	BotsAddToChat    bool
}

// GroupsSetSettings устанавливает настройки сообщества для того, чтобы воспользоваться этим методом, вы должны быть администратором сообщества.
func (api *API) GroupsSetSettings(p GroupsSetSettingsParams) (bool, error) {
	resp, err := api.Request("groups.setSettings", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =============
//  GroupsUnban
// =============

// GroupsUnbanParams параметры метода GroupsUnban.
type GroupsUnbanParams struct {
	GroupID uint
	OwnerID int
}

// GroupsUnban убирает пользователя или группу из черного списка сообщества.
func (api *API) GroupsUnban(p GroupsUnbanParams) (bool, error) {
	resp, err := api.Request("groups.unban", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}
