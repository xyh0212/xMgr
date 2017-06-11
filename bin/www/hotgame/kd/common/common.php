<?php

function encode($str) {
	$str = str_replace("<", "&lt;", $str);
	$str = str_replace(">", "&gt;", $str);
	$str = str_replace("'", "''", $str);
	$str = str_replace('"', '&quot;', $str);
	$str = str_replace(" ", "&nbsp;", $str);
	$str = str_replace("\n", "<brs>", $str);
	return $str;
}

function decode($str) {
	$str = str_replace("&lt;", "<", $str);
	$str = str_replace("&gt;", ">", $str);
	$str = str_replace("''", "'", $str);
	//$str = str_replace('&quot;','"',$str);
	$str = str_replace("&nbsp;", " ", $str);
	$str = str_replace("<brs>", "\n", $str);
	return $str;
}

//escape编码， 解决中文传输
function escape($str) {
	preg_match_all("/[\xc2-\xdf][\x80-\xbf]+|[\xe0-\xef][\x80-\xbf]{2}|[\xf0-\xff][\x80-\xbf]{3}|[\x01-\x7f]+/e", $str, $r);
	//匹配utf-8字符，
	$str = $r[0];
	$l = count($str);
	for ($i = 0; $i < $l; $i++) {
		$value = ord($str[$i][0]);
		if ($value < 223) {
			$str[$i] = rawurlencode(utf8_decode($str[$i]));
			//先将utf8编码转换为ISO-8859-1编码的单字节字符，urlencode单字节字符.
			//utf8_decode()的作用相当于iconv("UTF-8","CP1252",$v)。
		} else {
			$str[$i] = "%u" . strtoupper(bin2hex(iconv("UTF-8", "UCS-2", $str[$i])));
		}
	}
	return join("", $str);
}

//unescape解码， 解决中文传输
function unescape($str) {
	$r = '';
	$str = rawurldecode($str);
	preg_match_all("/(?:%u.{4})|[^%]+/", $str, $r);
	$ar = $r[0];
	foreach ($ar as $k => $v) {
		if (substr($v, 0, 2) == "%u" && strlen($v) == 6) {
			$ar[$k] = iconv("UCS-2", "UTF-8", @pack("H4", substr($v, -4)));
		}

	}
	return join("", $ar);
}

function toLower($str) {
	$len = strlen($str);
	for ($i = 0; $i < $len; $i++) {
		if (ord($str{$i}) >= 65 && ord($str{$i}) <= 90) {
			$str{$i} = chr(ord($str{$i}) + 32);
		}
	}
	return $str;
}

function inject_check($str) {
	$tmp = eregi('select|insert|update|delete|\'|\/\*|\*|\.\.\/|\.\/|union|into|load_file|outfile', $str); // 进行过滤
	if ($tmp) {
		echo "<script>";
		echo "alert('非法访问！');";
		echo "parent.location.href='" . UI_BASE . "/logout.php';";
		echo "</script>";
		die();
	} else {
		return $str;
	}
}

function checkLogin() {
	if (!isset($_SESSION["user"])) {
		echo "<script>";
		echo "alert('登录超时！');";
		echo "parent.location.href='" . ACTION_BASE . "/logout.act.php';";
		echo "</script>";
		die();
	}
}

function GetIP() {
	if (!empty($_SERVER["HTTP_CLIENT_IP"])) {
		$cip = $_SERVER["HTTP_CLIENT_IP"];
	} else if (!empty($_SERVER["HTTP_X_FORWARDED_FOR"])) {
		$cip = $_SERVER["HTTP_X_FORWARDED_FOR"];
	} else if (!empty($_SERVER["REMOTE_ADDR"])) {
		$cip = $_SERVER["REMOTE_ADDR"];
	} else {
		$cip = "noip";
	}

	return $cip;
}

function FastPostRequest($url, $aryData, $optional_headers = null) {
	$data = http_build_query($aryData);
	$params = array('http' => array(
		'method' => 'POST',
		'content' => $data,
	));
	if ($optional_headers !== null) {
		$params['http']['header'] = $optional_headers;
	}
	$ctx = stream_context_create($params);
	$fp = @fopen($url, 'rb', false, $ctx);
	if (!$fp) {
		throw new Exception("Problem with $url, $php_errormsg");
	}
	$response = @stream_get_contents($fp);
	if ($response === false) {
		throw new Exception("Problem reading data from $url, $php_errormsg");
	}
	return $response;
}

function AddPayLog($sdk_type, $str) {
	if (!is_dir('paylog')) {
		mkdir('paylog', 0777);
	}
	// 使用最大权限0777创建文件
	file_put_contents("paylog/$sdk_type.log", date("Y-m-d") . ' ' . date("H:i:s") . "\t" . GetIP() . ' ' . $sdk_type . ' ' . $str . "\n", FILE_APPEND | LOCK_EX);
}

function AddPayErrLog($sdk_type, $str) {
	$date = date("Y-m-d");
	if (!is_dir('paylog')) {
		mkdir('paylog', 0777);
	}
	// 使用最大权限0777创建文件
	file_put_contents("paylog/$sdk_type-err-$date.log", date("Y-m-d") . ' ' . date("H:i:s") . "\t" . GetIP() . ' ' . $sdk_type . ' ' . $str . "\n", FILE_APPEND | LOCK_EX);
}

function AddErrLog2($sdk_type, $str) {
	if (!is_dir('log/')) {
		mkdir('log/', 0777);
	}
	// 使用最大权限0777创建文件
	file_put_contents("log/$sdk_type-err.log", date("Y-m-d") . ' ' . date("H:i:s") . "\t" . GetIP() . ' ' . $sdk_type . ' ' . $str . "\n", FILE_APPEND | LOCK_EX);
}

function AddInfoLog2($sdk_type, $str) {
	if (!is_dir('log/')) {
		mkdir('log/', 0777);
	}
	// 使用最大权限0777创建文件
	file_put_contents("log/$sdk_type-info.log", date("Y-m-d") . ' ' . date("H:i:s") . "\t" . GetIP() . ' ' . $sdk_type . ' ' . $str . "\n", FILE_APPEND | LOCK_EX);
}

function getval($ary, $str) {
	$val = !empty($ary[$str]) ? $ary[$str] : null;
	return $val;
}

function AddErrLog($sdk_type, $str) {
	$date = date("Y-m-d");
	if (!is_dir('log/')) {
		mkdir('log/', 0777);
	}
	// 使用最大权限0777创建文件
	file_put_contents("log/$sdk_type-err-$date.log", ' ' . $date . " " . date("H:i:s") . "\t" . GetIP() . ' ' . $sdk_type . ' ' . $str . "\n", FILE_APPEND | LOCK_EX);
}

function AddInfoLog($sdk_type, $str) {
	$date = date("Y-m-d");
	if (!is_dir('log/')) {
		mkdir('log/', 0777);
	}
	// 使用最大权限0777创建文件
	file_put_contents("log/$sdk_type-info-$date.log", ' ' . $date . " " . date("H:i:s") . "\t" . GetIP() . ' ' . $sdk_type . ' ' . $str . "\n", FILE_APPEND | LOCK_EX);
}

function http_post($url, $data) {
	$ch = curl_init();
	curl_setopt($ch, CURLOPT_URL, $url);
	curl_setopt($ch, CURLOPT_HEADER, 0);
	curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
	curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
	curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
	curl_setopt($ch, CURLOPT_POST, 1);
	curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
	curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-type:application/json;charset=utf-8'));
	$rtn = curl_exec($ch);
	if ($errno = curl_errno($ch)) {
		throw new Exception(curl_error($ch), $errno);
	}
	curl_close($ch);
	return $rtn;
}
// $strPost = sprintf("gamekey=%s&timestamp=%s&cp=%s&token=%s&_sign=%s", $sdk_info['appkey'], $t, $sdk_info['cp'], $token, $strSign2);

function http_post2($url, $data) {
	$ch = curl_init();
	curl_setopt($ch, CURLOPT_URL, $url);
	curl_setopt($ch, CURLOPT_HEADER, 0);
	curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
	curl_setopt($ch, CURLOPT_TIMEOUT, 30);
	curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 30);
	curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
	curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
	curl_setopt($ch, CURLOPT_POST, 1);
	curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
	// curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-type:application/json;charset=utf-8'));
	$rtn = curl_exec($ch);

	if ($errno = curl_errno($ch)) {
		throw new Exception(curl_error($ch), $errno);
	}
	curl_close($ch);
	return $rtn;
}

function get_retmsg($code, $msg) {
	// {"code”:1,”result”: 充值成功}
	$arrayName = array(
		'code' => $code,
		'result' => urlencode($msg),
	);
	return json_encode($arrayName);
}
