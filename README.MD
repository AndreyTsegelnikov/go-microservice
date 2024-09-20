## Запуск сервиса

go run cmd/main.go 

## Запрос получения текущего времени

```curl
curl  -X GET \
  'http://localhost:8080/time' \
  --header 'Accept: */*' \
```