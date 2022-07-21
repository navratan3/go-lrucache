package cache

type Strategy interface {
	Get(key int) interface{}
	Put(key int, item interface{})
	Init(capacity int)
	Print()
}
