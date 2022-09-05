<script>
import TheWelcome from '../components/TheWelcome.vue'

export default {
  data() {
    return {
      // sudoku: new Array(81).fill(null)
      sudoku: new Array(81).fill(null)
    }
  },
  methods: {
    onkeypress: function (cellId, event) {
      console.log(cellId, event)

      // Source: https://www.geeksforgeeks.org/how-to-force-input-field-to-enter-numbers-only-using-javascript/
      // Only ASCII character in that range allowed
      // var ASCIICode = (event.which) ? event.which : event.keyCode
      // if (ASCIICode > 31 && (ASCIICode < 48 || ASCIICode > 57)) {
      //   event.preventDefault();
      //   return false
      // }

      // TODO: support backspace and delete keys
      // TODO: auto tab when number is filled

      var key = event.key;
      if (!["0", "1","2","3","4","5","6","7","8","9"].includes(key)) {
        event.preventDefault();
      } else {
        this.sudoku[cellId] = parseInt(event.key)
      }
    },
    async solve() {
      console.log("SOLVE!")
      console.log([...this.sudoku])
      
      try {
        const res = await fetch('http://localhost:8080/solve', { // TODO: get URL from env var
          method: "POST",
          headers: {"Content-Type": "application/json"},
          body: JSON.stringify({'sudoku': this.sudoku.map(i => i ?? 0)}),
          mode: 'no-cors' // TODO: remove
        });
        var answer = (await res.json())
      } catch (error) {
        var answer = 'Error! Could not reach the API. ' + error
      }

      console.log(answer)
    }
  }
}
</script>

<template>
  <main style="margin-left: auto margin-right: auto">
    <table class="table">
      <tr v-for="(v1, row) in 9" :key="row">
        <!-- <td v-for="(v2, column) in 9" :key="column" class="cell">
          <span>{{ row * 9 + column }}</span>
        </td> -->
        <td v-for="(v2, column) in 9" :key="column" :set="cellId = row * 9 + column"
          :class="{
            'cell': true,
            'cell-border-right': (column == 2 || column == 5),
            'cell-border-left': (column == 3 || column == 6),
            'cell-border-bottom': (row == 2 || row == 5),
            'cell-border-top': (row == 3 || row == 6),
          }">
          <input type="text" class="button" @keypress.prevent="(event) => onkeypress(row * 9 + column, event)" v-model="sudoku[cellId]"/>
        </td>

        <!-- <td v-for="column in 9" :key="column" class="cell">
          <input type="text" class="button" @keypress.prevent="(event) => onkeypress(row * 9 + column, event)" v-model="sudoku[row * 9 + column]"/>
        </td> -->
      </tr>
    </table>

    <button @click="solve()">Solve</button>

    <!-- <p><button>Solve</button></p>
    <p></p>
    <button>Solve next cell</button>
    <p></p>
    <button>Hint</button>
    <p></p>
    <button>Check</button>
    <p></p>
    <button>Show options / candidates</button> -->
  </main>
</template>

<style>
.table {
  margin-left: auto;
  margin-right: auto;
  border: 2px solid;
  border-spacing: 0;
}
.cell {
  padding: 0px;
}
.button {
  width: fit-content;
  display: inline-block;
  width: 40px;
  height: 40px;
  background: lightsteelblue;
  text-align: center;
  border: 1px solid;
  outline: none;
}
.cell-border-right {
  border-right: 1px solid;
}
.cell-border-left {
  border-left: 1px solid;
}
.cell-border-bottom {
  border-bottom: 1px solid;
}
.cell-border-top {
  border-top: 1px solid;
}
</style>