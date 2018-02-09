# GOSU
go-osu library for reading (and writing) osu! DB binary blobs.

# USAGE
```
import (
  "os"
  "github.com/Stymphalian/gosu"
)

func Foo() {
  path := "data/osu!.db"
  file, _ := os.Open(path)
  defer file.Close()

  var db gosu.OsuDb
  err = db.UnmarshalBinary(file)
  if err != nil {
    return
  }

  fmt.Println(db.Version)
}
```

# INSTALLATION
go get -v github.com/Stymphalian/gosu

# INFO
__LICENSE:__ MIT \
__Last Updated__: 2018/02/04

TODO(jordanyu): sqlite DB
TODO(jordanyu): web interface to manage collections
