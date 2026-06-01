cpygiturl=${CPY_GIT:-"https://github.com/strumati/cpy"}

git clone "$cpygiturl"
cd cpy
./build.sh
cd ..
go build main.go