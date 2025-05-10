const { ref } = Vue;


export default {
  setup() {
    const grid = ref(Array(81).fill(""));

    const generate = async () => {
      const response = await fetch('http://localhost:8081/gen', { method: 'GET' });
      const data = await response.json();
      grid.value = data.Grid.split("").map(cell => cell === "0" ? "" : cell);
    };

    return {
      grid,
      generate,
    };

  }
}
