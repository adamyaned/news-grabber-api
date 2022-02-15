
# News grabber API

News grabber API description

## API Reference

#### Get news sources

```http
  GET /sources/${category}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `category`      | `string` | **Required**. Category of news (business, entertainment, general, health, sports, technology)|


#### Collect news

```http
  GET /search/${category}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `category`      | `string` | **Required**. Category of news (business, entertainment, general, health, sports, technology)|

#### Get news

```http
  GET /result/${category}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `category`      | `string` | **Required**. Category of news (business, entertainment, general, health, sports, technology)|


## Used API's

- [News API](https://newsapi.org)


## Authors

- [@adamyaned](https://github.com/adamyaned)

