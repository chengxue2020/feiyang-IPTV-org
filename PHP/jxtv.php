<?php
date_default_timezone_set("Asia/Shanghai");
$id = isset($_GET['id'])?$_GET['id']:'jxws';
$n = [
    'jxws' => 'tv_jxtv1.m3u8',//江西卫视
    'jxds' => 'tv_jxtv2.m3u8',//江西都市
    'jxjs' => 'tv_jxtv3.m3u8',//江西经视
    'jxjshd' => 'tv_jxtv3_hd.m3u8',//江西经视高清
    'jxys' => 'tv_jxtv4.m3u8',//江西影视
    'jxgg' => 'tv_jxtv5.m3u8',//江西公共
    'jxse' => 'tv_jxtv6.m3u8',//江西少儿
    'jxxw' => 'tv_jxtv7.m3u8',//江西新闻
    'jxyd' => 'tv_jxtv8.m3u8',//江西移动
    'fsgw' => 'tv_fsgw.m3u8',//江西风尚购物
    'jxtc' => 'tv_taoci.m3u8',//江西陶瓷
];
$timestamp = time();
function etag() {
    $e = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678oOLl9gqVvUuI1";
    $a = "";
    for ($i = 0; $i < 8; $i++) {
        $a .= $e[rand(0,  strlen($e)-1)];
    }
    return $a;
}
$etag = etag();
$auth = md5('1609229748'.$n[$id].$etag .'233face@12&^a');
$token_url = "https://app.jxntv.cn/Qiniu/liveauth/getPCAuth.php";
$info = 't=1609229748&stream='.$n[$id].'&uuid=0f408901f01b';
$headers = array(
    'etag:'.$etag,
    'Authorization:'.$auth,
    "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36",
    "origin: http://www.jxntv.cn",
    "Referer: http://www.jxntv.cn/",
    "Content-Type: application/x-www-form-urlencoded;charset=UTF-8",
);
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $token_url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($ch, CURLOPT_POST, TRUE);
curl_setopt($ch, CURLOPT_POSTFIELDS, $info);
curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);

$response = curl_exec($ch);
curl_close($ch);
$responseArray = json_decode($response, true);
$t = $responseArray['t'];
$token = $responseArray['token'];
$playurl = "https://yun-live.jxtvcn.com.cn/live-jxtv/".$n[$id]."?source=pc&t=".$t."&token=".$token;
header('Location:'.$playurl);