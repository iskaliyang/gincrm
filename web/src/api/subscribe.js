import request from "../axios/index";

export function getSubscribeInfo(param) {
    return request({
        url: '/api/subscribe/info',
        method: 'get',
        params: param,
    })
}

export function subscribePay(param) {
    return request({
        url: '/api/subscribe/pay',
        method: 'post',
        data: param,
    })
}