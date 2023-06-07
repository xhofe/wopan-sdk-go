# FCloudProductOrdListQry

### Req

```json
{
  "header": {
    "key": "FCloudProductOrdListQry",
    "resTime": 1686118434324,
    "reqSeq": 114879,
    "channel": "wohome",
    "version": "",
    "sign": "ae8fa92b643f9b5e885168b6d8b51c3e"
  },
  "body": {
    "qryType": "1",
    "clientId": "1001000021"
  }
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "d5ea7ac0656d4e6399feec736be73be0",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": {
      "fcloudProductOrds": [
        {
          "activiteCode": "",
          "appStorePackageDesc": "",
          "appStorePackageFee": "",
          "appStoreProductId": "",
          "applyTime": "20210906192439",
          "applyTimeFormate": "2021.09.06",
          "cbssOrderId": "30202109061924398853329735874469",
          "city": "710",
          "clientId": "",
          "days": "",
          "descUrl": "",
          "effectState": "1",
          "effectiveDays": 2399,
          "expireTime": "20230630235959",
          "expireTimeFormate": "2029.12.31",
          "fee": "",
          "isAppStorePay": "",
          "isAutoSub": "0",
          "isExpire": "0",
          "isNewPackage": "0",
          "isOnline": "1",
          "isPlus": "0",
          "isShowExpireTips": "0",
          "orderId": "",
          "orderState": "0",
          "orderStatus": "02",
          "packageDesc": "0",
          "packageProductCode": "20200217160000",
          "packageProductId": "20200217160000",
          "payMethod": "",
          "payTransactionId": "",
          "payType": "0",
          "province": "71",
          "remainDays": "0",
          "signStatus": "",
          "source": "6",
          "subTime": "20210906192254",
          "subTimeFormate": "2021.09.06",
          "subType": "",
          "userId": "186****7660",
          "vipDesc": "普通会员",
          "vipDescNew": "普通会员",
          "vipExpireTimeLabel": "永久",
          "vipLevel": "0"
        },
        {
          "activiteCode": "",
          "appStorePackageDesc": "",
          "appStorePackageFee": "",
          "appStoreProductId": "",
          "applyTime": "20210906192439",
          "applyTimeFormate": "2021.09.06",
          "cbssOrderId": "30202109061924398853329735874469",
          "city": "710",
          "clientId": "",
          "days": "",
          "descUrl": "",
          "effectState": "1",
          "effectiveDays": 2399,
          "expireTime": "20230630235959",
          "expireTimeFormate": "2029.12.31",
          "fee": "",
          "isAppStorePay": "",
          "isAutoSub": "0",
          "isExpire": "0",
          "isNewPackage": "0",
          "isOnline": "1",
          "isPlus": "0",
          "isShowExpireTips": "0",
          "orderId": "",
          "orderState": "0",
          "orderStatus": "02",
          "packageDesc": "0",
          "packageProductCode": "20200217160000",
          "packageProductId": "20200217160000",
          "payMethod": "",
          "payTransactionId": "",
          "payType": "0",
          "province": "71",
          "remainDays": "0",
          "signStatus": "",
          "source": "6",
          "subTime": "20210906192254",
          "subTimeFormate": "2021.09.06",
          "subType": "",
          "userId": "186****7660",
          "vipDesc": "普通会员",
          "vipDescNew": "普通会员",
          "vipExpireTimeLabel": "永久",
          "vipLevel": "0"
        }
      ],
      "maxVipLevel": "3",
      "isShowInlet": "0"
    }
  }
}
```

# QueryCloudUsageInfo

### Req

```json
{
  "header": {
    "key": "QueryCloudUsageInfo",
    "resTime": 1686118434324,
    "reqSeq": 188014,
    "channel": "wohome",
    "version": "",
    "sign": "4b12ba49e760f16373c5e354a07f0685"
  },
  "body": {
    "param": "1GuP7b4Bf/aS/nPfM56MQFw54IzHa6dBnWWkDQv7skksWX71heudEwrh8NgqAKBnr1RDqP9EBs4XD6zLV5eogg==",
    "secret": true
  }
}
```

#### param

```json
{ "phoneNum": "186****5244", "clientId": "1001000021" }
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "efac4d4b15a9480bab47907e2096700f",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "I02moQlq44kMrImZhftqX2yYleUrCSZXdZ0H9+dKBEUSr+ekwzx7+e21mNJHrnGBI+WhXaoA41KfPsJqV6HSGdpsj16AtEfxS79nCh1YLATlgLuTUkgN/LygywDPmKE6kXQNMIbCFi5W2wlsXsZAjyLtJGPOOa3RexOGajInHwNLgwep8uazK4fnvYThiuL4Qh2KLQlivNMbapFZTiFOKzz6UfvD+RtWYkU1Pc2E2AgNsBnAM13EaBL/o3bdMY2A3+mCG5chN27gtmigfqjsFMCWzgv2q4CAvgFC9pykac1jCNenIfL1hQOGJKGP58HUpstbGZFa10XK20NY2opaqllsl+ojcyxnRTAGx9/JzOymBFU3j4N22IodRh0JxmNWCWeYsckRla5EpRA8sksNMF1sUwgYvBXcFTqiVYYRFpp0v4mxpjB4aEmt346J7WDz"
  }
}
```

#### data

```json
{
  "code": "0000",
  "usageInfo": {
    "totalSize": "10485760",
    "usedSize": 3124,
    "imageSize": 3124,
    "videoSize": 0,
    "audioSize": 0,
    "textSize": 0,
    "otherSize": 0,
    "byteUsedSize": 3199878,
    "byteTotalSize": "10737418240"
  },
  "vipLevel": "0",
  "expireTime": "20291231235959",
  "applyTime": "20210906192439",
  "payType": "0",
  "source": "6",
  "orderState": "0",
  "status": "1"
}
```

# FCloudProductPackage

### Req

```json
{
  "header": {
    "key": "FCloudProductPackage",
    "resTime": 1686118434324,
    "reqSeq": 123983,
    "channel": "wohome",
    "version": "",
    "sign": "898ce3b15fa95ab3de6d32785b2b4c99"
  },
  "body": {
    "param": {
      "clientId": "1001000021",
      "vipLevel": "0"
    },
    "clientId": "1001000021"
  }
}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "cc74b16ca9b44662a24b93eba786376f",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": [
      {
        "vipLevel": "0",
        "vipDesc": "普通会员",
        "familyCount": "10",
        "recycleBin": "7",
        "batchUploadCount": "100",
        "capacity": "10240",
        "uploadFileSize": "1024",
        "memberCount": "20",
        "transientCount": "50",
        "vipWeight": "0",
        "packages": [
          {
            "appStoreContractProductId": "",
            "appStoreDesc": "",
            "appStorePrice": "",
            "appStoreProductId": "",
            "applyTime": "2023.06.07",
            "batchUploadCount": "100",
            "capacity": "10240",
            "days": "0",
            "desc": "0",
            "descPic": "",
            "expireTime": "2023.06.07",
            "familyCount": "10",
            "label": "",
            "memberCount": "20",
            "months": "",
            "order": 0,
            "originalPrice": "",
            "payType": "0",
            "price": "0",
            "productCode": "20200217160000",
            "productId": "20200217160000",
            "recycleBin": "7",
            "tips": "",
            "transientCount": "50",
            "uploadFileSize": "1024",
            "vipDesc": "普通会员",
            "vipLevel": "0",
            "vipWeight": "0"
          }
        ]
      },
      {
        "vipLevel": "4",
        "vipDesc": "乐享会员",
        "familyCount": "10",
        "recycleBin": "7",
        "batchUploadCount": "100",
        "capacity": "102400",
        "uploadFileSize": "1024",
        "memberCount": "20",
        "transientCount": "50",
        "vipWeight": "1000",
        "packages": [
          {
            "appStoreContractProductId": "",
            "appStoreDesc": "",
            "appStorePrice": "",
            "appStoreProductId": "",
            "applyTime": "2023.06.07",
            "batchUploadCount": "100",
            "capacity": "102400",
            "days": "30",
            "desc": "2元/月",
            "descPic": "",
            "expireTime": "2023.07.07",
            "familyCount": "10",
            "label": "",
            "memberCount": "20",
            "months": "",
            "order": 0,
            "originalPrice": "",
            "payType": "0",
            "price": "2",
            "productCode": "lxcloud01",
            "productId": "lxcloud01",
            "recycleBin": "7",
            "tips": "",
            "transientCount": "50",
            "uploadFileSize": "1024",
            "vipDesc": "乐享会员",
            "vipLevel": "4",
            "vipWeight": "1000"
          }
        ]
      },
      {
        "vipLevel": "5",
        "vipDesc": "畅享会员",
        "familyCount": "10",
        "recycleBin": "7",
        "batchUploadCount": "100",
        "capacity": "512000",
        "uploadFileSize": "1024",
        "memberCount": "20",
        "transientCount": "50",
        "vipWeight": "2000",
        "packages": [
          {
            "appStoreContractProductId": "",
            "appStoreDesc": "",
            "appStorePrice": "",
            "appStoreProductId": "",
            "applyTime": "2023.06.07",
            "batchUploadCount": "100",
            "capacity": "512000",
            "days": "30",
            "desc": "6元/月",
            "descPic": "",
            "expireTime": "2023.07.07",
            "familyCount": "10",
            "label": "",
            "memberCount": "20",
            "months": "",
            "order": 0,
            "originalPrice": "",
            "payType": "0",
            "price": "6",
            "productCode": "cxcloud01",
            "productId": "cxcloud01",
            "recycleBin": "7",
            "tips": "",
            "transientCount": "50",
            "uploadFileSize": "1024",
            "vipDesc": "畅享会员",
            "vipLevel": "5",
            "vipWeight": "2000"
          }
        ]
      },
      {
        "vipLevel": "1",
        "vipDesc": "尊享会员",
        "familyCount": "30",
        "recycleBin": "15",
        "batchUploadCount": "1000",
        "capacity": "2097152",
        "uploadFileSize": "5120",
        "memberCount": "30",
        "transientCount": "500",
        "vipWeight": "3000",
        "packages": [
          {
            "appStoreContractProductId": "ct90780130",
            "appStoreDesc": "12元/月",
            "appStorePrice": "12",
            "appStoreProductId": "90780130",
            "applyTime": "2023.06.07",
            "batchUploadCount": "1000",
            "capacity": "2097152",
            "days": "30",
            "desc": "10元/月",
            "descPic": "",
            "expireTime": "2023.07.07",
            "familyCount": "30",
            "label": "超值",
            "memberCount": "30",
            "months": "",
            "order": 1,
            "originalPrice": "",
            "payType": "0",
            "price": "10",
            "productCode": "91027196",
            "productId": "91027196",
            "recycleBin": "15",
            "tips": "",
            "transientCount": "500",
            "uploadFileSize": "5120",
            "vipDesc": "尊享会员",
            "vipLevel": "1",
            "vipWeight": "3000"
          },
          {
            "appStoreContractProductId": "ct91027219",
            "appStoreDesc": "30元/季",
            "appStorePrice": "30",
            "appStoreProductId": "91027219",
            "applyTime": "2023.06.07",
            "batchUploadCount": "1000",
            "capacity": "2097152",
            "days": "90",
            "desc": "30元/季",
            "descPic": "",
            "expireTime": "2023.09.05",
            "familyCount": "30",
            "label": "超值",
            "memberCount": "30",
            "months": "",
            "order": 2,
            "originalPrice": "",
            "payType": "3",
            "price": "30",
            "productCode": "91027219",
            "productId": "91027219",
            "recycleBin": "15",
            "tips": "",
            "transientCount": "500",
            "uploadFileSize": "5120",
            "vipDesc": "尊享会员",
            "vipLevel": "1",
            "vipWeight": "3000"
          },
          {
            "appStoreContractProductId": "ct91027224",
            "appStoreDesc": "123元/年",
            "appStorePrice": "123",
            "appStoreProductId": "91027224",
            "applyTime": "2023.06.07",
            "batchUploadCount": "1000",
            "capacity": "2097152",
            "days": "360",
            "desc": "120元/年",
            "descPic": "",
            "expireTime": "2024.06.01",
            "familyCount": "30",
            "label": "超值",
            "memberCount": "30",
            "months": "",
            "order": 3,
            "originalPrice": "",
            "payType": "1",
            "price": "120",
            "productCode": "91027224",
            "productId": "91027224",
            "recycleBin": "15",
            "tips": "",
            "transientCount": "500",
            "uploadFileSize": "5120",
            "vipDesc": "尊享会员",
            "vipLevel": "1",
            "vipWeight": "3000"
          }
        ]
      },
      {
        "vipLevel": "6",
        "vipDesc": "白银会员",
        "familyCount": "30",
        "recycleBin": "15",
        "batchUploadCount": "1000",
        "capacity": "4194304",
        "uploadFileSize": "5120",
        "memberCount": "30",
        "transientCount": "500",
        "vipWeight": "4000",
        "packages": [
          {
            "appStoreContractProductId": "",
            "appStoreDesc": "",
            "appStorePrice": "",
            "appStoreProductId": "",
            "applyTime": "2023.06.07",
            "batchUploadCount": "1000",
            "capacity": "4194304",
            "days": "30",
            "desc": "15元/月",
            "descPic": "",
            "expireTime": "2023.07.07",
            "familyCount": "30",
            "label": "",
            "memberCount": "30",
            "months": "",
            "order": 1,
            "originalPrice": "",
            "payType": "0",
            "price": "15",
            "productCode": "bycloud01",
            "productId": "bycloud01",
            "recycleBin": "15",
            "tips": "",
            "transientCount": "500",
            "uploadFileSize": "5120",
            "vipDesc": "白银会员",
            "vipLevel": "6",
            "vipWeight": "4000"
          },
          {
            "appStoreContractProductId": "",
            "appStoreDesc": "",
            "appStorePrice": "",
            "appStoreProductId": "",
            "applyTime": "2023.06.07",
            "batchUploadCount": "1000",
            "capacity": "4194304",
            "days": "90",
            "desc": "45元/季",
            "descPic": "",
            "expireTime": "2023.09.05",
            "familyCount": "30",
            "label": "",
            "memberCount": "30",
            "months": "",
            "order": 2,
            "originalPrice": "",
            "payType": "3",
            "price": "45",
            "productCode": "bycloud03",
            "productId": "bycloud03",
            "recycleBin": "15",
            "tips": "",
            "transientCount": "500",
            "uploadFileSize": "5120",
            "vipDesc": "白银会员",
            "vipLevel": "6",
            "vipWeight": "4000"
          },
          {
            "appStoreContractProductId": "",
            "appStoreDesc": "",
            "appStorePrice": "",
            "appStoreProductId": "",
            "applyTime": "2023.06.07",
            "batchUploadCount": "1000",
            "capacity": "4194304",
            "days": "360",
            "desc": "140元/年",
            "descPic": "",
            "expireTime": "2024.06.01",
            "familyCount": "30",
            "label": "",
            "memberCount": "30",
            "months": "",
            "order": 3,
            "originalPrice": "",
            "payType": "1",
            "price": "140",
            "productCode": "bycloud05",
            "productId": "bycloud05",
            "recycleBin": "15",
            "tips": "",
            "transientCount": "500",
            "uploadFileSize": "5120",
            "vipDesc": "白银会员",
            "vipLevel": "6",
            "vipWeight": "4000"
          }
        ]
      },
      {
        "vipLevel": "2",
        "vipDesc": "黄金会员",
        "familyCount": "50",
        "recycleBin": "30",
        "batchUploadCount": "5000",
        "capacity": "6291456",
        "uploadFileSize": "20480",
        "memberCount": "50",
        "transientCount": "1000",
        "vipWeight": "5000",
        "packages": [
          {
            "appStoreContractProductId": "ct90780132",
            "appStoreDesc": "18元/月",
            "appStorePrice": "18",
            "appStoreProductId": "90780132",
            "applyTime": "2023.06.07",
            "batchUploadCount": "5000",
            "capacity": "6291456",
            "days": "30",
            "desc": "18元/月",
            "descPic": "",
            "expireTime": "2023.07.07",
            "familyCount": "50",
            "label": "限时特惠",
            "memberCount": "50",
            "months": "",
            "order": 4,
            "originalPrice": "",
            "payType": "0",
            "price": "18",
            "productCode": "91027199",
            "productId": "91027199",
            "recycleBin": "30",
            "tips": "",
            "transientCount": "1000",
            "uploadFileSize": "20480",
            "vipDesc": "黄金会员",
            "vipLevel": "2",
            "vipWeight": "5000"
          },
          {
            "appStoreContractProductId": "ct91027207",
            "appStoreDesc": "60元/季",
            "appStorePrice": "60",
            "appStoreProductId": "91027207",
            "applyTime": "2023.06.07",
            "batchUploadCount": "5000",
            "capacity": "6291456",
            "days": "90",
            "desc": "54元/季",
            "descPic": "",
            "expireTime": "2023.09.05",
            "familyCount": "50",
            "label": "限时特惠",
            "memberCount": "50",
            "months": "",
            "order": 5,
            "originalPrice": "",
            "payType": "3",
            "price": "54",
            "productCode": "91027207",
            "productId": "91027207",
            "recycleBin": "30",
            "tips": "",
            "transientCount": "1000",
            "uploadFileSize": "20480",
            "vipDesc": "黄金会员",
            "vipLevel": "2",
            "vipWeight": "5000"
          },
          {
            "appStoreContractProductId": "ct91027230",
            "appStoreDesc": "163元/年",
            "appStorePrice": "163",
            "appStoreProductId": "91027230",
            "applyTime": "2023.06.07",
            "batchUploadCount": "5000",
            "capacity": "6291456",
            "days": "360",
            "desc": "160元/年",
            "descPic": "",
            "expireTime": "2024.06.01",
            "familyCount": "50",
            "label": "限时特惠",
            "memberCount": "50",
            "months": "",
            "order": 6,
            "originalPrice": "",
            "payType": "1",
            "price": "160",
            "productCode": "91027230",
            "productId": "91027230",
            "recycleBin": "30",
            "tips": "",
            "transientCount": "1000",
            "uploadFileSize": "20480",
            "vipDesc": "黄金会员",
            "vipLevel": "2",
            "vipWeight": "5000"
          }
        ]
      },
      {
        "vipLevel": "3",
        "vipDesc": "钻石会员",
        "familyCount": "99",
        "recycleBin": "60",
        "batchUploadCount": "9999999",
        "capacity": "8388608",
        "uploadFileSize": "9999999",
        "memberCount": "99",
        "transientCount": "9999999",
        "vipWeight": "6000",
        "packages": [
          {
            "appStoreContractProductId": "ct90780131",
            "appStoreDesc": "40元/月",
            "appStorePrice": "40",
            "appStoreProductId": "90780131",
            "applyTime": "2023.06.07",
            "batchUploadCount": "9999999",
            "capacity": "8388608",
            "days": "30",
            "desc": "38元/月",
            "descPic": "",
            "expireTime": "2023.07.07",
            "familyCount": "99",
            "label": "限时特惠",
            "memberCount": "99",
            "months": "",
            "order": 7,
            "originalPrice": "",
            "payType": "0",
            "price": "38",
            "productCode": "91027202",
            "productId": "91027202",
            "recycleBin": "60",
            "tips": "",
            "transientCount": "9999999",
            "uploadFileSize": "9999999",
            "vipDesc": "钻石会员",
            "vipLevel": "3",
            "vipWeight": "6000"
          },
          {
            "appStoreContractProductId": "ct91027211",
            "appStoreDesc": "118元/季",
            "appStorePrice": "118",
            "appStoreProductId": "91027211",
            "applyTime": "2023.06.07",
            "batchUploadCount": "9999999",
            "capacity": "8388608",
            "days": "90",
            "desc": "114元/季",
            "descPic": "",
            "expireTime": "2023.09.05",
            "familyCount": "99",
            "label": "限时特惠",
            "memberCount": "99",
            "months": "",
            "order": 8,
            "originalPrice": "",
            "payType": "3",
            "price": "114",
            "productCode": "91027211",
            "productId": "91027211",
            "recycleBin": "60",
            "tips": "",
            "transientCount": "9999999",
            "uploadFileSize": "9999999",
            "vipDesc": "钻石会员",
            "vipLevel": "3",
            "vipWeight": "6000"
          },
          {
            "appStoreContractProductId": "ct91027234",
            "appStoreDesc": "488元/年",
            "appStorePrice": "488",
            "appStoreProductId": "91027234",
            "applyTime": "2023.06.07",
            "batchUploadCount": "9999999",
            "capacity": "8388608",
            "days": "360",
            "desc": "456元/年",
            "descPic": "",
            "expireTime": "2024.06.01",
            "familyCount": "99",
            "label": "限时特惠",
            "memberCount": "99",
            "months": "",
            "order": 9,
            "originalPrice": "",
            "payType": "1",
            "price": "456",
            "productCode": "91027234",
            "productId": "91027234",
            "recycleBin": "60",
            "tips": "",
            "transientCount": "9999999",
            "uploadFileSize": "9999999",
            "vipDesc": "钻石会员",
            "vipLevel": "3",
            "vipWeight": "6000"
          }
        ]
      }
    ]
  }
}
```

# ClassifyRule

### Req

```json
{
  "header": {
    "key": "ClassifyRule",
    "resTime": 1686118434325,
    "reqSeq": 163343,
    "channel": "wohome",
    "version": "",
    "sign": "e4b0b3c4a4c29c5e74c720c82dc14417"
  },
  "body": {
    "param": "kkVvgDMX2gdSuovAHMywlQ==",
    "key": true
  }
}
```

#### param

```json
{}
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "2136cfca4a96449baa790c5572935dd8",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "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"
  }
}
```

#### data

```json
{
  "fileTypes": {
    "jpg": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "jpeg": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "png": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "bmp": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "tiff": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "tif": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "psd": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "pcd": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "raw": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "tga": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "livp": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "heic": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "ico": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "webp": {
      "subType": "image",
      "ability": "1",
      "type": "1"
    },
    "gif": {
      "subType": "gif",
      "ability": "1",
      "type": "1"
    },
    "asf": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "avi": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "dv": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "flv": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "mkv": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "mov": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "mp4": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "rm": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "swf": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "ts": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "vob": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "wmv": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "rmvb": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "mpg": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "webm": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "ogv": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "m4v": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "f4v": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "3gp": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "m3u8": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "dat": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "m3us": {
      "subType": "video",
      "ability": "1",
      "type": "2"
    },
    "ac3": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "flac": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "m4a": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "mp2": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "mp3": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "wav": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "wma": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "ape": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "mpc": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "tta": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "ogg": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "amr": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "aac": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "aiff": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "au": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "mka": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "wv": {
      "subType": "audio",
      "ability": "0",
      "type": "3"
    },
    "txt": {
      "subType": "text",
      "ability": "0",
      "type": "4"
    },
    "rtf": {
      "subType": "doc",
      "ability": "0",
      "type": "4"
    },
    "doc": {
      "subType": "word",
      "ability": "0",
      "type": "4"
    },
    "docx": {
      "subType": "word",
      "ability": "0",
      "type": "4"
    },
    "ppt": {
      "subType": "ppt",
      "ability": "0",
      "type": "4"
    },
    "pptx": {
      "subType": "ppt",
      "ability": "0",
      "type": "4"
    },
    "xls": {
      "subType": "excel",
      "ability": "0",
      "type": "4"
    },
    "xlsx": {
      "subType": "excel",
      "ability": "0",
      "type": "4"
    },
    "md": {
      "subType": "doc",
      "ability": "0",
      "type": "4"
    },
    "pdf": {
      "subType": "pdf",
      "ability": "0",
      "type": "4"
    },
    "hlp": {
      "subType": "doc",
      "ability": "0",
      "type": "4"
    },
    "csv": {
      "subType": "csv",
      "ability": "0",
      "type": "4"
    },
    "7z": {
      "subType": "zip",
      "ability": "0",
      "type": "5"
    },
    "tar": {
      "subType": "zip",
      "ability": "0",
      "type": "5"
    },
    "gz": {
      "subType": "zip",
      "ability": "0",
      "type": "5"
    },
    "zip": {
      "subType": "zip",
      "ability": "0",
      "type": "5"
    },
    "bz2": {
      "subType": "zip",
      "ability": "0",
      "type": "5"
    },
    "xz": {
      "subType": "zip",
      "ability": "0",
      "type": "5"
    },
    "wim": {
      "subType": "zip",
      "ability": "0",
      "type": "5"
    },
    "rar": {
      "subType": "zip",
      "ability": "0",
      "type": "5"
    },
    "exe": {
      "subType": "other",
      "ability": "0",
      "type": "5"
    },
    "xmind": {
      "subType": "xmind",
      "ability": "0",
      "type": "5"
    }
  },
  "fileIcons": {
    "app": {
      "zip": "https://panservice.mail.wo.cn/h5/icon/APP/zip_icon.png",
      "image": "https://panservice.mail.wo.cn/h5/icon/APP/%e5%9b%be%e7%89%87_icon.png",
      "other": "https://panservice.mail.wo.cn/h5/icon/APP/%E6%9C%AA%E7%9F%A5%E6%96%87%E4%BB%B6icon.png",
      "xmind": "https://panservice.mail.wo.cn/h5/icon/APP/xmind_icon.png",
      "gif": "https://panservice.mail.wo.cn/h5/icon/APP/gif_icon.png",
      "csv": "https://panservice.mail.wo.cn/h5/icon/APP/csv_icon.png",
      "video": "https://panservice.mail.wo.cn/h5/icon/APP/%e8%a7%86%e9%a2%91_icon.png",
      "excel": "https://panservice.mail.wo.cn/h5/icon/APP/xls_icon.png",
      "pdf": "https://panservice.mail.wo.cn/h5/icon/APP/pdf_icon.png",
      "ppt": "https://panservice.mail.wo.cn/h5/icon/APP/ppt_icon.png",
      "doc": "https://panservice.mail.wo.cn/h5/icon/APP/doc_icon.png",
      "audio": "https://panservice.mail.wo.cn/h5/icon/APP/%e9%9f%b3%e9%a2%91icon.png",
      "text": "https://panservice.mail.wo.cn/h5/icon/APP/txt_icon.png",
      "word": "https://panservice.mail.wo.cn/h5/icon/APP/word_icon.png"
    },
    "h5": {}
  }
}
```

# GetZoneInfo

### Req

```json
{
  "header": {
    "key": "GetZoneInfo",
    "resTime": 1686118434325,
    "reqSeq": 170890,
    "channel": "wohome",
    "version": "",
    "sign": "326b96ee368b30e4e9cf7bd31e305880"
  },
  "body": {
    "param": "nanNmnX3m4Anrp1DjXylISnti/HBwPrVFTjINwpyf1w=",
    "key": true
  }
}
```

#### param

```json
{ "appId": "10000001" }
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "983cfcaf29e44229ad3c75bd15b16229",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "Nm2QfrArNNXyA2JPvLzW7p6qTqISUTlUdcZCH81YaZ6N4Fy8ftbKSv6ujiSx4NUP"
  }
}
```

#### data

```json
{ "url": "https://gxupload.pan.wo.cn:8443" }
```

# QuerySysConfig

### Req

```json
{
  "header": {
    "key": "QuerySysConfig",
    "resTime": 1686118435460,
    "reqSeq": 136092,
    "channel": "wohome",
    "version": "",
    "sign": "e82bdd9c892e2a5655f5bc33f03f5ba3"
  },
  "body": {
    "param": "9rk8n0yxVyU77n0GXkIIVXiC2TbdyRoWLZYoiFXrxWh84/MBXQXLZb4IRiAc7V6oo1NoZ7knEvpuF91mPPAOlLmxpGcTCuwskaSlhKaZtqg=",
    "secret": true
  }
}
```

#### param

```json
{ "clientType": "1", "application": "WO_PAN", "clientId": "1001000021" }
```

### Resp

```json
{
  "STATUS": "200",
  "MSG": "服务调用成功！",
  "LOGID": "eebcc42481854d13adced72779881c8d",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "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"
  }
}
```

#### data

```json
{
  "config": {
    "vip_level_file_num_limit": {
      "0": 100,
      "1": 1000,
      "2": 5000,
      "3": 0,
      "4": 100,
      "5": 100,
      "6": 1000
    },
    "file_max_upload_alert": 5000,
    "vip_activity_icon": {
      "0": "https://panservice.mail.wo.cn/upload/files/pic/wohome/1680076451127.png",
      "1": "",
      "2": "",
      "3": "",
      "4": "",
      "5": "",
      "6": ""
    },
    "file_max_select": 100,
    "title_icon": {
      "谨慎的收藏家": "https://panservice.mail.wo.cn/upload/files/pic/wohome/1680076103413.png",
      "空间守护者": "https://panservice.mail.wo.cn/upload/files/pic/wohome/1680076114414.png",
      "江湖历练标兵": "https://panservice.mail.wo.cn/upload/files/pic/wohome/1680076123905.png",
      "满满的王者实力": "https://panservice.mail.wo.cn/upload/files/pic/wohome/1680076128834.png",
      "云游小卫士": "https://panservice.mail.wo.cn/upload/files/pic/wohome/1680076119036.png",
      "最美存储大师": "https://panservice.mail.wo.cn/upload/files/pic/wohome/1680076095185.png"
    },
    "vip_activity_url": {
      "0": "https://u.10010.cn/qAxi9",
      "1": "https://u.10010.cn/qAxHX",
      "2": "https://u.10010.cn/qAxHX",
      "3": "",
      "4": "https://u.10010.cn/qAxHX",
      "5": "https://u.10010.cn/qAxHX",
      "6": "https://u.10010.cn/qAxHX"
    },
    "vip_activity_desc": {
      "0": "尊享会员2T超大存储限时免费领取，快去看看>>>",
      "1": "了解更多会员权益详情，点击查看>>>",
      "2": "钻石会员享文件传输无大小限制，避免烦恼，去升级>>>",
      "3": "您已享有最高级别权益，感谢您的支持与信任。",
      "4": "畅享会员低至6元/月，更大空间等你享用，去了解>>>",
      "5": "尊享会员享2T超大空间任你享用，去了解>>>",
      "6": "了解更多会员权益详情，点击查看>>>"
    },
    "title_desc": {
      "谨慎的收藏家": "积少成多，争取每天进步一点点，加油！",
      "空间守护者": "你还在犹豫吗？世界很大，是时候该勇敢的踏出第一步了！",
      "江湖历练标兵": "再冷的石头，坐上三年也会暖，何况您已在熟的路上了，加油！",
      "满满的王者实力": "天助自助者，您已用实力证明了自己的选择是正确的，继续保持哦！",
      "云游小卫士": "作为新秀，您的表现颇为神勇，继续努力，未来可期",
      "最美存储大师": "您已达到伟大而高贵的时刻，请继续保持哦～"
    }
  }
}
```

# FamilyUserCurrentEncode

### Req

```json
{
  "header": {
    "key": "FamilyUserCurrentEncode",
    "resTime": 1686118435496,
    "reqSeq": 164667,
    "channel": "wohome",
    "version": "",
    "sign": "d8d0332ab0f2423254083708d12bd6eb"
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
  "LOGID": "50f30af2d857438fa6ef7bfaa6a1e933",
  "RSP": {
    "RSP_CODE": "0000",
    "RSP_DESC": "成功",
    "DATA": "Lf40eIyAXOPhSckvP+VQyeTRhshqJIot4nfOYpc/HRmHgBTVlsdwrnAiZAIh1ba7+uYiRWJJiPku4CoDMc2vONjMuDqIBSiFMqK2f8nDr/gJIMQ93tXA4dKdbPqIWS4lvHwQH0+VKb8rUTHGkGem4JlKT3FGbagafIGcKwx8OEhkdY4003TYITxmGKRWl7Px10ron68aJaFZZu2Sw7h7d9y5YLsW/P5uk8QRB85pTRxrClStEDlhGQLue/b+o7iw3nnU7c0k8WFnjfF5b2DXE5BFqJ+k5/SZ0Gd4G4YH0QP7ZD7ybS16MnX9+OzYkRslGo0BB9/dadKviulmH9zzOcAUANTiwBEYB4Bq8YIqZ39qarisqjSX0LPN5Q+C6Wsi"
  }
}
```

#### data

```json
{
  "count": "1",
  "defaultHomeId": 9263097,
  "defaultHomeName": "186****5244的家",
  "groupHeadUrl": "  https://iotpservice.smartont.net/upload/headPortrait/person_1600740115826232.png",
  "groupName": "186****5244的家",
  "id": 9263097,
  "memberRole": "1",
  "owner  Id": "186****5244",
  "unreadFlag": "1"
}
```
