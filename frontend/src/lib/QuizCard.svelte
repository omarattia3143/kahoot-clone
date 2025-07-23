<script lang="ts">
    import type {Quiz} from "../models/models";
    import Button from "./Button.svelte";

    export let quiz: Quiz;
    export let host: () => void;
</script>

<div class="flex justify-between bg-white rounded p-4 items-center border-2 m-2 border-green-500">
    <p>{quiz.name}</p>
    <Button onclick={host}>Host</Button>
</div>