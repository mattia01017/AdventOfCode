const readline = require("readline");
const ChooseSignRPS = require("./game2.1");
const ChooseResultRPS = require("./game2.2");

var rl = readline.createInterface({
    input: process.stdin
});

var chooseSignRPS = new ChooseSignRPS();
var chooseResultRPS = new ChooseResultRPS();
var points1 = 0;
var points2 = 0;

rl.addListener("line", input => {
    points1 += chooseSignRPS.points(input);
    points2 += chooseResultRPS.points(input);
});

rl.addListener("close", () => {
    console.log("Punteggi prima parte: %d", points1);
    console.log("Punteggi seconda parte: %d", points2);
});
