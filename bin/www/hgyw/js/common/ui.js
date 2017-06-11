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
    this.click=function(func){
        bindEvent(pDropList, "change", func);
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
    this.IdAndName=pArrInfo

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
    var pArrInfo=arrInfo
    //var pVorH=""//默认为横向
    var pCheckBoxGroup = create(arrInfo,"span")



    this.setVertical=function(){
        pVorH="div"
        pCheckBoxGroup=null
        this.CheckBoxGroup =create(pArrInfo,"div")

    }
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


    function create(arr,pVorH) {

        var pDiv = document.createElement("div")
        for ( var j = -1; j < arr.length; j++ ) {
            var span = document.createElement(pVorH)
            //alert(pVorH)
            var checkBox = document.createElement("input")
            var b = document.createElement("b")
            checkBox.type = "checkbox"
            if ( j == -1 ) {
                var br = document.createElement("br")
                b.innerHTML = "全部";
                span.appendChild(checkBox);
                span.appendChild(b);
                span.appendChild(br);
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
function TableHandle(arrS) {
    var pTable = create(arrS)

    this.Table = pTable

    function create(arr) {
        if (!multiarr(arr)) {
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

        var table = document.createElement("table")
        table.className = "datalist"
        var tr = document.createElement("tr")
        for (var j = 0; j < arr[0].length; j++) {
            var th = document.createElement("th")
            th.innerHTML = arr[0][j]
            tr.appendChild(th)
        }
        table.appendChild(tr);
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
    //设置表的头部字段名
    this.setFirstRow = function (arr) {
        if (arr.constructor != Array) {
            return
        }
        var node1 = pTable.getElementsByTagName("tr")[0].cloneNode(true);
        var rowLen = pTable.rows.length
        if  (rowLen<2){
            for (var i = 0; i < node1.childNodes.length; i++) {
                node1.childNodes[i].innerHTML = arr[i]
            }
            pTable.insertBefore(node1, pTable.getElementsByTagName("tr")[0])
            pTable.deleteRow(1)
            return
        }
        var node2 = pTable.getElementsByTagName("tr")[1].cloneNode(true);
        for (var i = 0; i < node1.childNodes.length; i++) {
            node2.childNodes[i].innerHTML = node1.childNodes[i].innerHTML
        }
        pTable.insertBefore(node2, pTable.getElementsByTagName("tr")[1])
        for (var i = 0; i < node1.childNodes.length; i++) {
            node1.childNodes[i].innerHTML = arr[i]
        }
        pTable.replaceChild(node1, pTable.getElementsByTagName("tr")[0]);
        pTable.deleteRow(1)
    }
    //插入列自己实现
    this.insertColumn = function (func) {
        var rows = pTable.rows;
        //表格行数
        var rowsLen = rows.length ;
        //表格列数
        //元素和列名，fun
        var cellsLen = rows.item(0).cells.length ;
        func(rows)
    }
    //插入列的button
    this.insertColumBtn=function(){
        var rows = pTable.rows;
        var rowsLen = rows.length ;//表格行数
        var cellsLen = rows.item(0).cells.length ;//表格列数
        var arr = ["删除","编辑"]
        var arrClassName = ["deleteName","editName"]
        for (var i=0;i<arr.length;i++){
            var th=document.createElement("th")
            th.innerHTML=arr[i]
            rows[0].insertBefore(th, rows[0].cells[cellsLen+i]);
        }
        for (var i=0;i<arr.length;i++){
            for (var j=1;j<rowsLen;j++){
                var td = rows[j].insertCell(cellsLen+i)
                var btn =document.createElement("input")
                btn.type="button"
                btn.value=arr[i]
                btn.name=arrS[j][0]
                btn.className=arrClassName[i]
                td.appendChild(btn)
            }
        }
    }
    //插入行的text
    this.insertRowText=function(id) {
        var rows = pTable.rows;
        var rowsLen = rows.length;//表格行数
        var cellsLen = rows.item(0).cells.length - 2;//表格列数
        for (var i = 1; i < rowsLen; i++) {
            var value =rows[i].cells[cellsLen+1].childNodes[0]
            if (value.value=="保存"){
                reNewTd(rows[i].cells)
                value.value="编辑"
                ///value.name="btnName"
                value.className="editName"
            }
        }
        for (var i = 1; i < rowsLen; i++) {
            var value = rows[i].cells[0].innerHTML
            if (value==id){
                initText(rows,i);
                var btn = rows[i].cells[cellsLen+1].childNodes[0]
                btn.value="保存"
                btn.className="save"
                return
            }
        }
    }
    //
    this.getRowTextValue=function(){
        var rows = pTable.rows;
        var rowsLen = rows.length;//表格行数
        var cellsLen = rows.item(0).cells.length - 2;
        for (var i = 1; i < rowsLen; i++) {
            var text = rows[i].cells[cellsLen + 1].childNodes[0]
            if (text.className =="save"){
               // alert(text.className)
                var cells= rows[i].cells
                var arr = []
                for (var i=1;i<cells.length-2;i++){
                    arr.push(cells[i].childNodes[0].value)
                    //cells[i].innerHTML= cells[i].childNodes[0].value
                }
                return arr
            }
        }
    }
    ///把文本框中的值还原到td
    function reNewTd(cells){
        for (var i=1;i<cells.length-2;i++){
            cells[i].innerHTML= cells[i].childNodes[0].value
        }
    }
    ///把td中的值放到文本框中
    var arrSize=getRowMaxNumChar(arrS)
    this.setTextSize=function(arr){
        arrSize =arr
    }
    function initText(rows,i){
        var fontLen =getRowMaxNumChar(arrS)
        var cellsLen = rows.item(0).cells.length - 2;//表格列数
            for (var j = 1; j < cellsLen; j++) {
                var td = rows[i].cells[j]
                var btn = document.createElement("input")
                if (td.value==null){
                    td.value=td.innerHTML
                }
                td.innerHTML=null
                btn.type = "text"
                btn.size=arrSize[j]
                btn.value = td.value
                td.appendChild(btn)
            }

    }
    this.click=function(e){
        var name = e.target.className
        if (name=="deleteName"){
            alert(e.target.name)
        }else if (name=="editName"){
            this.insertRowText(e.target.name)
        }

    }
    //返回每列最长的字符串长度
    function getRowMaxNumChar(Gjson){
        var arrLen = new Array()
        for (var i = 0; i <Gjson[0].length; i++) {
            var len = 1
            for (var j = 0; j <Gjson.length; j++) {
                if(Gjson[j][i].length>len){
                    len = Gjson[j][i].length
                }
            }
            arrLen[i]=len
        }
        return arrLen
    }
}

function multiarr(arr){
        for (i=0,len=arr.length;i<len;i++)
            if(arr[i] instanceof Array)return true;
        return false;
}