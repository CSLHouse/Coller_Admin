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
		hasPhone: false,
		openId: ''
	},
	mutations: {
		refreshLoginSession() {
			let _this = this
			uni.login({
				provider: 'weixin', //使用微信登录
				success: function (loginRes) {
					wxLogin({code: loginRes.code}).then(res => {
						console.log("------login---res--------", res.data)
						if (res.code == 0) {
							uni.setStorageSync('OpenId', res.data.openid);
							// wx.setStorageSync("WxCode", loginRes.code)
							// wx.setStorageSync("WxCodeTime", (new Date()).getTime())
							_this.state.openId = res.data.openid
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
			    key: 'userInfo',  
			    data: provider  
			})
			state.openId = provider.openid
		},
		setPhone(state, provider) {
			state.hasPhone = provider;
		},
		logout(state) {
			state.hasLogin = false;
			state.hasPhone = false;
			state.userInfo = {};
			uni.removeStorage({  
                key: 'userInfo'  
            });
			uni.removeStorage({
			    key: 'token'  
			})
		}
	},
	actions: {
	
	}
})

export default store
