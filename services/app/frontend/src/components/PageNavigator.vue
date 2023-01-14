<template>
  <h1>{{ title }}</h1>
  <div class="card">
    <div style="width:400px;height:130px;margin-top:20px;border-style: dotted;">
      <br>
      <span>Firstmame: {{ firstname }}</span> <br>
      <span>Age: {{ age }}</span> <br>
    </div><br><br>
    <label> Enter Firstname </label><br>
    <input
      v-model="search"
      type="text"
      style="font-size:20px;border-radius:10px;"
      placeholder=" Name ..."
    > <br> <br>
    <button
      type="button"
      @click="getAge"
    >
      Guess Age
    </button>
    <br> <br> <br>
    <input
      type="radio"
      value="pop"
    >
    <label>Save my data</label>
  </div>
</template>
<script setup>
defineProps({
  title: {
    type: String,
    required: true
  }
})
</script>
<script>
export default {
  data() {
    return {
      search:"",
      firstname:"",
      age:"",
    }
  },
  computed: {
    getAge() {
      return fetch('https://api.agify.io/?name='+ this.search)
          .then(response => response.json())
          .then(data => {
            this.age = data.age
            this.firstname = data.name
            this.search=""
          })
    }
  }
}
</script>