@echo off
gopherjs build
copy sim3.js web\go /Y
copy sim3.js.map web\go /Y