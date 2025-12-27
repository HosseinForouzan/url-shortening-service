# URL Shortening Service

This shortening url API help users to shorten their long url and use short url instead of it.

## Features
- 🔗 URL shortening and redirection
- ⚡ High-performance redirects with Redis caching
- 📊 Visit tracking and basic analytics
- 🗄️ PostgreSQL persistence
- 🌐 RESTful API
- 🐳 Docker support

## Tech Stack
- Golang
- PostgreSQL
- Docker
- Redis

## Instalation
```bash
git clone https://github.com/HosseinForouzan/url-shortening-service
cd url-shortening-service
go mod tidy
```

### Redirect to Original URL

**GET** `/{short_code}`

- Example: Short URL `http://localhost:8080/abc123`
- Behavior: Redirects to the original URL
- HTTP Status: 302 (temporary)


## API Reference
### Base URL
http://localhost:8080/

### Create Short URL
#### Request Body

**POST** `/shorten`
```json
{
  "url": "https://www.example.com/some/long/url"
}
```
#### Response

```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### Retrieve Original URL

**GET** `/shorten/abc123`

#### Response
```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```
### Retrieve Original URL

**PUT** `/shorten/abc123`
#### Response

```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### Retrieve Original URL

**DELETE** `/shorten/abc123`


### Retrieve URL Statistics

**GET** `/shorten/abc123/stats`

#### Response

```json
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z",
  "visits" : 10
}
```
