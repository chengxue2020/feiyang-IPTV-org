<?php

$host = "http://192.168.10.1:35455"; //这里是你的bilibili代理程序域名

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

$path = mk_dir('./bililive/') . 'bililive' . '.m3u';

function bilibili($requesturl)
{
    $header = array(
        'upgrade-insecure-requests: 1',
        'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36',
    );
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $requesturl);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);
    curl_setopt($ch, CURLOPT_HTTPHEADER, $header);
    curl_setopt($ch, CURLOPT_HEADER, FALSE);
    $content = curl_exec($ch);
    curl_close($ch);
    return $content;
}

$res = "";
$i = 1;

do {
    $apires = bilibili("https://api.live.bilibili.com/xlive/web-interface/v1/second/getUserRecommend?page=$i&page_size=300&platform=web");
    $has_more = json_decode($apires, TRUE)["data"]['has_more'];
    $list = json_decode($apires, TRUE)["data"]["list"];
    foreach ($list as $value) {
        $res .= "#EXTINF:-1 tvg-logo=\"" . $value['face'] . "\"" . " group-title=\"{$value['parent_name']}\"," . " {$value['uname']}" . PHP_EOL;
        $res .= "$host/bilibili/" . $value['roomid'] . PHP_EOL;
    }
    $i++;

} while ($has_more == 1);

$str = <<<EOD
#EXTM3U
$res
EOD;

file_put_contents($path, $str);