new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        chatContent: '', // A running list of chat messages displayed on the screen
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var element = document.getElementById('chat-messages');
            var msg = JSON.parse(e.data);
            self.chatContent += '<div class="chip">' // Avatar
                + msg.title
                + '</div>'
                + emojione.toImage(msg.message) + '<br/>'; // Parse emojis
            // element.innerHTML = msg.message;
        });
    },
});