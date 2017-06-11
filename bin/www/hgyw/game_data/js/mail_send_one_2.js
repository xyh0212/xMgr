// JavaScript Document
//////////////////////////////////////////////////////////////
/////事件//////////////////////////////////////////////////////////////
//启动时加载
window.onload=function(){
    innitData()
    var json = {"table_id":1};
    HttpConn("/id_name",json,initConn);
    packageClick()
    // insert into d_package_item(`user_id`,`package`,`title_txt`,`desc_txt`,`itemtype_main`,`itemtype_param`,`num`,`time_t`)
    // values ("UserID","1","标题文字","内容文字","物品主类型","物品参数","个数",UNIX_TIMESTAMP())
}

function initConn(data){
    if (data.cmd==1){
        var  info= JSON.parse(data.info);
        createDropList(info.info)
       // alert(info.info);
        return
    }
    if (data.cmd==0){
        alert(data.info)
        return
    }
}


//创建下拉菜单
function createDropList(arrInfo){
    var obj = document.getElementById("select_gs");
    obj.innerHTML=null;
    for (var j=0;j<arrInfo.length;j++){
        var op= document.createElement("option")
        op.value=arrInfo[j][0]
        op.innerHTML= arrInfo[j][1];
        obj.appendChild(op);
    }
}
function innitData(){
    document.getElementById("gs").innerHTML=getUrlValue("gs");
    document.getElementById("uid").innerHTML=getUrlValue("tag0");
    document.getElementById("player_name").innerHTML=getUrlValue("tag1");
    document.getElementById("player_level").innerHTML=getUrlValue("tag2");
}
function rec(data){
    if (data.cmd==1){
        alert(data.info);
        return
    }
    if (data.cmd==10){
        alert(data.info)
        //document.write(data.info)
        return
    }
}

//////////事件//////////////////////////////////////////////////////////////////////
function sendEmailClick(){
    var myselect=document.getElementById("select_gs");
    var index=myselect.selectedIndex;
    var s1 =myselect.options[index].value;
    var s2 =myselect.options[index].text;
    var radio = document.getElementById("radio")
    var json ='{"gs":"'+getUrlValue("gsid")+'","user_id":"'+document.getElementById("uid").innerHTML+'",'
    json+='"title_txt":"'+document.getElementById("title_txt").value+'",'
    json+='"desc_txt":"'+document.getElementById("desc_txt").value+'",'
    if (radio.checked==true){
        json+='"item_type_main":"999999",'
        json+='"item_type_param":"'+document.getElementById("package_id").value+'",'
        json+='"num":"1"}'

    }else {
        json+='"item_type_main":"'+s1+'",'
        json+='"item_type_param":"0",'
        json+='"num":"'+document.getElementById("gift_count").value+'"}'
    }

    HttpConn("/e_mail_skip",JSON.parse(json),rec);

}
//{"gs":"1","user_id":"1600001000001","title_txt":"请输入邮件标题","desc_txt":"请输入邮件内容","item_type_main":"999999","item_type_param":"111111","num":"1"}
function giftClick(){
    $(function(){
        $('#gift_count').attr('disabled',false);
        $('#select_gs').attr('disabled',false);
        $('#package_id').attr('disabled',true);
    });

}

function packageClick(){
    $(function(){
        $('#gift_count').attr('disabled',true);
        $('#select_gs').attr('disabled',true);
        $('#package_id').attr('disabled',false);
    });
}