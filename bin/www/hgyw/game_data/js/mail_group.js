// JavaScript Document
//////////////////////////////////////////////////////////////
/////事件//////////////////////////////////////////////////////////////
//启动时加载
var gCheckBoxGroupHanle
var gDropListHandle
window.onload=function(){
    //innitData()
    var json1 = {"table_id":1};
    HttpConn("/id_name",json1,initConn);
    var json2 = {"table_id":2};
    HttpConn("/id_name",json2,initGs);
    //packageClick()

}
function initGs(data){
    if (data.cmd==1){
        var  info= JSON.parse(data.info);
        gCheckBoxGroupHanle = new CheckBoxGroupHanle(info.info)
        document.getElementById("gs").appendChild(gCheckBoxGroupHanle.CheckBoxGroup)
        //createGsCheckBox(info.info)
        return
    }
    if (data.cmd==10){
        alert(data.info)
        return
    }
}

function initConn(data){
    if (data.cmd==1){
        var  info= JSON.parse(data.info);
        gDropListHandle=new DropListHandle(info.info)
        document.getElementById("gift_type").appendChild(gDropListHandle.DropList)
        //createDropList(info.info)
        return
    }
    if (data.cmd==10){
        alert(data.info)
        return
    }
}
document.onchange=function(e){
    gCheckBoxGroupHanle.onclick(e)
}


function rec(data){
    if (data.cmd==1){
        alert(data.info);
        return
    }
    if (data.cmd==10){
        document.write(data.info)
        return
    }
}

//////////事件//////////////////////////////////////////////////////////////////////
function sendEmailClick(){
    var radio = document.getElementById("radio")
    var json ='{"gs":"'+gCheckBoxGroupHanle.getValue()+'",'
   // json +='","user_id":"'+document.getElementById("uid").innerHTML+'",'
    json+='"title_txt":"'+document.getElementById("title_txt").value+'",'
    json+='"desc_txt":"'+document.getElementById("desc_txt").value+'",'
    if (radio.checked==true){
        json+='"item_type_main":"999999",'
        json+='"item_type_param":"'+document.getElementById("package_id").value+'",'
        json+='"num":"1"}'
        //alert(json)

    }else {
        json+='"item_type_main":"'+gDropListHandle.getValue()+'",'
        json+='"item_type_param":"0",'
        json+='"num":"'+document.getElementById("gift_count").value+'"}'
    }

    HttpConn("/e_mail_group",JSON.parse(json),rec);

}

//insert into d_package_item(user_id,package,title_id,desc_id,itemtype_main,itemtype_param,num,time_t,title_txt,desc_txt)
//SELECT id,1,0,0,物品主类型,物品参数,物品个数,UNIX_TIMESTAMP(),标题文字，内容文字
//from d_user where login_time+30*24*60*60>UNIX_TIMESTAMP()
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