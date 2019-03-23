# Websocket client
## Created for testing Websocket live chat

View: index.html
Working with Websocket and getting chat history: js/main.js

# Connecting with WebSocket
```
this.ws = new WebSocket('ws://localhost:8000/api/chat/live?token=' + this.token)
```

# Listening for new messages (requires previous step)
```
this.ws.addEventListener('message', ({data}) => {
            const message = JSON.parse(data);
            // Timestamp to hours:minutes
            var tmp = new Date(message.createdAt)
            message.createdAt = tmp.getHours() + ":" + (tmp.getMinutes() < 10 ? '0'+tmp.getMinutes() : tmp.getMinutes())
            this.messages.unshift(message)
        })
```


# Sending message
```
this.ws.send(
    JSON.stringify({
        message: this.message
    })
)
```
You do not need to send anything more than message because server will recognize who you are by Token.


# Getting every message that has been ever sent
GET /api/chat/history
Needs Authorization header with value Basic <token>
