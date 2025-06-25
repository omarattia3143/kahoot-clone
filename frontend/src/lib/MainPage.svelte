<script lang="ts">
  import QuizCard from "./QuizCard.svelte";
  import Button from "./Button.svelte";

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
    const ws = new WebSocket("ws://localhost:3000/ws");
    ws.onopen = () => {
      ws.send("Hello");
    };
    ws.onmessage = (event) => {
      console.log(event.data);
    };
  }
</script>

<div>
  <Button on:click={getQuizzes}>Get Quizzes</Button>
  <Button on:click={webSocketConnect}>Connect</Button>
</div>

{#each quizzes as quiz}
  <QuizCard {quiz}/>
{/each}
