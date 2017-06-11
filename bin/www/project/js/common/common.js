//window.onload = function() {
   //alert("hello");
    //FileLoading.css("body.css")
    //FileLoading.js("common/constant.js")
   // FileLoading.js("common/auto_websoket.js")
//}

var   FileLoading = {
    css: function(path){
        if(!path || path.length === 0){
            throw new Error('argument "path" is required !');
        }
        var head = document.getElementsByTagName('head')[0];
        var link = document.createElement('link');
        link.href = path;
        link.rel = 'stylesheet';
        link.type = 'text/css';
        head.appendChild(link);
    },
    js: function(path){
        if(!path || path.length === 0){
            throw new Error('argument "path" is required !');
        }
        var head = document.getElementsByTagName('head')[0];
        var script = document.createElement('script');
        script.src = path;
        script.type = 'text/javascript';
        head.appendChild(script);
        //alert("suc")
    }
}




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
    //get_par =decodeURI(get_par);
    return get_par;
}
//去除左右空格
function trim(s){
    return s.replace(/(^\s*)|(\s*$)/g, "");
}