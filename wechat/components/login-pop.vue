<template>
	<view v-if='!hasLogin'>
		<div @click="closePop">
		    <!-- 需要粘贴的部分，替换掉【授权登录】按钮这一行代码 -->
		    <div class="modal-mask" >
		    </div>
		    <div class="modal-dialog">
		      <div class="modal-content">
				<image class="img" src="/static/hot_product_banner.png"></image>
		        <div class="content-text">
		          <p class="key-bold-tip">猪迪克星动乐园</p>
		          <p class="key-bold">星动乐园，为每一个星星般美好的孩子而打造</p>
		          <p class="little-tip">我们的生活圈：</p>
		          <p class="little-content">
		            致力于打造附近几个小区自己的生活圈，
		          </p>
		          <p class="little-content">
		            一店消费，多家优惠，欢迎体验
		          </p>
		        </div>
		      </div>
		      <div class="modal-footer">
		         <!-- 小程序集成的API，通过button来授权登录 -->
		         <!-- <button open-type="getUserInfo" lang="zh_CN" class='btn' @getuserinfo="login">授权登录</button> -->
				 <button class='btn' lang="zh_CN" @click="wxGetUserInfo"> 微信一键登录 </button>
		      </div>
		    </div>
		</div>
	</view>
</template>

<script>
	import {
		mapMutations, mapState
	} from 'vuex';
	import {
		UpdateWxUserInfo, wxLogin, wxRefreshLogin
	} from '@/api/member.js';
	export default {
		name:"login-pop",
		data() {
			return {
				
			};
		},
		computed: {
			...mapState(['hasLogin']),
		},
		methods: {
			...mapMutations(['login', 'refreshLoginSession']),
			//登录
			// wxLoginRefreshSession() {
			// 	let _this = this;
			// 	// 1.wx获取登录用户code
			// 	uni.login({
			// 		provider: 'weixin',
			// 		success: function(loginRes) {
			// 			let code = loginRes.code;
			// 			if (loginRes.code) {
			// 				wxLogin({
			// 					code: loginRes.code,
			// 				}).then(res => {
			// 					console.log("------login---res--------", res.data)
			// 					uni.setStorageSync('OpenId', res.data.openid);
			// 					// uni.setStorageSync('SessionKey', res.data.session_key);
			// 				}).catch((e) => {
			// 					console.log("------login---res--------", e)
			// 				});
			// 			} else {
			// 				uni.showToast({ title: '获取code失败', duration: 2000 })
			// 				console.log('获取code失败' + loginRes.data)
			// 			}
			// 		},
			// 		fail: (res) => {
			// 			uni.showToast({ title: '获取code失败,请重新打开小程序！', duration: 2000 })
			// 			console.log('获取code失败' + res.data)
			// 		}
			// 	});
			// },
			initWXLogin() {
				const openid = uni.getStorageSync('OpenId');
				if (!openid) {
					this.refreshLoginSession()
					console.log("==initWXLogin=openid=", openid)
				}
			},
			async wxGetUserInfo() {
				let _this = this;
				this.initWXLogin()
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
									_this.updateUserInfo(infoRes.userInfo.nickName, infoRes.userInfo.avatarUrl, infoRes.userInfo.gender)
								},
								fail(res) {
									uni.showToast({ title: '获取个人信息失败', duration: 2000 })
								}
							});
						}
					}
				})
			},
			updateUserInfo (nickName, avatarUrl, gender) {
				let _this = this
				UpdateWxUserInfo({
					avatarUrl: avatarUrl,
					nickName: nickName,
					gender: gender,
					openId: uni.getStorageSync('OpenId'),
				}).then(res=>{
					if (res.code == 0) {
						console.log("------wxUserInfo:", res)
						wx.setStorageSync("UserInfo", res.data.user)
						// wx.setStorageSync("HadLogin", true)
						// wx.setStorageSync("Token", res.data.token)
						// wx.setStorageSync("TokenTime", (new Date()).getTime())
						
						uni.getStorage({
							key: 'UserInfo',
							success: (res) => {
								this.login(res.data);
							}
						});
						// 传递给父组件 关闭弹窗
						this.$emit('success', true);
						_this.getToken()
					}
					else {
						console.log(res);
					}
				});
			},
			getToken() {
				let _this = this
				wxRefreshLogin({openId: _this.$store.state.openId}).then(res => {
					if (res.code == 0) {
						const userinfo = res.data
						wx.setStorageSync("UserInfo", userinfo.user)
						this.login(userinfo.user);
						wx.setStorageSync("Token", userinfo.token)
						wx.setStorageSync("TokenTime", (new Date()).getTime())
						if (userinfo.user.phone.length == 11) {
							uni.setStorageSync('HadPhone', true)
						}
					}
				}).catch(errors => {
					console.log("------wxRefreshLogin---errors--------", errors)
				});
			},
			closePop() {
				// 传递给父组件 关闭弹窗
				this.$emit('close', true);
			},
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
  top: 30%;
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