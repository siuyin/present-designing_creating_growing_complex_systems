package main

import "html/template"

var adderForm = template.HTML(`
<form class="pure-form">
    <fieldset>
        <legend>Adding two whole numbers</legend>

        <input id="adder-a" type="number" step="1" placeholder="0">
        <input id="adder-b" type="number" step="1" placeholder="0">

        <button type="button" class="pure-button pure-button-primary"
          id="adder-btn"
        >Go</button>
    </fieldset>
</form>
`)
