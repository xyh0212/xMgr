//"/upload?user=lkdfgl&name=kjdf&filekey=tfhhcgh" 
function HttpFileConnetion (fileUrl){
$(function(){
  alert(fileUrl);
 var options={
            url:fileUrl,
            type:"post",
            success:function(mes){
               selectUploadCmdPram(mes);
            }
        };
        $("#form2").submit(function() {
            $(this).ajaxSubmit(options);
                return false;
        });

});
}
function HttpFileConn (fileUrl,func){
    $(function(){
        var options={
            url:fileUrl+"&session="+getCookie("session"),
            type:"post",
            success:function(mes){
                func(JSON.parse(mes));
            }
        };
        $("#form2").submit(function() {
            $(this).ajaxSubmit(options);
            return false;
        });

    });
}
//发送 post
function HttpConnetion(filepath){
var  session =	getCookie("session")
var path = "/post?path="+filepath+"&"+getCheckedStr()+"&session="+session
 $(document).ready(function(){
    $.post(path,json, function(data,status){
	selectCmdPram(data)
	document.write(data);
     // alert("数据：" + data + "\n状态：" + status);
    });
});
}

//发送 post
function HttpGetConn(filepath){
    var path = filepath+"&session="+getCookie("session")
    $(document).ready(function(){
        $.post(path,{}, function(data,status){
            //selectCmdPram(data)
            document.write(data);
            // alert("数据：" + data + "\n状态：" + status);
        });
    });
}



//发送 post
function HttpConn(url,json,func){
    json["session"] = getCookie("session");
    $(document).ready(function(){
        $.post(url,json, function(data,status){
            func(JSON.parse(data))
        });
    });
}

function rec(data){
    if (data.cmd == 1) {
        alert(data.info)
    }
    if (data.cmd == 10) {
        alert("错误："+data.code+"\n信息："+data.info)
        return
    }

}

function WebSocketConn(func) {
    if ("WebSocket" in window) {
        var url = window.location.hostname + ':' + window.location.port
        url = url.replace("//", "")
        // Let us open a web socket
        ws = new WebSocket("ws://" + url + "/ws");
        ws.onopen = function () {
            sendSessionJsonStr();
        };
        //接收到服务端的数据
        ws.onmessage = function (evt) {
            var received_msg = evt.data;
            var pJson = JSON.parse(received_msg);
            func(ws,pJson); //查询是哪个cmd

        }
        ws.onclose = function () {
            sendSessionJsonStr();
        };
    } else {
        // The browser doesn't support WebSocket
        alert("这不是websocket");
    }

    function sendSessionJsonStr() {
        var cookie = getCookie("session");
        var str = '{"cmd":5,"param":{"session":"' + cookie + '"}}';
        ws.send(str);
    }
}





























