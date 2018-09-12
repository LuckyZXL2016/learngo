# Go语言学习

## 一、Go：变量定义（[代码](https://github.com/ZzXxL1994/learngo/blob/master/basic/basic)）
**1. 变量定义要点：**
- 变量类型写在变量名之后
- 编译器可推测变量类型
- 没有char，只有rune
- 原生支持复数类型

**2. 内建变量类型：**
- bool, string
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr(指针)
- byte, rune(就是char类型，为了应付多语言的一字节来设计，32位，与int32可混用)
- float32, float64, complex64, complex128(表示实数，128表示实部和虚部分别是64位)

<br>

## 二、Go：条件和循环（[代码1](https://github.com/ZzXxL1994/learngo/tree/master/basic/branch)，[代码2](https://github.com/ZzXxL1994/learngo/tree/master/basic/loop)）
**条件、循环要点：**
- for, if后面的条件没有括号
- if条件里也可以定义变量
- 没有while
- switch不需要break，也可以直接switch多个条件

<br>

## 三、Go：函数和指针（[代码](https://github.com/ZzXxL1994/learngo/tree/master/basic/func)）
**1. 函数要点：**
- 返回值类型写在最后面
- 可返回多个值
- 函数作为参数
- 没有默认参数、可选参数等，只有可变参数列表

**2. 值传递和引用传递：**
- cpp中可以值传递也可以引用传递（&）
- python和java除了系统内建类型是值传递以外，都是引用传递
- go只有值传递，要改变值必须用指针参数

<br>

## 四、Go：内建容器（[代码](https://github.com/ZzXxL1994/learngo/tree/master/container)）
**1. 数组：**
- 数组是值类型
- [10]int和[20]int是不同的类型
- 调用func f(arr [10]int)会拷贝数组。想改变原数组的值，用指针类型参数
- go中一般不直接使用数组

**2. 切片：**

**3. Map：**
- map的操作：
  - 创建：make(map[string]int)
  - 获取元素：m[key]
  - key不存在时，获得Value类型的初始值
  - 用value, ok:=m[key]来判断是否存在key
  - 用delete删除一个key
  - 使用len获得元素个数
- map的遍历：
  - 用range遍历key，或者遍历key-value对
  - 不保证遍历顺利，如需顺序，需手动对key排序
- map的key：
  - map使用哈希表，必须可以比较相等
  - 除了slice，map，function的内建类型都可以作为key
  - Struct类型中只要不包含上述三个字段，也可作为key
  
**4. rune（go中的char）：**
- 使用range遍历pos-rune对
- 使用utf8.RuneCountInString获得字符数量
- 使用len获得字节长度
- 使用[]byte获得字节

<br>

## 五、Go：面向对象（[代码1](https://github.com/ZzXxL1994/learngo/tree/master/tree)，[代码2](https://github.com/ZzXxL1994/learngo/tree/master/queue)）
**1. Go不能算是一种面向对象的语言：**
- go只支持封装，不支持继承和多态
- go只有struct，没有class

**2. 结构（struct）创建在堆上还是栈上？**
- cpp中，局部变量分配在栈上，在外界也要使用的变量要分配到堆上，并且要手动释放
- java中，对象都分配在堆上，有对应的垃圾回收机制
- go中不需要知道分配在堆上还是栈上。比如返回了局部变量的地址，那么是分配在堆上，并有对应的垃圾回收，这些都是编译器自己识别并实现的

**3. 值接收者 vs 指针接收者：**
- 要改变内容必须使用指针接收者
- 结构过大也考虑使用指针接收者
- 一致性：如果有指针接收者，最好都是指针接收者
- 值接收者是go语言特有，值/指针接收者均可接受值/指针

**4. 封装：**
- 名字一般使用CamelCase形式
- 首字母大写：public
- 首字母小写：private
- 首字母大小写是针对包(package)来说的

**5. 包：**
- 每个目录一个包
- main包包含可执行入口，如果这个包里有main函数，那么这个目录下只能有一个main包
- 为结构定义的方法必须放在同一个包内，可以是不同文件

**6. 扩展系统类型或者别人的类型：**
- 定义别名（见[代码](https://github.com/ZzXxL1994/learngo/blob/master/queue/queue.go)）
- 使用组合（见[代码](https://github.com/ZzXxL1994/learngo/blob/master/tree/entry/entry.go)）

<br>

## 六、Go：面向接口（[代码](https://github.com/ZzXxL1994/learngo/tree/master/retriever)）
**1. go语言的duck typing：**
- 具有python，cpp的duck typing的灵活性
- 又具有java的类型检查

**2. 接口变量里有什么?**
- 实现者的类型
- 实现者的值（或指针指向实现者）

**3. 注意：**
- 接口变量里自带指针，接口变量本身采用值传递，几乎不需要使用接口的指针
- 指针接收者实现只能以指针方式使用；值接收者都可以

<br>

## 七、Go：函数式编程（[代码](https://github.com/ZzXxL1994/learngo/tree/master/functional)）
**1. 函数式编程 vs 函数指针：**
- 函数是一等公民：参数，变量，返回值都可以是函数
- 高阶函数
- 函数->闭包

**2. “正统”函数式编程：**
- 不可变性：不能有状态，只有常量和函数
- 函数只能有一个参数
- go没有以上规定

**3. go闭包的应用：**
- 更为自然，不需要修饰如何访问自由变量
- 没有lambda表达式，但是有匿名函数

<br>

## 八、Go：goroutine（[代码](https://github.com/ZzXxL1994/learngo/tree/master/goroutine)）
**1. 与协程Coroutine相似：**
- 轻量级“线程”
- 非抢占式多任务处理，由协程主动交出控制权
- 编译器/解释器/虚拟机层面的多任务
- 多个协程可能在一个或多个线程上运行

**2. goroutine的定义：**
- 任何函数只需加上go就能送给调度器运行
- 不需要在定义时区分是否是异步函数
- 调度器在合适的点进行切换
- 使用-race来检测数据访问冲突

**3. goroutine可能的切换点（参考）：**
- I/O, select, channel
- runtime.Gosched()
- 等待锁，函数调用（有时）






