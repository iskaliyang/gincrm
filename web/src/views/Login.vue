<template>
  <a-form :model="formData" @finish="Login" @finishFailed="LoginFailed" style="width: 65%;">
    <a-form-item>
      <div class="title">账号登录</div>
    </a-form-item>
    <a-form-item name="email" :rules="rules.name">
      <a-input v-model:value="formData.email" size="large" placeholder="邮箱">
        <template #prefix>
          <UserOutlined class="site-form-item-icon" />
        </template>
      </a-input>
    </a-form-item>
    <a-form-item name="password" :rules="rules.password">
      <a-input-password v-model:value="formData.password" size="large" placeholder="密码">
        <template #prefix>
          <LockOutlined class="site-form-item-icon" />
        </template>
      </a-input-password>
    </a-form-item>
    <a-form-item>
      <a-form-item name="remember" no-style>
        <a-checkbox v-model:checked="formData.remember" style="float: left;">记住我</a-checkbox>
      </a-form-item>
      <a class="login-form-forgot" style="float: right;" @click="forgotPassword">忘记密码？</a>
    </a-form-item>
    <a-form-item>
      <a-button size="large" type="primary" html-type="submit" class="login-form-button" style="width: 100%">登录
      </a-button>
      <br />
    </a-form-item>
    <a-form-item>还没有账号？<a @click="toRegister"> 立即注册</a></a-form-item>
  </a-form>
</template>

<script setup>
import { reactive } from "vue"
import { useRouter } from "vue-router"
import { message } from 'ant-design-vue';
import { LockOutlined, UserOutlined } from '@ant-design/icons-vue';
import { userLogin } from "@/api/user.js";

const router = useRouter()

const rules = {
  name: [
    { required: true, message: '请输入邮箱!', trigger: "change" }
  ],
  password: [
    { required: true, message: '请输入密码!', trigger: "change" },
    { min: 6, max: 12, message: '密码长度需要在6~12之间', trigger: 'blur' },
  ],
}

// 表单数据
const formData = reactive({
  email: '3285844058@qq.com',
  password: '123456',
  remember: true,
})

// 用户登录
const Login = () => {
  let param = {
    email: formData.email,
    password: formData.password
  }
  userLogin(param).then((res) => {
    if (res.data.code === 0) {
      localStorage.setItem('uid', res.data.data.uid)
      localStorage.setItem('token', res.data.data.token)
      router.push("/home")
    }
    if (res.data.code === 10002) {
      message.error('用户不存在');
    }
    if (res.data.code === 10003) {
      message.error('用户名或密码错误');
    }
  })
};

const LoginFailed = msg => {
  console.log('登录失败: ', msg);
};

const toRegister = () => {
  router.push("/register")
}

const forgotPassword = () => {
  router.push("/password")
}

</script>

<style scoped>
.title {
  color: #303133;
  font-size: 30px;
  line-height: 65px;
  font-weight: 500;
}

.site-form-item-icon {
  color: rgba(0, 0, 0, .45);
}
</style>