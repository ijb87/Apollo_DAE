<template>
  <div style="margin-top: 50px;">
    <p v-if="currentUsername == null" style="margin-left: 50px; font-family: 'IBM Plex Serif', serif; font-size: 20px">
      Enter a username above to see their Github data...
    </p>
    <p v-else style="margin-left: 50px; font-family: 'IBM Plex Serif', serif; font-size: 20px;">
      Below are the results for {{ currentUsername }}.
      <github-user-data :data="githubData[currentUsername]"></github-user-data>
    </p>
  </div>
</template>
<script src="./script.js"></script>
<style scoped src="./style.css"></style>
