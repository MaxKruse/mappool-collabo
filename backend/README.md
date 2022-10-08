# Backend

## Features

- [x] osu bancho login
- [x] create tournament
- [x] create mappool
- [x] permissions (groups)
- [x] suggestions to mappool
- [x] feedback for maps
- [x] voting for maps
- [x] add map to mappool
- [x] replay upload

### Endpoints

Unless otherwised specified, all endpoints require authentication.

Authentiation is to be provided via `Authorization: Bearer <token>` where the token can be aquired by navigating to `/oauth/login`. It will redirect to a frontend, passing the token as a query parameter `/login?token=<token>`.

#### Users

```json
{
  "id": 1199528,
  "avatar_url": "https://a.ppy.sh/1199528?1654635999.jpeg",
  "username": "[BH]Lithium"
}
```

1. GET/user/ - Lists all users
2. GET /user/{id} - Get a specific user
3. GET /user/self - Get the current user

#### Tournaments

```json
{
  "id": 3,
  "name": "Another One",
  "description": "No cock",
  "owner": {
    "id": 1199528,
    "avatar_url": "https://a.ppy.sh/1199528?1654635999.jpeg",
    "username": "[BH]Lithium"
  },
  "testplayers": [
    {
      "id": 1199528,
      "avatar_url": "https://a.ppy.sh/1199528?1654635999.jpeg",
      "username": "[BH]Lithium"
    }
  ],
  "mappoolers": [
    {
      "id": 1199528,
      "avatar_url": "https://a.ppy.sh/1199528?1654635999.jpeg",
      "username": "[BH]Lithium"
    }
  ],
}
``` 

1. GET /tournament/ - Lists all tournaments
2. GET /tournament/{id} - Get a specific tournament (includes mappool)
3. POST /tournament/ - Create a tournament
4. PUT /tournament/{id} - Update a tournament
5. DELETE /tournament/{id} - Delete a tournament