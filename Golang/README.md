# **使用说明：**  
## 首先说一下为什么用Golang重写，主要还是考虑后期做公共API，Golang更加稳定，效率也高一些，个人用还是优选PHP，因为现成的东西太多了
## 暂时只提供linux-amd64，去action中下载二进制执行文件，然后解压并直接执行`./allinone`，建议搭配进程守护工具进行使用，其它平台自行编译，如果你可以帮我PR交叉编译build.yml以及多平台Dockerfile，我将感激不尽！  
## **抖音：**
### 1，抖音手机客户端进入直播间后，点击右下角三个点，点击分享，点击复制链接，然后运行并访问：
```
http://你的IP:35455/douyin?url=https://v.douyin.com/xxxxxx(&quality=xxxx)
```
其中&quality参数默认origin原画，可以省略，也可以手动指定：uhd、origin、hd、sd、ld
### 2，抖音电脑端需要打开抖音网页版复制`(live.douyin.com/)xxxxxx`，只需要复制后面的xxxxx即可
```
http://你的IP:35455/douyin/xxxxx
```
## **斗鱼：**
### 1，可选m3u8和flv两种流媒体传输方式（不加stream参数默认flv）：
```
http://你的IP:35455/douyin/xxxxx(?stream=hls)
```
### 2，选择flv时可选择不同cdn（不加cdn参数默认`akm-tct.douyucdn.cn`）
```
http://你的IP:35455/douyin/xxxxx(?cdn=ws-tct)
```
## bilibili，虎牙等更多平台后续会陆续添加
