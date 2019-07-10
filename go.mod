module app

require (
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/module/lcd v0.0.0
)

replace github.com/module/lcd => ./lcd

go 1.12
