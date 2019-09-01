package vkapi

import (
	"math/rand"
)

// =====================
//  MessagesAddChatUser
// =====================

// MessagesAddChatUserParams параметры метода MessagesAddChatUser.
type MessagesAddChatUserParams struct {
	ChatID uint
	UserID uint
}

// MessagesAddChatUser добавляет в мультидиалог нового пользователя. чтобы пользователю вернуться в беседу, которую он сам покинул, достаточно отправить сообщение в неё (если есть свободные места), либо вызвать этот метод, передав свой идентификатор в параметре user_id.
func (api *API) MessagesAddChatUser(p MessagesAddChatUserParams) (bool, error) {
	resp, err := api.Request("messages.addChatUser", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ================================
//  MessagesAllowMessagesFromGroup
// ================================

// MessagesAllowMessagesFromGroupParams параметры метода MessagesAllowMessagesFromGroup.
type MessagesAllowMessagesFromGroupParams struct {
	GroupID uint
	Key     string
}

// MessagesAllowMessagesFromGroup позволяет разрешить отправку сообщений от сообщества текущему пользователю.
func (api *API) MessagesAllowMessagesFromGroup(p MessagesAllowMessagesFromGroupParams) (bool, error) {
	resp, err := api.Request("messages.allowMessagesFromGroup", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ====================
//  MessagesCreateChat
// ====================

// MessagesCreateChatParams параметры метода MessagesCreateChat.
type MessagesCreateChatParams struct {
	UserIDS []int
	Title   string
}

// MessagesCreateChatResp структура, возвращаемая методом MessagesCreateChat.
type MessagesCreateChatResp struct {
	ChatID int `json:"chat_id"`
}

// MessagesCreateChat создаёт беседу с несколькими участниками.
func (api *API) MessagesCreateChat(p MessagesCreateChatParams) (int, error) {
	resp, err := api.Request("messages.createChat", p, new(int))
	if err != nil {
		return 0, err
	}
	return resp.(int), nil
}

// ================
//  MessagesDelete
// ================

// MessagesDeleteParams параметры метода MessagesDelete.
type MessagesDeleteParams struct {
	MessageIDS   []int
	Spam         bool
	GroupID      uint
	DeleteForAll bool
}

// MessagesDelete удаляет сообщение.
func (api *API) MessagesDelete(p MessagesDeleteParams) (map[string]bool, error) {
	resp, err := api.Request("messages.delete", p, new(interface{}))
	if err != nil {
		return nil, err
	}

	temp := resp.(map[string]interface{})
	m := make(map[string]bool, len(temp))
	for k, v := range temp {
		m[k] = toBool(int(v.(float64)))
	}

	return m, nil
}

// =========================
//  MessagesDeleteChatPhoto
// =========================

// MessagesDeleteChatPhotoParams параметры метода MessagesDeleteChatPhoto.
type MessagesDeleteChatPhotoParams struct {
	ChatID  uint
	GroupID uint
}

// MessagesDeleteChatPhotoResp структура, возвращаемая методом MessagesDeleteChatPhoto.
type MessagesDeleteChatPhotoResp struct {
	MessageID int   `json:"message_id"`
	Chat      *Chat `json:"chat"`
}

// MessagesDeleteChatPhoto позволяет удалить фотографию мультидиалога.
func (api *API) MessagesDeleteChatPhoto(p MessagesDeleteChatPhotoParams) (*MessagesDeleteChatPhotoResp, error) {
	resp, err := api.Request("messages.deleteChatPhoto", p, new(MessagesDeleteChatPhotoResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesDeleteChatPhotoResp), nil
}

// ============================
//  MessagesDeleteConversation
// ============================

// MessagesDeleteConversationParams параметры метода MessagesDeleteConversation.
type MessagesDeleteConversationParams struct {
	UserID  string
	PeerID  int
	GroupID uint
}

// MessagesDeleteConversationResp структура, возвращаемая методом MessagesDeleteConversation.
type MessagesDeleteConversationResp struct {
	LastDeletedID int `json:"last_deleted_id"`
}

// MessagesDeleteConversation удаляет беседу.
func (api *API) MessagesDeleteConversation(p MessagesDeleteConversationParams) (*MessagesDeleteConversationResp, error) {
	resp, err := api.Request("messages.deleteConversation", p, new(MessagesDeleteConversationResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesDeleteConversationResp), nil
}

// ===============================
//  MessagesDenyMessagesFromGroup
// ===============================

// MessagesDenyMessagesFromGroupParams параметры метода MessagesDenyMessagesFromGroup.
type MessagesDenyMessagesFromGroupParams struct {
	GroupID uint
}

// MessagesDenyMessagesFromGroup позволяет запретить отправку сообщений от сообщества текущему пользователю.
func (api *API) MessagesDenyMessagesFromGroup(p MessagesDenyMessagesFromGroupParams) (bool, error) {
	resp, err := api.Request("messages.denyMessagesFromGroup", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ==============
//  MessagesEdit
// ==============

// MessagesEditParams параметры метода MessagesEdit.
type MessagesEditParams struct {
	PeerID              int
	Message             string
	MessageID           uint
	Lat                 float32
	Long                float32
	Attachment          string
	KeepForwardMessages bool
	KeepSnippets        bool
	GroupID             uint
	DontParseLinks      bool
}

// MessagesEdit редактирует сообщение.
func (api *API) MessagesEdit(p MessagesEditParams) (bool, error) {
	resp, err := api.Request("messages.edit", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ==================
//  MessagesEditChat
// ==================

// MessagesEditChatParams параметры метода MessagesEditChat.
type MessagesEditChatParams struct {
	ChatID uint
	Title  string
}

// MessagesEditChat изменяет название беседы.
func (api *API) MessagesEditChat(p MessagesEditChatParams) (bool, error) {
	resp, err := api.Request("messages.editChat", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ====================================
//  MessagesGetByConversationMessageID
// ====================================

// MessagesGetByConversationMessageIDParams параметры метода MessagesGetByConversationMessageID.
type MessagesGetByConversationMessageIDParams struct {
	PeerID                 int
	ConversationMessageIDS []int
	Extended               bool
	Fields                 string
	GroupID                uint
}

// MessagesGetByConversationMessageIDResp структура, возвращаемая методом MessagesGetByConversationMessageID.
type MessagesGetByConversationMessageIDResp struct {
	Count int       `json:"count"`
	Items []Message `json:"items"`
}

// MessagesGetByConversationMessageID возвращает сообщения по их идентификаторам в рамках беседы или диалога.
func (api *API) MessagesGetByConversationMessageID(p MessagesGetByConversationMessageIDParams) (*MessagesGetByConversationMessageIDResp, error) {
	resp, err := api.Request("messages.getByConversationMessageId", p, new(MessagesGetByConversationMessageIDResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetByConversationMessageIDResp), nil
}

// =================
//  MessagesGetByID
// =================

// MessagesGetByIDParams параметры метода MessagesGetByID.
type MessagesGetByIDParams struct {
	MessageIDS    []int
	PreviewLength uint
	Extended      bool
	Fields        string
	GroupID       uint
}

// MessagesGetByIDResp структура, возвращаемая методом MessagesGetByID.
type MessagesGetByIDResp struct {
	Count int       `json:"count"`
	Items []Message `json:"items"`
}

// MessagesGetByID возвращает сообщения по их идентификаторам.
func (api *API) MessagesGetByID(p MessagesGetByIDParams) (*MessagesGetByIDResp, error) {
	resp, err := api.Request("messages.getById", p, new(MessagesGetByIDResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetByIDResp), nil
}

// =================
//  MessagesGetChat
// =================

// MessagesGetChatParams параметры метода MessagesGetChat.
type MessagesGetChatParams struct {
	ChatID   uint
	ChatIDS  []int
	Fields   string
	NameCase string
}

// MessagesGetChat возвращает информацию о беседе.
func (api *API) MessagesGetChat(p MessagesGetChatParams) ([]Chat, error) {
	if p.ChatID != 0 {
		p.ChatIDS = append(p.ChatIDS, int(p.ChatID))
		p.ChatID = 0
	}

	resp, err := api.Request("messages.getChat", p, new([]Chat))
	if err != nil {
		return nil, err
	}
	return resp.([]Chat), nil
}

// ========================
//  MessagesGetChatPreview
// ========================

// MessagesGetChatPreviewParams параметры метода MessagesGetChatPreview.
type MessagesGetChatPreviewParams struct {
	PeerID uint
	Link   string
	Fields string
}

// MessagesGetChatPreviewResp структура, возвращаемая методом MessagesGetChatPreview.
type MessagesGetChatPreviewResp struct {
	Preview *struct {
		AdminID int    `json:"admin_id"`
		Members []int  `json:"members"`
		Title   string `json:"title"`
		Photo   *struct {
			Photo50  string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		} `json:"photo"`
		LocalID int `json:"local_id"`
	} `json:"preview"`
	Profiles []User  `json:"profiles"`
	Groups   []Group `json:"groups"`
	Emails   []struct {
		ID      int    `json:"id"`
		Address string `json:"address"`
	} `json:"emails"`
}

// MessagesGetChatPreview получает данные для превью чата с приглашением по ссылке.
func (api *API) MessagesGetChatPreview(p MessagesGetChatPreviewParams) (*MessagesGetChatPreviewResp, error) {
	resp, err := api.Request("messages.getChatPreview", p, new(MessagesGetChatPreviewResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetChatPreviewResp), nil
}

// ================================
//  MessagesGetConversationMembers
// ================================

// MessagesGetConversationMembersParams параметры метода MessagesGetConversationMembers.
type MessagesGetConversationMembersParams struct {
	PeerID  int
	Fields  string
	GroupID uint
}

// MessagesGetConversationMembersResp структура, возвращаемая методом MessagesGetConversationMembers.
type MessagesGetConversationMembersResp struct {
	Count int `json:"count"`
	Items []struct {
		MemberID  int   `json:"member_id"`
		InvitedBy int   `json:"invited_by"`
		JoinDate  int64 `json:"join_date"`
		IsAdmin   bool  `json:"is_admin"`
		CanKick   bool  `json:"can_kick"`
	} `json:"items"`
	Profiles []User  `json:"profiles"`
	Groups   []Group `json:"groups"`
}

// MessagesGetConversationMembers позволяет получить список участников беседы.
func (api *API) MessagesGetConversationMembers(p MessagesGetConversationMembersParams) (*MessagesGetConversationMembersResp, error) {
	resp, err := api.Request("messages.getConversationMembers", p, new(MessagesGetConversationMembersResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetConversationMembersResp), nil
}

// ==========================
//  MessagesGetConversations
// ==========================

// MessagesGetConversationsParams параметры метода MessagesGetConversations.
type MessagesGetConversationsParams struct {
	Offset         uint
	Count          uint
	Filter         string
	Extended       bool
	StartMessageID uint
	Fields         string
	GroupID        uint
}

// MessagesGetConversationsResp структура, возвращаемая методом MessagesGetConversations.
type MessagesGetConversationsResp struct {
	Count int `json:"count"`
	Items []struct {
		Conversation *Conversation `json:"conversation"`
		LastMessage  *Message      `json:"last_message"`
	} `json:"items"`
	UnreadCount int     `json:"unread_count"`
	Profiles    []User  `json:"profiles"`
	Groups      []Group `json:"groups"`
}

// MessagesGetConversations возвращает список бесед пользователя.
func (api *API) MessagesGetConversations(p MessagesGetConversationsParams) (*MessagesGetConversationsResp, error) {
	resp, err := api.Request("messages.getConversations", p, new(MessagesGetConversationsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetConversationsResp), nil
}

// ==============================
//  MessagesGetConversationsByID
// ==============================

// MessagesGetConversationsByIDParams параметры метода MessagesGetConversationsByID.
type MessagesGetConversationsByIDParams struct {
	PeerIDS  []int
	Extended bool
	Fields   string
	GroupID  uint
}

// MessagesGetConversationsByIDResp структура, возвращаемая методом MessagesGetConversationsByID.
type MessagesGetConversationsByIDResp struct {
	Count int            `json:"count"`
	Items []Conversation `json:"items"`
}

// MessagesGetConversationsByID позволяет получить беседу по её идентификатору.
func (api *API) MessagesGetConversationsByID(p MessagesGetConversationsByIDParams) (*MessagesGetConversationsByIDResp, error) {
	resp, err := api.Request("messages.getConversationsById", p, new(MessagesGetConversationsByIDResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetConversationsByIDResp), nil
}

// ====================
//  MessagesGetHistory
// ====================

// MessagesGetHistoryParams параметры метода MessagesGetHistory.
type MessagesGetHistoryParams struct {
	Offset         int
	Count          uint
	UserID         int
	PeerID         int
	StartMessageID int
	Rev            bool
	Extended       bool
	Fields         string
	GroupID        uint
}

// MessagesGetHistoryResp структура, возвращаемая методом MessagesGetHistory.
type MessagesGetHistoryResp struct {
	Count   int       `json:"count"`
	Unread  int       `json:"unread"`
	Skipped int       `json:"skipped"`
	Items   []Message `json:"items"`
	InRead  int       `json:"in_read"`
	OutRead int       `json:"out_read"`
}

// MessagesGetHistory возвращает историю сообщений для указанного диалога.
func (api *API) MessagesGetHistory(p MessagesGetHistoryParams) (*MessagesGetHistoryResp, error) {
	resp, err := api.Request("messages.getHistory", p, new(MessagesGetHistoryResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetHistoryResp), nil
}

// ===============================
//  MessagesGetHistoryAttachments
// ===============================

// MessagesGetHistoryAttachmentsParams параметры метода MessagesGetHistoryAttachments.
type MessagesGetHistoryAttachmentsParams struct {
	PeerID           int
	MediaType        string
	StartFrom        string
	Count            uint
	PhotoSizes       bool
	Fields           string
	GroupID          uint
	PreserveOrder    bool
	MaxForwardsLevel uint
}

// MessagesGetHistoryAttachmentsResp структура, возвращаемая методом MessagesGetHistoryAttachments.
type MessagesGetHistoryAttachmentsResp struct {
	Items []struct {
		MessageID  int         `json:"message_id"`
		Attachment *Attachment `json:"attachment"`
	} `json:"items"`
	NextFrom string `json:"next_from"`
}

// MessagesGetHistoryAttachments возвращает материалы диалога или беседы.
func (api *API) MessagesGetHistoryAttachments(p MessagesGetHistoryAttachmentsParams) (*MessagesGetHistoryAttachmentsResp, error) {
	resp, err := api.Request("messages.getHistoryAttachments", p, new(MessagesGetHistoryAttachmentsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetHistoryAttachmentsResp), nil
}

// ==============================
//  MessagesGetImportantMessages
// ==============================

// MessagesGetImportantMessagesParams параметры метода MessagesGetImportantMessages.
type MessagesGetImportantMessagesParams struct {
	Count          uint
	Offset         uint
	StartMessageID uint
	PreviewLength  uint
	Fields         string
	Extended       bool
	GroupID        uint
}

// MessagesGetImportantMessagesResp структура, возвращаемая методом MessagesGetImportantMessages.
type MessagesGetImportantMessagesResp struct {
	Messages *struct {
		Count int       `json:"count"`
		Items []Message `json:"items"`
	} `json:"messages"`
}

// MessagesGetImportantMessages возвращает список важных сообщений пользователя.
func (api *API) MessagesGetImportantMessages(p MessagesGetImportantMessagesParams) (*MessagesGetImportantMessagesResp, error) {
	resp, err := api.Request("messages.getImportantMessages", p, new(MessagesGetImportantMessagesResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetImportantMessagesResp), nil
}

// =======================
//  MessagesGetInviteLink
// =======================

// MessagesGetInviteLinkParams параметры метода MessagesGetInviteLink.
type MessagesGetInviteLinkParams struct {
	PeerID  uint
	Reset   bool
	GroupID uint
}

// MessagesGetInviteLinkResp структура, возвращаемая методом MessagesGetInviteLink.
type MessagesGetInviteLinkResp struct {
	Link string `json:"link"`
}

// MessagesGetInviteLink получает ссылку для приглашения пользователя в беседу. Только создатель беседы имеет доступ к ссылке на беседу.
func (api *API) MessagesGetInviteLink(p MessagesGetInviteLinkParams) (*MessagesGetInviteLinkResp, error) {
	resp, err := api.Request("messages.getInviteLink", p, new(MessagesGetInviteLinkResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetInviteLinkResp), nil
}

// =========================
//  MessagesGetLastActivity
// =========================

// MessagesGetLastActivityParams параметры метода MessagesGetLastActivity.
type MessagesGetLastActivityParams struct {
	UserID int
}

// MessagesGetLastActivityResp структура, возвращаемая методом MessagesGetLastActivity.
type MessagesGetLastActivityResp struct {
	Online int   `json:"online"`
	Time   int64 `json:"time"`
}

// MessagesGetLastActivity возвращает текущий статус и дату последней активности указанного пользователя.
func (api *API) MessagesGetLastActivity(p MessagesGetLastActivityParams) (*MessagesGetLastActivityResp, error) {
	resp, err := api.Request("messages.getLastActivity", p, new(MessagesGetLastActivityResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetLastActivityResp), nil
}

// ============================
//  MessagesGetLongPollHistory
// ============================

// MessagesGetLongPollHistoryParams параметры метода MessagesGetLongPollHistory.
type MessagesGetLongPollHistoryParams struct {
	Ts            uint
	Pts           uint
	PreviewLength uint
	Onlines       bool
	Fields        string
	EventsLimit   uint
	MsgsLimit     uint
	MaxMsgID      uint
	GroupID       uint
	LPVersion     uint
	LastN         uint
	Credentials   bool
}

// MessagesGetLongPollHistoryResp структура, возвращаемая методом MessagesGetLongPollHistory.
type MessagesGetLongPollHistoryResp struct {
	History  [][]int `json:"history"`
	Messages *struct {
		Count int       `json:"count"`
		Items []Message `json:"items"`
	} `json:"messages"`
	Profiles []User `json:"profiles"`
	NewPts   int    `json:"new_pts"`
}

// MessagesGetLongPollHistory возвращает обновления в личных сообщениях пользователя. Для ускорения работы с личными сообщениями может быть полезно кешировать уже загруженные ранее сообщения на мобильном устройстве / пк пользователя, чтобы не получать их повторно при каждом обращении. Этот метод помогает осуществить синхронизацию локальной копии списка сообщений с актуальной версией.
func (api *API) MessagesGetLongPollHistory(p MessagesGetLongPollHistoryParams) (*MessagesGetLongPollHistoryResp, error) {
	resp, err := api.Request("messages.getLongPollHistory", p, new(MessagesGetLongPollHistoryResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetLongPollHistoryResp), nil
}

// ===========================
//  MessagesGetLongPollServer
// ===========================

// MessagesGetLongPollServerParams параметры метода MessagesGetLongPollServer.
type MessagesGetLongPollServerParams struct {
	NeedPts   bool
	GroupID   uint
	LPVersion uint
}

// MessagesGetLongPollServerResp структура, возвращаемая методом MessagesGetLongPollServer.
type MessagesGetLongPollServerResp struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     int    `json:"ts"`
}

// MessagesGetLongPollServer возвращает данные, необходимые для подключения к long poll серверу. Long poll позволит вам моментально узнавать о приходе новых сообщений и других событиях.
func (api *API) MessagesGetLongPollServer(p MessagesGetLongPollServerParams) (*MessagesGetLongPollServerResp, error) {
	resp, err := api.Request("messages.getLongPollServer", p, new(MessagesGetLongPollServerResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesGetLongPollServerResp), nil
}

// ====================================
//  MessagesIsMessagesFromGroupAllowed
// ====================================

// MessagesIsMessagesFromGroupAllowedParams параметры метода MessagesIsMessagesFromGroupAllowed.
type MessagesIsMessagesFromGroupAllowedParams struct {
	GroupID uint
	UserID  uint
}

// MessagesIsMessagesFromGroupAllowedResp структура, возвращаемая методом MessagesIsMessagesFromGroupAllowed.
type MessagesIsMessagesFromGroupAllowedResp struct {
	IsAllowed int `json:"is_allowed"`
}

// MessagesIsMessagesFromGroupAllowed возвращает информацию о том, разрешена ли отправка сообщений от сообщества пользователю.
func (api *API) MessagesIsMessagesFromGroupAllowed(p MessagesIsMessagesFromGroupAllowedParams) (*MessagesIsMessagesFromGroupAllowedResp, error) {
	resp, err := api.Request("messages.isMessagesFromGroupAllowed", p, new(MessagesIsMessagesFromGroupAllowedResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesIsMessagesFromGroupAllowedResp), nil
}

// ==============================
//  MessagesJoinChatByInviteLink
// ==============================

// MessagesJoinChatByInviteLinkParams параметры метода MessagesJoinChatByInviteLink.
type MessagesJoinChatByInviteLinkParams struct {
	Link string
}

// MessagesJoinChatByInviteLinkResp структура, возвращаемая методом MessagesJoinChatByInviteLink.
type MessagesJoinChatByInviteLinkResp struct {
	ChatID int `json:"chat_id"`
}

// MessagesJoinChatByInviteLink позволяет присоединиться к чату по ссылке-приглашению.
func (api *API) MessagesJoinChatByInviteLink(p MessagesJoinChatByInviteLinkParams) (*MessagesJoinChatByInviteLinkResp, error) {
	resp, err := api.Request("messages.joinChatByInviteLink", p, new(MessagesJoinChatByInviteLinkResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesJoinChatByInviteLinkResp), nil
}

// ====================================
//  MessagesMarkAsAnsweredConversation
// ====================================

// MessagesMarkAsAnsweredConversationParams параметры метода MessagesMarkAsAnsweredConversation.
type MessagesMarkAsAnsweredConversationParams struct {
	PeerID   int
	Answered bool
	GroupID  uint
}

// MessagesMarkAsAnsweredConversation помечает беседу как отвеченную либо снимает отметку.
func (api *API) MessagesMarkAsAnsweredConversation(p MessagesMarkAsAnsweredConversationParams) (bool, error) {
	resp, err := api.Request("messages.markAsAnsweredConversation", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =========================
//  MessagesMarkAsImportant
// =========================

// MessagesMarkAsImportantParams параметры метода MessagesMarkAsImportant.
type MessagesMarkAsImportantParams struct {
	MessageIDS []int
	Important  uint
}

// MessagesMarkAsImportant помечает сообщения как важные либо снимает отметку.
func (api *API) MessagesMarkAsImportant(p MessagesMarkAsImportantParams) ([]int, error) {
	resp, err := api.Request("messages.markAsImportant", p, new([]int))
	if err != nil {
		return nil, err
	}
	return resp.([]int), nil
}

// =====================================
//  MessagesMarkAsImportantConversation
// =====================================

// MessagesMarkAsImportantConversationParams параметры метода MessagesMarkAsImportantConversation.
type MessagesMarkAsImportantConversationParams struct {
	PeerID    int
	Important bool
	GroupID   uint
}

// MessagesMarkAsImportantConversation помечает беседу как важную либо снимает отметку.
func (api *API) MessagesMarkAsImportantConversation(p MessagesMarkAsImportantConversationParams) (bool, error) {
	resp, err := api.Request("messages.markAsImportantConversation", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ====================
//  MessagesMarkAsRead
// ====================

// MessagesMarkAsReadParams параметры метода MessagesMarkAsRead.
type MessagesMarkAsReadParams struct {
	MessageIDS     []int
	PeerID         int
	StartMessageID uint
	GroupID        uint
}

// MessagesMarkAsRead помечает сообщения как прочитанные.
func (api *API) MessagesMarkAsRead(p MessagesMarkAsReadParams) (bool, error) {
	resp, err := api.Request("messages.markAsRead", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =============
//  MessagesPin
// =============

// MessagesPinParams параметры метода MessagesPin.
type MessagesPinParams struct {
	PeerID    int
	MessageID uint
}

// MessagesPin закрепляет сообщение.
func (api *API) MessagesPin(p MessagesPinParams) (*PinnedMessage, error) {
	resp, err := api.Request("messages.pin", p, new(PinnedMessage))
	if err != nil {
		return nil, err
	}
	return resp.(*PinnedMessage), nil
}

// ========================
//  MessagesRemoveChatUser
// ========================

// MessagesRemoveChatUserParams параметры метода MessagesRemoveChatUser.
type MessagesRemoveChatUserParams struct {
	ChatID   uint
	UserID   int
	MemberID int
}

// MessagesRemoveChatUser исключает из мультидиалога пользователя, если текущий пользователь или сообщество является администратором беседы либо текущий пользователь пригласил исключаемого пользователя. Также может быть использован для выхода текущего пользователя из беседы, в которой он состоит. чтобы пользователю вернуться в беседу, достаточно отправить сообщение в неё (если есть свободные места).
func (api *API) MessagesRemoveChatUser(p MessagesRemoveChatUserParams) (bool, error) {
	resp, err := api.Request("messages.removeChatUser", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// =================
//  MessagesRestore
// =================

// MessagesRestoreParams параметры метода MessagesRestore.
type MessagesRestoreParams struct {
	MessageID uint
	GroupID   uint
}

// MessagesRestore восстанавливает удаленное сообщение.
func (api *API) MessagesRestore(p MessagesRestoreParams) (bool, error) {
	resp, err := api.Request("messages.restore", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ================
//  MessagesSearch
// ================

// MessagesSearchParams параметры метода MessagesSearch.
type MessagesSearchParams struct {
	Q             string
	PeerID        int
	Date          int
	PreviewLength uint
	Offset        uint
	Count         uint
	Extended      bool
	Fields        string
	GroupID       uint
}

// MessagesSearchResp структура, возвращаемая методом MessagesSearch.
type MessagesSearchResp struct {
	Count    int       `json:"count"`
	Items    []Message `json:"items"`
	Profiles []User    `json:"profiles"`
	Groups   []Group   `json:"groups"`
}

// MessagesSearch возвращает список найденных личных сообщений текущего пользователя по введенной строке поиска.
func (api *API) MessagesSearch(p MessagesSearchParams) (*MessagesSearchResp, error) {
	resp, err := api.Request("messages.search", p, new(MessagesSearchResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesSearchResp), nil
}

// =============================
//  MessagesSearchConversations
// =============================

// MessagesSearchConversationsParams параметры метода MessagesSearchConversations.
type MessagesSearchConversationsParams struct {
	Q        string
	Count    uint
	Extended bool
	Fields   string
	GroupID  uint
}

// MessagesSearchConversationsResp структура, возвращаемая методом MessagesSearchConversations.
type MessagesSearchConversationsResp struct {
	Count int            `json:"count"`
	Items []Conversation `json:"items"`
}

// MessagesSearchConversations позволяет искать диалоги.
func (api *API) MessagesSearchConversations(p MessagesSearchConversationsParams) (*MessagesSearchConversationsResp, error) {
	resp, err := api.Request("messages.searchConversations", p, new(MessagesSearchConversationsResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesSearchConversationsResp), nil
}

// ==============
//  MessagesSend
// ==============

// MessagesSendParams параметры метода MessagesSend.
type MessagesSendParams struct {
	UserID          int
	RandomID        uint64
	PeerID          int
	Domain          string
	ChatID          uint
	UserIDS         []int
	Message         string
	Lat             float32
	Long            float32
	Attachment      string
	ReplyTo         int
	ForwardMessages []int
	Forward         interface{}
	StickerID       uint
	GroupID         uint
	Keyboard        *Keyboard
	Payload         string
	DontParseLinks  bool
	DisableMentions bool
}

// MessagesSendResp структура, возвращаемая методом MessagesSend.
type MessagesSendResp []struct {
	PeerID    int `json:"peer_id"`
	MessageID int `json:"message_id"`
	Error     *struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"error"`
}

// MessagesSend отправляет сообщение. Если RandomID не задано, ему автоматически присваивается случайное значение. Возвращает идентификатор отправленного сообщения, если передан UserID или *MessagesSendResp, если передан UserIDS.
func (api *API) MessagesSend(p MessagesSendParams) (interface{}, error) {
	if p.RandomID == 0 {
		p.RandomID = rand.Uint64()
	}

	var holder interface{}
	switch len(p.UserIDS) > 0 {
	case true:
		holder = new(MessagesSendResp)
	default:
		holder = new(int)
	}

	resp, err := api.Request("messages.send", p, holder)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// =====================
//  MessagesSetActivity
// =====================

// MessagesSetActivityParams параметры метода MessagesSetActivity.
type MessagesSetActivityParams struct {
	UserID  string
	Type    string
	PeerID  int
	GroupID uint
}

// MessagesSetActivity изменяет статус набора текста пользователем в диалоге.
func (api *API) MessagesSetActivity(p MessagesSetActivityParams) (bool, error) {
	resp, err := api.Request("messages.setActivity", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}

// ======================
//  MessagesSetChatPhoto
// ======================

// MessagesSetChatPhotoParams параметры метода MessagesSetChatPhoto.
type MessagesSetChatPhotoParams struct {
	File string
}

// MessagesSetChatPhotoResp структура, возвращаемая методом MessagesSetChatPhoto.
type MessagesSetChatPhotoResp struct {
	MessageID int   `json:"message_id"`
	Chat      *Chat `json:"chat"`
}

// MessagesSetChatPhoto позволяет установить фотографию мультидиалога, загруженную с помощью метода photos.getchatuploadserver.
func (api *API) MessagesSetChatPhoto(p MessagesSetChatPhotoParams) (*MessagesSetChatPhotoResp, error) {
	resp, err := api.Request("messages.setChatPhoto", p, new(MessagesSetChatPhotoResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessagesSetChatPhotoResp), nil
}

// ===============
//  MessagesUnpin
// ===============

// MessagesUnpinParams параметры метода MessagesUnpin.
type MessagesUnpinParams struct {
	PeerID  int
	GroupID uint
}

// MessagesUnpin открепляет сообщение.
func (api *API) MessagesUnpin(p MessagesUnpinParams) (bool, error) {
	resp, err := api.Request("messages.unpin", p, new(int))
	if err != nil {
		return false, err
	}
	return toBool(resp.(int)), nil
}
