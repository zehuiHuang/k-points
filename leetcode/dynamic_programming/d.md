#动态规划-背包问题

## 1、理解01背包问题和完全背包问题
dp[i]
weight数组的大小 就是物品个数

01背包二维数组推导公式：
```
dp[i][j]= max(dp[i-1][j],dp[i-1][j-weight[i]]+value[i])
# dp[i][j] 表示从0～i任意选择物品（每个物品最多选一次），那么装满容量为j的背包时的最大价值
#循环遍历：选择先遍历物品，在遍历背包容量（谁先谁后都可以）
for i:=1;i<len(weight);i++{
  for j:=0;j <= bigBag;j++{
    if j<weight[i]{
       dp[i][j]=dp[i-1][j]
    }else{
       dp[i][j]=max(dp[i-1][j],dp[i-1][j-weight[i]]+value[i])
    }
  }
}



```

01背包一维数组推导公式：
```
dp[j] = max(dp[j], dp[j - weight[i]] + value[i])
//循环遍历：先遍历物品，在遍历容量，且容量要倒叙遍历（防止前面被重复赋值）
for i:=0;i<len(weight);i++{
  for j:=bigBag;j>=weight[i];j--{
      dp[j]=max(dp[j],dp[j-wight[i]]+value[i])
  }
}

```


完全背包问题二维数组推导公式：
```
dp[i][j]=max(dp[i-1][j],dp[i][j-weight[i]]+value[i])
# dp[i][j] 表示从0～i任意选择物品（每个物品可以选择无限次），那么装满容量为j的背包时的最大价值
# 循环遍历：先遍历物品，在遍历容量
for i:=1;i<len(weight);i++{
    for j:=0;j<=bigBag;j++{
        if j<weight[i] {
           dp[i][j]=dp[i-1][j]
        }else{
           p[i][j]=max(dp[i-1][j],dp[i][j-weight[i]]+value[i])
        }
    }
}
```

完全背包问题一维数组推导公式：
```
dp[j] = max(dp[j], dp[j - weight[i]] + value[i])
#循环遍历：先遍历物品，在遍历容量
for i:=0;i<len(weight);i++{
    for j:=0;j<=bigBag;j++{
        if j>=weight[i]{
           dp[j]=max(dp[j],dp[j-weight[i]]+value[i])
        }
    }
}

```


## 2、理解01背包问题延伸的组合数和完全背包问题延伸的组合数
---容量为j的背包不放物品i有几种方法 + 容量为j的背包放物品i有几种方法

01背包的延伸问题(组合数)二维数组推导公式：
```
dp[i][j]=dp[i-1][j]+dp[i-1][j-weight[i]]
//循环遍历：先遍历物品，在遍历容量
for i:=1;i<len(weight);i++{
    for j:=0;j<=bigBag;j++{
        if j<weight[i]{
            dp[i][j]=dp[i-1][j]
        }else{
            dp[i][j]=dp[i-1][j]+dp[i-1][j-weight[i]]
        }
    }
}

```

01背包的延伸问题(组合数)一维数组推导公式：
```
dp[j] += dp[j-weight[i]]
#循环遍历：先遍历物品，然后倒叙遍历容量
for i:=1;i<len(weight);i++{
    for j:=bigBag;j>=weight[i];j--{
         dp[j] +=dp[j-weight[i]]
    }
}
```

完全背包的延伸问题(组合数)二维数组推导公式：
```
dp[i][j]=dp[i-1][j]+dp[i][j-weight[i]]

//循环遍历：先遍历物品，在遍历容量
for i:=1;i<len(weight);i++{
    for j:=0;j<=bigBag;j++{
        if j<weight[i]{
            dp[i][j]=dp[i-1][j]
        }else{
            dp[i][j]=dp[i-1][j]+dp[i][j-weight[i]]
        }
    }
}
```

完全背包的延伸问题(组合数)一维数组推导公式：
```
dp[j] += dp[j - weight[i]]
//循环遍历：先遍历物品，在遍历容量
for i:=0;i<len(weight);i++{
    for j:=weight[i];j<=bigBag;j++{
        dp[j] += dp[j - weight[i]]
    }
}
```


## 总结：
关于01背包问题，二维数组双层遍历中的物品和容量谁都可以在外层，而一维数组双层遍历必须要求外层是物品，内层是容量，且容量必须倒叙遍历
关于完全背包，和01背包的区别是dp[i][j-weight[i]]+value[i]（完全背包）, dp[i-1][j-weight[i]]+value[i]（01背包）
关于组合数:
01背包 二维数组： dp[i][j]=dp[i-1][j]+dp[i-1][j-weight[i]] ,一维数组dp[j]+=dp[j-weight[i]]
完全背包二维数组： dp[i][j]=dp[i-1][j]+dp[i][j-weight[i]],  一维数组dp[j]+=dp[j-weight[i]]
虽然组合数中的一维数组的推导公式一样，但是遍历过程不一样，01背包内层遍历的容量要倒叙，完全背包的内层遍历无要求

关于组合，如果不需要考虑数据的顺序，那么外层循环为物品，内层循环为背包容量。
如果要考虑数据的顺序（比如 「1，3」和：「3，1」不同），那么要求外层为背包容量，内层为物品

## 完全背包组合问题考虑顺序:
```
//先背包在物品
for j:=0;j<=target;j++ {
		for i:=0 ;i < len(nums);i++ {
			if j >= weight[i] {
				dp[j] += dp[j-weight[i]]
			}
		}
	}
```

## 完全背包组合问题不考虑顺序:
```
//先物品在背包
for i:=0;i<len(nums);i++{
        for j:=nums[i];j<=target;j++{
            dp[j]+=dp[j-nums[i]]
        }
    }
```




## 规律
动态规划 组合数
# 1、从不重复且可重复选的集合中选择等于目标值的组合数
```
不重复说明不需要特殊处理,可重复选,说明index的下一个迭代不需要+1
模版如下:
var dfs func(sum, index int, path []int)
dfs = func(sum, index int, path []int) {
      //条件判定,并收集结果,并return
      if 条件{
      
      }
      for i~n{
      //裁剪
      }
  }
	
	
```
2、从重复且不可重复选的集合中选择等于目标值的组合数--- 不可重复选的组合数
```
重复说明需要特殊处理,不重复选,说明index的下一个个迭代需要+1
var dfs func(sum, index int, path []int)
dfs = func(sum, index int, path []int) {
      //条件判定,并收集结果,并return
      if 条件{
      
      }
      //先排好序,对相邻选过的不要再选
      for i~n{
      //裁剪
if i > index && candidates[i] == candidates[i-1] {
                        continue
     } }
     ....           
```