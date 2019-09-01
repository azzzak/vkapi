package vkapi

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-resty/resty/v2"
)

// AlbumUploadResp фото в альбоме.
type AlbumUploadResp struct {
	Server     int    `json:"server"`
	PhotosList string `json:"photos_list"`
	Hash       string `json:"hash"`
	AID        int    `json:"aid"`
} //file1-5

// AlbumUploadFromReader загрузка фото в альбом из io.Reader.
func AlbumUploadFromReader(uploadURL, filename string, reader io.Reader) (*AlbumUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(AlbumUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*AlbumUploadResp), nil
}

// AlbumUploadFromFile загрузка фото в альбом из файла.
func AlbumUploadFromFile(uploadURL, path string) (*AlbumUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return AlbumUploadFromReader(uploadURL, filepath.Base(path), file)
}

// WallUploadResp фото на стене.
type WallUploadResp struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
} //photo

// AvatarUploadResp главное фото пользователя или сообщества.
type AvatarUploadResp struct {
	Server      int    `json:"server"`
	Photo       string `json:"photo"`
	Mid         int    `json:"mid"`
	Hash        string `json:"hash"`
	MessageCode int    `json:"message_code"`
	ProfileAID  int    `json:"profile_aid"`
} //photo

// MessageUploadResp фото в сообщении.
type MessageUploadResp struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
} //photo

// MessageUploadFromReader загрузка фото в сообщение из io.Reader.
func MessageUploadFromReader(uploadURL, filename string, reader io.Reader) (*MessageUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(MessageUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*MessageUploadResp), nil
}

// MessageUploadFromFile загрузка фото в сообщение из файла.
func MessageUploadFromFile(uploadURL, path string) (*MessageUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return MessageUploadFromReader(uploadURL, filepath.Base(path), file)
}

// ChatCoverUploadResp главное фото для чата.
type ChatCoverUploadResp struct {
	Response int `json:"response"`
} //file

// ProductUploadResp фото товара.
type ProductUploadResp struct {
	Server   int    `json:"server"`
	Photo    string `json:"photo"`
	Hash     string `json:"hash"`
	CropData string `json:"crop_data"`
	CropHash string `json:"crop_hash"`
} //file

// ProductSetUploadResp фото подборки товаров.
type ProductSetUploadResp struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	GID    int    `json:"gid"`
	Hash   string `json:"hash"`
} //file

// AudioUploadResp аудиозапись.
type AudioUploadResp struct {
	Redirect string `json:"redirect"`
	Server   int    `json:"server"`
	Audio    string `json:"audio"`
} //file

// VideoUploadResp видеозапись.
type VideoUploadResp struct {
	Size    int `json:"size,omitempty"`
	VideoID int `json:"video_id,omitempty"`
} // video_file

// DocUploadResp документ.
type DocUploadResp struct {
	File string `json:"file,omitempty"`
} // file

// GroupCoverUploadResp обложка сообщества.
type GroupCoverUploadResp struct {
	Photo string `json:"photo"`
	Hash  string `json:"hash"`
} // photo

func uploadPhoto(uploadURL, filename string, reader io.Reader, holder interface{}) (interface{}, error) {
	var fieldName string
	switch holder.(type) {
	case *MessageUploadResp, *WallUploadResp, *AvatarUploadResp, *GroupCoverUploadResp:
		fieldName = "photo"
	case *ChatCoverUploadResp, *AudioUploadResp, *DocUploadResp, *ProductUploadResp, *ProductSetUploadResp:
		fieldName = "file"
	case *VideoUploadResp:
		fieldName = "video_file"
	case *AlbumUploadResp:
		fieldName = "file1"
	}

	client := resty.New()
	resp, err := client.R().
		SetFileReader(fieldName, filename, reader).
		Post(uploadURL)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp.Body(), holder)
	if err != nil {
		return nil, err
	}
	return holder, nil
}

// Keyboard структура клавиатуры бота.
type Keyboard struct {
	OneTime bool       `json:"one_time"`
	Buttons [][]Button `json:"buttons"`
}

// Button структура кнопки клавиатуры бота.
type Button struct {
	Action Action `json:"action"`
	Color  string `json:"color"`
}

const (
	// ColorPrimary синяя кнопка, обозначает основное действие
	ColorPrimary = "primary"
	// ColorSecondary обычная белая кнопка
	ColorSecondary = "secondary"
	// ColorNegative опасное действие, или отрицательное действие (отклонить, удалить и т.д.)
	ColorNegative = "negative"
	// ColorPositive согласиться, подтвердить
	ColorPositive = "positive"
)

// Action структура действия кнопки клавиатуры бота.
type Action struct {
	Type    string `json:"type"`
	Label   string `json:"label,omitempty"`
	Payload string `json:"payload,omitempty"`
	Hash    string `json:"hash,omitempty"`
	AppID   string `json:"app_id,omitempty"`
	OwnerID string `json:"owner_id,omitempty"`
}

const (
	// ActionText текстовая кнопка
	ActionText = "text"
	// ActionLocation отправляет местоположение в диалог с ботом/беседу; занимает всю ширину клавиатуры
	ActionLocation = "location"
	// ActionVKPay открывает окно оплаты VK Pay с предопределёнными параметрами. Кнопка называется “Оплатить через VK pay”, VK pay отображается как логотип; занимает всю ширину клавиатуры
	ActionVKPay = "vkpay"
	// ActionVKApps лткрывает указанное приложение VK Apps; занимает всю ширину клавиатуры
	ActionVKApps = "open_app"
)

// NewKeyboard новая клавиатура.
func NewKeyboard(hide bool) *Keyboard {
	return &Keyboard{
		OneTime: hide,
		Buttons: [][]Button{},
	}
}

// NewTextButton новая кнопка.
func NewTextButton(label, color string, payload ...string) Button {
	var p string
	if len(payload) > 0 {
		p = payload[0]
	}

	return Button{
		Action: Action{
			Type:    ActionText,
			Label:   label,
			Payload: p,
		},
		Color: color,
	}
}

// ButtonsRow ряд кнопок на клавиатуре.
func (k *Keyboard) ButtonsRow(b ...Button) *Keyboard {
	k.Buttons = append(k.Buttons, b)
	return k
}

// MakeAttachment аттачмент.
func MakeAttachment(t string, id, oid int) string {
	return fmt.Sprintf("%s%d_%d", t, id, oid)
}
