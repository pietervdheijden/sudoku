<script>
export default {
  data() {
    return {
      backendUrl: import.meta.env.VITE_BACKEND_URL,
      infoMessage: "",
      errorMessage: "",
      sudoku: new Array(81),
      // Test code:
      // sudoku: "000920000040851000256003091100085409098730162000200530007060900900002680080090054".split('').map(c => {
      //   var value = parseInt(c)
      //   if (value == 0) {
      //     return null
      //   } else {
      //     return value
      //   }
      // }),
      options: null
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
      if(event.shiftKey && event.keyCode == 9) { // shift tab
        if (cellId > 0) {
          this.$refs.cell[cellId-1].focus();
        }
        return;
      }
      if (event.keyCode == 9) { // tab
        if (cellId < 80) {
          this.$refs.cell[cellId+1].focus();
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
      console.log(`Solve sudoku: ${this.sudoku.join('')}`)
      this.infoMessage = "";
      this.errorMessage = "";

      try {
        const response = await fetch(`${this.backendUrl}/api/v1/solve`, {
          method: "POST",
          headers: {"Content-Type": "application/json"},
          body: JSON.stringify({'sudoku': this.sudoku.map(i => i ?? 0)}),
        });
        const statusCode = response.status;
        const json = (await response.json())
        if (statusCode == 200) {
          this.sudoku = json.sudoku;
          this.infoMessage = "Sudoku is solved!"
        } else {
          this.errorMessage = json.message
        }
      } catch (error) {
        this.errorMessage = `API returned error: ${error}` 
      }
    },
    async hint() {
      console.log("Hint")
      this.infoMessage = "Hint"
    },
    async checkSolution() {
      console.log("Check solution")
      this.infoMessage = "Check solution"
    },
    async showOptions() {
      console.log("Show options")
      this.infoMessage = "Show options"
    }
  }
}
</script>

<template>
  <main style="margin-left: auto margin-right: auto text-align: center">
    <table class="table">
      <tr v-for="(_, row) in 9" :key="row">
        <td v-for="(_, column) in 9"
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
            class="cell-input"
            @keydown.prevent="(event) => onkeydown(row * 9 + column, event)"
            v-model="sudoku[row * 9 + column]"
            ref="cell" />
        </td>
      </tr>
    </table>

    <button @click="solve()" class="button">Solve</button>
    <!--
      TODO: implement buttons
    <button @click="hint()" class="button">Hint</button>
    <button @click="checkSolution()" class="button">Check solution</button>
    <button @click="showOptions()" class="button">Show options</button>
    -->
    
    <span class="info-box">{{ infoMessage }}</span>
    <span class="error-box">{{ errorMessage }}</span>
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
.cell-input {
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
.info-box {
  margin: auto;
}
.error-box {
  color: red;
}
.button {
  background-color: blueviolet;
  margin: auto;
  padding: 10px;
  display: block;
}
</style>