<template>
  <div>
    <h1>{{ title }}</h1>
    <p>{{ greeting() }}</p>
  </div>
</template>

<script>
export default {
  data () {
    return {
    title: 'Your first Vue file, woooo!'
    }
  },
  methods: {
    greeting: function(){
    return 'heeeeey cowbow!'
    }
  }
}
</script>

<style>
</style>
