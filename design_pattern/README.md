## 设计模式
``` https://www.liaoxuefeng.com/wiki/1252599548343744/1264742167474528 ```

### 开闭原则
由Bertrand Meyer提出的开闭原则（Open Closed Principle）是指，软件应该对扩展开放，而对修改关闭。这里的意思是在增加新功能的时候，能不改代码就尽量不要改，如果只增加代码就完成了新功能，那是最好的。

### 里氏替换原则
里氏替换原则是Barbara Liskov提出的，这是一种面向对象的设计原则，即如果我们调用一个父类的方法可以成功，那么替换成子类调用也应该完全可以运行。

设计模式把一些常用的设计思想提炼出一个个模式，然后给每个模式命名，这样在使用的时候更方便交流。GoF把23个常用模式分为创建型模式、结构型模式和行为型模式三类

## 图说设计模式
https://design-patterns.readthedocs.io/zh_CN/latest/index.html


## Go 语言设计模式

[![Build Status](https://travis-ci.org/senghoo/golang-design-pattern.svg?branch=master)](https://travis-ci.org/senghoo/golang-design-pattern)

Go 语言设计模式的实例代码

### 创建型模式

* [简单工厂模式（Simple Factory）](https://github.com/senghoo/golang-design-pattern/tree/master/00_simple_factory)
* [工厂方法模式（Factory Method）](https://github.com/senghoo/golang-design-pattern/tree/master/04_factory_method)
* [抽象工厂模式（Abstract Factory）](https://github.com/senghoo/golang-design-pattern/tree/master/05_abstract_factory)
* [创建者模式（Builder）](https://github.com/senghoo/golang-design-pattern/tree/master/06_builder)
* [原型模式（Prototype）](https://github.com/senghoo/golang-design-pattern/tree/master/07_prototype)
* [单例模式（Singleton）](https://github.com/senghoo/golang-design-pattern/tree/master/03_singleton)

### 结构型模式

* [外观模式（Facade）](https://github.com/senghoo/golang-design-pattern/tree/master/01_facade)
* [适配器模式（Adapter）](https://github.com/senghoo/golang-design-pattern/tree/master/02_adapter)
* [代理模式（Proxy）](https://github.com/senghoo/golang-design-pattern/tree/master/09_proxy)
* [组合模式（Composite）](https://github.com/senghoo/golang-design-pattern/tree/master/13_composite)
* [享元模式（Flyweight）](https://github.com/senghoo/golang-design-pattern/tree/master/18_flyweight)
* [装饰模式（Decorator）](https://github.com/senghoo/golang-design-pattern/tree/master/20_decorator)
* [桥模式（Bridge）](https://github.com/senghoo/golang-design-pattern/tree/master/22_bridge)

### 行为型模式

* [中介者模式（Mediator）](https://github.com/senghoo/golang-design-pattern/tree/master/08_mediator)
* [观察者模式（Observer）](https://github.com/senghoo/golang-design-pattern/tree/master/10_observer)
* [命令模式（Command）](https://github.com/senghoo/golang-design-pattern/tree/master/11_command)
* [迭代器模式（Iterator）](https://github.com/senghoo/golang-design-pattern/tree/master/12_iterator)
* [模板方法模式（Template Method）](https://github.com/senghoo/golang-design-pattern/tree/master/14_template_method)
* [策略模式（Strategy）](https://github.com/senghoo/golang-design-pattern/tree/master/15_strategy)
* [状态模式（State）](https://github.com/senghoo/golang-design-pattern/tree/master/16_state)
* [备忘录模式（Memento）](https://github.com/senghoo/golang-design-pattern/tree/master/17_memento)
* [解释器模式（Interpreter）](https://github.com/senghoo/golang-design-pattern/tree/master/19_interpreter)
* [职责链模式（Chain of Responsibility）](https://github.com/senghoo/golang-design-pattern/tree/master/21_chain_of_responsibility)
* [访问者模式（Visitor）](https://github.com/senghoo/golang-design-pattern/tree/master/23_visitor)
