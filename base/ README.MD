## 写出高效的代码


## 内存管理相关
1、日志限流:防止错误引发的日志风暴把系统打挂
2、利用对象池技术(甚至复用切片,避免make的开销)
```
var uint64sPool sync.Pool

type Uint64s struct {
	A []uint64
}

func GetUint64s(size int) *Uint64s {
	v := uint64sPool.Get()
	if v == nil {
		return &Uint64s{
			A: make([]uint64, size),
		}
	}
	is := v.(*Uint64s)
	// 关键技巧：复用底层数组，仅调整切片长度
	// 避免了重新 make([]uint64) 的开销
	is.A = SetLength(is.A, size)
	return is
}

func SetLength[T any](a []T, newLen int) []T {
	// 如果新长度大于当前容量，则扩容
	if n := newLen - cap(a); n > 0 {
		a = append(a[:cap(a)], make([]T, n)...)
	}
    //否则,进行切片截取到合适长度
	return a[:newLen]
}

```
3、解决对象池可能的不可控隐私:控制大对象的总数时,利用缓冲的channel更靠谱
```
type inmemoryPart struct {
	creationTime uint64
}

func (mp *inmemoryPart) Reset() {
	mp.creationTime = 0
}

// 容量严格限制为 CPU 核数，防止内存无限膨胀
// cgroup.AvailableCPUs()
var mpPool = make(chan *inmemoryPart, runtime.NumCPU())

// getInmemoryPart 从对象池中获取 inmemoryPart 实例
func getInmemoryPart() *inmemoryPart {
	select {
	case mp := <-mpPool: // 尝试从池中获取
		return mp
	default:
		return &inmemoryPart{} // 池空了，才新建
	}
}

// putInmemoryPart 归还 inmemoryPart 实例到对象池
func putInmemoryPart(mp *inmemoryPart) {
	mp.Reset()
	select {
	case mpPool <- mp: // 尝试归还
	default:
		// 池满了，直接丢弃，等待 GC 回收
	}
}
```
4、极致的复用切片:[:0],即清空切片,长度变为0,保留底层数组的容量

5、分层缓冲区:创建多个对象池,根据对象的大小选择合适的池子

## 并发相关
1、分片锁,减少锁竞争
2、原子操作使用 无锁编程cas	
```
func TestNameCaptor2(t *testing.T) {
	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.StoreUint64(&counter, counter+1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter = ", counter)
}
```
3、本地 Worker Pool模式, 更具CPU数定义worker数,每个Worker Pool只处理自己的任务(通过channel实现),空闲时在帮助其他处理
4、利用channel进行限流
```
select {
case tb.inmemoryPartsLimitCh <- struct{}{}:  // 尝试发送
default:  // 如果无法立即发送（通道已满）
    tb.inmemoryPartsLimitReachedCount.Add(1)  // 记录限制到达次数
    select {
    case tb.inmemoryPartsLimitCh <- struct{}{}:  // 阻塞等待直到可以发送
    case <-tb.stopCh:  // 或者收到停止信号
    }
}
```
