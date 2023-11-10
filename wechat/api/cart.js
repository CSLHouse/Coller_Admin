import request from '@/utils/requestUtil'

export function addCartItem(data) {
	return request({
		method: 'POST',
		url: '/product/cart',
		data: data
	})
}

export function fetchCartList() {
	return request({
		method: 'GET',
		url: '/product/cart/list'
	})
}

export function deletCartItem(params) {
	return request({
		method: 'DELETE',
		url: '/product/cart',
		params:params
	})
}

export function deletCartItemWithList(params) {
	return request({
		method: 'DELETE',
		url: '/product/carts',
		params:params
	})
}

export function updateQuantity(params) {
	return request({
		method: 'PUT',
		url: '/product/cart',
		params:params
	})
}

export function clearCartList() {
	return request({
		method: 'POST',
		url: '/cart/clear'
	})
}