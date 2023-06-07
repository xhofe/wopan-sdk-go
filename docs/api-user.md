# PcWebLogin

### Req

```json
{
  "header": {
    "key": "PcWebLogin",
    "resTime": 1686118416490,
    "reqSeq": 178232,
    "channel": "api-user",
    "sign": "2df72fb5b35f9b5535e06b786a5810c0",
    "version": ""
  },
  "body": {
    "param": "zb2sV3uNfNxSdGJlZIuto/07ECO5+XwXgwh3biexRDsKpMImBQOeFEza0BH0OlfCvXScOZwcue7HU68qzIsB4rV2dQF6gecJRFtUf6CF+jHdfgmW5IVTvvGiKeYiUCIXqwOzfdrgWCiCHqG81brYS3P3Z9ZnnBSAYQD6q9Yw4qw=",
    "clientId": "1001000021",
    "secret": true
  }
}
```

#### param

```json
{
  "phone": "186****5244",
  "password": "this is password",
  "uuid": "",
  "verifyCode": "",
  "clientSecret": "XFmi9GS2hzk98jGX"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "d7e87d6a95614e8eaef4a2f22bfaf7a6",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": {
      "needSmsCode": "1"
    }
  }
}
```

# PcLoginVerifyCode

### Req

```json
{
  "header": {
    "key": "PcLoginVerifyCode",
    "resTime": 1686118433825,
    "reqSeq": 153713,
    "channel": "api-user",
    "sign": "c17d5b143168fce8dd5e86a4ed5d5d74",
    "version": ""
  },
  "body": {
    "param": "zb2sV3uNfNxSdGJlZIuto4JcpoFt3BgCZ6hy0mZI+IyuItwv33Ftg6SGaDc4c65fHQVowbewSxDw4h06CrLu1FiwSy+2+Q6T1soAI5RnMA4p3gXcwNcJZRwoYbZSyVZ31zs4PrXdxRbOC5/xPlQWZZ03lkT+JutfI6OH+7SvYkIK0Yg6y8dgZrxOAppLrjwp",
    "clientId": "1001000021",
    "secret": true
  }
}
```

#### param

```json
{
  "phone": "186****5244",
  "messageCode": "634978",
  "verifyCode": null,
  "uuid": null,
  "clientSecret": "XFmi9GS2hzk98jGX",
  "password": "this is password"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "1e664e05d9124feabcf08e3dba3d10cb",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": {
      "access_token": "c3f0362b-xxxx-4997-a9c2-0478702a5fea",
      "refresh_token": "457192a6-xxxx-485e-8ab1-0922c17cca69",
      "expires_in": 604799
    }
  }
}
```

# AppQueryUser

### Req

```json
{
  "header": {
    "key": "AppQueryUser",
    "resTime": 1686118434179,
    "reqSeq": 106207,
    "channel": "api-user",
    "sign": "ce928821b32a3808925bd1343f26e78f",
    "version": ""
  },
  "body": {
    "param": "/OP/co9mYqLkkDks24f5wyqYIeq7+K+bJcKXsaRgM61hjy8vBTDNvjfqXhGvMcdTcVudN8Tf6m6nOOGpRPY5PA==",
    "clientId": "1001000021",
    "secret": true
  }
}
```

#### param

```json
{ "accessToken": "c3f0362b-xxxx-4997-a9c2-0478702a5fea" }
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "2d38989e8d1145e8a9fd911b60775ff9",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "OUiQMdAaILKTyrcyS0RSd9Kdd86NnTndTnZ45C1iz1qqLlaMbPA7IEhEsyowFNSgj5lsTwXTWnA8e/gZsV1QW39oK1ZUn0Nw8ACNwHgUNsXSc/UjZ2vPf+mlXqCN2HL9JQLEddvvZER319qeaxrHI1A/t7XZOgX9YChhO9QR1tiakwLxKWaIDH4AnAmR7WSBmo2dWFoRzvJBibTCD3eSAjqGX6hOsZ2h5qyINDtS303qWhQRsYPIu3AVQ9RSub5OJpiLPUSIKS1vNQ9r/m2AgT5lLoSfNkOdoIHNhWbJqfC1elT1JC4IlybttsXQPqj60FK36ujvSLwsvXse+SKa2Q=="
  }
}
```

#### data

```json
{
  "userId": "186****5244",
  "headUrl": "https://panservice.mail.wo.cn/upload/headPortrait/person_1600762418973144.png",
  "userName": "wo_Kbc7Hb",
  "sex": "",
  "birthday": "",
  "isModify": "0",
  "isHeadModify": "0",
  "isSetPassword": "1",
  "registerTime": "20220221185416"
}
```
