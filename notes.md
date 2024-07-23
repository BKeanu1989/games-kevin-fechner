## bugs

panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x1 addr=0x38 pc=0x8af712]

goroutine 1 [running]:
database/sql.(*Stmt).ExecContext(0xc71aa0?, {0xc749a8?, 0xf855a0?}, {0x0?, 0x1?, 0x1?})
        C:/Program Files/Go/src/database/sql/sql.go:2634 +0x52
database/sql.(*Stmt).Exec(...)
        C:/Program Files/Go/src/database/sql/sql.go:2658
games-go-blueprint/internal/database.CreateTables()
        C:/Users/Kevin/Projects/games.kevin-fechner.site/games-go/internal/database/database.go:94 +0x15e
main.main()
        C:/Users/Kevin/Projects/games.kevin-fechner.site/games-go/cmd/api/main.go:18 +0x85
exit status 2

...
wow
2024/06/26 23:23:22 Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub


---
admin powershell
set CGO_ENABLED=1 worked ...
after that go wont compile at all. -> cant set it back to 0 with
set CGO_ENABLED=0 ...


---
in dockerfile
if cgo is enabled -> executable does exist, but gives the message
```txt
/bin/sh: . /games-kevin-fechner: not found
```

https://slavomirj.medium.com/docker-go-and-cgo-application-build-8f295ab0fe07

