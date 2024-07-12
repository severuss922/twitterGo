git add .
git commit -m "ultimo commit"
git push
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -tags lambda.norpc -o bootstrap main.go
mmarq\Go\bin\build-lambda-zip.exe -o myFunction.zip bootstrap