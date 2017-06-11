<?php
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/common/db.php";
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/sdk/sdk_common.php";
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/sdk/cb/login_common.php";
login();
function login()
{
    $ArrInfo = $_GET;
    AddInfoLog(sdk_type, getUrl());
    //检查彩果用户id
    if (!array_key_exists("uid", $ArrInfo)) {
        echo "彩果用户id不存在";
        AddErrLog(sdk_type, "彩果用户id不存在");
        return;
    }
    //检查彩果昵称
    if (!array_key_exists("uname", $ArrInfo)) {
        echo "彩果昵称不存在";
        AddErrLog(sdk_type, "彩果昵称不存在");
        return;
    }
    //检查签名
    if (!array_key_exists("sign", $ArrInfo)) {
        echo "签名不存在";
        AddErrLog(sdk_type, "签名不存在");
        return;
    }
    //组建要MD5的字符串
    $md5Str = "uid=".$_GET["uid"] . "uname=".$_GET["uname"] . key;

    // 验证签名
    if (md5($md5Str) != $_GET["sign"]) {
        echo "验证签名错误";
        AddErrLog(sdk_type, "验证签名错误");
        return;
    }
    $ret = ProcessCreateToken($ArrInfo["uname"], sdk_type);
    echo $ret;
    //成功登入
    AddInfoLog(sdk_type, "成功登入");



}


?>