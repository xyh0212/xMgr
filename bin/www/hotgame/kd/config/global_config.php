<?php
session_start();
header("Expires: Sat, 1 Jan 2005 00:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . "GMT");
header("Cache-Control: no-cache, must-revalidate");
header("Pragma: no-cache");
header("Content-type:text/html;charset=utf-8");
// date_default_timezone_set('PRC'); //时区设置
error_reporting(E_ALL ^ E_NOTICE);
//error_reporting(0);

/**********网站路径配置**********/
$SITE_DOCUMENT = "/hotgame/kd"; //程序子目录
$SITE_ROOT     = $_SERVER['DOCUMENT_ROOT'] . $SITE_DOCUMENT;
$SITE_BASE     = $SITE_ROOT; //网站程序根目录
$CLASS_BASE    = $SITE_BASE . "/classes"; //类目录
$COMMON_BASE   = $SITE_BASE . "/common"; //公用目录根目录

// 帐服的数据库
static $aryAccountDB = array(
	1 => array(
		'ip'   => '127.0.0.1:33123', //数据库服务器地址
		'user' => 'root2', //数据库用户名
		'psw'  => '123456', //数据库密码
		'db'   => 'hotgame_as', //数据库名称
		'code' => 'UTF8', //数据库编码方式
	),
);

static $aryCzDB = array(
	1 => array(
		"ip"   => "127.0.0.1:33123",
		'user' => 'root2', //数据库用户名
		'psw'  => '123456', //数据库密码
		"db"   => "hotgame_cz",
		"code" => "UTF8",
	),
);
