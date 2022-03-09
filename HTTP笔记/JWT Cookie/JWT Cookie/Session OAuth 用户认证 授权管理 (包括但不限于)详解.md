# JWT Cookie/Session OAuth 用户认证 授权管理 (包括但不限于)详解

先来梳理一下用户交互&登陆全流程吧

## 用户登陆&交互全流程

1. 第一步先进行**身份认证**，身份认证有如下几种方式：
	* 账户密码
	* 验证码(邮箱、谷歌验证……)
	* 生物验证(人脸识别、指纹解锁……)
2. 认证成功后，为了让服务器知道后续操作是谁，就需要**维持上下文状态**，实现方式也有如下几种：
	* 使用 Cookie/Session 来维持
	* 使用 JWT 

***

## 身份认证

身份认证目的是**让服务器知道你到底是谁，随后服务器再根据你的身份，开放相应的权限**。稍微完善的网站，都会提供用户多个登陆途径，以防某单一方法失败，导致用户无法使用整个网站。

### 账户密码

账户密码是最传统的一种方式，这种方式是存在一些风险的，有这么几种情况可能会出现密码泄露：

* 客户端：暴力破解
    > 黑客通过程序不断的穷举，直至成功试出正确的密码。
* 传输层：密码传输过程中被抓包
    > 密码明文传输，或者黑客拥有一个MD5库，可以将简单的密码进行破解。
* 服务端：数据库被攻破，密码被读取或者修改

对应解决方案：

* 客户端：增加密码防护机制，比如连续5次错误后暂停5分钟……

* 传输层：客户端将密钥加密，同时使用加密传输协议，比如[HTTPS](TODO)

* 服务端：

    1. 不使用明文存储密码
	2. 加盐字段
	3. 使用不可逆的加密算法计算密文，存储密文，几种加密算法简介🔗：[加密算法](TODO)
	4. 将密文与账号ID绑定，计算密文的时候fn(密码+盐+账号ID)，这样即使脱库，攻击者也无法通过更换密钥的方式来控制账户

服务端的方案比较抽象，下面举一个案例说明一下。

#### XXX项目是如何存储密钥的？

这是一张存储管理员信息数据库表：

```sql
-- 管理员表
DROP TABLE IF EXISTS `t_admin`;
CREATE TABLE `t_admin` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	`address` CHAR(42) UNIQUE NOT NULL COMMENT '管理员地址',
	`salt` VARCHAR(18) NOT NULL COMMENT '生成摘要的盐',
	`digest` CHAR(64) NOT NULL COMMENT '摘要',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

其中：`salt`是盐，`digest`是用来校验的摘要，也就是fn(密码+盐+账号ID)计算出来的结果。

下面是登陆验证时的代码：

```golang
type LoginParam struct {
   Address  string `validate:"required"`
   Password string `validate:"required"`
}

func (l *LoginParam) Verify() (bool, error) {
   m := model.TAdmin{
	  Address: l.Address,
   }
   one, err := m.One()
   if err != nil {
	  return false, errors.Wrap(err, "TAdmin One error")
   }

   //  计算摘要的方式在这里
   digest := sha256.Sum256([]byte(strconv.Itoa(int(one.ID)) + one.Salt + l.Password))

   return strings.ToUpper(hex.EncodeToString(digest[:])) == strings.ToUpper(one.Digest), nil
}
```

解释一下上述代码最关键的一行：使用sha256摘要算法，对“账号ID+盐+密码”进行字符串拼接后的字符串计算摘要，随后再与数据库中记录的摘要相比较。

#### 这么做的好处在哪？

- 首先，没有存储明文，这样即使数据库被攻击，黑客也无法得到用户的密码。

- 再者，即使黑客将所有账户的摘要替换成他自己账户的摘要，因为计算的时候，需要与账户的ID拼接，所以验证的时候，摘要还是不正确的。

- 最后，若黑客将ID和盐都修改，那么原本攻击的那个账户所有的权限与资源就都失效了，因为业务逻辑上是与ID绑定的。

这样，即使黑客获取了数据库用户表的信息，也无法登陆其他账户。