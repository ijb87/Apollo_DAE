<template>
  <form v-on:submit.prevent="onSubmit" style="margin-left: 50px; margin-top: 50px;">
    <input class="input is-rounded" type="text" v-model="username" placeholder="Enter a github username here..." style="width: 300px;"  />
    <button class="button is-dark is-small" type="submit" style="width: 50px; top: 7px; left: 3px;">Go!</button>
  </form>
</template>

<script src="./script.js"></script>
<style scoped src="./style.css"></style>
