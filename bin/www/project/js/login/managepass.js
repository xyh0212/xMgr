// JavaScript Document
//////////////////////////////////////////////////////////////
//websoket����
var ws
function NewWebSocket() {

    if ("WebSocket" in window) {
        // Let us open a web socket
        ws = new WebSocket("ws://127.0.0.1:8888/ws");
        ws.onopen = function() { 
        sendLoginJsonStr();	
        };
        //���յ�����˵�����
        ws.onmessage = function(evt) {
        	 var received_msg = evt.data;
             var pJson = JSON.parse(received_msg);
        	  selectCmdPram(pJson);//��ѯ���ĸ�cmd
    	};
        ws.onclose = function() {
            //alert("closed");
        };
   
    }else {
        // The browser doesn't support WebSocket
        alert("�ⲻ��websocket");
    }
}
//���͵����json
function sendLoginJsonStr() {
	var gPassword = document.getElementById("password").value;
    var gUser= document.getElementById("user").value;
   if(gPassword=="" && gUser==""){
   	    alert("Password or user is null");
   	    return "";
   }else{
	  gPassword = hex_md5(gPassword);
	  //document.write(hex_md5("1"))
	  //alert( );
      var dom = '{"cmd":1,"param":{"name":"' + gUser + '","pass":"' + gPassword + '"}}';
	  ws.send(dom);
    }
}



//////////////////////
/////����
function login(pJson) { 
  document.cookie = 'session='+pJson.session ;
  document.cookie = 'name='+document.getElementById("user").value;
  document.cookie = 'privilige='+pJson.privilige;
  //alert(getCookie("session"));
 location.href = 'mgr.html'; 
}
////////////////////////
//������Ϣ
function selectCmdPram(pJson){

switch(pJson.cmd){
  case 1:
  login(pJson);
  break;
  case 10:
  alert(pJson.cmd+pJson.code+pJson.info);
  //��ѯsql�Ľ��
  break;
    }

}



//发送 post
function HttpConnetion(filepath){
    //  filepath="http://localhost:8080/"
    var gPassword = document.getElementById("password").value;
    var gUser= document.getElementById("user").value;
    var  session =	"dsgk"
    var path = "/post?path="+filepath+"&"+"sdgflskd"+"&session="+session
    path="http://localhost:8080"
//alert(path);
    var json = {"用户名":gUser,"密码":gPassword};
    // alert(path);
    $(document).ready(function(){
        $.post(path,json, function(data,status){
            alert(data);
            //  selectCmdPram(data)
            // document.write(data);
            // alert("数据：" + data + "\n状态：" + status);
        });
    });
}