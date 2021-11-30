# XUTAPD

# 分支管理办法（暂定）

定义如下分支：

- main: 主分支
- dev：开发分支
- feat： 功能分支
- release：发布分支
- hotfix: bug修复分支

![img](http://static.oschina.net/uploads/img/201302/25142840_pKcL.png)

### main:主版本分支：

源于：

- 主动创建
- 被release合并

可能被merge到

- dev
- release

命名：main

### dev:开发分支

源于：

- main
- release

可能被合并到

- release

命名：dev

### feat:特性分支

源于：

- dev

可能被合并到

- dev

命名：feat-*

### release:发布分支

源于：

- 被dev合并

可能被合并到

- dev
- master

命名规范:release-*

### Hotfix:线上修复分支

源于：

- master

可能被合并到

- develop
- master

> 如果在hotfix准备合并时存在存活的release分支，则将release分支取代dev分支

命名：hotfix-*

### 概述

> 我们开始开发直到拿出第一个原型版本之前都直接在dev分支上进行直接开发，但不要合并到main分支

1.main分支，用于生产环境，不接受出了release外分支的任何merge操作，每一次出现新的main分支必须打上tag

2.dev分支只有一个，所有开发于dev分支上进行，并在dev上拉取新分支,

当dev测试的差不多比较稳定时，合并入release分支，再由release分支合并入master

3.feat分支可以有多个，只从dev中拉取，特性功能完成后只能合并到dev分支

4.release分支，为稳定的带发布分支，只从稳定的dev中拉去，一般只合并到main分支

出现特殊情况如在此分支上修复了bug，需要另外合并到dev分支上

5.Hotfix分支，用于修复出现在已经上线的main分支中的bug，bug修复完成后需要合并到新的main分支，并打上修复缺陷tag，还需要合并到dev分支或release分支。

