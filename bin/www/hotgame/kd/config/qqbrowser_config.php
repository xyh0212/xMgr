<?php
require_once "global_config.php";

// 各个游服对应的数据库权限
static $aryGS = array();
//模板
$aryGS[0] = array(
	"ip"   => "127.0.0.1:33123",
	"user" => "hotgame_web",
	"psw"  => "hotgame82yearlf.HL",
	"db"   => "hotgame_hl",
	"code" => "UTF8",
);

$aryGS[2001] = array(
	"ip"     => "119.29.75.178:33123",
	'online' => 1,
);
