<template>
  <div class="todos">
    <h1>Todos</h1>
    <p>
      Type a new todo in the the box below and press enter to add it to the list. To delete a todo, simply double click on it.
      </p>
      <!-- added an explanation of how to use the todo list -->
    <form v-on:submit="submitForm">
      <input type="text" v-model="text" />
      </form>
    <ul>
      <li v-for="todo in todos">
        <input class="toggle" type="checkbox" v-model="todo.completed" />
        <span :class="{completed: todo.completed}">
          <label @dblclick="deleteTodo(todo)">
            {{ todo.text }}
          </label>
        </span>
        </li>
      </ul>
  </div>
</template>

<script>
export default {
  name: 'todos',
  data () {
    return {
      text: '',
      todos: [
        {
        text: 'My Todo One',
        completed: false
      },
      {
      text: 'My Todo Two',
      completed: false
      },
      {
      text: 'My Todo Three',
      completed: false
      }
      ]
    }
  },
  methods: {
    submitForm(e) {
      this.todos.push({
        text: this.text,
        completed: false
      })
      e.preventDefault()
    },
    deleteTodo(todo){
      this.todos.splice(this.todos.indexOf(todo), 1)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  ul {
    list-style: none;
    padding: 0;
  }

  li {
    padding: 10px 10px;
    border-bottom: 1px #ccc solid;
  }

  .completed {
    text-decoration: line-through;
  }
</style>
