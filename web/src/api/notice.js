import request from '../axios/index'

// 获取通知列表
export function updateNotice(param) {
  return request({
    url: '/api/notice/update',
    method: 'put',
    data: param,
  })
}

// 获取通知列表
export function getNoticeCount(param) {
  return request({
    url: '/api/notice/count',
    method: 'get',
    params: param,
  })
}

// 删除全部通知
export function deleteNotice(param) {
  return request({
    url: '/api/notice/delete',
    method: 'delete',
    data: param,
  })
}

// 获取通知列表
export function getNoticeList(param) {
  return request({
    url: '/api/notice/list',
    method: 'get',
    params: param,
  })
}
