const { ref } = Vue;


export default {
  setup() {
    const grid = ref(Array(81).fill(""));
    const status = ref("");

    const clear = () => {
      grid.value = Array(81).fill("");
      status.value = "";
    };

    const generate = async () => {
      const response = await fetch('http://localhost:8081/gen', { method: 'GET' });
      const data = await response.json();
      grid.value = data.Grid.split("").map(cell => cell === "0" ? "" : cell);
    };

    const solve = async () => {
      const gridString = grid.value.map(cell => cell === "" ? "0" : cell).join("");
      const response = await fetch(`http://localhost:8081/solve?input=${gridString}`, {
        method: 'GET',
      });
      const data = await response.json();
      if (data.Code === 200) {
        grid.value = data.Solutions[0].split("").map(cell => cell === "0" ? "" : cell);
        status.value = `Found ${data.Solutions.length} solutions`;
      } else {
        status.value = "Error: " + data.Message;
      }
    }

    const onInput = (index) => {
      const value = grid.value[index];
      if (!/^[1-9]$/.test(value)) {
        grid.value[index] = "";
      }
    }

    return {
      grid,
      status,
      clear,
      generate,
      solve,
      onInput,
    };

  }
}
