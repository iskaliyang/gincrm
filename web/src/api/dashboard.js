import request from "../axios/index";

export function getSummary(param) {
    return request({
        url: '/api/dashboard/summary',
        method: 'get',
        params: param,
    })
}
