package vkapi

// ==================
//  PhotosConfirmTag
// ==================

// PhotosConfirmTagParams параметры метода PhotosConfirmTag.
type PhotosConfirmTagParams struct {
	OwnerID int
	PhotoID int
	TagID   int
}

// PhotosConfirmTag подтверждает отметку на фотографии.
func (api *API) PhotosConfirmTag(p PhotosConfirmTagParams) (bool, error) {
	resp, err := api.Request("photos.confirmTag", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ============
//  PhotosCopy
// ============

// PhotosCopyParams параметры метода PhotosCopy.
type PhotosCopyParams struct {
	OwnerID   int
	PhotoID   uint
	AccessKey string
}

// PhotosCopy позволяет скопировать фотографию в альбом "сохраненные фотографии".
func (api *API) PhotosCopy(p PhotosCopyParams) (int, error) {
	resp, err := api.Request("photos.copy", p, new(int))
	if err != nil {
		return 0, err
	}
	return resp.(int), nil
}

// ===================
//  PhotosCreateAlbum
// ===================

// PhotosCreateAlbumParams параметры метода PhotosCreateAlbum.
type PhotosCreateAlbumParams struct {
	Title              string
	GroupID            int
	Description        string
	PrivacyView        string
	PrivacyComment     string
	UploadByAdminsOnly bool
	CommentsDisabled   bool
}

// PhotosCreateAlbumResp структура, возвращаемая методом PhotosCreateAlbum.
type PhotosCreateAlbumResp struct {
	ID          int    `json:"id"`
	ThumbID     int    `json:"thumb_id"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     int64  `json:"created"`
	Updated     int64  `json:"updated"`
	Size        int    `json:"size"`
	PrivacyView *struct {
		Category string `json:"category"`
		Owners   *struct {
			Allowed  []int `json:"allowed"`
			Excluded []int `json:"excluded"`
		} `json:"owners,omitempty"`
		Lists *struct {
			Allowed  []int `json:"allowed"`
			Excluded []int `json:"excluded"`
		} `json:"lists"`
	} `json:"privacy_view"`
	PrivacyComment *struct {
		Category string `json:"category"`
		Owners   *struct {
			Allowed  []int `json:"allowed"`
			Excluded []int `json:"excluded"`
		} `json:"owners,omitempty"`
		Lists *struct {
			Allowed  []int `json:"allowed"`
			Excluded []int `json:"excluded"`
		} `json:"lists"`
	} `json:"privacy_comment"`
	UploadByAdminsOnly int `json:"upload_by_admins_only"`
	CommentsDisabled   int `json:"comments_disabled"`
	CanUpload          int `json:"can_upload"`
}

// PhotosCreateAlbum создает пустой альбом для фотографий.
func (api *API) PhotosCreateAlbum(p PhotosCreateAlbumParams) (*PhotosCreateAlbumResp, error) {
	resp, err := api.Request("photos.createAlbum", p, new(PhotosCreateAlbumResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosCreateAlbumResp), nil
}

// =====================
//  PhotosCreateComment
// =====================

// PhotosCreateCommentParams параметры метода PhotosCreateComment.
type PhotosCreateCommentParams struct {
	OwnerID        int
	PhotoID        int
	Message        string
	Attachments    string
	FromGroup      bool
	ReplyToComment int
	StickerID      uint
	AccessKey      string
	GUID           string
}

// PhotosCreateComment создает новый комментарий к фотографии.
func (api *API) PhotosCreateComment(p PhotosCreateCommentParams) (int, error) {
	resp, err := api.Request("photos.createComment", p, new(int))
	if err != nil {
		return 0, err
	}
	return resp.(int), nil
}

// ==============
//  PhotosDelete
// ==============

// PhotosDeleteParams параметры метода PhotosDelete.
type PhotosDeleteParams struct {
	OwnerID int
	PhotoID uint
}

// PhotosDelete удаление фотографии.
func (api *API) PhotosDelete(p PhotosDeleteParams) (bool, error) {
	resp, err := api.Request("photos.delete", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===================
//  PhotosDeleteAlbum
// ===================

// PhotosDeleteAlbumParams параметры метода PhotosDeleteAlbum.
type PhotosDeleteAlbumParams struct {
	AlbumID int
	GroupID int
}

// PhotosDeleteAlbum удаляет указанный альбом для фотографий у текущего пользователя.
func (api *API) PhotosDeleteAlbum(p PhotosDeleteAlbumParams) (bool, error) {
	resp, err := api.Request("photos.deleteAlbum", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =====================
//  PhotosDeleteComment
// =====================

// PhotosDeleteCommentParams параметры метода PhotosDeleteComment.
type PhotosDeleteCommentParams struct {
	OwnerID   int
	CommentID int
}

// PhotosDeleteComment удаляет комментарий к фотографии.
func (api *API) PhotosDeleteComment(p PhotosDeleteCommentParams) (bool, error) {
	resp, err := api.Request("photos.deleteComment", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ============
//  PhotosEdit
// ============

// PhotosEditParams параметры метода PhotosEdit.
type PhotosEditParams struct {
	OwnerID      int
	PhotoID      uint
	Caption      string
	Latitude     float32
	Longitude    float32
	PlaceStr     string
	FoursquareID string
	DeletePlace  bool
}

// PhotosEdit редактирует описание или геометку у фотографии.
func (api *API) PhotosEdit(p PhotosEditParams) (bool, error) {
	resp, err := api.Request("photos.edit", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =================
//  PhotosEditAlbum
// =================

// PhotosEditAlbumParams параметры метода PhotosEditAlbum.
type PhotosEditAlbumParams struct {
	AlbumID            int
	Title              string
	Description        string
	OwnerID            int
	PrivacyView        string
	PrivacyComment     string
	UploadByAdminsOnly bool
	CommentsDisabled   bool
}

// PhotosEditAlbum редактирует данные альбома для фотографий.
func (api *API) PhotosEditAlbum(p PhotosEditAlbumParams) (bool, error) {
	resp, err := api.Request("photos.editAlbum", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===================
//  PhotosEditComment
// ===================

// PhotosEditCommentParams параметры метода PhotosEditComment.
type PhotosEditCommentParams struct {
	OwnerID     int
	CommentID   int
	Message     string
	Attachments string
}

// PhotosEditComment изменяет текст комментария к фотографии. Обратите внимание, что редактирование комментария доступно только в течение суток после его создания.
func (api *API) PhotosEditComment(p PhotosEditCommentParams) (bool, error) {
	resp, err := api.Request("photos.editComment", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===========
//  PhotosGet
// ===========

// PhotosGetParams параметры метода PhotosGet.
type PhotosGetParams struct {
	OwnerID    int
	AlbumID    string
	PhotoIDS   []int
	Rev        bool
	Extended   bool
	FeedType   string
	Feed       int
	PhotoSizes bool
	Offset     uint
	Count      uint
}

// PhotosGetResp структура, возвращаемая методом PhotosGet.
type PhotosGetResp struct {
	Count int
	Items []Photo
}

// PhotosGet возвращает список фотографий в альбоме.
func (api *API) PhotosGet(p PhotosGetParams) (*PhotosGetResp, error) {
	resp, err := api.Request("photos.get", p, new(PhotosGetResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetResp), nil
}

// =================
//  PhotosGetAlbums
// =================

// PhotosGetAlbumsParams параметры метода PhotosGetAlbums.
type PhotosGetAlbumsParams struct {
	OwnerID    int
	AlbumIDS   []int
	Offset     uint
	Count      uint
	NeedSystem bool
	NeedCovers bool
	PhotoSizes bool
}

// PhotosGetAlbumsResp структура, возвращаемая методом PhotosGetAlbums.
type PhotosGetAlbumsResp struct {
	Count int `json:"count"`
	Items []struct {
		ID          int    `json:"id"`
		ThumbID     int    `json:"thumb_id"`
		OwnerID     int    `json:"owner_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Created     int64  `json:"created"`
		Updated     int64  `json:"updated"`
		Size        int    `json:"size"`
		CanUpload   int    `json:"can_upload"`
		PrivacyView *struct {
			Category string `json:"category"`
			Owners   *struct {
				Allowed  []int `json:"allowed"`
				Excluded []int `json:"excluded"`
			} `json:"owners,omitempty"`
			Lists *struct {
				Allowed  []int `json:"allowed"`
				Excluded []int `json:"excluded"`
			} `json:"lists"`
		} `json:"privacy_view"`
		PrivacyComment *struct {
			Category string `json:"category"`
			Owners   *struct {
				Allowed  []int `json:"allowed"`
				Excluded []int `json:"excluded"`
			} `json:"owners,omitempty"`
			Lists *struct {
				Allowed  []int `json:"allowed"`
				Excluded []int `json:"excluded"`
			} `json:"lists"`
		} `json:"privacy_comment"`
		UploadByAdminsOnly int    `json:"upload_by_admins_only"`
		CommentsDisabled   int    `json:"comments_disabled"`
		ThumbSrc           string `json:"thumb_src"`
	} `json:"items"`
}

// PhotosGetAlbums возвращает список фотоальбомов пользователя или сообщества.
func (api *API) PhotosGetAlbums(p PhotosGetAlbumsParams) (*PhotosGetAlbumsResp, error) {
	resp, err := api.Request("photos.getAlbums", p, new(PhotosGetAlbumsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetAlbumsResp), nil
}

// ======================
//  PhotosGetAlbumsCount
// ======================

// PhotosGetAlbumsCountParams параметры метода PhotosGetAlbumsCount.
type PhotosGetAlbumsCountParams struct {
	UserID  int
	GroupID int
}

// PhotosGetAlbumsCount возвращает количество доступных альбомов пользователя или сообщества.
func (api *API) PhotosGetAlbumsCount(p PhotosGetAlbumsCountParams) (int, error) {
	resp, err := api.Request("photos.getAlbumsCount", p, new(int))
	if err != nil {
		return 0, err
	}
	return resp.(int), nil
}

// ==============
//  PhotosGetAll
// ==============

// PhotosGetAllParams параметры метода PhotosGetAll.
type PhotosGetAllParams struct {
	OwnerID         int
	Extended        bool
	Offset          uint
	Count           uint
	PhotoSizes      bool
	NoServiceAlbums bool
	NeedHidden      bool
	SkipHidden      bool
}

// PhotosGetAllResp структура, возвращаемая методом PhotosGetAll.
type PhotosGetAllResp struct {
	Count int     `json:"count"`
	Items []Photo `json:"items"`
	More  int     `json:"more"`
}

// PhotosGetAll возвращает все фотографии пользователя или сообщества в антихронологическом порядке.
func (api *API) PhotosGetAll(p PhotosGetAllParams) (*PhotosGetAllResp, error) {
	resp, err := api.Request("photos.getAll", p, new(PhotosGetAllResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetAllResp), nil
}

// ======================
//  PhotosGetAllComments
// ======================

// PhotosGetAllCommentsParams параметры метода PhotosGetAllComments.
type PhotosGetAllCommentsParams struct {
	OwnerID   int
	AlbumID   uint
	NeedLikes bool
	Offset    uint
	Count     uint
}

// PhotosGetAllCommentsResp структура, возвращаемая методом PhotosGetAllComments.
type PhotosGetAllCommentsResp struct {
	Count int       `json:"count"`
	Items []Comment `json:"items"`
}

// PhotosGetAllComments возвращает отсортированный в антихронологическом порядке список всех комментариев к конкретному альбому или ко всем альбомам пользователя.
func (api *API) PhotosGetAllComments(p PhotosGetAllCommentsParams) (*PhotosGetAllCommentsResp, error) {
	resp, err := api.Request("photos.getAllComments", p, new(PhotosGetAllCommentsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetAllCommentsResp), nil
}

// ===============
//  PhotosGetByID
// ===============

// PhotosGetByIDParams параметры метода PhotosGetByID.
type PhotosGetByIDParams struct {
	Photos     string
	Extended   bool
	PhotoSizes bool
}

// PhotosGetByID возвращает информацию о фотографиях по их идентификаторам.
func (api *API) PhotosGetByID(p PhotosGetByIDParams) ([]Photo, error) {
	resp, err := api.Request("photos.getById", p, new([]Photo))
	if err != nil {
		return nil, err
	}
	return resp.([]Photo), nil
}

// ===========================
//  PhotosGetChatUploadServer
// ===========================

// PhotosGetChatUploadServerParams параметры метода PhotosGetChatUploadServer.
type PhotosGetChatUploadServerParams struct {
	ChatID    uint
	CropX     uint
	CropY     uint
	CropWidth uint
}

// PhotosGetChatUploadServerResp структура, возвращаемая методом PhotosGetChatUploadServer.
type PhotosGetChatUploadServerResp struct {
	UploadURL string `json:"upload_url"`
}

// PhotosGetChatUploadServer позволяет получить адрес для загрузки обложки чата.
func (api *API) PhotosGetChatUploadServer(p PhotosGetChatUploadServerParams) (*PhotosGetChatUploadServerResp, error) {
	resp, err := api.Request("photos.getChatUploadServer", p, new(PhotosGetChatUploadServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetChatUploadServerResp), nil
}

// ===================
//  PhotosGetComments
// ===================

// PhotosGetCommentsParams параметры метода PhotosGetComments.
type PhotosGetCommentsParams struct {
	OwnerID        int
	PhotoID        int
	NeedLikes      bool
	StartCommentID uint
	Offset         int
	Count          uint
	Sort           string
	AccessKey      string
	Extended       bool
	Fields         string
}

// PhotosGetCommentsResp структура, возвращаемая методом PhotosGetComments.
type PhotosGetCommentsResp struct {
	Count      int       `json:"count"`
	Items      []Comment `json:"items"`
	Profiles   []User    `json:"profiles"`
	Groups     []Group   `json:"groups"`
	RealOffset int       `json:"real_offset"`
}

// PhotosGetComments возвращает список комментариев к фотографии.
func (api *API) PhotosGetComments(p PhotosGetCommentsParams) (*PhotosGetCommentsResp, error) {
	resp, err := api.Request("photos.getComments", p, new(PhotosGetCommentsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetCommentsResp), nil
}

// ==================================
//  PhotosGetMarketAlbumUploadServer
// ==================================

// PhotosGetMarketAlbumUploadServerParams параметры метода PhotosGetMarketAlbumUploadServer.
type PhotosGetMarketAlbumUploadServerParams struct {
	GroupID uint
}

// PhotosGetMarketAlbumUploadServerResp структура, возвращаемая методом PhotosGetMarketAlbumUploadServer.
type PhotosGetMarketAlbumUploadServerResp struct {
	UploadURL string `json:"upload_url"`
}

// PhotosGetMarketAlbumUploadServer возвращает адрес сервера для загрузки фотографии подборки товаров в сообществе. После успешной загрузки вы можете сохранить фотографию, воспользовавшись методом photos.savemarketalbumphoto.
func (api *API) PhotosGetMarketAlbumUploadServer(p PhotosGetMarketAlbumUploadServerParams) (*PhotosGetMarketAlbumUploadServerResp, error) {
	resp, err := api.Request("photos.getMarketAlbumUploadServer", p, new(PhotosGetMarketAlbumUploadServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetMarketAlbumUploadServerResp), nil
}

// =============================
//  PhotosGetMarketUploadServer
// =============================

// PhotosGetMarketUploadServerParams параметры метода PhotosGetMarketUploadServer.
type PhotosGetMarketUploadServerParams struct {
	GroupID   uint
	MainPhoto bool
	CropX     uint
	CropY     uint
	CropWidth uint
}

// PhotosGetMarketUploadServerResp структура, возвращаемая методом PhotosGetMarketUploadServer.
type PhotosGetMarketUploadServerResp struct {
	UploadURL string `json:"upload_url"`
}

// PhotosGetMarketUploadServer возвращает адрес сервера для загрузки фотографии товара. После успешной загрузки вы можете сохранить фотографию, воспользовавшись методом photos.savemarketphoto.
func (api *API) PhotosGetMarketUploadServer(p PhotosGetMarketUploadServerParams) (*PhotosGetMarketUploadServerResp, error) {
	resp, err := api.Request("photos.getMarketUploadServer", p, new(PhotosGetMarketUploadServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetMarketUploadServerResp), nil
}

// ===============================
//  PhotosGetMessagesUploadServer
// ===============================

// PhotosGetMessagesUploadServerParams параметры метода PhotosGetMessagesUploadServer.
type PhotosGetMessagesUploadServerParams struct {
	PeerID int
}

// PhotosGetMessagesUploadServerResp структура, возвращаемая методом PhotosGetMessagesUploadServer.
type PhotosGetMessagesUploadServerResp struct {
	UploadURL string `json:"upload_url"`
	AlbumID   int    `json:"album_id"`
	GroupID   int    `json:"group_id"`
}

// PhotosGetMessagesUploadServer возвращает адрес сервера для загрузки фотографии в личное сообщение. После успешной загрузки вы можете сохранить фотографию, воспользовавшись методом photos.savemessagesphoto.
func (api *API) PhotosGetMessagesUploadServer(p PhotosGetMessagesUploadServerParams) (*PhotosGetMessagesUploadServerResp, error) {
	resp, err := api.Request("photos.getMessagesUploadServer", p, new(PhotosGetMessagesUploadServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetMessagesUploadServerResp), nil
}

// ==================
//  PhotosGetNewTags
// ==================

// PhotosGetNewTagsParams параметры метода PhotosGetNewTags.
type PhotosGetNewTagsParams struct {
	Offset int
	Count  uint
}

// PhotosGetNewTagsResp структура, возвращаемая методом PhotosGetNewTags.
type PhotosGetNewTagsResp struct {
	Count int     `json:"count"`
	Items []Photo `json:"items"`
}

// PhotosGetNewTags возвращает список фотографий, на которых есть непросмотренные отметки.
func (api *API) PhotosGetNewTags(p PhotosGetNewTagsParams) (*PhotosGetNewTagsResp, error) {
	resp, err := api.Request("photos.getNewTags", p, new(PhotosGetNewTagsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetNewTagsResp), nil
}

// ======================================
//  PhotosGetOwnerCoverPhotoUploadServer
// ======================================

// PhotosGetOwnerCoverPhotoUploadServerParams параметры метода PhotosGetOwnerCoverPhotoUploadServer.
type PhotosGetOwnerCoverPhotoUploadServerParams struct {
	GroupID uint
	CropX   uint
	CropY   uint
	CropX2  uint
	CropY2  uint
}

// PhotosGetOwnerCoverPhotoUploadServerResp структура, возвращаемая методом PhotosGetOwnerCoverPhotoUploadServer.
type PhotosGetOwnerCoverPhotoUploadServerResp struct {
	UploadURL string `json:"upload_url"`
}

// PhotosGetOwnerCoverPhotoUploadServer получает адрес для загрузки обложки сообщества.
func (api *API) PhotosGetOwnerCoverPhotoUploadServer(p PhotosGetOwnerCoverPhotoUploadServerParams) (*PhotosGetOwnerCoverPhotoUploadServerResp, error) {
	resp, err := api.Request("photos.getOwnerCoverPhotoUploadServer", p, new(PhotosGetOwnerCoverPhotoUploadServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetOwnerCoverPhotoUploadServerResp), nil
}

// =================================
//  PhotosGetOwnerPhotoUploadServer
// =================================

// PhotosGetOwnerPhotoUploadServerParams параметры метода PhotosGetOwnerPhotoUploadServer.
type PhotosGetOwnerPhotoUploadServerParams struct {
	OwnerID int
}

// PhotosGetOwnerPhotoUploadServerResp структура, возвращаемая методом PhotosGetOwnerPhotoUploadServer.
type PhotosGetOwnerPhotoUploadServerResp struct {
	UploadURL string `json:"upload_url"`
}

// PhotosGetOwnerPhotoUploadServer возвращает адрес сервера для загрузки главной фотографии на страницу пользователя или сообщества. После удачной загрузки вы можете воспользоваться методом photos.saveownerphoto.
func (api *API) PhotosGetOwnerPhotoUploadServer(p PhotosGetOwnerPhotoUploadServerParams) (*PhotosGetOwnerPhotoUploadServerResp, error) {
	resp, err := api.Request("photos.getOwnerPhotoUploadServer", p, new(PhotosGetOwnerPhotoUploadServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetOwnerPhotoUploadServerResp), nil
}

// ===============
//  PhotosGetTags
// ===============

// PhotosGetTagsParams параметры метода PhotosGetTags.
type PhotosGetTagsParams struct {
	OwnerID   int
	PhotoID   int
	AccessKey string
}

// PhotosGetTagsResp структура, возвращаемая методом PhotosGetTags.
type PhotosGetTagsResp []struct {
	UserID     int     `json:"user_id"`
	ID         int     `json:"id"`
	PlacerID   int     `json:"placer_id"`
	TaggedName string  `json:"tagged_name"`
	Date       int64   `json:"date"`
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
	X2         float32 `json:"x2"`
	Y2         float32 `json:"y2"`
	Viewed     int     `json:"viewed"`
}

// PhotosGetTags возвращает список отметок на фотографии.
func (api *API) PhotosGetTags(p PhotosGetTagsParams) (*PhotosGetTagsResp, error) {
	resp, err := api.Request("photos.getTags", p, new(PhotosGetTagsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetTagsResp), nil
}

// =======================
//  PhotosGetUploadServer
// =======================

// PhotosGetUploadServerParams параметры метода PhotosGetUploadServer.
type PhotosGetUploadServerParams struct {
	AlbumID int
	GroupID int
}

// PhotosGetUploadServerResp структура, возвращаемая методом PhotosGetUploadServer.
type PhotosGetUploadServerResp struct {
	UploadURL string `json:"upload_url"`
	AlbumID   int    `json:"album_id"`
	UserID    int    `json:"user_id"`
}

// PhotosGetUploadServer возвращает адрес сервера для загрузки фотографий. После успешной загрузки фотография может быть сохранена с помощью метода photos.save.
func (api *API) PhotosGetUploadServer(p PhotosGetUploadServerParams) (*PhotosGetUploadServerResp, error) {
	resp, err := api.Request("photos.getUploadServer", p, new(PhotosGetUploadServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetUploadServerResp), nil
}

// =====================
//  PhotosGetUserPhotos
// =====================

// PhotosGetUserPhotosParams параметры метода PhotosGetUserPhotos.
type PhotosGetUserPhotosParams struct {
	UserID   uint
	Offset   uint
	Count    uint
	Extended bool
	Sort     string
}

// PhotosGetUserPhotosResp структура, возвращаемая методом PhotosGetUserPhotos.
type PhotosGetUserPhotosResp struct {
	Count int     `json:"count"`
	Items []Photo `json:"items"`
}

// PhotosGetUserPhotos возвращает список фотографий, на которых отмечен пользователь.
func (api *API) PhotosGetUserPhotos(p PhotosGetUserPhotosParams) (*PhotosGetUserPhotosResp, error) {
	resp, err := api.Request("photos.getUserPhotos", p, new(PhotosGetUserPhotosResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetUserPhotosResp), nil
}

// ===========================
//  PhotosGetWallUploadServer
// ===========================

// PhotosGetWallUploadServerParams параметры метода PhotosGetWallUploadServer.
type PhotosGetWallUploadServerParams struct {
	GroupID int
}

// PhotosGetWallUploadServerResp структура, возвращаемая методом PhotosGetWallUploadServer.
type PhotosGetWallUploadServerResp struct {
	UploadURL string `json:"upload_url"`
	AlbumID   int    `json:"album_id"`
	UserID    int    `json:"user_id"`
}

// PhotosGetWallUploadServer возвращает адрес сервера для загрузки фотографии на стену пользователя или сообщества. После успешной загрузки вы можете сохранить фотографию, воспользовавшись методом photos.savewallphoto.
func (api *API) PhotosGetWallUploadServer(p PhotosGetWallUploadServerParams) (*PhotosGetWallUploadServerResp, error) {
	resp, err := api.Request("photos.getWallUploadServer", p, new(PhotosGetWallUploadServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosGetWallUploadServerResp), nil
}

// =================
//  PhotosMakeCover
// =================

// PhotosMakeCoverParams параметры метода PhotosMakeCover.
type PhotosMakeCoverParams struct {
	OwnerID int
	PhotoID int
	AlbumID int
}

// PhotosMakeCover делает фотографию обложкой альбома.
func (api *API) PhotosMakeCover(p PhotosMakeCoverParams) (bool, error) {
	resp, err := api.Request("photos.makeCover", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ============
//  PhotosMove
// ============

// PhotosMoveParams параметры метода PhotosMove.
type PhotosMoveParams struct {
	OwnerID       int
	TargetAlbumID int
	PhotoID       int
}

// PhotosMove переносит фотографию из одного альбома в другой.
func (api *API) PhotosMove(p PhotosMoveParams) (bool, error) {
	resp, err := api.Request("photos.move", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ==============
//  PhotosPutTag
// ==============

// PhotosPutTagParams параметры метода PhotosPutTag.
type PhotosPutTagParams struct {
	OwnerID uint
	PhotoID uint
	UserID  int
	X       float32
	Y       float32
	X2      float32
	Y2      float32
}

// PhotosPutTag добавляет отметку на фотографию.
func (api *API) PhotosPutTag(p PhotosPutTagParams) (int, error) {
	resp, err := api.Request("photos.putTag", p, new(int))
	if err != nil {
		return 0, err
	}
	return resp.(int), nil
}

// =================
//  PhotosRemoveTag
// =================

// PhotosRemoveTagParams параметры метода PhotosRemoveTag.
type PhotosRemoveTagParams struct {
	OwnerID int
	PhotoID int
	TagID   int
}

// PhotosRemoveTag удаляет отметку с фотографии.
func (api *API) PhotosRemoveTag(p PhotosRemoveTagParams) (bool, error) {
	resp, err := api.Request("photos.removeTag", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =====================
//  PhotosReorderAlbums
// =====================

// PhotosReorderAlbumsParams параметры метода PhotosReorderAlbums.
type PhotosReorderAlbumsParams struct {
	OwnerID int
	AlbumID int
	Before  int
	After   int
}

// PhotosReorderAlbums меняет порядок альбома в списке альбомов пользователя.
func (api *API) PhotosReorderAlbums(p PhotosReorderAlbumsParams) (bool, error) {
	resp, err := api.Request("photos.reorderAlbums", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =====================
//  PhotosReorderPhotos
// =====================

// PhotosReorderPhotosParams параметры метода PhotosReorderPhotos.
type PhotosReorderPhotosParams struct {
	OwnerID int
	PhotoID int
	Before  int
	After   int
}

// PhotosReorderPhotos меняет порядок фотографии в списке фотографий альбома пользователя.
func (api *API) PhotosReorderPhotos(p PhotosReorderPhotosParams) (bool, error) {
	resp, err := api.Request("photos.reorderPhotos", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ==============
//  PhotosReport
// ==============

// PhotosReportParams параметры метода PhotosReport.
type PhotosReportParams struct {
	OwnerID int
	PhotoID uint
	Reason  uint
}

// PhotosReport позволяет пожаловаться на фотографию.
func (api *API) PhotosReport(p PhotosReportParams) (bool, error) {
	resp, err := api.Request("photos.report", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =====================
//  PhotosReportComment
// =====================

// PhotosReportCommentParams параметры метода PhotosReportComment.
type PhotosReportCommentParams struct {
	OwnerID   int
	CommentID uint
	Reason    uint
}

// PhotosReportComment позволяет пожаловаться на комментарий к фотографии.
func (api *API) PhotosReportComment(p PhotosReportCommentParams) (bool, error) {
	resp, err := api.Request("photos.reportComment", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ===============
//  PhotosRestore
// ===============

// PhotosRestoreParams параметры метода PhotosRestore.
type PhotosRestoreParams struct {
	OwnerID int
	PhotoID uint
}

// PhotosRestore восстанавливает удаленную фотографию.
func (api *API) PhotosRestore(p PhotosRestoreParams) (bool, error) {
	resp, err := api.Request("photos.restore", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ======================
//  PhotosRestoreComment
// ======================

// PhotosRestoreCommentParams параметры метода PhotosRestoreComment.
type PhotosRestoreCommentParams struct {
	OwnerID   int
	CommentID int
}

// PhotosRestoreComment восстанавливает удаленный комментарий к фотографии.
func (api *API) PhotosRestoreComment(p PhotosRestoreCommentParams) (bool, error) {
	resp, err := api.Request("photos.restoreComment", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ============
//  PhotosSave
// ============

// PhotosSaveParams параметры метода PhotosSave.
type PhotosSaveParams struct {
	AlbumID    int
	GroupID    int
	Server     int
	PhotosList string
	Hash       string
	Latitude   float32
	Longitude  float32
	Caption    string
}

// PhotosSave сохраняет фотографии после успешной загрузки.
func (api *API) PhotosSave(p PhotosSaveParams) ([]Photo, error) {
	resp, err := api.Request("photos.save", p, new([]Photo))
	if err != nil {
		return nil, err
	}
	return resp.([]Photo), nil
}

// ============================
//  PhotosSaveMarketAlbumPhoto
// ============================

// PhotosSaveMarketAlbumPhotoParams параметры метода PhotosSaveMarketAlbumPhoto.
type PhotosSaveMarketAlbumPhotoParams struct {
	GroupID uint
	Photo   string
	Server  uint
	Hash    string
}

// PhotosSaveMarketAlbumPhoto сохраняет фотографии после успешной загрузки на URI, полученный методом photos.getmarketalbumuploadserver.
func (api *API) PhotosSaveMarketAlbumPhoto(p PhotosSaveMarketAlbumPhotoParams) ([]Photo, error) {
	resp, err := api.Request("photos.saveMarketAlbumPhoto", p, new([]Photo))
	if err != nil {
		return nil, err
	}
	return resp.([]Photo), nil
}

// =======================
//  PhotosSaveMarketPhoto
// =======================

// PhotosSaveMarketPhotoParams параметры метода PhotosSaveMarketPhoto.
type PhotosSaveMarketPhotoParams struct {
	GroupID  uint
	Photo    string
	Server   int
	Hash     string
	CropData string
	CropHash string
}

// PhotosSaveMarketPhoto сохраняет фотографии после успешной загрузки на URI, полученный методом photos.getmarketuploadserver.
func (api *API) PhotosSaveMarketPhoto(p PhotosSaveMarketPhotoParams) ([]Photo, error) {
	resp, err := api.Request("photos.saveMarketPhoto", p, new([]Photo))
	if err != nil {
		return nil, err
	}
	return resp.([]Photo), nil
}

// =========================
//  PhotosSaveMessagesPhoto
// =========================

// PhotosSaveMessagesPhotoParams параметры метода PhotosSaveMessagesPhoto.
type PhotosSaveMessagesPhotoParams struct {
	Photo  string
	Server int
	Hash   string
}

// PhotosSaveMessagesPhotoResp структура, возвращаемая методом PhotosSaveMessagesPhoto.
type PhotosSaveMessagesPhotoResp []struct {
	ID       int    `json:"id"`
	PID      int    `json:"pid"`
	AID      int    `json:"aid"`
	OwnerID  int    `json:"owner_id"`
	Src      string `json:"src"`
	SrcBig   string `json:"src_big"`
	SrcSmall string `json:"src_small"`
	SrcXbig  string `json:"src_xbig"`
	SrcXxbig string `json:"src_xxbig"`
	Created  int64  `json:"created"`
}

// PhotosSaveMessagesPhoto сохраняет фотографию после успешной загрузки на URI, полученный методом photos.getmessagesuploadserver.
func (api *API) PhotosSaveMessagesPhoto(p PhotosSaveMessagesPhotoParams) (*PhotosSaveMessagesPhotoResp, error) {
	resp, err := api.Request("photos.saveMessagesPhoto", p, new(PhotosSaveMessagesPhotoResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosSaveMessagesPhotoResp), nil
}

// ===========================
//  PhotosSaveOwnerCoverPhoto
// ===========================

// PhotosSaveOwnerCoverPhotoParams параметры метода PhotosSaveOwnerCoverPhoto.
type PhotosSaveOwnerCoverPhotoParams struct {
	Hash  string
	Photo string
}

// PhotosSaveOwnerCoverPhotoResp структура, возвращаемая методом PhotosSaveOwnerCoverPhoto.
type PhotosSaveOwnerCoverPhotoResp []struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// PhotosSaveOwnerCoverPhoto сохраняет изображение для обложки сообщества после успешной загрузки.
func (api *API) PhotosSaveOwnerCoverPhoto(p PhotosSaveOwnerCoverPhotoParams) (*PhotosSaveOwnerCoverPhotoResp, error) {
	resp, err := api.Request("photos.saveOwnerCoverPhoto", p, new(PhotosSaveOwnerCoverPhotoResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosSaveOwnerCoverPhotoResp), nil
}

// ======================
//  PhotosSaveOwnerPhoto
// ======================

// PhotosSaveOwnerPhotoParams параметры метода PhotosSaveOwnerPhoto.
type PhotosSaveOwnerPhotoParams struct {
	Server string
	Hash   string
	Photo  string
}

// PhotosSaveOwnerPhotoResp структура, возвращаемая методом PhotosSaveOwnerPhoto.
type PhotosSaveOwnerPhotoResp struct {
	PhotoHash     string `json:"photo_hash"`
	PhotoSrc      string `json:"photo_src"`
	PhotoSrcBig   string `json:"photo_src_big"`
	PhotoSrcSmall string `json:"photo_src_small"`
}

// PhotosSaveOwnerPhoto позволяет сохранить главную фотографию пользователя или сообщества. Адрес для загрузки фотографии вы можете получить с помощью метода photos.getownerphotouploadserver.
func (api *API) PhotosSaveOwnerPhoto(p PhotosSaveOwnerPhotoParams) (*PhotosSaveOwnerPhotoResp, error) {
	resp, err := api.Request("photos.saveOwnerPhoto", p, new(PhotosSaveOwnerPhotoResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosSaveOwnerPhotoResp), nil
}

// =====================
//  PhotosSaveWallPhoto
// =====================

// PhotosSaveWallPhotoParams параметры метода PhotosSaveWallPhoto.
type PhotosSaveWallPhotoParams struct {
	UserID    uint
	GroupID   uint
	Photo     string
	Server    int
	Hash      string
	Latitude  float32
	Longitude float32
	Caption   string
}

// PhotosSaveWallPhoto сохраняет фотографии после успешной загрузки на URI, полученный методом photos.getwalluploadserver.
func (api *API) PhotosSaveWallPhoto(p PhotosSaveWallPhotoParams) ([]Photo, error) {
	resp, err := api.Request("photos.saveWallPhoto", p, new([]Photo))
	if err != nil {
		return nil, err
	}
	return resp.([]Photo), nil
}

// ==============
//  PhotosSearch
// ==============

// PhotosSearchParams параметры метода PhotosSearch.
type PhotosSearchParams struct {
	Q         string
	Lat       float32
	Long      float32
	StartTime int64
	EndTime   int64
	Sort      bool
	Offset    uint
	Count     uint
	Radius    uint
}

// PhotosSearchResp структура, возвращаемая методом PhotosSearch.
type PhotosSearchResp struct {
	Count int     `json:"count"`
	Items []Photo `json:"items"`
}

// PhotosSearch осуществляет поиск изображений по местоположению или описанию.
func (api *API) PhotosSearch(p PhotosSearchParams) (*PhotosSearchResp, error) {
	resp, err := api.Request("photos.search", p, new(PhotosSearchResp))
	if err != nil {
		return nil, err
	}
	return resp.(*PhotosSearchResp), nil
}
