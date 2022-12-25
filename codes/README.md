# **使用说明：**
## **为了保护仓库PHP源码，已经对源码进行不可逆的加密处理，使用方法如下：**
### **推荐使用肥羊PHP集成解密扩展Docker镜像：**  
**一键运行：**  
```
docker run -d --restart unless-stopped --privileged=true -p 5678:80 --name php-env youshandefeiyang/php-env
```  
**然后执行：**   
```
docker cp /你的本地PHP文件地址/xxx.php php-env:/var/www/html/
```   
**访问地址：**
```
http://你的IP:5678/xxx.php?id=xxx&xx=xxx...
```
### **如果直接运行，小白建议使用宝塔面板，PHP版本需要为8.1：**  
1.首先将phpso文件夹中的扩展下载至你本地指定目录，比如`/home/php-so/tonyenc.so`  
2.然后在PHP主配置文件添加`extension = /home/php-so/tonyenc.so`如果你不懂，可以用宝塔面板部署，如下图所示：
![](https://raw.githubusercontent.com/youshandefeiyang/IPTV/main/logo/jiami.jpg)
3.添加完成后记得保存并重启服务器，然后访问PHP或使用PHP-CLI即可生效！  
### `gudou.php`使用方法：  
1.需要在ROOT的安卓手机安装抓包工具（抓包精灵/小黄鸟）  
2.安装9004版谷豆（腾讯应用宝/豌豆助手历史版本下载）  
3.注册并登录自己的手机号码，开启抓包后筛选关键词aut002，然后puser是你的手机号，你需要记录`ptoken`和`pserialnumber`以及`cid`  
4.谷豆源码只适合安卓版本，最后访问∶  
```
http://xxx.xxx.xxx:xxx/gudou.php?phone=你的手机号&ptoken=你的ptoken&pserialnumber=你的pserialnumber&cid=你的cid&id=谷豆频道对应id
```
