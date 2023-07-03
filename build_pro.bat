
@echo off
for /F %%i in ('git rev-parse --short HEAD') do ( set commitid=%%i)
echo %commitid%
for /F %%j in ('git symbolic-ref --short -q HEAD') do ( set branch=%%j)
echo %branch%
Set dt=%Date:~0,4%%Date:~5,2%%Date:~8,2%
echo %dt%
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -tags nethttpomithttp2 -o student -ldflags "-s -w -X gitee.com/super_sky/mkh_utils.CompanyLogo=maxt -X gitee.com/super_sky/mkh_utils.Version=v3.0.2 -X gitee.com/super_sky/mkh_utils.ServerType=saas -X gitee.com/super_sky/mkh_utils.AppEnv=pro -X gitee.com/super_sky/mkh_utils.Branch=%branch% -X gitee.com/super_sky/mkh_utils.Commit=%commitid% -X gitee.com/super_sky/mkh_utils.BuildTime=%dt%" main.go
@echo off
echo Build Finish. %dt%
pause