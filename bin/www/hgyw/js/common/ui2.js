//class
//arrInfo:是一个二维数组 var arrInfo = [["0-1","0-2"],["1-1","1-2"],["2-1","2-2"]]
// var obj = new  DropListHandle();
//obj.DropList //为一个下拉框
//obj.getContent() //为一个下拉框选中的内容
//obj.getValue() //为一个下拉框选中的值
function DropListHandle(arrInfo) {
    var pDropList = create(arrInfo)
    var pArrInfo  = arrInfo
    this.getAllValue=function(){
        var arr = new Array()
        for ( var i = 0; i < pArrInfo.length; i++ ) {
            arr[i] = pArrInfo[i][0]
        }
        return arr
    }

    this.setItemAll = function () {
        var op = document.createElement("option")
        op.innerHTML = "全部";
        op.value = "";
        if ( !pDropList.hasChildNodes() ) {
            pDropList.appendChild(op);
            return
        }
        pDropList.insertBefore(op, pDropList.childNodes[0])

    }
    this.getContent = function () {
        var index = pDropList.selectedIndex;
        return pDropList.options[index].text;
    }
    this.DropList = pDropList
    this.getValue = function () {
        var index = pDropList.selectedIndex;
        return pDropList.options[index].value;
    }


    function create(arr) {
        var d = document.createElement("select")
        for ( var j = 0; j < arr.length; j++ ) {
            var op = document.createElement("option")
            op.value = arr[j][0]
            op.innerHTML = arr[j][1];
            d.appendChild(op);
        }
        return d
    }
}
//class
// arrS: 数组 var arrS = ["条数","页数"，"总页数"]
function bindEvent(obj, ev, fn){
    if (obj.attachEvent) {
        obj.attachEvent('on' + ev, fn);
    }
    else
    {
        obj.addEventListener(ev, fn, false);
    }
}
function PageHandle(arrS) {

    var arrPage = arrS  //arr: var arr = ["第几页","共几页","共几条"]
    var pPagination = createPagination()

    this.Pagination = pPagination

    this.setPage = function (value) {
        pPagination.getElementsByTagName("input")[0].value = value
    }
    this.getPage = function () {
        return pPagination.getElementsByTagName("input")[0].value
    }
    this.getNumPage = function () {
        return arrPage[1]
    }
////////////////////////////////

    this.innitEvent = function (gPageHandle, func) {
        var obj = new eventFun(gPageHandle, func)
        var arrSpan = pPagination.childNodes[0].getElementsByTagName("span")
        bindEvent(arrSpan[0], "click", obj.startPage);
        bindEvent(arrSpan[1], "click", obj.lastPage);
        bindEvent(arrSpan[2], "click", obj.submit);
        bindEvent(arrSpan[3], "click", obj.nextPage);
        bindEvent(arrSpan[4], "click", obj.endPage);
    }

    function createPagination() {
        pPagination = document.createElement("div")
        pPagination.appendChild(getPagination1())
        pPagination.appendChild(getPagination2())
        pPagination.style.width = "350px"
        return pPagination
    }

    function getPagination1() {
        var pPagination1 = document.createElement("div")
        pPagination1.className = "input-group"
        var arrStr = ["首页", "上一页", "go", "下一页", "尾页"]
        for ( var i = 0; i < arrStr.length; i++ ) {
            pPagination1.appendChild(createSpan(arrStr[i]))
        }
        var text = document.createElement("input")
        text.type = "text"
        text.className = "form-control"
        text.value = arrPage[0]
        pPagination1.insertBefore(text, pPagination1.childNodes[2]);
        return pPagination1
    }

    function getPagination2() {
        var arrStr = ["/", "页,共", "记录"]
        var pPagination2 = document.createElement("div")
        for ( var i = 0; i < arrPage.length; i++ ) {
            var span = document.createElement("span")
            var a = document.createElement("a")
            a.innerHTML = arrPage[i]
            span.innerHTML = arrStr[i]
            pPagination2.appendChild(a)
            pPagination2.appendChild(span)

        }
        return pPagination2
    }

    function createSpan(str) {
        var span = document.createElement("span")
        var a = document.createElement("a")
        a.innerHTML = str;
        span.appendChild(a);
        span.className = "input-group-addon"
        return span;
    }

    function eventFun(gPageHandle, func) {
        this.startPage = function () {
            func(1)
            gPageHandle.setPage(1)
        }
        this.nextPage = function () {
            var num = Number(gPageHandle.getPage())
            if ( num == gPageHandle.getNumPage() ) {
                return
            }
            func(num + 1)
            gPageHandle.setPage(num + 1)
        }
        this.lastPage = function () {
            var num = Number(gPageHandle.getPage())
            if ( num < 1 ) {
                gPageHandle.setPage(0)
                return
            }
            func(num - 1)
            gPageHandle.setPage(num - 1)
        }
        this.submit = function () {
            var page =gPageHandle.getPage()
            if (page>gPageHandle.getNumPage()||page<1){
                alert("可输入的页数范围:\n1-"+gPageHandle.getNumPage()+"\n请从新输入")
                gPageHandle.setPage(1)
                return
            }
            func(page)

        }
        this.endPage = function () {
            var page = gPageHandle.getNumPage()
            func(page)
            gPageHandle.setPage(page)
        }
    }
}

//class
//arrInfo:是一个二维数组 var arrInfo = [["0-1","0-2"],["1-1","1-2"],["2-1","2-2"]]
// var obj = new  DropListHandle();
//obj.CheckBoxGroup //为一个复选框组
//obj.getContent() //复选框组选中的内容 return 字符串数组
//obj.getValue() //复选框组选中的值     return 字符串数组
function CheckBoxGroupHanle(arrInfo) {
    var pCheckBoxGroup = create(arrInfo)
    this.CheckBoxGroup = pCheckBoxGroup

    this.getContent = function () {
        var arr = new Array()
        var arrcheck = pCheckBoxGroup.getElementsByTagName("input")
        var arrNode = pCheckBoxGroup.getElementsByTagName("b")
        for ( var i = 1; i < arrcheck.length; i++ ) {
            if ( arrcheck[i].checked ) {
                arr.push(arrNode[i].innerHTML)
            }
        }
        return arr
    }

    this.onclick = function (e) {
        var check = pCheckBoxGroup.childNodes[0].childNodes[0];
        if ( e.target != check ) {
            return
        }
        var arrcheck = pCheckBoxGroup.getElementsByTagName("input")
        if ( check.checked ) {
            for ( var i = 1; i < arrcheck.length; i++ ) {
                arrcheck[i].checked = true
            }
        }
        if ( !check.checked ) {
            for ( var i = 1; i < arrcheck.length; i++ ) {
                arrcheck[i].checked = false
            }
        }

    }
    this.getValue = function () {
        var arr = new Array()
        var arrcheck = pCheckBoxGroup.getElementsByTagName("input")
        var arrNode = pCheckBoxGroup.getElementsByTagName("span")
        for ( var i = 1; i < arrcheck.length; i++ ) {
            if ( arrcheck[i].checked ) {
                arr.push(arrNode[i].value)
            }
        }
        return arr
    }


    function create(arr) {
        var pDiv = document.createElement("div")
        for ( var j = -1; j < arr.length; j++ ) {
            var span = document.createElement("span")
            var checkBox = document.createElement("input")
            var b = document.createElement("b")
            checkBox.type = "checkbox"
            if ( j == -1 ) {
                b.innerHTML = "全部";
                span.appendChild(checkBox);
                span.appendChild(b);
                pDiv.appendChild(span);
                continue
            }
            checkBox.name = "checkBoxName"
            span.value = arr[j][0]
            b.innerHTML = arr[j][1];
            span.appendChild(checkBox);
            span.appendChild(b);
            pDiv.appendChild(span);
        }
        return pDiv
    }


}
//class
// arrS: 二维数组 var arrS = [["条数","页数"，"总页数"],["条数","页数"，"总页数"],["条数","页数"，"总页数"]]
function TableHandle(arrS)  {
    var pTable = create(arrS)

    this.Table = pTable

    function create(arr) {
        if (!multiarr(arr)) {
           return createTableTitle(arr)
        }
        var table =createTableTitle(arr[0])
        for (var i = 1; i < arr.length; i++) {
            var tr = document.createElement("tr")
            for (var j = 0; j < arr[i].length; j++) {
                var td = document.createElement("td")
                td.innerHTML = arr[i][j]
                tr.appendChild(td)
            }
            table.appendChild(tr);
        }
        return table
    }
    ///创建表的头部字段名
    function  createTableTitle(arr){
        var table = document.createElement("table")
        table.className = "datalist"
        var tr = document.createElement("tr")
        for (var j = 0; j < arr.length; j++) {
            var th = document.createElement("th")
            th.innerHTML = arr[j]
            tr.appendChild(th)
        }
        table.appendChild(tr);
        return table
    }
    this.setFirstRow = function (arr) {
        if (arr.constructor != Array) {
            return
        }
        var node1 = pTable.getElementsByTagName("th")
        for (var i = 0; i < node1.length; i++) {
            node1[i].innerHTML = arr[i];
        }
    }
    this.insertColumn = function (func) {
        var rows = pTable.rows;
        //表格行数
        var rowsLen = rows.length ;
        //表格列数
        //元素和列名，fun
        var cellsLen = rows.item(0).cells.length ;
        func(rows)
        /*

        var th = document.createElement("th");
        th.innerHTML = "检查";
        rows[0].insertBefore(th, rows[0].cells[cellsLen]);
        for (var i = 1; i < rows.length; i++) {
            var td = rows[i].insertCell(cellsLen);
            td.appendChild(ck);
            func(rowsLen,cellsLen)

            var ck = document.createElement("input");
            ck.type = "button";
            ck.value="补单"
            td.appendChild(ck);
        }
*/
    }

    this.insertColumnElement=function(func){
        var rows = pTable.rows;      //行数组
        var rowsLen = rows.length ;     //表格行数
        var cellsLen = rows.item(0).cells.length ;//表格列数
        var arr =["编辑","删除"]
        for(var i=0;i<arr.length;i++){
            var th = document.createElement("th");
            th.innerHTML = arr[i];
            rows[0].insertBefore(th, rows[0].cells[cellsLen+i]);
        }
        for (var i = 0; i <arr.length; i++) {
            for ( var j = 1; j < rows.length; j++ ) {
                var td = rows[j].insertCell(cellsLen+i);
                func(td)
            }
        }
    }


}

function multiarr(arr){
        for (i=0,len=arr.length;i<len;i++)
            if(arr[i] instanceof Array)return true;
        return false;
}
function ElementHandle(){
    this.CreateElement=function(obj){
        var ele = document.createElement(obj.eleName);
        ele.id=obj.id;
        ele.className=obj.name;
        ele.value=obj.value;
        ele.title=obj.title;
        ele.click=obj.click;
        ele.src=obj.src;
       return ele
    }
    this.getElement=function(id){
        return document.getElementById(id)
    }




}