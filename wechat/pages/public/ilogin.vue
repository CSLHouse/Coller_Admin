<template>
	<view class="container">
		<view class="left-bottom-sign"></view>
		<view class="back-btn yticon icon-zuojiantou-up" @click="navBack"></view>
		<view class="right-top-sign"></view>
		<!-- 设置白色背景防止软键盘把下部绝对定位元素顶上来盖住输入框等 -->
		<view class="wrapper">
			<!-- <view class="left-top-sign">LOGIN</view> -->
			<view class="welcome">
				欢迎使用猪迪克！
			</view>
			
			<!-- #ifdef MP-WEIXIN -->
			<view v-if="isCanUse">
				<view>
					<view class='header'>
						<image src='../../static/wx_login.png'></image>
					</view>
					<view class='content'>
						<view>申请获取以下权限</view>
						<text>获得你的公开信息(昵称，头像、地区等)</text>
					</view>
					<!--新版登录方式-->
					<button class='bottom' type='primary' @click="wxGetUserInfo"> 微信一键登录 </button>
					<!--旧版登录方式-->
					<!-- <button class='bottom' type='primary' open-type="getUserInfo" withCredentials="true" lang="zh_CN" @getuserinfo="wxGetUserInfo">
						授权登录
					</button> -->
				</view>
			</view>
			<view class="moadl-mask" v-if="showModal">
				<view class='header'>
					<image src='../../static/wx_login.png'></image>
				</view>
				<view class='content'>
					<view>绑定手机号</view>
					<text>会员注册</text>
				</view>
				
				<button class='bottom' type='primary' open-type='getPhoneNumber' @getphonenumber="getPhoneNumber">
					微信用户一键绑定
				 <image src='../images/showWx.png' class='iconWx'></image>微信用户一键绑定
				</button>
			</view>
			<!-- #endif -->
			
			<!-- <view class="input-content">
				<view class="input-item">
					<text class="tit">用户名</text>
					<input type="text" v-model="username" placeholder="请输入用户名" maxlength="11"/>
				</view>
				<view class="input-item">
					<text class="tit">密码</text>
					<input type="text" v-model="password" placeholder="8-18位不含特殊字符的数字、字母组合" placeholder-class="input-empty" maxlength="20"
					 password @confirm="toLogin" />
				</view>
			</view>
			<button class="confirm-btn" @click="toLogin" :disabled="logining">登录</button>
			<button class="confirm-btn2" @click="toRegist" >获取体验账号</button>
			<view class="forget-section" @click="toRegist">
				忘记密码?
			</view> -->
		</view>
		<!-- <view class="register-section">
			还没有账号?
			<text @click="toRegist">马上注册</text>
		</view> -->
	</view>
</template>

<script>
	import {
		mapMutations
	} from 'vuex';
	import {
		memberLogin,memberInfo,wxLogin,getWXPhoneNumber,UpdateWxUserInfo
	} from '@/api/member.js';
	export default {
		data() {
			return {
				sessionKey: '',
				openId: '',
				nickName: null,
				avatarUrl: null,
				isCanUse: uni.getStorageSync('isCanUse')||true,//默认为true
				gender: 0,
				showModal: false,
			}
		},
		onLoad() {
			this.initWXLogin();
		},
		methods: {
			...mapMutations(['login']),
			navBack() {
				uni.navigateBack();
			},
			toRegist() {
				uni.navigateTo({url:'/pages/public/register'});
			},
			async toLogin() {
				this.logining = true;
				memberLogin({
					username: this.username,
					password: this.password
				}).then(response => {
					let token = response.data.tokenHead+response.data.token;
					uni.setStorageSync('token',token);
					uni.setStorageSync('username',this.username);
					uni.setStorageSync('password',this.password);
					memberInfo().then(response=>{
						this.login(response.data);
						uni.navigateBack();
					});
				}).catch(() => {
					this.logining = false;
				});
			},
			//第一授权获取用户信息===》按钮触发
			async wxGetUserInfo() {
				let _this = this;
				uni.showModal({
					title: '登录提示',
					content: '您需要授权微信登录后才能正常使用小程序功能',
					success(res) {
						if (res.confirm) {
							uni.getUserProfile ({
								provider: 'weixin',
								desc: '获得你的昵称，头像、地区',
								success: function(infoRes) {
									console.log("-------wxGetUserInfo--infoRes:", infoRes)
									_this.nickName = infoRes.userInfo.nickName; //昵称
									_this.avatarUrl = infoRes.userInfo.avatarUrl; //头像
									_this.gender = infoRes.userInfo.gender;	// 性别
									uni.setStorageSync('userInfo', infoRes.userInfo);
									_this.updateUserInfo()
									try {
										_this.isCanUse = false
										uni.setStorageSync('isCanUse', false);//记录是否第一次授权  false:表示不是第一次授权
										_this.showDialogBtn();//调用一键获取手机号弹窗
									} catch (e) {}
								},
								fail(res) {
									uni.showToast({ title: '获取个人信息失败', duration: 2000 })
								}
							});
						}
					}
				})
			},
			// 显示一键获取手机号弹窗
			showDialogBtn: function () {
				this.showModal = true	//修改弹窗状态为true，即显示
			},
			// 隐藏一键获取手机号弹窗
			hideModal: function () {
				this.showModal = false	//修改弹窗状态为false,即隐藏
			},
			
			//绑定手机
			async getPhoneNumber(e) {
				var _this = this;
				if(e.detail.errMsg == "getPhoneNumber:ok") {
					uni.checkSession({
					  success: function (res) {
						 if	(res.errMsg == 'checkSession:ok') {
							 _this.requestWxPhoneNumber({
								encryptedData: e.detail.encryptedData,
								iv: e.detail.iv,
								sessionKey: wx.getStorageSync("SessionKey"),
								openId: wx.getStorageSync("OpenId"),
							 })
						 }
					  },
					  fail: function () {
						  _this.initWXLogin()
						  _this.requestWxPhoneNumber({
						  	encryptedData: e.detail.encryptedData,
						  	iv: e.detail.iv,
						  	sessionKey: wx.getStorageSync("SessionKey"),
						  	openId: wx.getStorageSync("OpenId"),
						  })
					  }
					})
				} else {
					uni.showToast({ title: '取消注册会员', duration: 2000 })
				}
			},
			requestWxPhoneNumber(data) {
				getWXPhoneNumber(data).then(res=>{
					console.log("------getWXPhoneNumber----res---", res)
					if (res.code == 0) {
						console.log("登录成功")
						_this.hideModal();
						console.log(res.data.phoneNumber)//成功后打印微信手机号
					}
					else {
						uni.showToast({ title: '注册会员失败', duration: 2000 })
					}
					uni.navigateBack();
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
								uni.setStorageSync('SessionKey', res.data.session_key);
								// _this.openId = res.data.openid
								// _this.sessionKey = res.data.session_key
							}).catch((e) => {
								console.log("------login---res--------", e)
							});
						} else {
							uni.showToast({ title: '获取code失败', duration: 2000 })
							console.log('获取code失败' + loginRes.errMsg)
							_this.wxLogin()
						}
					},
					fail: (res) => {
						uni.showToast({ title: '获取code失败', duration: 2000 })
						console.log('获取code失败' + res.errMsg)
						_this.wxLogin()
					}
				});
			},
			updateUserInfo () {
				let userInfo = wx.getStorageSync("userInfo")
				UpdateWxUserInfo({
					avatarUrl: userInfo.avatarUrl,
					nickName: userInfo.nickName,
					gender: userInfo.gender,
					openId: wx.getStorageSync("OpenId"),
				}).then(res=>{
					if (res.code == 0) {
						console.log("------wxUserInfo:", res)
						console.log("登录成功")
						this.login(res.data);
					}
					else {
						console.log(res);
					}
				});
			},
		},

	}
</script>

<style lang='scss'>
	page {
		background: #fff;
	}
	
	.header {
		margin: 90rpx 0 90rpx 50rpx;
		border-bottom: 1px solid #ccc;
		text-align: center;
		width: 650rpx;
		height: 300rpx;
		line-height: 450rpx;
	}

	.header image {
		width: 200rpx;
		height: 200rpx;
	}

	.content {
		margin-left: 50rpx;
		margin-bottom: 90rpx;
	}

	.content text {
		display: block;
		color: #9d9d9d;
		margin-top: 40rpx;
	}

	.bottom {
		border-radius: 80rpx;
		margin: 70rpx 50rpx;
		font-size: 35rpx;
	}
		
	.container {
		padding-top: 115px;
		position: relative;
		width: 100vw;
		height: 100vh;
		overflow: hidden;
		background: #fff;
	}

	.wrapper {
		position: relative;
		z-index: 90;
		background: #fff;
		padding-bottom: 40upx;
	}

	.back-btn {
		position: absolute;
		left: 40upx;
		z-index: 9999;
		padding-top: var(--status-bar-height);
		top: 40upx;
		font-size: 40upx;
		color: $font-color-dark;
	}

	.left-top-sign {
		font-size: 120upx;
		color: $page-color-base;
		position: relative;
		left: -16upx;
	}

	.right-top-sign {
		position: absolute;
		top: 80upx;
		right: -30upx;
		z-index: 95;

		&:before,
		&:after {
			display: block;
			content: "";
			width: 400upx;
			height: 80upx;
			background: #b4f3e2;
		}

		&:before {
			transform: rotate(50deg);
			border-radius: 0 50px 0 0;
		}

		&:after {
			position: absolute;
			right: -198upx;
			top: 0;
			transform: rotate(-50deg);
			border-radius: 50px 0 0 0;
			/* background: pink; */
		}
	}

	.left-bottom-sign {
		position: absolute;
		left: -270upx;
		bottom: -320upx;
		border: 100upx solid #d0d1fd;
		border-radius: 50%;
		padding: 180upx;
	}

	.welcome {
		position: relative;
		left: 50upx;
		top: -90upx;
		font-size: 46upx;
		color: #555;
		text-shadow: 1px 0px 1px rgba(0, 0, 0, .3);
	}

	.input-content {
		padding: 0 60upx;
	}

	.input-item {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		justify-content: center;
		padding: 0 30upx;
		background: $page-color-light;
		height: 120upx;
		border-radius: 4px;
		margin-bottom: 50upx;

		&:last-child {
			margin-bottom: 0;
		}

		.tit {
			height: 50upx;
			line-height: 56upx;
			font-size: $font-sm+2upx;
			color: $font-color-base;
		}

		input {
			height: 60upx;
			font-size: $font-base + 2upx;
			color: $font-color-dark;
			width: 100%;
		}
	}

	.confirm-btn {
		width: 630upx;
		height: 76upx;
		line-height: 76upx;
		border-radius: 50px;
		margin-top: 70upx;
		background: $uni-color-primary;
		color: #fff;
		font-size: $font-lg;

		&:after {
			border-radius: 100px;
		}
	}
	
	.confirm-btn2 {
		width: 630upx;
		height: 76upx;
		line-height: 76upx;
		border-radius: 50px;
		margin-top: 40upx;
		background: $uni-color-primary;
		color: #fff;
		font-size: $font-lg;
	
		&:after {
			border-radius: 100px;
		}
	}

	.forget-section {
		font-size: $font-sm+2upx;
		color: $font-color-spec;
		text-align: center;
		margin-top: 40upx;
	}

	.register-section {
		position: absolute;
		left: 0;
		bottom: 50upx;
		width: 100%;
		font-size: $font-sm+2upx;
		color: $font-color-base;
		text-align: center;

		text {
			color: $font-color-spec;
			margin-left: 10upx;
		}
	}
</style>
