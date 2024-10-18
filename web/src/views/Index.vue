<template>
  <div class="container">
    <a-row>
      <a-col :span="14">
        <div class="back"></div>
      </a-col>
      <a-col :span="10">
        <div class="form">

          <router-view v-slot="{ Component }">
            <transition name="fade">
              <component :is="Component" />
            </transition>
          </router-view>

        </div>
      </a-col>
    </a-row>
  </div>
</template>

<script>
export default {
  name: "Index"
}
</script>

<style scoped>
.container {
  width: 100%;
  height: 100vh;
  background: #FFFFFF;
}

.back {
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f9f9fb url('../assets/crm.png') no-repeat center;
}

.form {
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.links a {
  color: rgba(0, 0, 0, .45);
  padding: 5px 20px;
}

.links a:hover {
  color: #5b5b5c;
}
</style>