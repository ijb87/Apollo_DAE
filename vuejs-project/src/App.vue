<template>
  <div id="app">
    <meta charset="utf-8"><div class="top-container">
     <img class="top-cloud" src="img/cloud.png" alt="cloud-img">
     <h1>I'm Ihab.</h1>
     <p id="subtitle">a programmer.</p>
     <img class="bottom-cloud" src="img/cloud.png" alt="cloud-img">
     <img class="mountain" src="img/mountain.png" alt="mountain-img">
   </div>
<div class="middle-container">
  <div class="profile">
    <img src="img/profile.jpg" alt="Ihab Baghdadi Profile Picture" style="height: 150px; margin-top: 50px;">
    <h2>Hello.</h2>
    <p class="intro">Diligent and tech-savvy graduate with strong coding and web development skills. Problem solver with ability to collaborate with others to come up with solutions, address performance concerns, and debug issues. Strong time management skills, able to meet deadlines, and sees all projects through to completion. Coachable individual, who wants to grow professionally, continue to learn, and be a strong contributor.</p>
  </div>
  <hr>
  <!-- Skills -->
  <Skills></Skills>
  <hr>
  <div class="contact-me">
    <h2>Get In Touch</h2>
    <h3>Got a question for me? I'm all ears!</h3>
    <p class="contact-message">Send me an email and I will get back to you as soon as possible!</p>
    <a class="btn" href="ihabbaghdadi@live.com">CONTACT ME</a>
  </div>
</div>


<div class="bottom-container">
  <a class="footer-link" href="https://www.linkedin.com/in/ihab-baghdadi-web-developer/">LinkedIn</a>
  <a class="footer-link" href="https://ijb87.github.io/My-Website/">Website</a>
  <p class="copyright">© 2019 Ihab Baghdadi.</p>
</div>
</div>
</template>

<script>

import Skills from './components/Skills'

export default {
  name: 'App',
  components: {
    Skills
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
}
</style>
