# **使用说明：**
## **为了保护仓库PHP源码，已经对部分源码进行不可逆的加密处理，使用方法如下：**
### **推荐使用肥羊PHP集成解密扩展Docker镜像：**  
### **一键运行：**  
**amd64架构：**  
```
docker run -d --restart unless-stopped --privileged=true -p 5678:80 --name php-env youshandefeiyang/php-env
```  
**arm64架构：**  
```
docker run -d --restart unless-stopped --privileged=true -p 5678:80 --name php-env youshandefeiyang/php-env:arm64
```  
### **然后执行：**   
```
docker cp /你的本地PHP文件地址/xxx.php php-env:/var/www/html/
```   
### **访问地址：**
```
http://你的IP:5678/xxx.php?id=xxx&xx=xxx...
```
### 小白可以直接参考视频教程∶[点击观看](https://v1.mk/php)
