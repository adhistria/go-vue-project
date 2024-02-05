<template>
  <div class="container">
    <Search v-on:search-change="onSearchChange"/>
    <DatePicker v-on:start-date-change="onStartDateChange" v-on:end-date-change="onEndDateChange"/>
    <p> Total Amount: {{total_amount}} </p>
    <Table v-bind:items="items"/>
    
    <b-row class="justify-content-md-center" v-if="isShowPagination">
        <b-pagination
          v-model="currentPage"
          :total-rows="rows"
          :per-page="perPage"
          @input="fetchData"
          aria-controls="my-table"
        ></b-pagination>
    </b-row>

  

  </div>
</template>

<script>
import DatePicker from '../components/DatePicker.vue'
import Table from '../components/Table.vue'
import Search from '../components/Search.vue'
import axios from 'axios';
export default {
  name: 'OrderIndex',
  components: {
    DatePicker,
    Table,
    Search
  },
  data() {
      return {
        items: [],
        total_amount: 0,
        search: '',
        startDate: '',
        endDate: '',
        currentPage: 1,
        rows: 0,
        perPage: 5,
        isShowPagination: false,
      }
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      let base_url = "http://localhost:8080/orders" 
      axios.get(base_url, {params: {search: this.search, page: this.currentPage, start_date: this.startDate, end_date: this.endDate}})
        .then(response => {
          this.items = response.data.data;
          this.rows = response.data.total_rows;
          this.total_amount = response.data.total_amount;
          if (this.rows > 0) {
            this.isShowPagination = true
          }else{
            this.isShowPagination = false
          }
        })
        .catch(error => {
          console.error("Error fetching data", error)
        })
    },
    onSearchChange(search) {
      this.search = search
      this.fetchData()
    },
    onStartDateChange(startDate) {
      this.startDate = new Date(startDate).toISOString()
      this.fetchData()
    },
    onEndDateChange(endDate) {
      this.endDate = new Date(endDate).toISOString()
      this.fetchData()
    }
    
  }
}
</script>