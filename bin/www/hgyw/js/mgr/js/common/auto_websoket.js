// JavaScript Document
window.onload = function () {
    NewWebSocket();
}
var ws

function NewWebSocket() {
    // alert("closed");
    if ("WebSocket" in window) {
        var url = window.location.hostname + ':' + window.location.port
        url = url.replace("//", "")
        // Let us open a web socket
        ws = new WebSocket("ws://" + url + "/ws");
        // ws = new WebSocket("ws://127.0.0.1:8888/ws");
        ws.onopen = function () {
            sendSessionJsonStr();
        };
        //接收到服务端的数据
        ws.onmessage = function (evt) {

            var received_msg = evt.data;
            var pJson = JSON.parse(received_msg);
            // alert(pJson.cmd)
            selectCmdPram(pJson); //查询是哪个cmd

        }
        ws.onclose = function () {
            // alert("closed");
        };

    } else {
        // The browser doesn't support WebSocket
        alert("这不是websocket");
    }
}

function sendSessionJsonStr() {
    var cookie = getCookie("session");
    var str = '{"cmd":5,"param":{"session":"' + cookie + '"}}';
    ws.send(str);
}

function getCookie(name) {
    var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");

    if (arr = document.cookie.match(reg))

        return unescape(arr[2]);
    else
        return null;
}