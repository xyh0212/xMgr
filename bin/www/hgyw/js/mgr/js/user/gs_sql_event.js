function checkBoxAllClick(thisCheckbox) {
    if (thisCheckbox.checked) {
        for (var i = 0; i < checkboxs.length; i++) {
            checkboxs[i].checked = true;
        }
    } else {
        for (var i = 0; i < checkboxs.length; i++) {
            checkboxs[i].checked = false;
        }
    }
}
//点击按钮时发送数据给连接的html
function selectBtnClick(ss) {
    var sql = document.getElementById("textarea").value;
    sql = encodeURIComponent(sql);
    document.getElementById("allCheckbox").checked = false;
    var str = URL_SHOW_TABLE + '?CheckedJsonStr=' + getCheckedJsonStr() + '&sqlJson=' + sql;
    window.open(str);
    return str;
}

//格式："arry_s_gs_id" :[1,1,1]
function getCheckedJsonStr() {
    var boxs = new Array();
    var str = '"arry_s_gs_id": [';
    boxs = getCheckedValue();
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

//格式：arry_s_gs_id=1:2:3:4
function getCheckedStr() {
    var boxs = new Array();
    var str = 'info=';
    boxs = getCheckedValue();
    for (var i = 0; i < boxs.length; i++) {
        if (i == boxs.length - 1) {
            str += boxs[i].toString();

        } else {
            str += boxs[i].toString() + ':';
        }
    }
    return str;
}
//查找选中的checkbox
function getCheckedValue() {
    var boxs = new Array();
    var boxs = [];
    var j = 0;
    for (var i = 0; i < checkboxs.length; i++) {
        //alert(checkboxs.length);
        if (checkboxs[i].checked) {
            boxs[j] = Number(checkboxs[i].value);
            j++;
            //alert(boxs[i]);
        }
    }
    if (boxs.length == 0) {
        alert("您没有选中区域");
        return;
    }
    //alert(boxs[1]);
    return boxs;

}

//选择的是哪个
var ProductValue = "玩吧三国";//产品
function setProductValue(osel) {
    document.getElementById("allCheckbox").checked = false;
    var str = "";
    ProductValue = osel.options[osel.selectedIndex].text;
    AppendDiv(ProductValue);
}

function testClick() {

    window.open("test.html");

}













