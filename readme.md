**api文件格式化**
```shell
goctl api format user.api -dir .
```

**生成api服务**
```shell
goctl api go -api user.api -dir .
```
### 开发顺序
- 编写api文件,生成api服务
- 编写proto文件,生成rpc服务
- api/yaml 新增 rpc配置
- api/config 新增 rpc变量
- api/servicecontext.go 新增rpc变量，新增实例代码
- api/logic，编写业务逻辑
- 新建model目录，编写sql文件，生成crud+cahce代码：goctl model mysql ddl -c -src food.sql -dir .
- rpc/yaml 增加mysql/redis配置
- rpc/config
