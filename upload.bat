git add . 
git commit -m "Added process token logic and claims"
git push
go build main.go
rm -f main.zip
tar -a -cf main.zip main