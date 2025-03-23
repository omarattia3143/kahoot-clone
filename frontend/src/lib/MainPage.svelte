<script lang="ts">
  import Button from "./Button.svelte";
  import QuizCard from "./QuizCard.svelte";

  let quizzes: { _id: string; name: string }[] = [];

  async function getQuizzes() {
    const response = await fetch("http://localhost:3000/api/getquizzes");
    if (!response.ok) {
      alert("error: " + response.text);
      return;
    }
    quizzes = await response.json();
    console.log(quizzes);
  }

  function webSocketConnect() {
    const ws = new WebSocket("http://localhost:3000/ws");
    ws.onopen = () => {
      console.log("Connected to WebSocket server");
      ws.send("Hello from the client!");
    };
    ws.onmessage = (event) => {
      console.log("Message from server:", event.data);
    };
  }

  webSocketConnect();
  getQuizzes();
</script>

{#each quizzes as quiz}
  <QuizCard {quiz}/>
{/each}
