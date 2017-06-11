<?php

// p=1:口袋安卓
// p=2:地城无双
// p=3:乱戳三国
define('P_KD_ANDROID', 1);
define('P_DC', 2);
define('P_SG', 3);

// gift=1:关注
define('GIFT_ATTENTION', 1);

//sdk type
define('XSDKTYPE_WIN', 0);
define('XSDKTYPE_91', 3); //越狱-91
define('XSDKTYPE_NDUO', 5); //
define('XSDKTYPE_KY', 4); //越狱-快用
define('XSDKTYPE_360', 7); //安卓-360 服务端自己取uid https
define('XSDKTYPE_UC', 8); //安卓-uc 服务端自己取uid json
define('XSDKTYPE_BAIDU', 9); //安卓-百度多酷
define('XSDKTYPE_XIAOMI', 10); //安卓-小米
define('XSDKTYPE_DANGLE', 11); //安卓-当乐 服务端自己取uid get
define('XSDKTYPE_FORGAME', 12); //安卓-forgame post json md5md5
define('XSDKTYPE_WANDOUJIA', 13); //安卓-豌豆荚
define('XSDKTYPE_ANZHI', 14); //安卓-安智 post json base64 服务端自己取uid
define('XSDKTYPE_MUZHIWAN', 15); //安卓-拇指玩
define('XSDKTYPE_OPPO', 16); //安卓-oppo
define('XSDKTYPE_LIANXIANG', 17); //安卓-联想
define('XSDKTYPE_vivo', 18); //安卓-vivo
define('XSDKTYPE_HUAWEI', 19); //安卓-华为
define('XSDKTYPE_4399', 20); //安卓-4399
define('XSDKTYPE_8849', 21); //安卓-8849
define('XSDKTYPE_YINGYONGHUI', 22); //安卓-应用汇
define('XSDKTYPE_PPS', 23); //安卓-pps
define('XSDKTYPE_KUPAI', 24); //安卓-酷派
define('XSDKTYPE_JINLI', 25); //安卓-金立
define('XSDKTYPE_YOUKU', 26); //安卓-优酷

define('XSDKTYPE_ITOOLS', 27); //越狱-itools 服务端自己取uid get
define('XSDKTYPE_TONGBUTUI', 28); //越狱-同步 服务端自己取uid get
define('XSDKTYPE_XY', 29); //越狱-xy助手 post 返回json 客户端提供uid
define('XSDKTYPE_HAIMA', 30); //越狱-海马 服务端通过海马AccountName生成uid
define('XSDKTYPE_AISI', 31); //越狱-爱思助手 get 返回json 服务端自己取uid
define('XSDKTYPE_GUOPAN', 32); //越狱-果盘SDK get 返回true/false 客户端提供uid
define('XSDKTYPE_FORGAME_APPSTORE', 33); //正版-forgame
define('XSDKTYPE_FORGAME_ISO', 34); //越狱-forgame

//31 爱思
define('XSDKTYPE_GUOPAN_ANDOIRD', 35); //安卓-果盘SDK/xx助手 get 返回true/false 客户端提供uid
define('XSDKTYPE_SINA', 36); //安卓-新浪
define('XSDKTYPE_SOGOU', 37); //安卓-搜狗
define('XSDKTYPE_3G', 38); //安卓-3G
define('XSDKTYPE_37WAN', 39); //安卓-37玩
define('XSDKTYPE_37WAN_2', 40); //安卓-37玩2
define('XSDKTYPE_37WAN_3', 41); //安卓-37玩3
define('XSDKTYPE_3G_2', 42); //安卓-3G_2
define('XSDKTYPE_TYD', 53); //安卓-天奕达
define('XSDKTYPE_TYYY', 54); //安卓-天宇应用
define('XSDKTYPE_HXYX', 55); //安卓-航讯游戏
define('XSDKTYPE_YIGUO', 56); //安卓-艺果
define('XSDKTYPE_KAIFU', 57); //安卓-开服助手
define('XSDKTYPE_DINGKAI', 58); //安卓-鼎开
define('XSDKTYPE_KAOPU', 59); //安卓-靠谱助手
define('XSDKTYPE_SAIZONG', 60); //安卓-赛众
define('XSDKTYPE_HTC', 61); //安卓-HTC
define('XSDKTYPE_QXZ1', 62); //安卓-七匣子1
define('XSDKTYPE_QXZ2', 63); //安卓-七匣子2
define('XSDKTYPE_QXZ3', 64); //安卓-七匣子3
define('XSDKTYPE_QXZ4', 65); //安卓-七匣子4
define('XSDKTYPE_QXZ5', 66); //安卓-七匣子5
define('XSDKTYPE_QXZ6', 67); //安卓-七匣子6
define('XSDKTYPE_QXZ7', 68); //安卓-七匣子7
define('XSDKTYPE_YOULONG', 69); //安卓-游龙
define('XSDKTYPE_YOUYI', 70); //安卓-优亿市场
define('XSDKTYPE_ANFENG', 71); //安卓-安锋网
define('XSDKTYPE_TT1', 72); //安卓-TT游戏1
define('XSDKTYPE_TT2', 73); //安卓-TT游戏2
define('XSDKTYPE_TT3', 74); //安卓-TT游戏3
define('XSDKTYPE_TT4', 75); //安卓-TT游戏4
define('XSDKTYPE_KSW', 76); //安卓-看书网
define('XSDKTYPE_YMOW1', 77); //安卓-有米偶玩1
define('XSDKTYPE_YMOW2', 78); //安卓-有米偶玩2
define('XSDKTYPE_YMOW3', 79); //安卓-有米偶玩3
define('XSDKTYPE_SM', 89); //安卓-手盟
define('XSDKTYPE_KSW2', 90); //安卓-看书网2
define('XSDKTYPE_KSW3', 91); //安卓-看书网3
define('XSDKTYPE_BDTB', 92); //安卓-百度贴吧
define('XSDKTYPE_WXRD', 93); //安卓-无线热点
define('XSDKTYPE_SOGOULLQ', 94); //安卓-sogou浏览器
define('XSDKTYPE_49YOU', 95); //安卓-49游
define('XSDKTYPE_YYWAN', 96); //安卓-丫丫玩
define('XSDKTYPE_KUGOU', 97); //安卓-酷狗
define('XSDKTYPE_9388', 98); //正版-tw 9388
define('XSDKTYPE_IAPPLE', 99); //越狱-IApple
define('XSDKTYPE_9388_ANDROID', 100); //安卓-tw 9388
define('XSDKTYPE_9388_ANDROID_HK', 101); //安卓-香港 9388
define('XSDKTYPE_9388_MUMAYI', 101); //安卓-木蚂蚁
define('XSDKTYPE_9388_UUC', 102); //安卓-UU村
define('XSDKTYPE_9388_SYKF', 103); //安卓-手游咖啡（星云）
define('XSDKTYPE_EGRET', 104); //Egret
define('XSDKTYPE_GOSU', 105); //越南GOSU
define('XSDKTYPE_IPLAY99', 107); //9388子渠道:iPlay99 爱情公寓
define('XSDKTYPE_58PLAY', 108); //9388子渠道:58Play 赤霄
define('XSDKTYPE_58PLAY2', 109); //9388子渠道:58Play 赤霄2
define('XSDKTYPE_HAOMENG', 110); //好盟安卓
define('XSDKTYPE_58PLAYALL', 111); //9388子渠道:58Play 赤霄总渠道（包含赤霄108 109)
define('XSDKTYPE_HULI', 112); //越狱-狐狸助手
define('XSDKTYPE_CAIGUO', 113); //h5-彩果
define('XSDKTYPE_LUOBOWAN', 114); //h5-萝卜玩
define('XSDKTYPE_EGRETTW', 115); //Egret tw
define('XSDKTYPE_EGRET_WANBA', 116); //Egret 玩吧
define('XSDKTYPE_EGRET_QQ', 117); //Egret QQ
define('XSDKTYPE_H5MEITU', 118); //h5-美图
define('XSDKTYPE_EGRET_QQBROWSER', 119);//乱戳三国-QQ浏览器
define('XSDKTYPE_LCSG_QIDIAN', 120);//乱戳三国-起点
define('XSDKTYPE_LCSG_QQREADER', 121);//乱戳三国-QQ阅读

define('XSDKTYPE_CMMM', 203); //
define('XSDKTYPE_FEILIU', 204); //
define('XSDKTYPE_COOLPAD', 205); //
define('XSDKTYPE_OUWAN', 206); //
define('XSDKTYPE_NIUA', 207); //
define('XSDKTYPE_XUNLEI', 208); //
define('XSDKTYPE_DXWY', 209); //
define('XSDKTYPE_SKY', 210); //
define('XSDKTYPE_LIANTONG', 211); //
define('XSDKTYPE_DJGAME', 212); //
define('XSDKTYPE_DUOWAN', 213); //
define('XSDKTYPE_IWY', 214); //
define('XSDKTYPE_PPW', 215); //
define('XSDKTYPE_TIANYI', 216); //
define('XSDKTYPE_XMWAN', 217); //
define('XSDKTYPE_3GMH', 218); //
define('XSDKTYPE_TENCENT', 219); //
define('XSDKTYPE_9YOU', 220); //
define('XSDKTYPE_GFAN', 221); //
define('XSDKTYPE_MEIZU', 222); //
define('XSDKTYPE_LENOVO', 223); //
define('XSDKTYPE_DIANXIN', 224); //
define('XSDKTYPE_EWAN', 225); //
define('XSDKTYPE_KUWO', 226); //
define('XSDKTYPE_YOUXIN', 227); //
define('XSDKTYPE_YYW', 228); //
define('XSDKTYPE_XXZS', 229); //
define('XSDKTYPE_KOUDAI', 230); //
define('XSDKTYPE_LIEBAO', 231); //
define('XSDKTYPE_PJ', 232); //
define('XSDKTYPE_YOUYOU', 233); //
define('XSDKTYPE_UUCUN', 234); //
define('XSDKTYPE_BSJ', 235); //
define('XSDKTYPE_TKCX', 236); //
define('XSDKTYPE_LETV', 237); //
define('XSDKTYPE_07073', 238); //
define('XSDKTYPE_CMGB', 241); //
define('XSDKTYPE_CC', 242); //
define('XSDKTYPE_LBYX', 243); //
define('XSDKTYPE_TOUTIAO', 245); //
define('XSDKTYPE_YDWY', 246); //
define('XSDKTYPE_LHH', 247); //
// define('XSDKTYPE_HULI', 248); //
define('XSDKTYPE_ANQU', 249); //
define('XSDKTYPE_LEDI', 250); //
define('XSDKTYPE_SHANDOU', 251); //
define('XSDKTYPE_ZSE', 252); //
define('XSDKTYPE_YUJIA', 253); //
define('XSDKTYPE_MZYW', 254); //
define('XSDKTYPE_SHOUMENG', 255); //
define('XSDKTYPE_LSTY', 256); //
define('XSDKTYPE_LTWY', 257); //
define('XSDKTYPE_43GZ', 258); //
define('XSDKTYPE_BEE', 259); //
define('XSDKTYPE_ZHANGYUE', 260); //
define('XSDKTYPE_KYONG', 261); //
define('XSDKTYPE_SNAIL', 262); //
define('XSDKTYPE_MEITU', 263); //
define('XSDKTYPE_RXY', 264); //
define('XSDKTYPE_MISI', 265); //
define('XSDKTYPE_BIE', 266); //
define('XSDKTYPE_LANCH', 267); //
define('XSDKTYPE_WEIXIN_DC', 268); //
define('XSDKTYPE_WEIXIN_SG', 269); //
define('XSDKTYPE_BUTTERFLY',  270); //h5蝴蝶
define('XSDKTYPE_HAOTENG',  271); //h5豪腾

function ConvertNiuA2MySDKType($strChannel) {
	if ($strChannel == "qihoo") {
		return XSDKTYPE_360;
	} else if ($strChannel == "baidu") {
		return XSDKTYPE_BAIDU;
	} else if ($strChannel == "oppo") {
		return XSDKTYPE_OPPO;
	} else if ($strChannel == "mumayi") {
		return XSDKTYPE_9388_MUMAYI;
	} else if ($strChannel == "nduo") {
		return XSDKTYPE_NDUO;
	} else if ($strChannel == "cmmm") {
		return XSDKTYPE_CMMM;
	} else if ($strChannel == "4399") {
		return XSDKTYPE_4399;
	} else if ($strChannel == "amigo") {
		return XSDKTYPE_JINLI;
	} else if ($strChannel == "feiliu") {
		return XSDKTYPE_FEILIU;
	} else if ($strChannel == "yyh") {
		return XSDKTYPE_YINGYONGHUI;
	} else if ($strChannel == "mzw") {
		return XSDKTYPE_MUZHIWAN;
	} else if ($strChannel == "sina") {
		return XSDKTYPE_SINA;
	} else if ($strChannel == "coolpad") {
		return XSDKTYPE_COOLPAD;
	} else if ($strChannel == "ouwan") {
		return XSDKTYPE_OUWAN;
	} else if ($strChannel == "youyou") {
		return XSDKTYPE_YOUYOU;
	} else if ($strChannel == "sogou") {
		return XSDKTYPE_SOGOU;
	} else if ($strChannel == "qxz") {
		return XSDKTYPE_QXZ1;
	} else if ($strChannel == "niua") {
		return XSDKTYPE_NIUA;
	} else if ($strChannel == "xunlei") {
		return XSDKTYPE_XUNLEI;
	} else if ($strChannel == "dxwy") {
		return XSDKTYPE_DXWY;
	} else if ($strChannel == "uucun") {
		return XSDKTYPE_UUCUN;
	} else if ($strChannel == "xiaomi") {
		return XSDKTYPE_XIAOMI;
	} else if ($strChannel == "huawei") {
		return XSDKTYPE_HUAWEI;
	} else if ($strChannel == "uc") {
		return XSDKTYPE_UC;
	} else if ($strChannel == "anzhi") {
		return XSDKTYPE_ANZHI;
	} else if ($strChannel == "sky") {
		return XSDKTYPE_SKY;
	} else if ($strChannel == "liantong") {
		return XSDKTYPE_LIANTONG;
	} else if ($strChannel == "djgame") {
		return XSDKTYPE_DJGAME;
	} else if ($strChannel == "duowan") {
		return XSDKTYPE_DUOWAN;
	} else if ($strChannel == "iwy") {
		return XSDKTYPE_IWY;
	} else if ($strChannel == "ppw") {
		return XSDKTYPE_PPW;
	} else if ($strChannel == "anfeng") {
		return XSDKTYPE_ANFENG;
	} else if ($strChannel == "tianyi") {
		return XSDKTYPE_TIANYI;
	} else if ($strChannel == "youku") {
		return XSDKTYPE_YOUKU;
	} else if ($strChannel == "49you") {
		return XSDKTYPE_49YOU;
	} else if ($strChannel == "xmwan") {
		return XSDKTYPE_XMWAN;
	} else if ($strChannel == "3gmh") {
		return XSDKTYPE_3GMH;
	} else if ($strChannel == "37wan") {
		return XSDKTYPE_37WAN;
	} else if ($strChannel == "8849") {
		return XSDKTYPE_8849;
	} else if ($strChannel == "tencent") {
		return XSDKTYPE_TENCENT;
	} else if ($strChannel == "9you") {
		return XSDKTYPE_9YOU;
	} else if ($strChannel == "bsj") {
		return XSDKTYPE_BSJ;
	} else if ($strChannel == "wdj") {
		return XSDKTYPE_WANDOUJIA;
	} else if ($strChannel == "gfan") {
		return XSDKTYPE_GFAN;
	} else if ($strChannel == "djoy") {
		return XSDKTYPE_DANGLE;
	} else if ($strChannel == "meizu") {
		return XSDKTYPE_MEIZU;
	} else if ($strChannel == "lenovo") {
		return XSDKTYPE_LENOVO;
	} else if ($strChannel == "dianxin") {
		return XSDKTYPE_DIANXIN;
	} else if ($strChannel == "vivo") {
		return XSDKTYPE_vivo;
	} else if ($strChannel == "ewan") {
		return XSDKTYPE_EWAN;
	} else if ($strChannel == "kuwo") {
		return XSDKTYPE_KUWO;
	} else if ($strChannel == "youxin") {
		return XSDKTYPE_YOUXIN;
	} else if ($strChannel == "yyw") {
		return XSDKTYPE_YYW;
	} else if ($strChannel == "pps") {
		return XSDKTYPE_PPS;
	} else if ($strChannel == "htc") {
		return XSDKTYPE_HTC;
	} else if ($strChannel == "xxzs") {
		return XSDKTYPE_XXZS;
	} else if ($strChannel == "youlong") {
		return XSDKTYPE_YOULONG;
	} else if ($strChannel == "kugou") {
		return XSDKTYPE_KUGOU;
	} else if ($strChannel == "koudai") {
		return XSDKTYPE_KOUDAI;
	} else if ($strChannel == "liebao") {
		return XSDKTYPE_LIEBAO;
	} else if ($strChannel == "tkcx") {
		return XSDKTYPE_TKCX;
	} else if ($strChannel == "pj") {
		return XSDKTYPE_PJ;
	} else if ($strChannel == "letv") {
		return XSDKTYPE_LETV;
	} else if ($strChannel == "07073") {
		return XSDKTYPE_07073;
	} else if ($strChannel == "itools") {
		return XSDKTYPE_ITOOLS;
	} else if ($strChannel == "tbtui") {
		return XSDKTYPE_TONGBUTUI;
	} else if ($strChannel == "toutiao") {
		return XSDKTYPE_TOUTIAO;
	} else if ($strChannel == "cmgb") {
		return XSDKTYPE_CMGB;
	} else if ($strChannel == "cc") {
		return XSDKTYPE_CC;
	} else if ($strChannel == "lbyx") {
		return XSDKTYPE_LBYX;
	} else if ($strChannel == "haima") {
		return XSDKTYPE_HAIMA;
	} else if ($strChannel == "ydwy") {
		return XSDKTYPE_YDWY;
	} else if ($strChannel == "lhh") {
		return XSDKTYPE_LHH;
	} else if ($strChannel == "huli") {
		return XSDKTYPE_HULI;
	} else if ($strChannel == "anqu") {
		return XSDKTYPE_ANQU;
	} else if ($strChannel == "ledi") {
		return XSDKTYPE_LEDI;
	} else if ($strChannel == "shandou") {
		return XSDKTYPE_SHANDOU;
	} else if ($strChannel == "zse") {
		return XSDKTYPE_ZSE;
	} else if ($strChannel == "yujia") {
		return XSDKTYPE_YUJIA;
	} else if ($strChannel == "mzyw") {
		return XSDKTYPE_MZYW;
	} else if ($strChannel == "shoumeng") {
		return XSDKTYPE_SHOUMENG;
	} else if ($strChannel == "lsty") {
		return XSDKTYPE_LSTY;
	} else if ($strChannel == "ltwy") {
		return XSDKTYPE_LTWY;
	} else if ($strChannel == "43gz") {
		return XSDKTYPE_43GZ;
	} else if ($strChannel == "bee") {
		return XSDKTYPE_BEE;
	} else if ($strChannel == "zhangyue") {
		return XSDKTYPE_ZHANGYUE;
	} else if ($strChannel == "kyong") {
		return XSDKTYPE_KYONG;
	} else if ($strChannel == "snail") {
		return XSDKTYPE_SNAIL;
	} else if ($strChannel == "meitu") {
		return XSDKTYPE_MEITU;
	} else if ($strChannel == "rxy") {
		return XSDKTYPE_RXY;
	} else if ($strChannel == "misi") {
		return XSDKTYPE_MISI;
	} else if ($strChannel == "bie") {
		return XSDKTYPE_BIE;
	} else if ($strChannel == "lanch") {
		return XSDKTYPE_LANCH;
	} else if ($strChannel == "h5sgmeitu") {
		return XSDKTYPE_H5MEITU;
	} else {
		return XSDKTYPE_WIN;
	}

}

?>