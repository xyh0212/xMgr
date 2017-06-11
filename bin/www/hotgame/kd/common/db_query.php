<?php

require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/config/sdktype.php";
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/common/common.php";

// d_gift.state
define('TB_GIFT_STATE_NO', 0); //未发送
define('TB_GIFT_STATE_GIVED', 1); //已发送

// CREATE TABLE `d_gift` (
//   `id` bigint(20) NOT NULL,
//   `account_id` bigint(20) DEFAULT '0',
//   `p` int(4) DEFAULT '0',
//   `gift_type` int(4) DEFAULT '0',
//   `state` int(4) DEFAULT '0',
//   PRIMARY KEY (`id`),
//   KEY `acc` (`account_id`,`p`,gift_type)
// ) ENGINE=MyISAM;

function GetTitleAndBodyByGift($p, $gift_type, &$title_id, &$body_id, &$package_id) {
	switch ($p) {
	case P_SG:{
			switch ($gift_type) {
			case GIFT_ATTENTION:
				$title_id   = 8000014;
				$body_id    = 8100014;
				$package_id = 20021;
				break;
			default:
				break;
			}
		}
	default:
		break;
	}
}

function GetGiftStatusDB($db, $uid, $p, $gift_type) {
	$sql = "SELECT state FROM d_gift WHERE account_id='$uid' and p='$p' and gift_type='$gift_type'";

	$rowOld = $db->uniqueQuery($sql);
	if ($rowOld == null) {
		return 0;
	}

	return $rowOld['state'];
}

function GetSvrName($db, $svrid) {
	$sql = "SELECT name FROM s_server WHERE id='$svrid'";

	$rowOld = $db->uniqueQuery($sql);
	if ($rowOld == null) {
		return 0;
	}

	return $rowOld['name'];
}

function InsertCmd_Reload($user_id, $db) {
	$sql = "INSERT INTO `cmd_d` (`type`, `user_id`,txt) VALUES ('1', '$user_id','')";
	$db->uniqueQuery($sql);
}

function AddItemLog($gift_type, $userid, $package_id) {
	if (!is_dir('log/')) {
		mkdir('log/', 0777);
	}
	// 使用最大权限0777创建文件
	$str = "$userid,$package_id,$gift_type";
	file_put_contents("log/$sdk_type-item" . ' ' . date("Y-m-d") . ".log", date("Y-m-d") . ' ' . date("H:i:s") . "," . GetIP() . ',' . $sdk_type . ',' . $str . "\n", FILE_APPEND | LOCK_EX);
}

// 0:发送失败 1:发送成功 2:已发送
function AwardGiftDB($as_db, $uid, $sdk_uname, $p, $gift_type, $svr, $sdk) {
	$sql = "INSERT ignore d_gift(account_id, p, gift_type) values ('$uid','$p','$gift_type')";
	$as_db->execute($sql);

	$sql    = "SELECT state FROM d_gift WHERE account_id='$uid' and p='$p' and gift_type='$gift_type'";
	$rowOld = $as_db->uniqueQuery($sql);
	if ($rowOld == null || $rowOld['state'] == null) {
		return get_retmsg(-102, '数据错误');
	}
	$state = $rowOld['state'];
	if ($state != TB_GIFT_STATE_NO) {
		return get_retmsg(2, '已发送');
	}

	$gs_db          = new gsdb($svr);
	$sql            = "SELECT id from d_user where account_id in (select id from d_account where name='$sdk_uname' AND sdk_type=$sdk AND svr_id=$svr)";
	$rowOldUserInfo = $gs_db->uniqueQuery($sql);
	// echo $sql;
	if ($rowOldUserInfo == null || $rowOldUserInfo['id'] == 0) {
		return get_retmsg(-106, '帐号不存在');
	}
	$idUser = $rowOldUserInfo['id'];

	$title_id   = 0;
	$body_id    = 0;
	$package_id = 0;
	GetTitleAndBodyByGift($p, $gift_type, $title_id, $body_id, $package_id);
	if ($package_id == 0) {
		return get_retmsg(-104, '活动不存在');
	}

	$sql = "INSERT INTO d_package_item (`user_id`, `itemtype_main`,itemtype_param, `num`, `title_id`, `desc_id`, package) VALUES ($idUser, 999999, $package_id, 1, '$title_id', '$body_id', 1)";
	$gs_db->uniqueQuery($sql);

	$sql = sprintf("UPDATE d_gift SET state=%d WHERE account_id='$uid' and p='$p' and gift_type='$gift_type'", TB_GIFT_STATE_GIVED);
	$as_db->uniqueQuery($sql);

	InsertCmd_Reload($idUser, $gs_db);
	AddItemLog($idUser, $package_id, $gift_type);

	return get_retmsg(1, '发送成功');
}

function CheckAccountTokenDB($db, $uid, $token, &$sdk_uname, &$sdk) {
	$sql    = "SELECT id,token,name,sdk_type FROM d_accounts WHERE id='$uid'";
	$rowOld = $db->uniqueQuery($sql);
	if ($rowOld == null or $rowOld['token'] == null) {
		return get_retmsg(-101, 'account not exist');
	}

	if ($rowOld['token'] != $token) {
		return get_retmsg(-103, 'token error');
	}

	$sdk       = $rowOld['sdk_type'];
	$sdk_uname = $rowOld['name'];

	return null;
}
