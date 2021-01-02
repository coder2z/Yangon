module yangon

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.9.0
	github.com/mattn/go-sqlite3 v1.14.5
	github.com/myxy99/component v0.1.6
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.0
	gopkg.in/src-d/go-git.v4 v4.13.1
)

replace (
	github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
