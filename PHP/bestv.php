<?php
date_default_timezone_set("Asia/Shanghai");
$channel = empty($_GET['id']) ? "cctv16hd4k/15000000" : trim($_GET['id']);
$stream = "http://223.111.117.11/liveplay-kk.rtxapp.com/live/program/live/{$channel}/";
$timestamp = substr(time(), 0, 9) - 7;
$current = "#EXTM3U" . PHP_EOL;
$current .= "#EXT-X-VERSION:3" . PHP_EOL;
$current .= "#EXT-X-TARGETDURATION:3" . PHP_EOL;
$current .= "#EXT-X-MEDIA-SEQUENCE:{$timestamp}" . PHP_EOL;
for ($i = 0; $i < 3; $i++) {
    $timematch = $timestamp . '0';
    $timefirst = date('YmdH', $timematch);
    $current .= "#EXTINF:3," . PHP_EOL;
    $current .= $stream . $timefirst . "/" . $timestamp . ".ts" . PHP_EOL;
    $timestamp = $timestamp + 1;
}
header("Content-Type: application/vnd.apple.mpegurl");
header("Content-Disposition: attachment; filename=index.m3u8");
echo $current;