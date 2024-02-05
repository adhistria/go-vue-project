<template>
  <b-row class="my-1">
      <b-col sm=3>
        <b-form-datepicker v-model="startDate" :max="endDate" @input="emitStartDateChange" placeholder="StartDate" class="mb-2"></b-form-datepicker>
      </b-col>
      <b-col sm=3>
        <b-form-datepicker v-model="endDate" :min="startDate" @input="emitEndDateChange" placeholder="EndDate" class="mb-2"></b-form-datepicker>
      </b-col>
    </b-row>
</template>

<script>
  export default {
    data() {
      return {
        startDate: '',
        endDate: '',
      }
    }, 
    methods: {
      emitStartDateChange() {
        this.$emit('start-date-change', this.startDate);
      },
      emitEndDateChange() {
        this.$emit('end-date-change', this.endDate);
      }
    }
  }
</script>