<script>
  import logo from "./assets/images/logo-universal.png";
  import "svelte-material-ui/themes/svelte-dark.css";
  import {
    Search,
    GetPartitions,
    ListDir,
    OpenFile,
  } from "../wailsjs/go/main/App.js";
  import DataTable, { Head, Body, Row, Cell } from "@smui/data-table";
  import { onMount } from "svelte";
  import Tootltip from "./Components/Tootltip.svelte";
  import { Jumper } from "svelte-loading-spinners";
  let resultText = "Please enter your name below 👇";
  let name;
  let testResults;
  let drives = [];
  let files = [];
  let dirs = [];
  let showDrives = true;
  let searchTerm = "";
  let curDir = "";
  let loading = false;
  onMount(async () => {
    console.log("mounted");
    drives = await GetPartitions();
  });

  async function handleSearch() {
    if (searchTerm !== "") {
      loading = true;
      let result = await Search(curDir, searchTerm);
      loading = false;
      console.log(result);
      dirs = [];
      files = result;
      curDir = "";
      showDrives = false;
    }
  }
  async function handleFileClick(filePath) {
    console.log(filePath);
    await OpenFile(filePath);
  }
  async function handleBack() {
    if (curDir !== "" || !curDir.endsWith(":\\\\")) {
      //need to get from current dir string the prev string
      const previousDirectory = curDir.split("\\").slice(0, -1).join("\\");

      let result = await ListDir(previousDirectory);
      console.log(result);
      files = result.Files;
      dirs = result.Directories;
      curDir = previousDirectory;
      showDrives = false;
    }
  }
  async function listDir(path) {
    let result = await ListDir(path);
    console.log(result);
    files = result.Files;
    dirs = result.Directories;
    curDir = path;
    showDrives = false;
  }
</script>

<main style="overflow-y: hidden;">
  <div class="result" id="result">File Explorer</div>
  <div class="input-box" id="input">
    <div style="margin-bottom: 10px;">
      <input
        autocomplete="off"
        bind:value={searchTerm}
        class="input"
        id="name"
        type="text"
      />
      <button class="btn" on:click={handleSearch}>Search</button>
    </div>
    <div class="toolscontainer">
      <button
        class="btn"
        disabled={curDir === "" ? true : false}
        on:click={() => handleBack()}>Back</button
      >
      {#if drives.length > 0}
        {#each drives as drive}
          <button class="btn" on:click={() => listDir(drive)}>{drive}</button>
        {/each}
      {/if}
    </div>
    <br /> <br />
    {#if loading}
      <div style="display: flex; justify-content: center; width: 100%; ">
        <Jumper size="60" color="#FF3E00" unit="px" duration="1s" />
      </div>
    {/if}
    {#if !showDrives}
      <DataTable
        table$aria-label="People list"
        style="width: 90%; max-height: 500px"
      >
        <Head>
          <Row>
            <Cell>Name</Cell>
            <Cell>Path</Cell>
          </Row>
        </Head>
        <Body>
          {#each dirs as dir}
            <Row
              style="cursor: pointer"
              on:click={() => {
                listDir(dir.Filepath);
              }}
            >
              <Cell>{dir.Filename}</Cell>
              <Cell>{dir.Filepath}</Cell>
            </Row>
          {/each}
          {#each files as file}
            <Row
              on:click={() => handleFileClick(file.Filepath)}
              style="cursor: pointer"
            >
              <Cell>{file.Filename}</Cell>
              <Cell>{file.Filepath}</Cell>
            </Row>
          {/each}
        </Body>
      </DataTable>

      <!-- {#each dirs as dir}
      <div style="margin-bottom: 10px; display: inline-block"> 
        <Tootltip text={dir.Filename} >
                <button class="btn-res" on:click={() => listDir(dir.Filepath)}>{dir.Filename.length > 10 ? dir.Filename.slice(0,10) + '...' : dir.Filename}</button>
        </Tootltip>       
      </div>
      {/each}
      {#each files as file}
      <div style="margin-bottom: 10px; display: inline-block"> 
        <Tootltip text={file.Filename}>
          <button class="btn-res" on:click={() => handleFileClick(file.Filepath)}>{file.Filename.length > 10 ? file.Filename.slice(0,10) + '...' : file.Filename}</button>

        </Tootltip>       
      </div>
      {/each} -->
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
  .btn-res {
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
