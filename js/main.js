new Vue({
    el: "#app",
    data: function(){
        return {
            messages: [],
            message: '',
            token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjR9.sz7T_Jlg7jCC6ogiBHmZMUAVXn6rTkEaA9F3TVEh5u8',
            ws: null // There will be socket's instance
        }
    },
    methods: {
        submit: function(){
            if(!this.message || this.message.trim() == ''){
                alert("Wiadomość nie może być pusta")
            }
            this.ws.send(
                JSON.stringify({
                    message: this.message
                })
            )
            this.message = '';
        }
    },
    created: function() {
        // Get history
        fetch('http://localhost:8000/api/chat/history', {
            headers: {
                "Authorization": "Basic " + this.token
            }
        })
        .then(response => response.json())
        .then(response => {
            var self = this
            response.messages.forEach(function(el) {
                // Timestamp to hours:minutes
                var tmp = new Date(el.createdAt)
                el.createdAt = tmp.getHours() + ":" + (tmp.getMinutes() < 10 ? '0'+tmp.getMinutes() : tmp.getMinutes())
                
                self.messages.unshift(el)
            })
        })

        // Connect with Chat
        this.ws = new WebSocket('ws://localhost:8000/api/chat/live?token=' + this.token)
        this.ws.addEventListener('message', ({data}) => {
            const message = JSON.parse(data);
            // Timestamp to hours:minutes
            var tmp = new Date(message.createdAt)
            message.createdAt = tmp.getHours() + ":" + (tmp.getMinutes() < 10 ? '0'+tmp.getMinutes() : tmp.getMinutes())
            this.messages.unshift(message)
        })
    }
})