//"/upload?user=lkdfgl&name=kjdf&filekey=tfhhcgh" 
function HttpFileConnetion (fileUrl){
$(function(){
  //alert("hello");
 var options={
            url:fileUrl,
            type:"post",
            success:function(mes){
               selectUploadCmdPram(mes);
            }
        };
        $("#form2").submit(function() {
            $(this).ajaxSubmit(options);
                return false;
        });

});
}
//发送 post
function HttpConnetion(filepath){
var  session =	getCookie("session")
var path = "/post?path="+filepath+"&"+getCheckedStr()+"&session="+session
//alert(path);
var json = {name:"Donald Duck",city:"Duckburg"};
 $(document).ready(function(){
    $.post(path,json, function(data,status){
							  
	selectCmdPram(data)
	document.write(data);
     // alert("数据：" + data + "\n状态：" + status);
    });
});
}