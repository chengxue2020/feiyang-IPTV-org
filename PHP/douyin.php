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
$firsturl = 'https://live.douyin.com';
$apiurl = "https://live.douyin.com/webcast/web/enter/?aid=6383&web_rid=$liveid";
$cookietext = './' . mk_dir('cookies/') . md5(microtime() + $liveid) . '.' . 'txt';
$headers = array(
    'upgrade-insecure-requests: 1',
    'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36'
);
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $firsturl);
curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
curl_setopt($ch, CURLOPT_HEADER, TRUE);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
curl_setopt($ch, CURLOPT_FOLLOWLOCATION, TRUE);
curl_setopt($ch, CURLOPT_COOKIEJAR, $cookietext);
$mcontent = curl_exec($ch);
curl_close($ch);
preg_match('/Set-Cookie:(.*);/iU', $mcontent, $str);
$realstr = $str[1];
// echo $realstr;
$newheader = array(
    "Cookie:$realstr",
    'upgrade-insecure-requests: 1',
    'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36'
);
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $apiurl);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
curl_setopt($ch, CURLOPT_HTTPHEADER, $newheader);
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);
$data = curl_exec($ch);
curl_close($ch);
unlink($cookietext);
$dataArr = json_decode($data, true);
$hls_url = $dataArr['data']['data'][0]['stream_url']['hls_pull_url_map']['FULL_HD1'];
header('location:' . $hls_url);
