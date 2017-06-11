// JavaScript Document
//////////////////////////////////////////////////////////////
//websoket连接
//test()
var TheArray = [["0-1", "0-2"], ["1-1", "1-2"], ["2-1", "2-2"]]
var obg = new CheckBoxGroupHanle(TheArray)
function alertA() {

    document.getElementById("game_data").appendChild(obg.CheckBoxGroup)
    // alert(obg.content());
}
alertA()
obg.click=function(){
    alert("jjj")
}
function test() {
    alert(obg.getContent());
    alert(obg.getValue());
    document.getElementById("package_id").placeholder = "xyh"
    // alertA()
}
document.onchange=function(e){
    obg.onclick(e)
}
//arrInfo:是一个二维数组 var arrInfo = [["0-1","0-2"],["1-1","1-2"],["2-1","2-2"]]
// var obj = new  DropListHandle();
//obj.DropList //为一个下拉框
//obj.getContent() //为一个下拉框选中的内容
//obj.getValue() //为一个下拉框选中的值
function DropListHandle(arrInfo) {
    var pDropList = document.createElement("select")

    this.DropList = create(arrInfo)

    this.getContent = function () {
        var index = pDropList.selectedIndex;
        return pDropList.options[index].text;
    }

    this.getValue = function () {
        var index = pDropList.selectedIndex;
        return pDropList.options[index].value;
    }

    function create(arr) {
        for ( var j = 0; j < arr.length; j++ ) {
            var op = document.createElement("option")
            op.value = arr[j][0]
            op.innerHTML = arr[j][1];
            pDropList.appendChild(op);
        }
        return pDropList
    }
}

function CheckBoxGroupHanle(arrInfo) {
    var pDiv = document.createElement("div")
    this.CheckBoxGroup=create(arrInfo)

    this.getContent = function () {
        var arr = new Array()
        var arrcheck= pDiv.getElementsByTagName("input")
        var arrNode=pDiv.getElementsByTagName("b")
        for (var i=1 ;i<arrcheck.length;i++){
            if (arrcheck[i].checked){
                arr.push(arrNode[i].innerHTML)
            }
        }
        return arr
    }

    this.onclick= function(e){
        var  check= pDiv.childNodes[0].childNodes[0];
        if (e.target!=check){
            return
        }
        var arrcheck= pDiv.getElementsByTagName("input")
        if (check.checked){
            for (var i=1 ;i<arrcheck.length;i++){
                arrcheck[i].checked=true
            }
        }
        if (!check.checked){
            for (var i=1 ;i<arrcheck.length;i++){
                arrcheck[i].checked=false
            }
        }

    }
    this.getValue = function () {
        var arr = new Array()
        var arrcheck= pDiv.getElementsByTagName("input")
        var arrNode=pDiv.getElementsByTagName("span")
        for (var i=1 ;i<arrcheck.length;i++){
            if (arrcheck[i].checked){
                arr.push(arrNode[i].value)
            }
        }
        return arr
    }

    function create(arr) {
        for ( var j = -1; j < arr.length; j++ ) {
            var span = document.createElement("span")
            var checkBox = document.createElement("input")
            var b = document.createElement("b")
            checkBox.type = "checkbox"
            if (j==-1){
                b.innerHTML ="全部";
                span.appendChild(checkBox);
                span.appendChild(b);
                pDiv.appendChild(span);
                continue
            }
            checkBox.name="checkBoxName"
            span.value = arr[j][0]
            b.innerHTML = arr[j][1];
            span.appendChild(checkBox);
            span.appendChild(b);
            pDiv.appendChild(span);
        }
        return pDiv
    }
}

