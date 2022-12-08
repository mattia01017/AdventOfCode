const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin
});

function isVisible(i, j) {
    let currTree = trees[i][j]

    let rightVisibile = Math.max(...trees[i].slice(j + 1)) < currTree;
    let leftVisible = Math.max(...trees[i].slice(0, j)) < currTree;
    let upVisible = Math.max(...trees.map(x => x[j]).slice(0, i)) < currTree;
    let downVisible = Math.max(...trees.map(x => x[j]).slice(i + 1)) < currTree;

    return rightVisibile || leftVisible || upVisible || downVisible
}

function countVisibles() {
    counter = 0;
    for (let i = 0; i < trees.length; i++) {
        for (let j = 0; j < trees[i].length; j++) {
            if (isVisible(i, j))
                counter++;
        }
    }
    return counter
}

function calcBestScenicScore() {
    let bestScore = 0
    for (let i = 0; i < trees.length; i++) {
        for (let j = 0; j < trees[i].length; j++) {
            let m = trees[i][j]
            let tmpScore = [1, 1, 1, 1]
            let k;
            for (k = 1; i + k < trees.length && trees[i + k][j] < m; k++)
                tmpScore[0]++;
            if (i + k == trees.length) tmpScore[0]--;
            for (k = 1; i - k >= 0 && trees[i - k][j] < m; k++)
                tmpScore[1]++;
            if (i - k < 0) tmpScore[1]--;
            for (k = 1; j + k < trees[i].length && trees[i][j + k] < m; k++)
                tmpScore[2]++;
            if (j + k == trees.length) tmpScore[2]--;
            for (k = 1; j - k >= 0 && trees[i][j - k] < m; k++)
                tmpScore[3]++;
            if (j - k < 0) tmpScore[3]--;
            let totTmpScore = 1
            for (s of tmpScore) totTmpScore *= s;
            if (totTmpScore > bestScore) bestScore = totTmpScore;
        }
    }
    return bestScore;
}

var trees = Array()

rl.addListener('line', input => {
    trees.push(Array.from(input).map(x => Number(x)))
})

rl.addListener('close', () => {
    let res1 = countVisibles();
    console.log("Number of visible trees: " + res1);
    let res2 = calcBestScenicScore();
    console.log("Best scenic score: " + res2)
})
