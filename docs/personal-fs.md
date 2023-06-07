# QueryAllFiles (root folder)

### Req

```json
{
  "header": {
    "key": "QueryAllFiles",
    "resTime": 1686118435506,
    "reqSeq": 125620,
    "channel": "wohome",
    "version": "",
    "sign": "59ffab70dcb5eb7f56bd656940fb801a"
  },
  "body": {
    "param": "GtmyJSC9Dk/J6UPJgyxFwFo/rUcLdIZrx/5O+5K7amBqviX7ywF7E+wGB/nhsmYRdvP11hZmtUmSvdboF63GLX9dcj5sF1dKlyVRre2kci/xQh0ykcvDDreYtjBowWs+ATZAsv5jpq+z1vP2OISXmQ==",
    "secret": true
  }
}
```

#### param

```json
{
  "spaceType": "0",
  "parentDirectoryId": "0",
  "pageNum": 0,
  "pageSize": 50,
  "sortRule": 0,
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "346b2385d39249e3a625e8630f2f6ac9",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "b+fh7Uuv5WqECmt1B++feofiJE8QlaGAN0GEFe9UYlLgLetBsq0Vq3bfPJQlT89h/Nzfye/UwN/gOSEMS/GYUw0NVilHlz+gFRW5x0rtpsdqjcIF7B6qQnE8mYy2oLrwxku5XHZ3lwBsMdMg6yrwhviXg4I1cqKEAKE94xxMDUaOHAGxSwHmleOLJA9TDt7h5PRA6USeRQrdy/HIambvwY56UK8nwgdocAC/H19DR0hlPko7DAfTpSDps9rEBDebvlk+G88WF5f7S/LgOOwsULeSEf9WGqj+kO+shOkhj5T5IR+AQ/rQqCT8WadKgbig"
  }
}
```

#### data

```json
{
  "files": [
    {
      "familyId": 9263097,
      "fid": "0",
      "creator": "186****5244",
      "size": 0,
      "createTime": "20220221185658",
      "name": "test",
      "shootingTime": "",
      "id": "b2e9b430ed02493782065106291b0308",
      "type": 0,
      "thumbUrl": "",
      "fileType": "0"
    }
  ],
  "systemDirs": []
}
```

# QueryAllFiles (inner folder)

### Req

```json
{
  "header": {
    "key": "QueryAllFiles",
    "resTime": 1686143325563,
    "reqSeq": 186807,
    "channel": "wohome",
    "version": "",
    "sign": "af8a6ceee637f0a6271cb3629993ea79"
  },
  "body": {
    "param": "GtmyJSC9Dk/J6UPJgyxFwFo/rUcLdIZrx/5O+5K7amDJ6s2tQJ16EiHpgh8ehcPv6L+6uhobjCsp77CwZetdNRQ3cQPkJpeJLJq8YS47/GZtcJEy5V4q77RVPqyWVwhPabOyxx+qBkmXmWBufEHYzSCdV0JZU5ETeY/UfneDHtY+FvEYXdXAmRUJAUvzceaT",
    "secret": true
  }
}
```

#### param

```json
{
  "spaceType": "0",
  "parentDirectoryId": "b2e9b430ed02493782065106291b0308",
  "pageNum": 0,
  "pageSize": 50,
  "sortRule": 0,
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "cec591e4d67c4fcfb30733b17119b337",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "b+fh7Uuv5WqECmt1B++feofiJE8QlaGAN0GEFe9UYlLgLetBsq0Vq3bfPJQlT89h/Nzfye/UwN/gOSEMS/GYUw0NVilHlz+gFRW5x0rtpsd17aAFj+Dw1YhLzJyGj5ryWWYaWxSoY734Z2lMykzRPkRYj1iKYkcujBdxxC2MMAKTuAbkROfGuJp/s9IHTN1ufU4fCOpje/u53ISSAPxz07850gNB2Nu13/ANLP0MHORRPB2JUqg4JxDvCDSArA7XkYbCweUUPTQX+LzdrYLEllM3hYxt2gC4NROoq/UlQofrH5+0j0kj6SSvExU6pHq4uRsf9rdiqopfQfgPCcLENo2+Ld9LLdJc9XrYAVUtQRUsj/RL0aeM/V/gGvTByK5oBKVmI+pN2vLXehnjMpDkLOmwRCIrvQ5WQ8XuLPA7hUv6RyGKO3Z36lClQZoZQBHhv4N3QbKQ3RNXpuMeied20EdgHhOnCQFp1DPNuuqpT7mSnWpslFFwTI5vLcWRUkGCtGUU9S6z0AMrmzc4e1OxZ2a195ZGk6UfaZuBB9cdZvMxsAVMfmfaLE+5ZYkpQgu1iNgvQRRQR5NdOXBe0eIKpDUPsU1EMek5xIYjZLJUjUGmnyVaFVGnRXlZ3jzQDlvPSADPveLgjJXKmeUaK9NdhZwxSUPWpy3cmVyNDXnLxuepvL77B+0busyUHH8sXtwI2Iz2WONPdINP9Lkvz7Bl5cgIjH+4JH69Qfx/gVSeLTvZe1nAvusJscsgCrVLWjRIbn2KLovrGlumlEJHskduefztEr0Gcq/V1O1156EBu1D82lKzVgEp+ja29CZHXctB"
  }
}
```

#### data

```json
{
  "files": [
    {
      "familyId": 9263097,
      "fid": "0",
      "creator": "186****5244",
      "size": 0,
      "createTime": "20230607210848",
      "name": "inner_folder",
      "shootingTime": "",
      "id": "1bb8cfb12d7f4d4eaac1b3a5293a256f",
      "type": 0,
      "thumbUrl": "",
      "fileType": "0"
    },
    {
      "familyId": 0,
      "fid": "nmG2IzeBFZ5hRo0KwMzfhQvuEVhj5HbeHyDIgu5BndubQ06FNZe3bA/qu/Yrf/iq",
      "creator": "186****5244",
      "size": 1599939,
      "createTime": "20220221190017",
      "name": "bg.png",
      "shootingTime": "",
      "id": "6a06db1c48f54cb58adb538591b3c613",
      "type": 1,
      "thumbUrl": "https://du.smartont.net:8442/openapi/thumbnails?fid=nmG2IzeBFZ5hRo0KwMzfhQvuEVhj5HbeHyDIgu5BndubQ06FNZe3bA%2Fqu%2FYrf%2Fiq",
      "fileType": "1"
    }
  ]
}
```

# GetSearchDirectory

### Req

```json
{
  "header": {
    "key": "GetSearchDirectory",
    "resTime": 1686142912132,
    "reqSeq": 115425,
    "channel": "wohome",
    "version": "",
    "sign": "e9e0d26910655fd878c0471ee6f21488"
  },
  "body": {
    "param": "doFCig8jnOT0I8MM0nYaJocdJ83NH+TxNua0ZwcmDbzI6WxufMMNB7bC/+2xzNEtxEAuKR9lgQgmGBxNo+r8ent6f5ScB0cqBrPQxC8lUKI=",
    "secret": true
  }
}
```

#### param

```json
{ "directoryId": "b2e9b430ed02493782065106291b0308", "clientId": "1001000021" }
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "b413f67b734c4d259338cef4ec1f9f3b",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "aFjqLh22OTF79zuEpdaLbpo7bB0aBl2B2SR75B3ERj93ctlwtNAdlHJEUF4JpDdK/Tnu52LppXAzAR4YEAgiWQiOXRW1AtER/hQicd7iSyTq12CZlW24wxgipoRcvBZZDlu/ahKLPSkESd3HeZbyVA=="
  }
}
```

#### data

```json
{
  "directoryInfo": [
    {
      "directoryId": "b2e9b430ed02493782065106291b0308",
      "directoryName": "test"
    }
  ],
  "flag": false
}
```

# GetDownloadUrlV2

### Req

```json
{
  "header": {
    "key": "GetDownloadUrlV2",
    "resTime": 1686146042800,
    "reqSeq": 174342,
    "channel": "wohome",
    "version": "",
    "sign": "7dd0297e1bb0c1042ebb0eaecd777658"
  },
  "body": {
    "param": "DVt84prwHvHwQf1sM3V+C8MWh5yw3w7qjkHHNi1rRLKlMzWXmXbUCr9qPzIIAWMiJoi6Ehw+Rm2C08b9e4JXn65vW9tgO/vY40ByTO/qryiIZLqNF8M29cPyv4Ai/ax3EUJfx7E2GX+oCExMKg30Y/sQpGrhtkuBMzXs+M6hpks=",
    "secret": true
  }
}
```

#### param

```json
{
  "type": "1",
  "fidList": [
    "Aovuy_d/kEVpB7Or7Lu79/3Cu7SLXQOe8ud1JAnvBMINggjkzYANI7rUN2T3y9kJxi9AU4"
  ],
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "dbfe150fee5a4d41b2439e761c0c07e7",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "Yo2tB2COSjTseVRJ85LFn9HjXnPUCxcIXUFPpXE74Tf3qcSYovBup6T4aHgM7g9jIpE5syNB9GQK5UCXnQm3Oe/qkbJwWZnG0Qz/K47wRwPIE12Qd/pl7oSw/72ETNXXQYWJ0l1bC7lZAz02OucyCOhiFSjiPtHKLgNJ61M++ZuSeZHlQXq1zOS04/Vo1edOIpn2D/B6LPANPCTBUbBDkK6om0Z/YOLIy2mUAwyRGVku6SfVXZ7SyRayfi62O3+e88VKcPcVveXDCw7kKaQTXFWzpG/nlwkakQUt89pyhNfgsJfQxWVTyjLCE0JubPx14gB9O2MKkFkuerddkzIg7sUGEs7MkSbgDldgIYwmTiw="
  }
}
```

#### data

```json
{
  "type": 1,
  "list": [
    {
      "fid": "Aovuy_d/kEVpB7Or7Lu79/3Cu7SLXQOe8ud1JAnvBMINggjkzYANI7rUN2T3y9kJxi9AU4",
      "downloadUrl": "https://gxdownload.pan.wo.cn:8445/openapi/download?fid=nXDdE_9Lic/LcDX63dzibzQQ9jpOIV4YLDQG93rHjcCc/gThpSMUbHUhc/fN377XLBhxlCwEbiIclWoopVnRPmwYq0PA=="
    }
  ]
}
```

# GetDownloadUrl

### Req

```json
{
  "header": {
    "key": "GetDownloadUrl",
    "resTime": 1686143141212,
    "reqSeq": 145379,
    "channel": "wohome",
    "version": "",
    "sign": "754ae87c4658b47a228b624bfbc91fc9"
  },
  "body": {
    "param": "M3MidlpAX5xwxYhGz6n84GXgLUkxP/1lsW9vL1/TQE5zEu3UbKWbXr17iqxq97rS3pIjSXBMgNKcy2G2Vr79S6R1gx8VdslvVulpo4Gfx46+73C3d1wgbxhGahNwG9QZogtLxfoNY9AzFqtAFkgBAO/o05XcjF8yc32hs+FOzxY=",
    "secret": true
  }
}
```

#### param

```json
{
  "fidList": [
    "nmG2IzeBFZ5hRo0KwMzfhQvuEVhj5HbeHyDIgu5BndubQ06FNZe3bA/qu/Yrf/iq"
  ],
  "clientId": "1001000001",
  "spaceType": "0"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "2422a391172047c191c53f836d6d6662",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "4IE9/A3FVJxIGFgZ8CHjhXx0ecDNRbIDrrtwb4cpf9MNA7P6QzSlUWmcaC5+BTv8mJMoett2wsI1g009kUuXp73wvPtytRCRRo18wjLWUe5ynQ4yIy2v8mbd1Fvk13P6HqSqqbHJ1OMnq1XJ97BiL4Sqjz7nPwiFQFYvbj52712nKeFLAU6h26Yly4JdeeRTN4gvvx9Ll5iCPv4zClXTNzpuxIpAMZQbErwT87HMkK3jRzLYsx9BpxmcMpomIQrzkE/+ae7/Sh30CqiRaKp8RxKzEUylN66tvnTRchCZ3aMh7JLzjFee8xiZqVKWqJxrk2Q9um9uCMJNXPJFJeRpcw=="
  }
}
```

#### data

```json
[
  {
    "fid": "nmG2IzeBFZ5hRo0KwMzfhQvuEVhj5HbeHyDIgu5BndubQ06FNZe3bA/qu/Yrf/iq",
    "downloadUrl": "https://du.smartont.net:8445/openapi/download?fid=R0WE1_uqDEXtx9lmogx8Zf/vxr7zNDsu/OBPwNgLIka%2Bl0vdtIWQQ2aHV%2Bw73GCGgcoJz0CmVeWURStW8kPohAA/oCeA=="
  }
]
```

# CreateDirectory

### Req

```json
{
  "header": {
    "key": "CreateDirectory",
    "resTime": 1686143325362,
    "reqSeq": 172039,
    "channel": "wohome",
    "version": "",
    "sign": "160309e7e8636b23e361f5c39496958c"
  },
  "body": {
    "param": "GtmyJSC9Dk/J6UPJgyxFwET/2LUubXJ2I6ecBemhCiMIyz0jsJFp8eINaJ219ILdNG2/Y0U6pSn/OuNkk1N75cdvCqooXiHqqSmgndpu3JriqYj3u6wdJOG5AyEUeP4niH+eSGf8yyGQzh+s0tPRre45oCiMVzV8WR5DOQTO9SdIC1dJJZnX4liTN4J02pK3P/O2j/fvhuBlR3TKOOpDgA==",
    "secret": true
  }
}
```

#### param

```json
{
  "spaceType": "0",
  "familyId": "9263097",
  "parentDirectoryId": "b2e9b430ed02493782065106291b0308",
  "directoryName": "inner_folder",
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "bb7a573aef8e4c7185754593cb6d7682",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "CnnZck3ubj44rwSM3AM/v9zjE3ysakdbmIKE/VaoNe/rnV62Yhda97WP0JQM5174"
  }
}
```

#### data

```json
{ "id": "1bb8cfb12d7f4d4eaac1b3a5293a256f" }
```

# RenameFileOrDirectory

### Req

```json
{
  "header": {
    "key": "RenameFileOrDirectory",
    "resTime": 1686144095914,
    "reqSeq": 119050,
    "channel": "wohome",
    "version": "",
    "sign": "15b981684108b2a927d6c268775d612a"
  },
  "body": {
    "param": "GtmyJSC9Dk/J6UPJgyxFwGp9iZTS3ITlNospQy+ACXq+x9bYwe5f2MKIoduy27GePOHiZcXoWU38TUoIY726fhnp4tREw48TNnjh9ltV8w+gNe1lcQtZK2ffLata84GVGf193PA2F76eyDXnT77S79cqXqmUZw+aebXAUBQ8eeo=",
    "secret": true
  }
}
```

#### param

```json
{
  "spaceType": "0",
  "type": 1,
  "fileType": "1",
  "id": "78ee3248f33941b4815199f1b6daaab7",
  "name": "new_name.png",
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "e3d45f94fc694d94b08aacca109e1352",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": ""
  }
}
```

# MoveFile

### Req

```json
{
  "header": {
    "key": "MoveFile",
    "resTime": 1686144313361,
    "reqSeq": 150760,
    "channel": "wohome",
    "version": "",
    "sign": "0585b1b888e64355279f605a2cd1d0aa"
  },
  "body": {
    "param": "ihM2idvmmh1dINmyiXmuLf0cruAdowEkbijMDGjuAbDKdwtrPqs4DXXyrAZn76zKhwpqWdJYnXxDEJzf9v/suwR0B1d6OJWUDhbhn/N9yrQX5iH3TOlkZnrOvQNo60M1KZe7/4wllIT8hfjBGTojHvw+l+FTIzQYjLf/OOlhxgKOxUzv2Rw2TQGfroVfeVHcS9ronpbSi4apmasygAPJroBLNKWm84rN5GBjwbdBZFiK6YiuUxlMXNZH6bCxq+Ex",
    "secret": true
  }
}
```

#### param

```json
{
  "targetDirId": "57ad13bb15cf4e6e94799c330d06e9f6",
  "sourceType": "0",
  "targetType": "0",
  "dirList": [],
  "fileList": ["78ee3248f33941b4815199f1b6daaab7"],
  "secret": false,
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "791fae44a0814aa3ae10e59094bfbbad",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": ""
  }
}
```

# CopyFile

### Req

```json
{
  "header": {
    "key": "CopyFile",
    "resTime": 1686144417113,
    "reqSeq": 155691,
    "channel": "wohome",
    "version": "",
    "sign": "d1f88b0de379c41bb82b971df754c025"
  },
  "body": {
    "param": "ihM2idvmmh1dINmyiXmuLXf0xChFSz2aX4ZYl/+UFTdb6Y8/M2mKYKo8OWRWAVotAPsIwR0BnoRXCdkO/7HiZOb171dLHHbaCcSyiskFrjxrk8U/w3BxawvwppCLtnO70mZKzw7adrqRnOMWwl2vhaRlXJjcslaMSuyXypUZNvdIDyCKROIOkNo6zUW+zqJnK9tk6KuvhVuX6jJxshBpMWv8Ly0gSUGB2310m76bhSsE1K79wRQGNliNpuYISTiISaCsIwik6X34vYntPAS8Qw==",
    "secret": true
  }
}
```

#### param

```json
{
  "targetDirId": "8fa10be30f5e41b49a6abb516f5c7649",
  "sourceType": "0",
  "targetType": "0",
  "dirList": [],
  "fileList": ["78ee3248f33941b4815199f1b6daaab7"],
  "familyId": "9263097",
  "secret": false,
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "396bca037ea04e8abe323237127132cc",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": ""
  }
}
```

# DeleteFile

### Req

```json
{
  "header": {
    "key": "DeleteFile",
    "resTime": 1686144630496,
    "reqSeq": 124076,
    "channel": "wohome",
    "version": "",
    "sign": "5387bfadbfdd13f56a1500c5482f0160"
  },
  "body": {
    "param": "GtmyJSC9Dk/J6UPJgyxFwMkgZZ7g5CkKAWC7kMnqnuOlfmc9GlKAtW0IDhM9/rTRwlUBuV1Ip72yNUVG0wbQ+kxL64qNF0z68IvaiCTf9X2z1QIkKtANyfa3VmzDEWJ11FDKZpsgEvU54aUjfOtyDbg1E9r1kZoR+LNuNVQ1/mI=",
    "secret": true
  }
}
```

#### param

```json
{
  "spaceType": "0",
  "vipLevel": "0",
  "dirList": ["70431a1e535b405f956b6e451db9255d"],
  "fileList": [],
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "f4872ff17b4149668b702c8ade591f29",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": ""
  }
}
```

# EmptyRecycleData

### Req

```json
{
  "header": {
    "key": "EmptyRecycleData",
    "resTime": 1686144735505,
    "reqSeq": 180247,
    "channel": "wohome",
    "version": "",
    "sign": "9d599a6465923f91fb14965a2a031150"
  },
  "body": {
    "param": "/V0Cb8bUEmLHNGFHDLHFgyE/yfwzzLO7mIejUNl3qjw=",
    "secret": true
  }
}
```

#### param

```json
{ "clientId": "1001000021" }
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "aebd244d7a0e4ec28a1cfecf42209c4e",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": ""
  }
}
```

# Upload2C

url: https://gxupload.pan.wo.cn:8443/openapi/client/upload2C

### Req

```yaml
uniqueId: 1686143710632
accessToken: c3f0362b-xxxx-4997-a9c2-0478702a5fea
fileName: avatar.png
psToken: undefined
fileSize: 70207
totalPart: 1
partSize: 70207
partIndex: 1
channel: wocloud
directoryId: b2e9b430ed02493782065106291b0308
fileInfo: GtmyJSC9Dk/J6UPJgyxFwO4ig/Yz4Qqf4p4+Upq21W9uzbXB73q6PG9Hq5NmHteGtIvtRsZSIXUGTGYw9UfUfkH3+xUmGD5vkvKg0/0OMlt0GDqztiTuWIEyvSwytVKfyKfS7ul4VDLoYWoki3qGxaprTrj5rc9ZI5VIIA0sYzojBEOXyh1042WCoDz8w+CFSXV9f3/aQOYsfVt7BxxkPA==
file: (二进制)
```

#### fileInfo

```json
{
  "spaceType": "0",
  "directoryId": "b2e9b430ed02493782065106291b0308",
  "batchNo": "20230607211507",
  "fileName": "avatar.png",
  "fileSize": 70207,
  "fileType": "1"
}
```

### Resp

```json
{
  "code": "0000",
  "data": {
    "fid": "8S7lz_d/kEVpB7Or7Lu79/3Cu7SDDMovQixRedpc86A1W464MPIUCdwrpPDuXLfjpbFEnP"
  },
  "msg": "上传成功"
}
```
