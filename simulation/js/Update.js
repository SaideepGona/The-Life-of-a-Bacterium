//Saideep Gona
const fs = require("fs")
const process = require("process")
var os = require('os')
var child = require('child_process').execFile
type = os.type()
console.log(type)
if (type == "Darwin") {
  var executableName = "/simulation";
} else {
  var executableName = "/simulation.exe";
}

var pwd = process.cwd()
//var executableName = "/simulation"
var loaded = "no"
var newChecker
var processing = "no"


 $("#RunButton").hover(function(){                                         // Lights up the run button on hover
    $("#RunButton").css("background-color", "rgb(73, 98, 187)");
    $("#run_hover_message").fadeIn(100);                                   // Also fades in some text
  },
  function(){
    $("#RunButton").css("background-color", "");
    $("#run_hover_message").fadeOut(200);
    console.log("unhover")
  });

$("#RefreshButton").click(function(){
  location.reload()
  /*
 	 $("#two").load("two.html");
   $("#three").load("three.html");
   $("#four").load("four.html");
   $.ajaxSetup({ cache: false });
   */
});

help = 1

 $("#help").click(function(){
  if (help == 1) {
    $("#instructionbox").css("color", "white");
    help = 0
  } else if (help == 0) {
    $("#instructionbox").css("color", "#666");
    help = 1
  }
 });


$("#LoadButton").click(function(){
  $(".content").fadeIn(1000);
});

var runArgs = ['200.0', '3', '200', '30', '50', '2', '13', '0.5']

$("#UpdateButton").click(function(){

  var newString = document.getElementById("initialization").value
  var newArgs = newString.split(" | ")
  if (newArgs.length==runArgs.length) {
    runArgs=newArgs
  }
  console.log(runArgs)

});

$("#RunButton").click(function(){                                          // On click of "Run" fades in the loading button and runs go script
    if (processing == "yes") {
      return
    }

    processing = "yes"
    loaded = "no"
    $("#load_box").fadeIn(1000);
    $('.circle-loader').toggleClass('load-complete');
    $('.checkmark').toggle();
    $(".content").fadeOut(1000);
    console.log("should empty now")
    empty()
    //setTimeout(empty,1000);

    newChecker = setInterval(checkLoad, 5000, loaded);
    console.log(pwd)
    child(pwd+executableName, runArgs, function(err, data) {                            // Runs go script
    if(err){
       console.error(err);
       return;
    }
    processing = "no"
    console.log(data.toString());
  });
});                                                           // Changes when loading is successful

function empty() {
  $(".content").empty()
}

function checkLoad(loaded) {
    console.log("Checking if executable is done", processing)
    if (processing=="no") {
        clearInterval(newChecker)
        console.log("Something")
        $('.circle-loader').toggleClass('load-complete');
        $('.checkmark').toggle();
        $('.circle-loader').fadeOut(1000);
        var loaded = "yes"  //Prevents retoggling
    }
};

function topFunction() {        // Click to return to top
  document.body.scrollTop = 0;
  document.documentElement.scrollTop = 0;
}

$(document).ready(function(){                                               // On ready, makes the check_box hidden
    //loadCharts()
    $("#load_box").hide();
    $(".content").hide();
    //$("#instructionbox").hide();
    //$("#init").hide();
    $('.circle-loader').toggleClass('load-complete');
    $('.checkmark').toggle();

});
