RunCounter:
	go run ./counter/main.go
RunGauge:
	go run ./gauge/main.go
RunHistogram:
	go run ./histogram/main.go
Runsummary:
	go run ./summary/main.go
MacOS64Build:
	SET  CGO_ENABLED=0
	SET GOOS=darwin
	SET GOARCH=amd64
	go build main.go
Linux64Build:
	SET CGO_ENABLED=0
	SET GOOS=linux
	SET GOARCH=amd64
	go build main.go