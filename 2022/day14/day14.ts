import readline from "readline";
import { BottomWall } from "./BottomWall"
import { InfiniteWall } from "./InfiniteWall"

var rl = readline.createInterface({
    input: process.stdin
});

let paths: string[] = []
rl.addListener("line", (input: string) => {
    paths.push(input);
})

rl.addListener("close", () => {
    let infWall = new InfiniteWall(paths);
    let botWall = new BottomWall(paths);
    console.log(infWall.startFlow())
    console.log(botWall.startFlow())
})