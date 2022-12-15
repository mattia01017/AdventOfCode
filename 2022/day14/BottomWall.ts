import { Wall } from "./Wall"

export class BottomWall extends Wall {

    bottom: number;

    public constructor(rockpaths: string[]) {
        super(rockpaths);
        this.bottom = this.lowery + 2;
    }

    public startFlow(): number {
        let sandUnits = 0;
        while (this.dropSand()) {
            sandUnits++;
        }
        return sandUnits;
    }

    private dropSand(): boolean {
        let xs = 500
        let ys = 0

        if (this.grid.has("500,0"))
            return false;

        while (ys <= this.lowery) {
            if (this.grid.has(`${xs},${ys+1}`)) {
                if (!this.grid.has(`${xs-1},${ys+1}`)) {
                    xs--;
                } else if (!this.grid.has(`${xs+1},${ys+1}`)) {
                    xs++;
                } else {
                    break;
                }
            }
            ys++;
        }
        this.grid.add(`${xs},${ys}`)
        return true;
    }
}