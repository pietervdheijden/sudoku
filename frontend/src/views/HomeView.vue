<script>
export default {
  data() {
    return {
      sudoku: new Array(81)
    }
  },
  methods: {
    onkeydown: function(cellId, event) {
      console.log("onkeydown", event, event.key, event.keyCode, cellId)
      
      var key = event.key
      if (event.keyCode == 8) { // backspace
        console.log('Set cell to null', cellId)
        this.sudoku[cellId] = null
        return
      }
      if (event.keyCode == 37) { // arrow left
        if (cellId > 0) {
          this.$refs.cell[cellId-1].focus();
        }
        return;
      }
      if (event.keyCode == 39) { // arrow right
        if (cellId < 80) {
          this.$refs.cell[cellId+1].focus();
        }
        return;
      }
      if (event.keyCode == 38) { // arrow up
        if (cellId > 8) {
          this.$refs.cell[cellId-9].focus();
        }
        return;
      }
      if (event.keyCode == 40) { // arrow down
        if (cellId < 72) {
          this.$refs.cell[cellId+9].focus();
        }
        return;
      }
      if (event.keyCode == 9) { // tab
        if (cellId < 80) {
          this.$refs.cell[cellId+1].focus();
        }
        return;
      }
      if(event.shiftKey && event.keyCode == 9) { // shift tab
        if (cellId > 0) {
          this.$refs.cell[cellId-1].focus();
        }
        return;
      }
      if (!["1","2","3","4","5","6","7","8","9"].includes(key)) { // invalid char
        event.preventDefault();
        return;
      }

      // Valid input
      this.sudoku[cellId] = parseInt(event.key)
      if (cellId < 80) {
        this.$refs.cell[cellId+1].focus()
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
          mode: 'no-cors' // TODO: remove / replace (?)
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
            @keydown.prevent="(event) => onkeydown(row * 9 + column, event)"
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