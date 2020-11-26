## 应用场景（特别是生产级的Web服务中）

- API
- 慢处理交互

## 目的

- 想要通知所有的goroutine停止运行并返回
- 有时候想让main线程等待goroutine执行完毕（这时候就需要一些方法，让goroutine告诉main他们执行完毕了，那就需要用到通道）

## Context

Go中的context包允许你传递一个"conext"到您的程序，用下面列出的上下文信息，可以来执行运行和返回

- 超时
- 截至日期（deadline）
- 通道

如果你正在执行一个web请求或运行一个系统命令，定义一个`超时`对生产级系统通常是个好主意。因为您所依赖的API运行缓慢，你不希望在系统上备份（backup）请求，因为它可能最终会增加复杂并降低所有请求的执行效率。导致联级效应。这是超时或者截至日期context派上用场的地方

## 创建Context

context包支持以下方式创建和获得context

> `context.Background()`

该函数返回一个**空context**。只能用于高级别（在main或顶级请求处理中）。这能用于派生我们稍后谈及的其他context。

```go
ctx := context.Background()
```

> `context.TODO()`

该函数也是创建一个**空context**。也只能用于高级别或当您不确定使用什么context，或函数以后会更新以便接收一个context。这意味您（或维护者）计划将来要添加context搭配函数

context.TODO()与backgroud代码完全相同。不同的是，静态分析工具可以使用它来验证context是否正确传递，这是一个重要细节，因为静态分析工具可以帮助在早期发现潜在的错误，并且可以连接到CI/CD管道

```go
ctx := context.TODO()

// context包 源码
var (
    background = new(emptyCtx)
    todo = nex(emptyCtx)
)
```

> `context.WithValue(parent Context, key, val interface{}) (ctx Context, cancel CancelFunc)`

此函数接收context并返回`派生context`，其中值val和key关联，并通过`context树`与context一起传递。这意味着一旦获得带有值的context，从中派生的任何context都会获得此值。

不建议使用context值传递关键参数，而是函数应接收签名中的那些值，使其显式化。

```go
ctx := context.WithValue(context.Background(), key, "test")
```

>  `context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)`

这是它开始变得有趣的地方。此函数创建从传入的父context派生的新context。父context可以是后台context或传递给函数的context。

返回派生context和取消函数，只有创建它的函数才能调用取消函数来取消此context。如果你愿意，可以传递取消函数，但是强烈建议不要这么做。这可能导致取消函数的调用者没有意识到取消context的下游影响。可能存在源自此的其他context，这可能导致程序以意外的方式运行。简而言之，永远不要传递取消函数

```go
ctx, cancel := context.WithCancel(context.Background())
```

> `context.WithDeadline(parent Context, d time.Time) (ctx Context, cancel CancelFunc)`

此函数返回其父顶的派生context，当截至日期超过或取消函数被调用时，该context将被取消。

例如，你可以创建一个将在以后的某个时间自动取消的context，并在子函数中传递它。当因为截至日期耗尽而取消该context时，获此context的所有函数都会收到通知去停止并返回

```go
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))
```

> `context.WithTimeout(parent Context, timeout time.Duratrion) (ctx Context, cancel CancelFunc)`

此函数类似于context.WithDeadline，不同之处在于它将持续时间作为参数输入而不是时间对象。此函数返回派生context，如果调用取消函数或者超出持续时间，则会取消该派生context

```go
ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
```