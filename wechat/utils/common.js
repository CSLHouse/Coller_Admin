export function checkCodeValid() {
  const code = wx.getStorageSync("WxCode")
  const codeTime = wx.getStorageSync("WxCodeTime")
  console.log("----code:", code)
  console.log("----codeTime:", codeTime)
  console.log("----codeTime:", (new Date()).getTime(), codeTime + 27000 - (new Date()).getTime())
  //微信官方称code有效时间为五分钟，保险起见设置 4.5 分钟
  if (code && (codeTime + 27000 > (new Date()).getTime())) {
  	return true
  }
  return false
}

// 保留小数点数值后两位，尾数四舍五入
export function numFilter (value) {
	// 截取当前数据到小数点后两位
	let realVal = parseFloat(value).toFixed(2)
	return realVal
}
export default checkCodeValid