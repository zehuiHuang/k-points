参控 https://zhuanlan.zhihu.com/p/670878449

## commit、tree、blob 三类 object
1、blob     ->  文件
   value    -> 文件内容
   key      -> 基于value通过SHA-1生成的摘要

2、tree    -> 文件夹
   value   -> 所有blob的摘要信息,即文件夹下所有文件或文件夹的摘要信息
   key     -> 基于value通过SHA-1生成的摘要

3、c****ommit    -> 提交记录
   value     -> 父commit的key(merge时会存在多个)、author、committer、仓库文件夹tree对应的key、提交时携带的摘要信息
   key       -> 基****于 value 通过 SHA-1 生成的摘要


## 常用命令-------------------------------------------------------

```
git add 1.txt
```
# 
```
git status
```

提交
```
git commit -m ' first commit'
```

查询该 id的 类型，有 tree、blob、commit
```
git cat-file -t 38fd29 
```

查询该 id的 内容，每个类型（tree、blob、commit）的内容不同
```
git cat-file -p 38fd29
```

切换并创建 test 分支
```
git checkout -b test
```

commit提交记录 拓扑结构
```
git log --graph --oneline
```

切换到master分支，把test分支的代码merge到当前master分支，此操作会生成一个新的commit,其parent 分别指向两个分支最开始的Header指针的地址，master分支的Header指针 在指向最新的merge commit
```
git checkout master
git merge test
```
换到master分支，把test分支的代码rebase到当前master分支,此操作会进行变基操作，dev和master公共commit节点A，dev上A后面的所有节点直接拼接到master上A的commit节点后面，形成链，再把master原本A后面的拼上去
如下图：
![img_3.png](img_3.png)
```
git checkout master
git rebase dev
```

若遇到冲突 需要解决完冲突后，执行：[git add 冲突文件] 命令， 之后执行如下命名继续rebase流程
```

git rebase --continue
```
终止rebase
```agsl
git rebase --abort 
```


终止合并
```agsl
git merge --abort 
```

取到一个个 commit 对象并将其延伸到当前分支的尾部
```agsl
git checkout branch2
git cherry-pick 21d0 .. f114
```

冲突时，解决完冲突后执行[git add 冲突文件] 命令， 之后执行如下命名继续cherry-pick流程
```agsl
git cherry-pick --continue
```

终止cherry-pick
```agsl
git cherry-pick --abort
```

远端仓库推送
```
git pull
```

git push

git 

## 场景
1、只是本地Commit,但未提交到远端,要回退到某次Commit(该Commit之前的都未push到远端),使用git reset
执行后当与本地仓库有冲突时,如下图:
如图:![img_1.png](img_1.png)
手动处理即可
```
# 回退
 git reset 77014a....9cfbd9c
 
 # 取消,自动还原
 git revert --abort

```

2、已推送到远程的Commit,若要回退到指定的Commit，优先使用git revert
```
# 回退到指定commit后,会创建一个新的commit,旧的commit并不会丢弃(有冲突时需要处理冲突后在提交)
 git revert 77014a....9cfbd9c
```

3、依赖某分支上的某一个commitId创建一个分支
```
git checkout -b 分支名称  77014a....9cfbd9c
```










文档结构：

一、整体结构:
中央仓库、本地仓库
-- 本地仓库：工作区、暂存区、仓库

工作区-》暂存区   git add <file>... 
暂存区=》本地仓库  git commit 

二、版本控制
核心特性：提交和分支

三、底层存储结构（存储模型）
blob ->相当于文件
tree ->相当于文件夹
commit -> 版本，提交记录

过程描述、图形结合


四、远端push检查
1、校验规则：从提交的commit object 的parent向前遍历，直到找到一个parent的 commit object的key 等于远端最后一个 commit object key
2、