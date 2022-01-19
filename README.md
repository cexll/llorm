# llorm

# Install 

```go
go get -u github.com/go-ll/llorm
```

# Run 

> Start
```go
e := NewEngine("sqlite3", "ll.db")
s := e.NewSession()

type User struct {
    Name string
    Age int64
}

var users []User
```
> Insert 
```go
u1 := &User{Name: "Tom", Age: 18}
u2 := &User{Name: "Sam", Age: 25}
s.Insert(u1, u2, ...)
```
> Find
```go
s.Find(&users{})
```

> Update
```go
s.Where("Name = ?", "Tom").Update("Age", 30)
```

> Delete
```go
s.Where("Name = ?", "Tom").Delete()
```

> Chain
```go
s.Where("Age > 18").Limit(3).Find(&users)
```

> Transaction
```go
e.Transaction(func(s *session.Session) (result interface{}, err error) {
    _ = s.Model(&User{}).CreateTable()
    _, err = s.Insert(&User{"Tom", 18})
    return nil, errors.New("Error")
})
```

> Migrate
```go
e.Migrate(&User{})
```