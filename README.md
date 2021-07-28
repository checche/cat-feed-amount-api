# cat-feed-amount-api
猫の給餌量計算をするAPI

```bash

# running on localhost:8080
go run main.go

```

## example
```json

// POST to http://localhost:8080/
{
    "weight": 2.4,  // 体重(kg)
    "age": 4,  // 月齢
    "is_neutered": false, // 有無
    "kcal": 412,  // 餌のカロリー
    "per_gram": 100  // kcalが餌何グラムに対してのものか
}

```
