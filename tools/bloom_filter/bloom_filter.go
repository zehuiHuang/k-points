package bloom_filter

import (
	"fmt"
	"github.com/demdxx/gocast"
	"github.com/spaolacci/murmur3"
	"math"
)

// ------------------------哈希编码

type Encryptor struct {
}

func NewEncryptor() *Encryptor {
	return &Encryptor{}
}

func (e *Encryptor) Encrypt(origin string) int32 {
	hasher := murmur3.New32()
	_, _ = hasher.Write([]byte(origin))
	//01111111111111111111111111111111
	fmt.Println(math.MaxInt32)
	fmt.Println(hasher.Sum32())
	i32 := int32(hasher.Sum32() % math.MaxInt32) //因为int32的最大值为2^31-1，所以取模
	fmt.Println(i32)
	return i32
}

// ------------------------布隆过滤器服务

type LocalBloomService struct {
	m, k, n   int32 //m:位图长度；k：hash函数个数；n:存入的数量个数
	bitmap    []int
	encryptor *Encryptor
}

func NewLocalBloomService(m, k int32, encryptor *Encryptor) *LocalBloomService {
	LocalBloomService := &LocalBloomService{
		m:         m,
		k:         k,
		bitmap:    make([]int, m/32+1),
		encryptor: encryptor,
	}
	return LocalBloomService
}

func (l *LocalBloomService) Exist(val string) bool {
	for _, offset := range l.getKEncrypted(val) {
		index := offset >> 5     // 等价于 / 32
		bitOffset := offset & 31 // 等价于 % 32

		if l.bitmap[index]&(1<<bitOffset) == 0 {
			return false
		}
	}
	return true
}

func (l *LocalBloomService) getKEncrypted(val string) []int32 {
	encrypteds := make([]int32, 0, l.k)
	origin := val
	for i := 0; int32(i) < l.k; i++ {
		encrypted := l.encryptor.Encrypt(origin)
		encrypteds = append(encrypteds, encrypted%l.m)
		if int32(i) == l.k-1 {
			break
		}
		origin = gocast.ToString(encrypted)
	}
	return encrypteds
}

func (l *LocalBloomService) Set(val string) {
	l.n++
	for _, offset := range l.getKEncrypted(val) {
		//因为offset的值是encrypted%l.m而来的
		//offset其实是在位图中的一个位置，比如长度为128，offset为32， 数组为int[]{1,2,3,4,5},数组中每一个值都表示32个数据的存在与否
		//所以是128/32，即值32在数组中的作为就在数组中的第四个位置
		//至于数组中的具体值，只是代表32个数据存在与否的一种情况
		index := offset >> 5     // 等价于 / 32  //找到int[]的下标索引，
		bitOffset := offset & 31 // 等价于 % 32  // offset对32取余，计算出bitOffset在当前索引下32位bit的位置，
		// 同时也表示offset的值在32位的bit中的位置，比如20那就表示在第20位是有值的
		l.bitmap[index] |= (1 << bitOffset) //10000...01001 |  100000...0000  把bitOffset这个bit位置设置为1
	}
}
