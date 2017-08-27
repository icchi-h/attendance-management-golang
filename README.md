# 退勤管理コマンド with golang
入社・退社時刻をtsvファイルに出力するコマンド。  
月ごとにファイルを分割して記録する仕様。

# Usage
```bash
cd /path/to/your project/
git clone <https://github.com/icchi-h/attendance-management-golang>
cd attendance-management-golang
go build ams.go

# パスを通す場合
export PATH=$PATH:/path/to/your project/attendance-management-golang/

ams [start or finish]
```
