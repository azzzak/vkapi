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
}

// AlbumUpload загрузка фото в альбом из io.Reader.
func AlbumUpload(uploadURL, filename string, reader io.Reader) (*AlbumUploadResp, error) {
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
	return AlbumUpload(uploadURL, filepath.Base(path), file)
}

// WallUploadResp фото на стене.
type WallUploadResp struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

// WallUpload загрузка фото на стену из io.Reader.
func WallUpload(uploadURL, filename string, reader io.Reader) (*WallUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(WallUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*WallUploadResp), nil
}

// WallUploadFromFile загрузка фото на стену из файла.
func WallUploadFromFile(uploadURL, path string) (*WallUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return WallUpload(uploadURL, filepath.Base(path), file)
}

// AvatarUploadResp главное фото пользователя или сообщества.
type AvatarUploadResp struct {
	Server      int    `json:"server"`
	Photo       string `json:"photo"`
	Mid         int    `json:"mid"`
	Hash        string `json:"hash"`
	MessageCode int    `json:"message_code"`
	ProfileAID  int    `json:"profile_aid"`
}

// AvatarUpload загрузка главного фото пользователя или сообщества из io.Reader.
func AvatarUpload(uploadURL, filename string, reader io.Reader) (*AvatarUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(AvatarUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*AvatarUploadResp), nil
}

// AvatarUploadFromFile загрузка главного фото пользователя или сообщества из файла.
func AvatarUploadFromFile(uploadURL, path string) (*AvatarUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return AvatarUpload(uploadURL, filepath.Base(path), file)
}

// MessageUploadResp фото в сообщении.
type MessageUploadResp struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

// MessageUpload загрузка фото в сообщение из io.Reader.
func MessageUpload(uploadURL, filename string, reader io.Reader) (*MessageUploadResp, error) {
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
	return MessageUpload(uploadURL, filepath.Base(path), file)
}

// ChatCoverUploadResp главное фото для чата.
type ChatCoverUploadResp struct {
	Response int `json:"response"`
}

// ChatUpload загрузка главного фото для чата из io.Reader.
func ChatUpload(uploadURL, filename string, reader io.Reader) (*ChatCoverUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(ChatCoverUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*ChatCoverUploadResp), nil
}

// ChatUploadFromFile загрузка главного фото для чата из файла.
func ChatUploadFromFile(uploadURL, path string) (*ChatCoverUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ChatUpload(uploadURL, filepath.Base(path), file)
}

// ProductUploadResp фото товара.
type ProductUploadResp struct {
	Server   int    `json:"server"`
	Photo    string `json:"photo"`
	Hash     string `json:"hash"`
	CropData string `json:"crop_data"`
	CropHash string `json:"crop_hash"`
}

// ProductUpload загрузка фото товара из io.Reader.
func ProductUpload(uploadURL, filename string, reader io.Reader) (*ProductUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(ProductUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*ProductUploadResp), nil
}

// ProductUploadFromFile загрузка фото товара из файла.
func ProductUploadFromFile(uploadURL, path string) (*ProductUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ProductUpload(uploadURL, filepath.Base(path), file)
}

// ProductSetUploadResp фотоподборка товаров.
type ProductSetUploadResp struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	GID    int    `json:"gid"`
	Hash   string `json:"hash"`
}

// ProductSetUpload загрузка фотоподборки товаров из io.Reader.
func ProductSetUpload(uploadURL, filename string, reader io.Reader) (*ProductSetUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(ProductSetUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*ProductSetUploadResp), nil
}

// ProductSetUploadFromFile загрузка фотоподборки товаров из файла.
func ProductSetUploadFromFile(uploadURL, path string) (*ProductSetUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ProductSetUpload(uploadURL, filepath.Base(path), file)
}

// AudioUploadResp аудиозапись.
type AudioUploadResp struct {
	Redirect string `json:"redirect"`
	Server   int    `json:"server"`
	Audio    string `json:"audio"`
}

// AudioUpload загрузка аудиозаписи из io.Reader.
func AudioUpload(uploadURL, filename string, reader io.Reader) (*AudioUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(AudioUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*AudioUploadResp), nil
}

// AudioUploadFromFile загрузка аудиозаписи из файла.
func AudioUploadFromFile(uploadURL, path string) (*AudioUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return AudioUpload(uploadURL, filepath.Base(path), file)
}

// VideoUploadResp видеозапись.
type VideoUploadResp struct {
	Size    int `json:"size,omitempty"`
	VideoID int `json:"video_id,omitempty"`
}

// VideoUpload загрузка видеозаписи из io.Reader.
func VideoUpload(uploadURL, filename string, reader io.Reader) (*VideoUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(VideoUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*VideoUploadResp), nil
}

// VideoUploadFromFile загрузка видеозаписи из файла.
func VideoUploadFromFile(uploadURL, path string) (*VideoUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return VideoUpload(uploadURL, filepath.Base(path), file)
}

// DocUploadResp документ.
type DocUploadResp struct {
	File string `json:"file,omitempty"`
}

// DocUpload загрузка документа из io.Reader.
func DocUpload(uploadURL, filename string, reader io.Reader) (*DocUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(DocUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*DocUploadResp), nil
}

// DocUploadFromFile загрузка документа из файла.
func DocUploadFromFile(uploadURL, path string) (*DocUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return DocUpload(uploadURL, filepath.Base(path), file)
}

// GroupCoverUploadResp обложка сообщества.
type GroupCoverUploadResp struct {
	Photo string `json:"photo"`
	Hash  string `json:"hash"`
}

// GroupCoverUpload загрузка обложки сообщества из io.Reader.
func GroupCoverUpload(uploadURL, filename string, reader io.Reader) (*GroupCoverUploadResp, error) {
	resp, err := uploadPhoto(uploadURL, filename, reader, new(GroupCoverUploadResp))
	if err != nil {
		return nil, err
	}
	return resp.(*GroupCoverUploadResp), nil
}

// GroupCoverUploadFromFile загрузка обложки сообщества из файла.
func GroupCoverUploadFromFile(uploadURL, path string) (*GroupCoverUploadResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return GroupCoverUpload(uploadURL, filepath.Base(path), file)
}

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
