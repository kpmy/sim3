@echo off
gopherjs build
copy sim3.js lib\src\ /Y
copy sim3.js.map lib\src\ /Y