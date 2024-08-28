# List Artikel
```
    GET /api/article
```

Response
```json
    {
    "id": 0,
    "title": "string",
    "img_url": "string",
    "posted": "string",
    "like_count": 0,
    "comment_count": 0
    }

```

# View Artikel
```
    GET /api/article/{id}
```

```json
    {
    "id": 0,
    "title": "string",
    "img_url": "string",
    "posted": "string",
    "author": "string",
    "content": "string",
    "comment": [
      {"id": 0,
      "name": "string",
      "comment":  "string"}
    ]
    }

```

# List Komentar Artikel
```
    GET /api/article/{id}/comment
```

```json
{
  "comment": [
    {"id": 0,
      "name": "string",
      "comment":  "string"}
  ]
}
```

# React Artikel
```
    POST /api/article/{id}/react
```

Request Body
```json
{
  "react_id": 1
}
```

# List Master React
```
    POST /api/master/react
```

```json
{
  "react_id": 1,
  "name": "string",
  "img_url": "string"
}
```

# Share Artikel
```
     GET /api/article/{id}/share
```