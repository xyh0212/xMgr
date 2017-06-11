var ArrTr = new Array();
//var Gjson 
function sendQueryJson(){
	// {"cmd":1,"param":{"sql":"xxxxx"}}
	 var name= getCookie("name")
     var str = '{"cmd":6,"param":{"name":"'+name+'"}}';
     ws.send(str);
     
}

function selectCmdPram(pJson){
switch(pJson.cmd)
{
  case 5:
   sendQueryJson();
  break;
   case 6:
  Gjson = JSON.parse(pJson.json);
  for (var i = 0; i < Gjson.userjson.length; i++) {
        var tr = GetTr(Gjson.userjson[i]);
		ArrTr[i]=tr;
     }
  InitUI();
  break;
   case 7:
   alert("成功删除");
   location.reload() 
  break;
   case 8:
   alert("成功修改");
   location.reload() 
   case 9:  
   alert("成功添加"); 
   location.reload() 
   break;
  case 10:
  //查询sql的结果
  alert(pJson.cmd+pJson.info);
  break;
}
}


