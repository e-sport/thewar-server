var websocket = require('ws');

var ws = new websocket('ws://127.0.0.1:8000');

ws.onopen = function (){
    ws.send("test");
};

ws.onmessage = function(e){
    console.log(e);
};
