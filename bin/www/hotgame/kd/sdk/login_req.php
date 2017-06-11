<?php
require_once "sdk_common.php";

// 帐号登录请求
// 127.0.0.1:8088../login_req.php?sdk=15&uname=1&token=xxxx&p=1
// p=1:口袋
// p=2:地城无双
// p=3:乱戳三国
// code=1时resp才有意义
//返回{"code":1, "msg":"xxxx", "resp":{"uid":"int64","token":"32字节",sdkuname":"xxxx","sdktype":"int"}}

function ProcessSDKLogin() {
	$token = $_GET['token'];
	switch ($sdkType) {
	case 0:
		require_once "pc/login.php";
		$s = new PCService;
		return $s->CheckToken($uname, $token, $sdkType);
	return 0;
}
$ret = ProcessSDKLogin();
//TODO: 正式环境中跨域访问可能需要做控制
header("Access-Control-Allow-Origin:*"); //h5 需要跨域访问

echo $ret;
?>