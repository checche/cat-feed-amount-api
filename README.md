# cat-feed-amount-api
猫の給餌量計算をするAPI

```bash

# running on localhost:8080
go run main.go

```

## example
```json

// POST to http://localhost:8080/
// request body (json)
{
    "weight": 2.4,  // 体重(kg)
    "age": 4,  // 月齢
    "is_neutered": false, // 去勢/避妊手術の有無
    "kcal": 412,  // 餌のカロリー(kcal)
    "per_gram": 100  // kcalが餌何グラムに対してのものか
}

// response (json)
{
    "feed_amount": 81.90289707688197,  // 1日の給餌量(g)
    "feed_kcal": 337.43993595675374  // 1日のカロリー(kcal)
}

```
