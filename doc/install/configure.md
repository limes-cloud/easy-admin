#### 环境要求
- nodejs
- golang(1.19+)
- mysql
- redis


#### 服务安装

##### 拉去代码
```
git clone https://github.com/limeschool/easy-admin
```

##### 后端安装与启动
```
cd server
go mod tidy
go run main.go -c config.yaml
```


##### 前端安装与启动
```
cd ../web
npm install
npm run dev
```


