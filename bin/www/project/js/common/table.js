//["",""]["",""]  to  [tr,tr]
//创建table的tr数组
//二维数组的json转换为数组tr
function GjsonToArrTr(Gjson){
	ArrTr = new Array()
	pArrTd= new Array()
    for (var i = 0; i <Gjson.length; i++) {
		var tr =createElement("tr")
            tr.align="center";
        for (var j = 0; j < Gjson[i].length+2; j++) {
			  var td = createElement("td")
			  //td = createElementAndValue("td",Gjson[i][j])  	  
		      if (j==Gjson[i].length){
				       var  img = getImg("img/delete.png","deleteImage");
					   img.value= Gjson[i][0];
					   //img=changeImg(img,20);
                       td.appendChild(img);
				  }
			  else if(j==Gjson[i].length+1){
				       var  img = getImg("img/edit.png","editImage");
					   img.value= Gjson[i][0];
					   //img=changeImg(img,20);
                       td.appendChild(img);
				  }
			  else{
		              td = createElementAndValue("td",Gjson[i][j])  
		              //td.size = Gjson[i][j].length;
			  }
			  tr.appendChild(td);     
         }
	 	 	 
		 tr.value=Gjson[i][0]
	     ArrTr[i]= tr	     
    }
	return ArrTr;
}
//var arrLen = new Array() //

//创建编辑框的Tr   value为button的值
//num 表列数；value 点击按钮的值；arrlen 编辑框的字符个数的数组
function createEditTr(num,value,arrLen){	
    var tr =createElement("tr");
    tr.value= ",";
	for(var i = 0 ;i<arrLen.length;i++){
	 var td =createElement("td");
	 var edit = createElement("input");
	 edit.size =arrLen[i];
     edit.type="text";
	 td.appendChild(edit);
	 tr.appendChild(td);
	 }
	 //var td = createElementAndValue("td",va) 
	 var td =createElement("td");
	 var edit = createElementAndValue("input",value) 
	 edit.type="button"	 
	 edit.className  = "button";
	 
	 td.appendChild(edit);
	 tr.appendChild(td);
	 tr.childNodes[0].childNodes[0].disabled = true;
  return tr;
}

//创建元素并且赋值
function createElementAndValue(name,value){
	 var td = document.createElement(name)
     td.innerHTML =  value;
     td.value =  value;
	 return td	
}
//创建一个元素
function createElement(name){
	 var obj = document.createElement(name)
	 return obj	
}

	


