## Touch Backend 

**面向web3年轻人的第一款AI LLM大模型 + 社交软件**

产品目标：让全世界web3人都能找到自己的知心人

go语言开发版本：go 1.22.2

后端端口：8006 ，接口：POST

/api/login

```shell
curl --location '127.0.0.1:8006/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
"username":"liu",
"password":"test"
}'
```
成功返回状态码200， 成功

```
{
    "result": 200,
    "message": "登录成功",
    "data": {},
    "requestId": ""
}
```


