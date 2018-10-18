## 1. Go语言特性
静态型，编译型，自带垃圾回收，自带并发

***整体风格类似C语言，做了大量简化***

> Goroutine
> Go语言在语言层通过goroutine对函数实现并发执行。goroutine类似于线程但并非线程，goroutine会在go语言运行时进行自动调度。
> 

### 原生支持并发(关键字go)
Go语言的并发是基于goroutine。可以将goroutine理解为一种虚拟线程。Go语言运行时参与调度goroutine，并将goroutine合理地分配到每个CPU中，最大限度的使用CPU性能。
多个goroutine中，Go语言使用channel进行通信，程序可以将需要并发的程序设计为生产者消费者模式，将数据放入通道。通道的另外一端的代码将这些数据进行并发计算并返回结果

## 2. 基本语法和使用
### 变量
var 变量名 变量类型
批量声明：

`var (
  a int
  b string
)`

初始化：
var a = 100 或 a := 100

多重赋值特性
a,b = b,a

**匿名变量_，匿名变量不占用命名空间，也不占用内存**

### 数据类型
#### 1. 整形
指定长度：
int8,int16,int32,int64
uint8,uint16,uint32,uint64
自动匹配平台：int,uint
> 哪些情况使用int和uint
逻辑对整形范围没有特殊需求。例如对象的长度使用内置函数len()返回，这个长度可以根据不同平台的字节长度进行变化。实际使用中切片和map的元素数量等都可以用int来表示。
反之，在二进制传输，读写文件的描述结构时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和uint

#### 2. 浮点型
float32和float64

#### 3. 字符串
字符串在Go语言中以原生数据类型出现
默认utf-8
多行字符串，用反引号：`

#### 4. 字符
Go语言的字符有两种：
1. uint8，或者叫byte，代表ASCII码的一个字符
2. rune（int32），代表utf-8的一个字符

> unicode和utf-8的区别
***unicode是字符集，ASCII也是一种字符集***
字符集为每个字符分配一个唯一的ID，例如a在ASCII和unicode中都是97，“你”在unicode中是20320.在不同国家的字符集中，“你”的ID不同，但是在unicode里，字符的ID不会变
UTF-8是编码规则，将unicode中的Id以某种方式进行编码。utf-8是一种变长编码，1到4个字节不等。

#### 5. 切片（slice）- 能动态分配的空间
创建切片：
a := make([]int,3)

### 指针
指针概念在Go语言中被拆分为两个核心概念：
1. 类型指针，允许对这个指针类型的数据进行修改。传递数据使用指针，而无需拷贝数据。**类型指针不能进行偏移和运算**。
2. 切片，由指向起始元素的原始指针，元素数量和容量组成

受益于这样的约束和拆分，Go语言的指针类型变量拥有指针的高效访问，但又不会发生指针偏移，从而避免非法修改关键性数据问题。同时，垃圾回收也比较容易对不会发生偏移的指针进行检索和回收。

> 切片比原始指针具备更强大的特性，更为安全。切片发生越界时，运行时会报出宕机，并打出堆栈，而原始指针只会崩溃。

#### 注意，Go语言中函数都是传值，会进行拷贝，在函数中进行修改，原来的值不会发生改变

#### 创建指针
str := new(string)
*str = "ninja"
new()函数可以创建一个对应类型的指针，创建过程会分配内存。被创建的指针指向的值为默认值

### 变量生命周期
堆和栈：
堆分配内存和栈分配内存相比，堆适合不可预知大小的内存分配。但是为此付出的代价是分配速度较慢，而且会形成内存碎片

#### 变量逃逸-----自动决定变量分配方式，提高运行效率
> 堆和栈各有优缺点。在C/C++中，需要开发者自己学习如何进行内存分配，选用怎样的内存分配方式来适应不同的算法需求。
Go语言将整个过程整个到编译器中，命名为“变量逃逸分析”。这个技术由编译器分析代码的特征和代码生命期，决定应该在堆还是栈进行内存分配。

变量分析：
go run -gcflags "-m -l" pointer.go

### 字符串
len()函数：取字符串的字节长度（ASCII字符数量）
如果想要统计unicode字符数量，使用utf8.RuneCountInString()

#### 遍历：
ASCII字符串遍历直接使用下标
unicode字符串遍历用for range
 
#### 修改字符串
Go语言的字符串无法直接修改每一个字符元素，智能通过重新构造新的字符串片段并赋值给原来的字符串变量实现。
Go语言的字符串和Java一样，默认是不可变的。
字符串不可变有很多好处，如天生线程安全，大家使用的都是只读对象，无须加锁；此外，方便内存共享，而不必使用写时复制技术；字符串的hash'值也只需制作一份。
修改字符串时，可以通过将字符串转换为byte[]进行修改

#### 常见格式化动词
%v:按值得本来值输出
%+v:在%v的基础上，对结构体字段名和值进行展开
%#v:输出Go语言语法格式的值
%T:输出Go语言语法格式的类型和值

### 常量
const关键字
枚举： iota的用法
```
	type Weapon int


	const (
		Arrow Weapon = iota
		Shuriken
		SniperRifle
		Rifle
		Blower
	)

	fmt.Println(Arrow,Shuriken,SniperRifle,Rifle,Blower)

	var weapon Weapon = Blower

	fmt.Println(weapon)
```

### 类型别名
搞清楚类型别名和类型定义
类型别名：type intAlias = int
类型定义：type intAlias int
类型定义是定义了一个新的类型，可以给这个新类型加上对应的方法，而类型别名则不是，本质上还是原来的东西，只是名字变了

## 3.容器

### 数组
数组是一段固定长度的连续内存区域，在Go语言中，数组从声明时就已经确定，使用时可以修改数组成员，但是数组大小不可变化
var 数组名 【长度】类型
初始化： var team = [3]string{"12","32"}

### 切片
![21fadbaf.png](:storage/4cd53e04-4481-46d2-8bc5-a188bbe74902/21fadbaf.png)
Go语言切片的内部结构包括地址，大小和容量。切片一般用于快速地操作一块数据集合。如果将数据集合比作切糕的话，切片就是你要切得那一块。

切片默认指向一段连续的内存区域，可以是数组，也可以是切片本身。
从数组或者切片生成新的切片拥有如下特性：
* 取出元素不包括结束位置对应的索引，切片最后一个元素使用slice[len(slice)]获取
* 当缺省开始位置时，表示从连续区域开头到结束位置
* 当缺省结束位置时，表示从开始位置到整个连续区域末尾
* 两者同时缺省时，与切片本身等效
* 两者同时为0时，等效于空切片，一般用于切片复位
* 根据索引位置取切片slice元素值时，取值范围是0~len(slice)-1，超界会报运行时错误。生成切片时，结束位置可以填写len(slice)但不会报错

切片有点像C语言的指针，指针可以做运算，但代价是内存操作越界。切片在指针的基础上增加了大小，约束了切片对应的内存区域，切片使用中无法对切片内部的地址和大小做手动调整，因此切片比指针更安全，强大

声明切片：var name []T
#### 使用make函数：
```make([]T,size,cap)```
size：为这个类型分配多少个元素
cap：预分配的元素数量，这个值设定后不影响size，只是能提前分配空间，降低多次分配空间造成的性能问题

> 使用make函数生成的切片一定发生了内存分配操作。但给定开始与结束位置的切片只是将新的切片结构指向了已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作

#### append函数：动态为切片添加元素
Go语言的内建函数append可以为切片动态添加元素。每个切片会指向一片内存空间，这片空间能容纳一定数量的元素。当空间不能容纳一定数量的元素时，切片就会进行扩容，扩容操作往往发生在append操作时。(扩容会进行搬家，全部复制到一块薪内存)

#### copy函数：
使用go语言内建的copy函数，可以迅速地将一个切片的数据复制到另外一个切片空间中，copy函数的使用格式如下：

```copy(destSlice,srcSlice) int```
copy的返回值表示实际发生复制的元素个数

#### 删除元素：
没有专门的删除元素的api，要删除的话将前后的元素连起来
```seq = append(seq[:index],seq[index+1:]...)```

### map-建立事物关联的容器
> 大多数语言中map容器使用两种算法：散列表和平衡树
> 散列表可以简单描述为一个数组，数组的每一个元素是一个列表。根据散列函数获得每个元素的特征值，将特征值作为映射的键。如果特征值重复，表示元素发生碰撞。碰撞的元素将被放在同一个特征值的列表中进行保存。散列表查找复杂度为O(1)，最坏O(n)，散列表需要尽量避免元素碰撞以提高查找效率，这样就需要进行扩容，每次扩容，元素都要重新放置，较为耗时
> 平衡树类似于有父子关系的一个数据树，每个元素在放入树时，都要与一些节点进行比较。平衡树的查找复杂度始终为O(logn)
> 
#### 添加键值对并访问
定义map：
`map[keyType]valueType`
查找
`v,ok := scene[key]`
声明时填充内容：

```
m := map[string]string{
  "w":"forward",
  "a":"left",
  "d":"right",
  "s":"backward",
}
```

#### 遍历
```
for k,v := range scene {
  fmt.println(k,v)
}
```
只遍历值:
```
for _,v := range scene {
  fmt.println(v)
}
```
只遍历键：
```
for k := range scene {
  fmt.println(k)
}
```
#### 删除
`delete(map,key)`

#### 清空
没有提供相关api，重新弄一个就行

### sync.Map
```
    var scene sync.Map

	scene.Store("Greece",97)
	scene.Store("London",100)
	scene.Store("egypt",200)

	fmt.Println(scene.Load("London"))

	scene.Delete("London")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate: ",key,value)
		return true
	})
```

### list-列表：可以快速增删的非连续空间的容器
默认内建的是双向链表
#### 初始化
变量名 := list.New()
或
var 变量名 list.List

#### 插入
pushBack（ele）
pushFront（ele）
PushBackList( *List)
PushFrontList( *List)

#### 使用
```
func main() {

	l := list.New()
	//	尾部添加
	l.PushBack("canon")
	//头部添加
	l.PushFront(67)
	//尾部添加后保存元素句柄
	ele := l.PushBack("first")
	//	在ele之后添加
	l.InsertAfter("high", ele)
	//	在ele之前添加
	l.InsertBefore("noon", ele)
	//	删除元素
	l.Remove(ele)

}
```

#### 遍历
```
    for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
```


## 4.流程控制

#### if的特殊写法
```
if err:=Connect();err!=nil{
  fmt.Println(err)
  return
}
```
Connect是一个带返回值的函数，err:=Connect()是一个语句，执行Connect后，将错误保存到err变量中。err!=nil才是if的判断表达式。这种写法可以将返回值和判断放在一行进行处理，而且返回值的作用范围被限制在if，else的语句组合中
> 在编程中，变量在其实现了变量的功能后，作用范围越小，所造成的问题可能性越小，每一个变量代表一个状态，有状态的地方，状态就会被修改，函数的局部变量只会影响一个函数的执行，但全局变量可能会影响所有代码的执行状态，因此限制变量的作用范围对代码的稳定性有很大的帮助
>

#### for range
Go语言可以使用for range遍历数组，切片，字符串，map及通道。
* 数组，切片，字符串返回索引和值
* map返回键和值
* 通道只返回通道里的值


#### switch
可以使用条件判断，字符串等等，多个case还可以用逗号
默认每个case是独立的，跟C不一样，不会继续往下走，因此不用加break

#### goto
```
    for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				goto breakhere
			}
		}
	}

	return

	breakhere:
		fmt.Println("done")
```
可以用于统一错误处理，消除重复代码：
```
	err := firstCheckError()

	if err != nil {
		goto onExit
	}

	err= secondCheckError()

	if err!=nil {
		goto onExit
	}

	onExit:
		fmt.Println(err)
		exitProcess()
```

## 5.函数
Go语言支持普通函数，匿名函数和闭包，从设计上对函数进行了优化和改进，让函数使用起来更加方便
> Go语言的函数属于一级公民：
> * 函数本身可以作为值进行传递
> * 支持匿名函数和闭包
> * 函数可以满足接口
> 

### 函数声明
func 函数名 (参数列表) 返回参数列表 {
  函数体
}

### 返回值
Go语言支持多返回值，经常使用多返回值的最后一个返回参数来返回函数执行中的错误
```
func fun1 () (a,b int) {
  a = 1
  b = 2
  return 
}
```

### 参数传递
Go语言中传入和返回参数在调用和返回时都使用值传递，这里需要注意的是指针，切片和map等引用型对象指向的内容在参数传递中不会发生复制，而是将指针进行复制，类似于创建一次引用
**注意，结构体是值，不是指针**

### 函数变量-把函数作为值

### 字符串的链式处理-操作与数据分离的设计技巧
使用SQL从数据库获取数据时，可以对原始数据进行排序，分组和去重等操作。SQL将数据的操作与遍历过程作为两个部分进行隔离。这样操作和遍历过程就可以独立地进行设计
```
func main() {

	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	chain := []func(string)string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringProcess(list,chain)

	for _,str := range list{
		fmt.Println(str)
	}

}

func StringProcess(list []string, chain []func(string) string) {

	for index,str := range list{
		result:=str
		for _,proc := range chain{
			result = proc(result)
		}
		list[index] = result
	}
}

func removePrefix(str string) string {
	return strings.TrimPrefix(str,"go")
}
```
***有点像责任链模式***

### 匿名函数
匿名函数往往以变量的方式被传递。经常被用于实现回调函数，闭包等
**回调**
```
func main() {

	visit([]int{1,2,3,4,5}, func(ele int) {
		fmt.Println(ele)
	})

}

func visit(list []int, f func(ele int))  {
	for _,v := range list{
		f(v)
	}
}
```
**使用匿名函数进行操作封装**
```
var skillParam = flag.String("skill", "", "skill to perform")

func main() {

	flag.Parse()

	skill := map[string]func(){
		"fire": func() {
			fmt.Println("fire")
		},
		"run": func() {
			fmt.Println("run")
		},
		"fly": func() {
			fmt.Println("fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	}else {
		fmt.Println("skill not found")
	}

}
```

### 函数类型实现接口-把函数作为接口来调用
# 这部分没看懂

### 闭包：引用了外部变量的匿名函数
闭包就是一个函数引用另外一个函数的变量，因为变量被引用着所以不会被回收，因此可以用来封装一个私有变量。这是优点也是缺点，不必要的闭包只会徒增内存消耗！另外使用闭包也要注意变量的值是否符合你的要求，因为他就像一个***静态私有变量***一样。闭包通常会跟很多东西混搭起来，接触多了才能加深理解，这里只是开个头说说基础性的东西。

闭包: 
概念: 
闭包是可以包含自由变量（未绑定到特定对象）的代码块，这些变量不在这个代码块内或者任何全局上下文中定义，而是在定义代码块的环境中定义。要执行的代码块(由于自由变量包含在代码块中，所以这些自由变量以及他们引用的对象没有被释放)为自由变量提供绑定的计算环境(作用域)。

Go预言中的闭包:
```
package main
import(
  "fmt"
)

func main(){
   var j int = 5
   a:=func() (func()){
       var i int = 10
       return func(){
           fmt.Printf("i,j:%d,%d\n",i,j)
       }
   }()//将一个无需参数返回值为匿名函数的函数赋值给a()

   a()
   j*=2
  // i*=2这样是错的
   a()
}
```
如上所见，虽然j是局部变量但是只要闭包还在使用，那么被闭包引用的变量就会一直存在 
而i除了在内部匿名函数中可以访问外，无法通过其他方式处理，因此保证了i的安全性

闭包是由外部函数嵌套内部函数，内部函数引用了外部函数的变量。闭包形成后，当外部函数用完后，本该销毁的变量，可以被内部函数引用并且保留下来，作为自由变量使用，这就是闭包。

# Go语言闭包

## 一、函数式编程

### 1、函数式编程简介

函数式编程是一种编程模型，将计算机运算看作是数学中函数的计算，并且避免了状态以及变量的概念。
在面向对象思想产生前，函数式编程已经有数十年的历史。随着硬件性能的提升以及编译技术和虚拟机技术的改进，一些曾被性能问题所限制的动态语言开始受到关注，Python、Ruby和Lua等语言都开始在应用中崭露头角。动态语言因其方便快捷的开发方式成为很多人喜爱的编程语言，伴随动态语言的流行，函数式编程也开始流行。

### 2、函数式编程的特点

函数式编程的主要特点如下：
A、变量的不可变性： 变量一经赋值不可改变。如果需要改变，则必须复制出去，然后修改。
B、函数是一等公民： 函数也是变量，可以作为参数、返回值等在程序中进行传递。 
C、尾递归：如果递归很深的话，堆栈可能会爆掉，并导致性能大幅度下降。而尾递归优化技术（需要编译器支持）可以在每次递归时重用stack。

### 3、高阶函数

在函数式编程中，函数需要作为参数传递，即高阶函数。在数学和计算机科学中，高阶函数是至少满足下列一个条件的函数:
A、函数可以作为参数被传递
B、函数可以作为返回值输出

## 二、匿名函数

1、匿名函数简介

匿名函数是指不需要定义函数名的一种函数实现方式，匿名函数由一个不带函数名的函数声明和函数体组成。C和C++不支持匿名函数。

func(x，y int) int {
    return x + y
}
2、匿名函数的值类型

在Go语言中，所有的函数是值类型，即可以作为参数传递，又可以作为返回值传递。
匿名函数可以赋值给一个变量：

f := func() int {
    ...
}
定义一种函数类型：
type CalcFunc func(x, y int) int
函数可以作为值传递：
```
func AddFunc(x, y int) int {
return x + y
}

func SubFunc(x, y int) int {
   return x - y
}

...

func OperationFunc(x, y int, calcFunc CalcFunc) int {
   return calcFunc(x, y)
}

func main() {
   sum := OperationFunc(1, 2, AddFunc)
   difference := OperationFunc(1, 2, SubFunc)
   ...
}
```
函数可以作为返回值：
```
// 第一种写法
func add(x, y int) func() int {
   f := func() int {
      return x + y
   }
   return f
}

// 第二种写法
func add(x, y int) func() int {
   return func() int {
      return x + y
   }
}
```
当函数返回多个匿名函数时建议采用第一种写法：
```
func calc(x, y int) （func(int), func()) {
   f1 := func(z int) int {
      return (x + y) * z / 2
   }

   f2 := func() int {
      return 2 * (x + y)
   }
   return f1, f2
}
```
匿名函数的调用有两种方法：
```
// 通过返回值调用
func main() {
   f1, f2 := calc(2, 3)
   n1 := f1(10)
   n2 := f1(20)
   n3 := f2()
   fmt.Println("n1, n2, n3:", n1, n2, n3)
}

// 在匿名函数定义的同时进行调用：花括号后跟参数列表表示函数调用
func safeHandler() {
   defer func() {
      err := recover()
      if err != nil {
         fmt.Println("some exception has happend:", err)
      }
   }()
   ...
}
```
## 三、闭包

1、闭包的定义

函数可以嵌套定义（嵌套的函数一般为匿名函数），即在一个函数内部可以定义另一个函数。Go语言通过匿名函数支持闭包，C++不支持匿名函数，在C++11中通过Lambda表达式支持闭包。
闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)。
闭包只是在形式和表现上像函数，但实际上不是函数。函数是一些可执行的代码，函数代码在函数被定义后就确定，不会在执行时发生变化，所以一个函数只有一个实例。闭包在运行时可以有多个实例，不同的引用环境和相同的函数组合可以产生不同的实例。
所谓引用环境是指在程序执行中的某个点所有处于活跃状态的约束所组成的集合。约束是指一个变量的名字和其所代表的对象之间的联系。由于在支持嵌套作用域的语言中，有时不能简单直接地确定函数的引用环境，因此需要将引用环境与函数组合起来。

2、闭包的本质

闭包是包含自由变量的代码块，变量不在代码块内或者任何全局上下文中定义，而是在定义代码块的环境中定义。由于自由变量包含在代码块中，所以只要闭包还被使用，那么自由变量以及引用的对象就不会被释放，要执行的代码为自由变量提供绑定的计算环境。
闭包可以作为函数对象或者匿名函数。支持闭包的多数语言都将函数作为第一级对象，即函数可以存储到变量中作为参数传递给其它函数，能够被函数动态创建和返回。
```
func add(n int) func(int) int {
   sum := n
   f := func(x int) int {
      var i int = 2
      sum += i * x
      return sum
   }
   return f
}
```
add函数中函数变量为f，自由变量为sum，同时f为sum提供绑定的计算环境，sum和f组成的代码块就是闭包。add函数的返回值是一个闭包，而不仅仅是f函数的地址。在add闭包函数中，只有内部的匿名函数f才能访问局部变量i，而无法通过其它途径访问，因此闭包保证了i的安全性。
当分别用不同的参数(10, 20)注入add函数而得到不同的闭包函数变量时，得到的结果是隔离的，即每次调用add函数后都将生成并保存一个新的局部变量sum。
在函数式语言中，当内嵌函数体内引用到体外的变量时，将会把定义时涉及到的引用环境和函数体打包成一个整体（闭包）返回。
当每次调用add函数时都将返回一个新的闭包实例，不同实例之间是隔离的，分别包含调用时不同的引用环境现场。不同于函数，闭包在运行时可以有多个实例，不同的引用环境和相同的函数组合可以产生不同的实例。
从形式上看，匿名函数都是闭包。
函数只是一段可执行代码，编译后就固定，每个函数在内存中只有一份实例，得到函数的入口点便可以执行函数。
***对象是附有行为的数据，而闭包是附有数据的行为。***

3、闭包的使用

闭包经常用于回调函数，当IO操作（例如从网络获取数据、文件读写)完成的时候，会对获取的数据进行某些操作，操作可以交给函数对象处理。
除此之外，在一些公共的操作中经常会包含一些差异性的特殊操作，而差异性的操作可以用函数来进行封装。
```
package main

import "fmt"

func adder() func(int) int {
   sum := 0
   f := func(x int) int {
      sum += x
      return sum
   }
   return f
}

func main() {
   sum := adder()
   for i := 0; i < 10; i++ {
      fmt.Println(sum(i))
   }
}
```

## 四、闭包的应用
```
package main

import "fmt"

//普通闭包
func adder() func(int) int {
   sum := 0
   return func(v int) int {
      sum += v
      return sum
   }
}

//无状态、无变量的闭包
type iAdder func(int) (int, iAdder)
func adder2(base int) iAdder {
   return func(v int) (int, iAdder) {
      return base + v, adder2(base + v)
   }
}

//使用闭包实现斐波那契数列
func Fibonacci() func() int {
   a, b := 0, 1
   return func() int {
      a, b = b, a+b
      return a
   }
}

func main() {
   //普通闭包调用
   a := adder()
   for i := 0; i < 10; i++ {
      var s int =a(i)
      fmt.Printf("0 +...+ %d = %d\n",i, s)
   }
   //状态 无变量的闭包 调用
   b := adder2(0)
   for i := 0; i < 10; i++ {
      var s int
      s, b = b(i)
      fmt.Printf("0 +...+ %d = %d\n",i, s)
   }

   //调用斐波那契数列生成
   fib:=Fibonacci()
   fmt.Println(fib(),fib(),fib(),fib(),fib(),fib(),fib(),fib())
}

```

### 可变参数
func 函数名(固定参数列表, v...T)(返回参数列表) {
  函数体
}
特性：
* 可变参数一般被放置在函数列表的末尾，前面是固定参数列表，当没有固定参数时，所有变量将是可变参数
* v为可变参数变量，类型为[]T,也就是拥有多个T元素的切片
* T为可变参数的类型
***传递可变参数***
```
func main() {
	print(1,2,3)
}

func rawPrint(rawList ...interface{}) {
	for _,s := range rawList {
		fmt.Println(s)
	}
}

func print(slist ...interface{}) {
	rawPrint(slist...)
}
```

### 延迟执行语句defer
Go语言的defer语句会将其后面跟随的语句进行。在defer归属的函数即将返回时，将延迟处理的语句按defer得逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行

延迟调用是在defer所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时

#### 使用defer在函数退出时释放资源

#### 处理运行时发生的错误

#### Go语言的错误处理思想包含以下特征：
1. 一个可能造成错误的函数，需要返回值中返回一个错误接口（error）。如果调用是成功的，错误接口将返回nil，否则返回错误
2. 在函数调用后要检查错误，如果发生错误，进行必要的错误处理

#### 错误接口的定义
```
type error interface{
  Error() string
}
```

#### 自定义错误
```
var err = errors.New("this is an error")
```

#### errors包
```
func New(text string) error {
  return &errorString(text)
}

type errorString struct {
  s string
}

func (e *errorString) Error() string {
  return e.s
}
```
### 宕机-panic

panic和recover的关系
* 有panic没recover，程序宕机
* 有panic也有recover捕获，程序不会宕机，执行完对应的defer后，从宕机点退出当前函数后继续执行

## 6. 结构体
Go语言中没有类的概念，也不支持类的继承等面向对象的概念
Go语言的结构体与类都是复合结构体，但Go语言中结构体的内嵌配合接口比面向对象具有更高的扩展性和灵活性
Go语言不仅认为结构体能拥有方法，且每种自定义类型也可以拥有自己的方法

### 定义结构体
type 类型名 struct{
  字段1 类型
  字段2 类型
}

### 实例化
var ins T
ins := new（T）

> 在Go语言中，对结构体进行“&”取地址操作时，视为对该类型进行一次new的实例化操作：
> ins := &T{}
> 取地址实例化是最广泛的一种结构体实例化方式

### 初始化
1. 键值对初始化
```
type People struct {
  name string
  child *people
}

relation := &People{
  name:"爷爷",
  child:&People{
    name:"爸爸",
    child: &People{
      name:"我",
    },
  },
}

```
2. 多个值初始化
直接在括号里面写值，不用写字段名

3. 匿名结构体
```
package main

import "fmt"

func main() {

	msg := &struct {
		id int
		data string
	}{
		1024,
		"hello",
	}

	printMsgType(msg)

}

func printMsgType(msg *struct{
	id int
	data string
}) {

	fmt.Printf("%T\n",msg)

}
```

### 构造函数：结构体和类型的一系列初始化操作的函数封装
Go语言的类型或结构体没有构造函数的功能。结构体的初始化过程可以使用函数封装实现
Go语言中没有提供构造函数相关的特殊机制，用户根据自己的需求，将参数使用函数传递到结构体参数中即可完成构造函数的任务

### 方法
Go语言中的方法是一种作用于特定类型变量的函数。这种特定类型变量叫做接收器。
如果将这种特定类型理解为类或结构体，接收器的概念相当于this或self
在Go语言中，接收器的类型可以使任何类型，不仅仅是结构体，任何类型都可以拥有方法

#### 接收器
格式：
```
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
  函数体
}
```
* 接收器变量：接收器中的参数变量名在命名时，建议使用接收器类型名的第一个小写字母
* 接收器类型：分为指针类型和非指针类型

#### 指针型接收器
指针类型的接收器由一个结构体的指针组成，由于指针的特性，在调用方法时，修改接收器指针的任意成员变量，在方法结束后，都是有效的

#### 非指针型接收器
当方法使用非指针型接收器时，Go语言会在代码运行时将接收器的值复制一份，在非指针 型接收器的方法中可以获取接收器的成员值，但修改后无效

> 在计算机中，小对象由于值复制时的速度较快，所以适合使用非指针接收器，大对象由于复制性能较低，适合使用指针接收器。

### 类型内嵌与结构体内嵌
结构体允许其成员字段在声明时没有字段名而只有类型，这种形式的字段被称为类型内嵌或匿名字段
```
type Data struct{
  int 
  float32
  bool
}
```
类型内嵌其实仍然拥有自己的字段名，只是字段名就是其类型本身而已，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
结构体实例化后，如果匿名的字段类型为结构体，那么可以直接访问匿名结构体里的所有成员，这种方式被称为结构体内嵌

#### 面向对象的组合思想
```
func main() {
	b := new(Bird)
	fmt.Println("bird: ")
	b.Fly()
	b.Walk()

	h := new(Human)
	fmt.Println("human: ")
	h.Walk()
}

type Flying struct {
}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

type Walkable struct {
}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

type Human struct {
	Walkable
}

type Bird struct {
	Flying
	Walkable
}

```

## 7. 接口
***这一段得好好理解***
接口本身是调用方和实现方均遵守的一种协议，大家按照统一的方法命名参数类型和数量来协调逻辑处理的过程
Go语言中使用组合实现对象特性的描述。对象的内部使用结构体内嵌组合对象应该具有的特性，对外通用接口暴露能使用的特性
Go语言的接口设计是非侵入式的，接口编写者无需知道接口被哪些类型实现。而接口实现者只需知道实现的是什么样子的接口，但无需指明实现哪一个接口，或者接口应该由谁实现

### 声明接口
type 接口类型名 interface{
  方法名1(参数列表) 返回列表
  方法名2(参数列表) 返回列表
}
* 接口类型名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面加er如Writer，Stringer
* 方法名：当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包之外的代码访问

> Go语言中的每个接口中的方法的数量不会很多。Go语言希望通过一个接口精准地描述自己的功能，而通过多个接口的嵌入和组合的方式将简单的接口扩展为复杂的接口。
> 
### 实现接口的条件
1. 接口的方法和实现接口的类型方法格式一致
2. 接口中的所有方法均被实现

> Go语言的接口实现是隐式的，无需让实现接口的类型写出实现了哪些接口。这个设计被称为非侵入式设计
> 实现者在编写方法时，无法预测未来哪些方法会变为接口。一旦某个接口创建出来，要求旧的代码来实现这个接口时，就需要修改旧的代码的派生部分，这一般会造成雪崩式的重新编译
> 

### 接口的嵌套组合
接口和接口嵌套组合形成新接口，只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可以被调用

### 在接口和类型之间转换
Go语言中使用接口断言将接口转为另一个接口，也可以将接口转为另外的类型。接口的转换在开发中非常的常见，使用也非常频繁
```
t,ok := i.(T)
```
i代表接口变量
T为要转变的类型
t为转变后的变量

## 8. 包

Go语言的源码复用建立在包的基础上。Go语言的入口main函数所在的包叫main，main包想要引用别的代码，必须以包的方式进行引用

* 一个目录的下的同级文件归属一个包
* 包名可以与其目录不同名
* 包名为main的包为应用程序的入口包

init函数：
1. 每个源码可以使用一个init函数
2. init函数会在程序main函数执行前被自动调用
3. 调用顺序为main中引用的包，以深度优先顺序初始化
4. 同一个包中的多个init函数被调用的顺序不可预期
5. init函数不能被其他函数调用
6. 运行时，被最后导入的包会最先初始化并调用init函数


## 9. 并发
Go语言通过编译器运行时（runtime），从语言上支持了并发的特性。Go语言的并发通过goroutine特性完成。goroutine类似于线程，但是可以根据需要创建多个goroutine并发工作。goroutine是由Go语言的运行时调度完成，而线程是由操作系统调度完成。
Go语言还提供channel在多个goroutine间进行通信。

### goroutine：根据需要随时创建的“线程”
在编写socket网络程序时，需要提前准备一个线程池为每一个socket的收发包分配一个线程。开发人员需要在线程数量和CPU数量之间建立一个对应关系，以保证每个任务能及时被分配到CPU上进行处理，同时避免多个任务频繁地在线程间切换执行而损失效率
虽然线程池为逻辑编写者提供了线程分配的抽象机制。但是如果面对随时可能发生的并发和线程处理需求，线程池就不是非常直观和方便了。能否有一种机制：使用者分配足够多的任务，系统能自动帮助使用者把任务分配到CPU上，让这些任务尽量并发运作。这种机制被称为goroutine

Go程序从main包的main函数开始，在程序启动时，Go程序就会为main函数创建一个默认的goroutine

#### 使用普通函数创建goroutine
Go中使用go关键字为一个函数创建一个goroutine，一个函数可以被创建多个goroutine，一个goroutine必定对应一个函数
```
go 函数名 (参数列表)
```
使用go关键字创建goroutine时，被调用函数的返回值会被忽略
> 如果需要在goroutine中返回数据，使用通道
> 

#### 匿名函数创建goroutine
```
go func(参数列表){
  函数体
}(调用参数列表)
```
> 所有goroutine在main函数结束时会一同结束
> 

并发：把任务在不同的时间点交给处理器进行处理。在同一时间点，任务不会同时运行
并行：把每一个任务分配给每一个处理器独立完成，在同一时间点，任务一定是同时运行

#### goroutine和coroutine
C#，Lua，python语言都支持coroutine特性，coroutine和goroutine都可以将函数或者语句在独立的环境中运行，但是他们之间有两点不同
1. goroutine可能发生并行执行，但coroutine始终顺序执行
狭义的说，goroutine可能发生在多线程环境下，goroutine无法控制自己获取高优先度支持；coroutine始终在单线程，coroutine程序需要主动交出控制权，宿主才能获得控制权并将控制权交给其他coroutine
2. goroutine之间通过channel通信；coroutine使用yield和resume操作
goroutine和coroutine的概念和运行机制都脱胎于早起的操作系统。
coroutine的运行机制属于协作式任务管理，早期的操作系统要每一个应用必须遵守操作系统的任务处理规则，应用程序在不需要使用CPU时，会主动交出CPU的使用权。如果开发者无意间或者故意让应用程序长时间占用CPU，操作系统也无能为力，表现出来的效果就是计算机很容易死机
goroutine属于抢占式任务管理，已经和现有的多线程和多进程任务处理非常类似。应用程序对CPU的控制最终还需要由操作系统来管理，操作系统如果发现应用程序长时间大量占用CPU，那么用户有权终止这个任务

### 通道：在多个goroutine之间进行通信
Go语言中的channel是一种特殊的类型。在任何时候，同时只能有一个goroutine访问通道进行发送和获取数据。goroutine间通过通道就可以通信。
通道像一个传送带或者队列，总是遵循先入先出
```
通道实例 := make(chan 数据类型)
```
* 通道的收发操作在不同的两个goroutine间进行
* 接受将持续阻塞直到发送方发送数据
* 每次接受一个元素
* 发送将持续阻塞直到接受方接收数据

例子：阻塞收发，同步
```
func main() {
	c := make(chan int)
	go printer(c)
	for i := 1; i <= 10; i++ {
		c <- i
	}
	c <- 0
	<-c
}

func printer(c chan int) {
	for {
		data := <-c
		if data == 0 {
			break
		}
		fmt.Println(data)
	}

	c <- 0
}
```

#### 带缓冲的通道
带缓冲的通道在发送时无需等待接收方接受即可完成发送过程，并且不会发生阻塞，只有当存储空间满时才会发生阻塞。同理如果缓冲通道中有数据，接收时将不会发生阻塞，直到通道中没有数据可读时，通道将会再度阻塞
```
通道实例 := make(chan 通道类型，缓冲大小)
```
#### 通道多路复用
```
func main() {
	ch := make(chan string)
	go RPCServer(ch)

	recv, err := RPCClient(ch,"hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received: ", recv)
	}
}

func RPCClient(ch chan string, req string) (string, error) {
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("timeout")
	}
}

func RPCServer(ch chan string) {
	for {
		data := <-ch
		fmt.Println("server received: ", data)
		ch <- "roger"
	}
}
```

更多例子：
```
func main() {
	exit := make(chan int)
	fmt.Println("start")
	time.AfterFunc(time.Second, func() {
		fmt.Println("one second after")
		exit <- 0
	})
	<- exit
}
```
### 同步：保证并发环境下数据访问的正确性


