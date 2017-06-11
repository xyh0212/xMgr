// JavaScript Document
//////////////////////////////////////////////////////////////
/////事件//////////////////////////////////////////////////////////////
//启动时加载
window.onload=function(){
    var json = {"table_id": 2};
    HttpConn("/id_name", json, initConn);
}

var gDropListHandle
function initConn(data){
    if (data.cmd==1){
        var  info= JSON.parse(data.info);
        gDropListHandle = new DropListHandle(info.info)
        document.getElementById("gs").appendChild(gDropListHandle.DropList);
        //createOption(info.info);
        return
    }
    if (data.cmd==0){
        alert(data.info)
        return
    }
}
//查询玩家的点击事件
//查询玩家：select id,name,level from d_user where name="玩家名";
function gamePlayerClick(){
    var text = document.getElementById("player_name");
    alert(text.value.trim());
    //return
    HttpConn("/e_mail",{"gs_id":gDropListHandle.getValue(),"name":text.value.trim()},loginSkipEmali);
}

function  loginSkipEmali(data){
    if (data.cmd==1){
        var  json= JSON.parse(data.info);
        var url =packageUrl('mail_send_one_2.html',json.info)
        window.location=url+"&gs="+gDropListHandle.getContent()+'&gsid='+gDropListHandle.getValue()
        return
    }
    alert(data.info)

}



//组装url
//key为tag1，tag2.。。。。。。。
//arr：字符串数组为get的参数
function packageUrl(url,arr){
    var str = url+"?";
    var tag = "tag";
    for (var i=0;i<arr.length;i++){
        str +=tag+ i.toString()+'='+arr[i].toString()+"&"
    }
    str = str.substr(0,str.length-1)
   return str;
}

