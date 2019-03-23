### Crr API
Simple API for small SPA
Functions: register/login user, auth by JWT Token, websocket livechat, buying/extending the pass (no paying simulation)

#Endpoints
POST /api/user/register
Body: email, name, password
Returns if succeed: 
```
{
    "message": "Konto zostało utworzone",
    "status": true,
    "user": {
        "id": 5,
        "name": "Tomasz Dzielnik",
        "email": "raxxxasdasdasdndacc@gmail.com",
        "password": "",
        "ticketExpires": "0001-01-01T00:00:00Z",
        "createdAt": "2019-03-23T21:32:19.826466592+01:00",
        "lastVisit": "2019-03-23T20:32:19.828382725Z",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjV9.kosw1yonr7zSNkz1xq0OHFT2X-WpIO9JH6gWs4ApsXo"
    }
}
```

Returns if failed:
```
{
    "message": "Email jest już zajęty",
    "status": false
}
```

POST /api/user/login
Body: email, password
Returns if succeed: 
```
{
    "message": "Zalogowano",
    "status": true,
    "user": {
        "id": 4,
        "name": "Tomasz Dzielnik",
        "email": "raasdasdasdndacc@gmail.com",
        "password": "",
        "ticketExpires": "0001-01-01T00:00:00Z",
        "createdAt": "2019-03-23T19:12:13+01:00",
        "lastVisit": "2019-03-23T20:32:46.597097726Z",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjR9.sz7T_Jlg7jCC6ogiBHmZMUAVXn6rTkEaA9F3TVEh5u8"
    }
}
```

Returns if failed:
```
{
    "message": "Błędne hasło",
    "status": false
}
```

POST /api/user/refresh
Headers:
    Authorization: Basic <token>
Returns if succeed: 
```
{
    "message": "Odświeżono",
    "status": true,
    "user": {
        "id": 4,
        "name": "Tomasz Dzielnik",
        "email": "raasdasdasdndacc@gmail.com",
        "password": "",
        "ticketExpires": "0001-01-01T00:00:00Z",
        "createdAt": "2019-03-23T21:33:01+01:00",
        "lastVisit": "2019-03-23T21:33:01+01:00",
        "token": ""
    }
}
```

Returns if failed:
```
{
    "message": "Brak tokena",
    "status": false
}
```

POST /api/ticket/buy
Extend ticket expires date +30 from current value
By default value is null, so first call will set time.Now()+30days
Headers:
    Authorization: Basic <token>
Returns if succeed: 
```
{
    "message": "Zakupiono bilet",
    "status": true,
    "ticketExpires": "2019-04-22T19:41:45.070670823Z"
}
```

GET /api/chat/history
Returns each message on chat
Headers:
    Authorization: Basic <token>
Returns if succeed: 
```
{
    "message": "Pobrano",
    "messages": [
        {
            "id": 1,
            "userId": 2,
            "name": "Tomasz Dzielnik",
            "message": "Siema!",
            "createdAt": "2019-03-23T11:46:30+01:00"
        },
        {
            "id": 2,
            "userId": 3,
            "name": "Tomasz Dzielnik",
            "message": "Hejo!",
            "createdAt": "2019-03-23T11:46:30+01:00"
        },
        {
            "id": 3,
            "userId": 2,
            "name": "Tomasz Dzielnik",
            "message": "Co tam?",
            "createdAt": "2019-03-23T19:29:48+01:00"
        },
        {
            "id": 4,
            "userId": 2,
            "name": "Tomasz Dzielnik",
            "message": "Halo",
            "createdAt": "2019-03-23T19:29:52+01:00"
        },
        {
            "id": 5,
            "userId": 2,
            "name": "Tomasz Dzielnik",
            "message": "Odbjuuur?",
            "createdAt": "2019-03-23T19:29:56+01:00"
        },
        {
            "id": 6,
            "userId": 2,
            "name": "Tomasz Dzielnik",
            "message": "Ejjj",
            "createdAt": "2019-03-23T19:29:59+01:00"
        }
    ],
    "status": true
}
```

WEBSOCKET /api/chat/live
Use like:
ws://localhost:8000/api/chat?token=<PASTE_HERE_JWT_TOKEN>
It allows to connect with websocket and use livechat.