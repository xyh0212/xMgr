<?php
require_once $_SERVER['DOCUMENT_ROOT'] . "/hotgame/kd/common/common.php";

abstract class db {
	protected $_conn; //单态连接

	protected $_error; //错误信息
	protected $_errno; //错误号
	protected $_perPageRecord; //每页显示几条数据
	protected $_currentTime;
	protected $_recordCount;
	public function __construct() {

	}
	//根据id查找一条记录
	public function findById($table, $id) {
		$sql = "select * from {$table} where id={$id}";
		return $this->uniqueQuery($sql);
	}
	// 连数据库
	abstract protected function _connect();

	//检查并执行查询
	private function _check_query($result, $sql) {
		if (!$result) {
			$this->_error = mysql_error();
			$this->_errno = mysql_errno();
			$this->_error("invalid SQL: " . $sql);
		}
	}
	//设置每页显示的数量(用于分页,当query的offset为空时,此值无用
	public function setPerPageRecord($perPageRecord) {
		$this->_perPageRecord = $perPageRecord;
	}
	//执行SQL并返回结果
	protected function _sendSQL($sql, $offset = false) {
		if (is_numeric($offset) && is_numeric($this->_perPageRecord)) {
			$sql = $sql . " limit {$offset}, " . $this->_perPageRecord;
		}
		$result = mysql_query($sql, $this->_conn);
		$this->_check_query($result, $sql);
		return $result;
	}
	//取得多条数据集
	public function query($sql, $offset = false) {
		$result = $this->_sendSQL($sql, $offset);
		$datas  = array();
		while ($row = mysql_fetch_array($result)) {
			$datas[] = $row;
		}
		return $datas;
	}
	//取得记录总数,假如不分页
	public function getRecordCount($sql) {
		$result                    = $this->_sendSQL($sql);
		return $this->_recordCount = mysql_num_rows($result);
	}
	//取得总页数
	public function getTotalPage() {
	}
	//取得唯一一条记录
	public function uniqueQuery($sql) {
		$result = $this->_sendSQL($sql);
		return mysql_fetch_array($result);
	}
	//取得多个值 (select单个字段时用)
	public function fetchValues($sql, $offset = false) {
		$result = $this->_sendSQL($sql, $offset);
		$datas  = array();
		while ($row = mysql_fetch_array($result)) {
			$datas[] = $row[0];
		}
		return $datas;
	}
	//错误处理
	private function _error($msg) {
		printf("</td></tr></table><b>Database error:</b> %s<br>\n", $msg);
		printf("<b>MySQL Error</b>: %s (%s)<br>\n", $this->_errno, $this->_error);
		die("Session halted.");
	}
	//取得某一值
	public function fetchValue($sql) {
		$result = $this->_sendSQL($sql);
		$value  = mysql_fetch_row($result);
		return $value[0];
	}
	//执行非查询语句返回id
	public function execute($sql) {
		$this->_sendSQL($sql);
		mysql_insert_id($this->_conn);
	}
	//执行非查询语句返回影响行数
	public function update($sql) {
		return $this->_sendSQL($sql);
	}
	//执行删除语句,要有返回值return,否则无法获得正确结果
	public function deletesql($sql) {
		return $this->_sendSQL($sql);
	}

	//关闭数据库
	public function close() {
		mysql_close($this->_conn);
	}
	//析构函数
	public function __destruct() {
		@$this->close();
	}
}

class account_db extends db {
	private $gsdb_host;
	private $gsdb_user;
	private $gsdb_pass;
	private $gsdb_name;
	private $gsdb_code;

	function __construct($idx) {
		global $aryAccountDB;
		$this->gsdb_host = $aryAccountDB[$idx]["ip"];
		$this->gsdb_user = $aryAccountDB[$idx]["user"];
		$this->gsdb_pass = $aryAccountDB[$idx]["psw"];
		$this->gsdb_name = $aryAccountDB[$idx]["db"];
		$this->gsdb_code = $aryAccountDB[$idx]["code"];
		$this->_connect();
	}

	protected function _connect() {
		if (!is_resource($this->_conn)) {
			$this->_conn = mysql_connect($this->gsdb_host, $this->gsdb_user, $this->gsdb_pass);
			mysql_select_db($this->gsdb_name, $this->_conn);
			db::execute("use " . $this->gsdb_name);
			mysql_query("set names " . $this->gsdb_code, $this->_conn);
		}
	}
}

class cz_db extends db {
	private $gsdb_host;
	private $gsdb_user;
	private $gsdb_pass;
	private $gsdb_name;
	private $gsdb_code;

	function __construct($idx) {
		global $aryCzDB;
		$this->gsdb_host = $aryAccountDB[$idx]["ip"];
		$this->gsdb_user = $aryAccountDB[$idx]["user"];
		$this->gsdb_pass = $aryAccountDB[$idx]["psw"];
		$this->gsdb_name = $aryAccountDB[$idx]["db"];
		$this->gsdb_code = $aryAccountDB[$idx]["code"];
		$this->_connect();
	}

	protected function _connect() {
		if (!is_resource($this->_conn)) {
			$this->_conn = mysql_connect($this->gsdb_host, $this->gsdb_user, $this->gsdb_pass);
			mysql_select_db($this->gsdb_name, $this->_conn);
			db::execute("use " . $this->gsdb_name);
			mysql_query("set names " . $this->gsdb_code, $this->_conn);
		}
	}
}

function GetDBCfg($idx) {
	global $aryGS;
	$cfg = $aryGS[$idx];
	if (!isset($cfg)) {
		$cfg = $aryGS[0];
	}
	$cfgDefault = $aryGS[0];

	$user = getval($cfg, 'user');
	if (!isset($user) || $user == '') {
		$user = $cfgDefault['user'];
	}

	$psw = getval($cfg, 'psw');
	if (!isset($psw) || $psw == '') {
		$psw = $cfgDefault['psw'];
	}

	$db = getval($cfg, 'db');
	if (!isset($db) || $db == '') {
		$db = "hotgame_hl$idx";
	}

	$code = getval($cfg, 'code');
	if (!isset($code) || $code == '') {
		$code = $cfgDefault['code'];
	}

	$ip = getval($cfg, 'ip');
	if (!isset($ip) || $ip == '') {
		$ip = "127.0.0.1:33123";
	}

	$cfg = array(
		"ip"   => $ip,
		"user" => $user,
		"psw"  => $psw,
		"db"   => $db,
		"code" => $code,
	);
	return $cfg;
}

class gsdb extends db {
	private $cfg;

	function __construct($idx) {
		$this->cfg = GetDBCfg($idx);
		// AddInfoLog('x', print_r($this->cfg,1));
		$this->_connect();
	}

	protected function _connect() {
		if (!is_resource($this->_conn)) {
			$cfg         = $this->cfg;
			$this->_conn = mysql_connect($cfg['ip'], $cfg['user'], $cfg['psw']);
			mysql_select_db($cfg['db'], $this->_conn);
			db::execute("use " . $cfg['db']);
			mysql_query("set names " . $cfg['code'], $this->_conn);
		}
	}
}
?>
