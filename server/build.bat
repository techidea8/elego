rd /s/q website-win
md website-win
go build -o website.exe
COPY mall.exe mall-win\
COPY app.conf website-win\
XCOPY html\*.* website-win\html\  /s /e
XCOPY asset\*.* website-win\asset\  /s /e
XCOPY view\*.* website-win\view\  /s /e
