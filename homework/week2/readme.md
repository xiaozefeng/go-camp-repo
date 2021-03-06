##  我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

sql.ErrNoRows 这个error 表示没有数据
在真实的场景中，sql查不到数据是很正常的，这个error应该在底层就被处理，直接返回空数据给上层，而不需要返回error
所以我的结论是：不应该wrap这个error抛给上层，

**代码**

```go
package week2

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

type PersonRepository struct {
}

func (pr *PersonRepository) QueryByName(name string) ([]byte, error) {
	return nil, sql.ErrNoRows
}

var pr = &PersonRepository{}

func ListPersontByName(name string) ([]*Person, error) {
	res, err := pr.QueryByName(name)
	if err != nil && IsNoRowsErr(err) {
		return []*Person{}, nil
	}
	var persons []*Person
	// fake unmarshal
	err = json.Unmarshal(res, &persons)
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func IsNoRowsErr(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
```

