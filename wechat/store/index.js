import Vue from 'vue'
import Vuex from 'vuex'
import {
		memberInfo,wxLogin, GetWxUserInfo
	} from '@/api/member.js';
	
Vue.use(Vuex)

const store = new Vuex.Store({
	state: {
		hasLogin: false,
		userInfo: {},
		openId: '',
		token: "",
		hadNickName: false,
	},
	mutations: {
		refreshLoginSession() {
			let _this = this
			uni.login({
				provider: 'weixin', //使用微信登录
				success: function (loginRes) {
					wxLogin({code: loginRes.code}).then(res => {
						console.log("------login---res--------", res)
						if (res.code == 0) {
							uni.setStorage({//缓存用户登陆状态
							    key: 'OpenId',  
							    data: res.data.openid
							})
							// wx.setStorageSync("WxCode", loginRes.code)
							// wx.setStorageSync("WxCodeTime", (new Date()).getTime())
							_this.state.openId = res.data.openid
							console.log("-[refreshLoginSession]-openId-", _this.state.openId)
						} else {
							uni.showToast({
								title: res.data,
								icon: 'none'
							})
						}
					}).catch(errors => {
						console.log("------login---errors--------", errors)
					});
				}
			});
		},
		login(state, provider) {
			state.hasLogin = true;
			state.userInfo = provider;
			uni.setStorage({//缓存用户登陆状态
			    key: 'UserInfo',  
			    data: provider
			})
		},
		logout(state) {
			state.hasLogin = false;
			state.userInfo = {};
			uni.removeStorage({  
                key: 'UserInfo'  
            });
			uni.removeStorage({
			    key: 'Token'  
			})
		}
	},
	actions: {
	
	}
})

export default store
