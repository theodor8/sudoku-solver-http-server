const { ref, watch } = Vue;


export default {
  setup() {
    const grid = ref(Array(81).fill(""));
    const status = ref("");
    const generateUnknowns = ref(30);
    const server = ref(`${window.location.hostname}:8081`);
    const importExportString = ref("");

    watch(generateUnknowns, (newValue) => {
      generateUnknowns.value = Math.max(0, Math.min(81, newValue));
    });

    const clear = () => {
      if (window.confirm("Are you sure you want to clear the grid?")) {
        grid.value = Array(81).fill("");
        status.value = "";
      }
    };

    const generate = async () => {
      if (generateUnknowns.value === "") {
        generateUnknowns.value = 30;
      }
      try {
        const response = await fetch(`http://${server.value}/gen?unknowns=${generateUnknowns.value}`, { method: 'GET' });
        const data = await response.json();
        if (!response.ok) {
          throw new Error(data.Message);
        }
        grid.value = data.Grid.split("").map(cell => cell === "0" ? "" : cell);
        status.value = "Grid generated successfully!";
      } catch (error) {
        status.value = "Error: " + error.message;
      }
    };

    const solve = async () => {
      const gridString = grid.value.map(cell => cell === "" ? "0" : cell).join("");
      try {
        const response = await fetch(`http://${server.value}/solve?input=${gridString}`, {
          method: 'GET',
        });
        const data = await response.json();
        if (!response.ok) {
          throw new Error(data.Message);
        }
        grid.value = data.Solutions[0].split("").map(cell => cell === "0" ? "" : cell);
        status.value = `Found ${data.Solutions.length} solutions`;
      } catch (error) {
        status.value = "Error: " + error.message;
      }
    }

    const onInput = (index, event) => {
      let value = grid.value[index];
      // Only keep the last entered digit if multiple are input
      if (typeof value === "string" && value.length > 1) {
        value = value.slice(-1);
        if (/^[1-9]$/.test(value)) {
          grid.value[index] = value;
        }
      }
      grid.value[index] = grid.value[index].replace(/[^1-9]/g, "");

      // Arrow key navigation
      if (event && event.type === "keydown") {
        let nextIndex = null;
        if (event.key === "ArrowRight") nextIndex = (index + 1) % 81;
        else if (event.key === "ArrowLeft") nextIndex = (index - 1 + 81) % 81;
        else if (event.key === "ArrowDown") nextIndex = (index + 9) % 81;
        else if (event.key === "ArrowUp") nextIndex = (index - 9 + 81) % 81;
        if (nextIndex !== null) {
          event.preventDefault();
          const nextInput = document.querySelectorAll('.cell')[nextIndex];
          if (nextInput) nextInput.focus();
        }
      }
    };

    const exportGrid = async () => {
      const gridString = grid.value.map(cell => cell === "" ? "0" : cell).join("");
      importExportString.value = gridString;
      status.value = "Grid exported.";
    };

    const importGrid = () => {
      const inputString = importExportString.value.trim();
      if (inputString.length !== 81 || !/^[0-9]{81}$/.test(inputString)) {
        status.value = "Invalid import string. Must be 81 digits (0-9).";
        return;
      }
      grid.value = inputString.split("").map(cell => cell === "0" ? "" : cell);
      status.value = "Grid imported successfully!";
    };



    return {
      grid,
      status,
      generateUnknowns,
      server,
      importExportString,

      clear,
      generate,
      solve,
      onInput,
      exportGrid,
      importGrid,
    };

  }
}
