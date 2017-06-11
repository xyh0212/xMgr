// JavaScript template using Windows Script Host
//var gdivs 
//var arrTr = new Array();
window.onload = function() {
    //alert(hex_md5("1"));	
	NewWebSocket();	
	document.onmousemove = function(e){
	var className = e.target.className;
         if(className== "deleteImage"){
            /*避免重复注册浪费性能*/
             className = "";
             e.target.onclick = function(){
			 sendDeleteUserJson(this.value);
             }
         }
		 if(className== "editImage"){
            className = "";
            e.target.onclick = function(){
               var len = ArrTr.length
               if (!haveTrEdit()){
                  len = ArrTr.length-1
               }
               	for(var i =0;i<len;i++){
						if (this.value  == ArrTr[i].value ) {
							 var tr =  createEditTr(5,"保存",[3,3,3]);
							for(var j = 0 ;j<tr.childNodes.length-1;j++){
							tr.childNodes[j].childNodes[0].value = ArrTr[i].childNodes[j].innerHTML;
							}
								  ArrTr.splice(i+1, 0,tr); 
								   InitUI();
								   break;
								} 
					 	}	
            }
         }
         if(className== "button"){
            className= "";
            e.target.onclick = function(){	
              if(this.value=="保存"){
              	 sendUserJson(8);
                 //sendGSJson(14);
              }
              if(this.value=="添加"){
              	  sendUserJson(9)
              	 // alert(this.value);
              //sendGSJson(12);
              }
            }
          }                                     
   }
   
}

function GetTr(gs_name){
        var tr = document.createElement("tr");
        tr.value = gs_name;
        var td5 = document.createElement("td")
        //td1.innerHTML =  gs_name;
        var td1 = document.createElement("td")
        td1.innerHTML =  gs_name;
        var td4 = document.createElement("td")
        	td4.innerHTML =  "******";
		var image2 = document.createElement("img")
        image2.src ="img/delete.png";
		image2.className  = "deleteImage";
		image2.value  = gs_name;
		image2 = changeImg(image2,IMAGE_LENGTH);
		var image3 = document.createElement("img")
		image3.className  = "editImage"
        image3.src ="img/edit.png";
		image3.value  = gs_name
		image3 = changeImg(image3,IMAGE_LENGTH);
		tr.appendChild(td5);
        tr.appendChild(td1);
         tr.appendChild(td4);
		var td2 = document.createElement("td")
		td2.appendChild(image2);
		var td3 = document.createElement("td")
		td3.appendChild(image3);
		
		tr.appendChild(td2);
		tr.appendChild(td3);
        return tr;
}
function addClick(){
	haveTrEdit();
//	var Arrlen = [3,3,3];
ArrTr.splice(ArrTr.length, 0,createEditTr(5,"添加",[3,3,3])); 
InitUI();	
//window.open("add_user.html");

}

function GetTrTitle(){
        var tr = document.createElement("tr");
        var td5 = document.createElement("td")
        td5.innerHTML =  "序号";
        var td1 = document.createElement("td")
        td1.innerHTML =  "用户名";
        var td4 = document.createElement("td")
        td4.innerHTML =  "密码";
		var td2 = document.createElement("td")
        td2.innerHTML =  "删除";
		var td3 = document.createElement("td")
        td3.innerHTML =  "修改";
        tr.appendChild(td5);
		tr.appendChild(td1);
		tr.appendChild(td4);
		tr.appendChild(td2);
		tr.appendChild(td3);
        return tr;
}
function InitUI(){
 var from1=  document.getElementById("form1");
    from1.innerHTML="";
	 var table = document.createElement("table")
	 table.setAttribute("border","1");
	 table.setAttribute("align","center");
	 var css = 'font-family: verdana,arial,sans-seriffont-size:11pxcolor:#333333border-width: 1pxborder-color: #999999;border-collapse: collapse;background-color:#c3dde0 ;'
    table.style.cssText =css ;
	 table.appendChild(GetTrTitle());
	 for (var i = 0; i < ArrTr.length; i++) {
	     table.appendChild(ArrTr[i]);
	 }
	 from1.appendChild(table);
}

function sendDeleteUserJson(name){
  if(confirm("确定删除用户："+name+"？")){
	   var str = '{"cmd":7,"param":{"name":"'+name+'"}}';
       ws.send(str);	
	  }
 
}

function sendUserJson(num){
	var gPassword 
    var gUser
	for(var i =0;i<ArrTr.length;i++){
	    if (ArrTr[i].value==","){
	      gUser =ArrTr[i].childNodes[1].childNodes[0].value;
	      gPassword =ArrTr[i].childNodes[2].childNodes[0].value;
	    }
	 }
	 alert(gUser);
	 gPassword = hex_md5(gPassword);
	 var str = '{"cmd":'+num+',"param":{"name":"' + gUser + '","pass":"' + gPassword + '"}}';
     ws.send(str);
}



function sendAddUserJson(name){
	var password = prompt("请输入修改的密码：");

	if (password == "".toString()){
		return ;
		}
	if(confirm("确定更改"+name+"的密码吗？")){
		password = hex_md5(password);
	   var str = '{"cmd":8,"param":{"name":"'+name+'","password":"'+password+'"}}';
       ws.send(str);	
	  } 
}



//判断 编辑框的Tr 是否存在 
//如果存在，那就删除,并返回true
//如果不存在，那就返回false
function haveTrEdit(){
   var len = ArrTr.length;
   for(var i =0;i<len;i++){
   	//当值为“，”时为edit的编辑框
    if (ArrTr[i].value==","){ 
    	ArrTr.splice(i, 1); 
    	return true
     }
   }
   return false
}