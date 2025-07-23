<script lang="ts">
  import QuizCard from "./QuizCard.svelte";
  import Button from "./Button.svelte";
  import type {Quiz} from "../models/models";

  let quizzes: Quiz[] = [];

  async function getQuizzes() {
    const response = await fetch("http://localhost:3000/api/getquizzes");
    if (!response.ok) {
      alert("error: " + await response.text());
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

  function hostQuiz(quiz: Quiz) {
    console.log("hosting: ", quiz)
  }
</script>

<div class="p-4">
  <Button onclick={getQuizzes}>Get Quizzes</Button>
<!--  <Button on:click={webSocketConnect}>Connect</Button>-->
</div>

<div class="flex p-4">
  <input class="border p-1" type="text" placeholder="Game Code">
  <Button class="ml-4" >Join</Button>
</div>

{#each quizzes as quiz}
  <QuizCard host={() => hostQuiz(quiz)} {quiz}/>
{/each}
