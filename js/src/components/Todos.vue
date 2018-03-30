<template>
  <div class="todos">
    <h1>TODOS</h1>
    <ul>
      <li v-for="todo in todos" :key="todo.id">
        {{ todo.name }}
      </li>
    </ul>
    <form v-on:submit.prevent="onSubmit">
      <input v-model="todo" placeholder="Add TODO" />
    </form>
  </div>
</template>

<script>
export default {
  name: 'Todos',
  data: { todo: null },
  computed: {
    todos () {
      return this.$store.state.todos
    }
  },
  methods: {
    onSubmit () {
      fetch('http://localhost:9000/api/todos', {
          headers: {
            'content-type': 'application/json'
          },
          method: 'POST',
          body: JSON.stringify({
            name: this.todo
          })
        })
        .then(response => response.json())
        .then(data => this.$store.commit('setTodos', data))
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
a {
  color: #42b983;
}
</style>
