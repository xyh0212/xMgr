<?php
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/common/db.php";
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/sdk/sdk_common.php";
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/sdk/cb/login_common.php";

login();
function login()
{
    AddInfoLog(sdk_type, getUrl());

    //检查用户名
    if (!array_key_exists("token", $_GET)) {
        echo "用户名不存在";
        AddInfoLog(sdk_type, "用户名不存在");
        return;
    }
    //获取jsjon
    $url = "http://wx.9g.com/open/userinfo?token=" . $_GET['token'];
    $ArrInfo = json_decode(file_get_contents($url), true);
    //检验
    if (array_key_exists("errcode", $ArrInfo)) {
        echo "检验出错";
        AddInfoLog(sdk_type, "检验出错");
        return;
    }

    $ret = ProcessCreateToken($ArrInfo["username"], sdk_type);
    echo $ret;
    //成功登入
    AddInfoLog(sdk_type, "成功登入");

}

?>