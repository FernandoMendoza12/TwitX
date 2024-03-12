git add . 
git commit -m "Added process token logic and claims"
git push
go build main.com
del main.zip
tar -a -cf main.zip main