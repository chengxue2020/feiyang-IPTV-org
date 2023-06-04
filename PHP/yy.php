<?php
date_default_timezone_set("Asia/Shanghai");
$firstrid = empty($_GET['id']) ? "17829" : trim($_GET['id']);
$quality = "4";

$checkUrl = "https://wap.yy.com/mobileweb/$firstrid";
$header = array(
    'Referer: https://wap.yy.com',
    'User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1'
);
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $checkUrl);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
curl_setopt($ch, CURLOPT_HTTPHEADER, $header);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);
$data = curl_exec($ch);
curl_close($ch);

function aes_decrypt($ciphertext, $key, $iv)
{
    return openssl_decrypt($ciphertext, 'AES-256-CBC', $key, 0, $iv);
}

$key = "abcdefghijklmnopqrstuvwxyz123456";
$iv = "1234567890123456";
$mediaurl = aes_decrypt("vcnTSiZsSUWtlZRxx+FuRnM7F1b1FlSVueFKcxewvKVbe9bXE49HXuq1dHha2K7cSic4yOuClWpau1RibQeO2g==", $key, $iv);

$realdata = "";
$rid = "";
$reg = "/md5Hash[\s\S]*?sid.*'(.*)'.*?getQuery/i";
preg_match($reg, $data, $realdata);
if (count($realdata) > 0) {
    $rid = $realdata[1];
} else {
    header('location:' . $mediaurl);
    exit();
}

$millis_13 = round(microtime(true) * 1000);
$millis_10 = time();
$data = '{"head":{"seq":' . $millis_13 . ',"appidstr":"0","bidstr":"121","cidstr":"' . $rid . '","sidstr":"' . $rid . '","uid64":0,"client_type":108,"client_ver":"5.14.13","stream_sys_ver":1,"app":"yylive_web","playersdk_ver":"5.14.13","thundersdk_ver":"0","streamsdk_ver":"5.14.13"},"client_attribute":{"client":"web","model":"","cpu":"","graphics_card":"","os":"chrome","osversion":"106.0.0.0","vsdk_version":"","app_identify":"","app_version":"","business":"","width":"1536","height":"864","scale":"","client_type":8,"h265":0},"avp_parameter":{"version":1,"client_type":8,"service_type":0,"imsi":0,"send_time":' . $millis_10 . ',"line_seq":-1,"gear":' . $quality . ',"ssl":1,"stream_format":0}}';
$url = "https://stream-manager.yy.com/v3/channel/streams?uid=0&cid=$rid&sid=$rid&appid=0&sequence=$millis_13&encode=json";
$headers = array(
    'Content-Type: text/plain;charset=UTF-8',
    'Referer: https://www.yy.com/',
    'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 Edg/106.0.1370.42'
);
$ch = curl_init($url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $data);

$result = curl_exec($ch);
$result = json_decode($result, true);
curl_close($ch);
if (array_key_exists('avp_info_res', $result)) {
    $a = $result['avp_info_res']['stream_line_addr'];
    $mediaurl = array_values($a)[0]['cdn_info']['url'];
    header('location:' . $mediaurl);
    exit();
} else {
    header('location:' . $mediaurl);
    exit();
}