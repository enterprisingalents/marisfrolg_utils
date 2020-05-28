package Marisfrolg_utils

type WXTemplate struct {

	/// <summary>
	/// 用户OPENID
	/// </summary>
	Touser string `json:"touser"`
	/// <summary>
	/// 模板ID
	/// </summary>
	Template_id string `json:"template_id"`
	/// <summary>
	/// 点击详情需要跳转的URL(模板跳转链接（海外帐号没有跳转能力）)
	/// </summary>
	Url string `json:"url"`
	/// <summary>
	/// 跳小程序所需数据，不需跳小程序可不用传该数据
	/// </summary>
	Miniprogram Miniprogram `json:"miniprogram"`
	/// <summary>
	/// 模板数据
	/// </summary>
	Data Data `json:"data"`
}


type Miniprogram struct {

	/// <summary>
	/// 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	/// </summary>
	Appid string `json:"appid"`
	/// <summary>
	/// 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
	/// </summary>
	Pagepath string `json:"pagepath"`
}

type Data struct {
	First First `json:"first"`
	/// <summary>
	/// 第一个关键字内容
	/// </summary>
	Keyword1 Keyword1 `json:"keyword1"`
	/// <summary>
	/// 第二个关键字内容
	/// </summary>
	Keyword2 Keyword2 `json:"keyword2"`
	/// <summary>
	/// 第三个关键字内容
	/// </summary>
	Keyword3 Keyword3 `json:"keyword3"`
	/// <summary>
	/// 第四个关键字内容
	/// </summary>
	Keyword4 Keyword4 `json:"keyword4"`
	/// <summary>
	/// 第五个关键字内容
	/// </summary>
	Keyword5 Keyword5 `json:"keyword5"`

	/// <summary>
	/// 第六个关键字内容
	/// </summary>
	Keyword6 Keyword6 `json:"keyword6"`

	/// <summary>
	/// 消费模板积分关键字内容
	/// </summary>
	Point Point `json:"point"`
	/// <summary>
	/// 消费模板时间关键字内容
	/// </summary>
	Time Time `json:"time"`
	/// <summary>
	/// 消费模板门店关键字内容
	/// </summary>
	Org Org `json:"org"`
	/// <summary>
	/// 消费模板类型关键字内容
	/// </summary>
	Types Types `json:"types"`
	/// <summary>
	/// 消费模板金额关键字内容
	/// </summary>
	Money Money `json:"money"`

	/// <summary>
	/// 积分兑换成功产品名关键字内容
	/// </summary>
	ProductType ProductType `json:"product_type"`

	/// <summary>
	/// 积分兑换成功产品名称关键字内容
	/// </summary>
	Name Name `json:"name"`

	/// <summary>
	/// 积分兑换成功积分（兑换积分）关键字内容
	/// </summary>
	Points Points `json:"points"`
	/// <summary>
	/// 积分兑换成功时间（兑换时间）关键字内容
	/// </summary>
	Date Date `json:"date"`

	/// <summary>
	/// 备注
	/// </summary>
	Remark Remark `json:"remark"`
}

type First struct {
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

type Keyword1 struct {
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

type Keyword2 struct {
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

type Keyword3 struct {
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}
type Keyword4 struct {
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}
type Keyword5 struct {
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}
type Keyword6 struct {
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

/// <summary>
/// 备注
/// </summary>
type Remark struct {

	/// <summary>
	/// 值
	/// </summary>
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

/// <summary>
/// 消费模板积分
/// </summary>
type Point struct {

	/// <summary>
	/// 值
	/// </summary>
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

/// <summary>
/// 消费模板时间
/// </summary>
type Time struct {

	/// <summary>
	/// 值
	/// </summary>
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

/// <summary>
/// 消费模板门店
/// </summary>
type Org struct {

	/// <summary>
	/// 值
	/// </summary>
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

/// <summary>
/// 消费模板类型
/// </summary>
type Types struct {
	/// <summary>
	/// 值
	/// </summary>
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

/// <summary>
/// 消费模板金额
/// </summary>
type Money struct {

	/// <summary>
	/// 值
	/// </summary>
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

/// <summary>
/// 积分兑换成功产品名
/// </summary>
type ProductType struct {

	/// <summary>
	/// 值
	/// </summary>
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

/// <summary>
/// 积分兑换成功产品名称
/// </summary>
type Name struct {

	/// <summary>
	/// 值
	/// </summary>
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

/// <summary>
/// 积分兑换成功积分（兑换积分）
/// </summary>
type Points struct {

	/// <summary>
	/// 值
	/// </summary>
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}

/// <summary>
/// 积分兑换成功时间（兑换时间）
/// </summary>
type Date struct {
	/// <summary>
	/// 值
	/// </summary>
	Value string `json:"value"`
	/// <summary>
	/// 模板内容字体颜色，不填默认为黑色
	/// </summary>
	Color string `json:"color"`
}
