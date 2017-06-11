// JavaScript Document
//////////////////////////////////////////////////////////////
//websoket连接
//test()
var obg  //CheckBoxGroup
window.onload=function(){
    var json = {"table_id":2};
    HttpConn("/id_name",json,initConn);

}
function initConn(data){
    if (data.cmd==1){
        var  info= JSON.parse(data.info);
        //createDropList(info.info)
        obg = new CheckBoxGroupHanle(info.info)
        document.getElementById("content").appendChild(obg.CheckBoxGroup)
        //alert("1")
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

function selectBtnClick(){
   var arr = obg.getValue()
    if (arr.length<=0){
        alert("您还没有选中区域")
        return
    }
    var sql = document.getElementById("textarea").value;
    sql = encodeURIComponent(sql);
   // alert(getCheckedJsonStr(arr))
    var str = '../show_table.html' + '?CheckedJsonStr=' + getCheckedJsonStr(arr) + '&sqlJson=' + sql;
    window.open(str);

}
//格式："arry_s_gs_id" :[1,1,1]
function getCheckedJsonStr(boxs) {
    var str = '"arry_s_gs_id": [';
    for (var i = 0; i < boxs.length; i++) {
        if (i == boxs.length - 1) {
            str += '"' + boxs[i] + '"';

        } else {
            str += '"' + boxs[i] + '",';
        }
    }
    str += ']';
    return str;
}