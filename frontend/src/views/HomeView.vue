<script>
import TheWelcome from '../components/TheWelcome.vue'

export default {
  data() {
    return {
      // sudoku: new Array(81).fill(null)
      sudoku: new Array(81)
    }
  },
  methods: {
    getCellId: function(cellId) {
      return `cell${cellId}`
    },
    onkeydown: function(cellId, event) {
      console.log(event, event.key, event.keyCode)

      if (event.keyCode == 8) {
        console.log('Set cell to null', cellId)
        this.sudoku[cellId] = null
        return
      }
    },
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
      var key = event.key;
      // console.log(event, event.key, event.keyCode)

      // if (event.keyCode == 8) {
      //   this.sudoku[cellId] = null
      //   return
      // }
      console.log(key)
      if (!["1","2","3","4","5","6","7","8","9"].includes(key)) {
        event.preventDefault();
        console.log("PREVENT DEFAULT")
        return
      }
      

      
      this.sudoku[cellId] = parseInt(event.key)
      if (cellId < 80) {
        this.$refs.cell[cellId+1].focus()
      }

      console.log("END!")
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
        <td v-for="(v2, column) in 9"
          :key="column"
          :class="{
            'cell': true,
            'cell-border-right': (column == 2 || column == 5),
            'cell-border-left': (column == 3 || column == 6),
            'cell-border-bottom': (row == 2 || row == 5),
            'cell-border-top': (row == 3 || row == 6),
          }">
          <input
            type="text"
            class="button"
            @keydown="(event) => onkeydown(row * 9 + column, event)"
            @keypress.prevent="(event) => onkeypress(row * 9 + column, event)"
            v-model="sudoku[row * 9 + column]"
            ref="cell" />
        </td>
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