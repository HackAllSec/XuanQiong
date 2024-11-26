# FAQ

1. Docker 版 MySQL 会出现无数据的情况，详情见：[Issues](https://github.com/HackAllSec/XuanQiong/issues/2)。
解决方法：
找到容器的配置文件，添加以下内容：
```
[mysqld]
character-set-server=utf8mb4
collation-server=utf8mb4_unicode_ci
```
![](images/faq/my_cnf.png)