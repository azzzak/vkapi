package vkapi

// User описывает пользователя. Набор полей может меняться в зависимости от вызванного метода и переданных в нем параметров.
type User struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Deactivated     string `json:"deactivated"`
	IsClosed        bool   `json:"is_closed"`
	CanAccessClosed bool   `json:"can_access_closed"`

	About                  string `json:"about"`
	Activities             string `json:"activities"`
	Bdate                  string `json:"bdate"`
	Blacklisted            int    `json:"blacklisted"`
	BlacklistedByMe        int    `json:"blacklisted_by_me"`
	Books                  string `json:"books"`
	CanPost                int    `json:"can_post"`
	CanSeeAllPosts         int    `json:"can_see_all_posts"`
	CanSeeAudio            int    `json:"can_see_audio"`
	CanSendFriendRequest   int    `json:"can_send_friend_request"`
	CanWritePrivateMessage int    `json:"can_write_private_message"`
	Career                 *struct {
		GroupID   int    `json:"group_id"`
		Company   string `json:"company"`
		CountryID int    `json:"country_id"`
		CityID    int    `json:"city_id"`
		CityName  string `json:"city_name"`
		From      int    `json:"from"`
		Until     int    `json:"until"`
		Position  string `json:"position"`
	} `json:"career"`
	City *struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"city"`
	CommonCount int         `json:"common_count"`
	Connections interface{} `json:"connections"`
	Contacts    *struct {
		MobilePhone string `json:"mobile_phone"`
		HomePhone   string `json:"home_phone"`
	} `json:"contacts"`
	Counters *struct {
		Albums        int `json:"albums"`
		Videos        int `json:"videos"`
		Audios        int `json:"audios"`
		Photos        int `json:"photos"`
		Notes         int `json:"notes"`
		Friends       int `json:"friends"`
		Groups        int `json:"groups"`
		OnlineFriends int `json:"online_friends"`
		MutualFriends int `json:"mutual_friends"`
		UserVideos    int `json:"user_videos"`
		Followers     int `json:"followers"`
		Pages         int `json:"pages"`
	} `json:"counters"`
	Country *struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"country"`
	CropPhoto *struct {
		Photo *Photo `json:"photo"`
		Crop  *struct {
			X  float32 `json:"x"`
			Y  float32 `json:"y"`
			X2 float32 `json:"x2"`
			Y2 float32 `json:"y2"`
		} `json:"crop"`
		Rect *struct {
			X  float32 `json:"x"`
			Y  float32 `json:"y"`
			X2 float32 `json:"x2"`
			Y2 float32 `json:"y2"`
		} `json:"rect"`
	} `json:"crop_photo"`
	Domain    string `json:"domain"`
	Education *struct {
		University     int    `json:"university"`
		UniversityName string `json:"university_name"`
		Faculty        int    `json:"faculty"`
		FacultyName    string `json:"faculty_name"`
		Graduation     int    `json:"graduation"`
	} `json:"education"`
	Exports          interface{} `json:"exports"`
	FirstNameNom     string      `json:"first_name_nom"`
	FirstNameGen     string      `json:"first_name_gen"`
	FirstNameDat     string      `json:"first_name_dat"`
	FirstNameAcc     string      `json:"first_name_acc"`
	FirstNameIns     string      `json:"first_name_ins"`
	FirstNameAbl     string      `json:"first_name_abl"`
	FollowersCount   int         `json:"followers_count"`
	FriendStatus     int         `json:"friend_status"`
	Games            string      `json:"games"`
	HasMobile        int         `json:"has_mobile"`
	HasPhoto         int         `json:"has_photo"`
	HomeTown         string      `json:"home_town"`
	Interests        string      `json:"interests"`
	IsFavorite       int         `json:"is_favorite"`
	IsFriend         int         `json:"is_friend"`
	IsHiddenFromFeed int         `json:"is_hidden_from_feed"`
	LastNameNom      string      `json:"last_name_nom"`
	LastNameGen      string      `json:"last_name_gen"`
	LastNameDat      string      `json:"last_name_dat"`
	LastNameAcc      string      `json:"last_name_acc"`
	LastNameIns      string      `json:"last_name_ins"`
	LastNameAbl      string      `json:"last_name_abl"`
	LastSeen         *struct {
		Time     int64 `json:"time"`
		Platform int   `json:"platform"`
	} `json:"last_seen"`
	Lists      string `json:"lists"`
	MaidenName string `json:"maiden_name"`
	Military   *struct {
		Unit      string `json:"unit"`
		UnitID    int    `json:"unit_id"`
		CountryID int    `json:"country_id"`
		From      int    `json:"from"`
		Until     int    `json:"until"`
	} `json:"military"`
	Movies     string `json:"movies"`
	Music      string `json:"music"`
	Nickname   string `json:"nickname"`
	Occupation *struct {
		Type string `json:"type"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"occupation"`
	Online   int `json:"online"`
	Personal *struct {
		Political  int         `json:"political"`
		Langs      interface{} `json:"langs"`
		Religion   string      `json:"religion"`
		InspiredBy string      `json:"inspired_by"`
		PeopleMain int         `json:"people_main"`
		LifeMain   int         `json:"life_main"`
		Smoking    int         `json:"smoking"`
		Alcohol    int         `json:"alcohol"`
	} `json:"personal"`
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200Orig string `json:"photo_200_orig"`
	Photo200     string `json:"photo_200"`
	Photo400Orig string `json:"photo_400_orig"`
	PhotoID      string `json:"photo_id"`
	PhotoMax     string `json:"photo_max"`
	PhotoMaxOrig string `json:"photo_max_orig"`
	Quotes       string `json:"quotes"`
	Relatives    []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"relatives"`
	Relation int `json:"relation"`
	Schools  []struct {
		ID            int    `json:"id"`
		Country       int    `json:"country"`
		City          int    `json:"city"`
		Name          string `json:"name"`
		YearFrom      int    `json:"year_from"`
		YearTo        int    `json:"year_to"`
		YearGraduated int    `json:"year_graduated"`
		Class         string `json:"class"`
		Speciality    string `json:"speciality"`
		Type          int    `json:"type"`
		TypeStr       string `json:"type_str"`
	} `json:"schools"`
	ScreenName   string `json:"screen_name"`
	Sex          int    `json:"sex"`
	Site         string `json:"site"`
	Status       string `json:"status"`
	Timezone     int    `json:"timezone"`
	Trending     int    `json:"trending"`
	TV           string `json:"tv"`
	Universities []struct {
		ID              int    `json:"id"`
		Country         int    `json:"country"`
		City            int    `json:"city"`
		Name            string `json:"name"`
		Faculty         int    `json:"faculty"`
		FacultyName     string `json:"faculty_name"`
		Chair           int    `json:"chair"`
		ChairName       string `json:"chair_name"`
		Graduation      int    `json:"graduation"`
		EducationForm   string `json:"education_form"`
		EducationStatus string `json:"education_status"`
	} `json:"universities"`
	Verified    int    `json:"verified"`
	WallDefault string `json:"wall_default"`

	// поле может добавить метод FriendsGetByPhones.
	Phone string `json:"phone"`
	// поле может добавить метод FriendsGetSuggestions.
	FoundWith interface{} `json:"found_with"`
	// поле может добавить метод MessagesGetChat.
	InvitedBy interface{} `json:"invited_by"`
}

// Group описывает сообщество. Набор полей может меняться в зависимости от вызванного метода и переданных в нем параметров.
type Group struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ScreenName   string `json:"screen_name"`
	IsClosed     int    `json:"is_closed"`
	Deactivated  string `json:"deactivated"`
	IsAdmin      int    `json:"is_admin"`
	AdminLevel   int    `json:"admin_level"`
	IsMember     int    `json:"is_member"`
	IsAdvertiser int    `json:"is_advertiser"`
	InvitedBy    int    `json:"invited_by"`
	Type         string `json:"type"`
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200     string `json:"photo_200"`

	Activity  string `json:"activity"`
	AgeLimits int    `json:"age_limits"`
	BanInfo   *struct {
		EndDate int64  `json:"end_date"`
		Comment string `json:"comment"`
	} `json:"ban_info"`
	CanCreateTopic int `json:"can_create_topic"`
	CanMessage     int `json:"can_message"`
	CanPost        int `json:"can_post"`
	CanSeeAllPosts int `json:"can_see_all_posts"`
	CanUploadDoc   int `json:"can_upload_doc"`
	CanUploadVideo int `json:"can_upload_video"`
	City           *struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"city"`
	Contacts *struct {
		UserID int    `json:"user_id"`
		Desc   string `json:"desc"`
		Phone  string `json:"phone"`
		Email  string `json:"email"`
	} `json:"contacts"`
	Counters *struct {
		Photos int `json:"photos"`
		Albums int `json:"albums"`
		Audios int `json:"audios"`
		Videos int `json:"videos"`
		Topics int `json:"topics"`
		Docs   int `json:"docs"`
	} `json:"counters"`
	Country *struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"country"`
	Cover *struct {
		Enabled int `json:"enabled"`
		Images  []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"images"`
	} `json:"cover"`
	CropPhoto *struct {
		Photo *Photo `json:"photo"`
		Crop  *struct {
			X  float32 `json:"x"`
			Y  float32 `json:"y"`
			X2 float32 `json:"x2"`
			Y2 float32 `json:"y2"`
		} `json:"crop"`
		Rect *struct {
			X  float32 `json:"x"`
			Y  float32 `json:"y"`
			X2 float32 `json:"x2"`
			Y2 float32 `json:"y2"`
		} `json:"rect"`
	} `json:"crop_photo"`
	Description       string `json:"description"`
	FixedPost         int    `json:"fixed_post"`
	HasPhoto          int    `json:"has_photo"`
	IsFavorite        int    `json:"is_favorite"`
	IsHiddenFromFeed  int    `json:"is_hidden_from_feed"`
	IsMessagesBlocked int    `json:"is_messages_blocked"`
	Links             *struct {
		ID       int    `json:"id"`
		URL      string `json:"url"`
		Name     string `json:"name"`
		Desc     string `json:"desc"`
		Photo50  string `json:"photo_50"`
		Photo100 string `json:"photo_100"`
	} `json:"links"`
	MainAlbumID int `json:"main_album_id"`
	MainSection int `json:"main_section"`
	Market      *struct {
		Enabled     int `json:"enabled"`
		PriceMin    int `json:"price_min"`
		PriceMax    int `json:"price_max"`
		MainAlbumID int `json:"main_album_id"`
		ContactID   int `json:"contact_id"`
		Currency    *struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"currency"`
		CurrencyText string `json:"currency_text"`
	} `json:"market"`
	MemberStatus int `json:"member_status"`
	MembersCount int `json:"members_count"`
	Place        *struct {
		ID        int     `json:"id"`
		Title     string  `json:"title"`
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
		Type      int     `json:"type"`
		Country   int     `json:"country"`
		City      int     `json:"city"`
		Address   string  `json:"address"`
	} `json:"place"`
	PublicDateLabel string      `json:"public_date_label"`
	Site            string      `json:"site"`
	StartDate       interface{} `json:"start_date"`
	FinishDate      int64       `json:"finish_date"`
	Status          string      `json:"status"`
	Trending        int         `json:"trending"`
	Verified        int         `json:"verified"`
	Wall            int         `json:"wall"`
	WikiPage        string      `json:"wiki_page"`
}

// Post описывает запись на стене пользователя или сообщества.
type Post struct {
	ID           int
	OwnerID      int
	FromID       int
	CreatedBy    int
	Date         int64  `json:"date"`
	Text         string `json:"text"`
	ReplyOwnerID int    `json:"reply_owner_id"`
	ReplyPostID  int    `json:"reply_post_id"`
	FriendsOnly  int    `json:"friends_only"`
	Comments     *struct {
		Count         int  `json:"count"`
		CanPost       int  `json:"can_post"`
		GroupsCanPost bool `json:"groups_can_post"`
		CanClose      int  `json:"can_close"`
		CanOpen       int  `json:"can_open"`
	} `json:"comments"`
	Likes *struct {
		Count      int `json:"count"`
		UserLikes  int `json:"user_likes"`
		CanLike    int `json:"can_like"`
		CanPublish int `json:"can_publish"`
	} `json:"likes"`
	Reposts *struct {
		Count        int `json:"count"`
		UserReposted int `json:"user_reposted"`
	} `json:"reposts"`
	Views *struct {
		Count int `json:"count"`
	} `json:"views"`
	PostType   string `json:"post_type"`
	PostSource *struct {
		Type     string `json:"type"`
		Platform string `json:"platform"`
		Data     string `json:"data"`
		URL      string `json:"url"`
	} `json:"post_source"`
	Attachments []Attachment `json:"attachments"`
	Geo         *struct {
		Type        string      `json:"type"`
		Coordinates interface{} `json:"coordinates"`
	} `json:"geo"`
	SignerID    int         `json:"signer_id"`
	CopyHistory interface{} `json:"copy_history"`
	CanPin      int         `json:"can_pin"`
	IsPinned    int         `json:"is_pinned"`
	CanDelete   int         `json:"can_delete"`
	CanEdit     int         `json:"can_edit"`
	CanArchive  bool        `json:"can_archive"`
	IsArchived  bool        `json:"is_archived"`
	MarkedAsAds int         `json:"marked_as_ads"`
	IsFavorite  bool        `json:"is_favorite"`
	PostponedID int         `json:"postponed_id"`
}

// Comment описывает комментарий к записи.
type Comment struct {
	ID             int          `json:"id"`
	FromID         int          `json:"from_id"`
	CanEdit        int          `json:"can_edit"`
	PostID         int          `json:"post_id"`
	OwnerID        int          `json:"owner_id"`
	Date           int64        `json:"date"`
	Text           string       `json:"text"`
	ReplyToUser    int          `json:"reply_to_user"`
	ReplyToComment int          `json:"reply_to_comment"`
	Attachments    []Attachment `json:"attachments"`
	ParentsStack   []int        `json:"parents_stack"`
	Thread         *struct {
		Count           int       `json:"count"`
		Items           []Comment `json:"items"`
		CanPost         bool      `json:"can_post"`
		ShowReplyButton bool      `json:"show_reply_button"`
		GroupsCanPost   bool      `json:"groups_can_post"`
	} `json:"thread"`
	Deleted bool `json:"deleted"`
	// поле для методов с параметром need_likes=true
	Likes *struct {
		Count     int `json:"count"`
		UserLikes int `json:"user_likes"`
		CanLike   int `json:"can_like"`
	} `json:"likes"`
	// поле для метода photos.getAllComments.
	PID int `json:"pid,omitempty"`
}

// Message описывает личное сообщение.
type Message struct {
	ID           int          `json:"id"`
	Date         int64        `json:"date"`
	PeerID       int          `json:"peer_id"`
	FromID       int          `json:"from_id"`
	Text         string       `json:"text"`
	RandomID     int64        `json:"random_id"`
	Ref          string       `json:"ref"`
	RefSource    string       `json:"ref_source"`
	Attachments  []Attachment `json:"attachments"`
	Important    bool         `json:"important"`
	Geo          *Geo         `json:"geo"`
	Payload      string       `json:"payload"`
	FwdMessages  []Message    `json:"fwd_messages"`
	ReplyMessage *Message     `json:"reply_message"`
	Action       *struct {
		Type     string `json:"type"`
		MemberID int    `json:"member_id"`
		Text     string `json:"text"`
		Email    string `json:"email"`
		Photo    *struct {
			Photo50  string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		}
	} `json:"action"`

	ConversationMessageID int  `json:"conversation_message_id"`
	IsHidden              bool `json:"is_hidden"`
	Out                   int  `json:"out"`
}

// PinnedMessage описывает закреплённое сообщение в беседе.
type PinnedMessage struct {
	ID          int         `json:"id"`
	Date        int64       `json:"date"`
	FromID      int         `json:"from_id"`
	Text        string      `json:"text"`
	Attachments string      `json:"attachments"`
	Geo         *Geo        `json:"geo"`
	FwdMessages interface{} `json:"fwd_messages"`
}

// Conversation описывает беседу с пользователем, сообществом или групповой чат.
type Conversation struct {
	Peer *struct {
		ID      int    `json:"id"`
		Type    string `json:"type"`
		LocalID int    `json:"local_id"`
	} `json:"peer"`
	InRead       int  `json:"in_read"`
	OutRead      int  `json:"out_read"`
	UnreadCount  int  `json:"unread_count"`
	Important    bool `json:"important"`
	Unanswered   bool `json:"unanswered"`
	PushSettings *struct {
		DisabledUntil   int64 `json:"disabled_until"`
		DisabledForever bool  `json:"disabled_forever"`
		NoSound         bool  `json:"no_sound"`
	} `json:"push_settings"`
	CanWrite *struct {
		Allowed bool `json:"allowed"`
		Reason  int  `json:"reason"`
	} `json:"can_write"`
	ChatSettings *struct {
		MembersCount  int            `json:"members_count"`
		Title         string         `json:"title"`
		PinnedMessage *PinnedMessage `json:"pinned_message"`
		State         string         `json:"state"`
		Photo         *struct {
			Photo50  string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		} `json:"photo"`
		ActiveIDS      []int `json:"active_ids"`
		IsGroupChannel bool  `json:"is_group_channel"`
	} `json:"chat_settings"`
}

// Chat описывает чат.
type Chat struct {
	ID           int    `json:"id"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	AdminID      int    `json:"admin_id"`
	Users        []int  `json:"users"`
	PushSettings *struct {
		Sound         int `json:"sound"`
		DisabledUntil int `json:"disabled_until"`
	} `json:"push_settings"`
	Photo50  string `json:"photo_50"`
	Photo100 string `json:"photo_100"`
	Photo200 string `json:"photo_200"`
	Left     int    `json:"left"`
	Kicked   int    `json:"kicked"`
}

// Note описывает заметку.
type Note struct {
	ID           int    `json:"id"`
	OwnerID      int    `json:"owner_id"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	Date         int64  `json:"date"`
	Comments     int    `json:"comments"`
	ReadComments int    `json:"read_comments"`
	ViewURL      string `json:"view_url"`
}

// Page описывает вики-страницу.
type Page struct {
	ID                       int    `json:"id"`
	GroupID                  int    `json:"group_id"`
	CreatorID                int    `json:"creator_id"`
	Title                    string `json:"title"`
	CurrentUserCanEdit       int    `json:"current_user_can_edit"`
	CurrentUserCanEditAccess int    `json:"current_user_can_edit_access"`
	WhoCanView               int    `json:"who_can_view"`
	WhoCanEdit               int    `json:"who_can_edit"`
	Edited                   int64  `json:"edited"`
	Created                  int64  `json:"created"`
	EditorID                 int    `json:"editor_id"`
	Views                    int    `json:"views"`
	Parent                   string `json:"parent"`
	Parent2                  string `json:"parent2"`
	Source                   string `json:"source"`
	HTML                     string `json:"html"`
	ViewURL                  string `json:"view_url"`
}

// MarketItem описывает товар.
type MarketItem struct {
	ID          int    `json:"id"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       *struct {
		Amount   int `json:"amount"`
		Currency *struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"currency"`
		Text string `json:"text"`
	} `json:"price"`
	Category *struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Section *struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"section"`
	} `json:"category"`
	ThumbPhoto   string `json:"thumb_photo"`
	Date         int64  `json:"date"`
	Availability int    `json:"availability"`
	IsFavorite   bool   `json:"is_favorite"`

	Photos     []Photo `json:"photos"`
	CanComment int     `json:"can_comment"`
	CanRepost  int     `json:"can_repost"`
	Likes      *struct {
		UserLikes int `json:"user_likes"`
		Count     int `json:"count"`
	} `json:"likes"`
	URL         string `json:"url"`
	ButtonTitle string `json:"button_title"`
}

// MarketAlbum описывает подборку товаров.
type MarketAlbum struct {
	ID          int    `json:"id"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Photo       *Photo `json:"photo"`
	Count       int    `json:"count"`
	UpdatedTime int64  `json:"updated_time"`
}

// Topic описывает обсуждение в сообществе.
type Topic struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Created      int64  `json:"created"`
	CreatedBy    int    `json:"created_by"`
	Updated      int64  `json:"updated"`
	UpdatedBy    int    `json:"updated_by"`
	IsClosed     int    `json:"is_closed"`
	IsFixed      int    `json:"is_fixed"`
	Comments     int    `json:"comments"`
	FirstComment string `json:"first_comment"`
	LastComment  string `json:"last_comment"`
}

// CommentBoard описывает комментарий из обсуждения в сообществе.
type CommentBoard struct {
	ID          int          `json:"id"`
	FromID      int          `json:"from_id"`
	Date        int64        `json:"date"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
	Likes       *struct {
		Count     int `json:"count"`
		UserLikes int `json:"user_likes"`
		CanLike   int `json:"can_like"`
	} `json:"likes"`
}

// App описывает приложение.
type App struct {
	ID                    int    `json:"id"`
	Title                 string `json:"title"`
	Icon278               string `json:"icon_278"`
	Icon139               string `json:"icon_139"`
	Icon150               string `json:"icon_150"`
	Icon75                string `json:"icon_75"`
	Banner560             string `json:"banner_560"`
	Banner1120            string `json:"banner_1120"`
	Type                  string `json:"type"`
	Section               string `json:"section"`
	AuthorURL             string `json:"author_url"`
	AuthorID              int    `json:"author_id"`
	AuthorGroup           int    `json:"author_group"`
	MembersCount          int    `json:"members_count"`
	PublishedDate         int64  `json:"published_date"`
	CatalogPosition       int    `json:"catalog_position"`
	International         int    `json:"international"`
	LeaderboardType       int    `json:"leaderboard_type"`
	GenreID               int    `json:"genre_id"`
	Genre                 string `json:"genre"`
	PlatformID            string `json:"platform_id"`
	IsInCatalog           int    `json:"is_in_catalog"`
	Friends               []int  `json:"friends"`
	Installed             int    `json:"installed"`
	IsHTML5App            int    `json:"is_html5_app"`
	ScreenOrientation     int    `json:"screen_orientation"`
	MobileControlsType    int    `json:"mobile_controls_type"`
	MobileViewSupportType int    `json:"mobile_view_support_type"`

	Description string  `json:"description"`
	ScreenName  string  `json:"screen_name"`
	Icon16      string  `json:"icon_16"`
	Screenshots []Photo `json:"screenshots"`
	PushEnabled int     `json:"push_enabled"`
}

// Poll описывает опрос.
type Poll struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Created  int64  `json:"created"`
	Question string `json:"question"`
	Votes    int    `json:"votes"`
	Answers  []struct {
		ID    int     `json:"id"`
		Text  string  `json:"text"`
		Votes int     `json:"votes"`
		Rate  float32 `json:"rate"`
	} `json:"answers"`
	Anonymous  bool   `json:"anonymous"`
	Multiple   bool   `json:"multiple"`
	AnswerIDS  []int  `json:"answer_ids"`
	EndDate    int64  `json:"end_date"`
	Closed     bool   `json:"closed"`
	IsBoard    bool   `json:"is_board"`
	CanEdit    bool   `json:"can_edit"`
	CanVote    bool   `json:"can_vote"`
	CanReport  bool   `json:"can_report"`
	CanShare   bool   `json:"can_share"`
	AuthorID   int    `json:"author_id"`
	Photo      *Photo `json:"photo"`
	Background *struct {
		ID     int    `json:"id"`
		Type   string `json:"type"`
		Angle  int    `json:"angle"`
		Color  string `json:"color"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Images []Size `json:"images"`
		Points []struct {
			Position float32 `json:"position"`
			Color    string  `json:"color"`
		} `json:"points"`
	} `json:"background"`
	Friends []int `json:"friends"`
}

// StatsFormat .
type StatsFormat struct {
	PeriodFrom string `json:"period_from"`
	PeriodTo   string `json:"period_to"`
	Visitors   *struct {
		Views    int `json:"views"`
		Visitors int `json:"visitors"`
	} `json:"visitors"`
	// ...
}

// Address описывает адрес в сообществе.
type Address struct {
	ID                int        `json:"id"`
	CountryID         int        `json:"country_id"`
	CityID            int        `json:"city_id"`
	Title             string     `json:"title"`
	Address           string     `json:"address"`
	AdditionalAddress string     `json:"additional_address"`
	Latitude          float32    `json:"latitude"`
	Longitude         float32    `json:"longitude"`
	MetroStationID    int        `json:"metro_station_id"`
	Phone             string     `json:"phone"`
	WorkInfoStatus    string     `json:"work_info_status"`
	Timetable         *Timetable `json:"timetable"`
	TimeOffset        int        `json:"time_offset"`
}

// Timetable расписание.
type Timetable struct {
	Mon *BusinessHours `json:"mon,omitempty"`
	Tue *BusinessHours `json:"tue,omitempty"`
	Wed *BusinessHours `json:"wed,omitempty"`
	Thu *BusinessHours `json:"thu,omitempty"`
	Fri *BusinessHours `json:"fri,omitempty"`
	Sat *BusinessHours `json:"sat,omitempty"`
	Sun *BusinessHours `json:"sun,omitempty"`
}

// BusinessHours рабочее время.
type BusinessHours struct {
	OpenTime       int `json:"open_time,omitempty"`
	CloseTime      int `json:"close_time,omitempty"`
	BreakOpenTime  int `json:"break_open_time,omitempty"`
	BreakCloseTime int `json:"break_close_time,omitempty"`
}

// Attachment .
type Attachment struct {
	Type        string       `json:"type"`
	Photo       *Photo       `json:"photo"`
	Video       *Video       `json:"video"`
	Audio       *Audio       `json:"audio"`
	Doc         *Doc         `json:"doc"`
	Link        *Link        `json:"link"`
	Note        *Note        `json:"note"`
	Poll        *Poll        `json:"poll"`
	Page        *Page        `json:"page"`
	Market      *MarketItem  `json:"market"`
	MarketAlbum *MarketAlbum `json:"market_album"`
	Wall        *Post        `json:"wall"`
	WallReply   *Comment     `json:"wall_reply"`
	Sticker     *Sticker     `json:"sticker"`
	Gift        *Gift        `json:"gift"`
	Album       *Album       `json:"album"`
	PrettyCard  *PrettyCard  `json:"pretty_card"`
	Event       *Event       `json:"event"`
}

// Photo описывает фотографию.
type Photo struct {
	ID      int    `json:"id"`
	AlbumID int    `json:"album_id"`
	OwnerID int    `json:"owner_id"`
	UserID  int    `json:"user_id"`
	Text    string `json:"text"`
	Date    int64  `json:"date"`
	Sizes   []Size `json:"sizes"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`

	Reposts *struct {
		Count int `json:"count"`
	} `json:"reposts"`
	Tags *struct {
		Count int `json:"count"`
	} `json:"tags"`
	Comments *struct {
		Count int `json:"count"`
	} `json:"comments"`
	CanComment int     `json:"can_comment"`
	CanRepost  int     `json:"can_repost"`
	Lat        float32 `json:"lat"`
	Long       float32 `json:"long"`
	PlacerID   int     `json:"placer_id"`
	TagCreated int64   `json:"tag_created"`
	TagID      int     `json:"tag_id"`

	// поля для методов Photos с параметром Extended=true
	Likes *struct {
		Count     int `json:"count"`
		UserLikes int `json:"user_likes"`
	}
	// поля для методов Photos с параметром skip_hidden=true,
	RealOffset int `json:"real_offset"`
	// поля для методов Photos с параметром need_hidden=true,
	Hidden int `json:"hidden"`
}

// Size формат описания размеров фотографии.
type Size struct {
	Type   string `json:"type"`
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Video описывает видеозапись.
type Video struct {
	ID             int    `json:"id"`
	OwnerID        int    `json:"owner_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Duration       int    `json:"duration"`
	Photo130       string `json:"photo_130"`
	Photo320       string `json:"photo_320"`
	Photo640       string `json:"photo_640"`
	Photo800       string `json:"photo_800"`
	Photo1280      string `json:"photo_1280"`
	FirstFrame130  string `json:"first_frame_130"`
	FirstFrame320  string `json:"first_frame_320"`
	FirstFrame640  string `json:"first_frame_640"`
	FirstFrame800  string `json:"first_frame_800"`
	FirstFrame1280 string `json:"first_frame_1280"`
	Date           int64  `json:"date"`
	AddingDate     int64  `json:"adding_date"`
	Files          *struct {
		MP240    string `json:"mp4_240"`
		MP360    string `json:"mp4_360"`
		MP480    string `json:"mp4_480"`
		MP720    string `json:"mp4_720"`
		External string `json:"external"`
	} `json:"files"`
	Views      int    `json:"views"`
	Comments   int    `json:"comments"`
	Player     string `json:"player"`
	Platform   string `json:"platform"`
	CanEdit    int    `json:"can_edit"`
	CanAdd     int    `json:"can_add"`
	IsPrivate  int    `json:"is_private"`
	AccessKey  string `json:"access_key"`
	Processing int    `json:"processing"`
	Live       int    `json:"live"`
	Upcoming   int    `json:"upcoming"`
	IsFavorite bool   `json:"is_favorite"`

	// поля для методов Photos с параметром Extended=true
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
	} `json:"privacy_view*"`
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
	CanComment int `json:"can_comment"`
	CanRepost  int `json:"can_repost"`
	Likes      *struct {
		Count     int `json:"count"`
		UserLikes int `json:"user_likes"`
	}
	Repeat int `json:"repeat"`
}

// Audio описывает аудиозапись.
type Audio struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	URL      string `json:"url"`
	LyricsID int    `json:"lyrics_id"`
	AlbumID  int    `json:"album_id"`
	GenreID  int    `json:"genre_id"`
	Date     int64  `json:"date"`
	NoSearch int    `json:"no_search"`
	IsHq     int    `json:"is_hq"`
}

// Doc описывает документ.
type Doc struct {
	ID      int    `json:"id"`
	OwnerID int    `json:"owner_id"`
	Title   string `json:"title"`
	Size    int    `json:"size"`
	Ext     string `json:"ext"`
	URL     string `json:"url"`
	Date    int64  `json:"date"`
	Type    int    `json:"type"`
	Preview *struct {
		Photo *struct {
			Sizes []Size `json:"sizes"`
		} `json:"photo"`
		Graffiti *struct {
			Src    string `json:"src"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"graffiti"`
		AudioMessage *struct {
			Duration int    `json:"duration"`
			Waveform int    `json:"waveform"`
			LinkOgg  string `json:"link_ogg"`
			LinkMp3  string `json:"link_mp3"`
		} `json:"preview"`
	}
}

// Link описывает прикрепленную ссылку в сообщении, комментарии или записи на стене.
type Link struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Caption     string `json:"caption"`
	Description string `json:"description"`
	Photo       *Photo `json:"photo"`
	Product     *struct {
		Price *struct {
			Amount   int `json:"amount"`
			Currency *struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"currency"`
			Text string `json:"text"`
		} `json:"price"`
	} `json:"product"`
	Button *struct {
		Title  string `json:"title"`
		Action *struct {
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"action"`
	} `json:"button"`
	PreviewPage string `json:"preview_page"`
	PreviewURL  string `json:"preview_url"`
}

// Sticker описывает стикер.
type Sticker struct {
	ProductID int `json:"product_id"`
	StickerID int `json:"sticker_id"`
	Images    []struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
	ImagesWithBackground []struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images_with_background"`
}

// Gift описывает подарок.
type Gift struct {
	ID       int    `json:"id"`
	Thumb256 string `json:"thumb_256"`
	Thumb96  string `json:"thumb_96"`
	Thumb48  string `json:"thumb_48"`
}

// Geo описывает геометку.
type Geo struct {
	Type        string `json:"type"`
	Coordinates *struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	} `json:"coordinates"`
	Place *struct {
		ID        int     `json:"id"`
		Title     string  `json:"title"`
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
		Created   int64   `json:"created"`
		Icon      string  `json:"icon"`
		Country   string  `json:"country"`
		City      string  `json:"city"`

		Type       int    `json:"type"`
		GroupID    int    `json:"group_id"`
		GroupPhoto string `json:"group_photo"`
		Checkins   int    `json:"checkins"`
		Updated    int64  `json:"updated"`
		Address    int    `json:"address"`
	} `json:"place"`
}

// Place описывает место.
type Place struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Created   int64   `json:"created"`
	Icon      string  `json:"icon"`
	Checkins  int     `json:"checkins"`
	Updated   int64   `json:"updated"`
	Type      int     `json:"type"`
	Country   int     `json:"country"`
	City      int     `json:"city"`
	Address   string  `json:"address"`
}

// Story описывает историю.
type Story struct {
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	Date      int64  `json:"date"`
	ExpiresAt int64  `json:"expires_at"`
	IsExpired bool   `json:"is_expired"`
	IsDeleted bool   `json:"is_deleted"`
	CanSee    int    `json:"can_see"`
	Seen      int    `json:"seen"`
	Type      string `json:"type"`
	Photo     *Photo `json:"photo"`
	Video     *Video `json:"video"`
	Link      *struct {
		Text string `json:"text"`
		URL  string `json:"url"`
	} `json:"link"`
	ParentStoryOwnerID int         `json:"parent_story_owner_id"`
	ParentStoryID      int         `json:"parent_story_id"`
	ParentStory        interface{} `json:"parent_story"`
	Replies            *struct {
		Count int `json:"count"`
		New   int `json:"new"`
	} `json:"replies"`
	CanReply   int    `json:"can_reply"`
	CanShare   int    `json:"can_share"`
	CanComment int    `json:"can_comment"`
	Views      int    `json:"views"`
	AccessKey  string `json:"access_key"`
}

// Album описывает альбом с фотографиями.
type Album struct {
	ID          int    `json:"id"`
	Thumb       *Photo `json:"thumb"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     int64  `json:"created"`
	Updated     int64  `json:"updated"`
	Size        int    `json:"size"`
}

// PrettyCard описывает массив элементов-карточек.
type PrettyCard struct {
	Cards []struct {
		CardID  interface{} `json:"card_id"`
		LinkURL string      `json:"link_url"`
		Title   string      `json:"title"`
		Images  []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"images"`
		Button   interface{} `json:"button"`
		Price    string      `json:"price"`
		PriceOld string      `json:"price_old"`
	} `json:"cards"`
}

// Event описывает встречу.
type Event struct {
	ID           int    `json:"id"`
	Time         int64  `json:"time"`
	MemberStatus int    `json:"member_status"`
	IsFavorite   bool   `json:"is_favorite"`
	Address      string `json:"address"`
	Text         string `json:"text"`
	ButtonText   string `json:"button_text"`
	Friends      []int  `json:"friends"`
}
