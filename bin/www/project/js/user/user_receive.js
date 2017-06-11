///////////////////////////////////////////////////////////////////////////////////////////////
//接收数据///////////////////////////////////////////////////////////////////////////////////////////////
function sendQueryGSJson(){
	// {"cmd":1,"param":{"sql":"xxxxx"}}
     // var str = '{"cmd":2}';
	  var str ='{"cmd":2,"param":{"tag":1}}'
      ws.send(str);
}

function selectCmdPram(pJson){
switch(pJson.cmd)
{
  case 2:
    json = JSON.parse(pJson.info); //再一次的解析json   
    initUI();
  break;
  case 3:
  break;
  case 5:
  //发送QueryGSJson
  sendQueryGSJson();
   break;
  case 10:
  //查询sql的结果
  alert(pJson.cmd);
  break;
}
}
///////////////////////////////////////////////////////////////////////////////////////////////
//界面UI///////////////////////////////////////////////////////////////////////////////////////////////
var checkboxs =new Array()
var Gdivs =new Array()	
var json	

function initUI(){
        // 创建 复选框 元素
        AppendDiv();
        //初始化下拉框
        var dat;
        for (var i = 0; i < json.num; i++) {
        dat = dat + "<option>" + json.p[i].name + "</option>";
        }
        document.getElementById("select1").innerHTML = dat;
}
//封装一个 div 子元素有
//<ckeckbox><a><br>
function GetDiv(gs_name,gs_id){
 //var from1=  document.getElementById("from1");
        var div = document.createElement("div");
        var checkbox = document.createElement("input");
        checkbox.type = "checkbox";
        checkbox.value = gs_id;
        var a = document.createElement("a")
        a.innerHTML =gs_name;
        var br = document.createElement("br")
        div.appendChild(checkbox);
        div.appendChild(a);
        div.appendChild(br);
        return div;
}

function AppendDiv(){
 var from1=  document.getElementById("from1");
    from1.innerHTML="";
    for (var i = 0; i < json.num; i++) {
        if (ProductValue == json.p[i].name) {
            for (var j = 0; j < json.p[i].gs_num; j++) {
                      var div = GetDiv(json.p[i].gs[j].name,json.p[i].gs[j].id);
            	      from1.appendChild(div);
            	      checkboxs[j] = div.firstChild;
            	      Gdivs[j]= div;	  
            	   }
            	  }
            }
}
///////////////////////////////////////////////////////////////////////////////////////////////
//展示结果///////////////////////////////////////////////////////////////////////////////////////////////
function showResult(gJSON){
   var gTextarea = document.getElementById("textarea").value;
   var str ="";
   for(var i = 0 ; i<gJSON.num;i++){
       if(ProductValue==gJSON.p[i].name){
       	   str='"gs_num":"'+gJSON.p[i].gs_num;
       	   	for(var j=0;j<gJSON.p[i].gs_num;j++){
	        str= str + '</br>"name":'+gJSON.p[i].gs[j].name+'</br>"id":'+gJSON.p[i].gs[j].name+'</br>';
	        str= str + '"gs_num":"';
  	        }
       }
    }
	inser(str)
}

function inser(str){
var ss = document.getElementById("textarea").value;
document.getElementById("resultdiv").innerHTML=	'Select languge:</br>'+ss+'</br>result:</br>'+str+'</br>textarea value:'
}