module GoMD

go 1.12

require (
	github.com/Lofanmi/pinyin-golang v0.0.0-20180808030053-30cdbfc8c2de
	github.com/astaxie/beego v1.11.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	google.golang.org/appengine v1.5.0 // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.1
)

replace gopkg.in/russross/blackfriday.v2 v2.0.1 => github.com/russross/blackfriday/v2 v2.0.1
