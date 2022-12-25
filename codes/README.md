# **使用说明：**
## **为了保护仓库PHP源码，已经对源码进行不可逆的加密处理，使用方法如下：**
### **推荐使用肥羊PHP集成解密扩展Docker镜像：**  
**一键运行：**  
```
docker run -d --restart unless-stopped --privileged=true -p 5678:80 --name php-env youshandefeiyang/php-env
```  
**然后执行：**   
```
docker cp /你的本地PHP文件目录/gdtv.php php-env:/var/www/html/
```   
**访问地址：**
```
http://你的IP:5678/gdtv.php?id=xxx&xx=xxx...
``` 
### `gudou.php`使用方法：  
1.需要在ROOT的安卓手机/安卓模拟器安装抓包工具（已在software目录提供），本套谷豆源码只适合安卓版本  
2.安装9004版谷豆（已在software目录提供）  
3.注册并登录自己的手机号码，开启抓包后筛选关键词aut002，然后puser是你的手机号，你需要记录`ptoken`和`pserialnumber`以及`cid`  
4.最后在IPTV播放器访问∶  
```
http://xxx.xxx.xxx:xxx/gudou.php?phone=你的手机号&ptoken=你的ptoken&pserialnumber=你的pserialnumber&cid=你的cid&id=谷豆频道对应id
```
