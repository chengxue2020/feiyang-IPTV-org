<?php

$liveid = $_GET['rid'] ?? null;
function mk_dir($newdir)
{
    $dir = $newdir;
    if (is_dir('./' . $dir)) {
        return $dir;
    } else {
        mkdir('./' . $dir, 0777, true);
        return $dir;
    }
}
$liveurl = "https://live.douyin.com/$liveid";
$cookietext = './' . mk_dir('cookies/') . md5(microtime()) . '.' . 'txt';
$headers = array(
    'upgrade-insecure-requests: 1',
    'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36'
);
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $liveurl);
curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
curl_setopt($ch, CURLOPT_HEADER, TRUE);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
curl_setopt($ch, CURLOPT_FOLLOWLOCATION, TRUE);
curl_setopt($ch, CURLOPT_COOKIEJAR, $cookietext);
$mcontent = curl_exec($ch);
curl_close($ch);
preg_match('/Set-Cookie:(.*);/iU', $mcontent, $str);
$realstr = $str[1];
$newheader = array(
    "cookie:$realstr",
    'upgrade-insecure-requests: 1',
    'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36'
);
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $liveurl);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
curl_setopt($ch, CURLOPT_HTTPHEADER, $newheader);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);
$data = curl_exec($ch);
curl_close($ch);
$realdata = urldecode($data);
unlink($cookietext);
$reg = "/\"roomid\"\:\"([0-9]+)\"/i";
preg_match($reg, $realdata, $roomid);
$nnreg = "/\"id_str\":\"{$roomid[1]}\"[\s\S]*?\"hls_pull_url\"/i";
preg_match($nnreg,$realdata,$newcontent);
$nnnreg = "/\"hls_pull_url_map\"[\s\S]*?}/i";
preg_match($nnnreg,$newcontent[0],$finalstr);
$mediaArr = json_decode('{' . $finalstr[0] . '}',true);
$hls_url = $mediaArr['hls_pull_url_map']['FULL_HD1'];
header('location:' . $hls_url);
