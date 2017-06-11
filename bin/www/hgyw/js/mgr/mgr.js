////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//动态的绑定事件  onmousedown
var gArrEleStr   //titlebar的显示路径
window.onload = function() {
    document.getElementById("show_user").innerHTML=getCookie("name");
    NewWebSocket();
    initUi();
    document.getElementById("home_id").style.display="block";
    //test();
}
function test(){

     var test = document.getElementById("test1")
     test.childNodes[1].style.display="none";
     return;
}
function initUi(){
    /*
    if (getCookie("privilige")!=1){
        var arrHeader =document.getElementsByName("header");
        arrHeader[2].style.display="none";
        arrHeader[3].style.display="none";
        //return;
    }*/
    var arrSideListDiv =document.getElementsByName("side_list");
    for (var i=0;i<arrSideListDiv.length;i++){
        arrSideListDiv[i].style.display="none";
    }

}
document.onmousedown = function(e){
   if (e.target.name=="header"){
       var arrSideListDiv =document.getElementsByName("side_list");
       for (var i=0;i<arrSideListDiv.length;i++){
        if(arrSideListDiv[i].title==e.target.title){
            showSideBarList(arrSideListDiv,i);
            createPathCatalog(e.target);
            createTitlePathElement();
            return;
        }

           if(e.target.title=="home"){
               e.target.style.display="block";
              // for (var i= 0;i<arrSideListDiv.length;i++){
                //   arrSideListDiv[i].style.display="block";
              // }
           }
       }
   }
    if (e.target.name=="side_list_third_a"){
        switch (e.target.id){
////////////////游戏数据//game_data////////////////////////////////////////////////////////////
            case "exc_sql":
                includeHtml("game_data/gs_sql.html");
                break;
            case "exc_sql2":
                includeHtml("game_data/gs_sql2.html");
                break;
            case  "exc_file_sql":
                includeHtml("game_data/gs_sqlfile.html");
                break;
            case  "exc_file_sql2":
                includeHtml("game_data/gs_sqlfile2.html");
                break;
            case  "emil_single_send":
                includeHtml("game_data/mail_send_one_1.html");
                break;
            case  "emil_group_send":
                includeHtml("game_data/mail_send_group.html");
                break;
            case  "player_list":
                includeHtml("game_data/player_list.html");
                break;
            case  "player_role":
                includeHtml("game_data/player_role.html");
                break;
            case  "player_channel_number":
                includeHtml("game_data/player_channel_number.html");
                break;
            case  "player_close":
                includeHtml("game_data/player_close.html");
                break;
            case  "player_resolve":
                includeHtml("game_data/player_resolve.html");
                break;
//////////////////////////////////////////////////////////////////////////////add_gs_mgr
            case  "user_mgr":
                includeHtml("user_mgr.html");
                break;
            case  "gs_mgr":
                includeHtml("gs/gs_mgr.html");
                break;
            case  "add_gs_mgr":
                includeHtml("gs/add_gs_mgr.html");
                break;

            case  "test":
                includeHtml("test.html");
                break;
///////////////////////充值管理//recharge////////////////////////////////////////////////////
            case  "query_order":
                includeHtml("recharge/query_order.html");
                break;
        }
        gArrEleStr.push(e.target.innerHTML)
        createTitlePathElement();
        gArrEleStr.pop()
    }
}
//验证用户是否过期，过期返回登入界面
//pJson 为后端传来的json
function selectCmdPram(pJson){
    if (pJson.cmd==10){
        alert("用户过期，请从新登入");
        document.location="login.html";
    }
}
//控制左边的显示列表
//arrTitleA:title_bar的‘a‘元素集合
//index：是arrTitleA的指定下标
function showSideBarList(arrSideListDiv,index){
    for (var i= 0;i<arrSideListDiv.length;i++){
        arrSideListDiv[i].style.display="block";
    }
    for (var i= 0;i<arrSideListDiv.length;i++){
        if (i!=index){
            arrSideListDiv[i].style.display="none";
        }
        if (i==index){
            showTitleBarEleBackgroud(arrSideListDiv[i]);
        }
    }
}
//设置title元素点击后的背景
//titleA:title_bar的‘a‘元素集合
function showTitleBarEleBackgroud(sideLeftDiv){
    var arrTitleLi = document.getElementsByName("title_li");
    for (var i= 0;i<arrTitleLi.length;i++){
        arrTitleLi[i].className=null;
    }
    for (var i= 0;i<arrTitleLi.length;i++){
        if (arrTitleLi[i].title==sideLeftDiv.title){
            arrTitleLi[i].className="active";

        }
    }
}
//titleA:titleBar元素<a>的的应用
//return:Array ["首页"，"SQL","执行SQL"]
function createPathCatalog(titleA){
    gArrEleStr= new Array();
    if(titleA.title!="home"){
        gArrEleStr.push("首页")
    }
    gArrEleStr.push(titleA.innerHTML)
    var obj=document.getElementsByName("side_list_second_a")
    for (var i= 0;i<obj.length;i++){
        if (titleA.title==obj[i].title){
            gArrEleStr.push(obj[i].innerHTML)
        }
    }
}

function createTitlePathElement(){
    var gTitlePathDiv = document.createElement("div");
    for (var i=0;i<gArrEleStr.length;i++){
        var a= document.createElement("a");
        var b= document.createElement("b");
        a.innerHTML=gArrEleStr[i];
        b.innerHTML=">";
        gTitlePathDiv.appendChild(a);
        gTitlePathDiv.appendChild(b);
    }
    document.getElementById("path_list_name").innerHTML=null;
    document.getElementById("path_list_name").appendChild(gTitlePathDiv);
}

function includeHtml(src){
    var obj=document.getElementById("content_right")
    obj.innerHTML=null;
    var frame=document.createElement("iframe")
    frame.src=src;
    frame.height=800
    frame.width="80%"
    frame.frameBorder="0"
    obj.appendChild(frame);


}