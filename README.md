<h1 align="center">
  <a href="https://doc.cms.talelin.com/">
  <img src="https://doc.cms.talelin.com/left-logo.png" width="250"/></a>
  <br>
  lin-cms-go
</h1>

<p align="center">
  <img src="https://img.shields.io/badge/go-%3e%3d1.14-blue.svg" alt="go version" data-canonical-src="https://img.shields.io/badge/go-%3e%3d1.14-blue.svg" style="max-width:100%;"></a>
  <a href="" rel="nofollow"><img src="https://img.shields.io/badge/fiber-2.21.*-green.svg" alt="fiber version" data-canonical-src="https://img.shields.io/badge/fiber-2.21.*-green.svg" style="max-width:100%;"></a>
  <img src="https://img.shields.io/badge/license-license--2.0-lightgrey.svg" alt="lisence" data-canonical-src="https://img.shields.io/badge/license-license--2.0-lightgrey.svg" style="max-width:100%;"></a>
</p>

# 简介

## 预防针

* 本项目非官方团队出品，仅出于学习、研究目的丰富下官方项目的语言支持
* 本项目采取后跟进官方团队功能的形式，即官方团队出什么功能，这边就跟进开发什么功能，开发者不必担心前端适配问题。
* 在上一点的基础上，我们会尝试加入一些自己的想法并实现。
* 局限于本人水平，有些地方还需重构，已经纳入了计划中，当然也会有我没考虑到的，希望有更多人参与进来一起完善。

## 专栏教程 (todo)

* lin cms 全家桶（lin-cms-vue & lin-cms-go）为一个前端应用实现内容管理系统。一套教程入门上手 vue、fiber 两大框架。

* 读者反馈：[《lin cms go &vue 教程》读者反馈贴](https://github.com/xushuhui/lin-cms-go/issues/47)

## 线上文档地址（完善中）

[https://github.com/xushuhui/lin-cms-go](https://github.com/xushuhui/lin-cms-go/)

## 线上 demo

可直接参考官方团队的线上 demo：[http://face.cms.talelin.com/](http://face.cms.talelin.com/)，用户名：super，密码：123456

## 什么是 lin cms？

> lin-cms 是林间有风团队经过大量项目实践所提炼出的一套**内容管理系统框架**。lin-cms 可以有效的帮助开发者提高 cms 的开发效率。

本项目是基于 fiber 5.1 的 lin cms 后端实现。

官方团队产品了解请访问 [talelin](https://github.com/talelin)

## lin cms 的特点

lin cms 的构筑思想是有其自身特点的。下面我们阐述一些 lin 的主要特点。

**lin cms 是一个前后端分离的 cms 解决方案**

这意味着，lin 既提供后台的支撑，也有一套对应的前端系统，当然双端分离的好处不仅仅在于此，我们会在后续提供 nodejs 和 php 版本的 lin。如果你心仪 lin，却又因为技术栈的原因无法即可使用，没关系，我们会在后续提供更多的语言版本。为什么 lin 要选择前后端分离的单页面架构呢？

首先，传统的网站开发更多的是采用服务端渲染的方式，需用使用一种模板语言在服务端完成页面渲染：比如 jinja2、jade 等。 服务端渲染的好处在于可以比较好的支持 seo，但作为内部使用的 cms 管理系统，seo 并不重要。

但一个不可忽视的事实是，服务器渲染的页面到底是由前端开发者来完成，还是由服务器开发者来完成？其实都不太合适。现在已经没有多少前端开发者是了解这些服务端模板语言的，而服务器开发者本身是不太擅长开发页面的。那还是分开吧，前端用最熟悉的 vue 写 js 和 css，而服务器只关注自己的 api 即可。

其次，单页面应用程序的体验本身就要好于传统网站。

更多关于 lin cms 的介绍请访问 [lin cms 线上文档](http://doc.cms.talelin.com/)

**框架本身已内置了 cms 常用的功能**

lin 已经内置了 cms 中最为常见的需求：用户管理、权限管理、日志系统等。开发者只需要集中精力开发自己的 cms 业务即可

## lin cms go 的特点

在当前项目的版本`(0.0.1)`中，特点更多来自于`fiber`框架本身带来的特点。通过充分利用框架的特性，实现高效的后端使用、开发，也就是说，只要你熟悉`fiber`框架，那么对于理解使用和二次开发本项目是没有难度的，即便对于框架的某些功能存在疑问也完全可以通过 fiber 官方的开发手册找到答案。当然我们更欢迎你通过 [issues](https://github.com/xushuhui/lin-cms-go/issues) 来向我们提问：)

在下一个版本中`(>0.0.1)`, 我们会在框架的基础上融入一些自己的东西来增强或者优化框架的使用、开发体验。

## 所需基础

由于 lin 采用的是前后端分离的架构，所以你至少需要熟悉 go 和 vue。

lin 的服务端框架是基于 fiber2.20 的，所以如果你比较熟悉 fiber 的开发模式，那将可以更好的使用本项目。但如果你并不熟悉 fiber，我们认为也没有太大的关系，因为框架本身已经提供了一套完整的开发机制，你只需要在框架下用 go 来编写自己的业务代码即可。照葫芦画瓢应该就是这种感觉。

但前端不同，前端还是需要开发者比较熟悉 vue 的。但我想以 vue 在国内的普及程度，绝大多数的开发者是没有问题的。这也正是我们选择 vue 作为前端框架的原因。如果你喜欢 react or angular，那么加入我们，为 lin 开发一个对应的版本吧。

# 快速开始

## server 端必备环境

* 安装 mysql（version： 5.7+）

* 安装 go 环境 (version： 1.14+)

## 获取工程项目

```bash
git clone https://github.com/xushuhui/lin-cms-go.git
```

> 执行完毕后会生成 lin-cms-go 目录

## 安装依赖包

```bash

go mod tidy

```

## 数据库配置

lin 需要你自己在 mysql 中新建一个数据库，名字由你自己决定。例如，新建一个名为` lin-cms `的数据库。接着，我们需要在工程中进行一项简单的配置。使用编辑器打开 lin 工程根目录下`/configs/config.yaml`，找到如下配置项：

```yaml
data:
  datasource:
    source: 用户名：密码@tcp（服务器地址）/数据库名？charset=utf8mb4&parseTime=True  
```

**请务必根据自己的实际情况修改此配置项**

## 导入数据

接下来使用你本机上任意一款数据库可视化工具，为已经创建好的`lin-cms`数据库运行 lin-cms-go 根目录下的`schema.sql`文件，这个 SQL 脚本文件将为为你生成一些基础的数据库表和数据。

## 运行

如果前面的过程一切顺利，项目所需的准备工作就已经全部完成，这时候你就可以试着让工程运行起来了。在工程的根目录打开命令行，输入：

```bash
go build -o //启动 Web 服务器
```

启动成功后会看到如下提示：

```bash
 Fiber v2.20.2
http://127.0.0.1:3000/

```

打开浏览器，访问``http://127.0.0.1:3000``，你会看到一个欢迎界面，至此，lin-cms-go 部署完毕，可搭配 [lin-cms-vue](https://github.com/TaleLin/lin-cms-vue) 使用了。

## 更新日志

[查看日志](http://github.com/xushuhui/lin-cms-go//)

## 常见问题

[查看常见问题](http://github.com/xushuhui/lin-cms-go/)

## 讨论交流

### 微信公众号

![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
