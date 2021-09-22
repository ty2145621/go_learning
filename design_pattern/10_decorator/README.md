# 装饰模式

装饰模式使用对象组合的方式动态改变或增加对象行为。
Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。
使用匿名组合，在装饰器中不必显式定义转调原对象方法。

这个case写的不好。最好是用两个平行互不干扰的方法

JAVA中的InputStream

```
Decorator模式的目的就是把一个一个的附加功能，用Decorator的方式给一层一层地累加到原始数据源上，最终，通过组合获得我们想要的功能。
例如：给FileInputStream增加缓冲和解压缩功能，用Decorator模式写出来如下：

// 创建原始的数据源:
InputStream fis = new FileInputStream("test.gz");
// 增加缓冲功能:
InputStream bis = new BufferedInputStream(fis);
// 增加解压缩功能:
InputStream gis = new GZIPInputStream(bis);

或者一次性写成这样：
InputStream input = new GZIPInputStream( // 第二层装饰
                        new BufferedInputStream( // 第一层装饰
                            new FileInputStream("test.gz") // 核心功能
                        ));
```