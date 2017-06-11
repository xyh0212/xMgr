
var GArrTr = new Array()//数组编辑框
var Colum //编辑框
var ArrFieldName //数组表的字段名
var ArrMaxLen =new Array()//数组表的字段名
////////////////////////////////////////////////////////////////////////////////////////////////
function selectCmdPram(pJson){
switch(pJson.cmd)
{
  case 15:
    json = JSON.parse(pJson.info); //再一次的解析json    
	var  Gjson = json.json
	Colum = Gjson[0].length;
	ArrFieldName = JSON.parse(pJson.field)
    GArrTr = GjsonToArrTr(Gjson);
    ArrMaxLen=getRowMaxNumChar(Gjson);
	//alert("1");
	initUI();
  break; 
  case 5:
  //发送QueryGSJson
  sendQueryGSJson();
   break;
  case 10:
 
  alert(pJson.info);
  break;
  case 12:
  sendQueryGSJson();
  initUI();
  break;
  case 13:
  sendQueryGSJson();
  initUI();
  break;
  case 14:
  sendQueryGSJson();
  initUI();
  break;
}
}

function  initUI(){
var from1=  document.getElementById("form1");
    from1.innerHTML=null
    from1.appendChild(createTable());
}

function createTable(){
	var table = createElement("table")
	table.setAttribute("border","1");
	table.setAttribute("align","center");
	table.style.color="blue";
	var css = 'font-family: verdana,arial,sans-seriffont-size:11pxcolor:#333333border-width: 1pxborder-color: #999999;border-collapse: collapse;background-color:#c3dde0 ;'
    table.style.cssText =css ;
    var tr =createElement("tr")
    	tr.align="center";
    var mycars=new Array("id","服务器名","区名","服务器组","gs_id","ip","port","数据库组","数据库名","开服时间","删除","添加");
    for (var i = 0; i <mycars.length; i++) {
       var td = createElement("td")
       	   td = createElementAndValue("td",mycars[i]) 
       	   tr.appendChild(td);
    }
    table.appendChild(tr);
	for (var i = 0;i<GArrTr.length;i++){
		table.appendChild(GArrTr[i]);
		var hander =table.rows[i+1].cells[9];
        table.rows[i+1].cells[9].innerHTML= splitStr(hander.innerHTML);       
		}
	for (var i = 0;i<GArrTr.length;i++){
	  if (GArrTr[i].value==","){
	   table.rows[i+1].cells[9].childNodes[0].value=splitStr(table.rows[i].cells[9].innerHTML) //splitStr(table.rows[i].cells[9].childNodes[0].value);
	   }
	}
	var ptr = GArrTr[GArrTr.length-1].childNodes;
	return table;
}

//判断 编辑框的Tr 是否存在 
function isTrEdit(){
   var len = GArrTr.length;
   for(var i =0;i<len;i++){
   	//当值为“，”时为edit的编辑框
    if (GArrTr[i].value==","){ 
    	GArrTr.splice(i, 1); 
     }
   }
    return len;
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
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//动态的绑定事件
document.onmousemove = function(e){
	var className = e.target.className;
                   if(className== "deleteImage"){
                       /*避免重复注册浪费性能*/

                       e.target.className = "";
                       e.target.onclick = function(){
						sendDeleteGSJson(parseInt(this.value))
						initUI();						   
                       }
                   }
				   if(className== "editImage"){
                       /*避免重复注册浪费性能*/

                       className= "";
                       e.target.onclick = function(){
                       	var len = isTrEdit();
      					for(var i =0;i<len;i++){
						if (this.value  == GArrTr[i].value ) {
							 var tr =  createEditTr(Colum,"保存",ArrMaxLen);
							for(var j = 0 ;j<tr.childNodes.length-1;j++){
							tr.childNodes[j].childNodes[0].value = GArrTr[i].childNodes[j].innerHTML;
                                //alert("sjsdhf");
							}
								  GArrTr.splice(i+1, 0,tr); 
								   initUI();
								   break;
								} 
					 	}

                       }

                   }
				    if(className== "button"){
                       /*避免重复注册浪费性能*/
                       className= "";
                       e.target.onclick = function(){	
                       if(this.value=="保存"){
                       sendGSJson(14);
                        
                       }
                       if(this.value=="添加"){
                        sendGSJson(12);
                        //initUI();
                       }
                       }
                   }
 }
 
function addOnClick(){
isTrEdit();
GArrTr.splice(GArrTr.length, 0,createEditTr(Colum,"添加",ArrMaxLen)); 
initUI();
} 
//向服务端发送数据
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//"cmd":14   Updata
//"cmd":12   Insert
//  var str = '{"cmd":12,"param":{"info":["id:56","name:dfg","name2:dlf","gs_group_id:14","gs_id:12","ip:6354613","gs_port:45","db_group_id:2","db_name:dfjh","start_time:2014-06-15 08:37:18"]}} '
function sendGSJson(num){
	var str ='{"cmd":'+num+',"param":{"info":['
	for(var i = 0 ;i<GArrTr.length;i++){
      if(GArrTr[i].value==","){
         var arrtr= GArrTr[i].childNodes;
         for(var j = 0 ;j<arrtr.length-1;j++){
              str +='"'+ ArrFieldName[j]+':'+trim(arrtr[j].childNodes[0].value)+'",';
         }
      }
    }
     str = str.substring(0,str.length-1);
     str += ']}}';
     ws.send(str);
}
function sendQueryGSJson(){
 var str = '{"cmd":15} '
  ws.send(str);
}
//var str = '{"cmd"13:,"param":{"id":"'+id+'"}}';
function sendDeleteGSJson(id){
  if(confirm("确定删除用户：？")){
	   var str = '{"cmd":13,"param":{"id":'+id+'}}';
       ws.send(str);	
	  }
}



