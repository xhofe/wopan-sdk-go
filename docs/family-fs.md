# QueryAllFiles (root folder)

### Req

```json
{
  "header": {
    "key": "QueryAllFiles",
    "resTime": 1686144942579,
    "reqSeq": 169687,
    "channel": "wohome",
    "version": "",
    "sign": "cf5eec2fa2799142aafa3d29ca22d23c"
  },
  "body": {
    "param": "RUHfQlrlkoNf8oVK8ofLwWAN/hnDzfaqiCKNuI8CQA41vYgavLuPF7+dPExXU5Mcho0xqtcIPpaeFVynnAfEoV7kZ0xRG2thTc6XnNXhHhroohXtVDFAUUvz2MkPySnB5XLzUeq+7zQ6y0JUWd+4hB96Ro4LXTUYph6MK2obLvU=",
    "secret": true
  }
}
```

#### param

```json
{
  "spaceType": "1",
  "parentDirectoryId": "0",
  "pageNum": 0,
  "pageSize": 50,
  "sortRule": 0,
  "familyId": "9263097",
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "32560d6cb1ec483f89f16b42b660820e",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "b+fh7Uuv5WqECmt1B++feofiJE8QlaGAN0GEFe9UYlLgLetBsq0Vq3bfPJQlT89h/Nzfye/UwN/gOSEMS/GYUw0NVilHlz+gFRW5x0rtpsdqjcIF7B6qQnE8mYy2oLrw3bwDOdLW9Wf3TfI5vK5dlShcJQmAYwB01eAGF+o4/iBePAz4l9i+u1liecCeuMoKU14acpaD2+zGnYUkj3j8eVahTwEe1i4I1IxrbfFZ+4U6mr8YV6L/wpd/9DU8cwAHZxNJNE0J4rYvw31k/GS34sK2bjIOVCljpmbTqdeODeY="
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
      "createTime": "20220221185510",
      "name": "test",
      "shootingTime": "",
      "id": "37135bfb5178467b857129a00931e24a",
      "type": 0,
      "thumbUrl": "",
      "fileType": "0"
    }
  ]
}
```

# QueryAllFiles (inner folder)

### Req

```json
{
  "header": {
    "key": "QueryAllFiles",
    "resTime": 1686145435020,
    "reqSeq": 102048,
    "channel": "wohome",
    "version": "",
    "sign": "db0eed7643a88f3826221dc2538c0171"
  },
  "body": {
    "param": "RUHfQlrlkoNf8oVK8ofLwWAN/hnDzfaqiCKNuI8CQA4i0oAb79adQXcUTbuil6axOPe+4PeZspE0LCXIRMT3yM1pv/LXrKVuteMkY6+RPDcmHiWmYc7JWx9hd5DsqxVH+PAWryHzBCUPkALFMhdST9Gtm5Rg3r5k86lnivIfUouGOffXFAKDhvSXoM+cDGNUt546xS8mbkNSIPXKvAf5gQ==",
    "secret": true
  }
}
```

#### param

```json
{
  "spaceType": "1",
  "parentDirectoryId": "37135bfb5178467b857129a00931e24a",
  "pageNum": 0,
  "pageSize": 50,
  "sortRule": 0,
  "familyId": "9263097",
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "82b027aff46149d3922218116ddde18d",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "b+fh7Uuv5WqECmt1B++feofiJE8QlaGAN0GEFe9UYlLgLetBsq0Vq3bfPJQlT89h/Nzfye/UwN/gOSEMS/GYUw0NVilHlz+gFRW5x0rtpsd17aAFj+Dw1YhLzJyGj5ryctFZ2UpX8crPdAcG+mhz2l7/ph2CRMEbsdanIIW6Bz+7g+Fxbbv8bT9+bWflIHS/NDQ2RkZhd04jwR0uLumUzk7mpDH0ouWaor5Zp0wrqDcdWAlTbcYm3WlEXRqWtHtNzk5aIDXR4yvRyVEWVqIlzs7pYFntDmduQt6Wu3+UYYKb8Dk64/glx8tC0lIIwSVusMf0nEmXDKBcpmFmnmj8SKm9Ab/YbRXInOAJG5f+2s3/BhOlyqFq4Me0XKKru1xk2xNRfJ9BSpiPxXyF+fc97uCtgW1YeATAp/YNkdVrmb4p6lkEhDk8jtekLXQ7R/K2rnNFAt5JqP7cdo3tcQBLcwO27QqwKJGiigCHQ+GLdoGrAg/r/zqM9+oW237DsAc/PSJQHN3QFqbeK7OvLAn5fzfHDvNFIwqV5V7vJ/Yu5g0WUkgXJIZgO9V24djGR0bw8s1JhGxzVW6P803op0FhPNFoRRA9vD0X/EeQH4raqqRDM0a3HpuvcXZlhwdHn5KtsQ+1X/oN2R8wi9vHBa/qbEtDCWSjcMKbZWg0aeg+zoCpT0Dw8Ez4nWiRTaEzB0EbCKM8lI1t9EM578M2go8nj/Ci3BAii46Hd69a33Yw5FekBEtvHbPcYTLCuKFvpK55hD0/nSpsQN2VodI9Mt5ddZLoTvwEvtkSRj/N1PG8uyfKTvgIDw8aZ7FfyQkUnOhY"
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
      "createTime": "20230607214351",
      "name": "inner_folder",
      "shootingTime": "",
      "id": "b8aa81b8d90c42e4b9e4109d9c8d2de5",
      "type": 0,
      "thumbUrl": "",
      "fileType": "0"
    },
    {
      "familyId": 9263097,
      "fid": "98+nz4OoVXNP9fuxk4BpBrtC9COF3ec21y+nNvqPTIPcBlbZe7uG7gZd3yx2Lu4F",
      "creator": "186****5244",
      "size": 1599939,
      "createTime": "20220221185522",
      "name": "bg.png",
      "shootingTime": "",
      "id": "983ee503fe9a4aadaf3daf2333a46b75",
      "type": 1,
      "thumbUrl": "https://du.smartont.net:8442/openapi/thumbnails?fid=98%2Bnz4OoVXNP9fuxk4BpBrtC9COF3ec21y%2BnNvqPTIPcBlbZe7uG7gZd3yx2Lu4F",
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
    "resTime": 1686145434904,
    "reqSeq": 131488,
    "channel": "wohome",
    "version": "",
    "sign": "29f4b35dde6d9e7eaff8c4e64005cd30"
  },
  "body": {
    "param": "doFCig8jnOT0I8MM0nYaJoFE+fOsoa2jz7353K/EvcicLO2we/6BJySOdhQE1xqVNhyfVEWeUjT69ltCvIRoOUSmeXtvEWIJvISsjKD3Hfs=",
    "secret": true
  }
}
```

#### param

```json
{ "directoryId": "37135bfb5178467b857129a00931e24a", "clientId": "1001000021" }
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
    "resTime": 1686145612714,
    "reqSeq": 154268,
    "channel": "wohome",
    "version": "",
    "sign": "fff4b0d818db261d43500b3a992f79b9"
  },
  "body": {
    "param": "DVt84prwHvHwQf1sM3V+C3IQXaT0FZoUciNfQkrxXoh8AMgdNK/6P5wESokUXDNYhgcpFxC2LCz8SbTsTsytAm3o/N4b2szcLe8sCvkwXLAyzgZF614B/syXT2zx58AwpcYfV5Q0/QX8qDDmAKxaQViXcDCWb9BUyNVWHLj2jzk=",
    "secret": true
  }
}
```

#### param

```json
{
  "type": "1",
  "fidList": [
    "98+nz4OoVXNP9fuxk4BpBrtC9COF3ec21y+nNvqPTIPcBlbZe7uG7gZd3yx2Lu4F"
  ],
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "6f9d8b4aa0824ec7b6086659bb4c8d5d",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "Yo2tB2COSjTseVRJ85LFn/18inrbX3t2HTSwFkURrBbg7/4hjznX9eIPjfVzDMJrEe+AudHxe7DHOKylI08oQXwcdViwGvi9Qq6ye+x0OaGUh3NhlVnUIclSQ2rwonNNza7CE+3CBmWgD9468BALyfGmyS4kYMJvyvg12thOHJX8gDpubBveTYlIq7C9gJUw1EZviuqYiu4kEjva0gbgebprtdfsTTJ+Lc+wyxO8aH74N0wPH6YNbIdU1Jtmh7M4WzvBKj3aXZfiXirEoErVcVFEj+5ydeAYFBR4e4y0vBzWW2HSw2YFMugC8FO3QocLFbhafx3XR25c984pf3aNGwWh87vJQS+v0LIQ7NHsmV0="
  }
}
```

#### data

```json
{
  "type": 1,
  "list": [
    {
      "fid": "98+nz4OoVXNP9fuxk4BpBrtC9COF3ec21y+nNvqPTIPcBlbZe7uG7gZd3yx2Lu4F",
      "downloadUrl": "https://du.smartont.net:8445/openapi/download?fid=9itmN_WiLhDRhDV2HK%2BJJXTvkdMpfAhC/rnPJU36IHW3toPndT45CwkaUUyB1mCDamGPweVC4ks7hiwjyRA9nfLDxOeg=="
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
    "resTime": 1686145789995,
    "reqSeq": 179986,
    "channel": "wohome",
    "version": "",
    "sign": "d7b872fccfc60d82e5a5706d698fc27e"
  },
  "body": {
    "param": "GyTtewjGpzFnNUxukTan2KnPv/Jo0np23JXyoBCKtaiz5TEsCv2pJsFCyv/OtewTOJ/nZhD/SOf1kxmx40wZYgiDbM/bKB1gdeRCYhXsXAEWLY3aVc9xlKUhiQIbCq290+mvjxS/fcF5p6Ikk+ZHhaxV2FkduSxTbtaW4LV7Q0s=",
    "secret": true
  }
}
```

#### param

```json
{
  "fidList": [
    "98+nz4OoVXNP9fuxk4BpBrtC9COF3ec21y+nNvqPTIPcBlbZe7uG7gZd3yx2Lu4F"
  ],
  "clientId": "1001000001",
  "spaceType": "1"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "3135a1c13ca64453a0b2c3a579978fdb",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "XtyvFKrfkLeUz/zso24J8vUlugBYjq3MJvczJm2+9pFVP0JUjHurP1ogTwstVL7ce9rsJ1fqS6k9YNj1AlLf56O2oicVOSFJpQm/G2IIDecfR4JTkvLjaEWoadBgaBZWBxt6Z3OF2D5+mVynkxuCw745yd8H/mDtYsH887NL+3hpQPTq3d7/PPVLpzVRp1B4pZOlkFxc3jhxH36FIYz4pP85KTq7/BngURoQ/tYIe2DCZ4dPTD/xWj338++n7daCzUNR6l+3q03fwbPVFfgsapCm99OYY4XKyPFcFk8TqMaFbZL9+DG5btktvdxucaKX7+B8EMkIBUILdXHO8XiQRA=="
  }
}
```

#### data

```json
[
  {
    "fid": "98+nz4OoVXNP9fuxk4BpBrtC9COF3ec21y+nNvqPTIPcBlbZe7uG7gZd3yx2Lu4F",
    "downloadUrl": "https://du.smartont.net:8445/openapi/download?fid=m4h0x_6vvZ6gwFt6939Ofxli%2BPAWddz3GsbK5oDqTR2zAt4ebowv5y6q8KA2Cn8lHevcHTwMXXHi%2BWmeMG8uDCUa0QqA=="
  }
]
```

# CreateDirectory

### Req

```json
{
  "header": {
    "key": "CreateDirectory",
    "resTime": 1686146328811,
    "reqSeq": 147094,
    "channel": "wohome",
    "version": "",
    "sign": "30ff0c4025a649750e8f75dbbb60297d"
  },
  "body": {
    "param": "RUHfQlrlkoNf8oVK8ofLwSILlW0vUpVHDISNlKLj6qV/q0p/OF4xX+Mv4peyNHLVQ/V8UTHHzr+kBekurHSQEooKIr3hSjUFVXNJ/LGFEL+WaBSJ1uMHU8TCEpisNOqfwPnWBRhG8lLF2OYFh6roF2gsM4vt7Ilk4fulke76SGk=",
    "secret": true
  }
}
```

#### param

```json
{
  "spaceType": "1",
  "familyId": "9263097",
  "parentDirectoryId": "0",
  "directoryName": "test_mkdir",
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "22cf6f29acc6469083b1b197e81369a9",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "olqWpzdKgTqTy0y/tsPiVNws2eDkiWDEfwzppQBNcVbmzdAFJFV3BqmGImrSxuEm"
  }
}
```

#### data

```json
{ "id": "ab8147dba859439abef9edec0987341b" }
```

# RenameFileOrDirectory

### Req

```json
{
  "header": {
    "key": "RenameFileOrDirectory",
    "resTime": 1686146447017,
    "reqSeq": 107355,
    "channel": "wohome",
    "version": "",
    "sign": "6c7130af16e0a0820fee1e369cc29d52"
  },
  "body": {
    "param": "RUHfQlrlkoNf8oVK8ofLwUfYXU/GWFZguMSyGssZA0OmY4Zf06frWmYn1UEOrTKDCNAyrEc28cCVFux/fzutFZayJBouD2gpS30/+5V3jg/+G1zkVBvBdPtzikUd5Dg90ZD3B9Q1qW8GkAmxCu+EAuN4QIZvBry5tyFJ3SqW4LMfLxotvoRjk2cjxHEYeIz9t+zK4AFzNI6yeDc58QXrMg==",
    "secret": true
  }
}
```

#### param

```json
{
  "spaceType": "1",
  "type": 0,
  "fileType": "0",
  "id": "ab8147dba859439abef9edec0987341b",
  "name": "rename_newname",
  "familyId": "9263097",
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "3403664275654279b3df6f91964452f0",
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
    "resTime": 1686146616560,
    "reqSeq": 141296,
    "channel": "wohome",
    "version": "",
    "sign": "ec4208bcbeaaccd09f4693293d523345"
  },
  "body": {
    "param": "ihM2idvmmh1dINmyiXmuLaM2YtQjHZsjKKxq4MJFwmw/ziVHmnVinnW4EZ/OUDC5ljxtT0ryFSxFGO1qpfXidznvEFR11vkTE6uy4W44LGawiqS1oFAQUWzuP1YtBbKORg+9ZJfgeB/22pKSnICTx1nWI5yaqKCXxaURHAEeiUNWKcE/8eXipMdzjiiMIKESaztKCIIM3Z3JQwxMqqe9m6MwzAA+MDEKj49qfp8RS35WjHu6mCEm1WoUSOnXfmz+ku5iZYULKXfET5NIbcgw9sfomx21mEP+r4DthHnw7HzbELE298hexyG2lv5zebID",
    "secret": true
  }
}
```

#### param

```json
{
  "targetDirId": "37135bfb5178467b857129a00931e24a",
  "sourceType": "1",
  "targetType": "1",
  "dirList": ["ab8147dba859439abef9edec0987341b"],
  "fileList": [],
  "familyId": "9263097",
  "fromFamilyId": "9263097",
  "secret": false,
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "2bb8990c816a4e16a829b699f5895c3a",
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
    "resTime": 1686146728802,
    "reqSeq": 159201,
    "channel": "wohome",
    "version": "",
    "sign": "00f0c43ff6bb93b733a3d8b7d18650c3"
  },
  "body": {
    "param": "ihM2idvmmh1dINmyiXmuLXN6+44xAmOmta/HiyiVNrO11HEyVshG05dnDvJnDvczAbJq12bndqNlQ/cPEQyeV5FG4+pG2Tdn1HSx0IovXG4pbHXbdgdDBuhpi03Dzwa6p6qvHE0Y9dIRq/dhmrlnTTicT6XYTIdCWJ9GKPB47lw2wxlQaSUytVL4BzCXN7T9By0Ijw9Su37bZtNwryEjuZi4WKJHsyrk2f6wCynjg/i99DpDwdFh1BBq3+ukHomIp2fN8M8gVTeRjdIZ7lA3xQ==",
    "secret": true
  }
}
```

#### param

```json
{
  "targetDirId": "ab8147dba859439abef9edec0987341b",
  "sourceType": "1",
  "targetType": "1",
  "dirList": ["b8aa81b8d90c42e4b9e4109d9c8d2de5"],
  "fileList": [],
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
  "LOGID": "ab7cf18a850c4da4a54af55f959e9d22",
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
    "resTime": 1686146818168,
    "reqSeq": 147076,
    "channel": "wohome",
    "version": "",
    "sign": "9fef701045d20ffe828626c1dee5cb85"
  },
  "body": {
    "param": "RUHfQlrlkoNf8oVK8ofLwV+mAYTdNZM63Z+XVO4+FnzTev0VfrIBkVNCFh/bUyXvdYH/cQr8xUkGCK8d2ah5KN6AMTFrIjbH4ZsWi2pNshLuvrZGZejcfKWc5o4GVTd2nwlyy03ZG71+pBZQnrzuGxrxa+j/DXl2h+/REMNz3AY=",
    "secret": true
  }
}
```

#### param

```json
{
  "spaceType": "1",
  "vipLevel": "0",
  "dirList": ["ab8147dba859439abef9edec0987341b"],
  "fileList": [],
  "clientId": "1001000021"
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "7c063be9dc154854b55747e81a567573",
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
uniqueId: 1686146963538
accessToken: c3f0362b-xxxx-4997-a9c2-0478702a5fea
fileName: Termius.exe
psToken: undefined
fileSize: 158384584
totalPart: 18
partSize: 8388608
partIndex: 7
channel: wocloud
directoryId: 37135bfb5178467b857129a00931e24a
fileInfo: RUHfQlrlkoNf8oVK8ofLwY0XYfmeLL5azg3jvGYjwTssmT9dt6K/lQaQoGYHLgrmPTUWQFqydYMI5DjdwTGuba0/4BCBNg47TQ9sgYKc5c6zrBtMNZ5+cfdOyw6Go3mD9DjcQj+I7QzSAEkhUu/UvmWv0o1gL73IB5fzYLfto57p1qDMHZscTyH5k/lekwVz/03F22CVDXH5emAG0xKh7eaSLe+l8AVx4mU+y/mr6rw=
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
    "fid": "73VcF_d/kEVpB7Or7Lu79/3Cu7SL1cPQOq8vtK7fhxtejmguBEUuq6AF1ee93IbAhKmPiU"
  },
  "msg": "上传成功"
}
```
