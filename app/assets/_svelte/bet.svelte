<script>
  import { create } from 'svelte';
  
  export let horse;
  export let amount;
  
  const handleSubmit = () => {
    // Make a bet with the horse and amount.
    console.log('Making a bet with horse:', horse, 'and amount:', amount);
  };
  
  const App = create({
    name: 'App',
    components: {},
    props: {},
    data: {
      horse: '',
      amount: '',
    },
    methods: {
      handleSubmit,
    },
  });
  
  export default App;
</script>
  
<h1>Make a Bet</h1>
  
<input
  type="text"
  placeholder="Horse"
  bind:value={horse}
/>
  
<input
  type="number"
  placeholder="Amount"
  bind:value={amount}
/>
  
<button on:click={handleSubmit}>Make Bet</button>