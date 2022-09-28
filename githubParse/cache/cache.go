package cache

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type Hash [sha256.Size]byte
type Function func(...any) (any, error)

type Cache struct {
	cached map[Hash]any
	*Parameters
}

func init() {
	nowFunc = time.Now
}

// New generate a Cache struct
// Parameters value `t` mustn't be above time.Now()
func New(parameters *Parameters) *Cache {
	c := new(Cache)
	c.cached = make(map[Hash]any)

	if parameters == nil {
		parameters = &Parameters{}
	}
	c.Parameters = parameters

	return c
}

func (c *Cache) Cache(fc Function, args ...any) (any, error) {
	f := c.getFunctionName(fc)
	result, ok := c.Get(f, args...)
	if ok && !c.CanBeUpdate() {
		return result, nil
	}
	hash := c.hash(c.buildKey(f, args))

	result, err := fc(args)
	if err != nil {
		return nil, err
	}

	c.cached[hash] = result

	return result, nil
}

func (c *Cache) Get(f string, args ...any) (any, bool) {
	hash := c.hash(c.buildKey(f, args))
	result, ok := c.cached[hash]

	return result, ok
}

func (c *Cache) CanBeUpdate() bool {
	if c.t.IsZero() {
		return false
	}

	return nowFunc().Sub(c.t) >= c.d*time.Second
}

func (c *Cache) hash(data []byte) Hash {
	return sha256.Sum256(data)
}

func (c *Cache) buildKey(f string, args ...any) []byte {
	argsMerged := c.argsToString(args)

	return []byte(argsMerged + f)
}

func (c *Cache) argsToString(args ...any) string {
	return fmt.Sprintf("%v", args)
}

// From https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
func (c *Cache) getFunctionName(temp any) string {
	strs := strings.Split(runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name(), ".")
	return strs[len(strs)-1]
}
