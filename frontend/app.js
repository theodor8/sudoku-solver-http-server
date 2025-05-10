const { ref } = Vue;


export default {
  setup() {
    const message = ref("Hello Vue 3!");
    const count = ref(0);

    const clickMe = () => {
      count.value++;
      if (count.value === 1) {
        message.value = "Button clicked";
      }
    };

    return {
      message,
      count,
      clickMe,
    };
  }
}
