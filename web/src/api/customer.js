import request from '../axios/index'

// 查询客户信息
export function queryCustomerInfo(param) {
    return request({
        url: '/api/customer/info',
        method: 'get',
        params: param,
    })
}

// 查询客户列表
export function queryCustomerList(param) {
    return request({
        url: '/api/customer/list',
        method: 'get',
        params: param,
    })
}

// 删除客户
export function deleteCustomer(param) {
    return request({
        url: '/api/customer/delete',
        method: 'delete',
        data: param,
    })
}

// 导出Excel表格
export function customerExport(param) {
    return request({
        url: '/api/customer/export',
        method: 'get',
        responseType: 'blob',
        params: param,
    })
}

// 新建客户
export function createCustomer(param) {
    return request({
        url: '/api/customer/create',
        method: 'post',
        data: param,
    })
}

// 更新客户
export function updateCustomer(param) {
    return request({
        url: '/api/customer/update',
        method: 'put',
        data: param,
    })
}

// 发送邮件给客户
export function sendMailToCustomer(param) {
    return request({
        url: '/api/customer/send',
        method: 'post',
        data: param,
    })
}

export function queryCustomerOption(param) {
    return request({
        url: '/api/customer/option',
        method: 'get',
        params: param,
    })
}