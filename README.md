# 退勤管理コマンド with golang
入社・退社時刻をtsvファイルに出力するコマンド。  
月ごとにファイルを分割して記録する仕様。

# Usage
## install
```bash
cd /path/to/your project/
git clone <https://github.com/icchi-h/attendance-management-golang>
cd attendance-management-golang
go build ams.go

# パスが通っているディレクトリにリンクを貼る
ln -s /path/to/your project/attendance-management-golang/ams /usr/local/bin/ams
```

プロジェクトのディレクトリにパスを通す場合は以下のコマンド

```bash
export PATH=$PATH:/path/to/your project/attendance-management-golang/
```

## Run

### Record
```
ams [start or finish]
```

### Show log file
```bash
$ ams show
Date	Start Time	Finish Time
2017/08/28(Mon)	09:56	21:39
```
