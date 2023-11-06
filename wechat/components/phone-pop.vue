<template>
	<view v-if='!hasPhone'>
		<div @click="closePop">
		    <!-- 需要粘贴的部分，替换掉【授权登录】按钮这一行代码 -->
		    <div class="modal-mask">
		    </div>
		    <div class="modal-dialog">
		      <div class="modal-content">
		        <image class="img" src="/static/hot_product_banner.png"></image>
		        <div class="content-text">
		          <p class="key-bold-tip">注册会员</p>
		          <p class="key-bold">注册成为会员享受更多优惠</p>
		          <p class="little-tip">我们的生活圈：</p>
		          <p class="little-content">
		            注册成为会员，一店消费，多家优惠，欢迎体验
		          </p>
		        </div>
		      </div>
		      <div class="modal-footer">
		         <!-- 小程序集成的API，通过button来授权登录 -->
		        <button class='btn' open-type='getPhoneNumber' @getphonenumber="getWXPhone">
		        	一键注册
		        </button>
		      </div>
		    </div>
		</div>
	</view>
</template>

<script setup>
	import {
		mapState
	} from 'vuex';
	import {
		wxLogin,getWXPhoneNumber
	} from '@/api/member.js';
	export default {
		name:"phone-pop",
		data() {
			return {
				
			};
		},
		computed: {
			...mapState(['hasPhone']),
		},
		
		methods: {
			decryptPhoneNumber: function(e) {
				console.log("-------decryptPhoneNumber------", e.detail)
			},
			//绑定手机
			getWXPhone: function(e) {
				console.log("------getPhoneNumber---", e.detail.code)
				var _this = this;
				if(e.detail.errMsg == "getPhoneNumber:ok") {
					console.log("------getPhoneNumber:ok---")
					// uni.checkSession({
					//   success: function (res) {
					// 	  console.log("------getWXPhoneNumber----checkSession--success-")
					// 	 if	(res.errMsg == 'checkSession:ok') {
					// 		 _this.requestWxPhoneNumber({
					// 			encryptedData: e.detail.encryptedData,
					// 			iv: e.detail.iv,
					// 			// sessionKey: wx.getStorageSync("SessionKey"),
					// 			openId: wx.getStorageSync("OpenId"),
					// 		 })
					// 	 }
					//   },
					//   fail: function () {
					// 		console.log("------getWXPhoneNumber----checkSession--fail---")
					// 		_this.initWXLogin()
					// 		uni.showToast({ title: '注册失败', duration: 2000 })
					// 		// _this.requestWxPhoneNumber({
					// 		// encryptedData: e.detail.encryptedData,
					// 		// iv: e.detail.iv,
					// 		// sessionKey: wx.getStorageSync("SessionKey"),
					// 		// openId: wx.getStorageSync("OpenId"),
					// 	  // })
					//   }
					// })
				} else {
					uni.showToast({ title: '取消注册会员', duration: 2000 })
				}
			},
			requestWxPhoneNumber(data) {
				getWXPhoneNumber(data).then(res=>{
					console.log("------getWXPhoneNumber----res---", res)
					if (res.code == 0) {
						uni.showToast({ title: '注册成功', duration: 2000 })
						console.log(res.data.phone)//成功后打印微信手机号
						this.$store.commit('setPhone', true)
						// 传递给父组件 关闭弹窗
						this.$emit('success', true);
					}
					else {
						uni.showToast({ title: '注册会员失败', duration: 2000 })
					}
				});
			},
　　　　　　//登录
			initWXLogin() {
				let _this = this;
			   // 1.wx获取登录用户code
				uni.login({
					provider: 'weixin',
					success: function(loginRes) {
						let code = loginRes.code;
						if (loginRes.code) {
							wxLogin({
								code: loginRes.code,
							}).then(res => {
								console.log("------login---res--------", res.data)
								uni.setStorageSync('OpenId', res.data.openid);
								wx.setStorageSync("WxCode", code)
								wx.setStorageSync("WxCodeTime", (new Date()).getTime())
							}).catch((e) => {
								console.log("------login---res--------", e)
							});
						} else {
							uni.showToast({ title: '获取code失败', duration: 2000 })
							console.log('获取code失败' + loginRes.errMsg)
						}
					},
					fail: (res) => {
						uni.showToast({ title: '获取code失败', duration: 2000 })
						console.log('获取code失败' + res.errMsg)
					}
				});
			},
			closePop() {
				// 传递给父组件 关闭弹窗
				this.$emit('close', true);
			},
		},
		mounted() {
			// this.initWXLogin()
		}
	}
</script>

<style>
.modal-mask {
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  left: 0;
  background: #000;
  opacity: 0.5;
  overflow: hidden;
  z-index: 9000;
  color: #fff;
}
.modal-dialog {
  box-sizing: border-box;
  width: 560rpx;
  overflow: hidden;
  position: fixed;
  top: 40%;
  left: 0;
  z-index: 9999;
  background: #fff;
  margin: -150rpx 95rpx;
  border-radius: 16rpx;
}
.modal-content {
  box-sizing: border-box;
  display: flex;
  padding: 0rpx 53rpx 50rpx 53rpx;
  font-size: 32rpx;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}
.content-tip {
  text-align: center;
  font-size: 36rpx;
  color: #333333;
}
.content-text {
  /* height:230px; */
  padding:10px 0px 10px 0px;
  font-size:14px;
}
.modal-footer {
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  border-top: 1px solid #e5e5e5;
  font-size: 16px;
  font-weight:bold;
  /* height: 45px; */
  line-height: 45px;
  text-align: center;
  background:#feb600;
}
.btn {
  width: 100%;
  height: 100%;
  background:#feb600;
  color:#FFFFFF;
  font-weight:bold;
}
.img {
  width: 560rpx;
  height:140rpx;
}
.little-tip {
  padding-top:15px;
  padding-bottom:3px;
  font-size: 14px;
  font-weight:bold;
  color: #feb600;
}
.little-content {
  padding-top:5px;
  font-size: 13px;
  color:#606060;
}
.key-bold-tip {
  padding-top:5px;
  font-size: 15px;
  font-weight:bold;
  color: #feb600;
}
.key-bold {
  padding-top:5px;
  font-size: 14px;
  /* font-weight:bold; */
}
</style>