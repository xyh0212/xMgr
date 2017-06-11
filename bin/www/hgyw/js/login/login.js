// JavaScript Document
//////////////////////////////////////////////////////////////
//websoket连接
var ws
function NewWebSocket() {

    if ("WebSocket" in window) {
        // Let us open a web socket
        var url =window.location.hostname+':'+window.location.port
        url=url.replace("//","")
        ws = new WebSocket("ws://"+url+"/ws");
        ws.onopen = function() {

        sendLoginJsonStr();
            //alert(ws);
        };
        //接收到服务端的数据
        ws.onmessage = function(evt) {
        	 var received_msg = evt.data;
             var pJson = JSON.parse(received_msg);
        	  selectCmdPram(pJson);//查询是哪个cmd
    	};
        ws.onclose = function() {
            //alert("closed");
        };
        ws.onerror =function(evt){
            alert(evt)
        }
   
    }else {
        // The browser doesn't support WebSocket
        alert("这不是websocket");
    }
}
//发送登入的json
function sendLoginJsonStr() {

	var gPassword = document.getElementById("password").value;
    var gUser= document.getElementById("user").value;
   if(gPassword=="" && gUser==""){
   	    alert("Password or user is null");
   	    return "";
   }else{
	  gPassword = hex_md5(gPassword);
      var dom = '{"cmd":1,"param":{"name":"' + gUser + '","pass":"' + gPassword + '"}}';
	  ws.send(dom);
    }
}



//////////////////////
/////登入
function login(pJson) { 
  document.cookie = 'session='+pJson.session ;
  document.cookie = 'name='+document.getElementById("user").value;
  document.cookie = 'privilige='+pJson.privilige;
 location.href = 'mgr.html';
}
////////////////////////
//接收消息
function selectCmdPram(pJson){

switch(pJson.cmd){
  case 1:
  login(pJson);
  break;
  case 10:
  alert(pJson.cmd+pJson.code+pJson.info);
  //查询sql的结果
  break;
    }

}
