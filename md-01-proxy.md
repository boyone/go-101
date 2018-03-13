# วิธีติดตั้ง Git Proxy
``git config --global http.proxy http://:@xxx.xx.xxx.x:9980``

``git config --global https.proxy http://:@xxx.xx.xxx.x:9980``

คำสั่งในการเอา git proxy ออก

``git config --global --unset http.proxy``

=========================================

## Option: วิธีการสร้าง shortcut เพื่อใช้ในการ set และ unset proxy
เปิดไฟล์ .gitconfig ปกติจะอยู่ที่ C:\\Users\<username>\.gitconfig
%USERPROFILE%.gitconfig

เพิ่มบรรทัดภายใต้ [alias]

## eg.

[alias]

showproxy = !git config --global --get http.proxy

settnproxy = !git config --global http.proxy http://@xxx.xx.xxx.x:9980

unsetproxy = !git config --global --unset http.proxy

## จากนั้นใน cmd จะสามารถเรียกใช้คำสั่งได้ดังนี้
``showproxy``

``settnproxy``

``unsetproxy``