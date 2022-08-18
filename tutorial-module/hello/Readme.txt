mkdir hello
cd hello
go mod init alpk.work/hello
vi hello.go
go mod edit -replace alpk.work/greetings=../greetings
cat go.mod 
go mod tidy
cat go.mod 
go run .

