#动态规划-背包问题

## 1、理解01背包问题和完全背包问题
dp[i]

01背包二维数组推导公式：
```
dp[i][j]= max(dp[i-1][j],dp[i-1][j-weight[i]]+value[i])
# dp[i][j] 表示从0～i任意选择物品（每个物品最多选一次），那么装满容量为j的背包时的最大价值
#循环遍历：选择先遍历物品，在遍历容量（谁先谁后都可以）
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
dp[j]=max(dp[i],dp[j-weight[i]]+value[i])
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
dp[j]=max(dp[j],dp[j-weight[i]]+value[i])
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
    for j:=bigBag;j>=weight[i];j++{
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
关于01背包问题，二维数组双层遍历中的物品和容量谁都可以在外层，而一维数组双层遍历必须要求外层是物品，内存是容量，且容量必须倒叙遍历
关于完全背包，和01背包的区别是dp[i][j-weight[i]]+value[i]（完全背包）, dp[i-1][j-weight[i]]+value[i]（01背包）
关于组合数:
01背包 二维数组： dp[i][j]=dp[i-1][j]+dp[i-1][j-weight[i]] ,一维数组dp[j]+=dp[j-weight[i]]
完全背包二维数组： dp[i][j]=dp[i-1][j]+dp[i][j-weight[i]],  一维数组dp[j]+=dp[j-weight[i]]
虽然组合数中的一维数组的推导公式一样，但是遍历过程不一样，01背包内层遍历的容量要倒叙，完全背包的内层遍历无要求