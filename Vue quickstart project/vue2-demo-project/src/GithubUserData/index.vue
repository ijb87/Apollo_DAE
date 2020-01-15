<template>
  <div v-if="data">
    <h4>Name: {{ data.name }}</h4>
    <p>Company: {{ data.company }}</p>
    <p>Number of repos: {{ data.public_repos }}</p>
  </div>
</template>
<script src="./script.js"></script>
<style scoped src="./style.css"></style>
