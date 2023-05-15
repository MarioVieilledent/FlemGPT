<script lang="ts">
  import { onMount } from "svelte";

  const url = "http://localhost:8080";

  let prompt: string = "";

  let results: string[] = [];

  let depth: number;
  let nbWords: number;
  let randPickOverBest: number;
  let coefOccur: number;
  let coefDistance: number;
  let coefWordLength: number;
  let minOccurToAccept: number;

  const keyUpPrompt = (event: any) => {
    if (event.key === "Enter") {
      api(`/prompt/${event.target.value}`, (res: any) => {
        results = [...results, res.result.join(" ")];
      });
    }
  };

  const modifyValue = (varName: string, event: any) => {
    api(`/set/${varName}/${event.target.value}`, (res: any) =>
      console.log(res)
    );
  };

  const api = (query: string, f: (res: any) => void): void => {
    fetch(url + query)
      .then((d) => d.json())
      .then(f);
  };

  onMount(() => {
    api("/get/all", (res: any) => {
      console.log(res);
      depth = res.depth;
      nbWords = res.nbWords;
      randPickOverBest = res.randPickOverBest;
      coefOccur = res.coefOccur;
      coefDistance = res.coefDistance;
      coefWordLength = res.coefWordLength;
      minOccurToAccept = res.minOccurToAccept;
    });
  });
</script>

<div class="f main">
  <div class="fc left">
    <span>depth</span>
    <input
      type="number"
      bind:value={depth}
      on:change={(event) => modifyValue("depth", event)}
    />
    <span>nbWords</span>
    <input
      type="number"
      bind:value={nbWords}
      on:change={(event) => modifyValue("nbWords", event)}
    />
    <span>randPickOverBest</span>
    <input
      type="number"
      bind:value={randPickOverBest}
      on:change={(event) => modifyValue("randPickOverBest", event)}
    />
    <span>coefOccur</span>
    <input
      type="number"
      bind:value={coefOccur}
      step="0.001"
      on:change={(event) => modifyValue("coefOccur", event)}
    />
    <span>coefDistance</span>
    <input
      type="number"
      bind:value={coefDistance}
      step="0.001"
      on:change={(event) => modifyValue("coefDistance", event)}
    />
    <span>coefWordLength</span>
    <input
      type="number"
      bind:value={coefWordLength}
      step="0.001"
      on:change={(event) => modifyValue("coefWordLength", event)}
    />
    <span>minOccurToAccept</span>
    <input
      type="number"
      bind:value={minOccurToAccept}
      on:change={(event) => modifyValue("minOccurToAccept", event)}
    />
  </div>

  <div class="fc right">
    <h1>FlemGPT</h1>
    <div class="fc conversations">
      {#each results as res}
        <div class="result">{res}</div>
      {/each}
    </div>
    <input
      type="text"
      placeholder="Prompt..."
      bind:value={prompt}
      on:keyup={(event) => keyUpPrompt(event)}
    />
  </div>
</div>

<style lang="scss">
  $left-width: 260px;

  .main {
    width: 100%;
    height: 100%;

    .left {
      width: $left-width;
      height: 100%;
      padding: 12px;
      box-sizing: border-box;
      background-color: #202123;
    }

    .right {
      width: calc(100% - $left-width);
      height: 100%;
      padding: 12px;
      box-sizing: border-box;
      background-color: #343541;
      justify-content: space-between;

      .conversations {
        overflow: auto;
      }
    }
  }
</style>
