////////////////////////////////
//设置玩家列表的显示表的内容
function setPlayerContent(rows) {
    //获取渠道的名字和id
    HttpConn("/id_name", {"table_id": 3}, f);
    function f(data) {
        if (data.cmd == 1) {
            var info = JSON.parse(data.info);
            arr = info.info
            //设置表的字段
            setTable(rows,arr)
            return
        }
        if (data.cmd == 10) {
            alert(data.info)
            return
        }
    }
    function setTable(rows,arr) {
        var rowsLen = rows.length;
        // var cellsLen = rows.item(0).cells.length;
        for (var i = 1; i < rowsLen; i++) {
            var value = rows[i].cells[3].innerHTML
            for (var j = 0; j < arr.length; j++) {
                if (value == arr[j][0]) {
                    rows[i].cells[3].innerHTML = arr[j][1]
                }
            }
            var money = Number(rows[i].cells[4].innerHTML)
            rows[i].cells[4].innerHTML = (money / 10).toFixed(2)
        }
    }


}