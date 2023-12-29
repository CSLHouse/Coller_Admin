<template>
	<view class="content b-t">
		<view v-if="!hasLogin || empty===true" class="empty">
			<view v-if="hasLogin" class="empty-tips">
				空空如也
			</view>
			<view v-else class="empty-tips">
				空空如也
				<view class="navigator" @click="navToLogin">去登陆></view>
			</view>
		</view>
		<view class="list b-b" v-for="(item, index) in cardList" :key="index" @click="showDitail(item)">
			<view class="wrapper">
				<view class="address-box">
					<text class="address">{{item.storeName}}</text>
				</view>
				<view class="u-box">
					<text class="name">剩余次数</text>
					<text class="mobile">{{item.remainTimes}}</text>
				</view>
			</view>
			<text class="yticon icon-bianji" @click.stop="showDitail(item)"></text>
		</view>
		<view v-if='isShowDitail'>
			<div @click="closePop">
			    <!-- 需要粘贴的部分，替换掉【授权登录】按钮这一行代码 -->
			    <div class="modal-mask">
			    </div>
			    <div class="modal-dialog">
			      <div class="modal-content">
			        <div class="content-text">
			          <p class="key-bold-tip">{{selectRow.combo.comboName}}</p>
					  <p class="key-bold">套餐：{{selectRow.remainTimes}}</p>
			          <p class="key-bold">剩余次数：{{selectRow.remainTimes}}</p>
					  <p class="key-bold">办理时间：{{selectRow.combo.UpdatedAt}}</p>
					  <p class="key-bold">到期日期：{{selectRow.deadline}}</p>
			          <p class="little-tip">请在生效期间使用，过期将自动清零</p>
			          <!-- <p class="little-content">
			            注册成为会员，一店消费，多家优惠，欢迎体验
			          </p> -->
			        </div>
			      </div>
			    </div>
			</div>
		</view>
	</view>
</template>

<script>
	import { mapState } from 'vuex';
	import { GetMemberCardList } from '@/api/member.js';
	export default {
		data() {
			return {
				cardList: [],
				isShowDitail: false,
				selectRow: {},
				empty: false,
			}
		},
		onLoad(option) {
			this.loadData();
		},
		computed: {
			...mapState(['hasLogin','userInfo'])
		},
		methods: {
			async loadData() {
				let _this = this
				if (_this.hasLogin) {
					GetMemberCardList({onlyId: _this.$store.state.userInfo.telephone}).then(response => {
						this.cardList = response.data;
					});
				}
			},
			//查看详情
			showDitail(item) {
				this.isShowDitail = true
				this.selectRow = item
			},
			closePop() {
				this.isShowDitail = false
			},
			navToLogin() {
				uni.reLaunch  ({
					url: '/pages/user/user'
				})
			},
		},
		watch: {
			//显示空白页
			cardList(e) {
				let empty = e.length === 0 ? true : false;
				if (this.empty !== empty) {
					this.empty = empty;
				}
			}
		},
	}
</script>

<style lang='scss'>
	page {
		padding-bottom: 120upx;
	}
	
	/* 空白页 */
	.empty {
		position: fixed;
		left: 0;
		top: 0;
		width: 100%;
		height: 100vh;
		padding-bottom: 100upx;
		display: flex;
		justify-content: center;
		flex-direction: column;
		align-items: center;
		background: #fff;
	
		image {
			width: 240upx;
			height: 160upx;
			margin-bottom: 30upx;
		}
	
		.empty-tips {
			display: flex;
			font-size: $font-sm+2upx;
			color: $font-color-disabled;
	
			.navigator {
				color: $uni-color-primary;
				margin-left: 16upx;
			}
		}
	}
	
	.content {
		position: relative;
	}

	.list {
		display: flex;
		align-items: center;
		padding: 20upx 30upx;
		;
		background: #fff;
		position: relative;
	}

	.wrapper {
		display: flex;
		flex-direction: column;
		flex: 1;
	}

	.address-box {
		display: flex;
		align-items: center;

		.tag {
			font-size: 24upx;
			color: $base-color;
			margin-right: 10upx;
			background: #fffafb;
			border: 1px solid #ffb4c7;
			border-radius: 4upx;
			padding: 4upx 10upx;
			line-height: 1;
		}

		.address {
			font-size: 30upx;
			color: $font-color-dark;
		}
	}

	.u-box {
		font-size: 28upx;
		color: $font-color-light;
		margin-top: 16upx;

		.name {
			margin-right: 30upx;
		}
	}

	.icon-bianji {
		display: flex;
		align-items: center;
		height: 80upx;
		font-size: 40upx;
		color: $font-color-light;
		padding-left: 30upx;
	}
	
	.icon-iconfontshanchu1 {
		display: flex;
		align-items: center;
		height: 80upx;
		font-size: 40upx;
		color: $font-color-light;
		padding-left: 30upx;
	}

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
	  text-align: center;
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

