import request from '../axios/index'

// 用户登录
export function userLogin(param) {
  return request({
    url: '/api/user/login',
    method: 'post',
    data: param,
  })
}

// 用户注册
export function userRegister(param) {
  return request({
    url: '/api/user/register',
    method: 'post',
    data: param,
  })
}

// 获取验证码
export function getVerifyCode(param) {
  return request({
    url: '/api/user/verifycode',
    method: 'get',
    params: param,
  })
}

// 忘记密码
export function userForgotPass(param) {
  return request({
    url: '/api/user/pass',
    method: 'post',
    data: param,
  })
}

// 注销账号
export function userDelete(param) {
  return request({
    url: '/api/user/delete',
    method: 'delete',
    data: param,
  })
}

// 获取用户信息
export function getUserInfo(param) {
  return request({
    url: '/api/user/info',
    method: 'get',
    params: param,
  })
}
