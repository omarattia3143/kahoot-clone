<script lang="ts">
  import QuizCard from "./QuizCard.svelte";
  import Button from "./Button.svelte";
  import type {Quiz} from "../models/models";

  let quizzes: Quiz[] = $state([]);
  let code: string = $state("");
  let msg: string = $state("");
  let ws: WebSocket | null = null;

  function getWebSocket() {
    if (!ws || ws.readyState === WebSocket.CLOSED) {
      ws = new WebSocket("ws://localhost:3000/ws");
      ws.onmessage = (event) => {
        console.log(event.data);
      };
    }
    return ws;
  }

  async function getQuizzes() {
    const response = await fetch("http://localhost:3000/api/getquizzes");
    if (!response.ok) {
      alert("error: " + await response.text());
      return;
    }
    quizzes = await response.json();
    console.log($state.snapshot(quizzes))
  }

  function hostQuiz(quiz: Quiz) {
    const socket = getWebSocket();
    socket.onopen = () => {
      socket.send(`host:${quiz.id}`);
    };
    if (socket.readyState === WebSocket.OPEN) {
      socket.send(`host:${quiz.id}`);
    }

    socket.onmessage = (event) => {
      msg = event.data;
    }
  }

  function joinQuiz(code: string) {
    const socket = getWebSocket();
    socket.onopen = () => {
      socket.send(`join:${code}`);
    };
    if (socket.readyState === WebSocket.OPEN) {
      socket.send(`join:${code}`);
    }
  }
</script>

<div class="p-4">
  <Button onclick={getQuizzes}>Get Quizzes</Button>
  <span> Message: {msg}</span>
</div>

<div class="flex p-4">
  <input class="border p-1" type="text" bind:value={code} placeholder="Game Code">
  <Button onclick={() => joinQuiz(code)} class="ml-4">Join</Button>
</div>

{#each quizzes as quiz}
  <QuizCard host={() => hostQuiz(quiz)} {quiz}/>
{/each}