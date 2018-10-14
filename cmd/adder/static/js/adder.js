document.addEventListener("DOMContentLoaded", function () {
  var btn = document.getElementById('adder-btn');
  var r = document.getElementById('adder-result');
  btn.onclick = function(e) {
    var a = document.getElementById('adder-a');
    var b = document.getElementById('adder-b');
    btn.classList.add("pure-button-disabled");
    var req = new XMLHttpRequest();
    req.onload = function() {
      if (req.readyState == req.DONE && req.status === 200 ) {
        r.textContent = req.responseText;
        btn.classList.remove("pure-button-disabled");
      }
    };
    req.open("GET","/adder?a="+a.value+"&b="+b.value);
    req.send();

  };
});
