<?php
date_default_timezone_set("Asia/Shanghai");
$channel = empty($_GET['id']) ? "emdy4k_8000" : trim($_GET['id']);
$array = explode("_", $channel);
if (isset($array[1])) {
    $stream = "http://[240e:944:6::22]/live2.rxip.sc96655.com/live/8ne5i_sccn,{$array[0]}_hls_pull_{$array[1]}K/";
} else {
    $stream = "http://[240e:944:6::22]/live2.rxip.sc96655.com/live/8ne5i_sccn,{$array[0]}_hls_pull_4000K/";
}
$timestamp = intval((time() - 60) / 6);
$current = "#EXTM3U" . "\r\n";
$current .= "#EXT-X-VERSION:3" . "\r\n";
$current .= "#EXT-X-TARGETDURATION:6" . "\r\n";
$current .= "#EXT-X-MEDIA-SEQUENCE:{$timestamp}" . "\r\n";
for ($i = 0; $i < 6; $i++) {
    $current .= "#EXTINF:6," . "\r\n";
    $current .= $stream . rtrim(chunk_split($timestamp, 3, "/"), "/") . ".ts" . "\r\n";
    $timestamp = $timestamp + 1;
}
header("Content-Disposition: attachment; filename=index.m3u8");
echo $current;
