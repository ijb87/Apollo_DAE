var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue!'
  }
})

app.message = 'I have changed the data!'
app.message = 'Latest update!'

// vue reads the last line and updates the message object