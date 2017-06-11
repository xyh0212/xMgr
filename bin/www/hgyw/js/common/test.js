// JavaScript Document

function dateClick(){
    var obj =document.getElementById("date").value
    obj=obj.replace("T"," ")
    date = obj.substring(0,19)
    date = date.replace(/-/g,'/')
    var timestamp = new Date(date).getTime()
    alert(formatDate(timestamp/1000))

}
window.onload = function () {
}
document.onclick=function(e){
    var name = e.target.className
    if (name=="deleteName"){
        var s=gTableHandle.getRowTextValue()
        alert(s)
    }else if (name=="editName"){
        gTableHandle.insertRowText(e.target.name)
    }else if (name=="save"){
        var s=gTableHandle.getRowTextValue()
        alert(s)
    }
}
///PaginationHandle
////////////////////////////////////
var gPageHandle
function btngPaginationHandleClick(){
    var arrI = [1, 2, 3]
    gPageHandle =new PageHandle(arrI)
    gPageHandle.innitEvent(gPageHandle,func)
    document.getElementById("pagination").appendChild(gPageHandle.Pagination)
}
function func(e){
    alert(e)
}

//////////////////////////////////
//TableHandle
var gTableHandle
function btnTableHandleClick(){
    var arrS = [["1", "2", "3"], ["4", "5", "6"], ["7", "8", "9"]]
    gTableHandle = new TableHandle(arrS)
    var ar = ["id","field1","field2"]
    gTableHandle.setFirstRow(ar)
    gTableHandle.insertColumBtn()

    document.getElementById("table").appendChild(gTableHandle.Table)
}
function test() {
    alert("hello")
}

bindEvent = function (obj, ev, fn) {
    if ( obj.attachEvent ) {
        obj.attachEvent('on' + ev, fn);
    }
    else {
        obj.addEventListener(ev, fn, false);
    }
}


/////////////////////////////////////////
//CheckBoxGroupHanle
var gCheckBoxGroupHanle

function btnCheckBoxGroupHanleClick() {
    var TheArray = [["0-1", "0-2"], ["1-1", "1-2"], ["2-1", "2-2"]]
    gCheckBoxGroupHanle = new CheckBoxGroupHanle(TheArray)
    document.getElementById("checkbox_group").appendChild(gCheckBoxGroupHanle.CheckBoxGroup)
}
document.onchange=function(e){
    gCheckBoxGroupHanle.onclick(e)
   // alert("sss")
}
/////////////////////////////////////////
//DropListHandle
var gDropListHandle
function btnDropListHandleClick() {
    var TheArray = [["0-1", "0-2"], ["1-1", "1-2"], ["2-1", "2-2"]]
    gDropListHandle = new DropListHandle(TheArray)
    document.getElementById("drop_list").appendChild(gDropListHandle.DropList)
}