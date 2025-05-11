const { ref, watch } = Vue;


export default {
  setup() {
    const grid = ref(Array(81).fill(""));
    const status = ref("");
    const generateUnknowns = ref(30);
    const server = ref("127.0.0.1:8081");

    watch(generateUnknowns, (newValue) => {
      generateUnknowns.value = Math.max(0, Math.min(81, newValue));
    });

    const clear = () => {
      grid.value = Array(81).fill("");
      status.value = "";
    };

    const generate = async () => {
      if (generateUnknowns.value === "") {
        generateUnknowns.value = 30;
      }
      const response = await fetch(`http://${server.value}/gen?unknowns=${generateUnknowns.value}`, { method: 'GET' });
      const data = await response.json();
      grid.value = data.Grid.split("").map(cell => cell === "0" ? "" : cell);
    };

    const solve = async () => {
      const gridString = grid.value.map(cell => cell === "" ? "0" : cell).join("");
      const response = await fetch(`http://${server.value}/solve?input=${gridString}`, {
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

    const onInput = (index, event) => {
      let value = grid.value[index];
      // Only keep the last entered digit if multiple are input
      if (typeof value === "string" && value.length > 1) {
        value = value.slice(-1);
        if (/^[1-9]$/.test(value)) {
          grid.value[index] = value;
        } else {
          grid.value[index] = grid.value[index].replace(/[^1-9]/g, "");
        }
      }
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
    }


    return {
      grid,
      status,
      generateUnknowns,
      server,

      clear,
      generate,
      solve,
      onInput,
    };

  }
}
