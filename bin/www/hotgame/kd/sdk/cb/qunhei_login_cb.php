<?php
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/common/db.php";
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/sdk/sdk_common.php";
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/sdk/cb/login_common.php";

login();
function login()
{
	//打印log
	AddInfoLog(sdk_type, getUrl());
	$ArrInfo = $_GET;

	//检查用户名
	if (!array_key_exists("username", $ArrInfo)) {
		echo "用户名不存在";
		AddErrLog(sdk_type, "用户名不存在");
		return;
	}
    //检查md5签名
	if (!array_key_exists("flag", $ArrInfo)) {
		echo "检查md5签名不存在";
		AddErrLog(sdk_type, "检查md5签名不存在");
		return;
	}
    //检验平台服务器的时间戳
	if (!array_key_exists("time", $ArrInfo)) {
		echo "平台服务器的时间戳不存在";
		AddErrLog(sdk_type, "平台服务器的时间戳不存在");
		return;
	}
    //检验是否成年
	if (!array_key_exists("isadult", $ArrInfo)) {
		$ArrInfo["isadult"] = -1;
	}
    //检查serverid
	if (!array_key_exists("serverid", $ArrInfo)) {
		$ArrInfo["serverid"] = 1;
	}
    //md5(username+serverid+isadult+time+key)
	$md5Str = $_GET["username"] . $_GET["serverid"] . $_GET["isadult"] . $_GET["time"] . key;
    // 验证签名
	if (md5($md5Str) != $_GET["flag"]) {
		echo "验证签名错误";
		AddErrLog(sdk_type, "验证签名错误");
		return;
	}

	$ret = ProcessCreateToken($ArrInfo["username"], sdk_type);
	echo $ret;
    //成功登入
	AddInfoLog(sdk_type, "成功登入");


}








?>