// JavaScript Document
//////////////////////////////////////////////////////////////
var  gGsid  //选择时的gsid
/////事件//////////////////////////////////////////////////////////////
//启动时加载
window.onload=function(){

    var json = {"id":"Donald Duckasksajdhgosadhohd","city":"Duckburg"};
    HttpConn("/send_gs",json,initConn);

}
//游服改变的事件
function gsOnChange(obj){
    gGsid = obj.value
    //ProductValue = osel.options[osel.selectedIndex].text;
   // alert(obj.options[obj.selectedIndex].text)
    gGsName =document.getElementById("select_gs")
    //alert(s.innerText)

}
function initConn(data){
    if (data.cmd==1){
        var  info= JSON.parse(data.info);
        createOption(info.info);
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
    HttpConn("/e_mail",{"gs_id":gGsid,"name":text.value.trim()},loginSkipEmali);
}

function  loginSkipEmali(data){
    if (data.cmd==1){
        var  json= JSON.parse(data.info);
        var url =packageUrl('mail_send_one_2.html',json.info)
        var myselect=document.getElementById("select_gs");
        var index=myselect.selectedIndex;
        var gGsName =myselect.options[index].text;
        //alert(gGsName)
        window.location=url+"&gs="+gGsName+'&gsid='+gGsid
        return
    }
    alert(data.info)

}
/////初始化UI//////////////////////////////////////////////////////////////
function createOption(arrInfo){
    var obj = document.getElementById("select_gs");
    obj.innerHTML=null;
    gGsid = arrInfo[0][0]
    gGsName=arrInfo[0][1]
        for (var j=0;j<arrInfo.length;j++){
            var op= document.createElement("option")
            op.value=arrInfo[j][0]
            op.innerHTML= arrInfo[j][1];
            obj.appendChild(op);
        }

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

