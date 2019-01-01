# vveq-api 

#### User

* `GET`  /api/user/usernameIsExists/:username  查询用户名是否存在
* `POST` /api/user 用户注册
* `POST` /api/user/login 用户登录
* `GET`  /api/user/:uid 获取用户信息

#### Verify

* `GET` /api/verify 获取验证码
* `POST` /api/verify 验证码校验

#### Like

### User

* `POST` /api/like 点赞/取消点赞

#### `GET` /api/user/usernameIsExists/:username  查询用户名是否存在

request

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
username|String|是|用户名

response

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
status|Int|是|1：用户名不存在，0：用户名已存在

#### `POST` /api/user 用户注册

request

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
from.username|String|是|用户名
from.password|String|是|密码
from.email|String|是|电子邮件
verifyForm.CaptchaType|String|是|audio：音频，character：运算表达式，为空表示数字
verifyForm.Id|String|是|系统生成的验证Id
verifyForm.VerifyValue|String|是|用户输入的验证码

response

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
status|Int|是|1：注册成功，0：注册失败，-1：验证码错误
data|Struct|否|注册成功返回用户信息

#### `POST` /api/user/login 用户登录

request

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
username|string|是|用户名
password|string|是|密码

response 

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
status|Int|是|1：成功，0：失败
data|Struct|否|登录成功返回用户信息


#### `GET`  /api/user/:uid 获取用户信息

request

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
uid|Int|是|用户id

response 

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
status|Int|是|1：成功，0：失败
data|Struct或Error|否|注册成功返回用户信息,失败返回失败信息

### Verify

#### `POST` /api/verify/getCaptcha 获取验证码

request

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
status|Int|是|1：成功，0：失败
CaptchaType|String|是|audio：音频，character：运算表达式，为空表示数字
Id|String|是|系统生成的验证Id
VerifyValue|String|是|用户输入的验证码

response

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
status|Int|是|1：成功，0：失败
data|String|否|Base64图片验证码
captchaId|String|否|验证码Id，用户验证码的校验

#### `POST` /api/verify/verifyCaptcha 验证码校验

request

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
CaptchaType|String|是|audio：音频，character：运算表达式，为空表示数字
Id|String|是|系统生成的验证Id
VerifyValue|String|是|用户输入的验证码

response

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
status|Int|是|1：成功，0：失败

#### `POST` /api/like 点赞/取消点赞

request

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
type|int|是|点赞对象，0-主题 1-回复
uid|int|是|用户id
pid|int|是|对象id

response

参数名|类型|是否必传|说明
:--|:--:|:--:|:--:
status|Int|是|1：成功，0：失败
