var gDropListGsGroup //
var gCheckBoxGroupHanleGs//
var gArrGs
window.onload = function () {
    var json = {"table_id": 4};
    HttpConn("/id_name", json, initGsGroup);
    WebSocketConn(initGs)
}
function initGsGroup(data) {
    if (data.cmd == 1) {
        var info = JSON.parse(data.info);
        //alert(info.info)
        gDropListGsGroup = new DropListHandle(info.info)
        document.getElementById("gs_group").appendChild(gDropListGsGroup.DropList);
        gDropListGsGroup.click(onChange)
        return
    }
    if (data.cmd == 10) {
        alert("错误："+data.code+"\n信息："+data.info)
        return
    }
}
function initGs(ws,data) {
    if (data.cmd == 5) {
        ws.send('{"cmd":15}')
        return
    }
    if (data.cmd == 15) {
        var json =JSON.parse(data.info)
        gArrGs=json.json
        //var s = gDropListGsGroup.getValue()
        onChange()
        return
    }
    if (data.cmd == 10) {
        alert("错误："+data.code+"\n信息："+data.info)
        return
    }
}
document.onclick=function(e){
    gCheckBoxGroupHanleGs.onclick(e)
}
function onChange(){
    var s =gDropListGsGroup.getValue()
    var arr = new Array()
    for (var i=0;i<gArrGs.length;i++){
        if (s==gArrGs[i][3]){
            arr.push([gArrGs[i][0],gArrGs[i][2]])
        }
    }
    gCheckBoxGroupHanleGs =  new CheckBoxGroupHanle(arr)
    document.getElementById("gs").innerHTML=null
    document.getElementById("gs").appendChild(gCheckBoxGroupHanleGs.CheckBoxGroup);

}
//点击按钮时发送数据给连接的html
function selectBtnClick() {
    var sql = document.getElementById("textarea").value;
    sql = encodeURIComponent(sql);
    //document.getElementById("allCheckbox").checked = false;
    var str = "../show_table.html" + '?CheckedJsonStr=' + getCheckedJsonStr() + '&sqlJson=' + sql;
    window.open(str);
    return str;
}
//格式："arry_s_gs_id" :[1,1,1]
function getCheckedJsonStr() {
    var boxs = new Array();
    var str = '"arry_s_gs_id": [';
    boxs = gCheckBoxGroupHanleGs.getValue();
    for (var i = 0; i < boxs.length; i++) {
        if (i == boxs.length - 1) {
            str += '"' + boxs[i].toString() + '"';
        } else {
            str += '"' + boxs[i].toString() + '",';
        }
    }
    str += ']';
    return str;
}