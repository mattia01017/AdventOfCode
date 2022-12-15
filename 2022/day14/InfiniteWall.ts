import {Wall} from "./Wall"

export class InfiniteWall extends Wall {

    public constructor(rockPaths: String[]) {
        super(rockPaths);
    }

    public startFlow(): number {
        let sandUnits = 0;
        while (this.dropSand()) {
            sandUnits++;
        }
        return sandUnits
    }

    private dropSand(): boolean {
        let xs = 500
        let ys = 0
        while (ys < this.lowery) {
            ys++;
            if (this.grid.has(`${xs},${ys}`)) {
                if (!this.grid.has(`${xs-1},${ys}`)) {
                    xs--;
                } else if (!this.grid.has(`${xs+1},${ys}`)) {
                    xs++;
                } else {
                    this.grid.add((`${xs},${ys-1}`))
                    return true;
                }
            }
        }
        return false;
    }
}