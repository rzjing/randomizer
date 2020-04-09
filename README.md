# randomizer

随机 .* 生成器。

## 已支持生成种类

* 指定区间内的随机数字
* 指定长度的随机字符串
* 随机 Mac 地址
* 随机 IPv4 地址
* 随机省份
* 随机城市

## 示例

```gotemplate
// 随机 数字
fmt.Println(handler.RInt64(0, 100))
22

// 随机 字符串
fmt.Println(handler.RString(8))
oIlXCIV7

// 随机 MAC 地址
fmt.Println(handler.RMac())
79:6e:4a:74:59:58

// 随机 IPv4 地址
fmt.Println(handler.RIPv4())
47.37.88.36

// 随机 省份
fmt.Println(handler.RProvince())
吉林省

// 随机 城市
fmt.Println(handler.RCity())
上饶市
```
