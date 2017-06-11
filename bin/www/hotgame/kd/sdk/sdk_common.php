<?php

require_once "sdk_config.php";
require_once "as.php";
require_once "sdk_common_test.php";

define("ERRJSON_ACCOUNT_CREATE_1403", '{"code":"1403","msg":"account create fail"}');
define("ERRJSON_SDKTYPE_1405", '{"code":"1405","msg":"sdk type not exist"}');
define("ERRJSON_SDKTYPE_1406", '{"code":"1406","msg":"sdk type not exist"}');
define("ERRJSON_NOTOKEN_1411", '{"code":"1411","msg":"token error"}');
define("ERRJSON_NOTOKEN_1412", '{"code":"1412","msg":"sdk return error"}');

// sdktype+bbbb

function gen_errmsg($code, $msg, $uid, $token) {
	// return sprintf('{"code":%d,"msg":"%s"}', $code, $msg);
	if ($code == 1) {
		$arrayName = array(
			'code' => $code,
			'msg'  => $msg,
			'resp' => array(
				'uid'   => $uid,
				'token' => $token),
		);
	} else {
		$arrayName = array(
			'code' => $code,
			'msg'  => $msg,
		);
	}

	return json_encode($arrayName);
}

function gen_errmsg2($code, $msg) {
	return sprintf('{"code":%d,"msg":"%s"}', $code, $msg);
}

function gen_login_suc($uid, $token, $last_svr, $sdkuid) {
	global $RECOMMEND_GSID;

	if ($last_svr != null) {
		$str = sprintf('{"code":1,"resp":{"uid":"%s","token":"%s", "gs":%s, "svr":%s, "sdkuid":"%s"}}', $uid, $token, $RECOMMEND_GSID, $last_svr, $sdkuid);
	} else {
		$str = sprintf('{"code":1,"resp":{"uid":"%s","token":"%s", "gs":%s, "sdkuid":"%s"}}', $uid, $token, $RECOMMEND_GSID, $sdkuid);
	}
	return $str;
	// $arrayName = array(
	//     'code' => 1,
	//     'resp' => array(
	//     'uid' => $uid,
	//     'token' => $token)
	//     );
	// return $arrayName;
}

function gen_login_suc_app($uid, $token) {
	global $RECOMMEND_GSID;
	$str = sprintf('{"code":1,"resp":{"uid":"%s","token":"%s", "gs":%s}}', $uid, $token, $APP_RECOMMEND_GSID);
	return $str;
	// $arrayName = array(
	//     'code' => 1,
	//     'resp' => array(
	//     'uid' => $uid,
	//     'token' => $token)
	//     );
	// return $arrayName;
}

function gen_gstoken_suc($uid, $sdkuname, $sdktype) {
	$str = sprintf('{"code":1,"resp":{"uid":"%s","sdkuname":"%s","sdktype":%d}}', $uid, $sdkuname, $sdktype);
	return $str;
}

// 32
function create_token() {
	global $AS_ID;
	$data  = $_SERVER['REMOTE_ADDR'] . time() . rand();
	$token = md5($data);
	$token = substr_replace($token, $AS_ID, 0, 1);
	return $token;
}

function ProcessCreateToken($uname, $sdk_type) {
	$my_token = create_token();
	return ProcessCreateTokenByToken($uname, $sdk_type, $my_token);
}

function GetLastSvr($id) {
	if (!array_key_exists("p", $_GET)) {
		return null;
	}
	$nProductType = $_GET['p'];

	$as_db  = new account_db(1);
	$sql    = sprintf('SELECT last_svr FROM d_info WHERE account_id=%d AND p=%d', $id, $nProductType);
	$rowOld = $as_db->uniqueQuery($sql);
	if ($rowOld == null) {
		return null;
	}

	$last_svr = $rowOld['last_svr'];
	return $last_svr;
}

function ProcessCreateTokenByToken($uname, $sdk_type, $my_token) {
	$as_db = new account_db(1);

	$sql = "INSERT ignore d_accounts(name, sdk_type, token, token_timestamp) values ('$uname',$sdk_type,'$my_token',UNIX_TIMESTAMP()) ON DUPLICATE KEY UPDATE token='$my_token',token_timestamp=UNIX_TIMESTAMP()";
	$as_db->execute($sql);

	$sql    = sprintf('SELECT id,token FROM d_accounts WHERE name=\'%s\' AND sdk_type=%d', $uname, $sdk_type);
	$rowOld = $as_db->uniqueQuery($sql);
	if ($rowOld == null or $rowOld['id'] == 0) {
		return ERRJSON_ACCOUNT_CREATE_1403;
	}

	$uid = $rowOld['id'];

	$last_svr = GetLastSvr($uid);

	return gen_login_suc($uid, $rowOld['token'], $last_svr, $uname);
}

function ProcessCreateTokenAPP($uname, $sdk_type) {
	$my_token = create_token();

	$db  = new account_db(1);
	$sql = "INSERT ignore d_accounts(name, sdk_type, token, token_timestamp) values ('$uname',$sdk_type,'$my_token',UNIX_TIMESTAMP()) ON DUPLICATE KEY UPDATE token='$my_token',token_timestamp=UNIX_TIMESTAMP()";
	$db->execute($sql);

	$sql    = sprintf('SELECT id,token FROM d_accounts WHERE name=\'%s\' AND sdk_type=%d', $uname, $sdk_type);
	$rowOld = $db->uniqueQuery($sql);
	if ($rowOld == null or $rowOld['id'] == 0) {
		return ERRJSON_ACCOUNT_CREATE_1403;
	}

	$uid = $rowOld['id'];

	return gen_login_suc_app($uid, $rowOld['token']);

}
