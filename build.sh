#Script builds the program
#!/usr/bin/env bash

p=`pwd`
for d in $(ls ./); do
  echo "building main/$d"
  cd $p/cm$d
  env GOOS=linux GOARCH=386 go build
done
cd $p
