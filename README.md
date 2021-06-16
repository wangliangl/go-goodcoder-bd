[背景]:
    在调研过程中，经常需要对一些网站进行定向抓取。由于go包含各种强大的库，使用go做定向抓取比较简单。请使用go开发一个迷你定向抓取器mini_spider，实现对种子链接的抓取，并把URL长相符合特定pattern的网页保存到磁盘上。

[程序运行]：
   ./mini_spider -c ../conf -l ../log

[配置文件spider.conf]：
    [spider]
    # 种子文件路径 
    urlListFile = ../data/url.data
    # 抓取结果存储目录 
    outputDirectory = ../output
    # 最大抓取深度(种子为0级)
    maxDepth = 1
    # 抓取间隔. 单位: 秒 
    crawlInterval =  1
    # 抓取超时. 单位: 秒 
    crawlTimeout = 1
    # 需要存储的目标网页URL pattern(正则表达式)
    targetUrl = .*.(htm|html)$
    # 抓取routine数 
    threadCount = 8
   
[种子文件为json格式，示例如下]：
   [
     "http://www.baidu.com",
     "http://www.sina.com.cn",
     ...
   ]  

[要求和注意事项]：
    1. 需要支持命令行参数处理。具体包含:  -h(帮助)、-v(版本)、-c(配置文件路径）、-l（日志文件路径，2个日志：mini_spider.log和mini_spider.wf.log)
    2. 抓取网页的顺序没有限制
    3. 单个网页抓取或解析失败，不能导致整个程序退出。需要在日志中记录下错误原因并继续。
    4. 当程序完成所有抓取任务后，必须优雅退出
    5. 从HTML提取链接时需要处理相对路径和绝对路径
    6. 需要能够处理不同字符编码的网页(例如utf-8或gbk)
    7. 网页存储时每个网页单独存为一个文件，以URL为文件名。注意对URL中的特殊字符，需要做转义
    8. 要求支持多routine并行抓取（注意：这里并不是指简单设置GOMAXPROCS>1)
    9. 代码严格遵守百度Golang编码规范V1.1（http://wiki.baidu.com/pages/viewpage.action?pageId=104882818）
    10. 不仅仅考察编程规范，同时也考察编程思路：代码的可读性和可维护性要好，合理进行pakcage、函数的设计和划分，多线程间的同步、信号量的使用等
    11. 完成相应的单元测试和使用demo。你的demo必须可运行，单元测试有效而且通过
    12. 注意控制抓取间隔（设计上可以针对单个routine控制，也可以针对站点进行控制），避免对方网站封禁百度IP。可以使用Ps Python CM委员会为Python good cooder考试提供测试抓取网站: http://pycm.baidu.com:8081。
    13. 日志库请使用http://icode.baidu.com/repos/baidu/go-lib/log
    14. 主配置文件读取使用https://github.com/go-gcfg/gcfg
    15. html解析请使用 https://go.googlesource.com/net/
    16. 编译环境配置请参考：http://wiki.baidu.com/pages/viewpage.action?pageId=515622823
    17. 对于配置文件的异常，需要进行检查和处理