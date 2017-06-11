<?php
define("sdk_type",270);
define("key","abcd");



//获取url的路劲
function getUrl(){
	$url = 'http://' . $_SERVER['SERVER_NAME'] . ':' . $_SERVER["SERVER_PORT"] . $_SERVER["REQUEST_URI"];
	return $url;
}

?>