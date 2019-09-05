package vkapi

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-resty/resty/v2"
)

// UploadAlbumResp фото в альбоме.
type UploadAlbumResp struct {
	Server     int    `json:"server"`
	PhotosList string `json:"photos_list"`
	Hash       string `json:"hash"`
	AID        int    `json:"aid"`
}

// UploadAlbum загрузка фото в альбом из io.Reader.
func UploadAlbum(uploadURL, filename string, reader io.Reader) (*UploadAlbumResp, error) {
	resp, err := upload(uploadURL, "file1", filename, reader, new(UploadAlbumResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadAlbumResp), nil
}

// UploadAlbumFromFile загрузка фото в альбом из файла.
func UploadAlbumFromFile(uploadURL, path string) (*UploadAlbumResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadAlbum(uploadURL, filepath.Base(path), file)
}

// UploadWallResp фото на стене.
type UploadWallResp struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

// UploadWall загрузка фото на стену из io.Reader.
func UploadWall(uploadURL, filename string, reader io.Reader) (*UploadWallResp, error) {
	resp, err := upload(uploadURL, "photo", filename, reader, new(UploadWallResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadWallResp), nil
}

// UploadWallFromFile загрузка фото на стену из файла.
func UploadWallFromFile(uploadURL, path string) (*UploadWallResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadWall(uploadURL, filepath.Base(path), file)
}

// UploadAvatarResp главное фото пользователя или сообщества.
type UploadAvatarResp struct {
	Server      int    `json:"server"`
	Photo       string `json:"photo"`
	Mid         int    `json:"mid"`
	Hash        string `json:"hash"`
	MessageCode int    `json:"message_code"`
	ProfileAID  int    `json:"profile_aid"`
}

// UploadAvatar загрузка главного фото пользователя или сообщества из io.Reader.
func UploadAvatar(uploadURL, filename string, reader io.Reader) (*UploadAvatarResp, error) {
	resp, err := upload(uploadURL, "photo", filename, reader, new(UploadAvatarResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadAvatarResp), nil
}

// UploadAvatarFromFile загрузка главного фото пользователя или сообщества из файла.
func UploadAvatarFromFile(uploadURL, path string) (*UploadAvatarResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadAvatar(uploadURL, filepath.Base(path), file)
}

// UploadMessageResp фото в сообщении.
type UploadMessageResp struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

// UploadMessage загрузка фото в сообщение из io.Reader.
func UploadMessage(uploadURL, filename string, reader io.Reader) (*UploadMessageResp, error) {
	resp, err := upload(uploadURL, "photo", filename, reader, new(UploadMessageResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadMessageResp), nil
}

// UploadMessageFromFile загрузка фото в сообщение из файла.
func UploadMessageFromFile(uploadURL, path string) (*UploadMessageResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadMessage(uploadURL, filepath.Base(path), file)
}

// UploadChatCoverResp главное фото для чата.
type UploadChatCoverResp struct {
	Response int `json:"response"`
}

// UploadChatCover загрузка главного фото для чата из io.Reader.
func UploadChatCover(uploadURL, filename string, reader io.Reader) (*UploadChatCoverResp, error) {
	resp, err := upload(uploadURL, "file", filename, reader, new(UploadChatCoverResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadChatCoverResp), nil
}

// UploadChatCoverFromFile загрузка главного фото для чата из файла.
func UploadChatCoverFromFile(uploadURL, path string) (*UploadChatCoverResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadChatCover(uploadURL, filepath.Base(path), file)
}

// UploadProductResp фото товара.
type UploadProductResp struct {
	Server   int    `json:"server"`
	Photo    string `json:"photo"`
	Hash     string `json:"hash"`
	CropData string `json:"crop_data"`
	CropHash string `json:"crop_hash"`
}

// UploadProduct загрузка фото товара из io.Reader.
func UploadProduct(uploadURL, filename string, reader io.Reader) (*UploadProductResp, error) {
	resp, err := upload(uploadURL, "file", filename, reader, new(UploadProductResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadProductResp), nil
}

// UploadProductFromFile загрузка фото товара из файла.
func UploadProductFromFile(uploadURL, path string) (*UploadProductResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadProduct(uploadURL, filepath.Base(path), file)
}

// UploadProductSetResp фотоподборка товаров.
type UploadProductSetResp struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	GID    int    `json:"gid"`
	Hash   string `json:"hash"`
}

// UploadProductSet загрузка фотоподборки товаров из io.Reader.
func UploadProductSet(uploadURL, filename string, reader io.Reader) (*UploadProductSetResp, error) {
	resp, err := upload(uploadURL, "file", filename, reader, new(UploadProductSetResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadProductSetResp), nil
}

// UploadProductSetFromFile загрузка фотоподборки товаров из файла.
func UploadProductSetFromFile(uploadURL, path string) (*UploadProductSetResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadProductSet(uploadURL, filepath.Base(path), file)
}

// UploadAudioResp аудиозапись.
type UploadAudioResp struct {
	Redirect string `json:"redirect"`
	Server   int    `json:"server"`
	Audio    string `json:"audio"`
}

// UploadAudio загрузка аудиозаписи из io.Reader.
func UploadAudio(uploadURL, filename string, reader io.Reader) (*UploadAudioResp, error) {
	resp, err := upload(uploadURL, "file", filename, reader, new(UploadAudioResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadAudioResp), nil
}

// UploadAudioFromFile загрузка аудиозаписи из файла.
func UploadAudioFromFile(uploadURL, path string) (*UploadAudioResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadAudio(uploadURL, filepath.Base(path), file)
}

// UploadVideoResp видеозапись.
type UploadVideoResp struct {
	Size    int `json:"size,omitempty"`
	VideoID int `json:"video_id,omitempty"`
}

// UploadVideo загрузка видеозаписи из io.Reader.
func UploadVideo(uploadURL, filename string, reader io.Reader) (*UploadVideoResp, error) {
	resp, err := upload(uploadURL, "video_file", filename, reader, new(UploadVideoResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadVideoResp), nil
}

// UploadVideoFromFile загрузка видеозаписи из файла.
func UploadVideoFromFile(uploadURL, path string) (*UploadVideoResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadVideo(uploadURL, filepath.Base(path), file)
}

// UploadDocResp документ.
type UploadDocResp struct {
	File string `json:"file,omitempty"`
}

// UploadDoc загрузка документа из io.Reader.
func UploadDoc(uploadURL, filename string, reader io.Reader) (*UploadDocResp, error) {
	resp, err := upload(uploadURL, "file", filename, reader, new(UploadDocResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadDocResp), nil
}

// UploadDocFromFile загрузка документа из файла.
func UploadDocFromFile(uploadURL, path string) (*UploadDocResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadDoc(uploadURL, filepath.Base(path), file)
}

// UploadGroupCoverResp обложка сообщества.
type UploadGroupCoverResp struct {
	Photo string `json:"photo"`
	Hash  string `json:"hash"`
}

// UploadGroupCover загрузка обложки сообщества из io.Reader.
func UploadGroupCover(uploadURL, filename string, reader io.Reader) (*UploadGroupCoverResp, error) {
	resp, err := upload(uploadURL, "photo", filename, reader, new(UploadGroupCoverResp))
	if err != nil {
		return nil, err
	}
	return resp.(*UploadGroupCoverResp), nil
}

// UploadGroupCoverFromFile загрузка обложки сообщества из файла.
func UploadGroupCoverFromFile(uploadURL, path string) (*UploadGroupCoverResp, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return UploadGroupCover(uploadURL, filepath.Base(path), file)
}

func upload(uploadURL, field, filename string, reader io.Reader, holder interface{}) (interface{}, error) {
	client := resty.New()
	resp, err := client.R().
		SetFileReader(field, filename, reader).
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
