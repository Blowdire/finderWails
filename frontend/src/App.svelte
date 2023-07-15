<script>
  import logo from './assets/images/logo-universal.png'
  import {Greet, Search, GetPartitions, ListDir} from '../wailsjs/go/main/App.js'
  import {onMount} from 'svelte'
  import Tootltip from './Components/Tootltip.svelte';
  let resultText = "Please enter your name below 👇"
  let name
  let testResults
  let drives = []
  let files = []
  let dirs = []
  let showDrives = true
  onMount(async () => {
    console.log('mounted')
    drives = await GetPartitions();
  })

  function greet() {
    Greet(name).then(result => resultText = result)
  }
 

   async function listDir(path){
    
    let result = await ListDir(path)
    console.log(result)
    files = result.Files
    dirs = result.Directories
    showDrives = false
  }
</script>

<main>

  <div class="result" id="result">File Explorer</div>
  <div class="input-box" id="input">
    <div style="margin-bottom: 10px;">
      <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
      <button class="btn" on:click={greet}>Search</button>
    </div>
    
    {#if showDrives && drives.length > 0}
      {#each drives as drive}
        <button class="btn" on:click={() => listDir(drive)}>{drive}</button>
      {/each}
    {/if} <br>
    {#if !showDrives}
      {#each dirs as dir}
      <div style="margin-bottom: 10px; display: inline-block">        
        <button class="btn-res">{dir.Filename}</button>
      </div>
      {/each}
      {#each files as file}
      <div style="margin-bottom: 10px; display: inline-block"> 
        <Tootltip text={file.Filename}>
          <button class="btn-res" >{file.Filename.length > 10 ? file.Filename.slice(0,10) + '...' : file.Filename}</button>

        </Tootltip>       
      </div>
      {/each}
    {/if}
  </div>
</main>

<style>

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }
  .btn-res{
    width: 150px;
    height: 30px;
    text-overflow: ellipsis;
    line-height: 30px;
    margin-top: 10px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }
  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }
  .btn-res:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    width: 80%;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
