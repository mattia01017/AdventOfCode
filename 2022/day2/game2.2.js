class ChooseSignRPS {
    #signs = {
        ROCK: {value: 1},
        PAPER: {value: 2},
        SCISSORS: {value: 3}
    };

    #opponentSigns = {
        'A': this.#signs.ROCK,
        'B': this.#signs.PAPER,
        'C': this.#signs.SCISSORS
    };

    hierarchy = new Map();

    constructor() {
        this.hierarchy.set(
            this.#signs.PAPER,
            {
                winsAgainst: this.#signs.ROCK,
                losesAgainst: this.#signs.SCISSORS
            }
        );
        this.hierarchy.set(
            this.#signs.ROCK,
            {
                winsAgainst: this.#signs.SCISSORS,
                losesAgainst: this.#signs.PAPER
            }
        );
        this.hierarchy.set(
            this.#signs.SCISSORS,
            {
                winsAgainst: this.#signs.PAPER,
                losesAgainst: this.#signs.ROCK
            }
        );
    }

    #parseOpponent(letter) {
        return this.#opponentSigns[letter];
    }

    #parseMe(letter, opponentSign) {
        switch (letter) {
            case 'X':
                return this.hierarchy.get(opponentSign).winsAgainst
            case 'Y':
                return opponentSign
            case 'Z':
                return this.hierarchy.get(opponentSign).losesAgainst
        }
    }

    points(match) {
        let opponentSign = this.#parseOpponent(match.charAt(0));
        let mySign = this.#parseMe(match.charAt(2), opponentSign);

        let points = mySign.value;
        if (this.hierarchy.get(mySign).winsAgainst == opponentSign) {
            points += 6;
        } else if (this.hierarchy.get(mySign).losesAgainst != opponentSign) {
            points += 3;
        }
        return points;
    }
}

module.exports = ChooseSignRPS;