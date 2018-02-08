The `meltflake` library provides a simple `Melt` function to parse and extract
information from Snowflake IDs. A simple `melter` utility is included.

```
go get -u github.com/nicholasknight/meltflake
go get -u github.com/nicholasknight/meltflake/melter
```

```
import "github.com/nicholasknight/meltflake"
...
flake := meltflake.Melt(381898139189116930, meltflake.Discord)
fmt.Print(flake)
```

For a slightly deeper example of using the melt library, look at
`melter/main.go`.
