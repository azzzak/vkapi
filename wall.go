package vkapi

// ===================
//  WallCloseComments
// ===================

// WallCloseCommentsParams параметры метода WallCloseComments.
type WallCloseCommentsParams struct {
	OwnerID int
	PostID  uint
}

// WallCloseComments выключает комментирование записи.
func (api *API) WallCloseComments(p WallCloseCommentsParams) (bool, error) {
	resp, err := api.Request("wall.closeComments", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===================
//  WallCreateComment
// ===================

// WallCreateCommentParams параметры метода WallCreateComment.
type WallCreateCommentParams struct {
	OwnerID        int
	PostID         uint
	FromGroup      uint
	Message        string
	ReplyToComment int
	Attachments    string
	StickerID      uint
	GUID           string
}

// WallCreateCommentResp структура, возвращаемая методом WallCreateComment.
type WallCreateCommentResp struct {
	CommentID    int   `json:"comment_id"`
	ParentsStack []int `json:"parents_stack"`
}

// WallCreateComment добавляет комментарий к записи на стене.
func (api *API) WallCreateComment(p WallCreateCommentParams) (*WallCreateCommentResp, error) {
	resp, err := api.Request("wall.createComment", p, new(WallCreateCommentResp))
	if err != nil {
		return nil, err
	}
	return resp.(*WallCreateCommentResp), nil
}

// ============
//  WallDelete
// ============

// WallDeleteParams параметры метода WallDelete.
type WallDeleteParams struct {
	OwnerID int
	PostID  uint
}

// WallDelete удаляет запись со стены.
func (api *API) WallDelete(p WallDeleteParams) (bool, error) {
	resp, err := api.Request("wall.delete", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===================
//  WallDeleteComment
// ===================

// WallDeleteCommentParams параметры метода WallDeleteComment.
type WallDeleteCommentParams struct {
	OwnerID   int
	CommentID uint
}

// WallDeleteComment удаляет комментарий к записи на стене.
func (api *API) WallDeleteComment(p WallDeleteCommentParams) (bool, error) {
	resp, err := api.Request("wall.deleteComment", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ==========
//  WallEdit
// ==========

// WallEditParams параметры метода WallEdit.
type WallEditParams struct {
	OwnerID             int
	PostID              uint
	FriendsOnly         bool
	Message             string
	Attachments         string
	Services            string
	Signed              bool
	PublishDate         int64
	Lat                 float32
	Long                float32
	PlaceID             uint
	MarkAsAds           bool
	CloseComments       bool
	PosterBkgID         uint
	PosterBkgOwnerID    int
	PosterBkgAccessHash string
}

// WallEditResp структура, возвращаемая методом WallEdit.
type WallEditResp struct {
	PostID int `json:"post_id"`
}

// WallEdit редактирует запись на стене.
func (api *API) WallEdit(p WallEditParams) (*WallEditResp, error) {
	resp, err := api.Request("wall.edit", p, new(WallEditResp))
	if err != nil {
		return nil, err
	}
	return resp.(*WallEditResp), nil
}

// ====================
//  WallEditAdsStealth
// ====================

// WallEditAdsStealthParams параметры метода WallEditAdsStealth.
type WallEditAdsStealthParams struct {
	OwnerID     int
	PostID      uint
	Message     string
	Attachments string
	Signed      bool
	Lat         float32
	Long        float32
	PlaceID     uint
	LinkButton  string
	LinkTitle   string
	LinkImage   string
	LinkVideo   string
}

// WallEditAdsStealth позволяет отредактировать скрытую запись. Создание скрытых записей возможно только в сообществах от имени группы, публичной страницы или мероприятия; пользователь должен обладать правами администратора или редактора.
func (api *API) WallEditAdsStealth(p WallEditAdsStealthParams) (bool, error) {
	resp, err := api.Request("wall.editAdsStealth", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =================
//  WallEditComment
// =================

// WallEditCommentParams параметры метода WallEditComment.
type WallEditCommentParams struct {
	OwnerID     int
	CommentID   uint
	Message     string
	Attachments string
}

// WallEditComment редактирует комментарий на стене.
func (api *API) WallEditComment(p WallEditCommentParams) (bool, error) {
	resp, err := api.Request("wall.editComment", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =========
//  WallGet
// =========

// WallGetParams параметры метода WallGet.
type WallGetParams struct {
	OwnerID  int
	Domain   string
	Offset   uint
	Count    uint
	Filter   string
	Extended bool
	Fields   string
}

// WallGetResp структура, возвращаемая методом WallGet.
type WallGetResp struct {
	Count    int     `json:"count"`
	Items    []Post  `json:"items"`
	Profiles []User  `json:"profiles"`
	Groups   []Group `json:"groups"`
}

// WallGet возвращает список записей со стены пользователя или сообщества. Обратите внимание, для этого метода действуют специальные ограничения на число запросов. https://vk.com/dev/data_limits
func (api *API) WallGet(p WallGetParams) (*WallGetResp, error) {
	resp, err := api.Request("wall.get", p, new(WallGetResp))
	if err != nil {
		return nil, err
	}
	return resp.(*WallGetResp), nil
}

// =============
//  WallGetByID
// =============

// WallGetByIDParams параметры метода WallGetById.
type WallGetByIDParams struct {
	Posts            string
	Extended         bool
	CopyHistoryDepth uint
	Fields           string
}

// WallGetByIDResp структура, возвращаемая методом WallGetById.
type WallGetByIDResp struct {
	Items    []Post  `json:"items"`
	Profiles []User  `json:"profiles"`
	Groups   []Group `json:"groups"`
}

// WallGetByID возвращает список записей со стен пользователей или сообществ по их идентификаторам. Возвращает []Post, если Extended=false или *WallGetByIDResp, если Extended=true.
func (api *API) WallGetByID(p WallGetByIDParams) (interface{}, error) {
	var holder interface{}
	switch p.Extended {
	case true:
		holder = new(WallGetByIDResp)
	default:
		holder = new([]Post)
	}

	resp, err := api.Request("wall.getById", p, holder)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ================
//  WallGetComment
// ================

// WallGetCommentParams параметры метода WallGetComment.
type WallGetCommentParams struct {
	OwnerID   int
	CommentID uint
	Extended  bool
	Fields    string
}

// WallGetCommentResp структура, возвращаемая методом WallGetComment.
type WallGetCommentResp struct {
	Items    []Comment `json:"items"`
	Profiles []User    `json:"profiles"`
	Groups   []Group   `json:"groups"`
}

// WallGetComment получает информацию о комментарии на стене.
func (api *API) WallGetComment(p WallGetCommentParams) (*WallGetCommentResp, error) {
	resp, err := api.Request("wall.getComment", p, new(WallGetCommentResp))
	if err != nil {
		return nil, err
	}
	return resp.(*WallGetCommentResp), nil
}

// =================
//  WallGetComments
// =================

// WallGetCommentsParams параметры метода WallGetComments.
type WallGetCommentsParams struct {
	OwnerID          int
	PostID           uint
	NeedLikes        bool
	StartCommentID   uint
	Offset           int
	Count            uint
	Sort             string
	PreviewLength    uint
	Extended         bool
	Fields           string
	CommentID        uint
	ThreadItemsCount uint
}

// WallGetCommentsResp структура, возвращаемая методом WallGetComments.
type WallGetCommentsResp struct {
	Count             int       `json:"count"`
	Items             []Comment `json:"items"`
	CurrentLevelCount int       `json:"current_level_count"`
	CanPost           bool      `json:"can_post"`
	ShowReplyButton   bool      `json:"show_reply_button"`
	GroupsCanPost     bool      `json:"groups_can_post"`
	RealOffset        int       `json:"real_offset"`
	Profiles          []User    `json:"profiles"`
	Groups            []Group   `json:"groups"`
}

// WallGetComments возвращает список комментариев к записи на стене.
func (api *API) WallGetComments(p WallGetCommentsParams) (*WallGetCommentsResp, error) {
	resp, err := api.Request("wall.getComments", p, new(WallGetCommentsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*WallGetCommentsResp), nil
}

// ================
//  WallGetReposts
// ================

// WallGetRepostsParams параметры метода WallGetReposts.
type WallGetRepostsParams struct {
	OwnerID int
	PostID  uint
	Offset  uint
	Count   uint
}

// WallGetRepostsResp структура, возвращаемая методом WallGetReposts.
type WallGetRepostsResp struct {
	Items    []Post  `json:"items"`
	Profiles []User  `json:"profiles"`
	Groups   []Group `json:"groups"`
}

// WallGetReposts позволяет получать список репостов заданной записи. Обратите внимание, получить список репостов можно только для записи, созданной текущим пользователем, или в сообществе, где текущий пользователь является администратором.
func (api *API) WallGetReposts(p WallGetRepostsParams) (*WallGetRepostsResp, error) {
	resp, err := api.Request("wall.getReposts", p, new(WallGetRepostsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*WallGetRepostsResp), nil
}

// ==================
//  WallOpenComments
// ==================

// WallOpenCommentsParams параметры метода WallOpenComments.
type WallOpenCommentsParams struct {
	OwnerID int
	PostID  uint
}

// WallOpenComments включает комментирование записи работает только с конкретными записями, комментирование которых было выключено с помощью wall.closecomments
func (api *API) WallOpenComments(p WallOpenCommentsParams) (bool, error) {
	resp, err := api.Request("wall.openComments", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =========
//  WallPin
// =========

// WallPinParams параметры метода WallPin.
type WallPinParams struct {
	OwnerID int
	PostID  uint
}

// WallPin закрепляет запись на стене (запись будет отображаться выше остальных).
func (api *API) WallPin(p WallPinParams) (bool, error) {
	resp, err := api.Request("wall.pin", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ==========
//  WallPost
// ==========

// WallPostParams параметры метода WallPost.
type WallPostParams struct {
	OwnerID           int
	FriendsOnly       bool
	FromGroup         bool
	Message           string
	Attachments       string
	Services          string
	Signed            bool
	PublishDate       int64
	Lat               float32
	Long              float32
	PlaceID           uint
	PostID            uint
	GUID              string
	MarkAsAds         bool
	CloseComments     bool
	MuteNotifications bool
}

// WallPostResp структура, возвращаемая методом WallPost.
type WallPostResp struct {
	PostID int `json:"post_id"`
}

// WallPost позволяет создать запись на стене, предложить запись на стене публичной страницы, опубликовать существующую отложенную запись. Чтобы создать предложенную запись, необходимо передать в owner_id идентификатор публичной страницы, в которой текущий пользователь не является руководителем. Для публикации предложенных и отложенных записей используйте параметр post_id, значение для которого можно получить методом wall.get с filter=suggests и postponed соответственно.
func (api *API) WallPost(p WallPostParams) (*WallPostResp, error) {
	resp, err := api.Request("wall.post", p, new(WallPostResp))
	if err != nil {
		return nil, err
	}
	return resp.(*WallPostResp), nil
}

// ====================
//  WallPostAdsStealth
// ====================

// WallPostAdsStealthParams параметры метода WallPostAdsStealth.
type WallPostAdsStealthParams struct {
	OwnerID     int
	Message     string
	Attachments string
	Signed      bool
	Lat         float32
	Long        float32
	PlaceID     uint
	GUID        string
	LinkButton  string
	LinkTitle   string
	LinkImage   string
	LinkVideo   string
}

// WallPostAdsStealth позволяет создать скрытую запись, которая не попадает на стену сообщества и в дальнейшем может быть использована для создания рекламного объявления типа "запись в сообществе". Cоздание скрытых записей возможно только в сообществах от имени группы, публичной страницы или мероприятия; пользователь должен обладать правами администратора или редактора. Обратите внимание — в сутки можно создать не более 100 скрытых записей.
func (api *API) WallPostAdsStealth(p WallPostAdsStealthParams) (int, error) {
	resp, err := api.Request("wall.postAdsStealth", p, new(int))
	if err != nil {
		return 0, err
	}
	return resp.(int), nil
}

// ===================
//  WallReportComment
// ===================

// WallReportCommentParams параметры метода WallReportComment.
type WallReportCommentParams struct {
	OwnerID   int
	CommentID uint
	Reason    uint
}

// WallReportComment позволяет пожаловаться на комментарий к записи.
func (api *API) WallReportComment(p WallReportCommentParams) (bool, error) {
	resp, err := api.Request("wall.reportComment", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ================
//  WallReportPost
// ================

// WallReportPostParams параметры метода WallReportPost.
type WallReportPostParams struct {
	OwnerID int
	PostID  uint
	Reason  uint
}

// WallReportPost позволяет пожаловаться на запись.
func (api *API) WallReportPost(p WallReportPostParams) (bool, error) {
	resp, err := api.Request("wall.reportPost", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ============
//  WallRepost
// ============

// WallRepostParams параметры метода WallRepost.
type WallRepostParams struct {
	Object            string
	Message           string
	GroupID           uint
	MarkAsAds         bool
	MuteNotifications bool
}

// WallRepostResp структура, возвращаемая методом WallRepost.
type WallRepostResp struct {
	Success      int `json:"success"`
	PostID       int `json:"post_id"`
	RepostsCount int `json:"reposts_count"`
	LikesCount   int `json:"likes_count"`
}

// WallRepost копирует объект на стену пользователя или сообщества.
func (api *API) WallRepost(p WallRepostParams) (*WallRepostResp, error) {
	resp, err := api.Request("wall.repost", p, new(WallRepostResp))
	if err != nil {
		return nil, err
	}
	return resp.(*WallRepostResp), nil
}

// =============
//  WallRestore
// =============

// WallRestoreParams параметры метода WallRestore.
type WallRestoreParams struct {
	OwnerID int
	PostID  uint
}

// WallRestore восстанавливает удаленную запись на стене пользователя или сообщества.
func (api *API) WallRestore(p WallRestoreParams) (bool, error) {
	resp, err := api.Request("wall.restore", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ====================
//  WallRestoreComment
// ====================

// WallRestoreCommentParams параметры метода WallRestoreComment.
type WallRestoreCommentParams struct {
	OwnerID   int
	CommentID uint
}

// WallRestoreComment восстанавливает удаленный комментарий к записи на стене.
func (api *API) WallRestoreComment(p WallRestoreCommentParams) (bool, error) {
	resp, err := api.Request("wall.restoreComment", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ============
//  WallSearch
// ============

// WallSearchParams параметры метода WallSearch.
type WallSearchParams struct {
	OwnerID    int
	Domain     string
	Query      string
	OwnersOnly bool
	Count      uint
	Offset     uint
	Extended   bool
	Fields     string
}

// WallSearchResp структура, возвращаемая методом WallSearch.
type WallSearchResp struct {
	Count    int     `json:"count"`
	Items    []Post  `json:"items"`
	Profiles []User  `json:"profiles"`
	Groups   []Group `json:"groups"`
}

// WallSearch позволяет искать записи на стене в соответствии с заданными критериями. Обратите внимание, для этого метода действуют специальные ограничения на число запросов. https://vk.com/dev/data_limits
func (api *API) WallSearch(p WallSearchParams) (*WallSearchResp, error) {
	resp, err := api.Request("wall.search", p, new(WallSearchResp))
	if err != nil {
		return nil, err
	}
	return resp.(*WallSearchResp), nil
}

// ===========
//  WallUnpin
// ===========

// WallUnpinParams параметры метода WallUnpin.
type WallUnpinParams struct {
	OwnerID int
	PostID  uint
}

// WallUnpin отменяет закрепление записи на стене.
func (api *API) WallUnpin(p WallUnpinParams) (bool, error) {
	resp, err := api.Request("wall.unpin", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}
