### 基础信息请阅读

#### 打包
+ go build -o rds_backup.exe -ldflags "-w -s" ./main.go

#### exe加壳(可选,需安装upx)
+ upx --backup --brute rds_backup.exe

