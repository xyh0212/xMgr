// JavaScript Document
//////////////////////////////////////////////////////////////
//websoket连接
//test()
var obg  //CheckBoxGroup
window.onload=function(){
    //var json = {"table_id":1000};
    //HttpConn("/table",json,initConn);
    WebSocketConn(initConn)
    var form = document.getElementById("form2");
    form.innerHTML=null
    addClick()
}
function initConn(ws,data){
    if (data.cmd==5){
        ws.send('{"cmd":15}')
        //alert(data.info)
        return
    }
    if (data.cmd==1){
        var  info= JSON.parse(data.info);
        obg = new CheckBoxGroupHanle(info.info)
        document.getElementById("content").appendChild(obg.CheckBoxGroup)
        return
    }
    if (data.cmd==0){
        alert(data.info)
        return
    }
}


document.onchange=function(e){
    obg.onclick(e)
}

//格式：arry_s_gs_id=1:2:3:4
function getCheckedStr() {
    var boxs = obg.getValue()
    var str = 'info=';
    for (var i = 0; i < boxs.length; i++) {
        if (i == boxs.length - 1) {
            str += boxs[i].toString();

        } else {
            str += boxs[i].toString() + ':';
        }
    }
    return str;
}



