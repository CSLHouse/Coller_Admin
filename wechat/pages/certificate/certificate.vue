<template>
	<view class="container">
		<view class="card" v-if="certificateList.length > 0">
			<image class="bg" src="/static/bg.jpg"></image>
			<text class="card-title">共通卡</text>
			<text class="card-content">凭此卡可在多个店铺享受至尊优惠</text>
		</view>
		<view v-else>
			<p>填充一张普通卡</p>
		</view>
		<view class="list b-b" v-for="(item, index) in certificateList" >
			<view class="wrapper" >
				<view class="address-box">
					<text class="address">{{item.storeName}}</text>
				</view>
			</view>
			<text class="count" :key="index" :style="{color: item.color}">{{item.count}}</text>
		</view>
		
	</view>
</template>

<script>
	import { mapState } from 'vuex';
	import { GetCertificateList } from '@/api/member.js';
	export default {
		data() {
			return {
				certificateList: []
			}
		},
		onLoad() {
			this.loadData();
		},
		computed: {
			...mapState(['hasLogin','userInfo'])
		},
		methods: {
			async loadData() {
				let _this = this
				if (_this.hasLogin) {
					console.log("-------_this.$store.state.openId------", _this.$store.state.userInfo)
					if (_this.$store.state.userInfo && _this.$store.state.userInfo.telephone) {
						console.log("-----openId---", _this.$store.state.openId)
						GetCertificateList({onlyId: _this.$store.state.userInfo.telephone}).then(response => {
							console.log("--response--", response)
							this.certificateList = response.data;
							this.certificateList.forEach(item => {
								if (item.isFirst) {
									item.color = "red";
								}
							})
						});
					}
					
				} else {
					uni.showToast({ title: '请先登录', duration: 2000 })
				}
			},
		}
	}
</script>

<style lang='scss'>
	.card {
		height: 520upx;
		padding: 90upx 30upx 0;
		position:relative;
		border-radius: 0.5rem;
		display:flex;
		justify-content:center;
		.bg{
			position:absolute;
			left: 0;
			top: 0;
			width: 100%;
			height: 100%;
			opacity: .8;
			/* background-color: yellow; */
		}
	}
	.card-title {
		position:absolute;
		box-sizing: border-box;
		text-align: center;
		margin:0 auto;
		font-size: 80upx;
		color: red;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		opacity: 1;
	}
	.card-content {
		position:absolute;
		margin: 35% auto;
		font-size:40upx;
		margin-top: 1;
	}
	.list {
		display: flex;
		align-items: center;
		padding: 20upx 30upx;
		margin-top: 20px;
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
			font-size: 40upx;
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
	.count {
		display: flex;
		align-items: center;
		/* height: 80upx; */
		font-size: 40upx;
		color: $font-color-light;
		padding-left: 30upx;
		border-radius: 0.5rem;
		/* color: red; */
	}
</style>
