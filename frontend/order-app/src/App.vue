<template>
  <div id="app">
    <OrderIndex/>
  </div>
</template>

<script>
import OrderIndex from './pages/OrderIndex.vue'

export default {
  name: 'App',
  components: {
    OrderIndex
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
