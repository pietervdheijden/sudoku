<script setup>
import TheWelcome from '../components/TheWelcome.vue'
</script>

<template>
  <main>
    <!-- <TheWelcome /> -->

    <br/>

    <table class="center">
      <tr v-for="row in 9" :key="row">
        <td v-for="column in 9" :key="column">
          {{ row }} {{ column }}
        </td>
      </tr>
    </table>

    <span v-for="n in 10" :key="n">{{ n }}</span>

    <p>Buttons:</p>

    <button>Solve</button>
    <button>Solve next cell</button>
    <button>Hint</button>
    <button>Check</button>
    <button>Show options / candidates</button>
  </main>
</template>

<style>
.center {
  margin-left: auto;
  margin-right: auto;
}
</style>