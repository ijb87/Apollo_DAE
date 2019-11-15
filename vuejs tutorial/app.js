new Vue({
    el: '#vue-app',
    data: {
        health: 100,
        ended: false,
        turnRed: false
    },
    methods: {
        punch: function(){
            this.health -= 10;
            if ( this.health <= 0 ){
                this.ended = true;
                window.alert("The game is finished. Start again by closing this alert.")
                this.ended = true;
            }
        },
        restart: function(){
            this.health = 100;
            this.ended = false;
        },
        playSound (sound) {
          if (sound) {
            var audio = new Audio(sound);
            audio.play();
          }
        }
    },
    computed: {

    }
});
