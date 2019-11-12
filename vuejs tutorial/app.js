new Vue({
  el:'#vue-app',
  data: {
    name: 'Shaun',
    job: 'Ninja'
  },
  methods: {
    greet: function(time) {
      this.job
      return 'Good' + time + ' ' + this.name;
    }
  }
});