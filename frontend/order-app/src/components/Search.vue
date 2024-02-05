<template>
  <b-container fluid>
    <b-row class="my-1">
      <b-col sm="3">
        <label for="input-search">Search</label>
      </b-col>
      <b-col sm="9">
        <b-form-input id="input-search" v-model="search" @change="emitInputChange"></b-form-input>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
  export default {
    data() {
      return {
        search: ''
      }
    },
    methods: {
      emitInputChange() {
        this.$emit('search-change', this.search);
      }
    }
  }
</script>