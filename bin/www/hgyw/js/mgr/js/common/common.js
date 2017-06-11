
//分离字符串
function splitStr(str){
	var arr = new Array();
    arr = str.split("+");
    return arr[0];
}


function getCookie(name)
{
    var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");

    if(arr=document.cookie.match(reg))

        return unescape(arr[2]);
    else
        return null;
}

function getImg(src,className){
	var image = createElement("img")
        image.src =src;
        image=changeImg(image,20);
		image.className  = className;
		return image
}
function changeImg(objImg ,len)
 {   
     var most = len;        //设置最大宽度
     if(objImg.width > most)
     {
         var scaling = 1-(objImg.width-most)/objImg.width;    
         //计算缩小比例
         objImg.width = objImg.width*scaling;
         //objImg.height = objImg.height;            //img元素没有设置高度时将自动等比例缩小
         objImg.height = objImg.height*scaling;    //img元素设置高度时需进行等比例缩小
         //alert(objImg.height);
     }
   return objImg;
 }
 function getUrlValue(key){
 	 var par =key
    //获取当前URL
    var local_url = document.location.href; 
    //获取要取得的get参数位置
    var get = local_url.indexOf(par +"=");
    if(get == -1){
        return false;   
    }   
    //截取字符串
    var get_par = local_url.slice(par.length + get + 1);    
    //判断截取后的字符串是否还有其他get参数
    var nextPar = get_par.indexOf("&");
    if(nextPar != -1){
        get_par = get_par.slice(0, nextPar);
    }
    get_par =decodeURI(get_par);
    return get_par;
}
//去除左右空格
function trim(s){
    return s.replace(/(^\s*)|(\s*$)/g, "");
}
//格式化时间
function formatDate(nS) {
    var now=new Date(parseInt(nS) * 1000);
    var year=now.getFullYear();
    var month=now.getMonth()+1;
    var date=now.getDate();
    var hour=now.getHours();
    var minute=now.getMinutes();
    var second=now.getSeconds();
    return year+"-"+month+"-"+date+" "+hour+":"+minute+":"+second;
}
//格式化时间 1602200900
function formatTime(nS) {
    var time = nS.toString()
    var len = time.length
    var year="20"+time.substring(0,len-8)
    var month=time.substring(2,len-6)
    var date=time.substring(4,len-4)
    var hour=time.substring(6,len-2)
    var minute=time.substring(8,len-0)
    return year+"-"+month+"-"+date+" "+hour+":"+minute;
}
////id:时间控件的值 //
//return:时间戳
function getStartDateValue(id){
    var value =document.getElementById(id).value
    if (value==""){
        return
    }
    value += ' 00:00:00';
    var timestamp = new Date(value).getTime()
    return timestamp/1000
}

////id:时间控件的值 //
//return:时间戳
function getEndDateValue(id){
    var value =document.getElementById(id).value
    if (value==""){
        return
    }
    value += ' 24:00:00';
    var timestamp = new Date(value).getTime()
    return timestamp/1000
}
function getStartTadayDate(){
    var d = new Date();
    var value = '2015-01-01' //d.getFullYear()+"-"+(d.getMonth()+1)+"-"+d.getDate();
    value += ' 00:00:00';
    var timestamp = new Date(value).getTime()
    return timestamp/1000
}

function getEndTadayDate(){
    var d = new Date();
    var value =d.getFullYear()+"-"+(d.getMonth()+1)+"-"+d.getDate();
    value += ' 24:00:00';
    var timestamp = new Date(value).getTime()
    return timestamp/1000
}