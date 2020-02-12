<template>
  <div class="skills">
    <h2>My Skills.</h2>
    <div class="skill-row" v-for="(obj, key) in skillsJSON" :key="key">
      <h3>~ {{ obj.subtitle }}</h3>
      <p>{{ obj.paragraph }}</p>
    </div>
    <!-- <div class="skill-row">
      <img class="skill-image" src="img/speech.png" alt="watch image">
      <h3>Strong Communicator</h3>
      <p>I've gotten certifications on Lynda.com (LinkedIn Learning) regarding writing and verbal communication. I learned a lot of ways to communicate effectively with others, including being assertive.</p>
    </div> -->
  </div>
</template>

<script>
  export default {
    data: () => ({
      skillsJSON: [
        {
          subtitle: 'Coding',
          paragraph: 'I learned basic and intermediate HTML, CSS3 and JavaScript in university. I have a certificate on Responsive Web Design from FreeCodeCamp.org and one on Web Programming Foundations from Lynda.com (LinkedIn Learning). I also ranked in the top 25% of those who took the TestDome certification test on HTML, CSS and JavaScript.'
        },
        {
          subtitle: 'Strong Communicator',
          paragraph: "I've gotten certifications on Lynda.com (LinkedIn Learning) regarding writing and verbal communication. I learned a lot of ways to communicate effectively with others, including being assertive."
        }
      ]
    })
  }
</script>