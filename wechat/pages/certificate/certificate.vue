<template>
	<view class="container">
		<view class="card" v-if="certificateList.length > 0">
			<image class="bg" src="/static/bg.jpg"></image>
			<text class="card-title">互通卡</text>
			<text class="card-content" style="color: green;">凭此卡可在多个店铺享受至尊优惠!</text>
		</view>
		<view class="card" v-else>
			<image class="bg" src="/static/bg.jpg" style="opacity: 0.3;"></image>
			<text class="card-title" style="font-size: 60upx;color: black;">暂无</text>
			<text class="card-content">猪迪克，让孩子们远离电视，尽享欢乐！</text>
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
					if (_this.$store.state.userInfo && _this.$store.state.userInfo.telephone) {
						GetCertificateList({onlyId: _this.$store.state.userInfo.telephone}).then(response => {
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
		background: $page-color-base;
		height: 480upx;
		padding: 40upx 0upx 0;
		position:relative;
		border-radius: 0.5rem;
		display:flex;
		justify-content:center;
		.bg{
			display: block;
			margin: 0 auto;
			width: 700upx;
			height: 420upx;
			border-radius: 5%;
		}
	}
	.card-title {
		position:absolute;
		box-sizing: border-box;
		text-align: center;
		margin:0 auto;
		font-size: 70upx;
		color: red;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		opacity: 1;
		top: 20%;
	}
	.card-content {
		position:absolute;
		/* margin: 35% auto; */
		font-size:36upx;
		margin-top: 1;
		top: 70%;
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
