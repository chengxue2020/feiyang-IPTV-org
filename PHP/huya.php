<?php
date_default_timezone_set("Asia/Shanghai");
$type = empty($_GET['type']) ? "nodisplay" : trim($_GET['type']);
$id = empty($_GET['id']) ? "shangdi" : trim($_GET['id']);
$cdn = empty($_GET['cdn']) ? "hwcdn" : trim($_GET['cdn']);
$media = empty($_GET['media']) ? "flv" : trim($_GET['media']);
$roomurl = "https://m.huya.com/" . $id;

function get_content($apiurl, $flag)
{
    if ($flag == "mobile") {
        $headers = array(
            'Content-Type: application/x-www-form-urlencoded',
            'User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1'
        );
    } else {
        $arr = [
            "appId" => 5002,
            "byPass" => 3,
            "context" => "",
            "version" => "2.4",
            "data" => new stdClass(),
        ];
        $postData = json_encode($arr);
        $headers = array(
            'Content-Type: application/json',
            'Content-Length: ' . strlen($postData),
            'upgrade-insecure-requests: 1',
            'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36'
        );
    }
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $apiurl);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
    curl_setopt($ch, CURLOPT_HTTPHEADER, $headers);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);
    if ($flag == "uid") {
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $postData);
    }
    $data = curl_exec($ch);
    curl_close($ch);
    return $data;
}

function aes_decrypt($ciphertext, $key, $iv)
{
    return openssl_decrypt($ciphertext, 'AES-256-CBC', $key, 0, $iv);
}

$key = "abcdefghijklmnopqrstuvwxyz123456";
$iv = "1234567890123456";
$mediaurl = aes_decrypt("vcnTSiZsSUWtlZRxx+FuRnM7F1b1FlSVueFKcxewvKVbe9bXE49HXuq1dHha2K7cSic4yOuClWpau1RibQeO2g==", $key, $iv);

$uid = json_decode(get_content("https://udblgn.huya.com/web/anonymousLogin", "uid"), true)["data"]["uid"];

function get_uuid()
{
    $now = intval(microtime(true) * 1000);
    $rand = rand(0, 1000) | 0;
    return intval(($now % 10000000000 * 1000 + $rand) % 4294967295);
}

function process_anticode($anticode, $uid, $streamname)
{
    parse_str($anticode, $q);
    $q["ver"] = "1";
    $q["sv"] = "2110211124";
    $q["seqid"] = strval(intval($uid) + intval(microtime(true) * 1000));
    $q["uid"] = strval($uid);
    $q["uuid"] = strval(get_uuid());
    $ss = md5("{$q["seqid"]}|{$q["ctype"]}|{$q["t"]}");
    $q["fm"] = base64_decode($q["fm"]);
    $q["fm"] = str_replace(["$0", "$1", "$2", "$3"], [$q["uid"], $streamname, $ss, $q["wsTime"]], $q["fm"]);
    $q["wsSecret"] = md5($q["fm"]);
    unset($q["fm"]);
    if (array_key_exists("txyp", $q)) {
        unset($q["txyp"]);
    }
    return http_build_query($q);
}

function format($livearr, $uid)
{
    $stream_info = ['flv' => [], 'hls' => []];
    $cdn_type = ['HY' => 'hycdn', 'AL' => 'alicdn', 'TX' => 'txcdn', 'HW' => 'hwcdn', 'HS' => 'hscdn', 'WS' => 'wscdn'];
    foreach ($livearr["roomInfo"]["tLiveInfo"]["tLiveStreamInfo"]["vStreamInfo"]["value"] as $s) {
        if ($s["sFlvUrl"]) {
            $stream_info["flv"][$cdn_type[$s["sCdnType"]]] = $s["sFlvUrl"] . '/' . $s["sStreamName"] . '.' . $s["sFlvUrlSuffix"] . '?' . process_anticode($s["sFlvAntiCode"], $uid, $s["sStreamName"]);
        }
        if ($s["sHlsUrl"]) {
            $stream_info["hls"][$cdn_type[$s["sCdnType"]]] = $s["sHlsUrl"] . '/' . $s["sStreamName"] . '.' . $s["sHlsUrlSuffix"] . '?' . process_anticode($s["sHlsAntiCode"], $uid, $s["sStreamName"]);
        }
    }
    return $stream_info;
}

$res = get_content($roomurl, "mobile");
$reg = "/\<script\> window.HNF_GLOBAL_INIT = (.*) \<\/script\>/i";
preg_match($reg, $res, $realres);
$realdata = json_decode($realres[1], true);

if (array_key_exists("exceptionType", $realdata)) {
    header('location:' . $mediaurl);
    exit();
} elseif ($realdata["roomInfo"]["eLiveStatus"] == 2) {
    $realurl = format($realdata, $uid);
    if ($type == "display") {
        print_r($realurl);
        exit();
    }
    if ($media == "flv") {
        switch ($cdn) {
            case $cdn:
                $mediaurl = $realurl["flv"][$cdn];
                break;
            default:
                $mediaurl = $realurl["flv"]["hwcdn"];
                break;
        }
    }
    if ($media == "hls") {
        switch ($cdn) {
            case $cdn:
                $mediaurl = $realurl["hls"][$cdn];
                break;
            default:
                $mediaurl = $realurl["hls"]["hwcdn"];
                break;
        }
    }
    header('location:' . $mediaurl);
    exit();
} elseif ($realdata["roomInfo"]["eLiveStatus"] == 3) {
    header('location:' . "http:" . base64_decode($realdata["roomProfile"]["liveLineUrl"]));
    exit();
} else {
    header('location:' . $mediaurl);
    exit();
}
