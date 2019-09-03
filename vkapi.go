package vkapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/huandu/xstrings"
)

// API структура клиента.
type API struct {
	Token  string
	Client *resty.Client
}

// Error ошибка при обращении к методам VK.
type Error struct {
	ErrorCode     int    `json:"error_code"`
	ErrorMsg      string `json:"error_msg"`
	RequestParams []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params"`
}

// Response ответ от методов VK.
type Response struct {
	Error    Error            `json:"error"`
	Response *json.RawMessage `json:"response"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewClientInit новый клиент.
func NewClientInit(token, lang string, timeout int) *API {
	return &API{
		Token: token,
		Client: resty.New().
			SetHostURL("https://api.vk.com/method").
			SetFormData(map[string]string{
				"access_token": token,
				"lang":         lang,
				"v":            "5.101",
			}).
			SetTimeout(time.Duration(timeout) * time.Second),
	}
}

// NewClient новый клиент с русским языком.
func NewClient(token string) *API {
	return NewClientInit(token, "ru", 15)
}

// NewClientWithLang новый клиент с выбранным языком.
func NewClientWithLang(token, lang string) *API {
	return NewClientInit(token, lang, 15)
}

// Request запрос к методам VK.
func (api *API) Request(method string, p, holder interface{}) (interface{}, error) {
	resp, err := api.Client.R().
		SetFormData(structToMap(p)).
		SetResult(&Response{}).
		Post("/" + method)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode()))
	}

	r := resp.Result().(*Response)
	if err := r.Error.ErrorMsg; err != "" {
		return nil, fmt.Errorf("Error code: %d; %s\n%+v", r.Error.ErrorCode, err, r.Error.RequestParams)
	}

	if err := json.Unmarshal(*r.Response, holder); err != nil {
		return nil, err
	}

	switch temp := holder.(type) {
	case *int:
		holder = *temp
	case *[]int:
		holder = *temp
	case *[]User:
		holder = *temp
	case *[]Group:
		holder = *temp
	case *[]Chat:
		holder = *temp
	case *[]Photo:
		holder = *temp
	case *interface{}:
		holder = *temp
	}

	return holder, nil
}

func structToMap(i interface{}) map[string]string {
	m := make(map[string]string)
	iVal := reflect.ValueOf(i)
	t := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		var v string
		switch f.Interface().(type) {
		case string:
			if v = f.String(); v == "" {
				continue
			}
		case int, int64:
			if f.Int() == 0 {
				continue
			}
			v = strconv.FormatInt(f.Int(), 10)
		case uint:
			if f.Uint() == 0 {
				continue
			}
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			if f.Float() == 0 {
				continue
			}
			v = strconv.FormatFloat(f.Float(), 'f', -1, 32)
		case bool:
			if f.Bool() == false {
				continue
			}
			v = "1"
		case []int:
			arr := f.Interface().([]int)
			if len(arr) == 0 {
				continue
			}
			var b bytes.Buffer
			for _, n := range arr {
				b.WriteString(strconv.Itoa(n) + ",")
			}
			v = strings.TrimSuffix(b.String(), ",")
		case *Keyboard, *Timetable:
			if !f.IsNil() {
				b, err := json.Marshal(f.Interface())
				if err != nil {
					log.Println("error:", err)
				}
				v = string(b)
			}
		}
		k := xstrings.ToSnakeCase(t.Field(i).Name)
		m[k] = v
	}
	return m
}

func toBool(i int) bool {
	return i == 1
}
