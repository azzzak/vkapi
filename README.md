# vkapi [![GoDoc](https://godoc.org/github.com/azzzak/vkapi?status.svg)](https://godoc.org/github.com/azzzak/vkapi) [![Go Report Card](https://goreportcard.com/badge/github.com/azzzak/vkapi)](https://goreportcard.com/report/github.com/azzzak/vkapi)

Golang-пакет для работы с API VK.
Версия протокола **101**.

## Получение токена

Варианты получения токена описаны на [странице](https://vk.com/dev/access_token).

Для получение пользовательского токена через _Implicit flow_ можно воспользоваться утилитой в папке `cmd`. Она принимает _id_ (идентификатор приложения) и _scope_ (настройки доступа). Например:

`go run main.go -id=1234567 -scope=photos,wall,groups`

После запуска надо открыть страницу [http://localhost:3000/token](http://localhost:3000/token), которая направит на сайт VK для подтверждения доступа. В случае согласия токен будет доступен в адресной строке.

Самый простой способ получить токен сообщества — создать его в настройках: `Управление > Работа с API > Создать ключ`.

## Правила

- Если параметр принимает значение **0** или **1**, то используется тип `bool`
- Для числовых идентификаторов используется тип `int`
- Для списка идентификаторов используется тип `[]int`
- Если параметр ожидает JSON-строку, то ему передается соответствующая структура, которая будет сериализована автоматически
- Если метод возвращает **1** в случае успеха, то используется тип `bool`

Иногда параметры могут изменять формат выдачи, в этих случаях возвращается тип `interface{}`, который перед использованием потребуется привести к нужному типу. Возможные варианты описаны в комментариях к методам.

## Пример

```Go
// создаем клиент для работы с API, передаем пользовательский токен
vk := vkapi.NewClient("token")

// создаем новый альбом
createParams := vkapi.PhotosCreateAlbumParams{
  Title: "New album",
}
album, _ := vk.PhotosCreateAlbum(createParams)

// получаем адрес для загрузки в новый альбом
getUploadURLParams := vkapi.PhotosGetUploadServerParams{
  AlbumID: album.ID,
}
upload, _ := vk.PhotosGetUploadServer(getUploadURLParams)

// загружаем фото по полученному адресу
u, _ := vkapi.AlbumUploadFromFile(upload.UploadURL, "photo.jpg")

// сохраняем загруженное фото
saveParams := vkapi.PhotosSaveParams{
  Server:     u.Server,
  AlbumID:    u.AID,
  PhotosList: u.PhotosList,
  Hash:       u.Hash,
}
photo, _ := vk.PhotosSave(saveParams)

// постим фото на стену
postParams := vkapi.WallPostParams{
  Attachments: vkapi.MakeAttachment("photo", photo[0].OwnerID, photo[0].ID),
}
vk.WallPost(postParams)
```

## Реализованные методы

| Группа методов |     | Готово / Всего |
| -------------- | :-: | :------------: |
| account        |     |   **0** / 19   |
| appWidgets     |     |   **0** / 8    |
| apps           |     |   **0** / 8    |
| auth           |     |   **0** / 2    |
| board          |     |   **0** / 13   |
| database       |  ✓  |  **12** / 12   |
| docs           |     |   **0** / 11   |
| fave           |     |   **0** / 23   |
| friends        |  ✓  |  **18** / 18   |
| gifts          |     |   **0** / 1    |
| groups         |  ✓  |  **45** / 45   |
| leadForms      |     |   **0** / 7    |
| likes          |     |   **0** / 4    |
| market         |     |   **0** / 24   |
| messages       |  ✓  |  **38** / 38   |
| newsfeed       |     |   **0** / 15   |
| notes          |     |   **0** / 10   |
| notifications  |     |   **0** / 3    |
| pages          |     |   **0** / 8    |
| photos         |  ✓  |  **46** / 46   |
| polls          |     |   **0** / 9    |
| prettyCards    |     |   **0** / 6    |
| search         |     |   **0** / 1    |
| stats          |     |   **0** / 3    |
| status         |  ✓  |   **2** / 2    |
| storage        |     |   **0** / 3    |
| stories        |     |   **0** / 14   |
| streaming      |     |   **0** / 5    |
| users          |     |   **0** / 6    |
| utils          |     |   **0** / 7    |
| video          |     |   **0** / 24   |
| wall           |  ✓  |  **23** / 23   |
| widgets        |     |   **0** / 2    |

## Использование остальных методов

Вызвать метод, который пока не реализован можно с помощью функции `Request`. Она принимает название метода, структуру с параметрами и ссылку на структуру, которая будет заполнена ответом метода.
