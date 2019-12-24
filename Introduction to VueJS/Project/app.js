Vue.component('name-item', {
  props: ['name'],
  template: '<li>{{ name.text }}</li>'
})

var app7 = new Vue({
  el: '#app-8',
  data: {
    nameList: [
      { id: 0, text: 'John' },
      { id: 1, text: 'Steve' },
      { id: 2, text: 'Albert' },
      { id: 3, text: 'Sam' },
      { id: 4, text: 'Max' },
      { id: 5, text: 'Ray' },
      { id: 6, text: 'Wendy' },
      { id: 7, text: 'Lauren' },
      { id: 8, text: 'Mary' },
      { id: 9, text: 'Frank' }
    ]
  }
})