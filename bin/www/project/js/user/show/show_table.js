
function selectCmdPram(pJson){
switch(pJson.cmd)
{
  case 3:
  //alert(pJson.info[0]) 
   var str =""; 
   for(var i=0;i<pJson.sql.length;i++){
    str += pJson.sql[i]+'</br>';
   }	  
  document.getElementById("resultdiv").innerHTML =str;
  break;
  case 5:
  //查询session
  ws.send(getQuerySqlCmdJson()); 
  break;
  case 10:
  //查询sql的结果
   //alert(pJson.info);
   var str = pJson.info  
   document.getElementById("resultdiv").innerHTML =str;
  
  break;
}
}

function getQuerySqlCmdJson(){
   var CheckedJsonStr= getUrlValue("CheckedJsonStr");
   var sql =getUrlValue("sqlJson");
    CheckedJsonStr = decodeURI(CheckedJsonStr)
    var str = '{"cmd":3,"param":{'+CheckedJsonStr+',"sql":"'+sql+'"}}';
    //alert(sql);
    return str;
}
