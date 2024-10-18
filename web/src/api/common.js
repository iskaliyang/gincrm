import request from '../axios/index'

// 文件移除
export function fileRemove(param) {
    return request({
        url: '/api/common/file/remove',
        method: 'delete',
        data: param,
    })
}