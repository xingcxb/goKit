# `goKit`


<p>
	<strong>
        一个用于偷懒少写代码的工具类
    </strong>
</p>

<p>
👉 <a href="https://xingcxb.com">https://xingcxb.com</a> 👈
</p>

-------------------------------------------------------------------------------
🫡 向`hutool`致敬
-------------------------------------------------------------------------------

## 📚简介

`goKit`是一个小而全的`Go`工具类库，通过将常用方法封装，降低相关`API`的学习成本，提高工作效率(😏主要是偷懒)，少干重复的活。

`goKit`能节省了开发人员对项目中公用类和公用工具方法的封装时间，使开发专注于业务，同时可以最大限度的避免封装不完善带来的`bug`。

## ⚠️ 注意事项

- 本项目为开源兴趣项目，使用时请自行验证问题，生产环境自己测试

## 🎁`goKit`名称的由来

`goKit` = `go` + `kit`，`go`表示语言；`kit`表示工具包。就粗暴的直译为go的工具包

### 🍺`goKit`如何改变我们的`coding`方式

`goKit`的目标是使用一个工具方法代替一段复杂代码，从而最大限度的避免“复制粘贴”代码的问题，彻底改变我们写代码的方式。

## 🛠️包含组件

一个`Go`基础工具文件，对文件、流、加密解密、转码、正则、线程、`XML`等方法进行封装，组成各种工具文件，同时提供以下组件：

## 📦安装

```shell
  go get https://github.com/xingcxb/goKit
```

> 🔔️注意：`goKit`支持`1.20+`

## 🏗️添砖加瓦

### 🎋分支说明

`goKit`的源码分为两个分支，功能如下：

| 分支       | 作用                          |
|----------|-----------------------------|
| `master` | 主分支，不接收任何`pr`或修改            |
| `dev`    | 开发分支，默认为下个版本的正式版本，接受修改或`pr` |

### 吐槽

#### 关于注释的解释
在`goland`中多行注释折叠时会导致完全看不见，让我不太舒服，所以第一行标准写法，参数只能是参考`Java`注释来了，话说意外的算好用


### 🐞提供bug反馈或建议

提交问题反馈请说明正在使用的`Go`版本呢、`goKit`版本和相关依赖库版本。

- [Github issue](https://github.com/xingcxb/goKit/issues)

### 🧬 贡献代码的步骤

1. 在`Github`上`fork`项目到自己的`repo`
2. 把`fork`过去的项目也就是你的项目`clone`到你的本地
3. 修改代码（记得一定要修改`v5-dev`分支）
4. `commit`后`push`到自己的库（`v5-dev`分支）
5. 登录`Github`在你首页可以看到一个`pull request`按钮，点击它，填写一些说明信息，然后提交即可。
6. 等待维护者合并

### 📐`PR`遵照的原则

`goKit`欢迎任何人为`goKit`添砖加瓦，贡献代码，不过维护者是一个强迫症患者，为了照顾病人，需要提交的`pr`（`pull request`）符合一些规范，规范如下：

1. 注释完备，尤其每个新增的方法应按照`Java`文档规范标明方法说明、参数说明、返回值说明等信息，必要时请添加单元测试，如果愿意，也可以加上你的大名。
2. `goKit`的缩进按照`IDEA`默认（`tab`）缩进，所以请遵守（不要和我争执空格与`tab`的问题，这是一个病人的习惯）。
3. 新加的方法不要使用第三方库的方法，`goKit`尽量遵循无依赖原则（除非偷懒的情况）。
4. 请`pull request`到`dev`分支。`master`是主分支，表示已经发布中央库的版本，这个分支不允许pr，也不允许修改。
5. 我们如果关闭了你的`issue`或`pr`，请不要诧异，这是我们保持问题处理整洁的一种方式，你依旧可以继续讨论，当有讨论结果时我们会重新打开。

-------------------------------------------------------------------------------

## ⭐`Star goKit`

[![Stargazers over time](https://starchart.cc/xingcxb/goKit.svg)](https://starchart.cc/xingcxb/goKit)


## 📜License

`MIT` 许可证 [LICENSE](LICENSE) ©️ 2023 xingcxb

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fxingcxb%2FgoKit.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fxingcxb%2FgoKit?ref=badge_shield)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fxingcxb%2FgoKit.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fxingcxb%2FgoKit?ref=badge_large)

## 🙏 感谢

![](jb_beam.png)
