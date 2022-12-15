type Coord = { x: number, y: number }

export abstract class Wall {

    protected grid: Set<string>;
    protected lowery: number;

    public constructor(rockPaths: String[]) {
        this.lowery = 0;
        this.grid = new Set()
        rockPaths.forEach(path => {
            let corners = path.split(" -> ");
            for (let i = 1; i < corners.length; i++) {
                let first = this.strToCoord(corners[i]);
                let second = this.strToCoord(corners[i - 1]);
                this.drawLine(first, second);
            }
        });
    }

    private strToCoord(s: string): Coord {
        let xy = s.split(",")
        let coord: Coord = { x: Number(xy[0]), y: Number(xy[1]) }
        return coord
    }

    private drawLine(coord1: Coord, coord2: Coord) {
        if (coord1.x == coord2.x) {
            if (coord2.y < coord1.y) {
                let t = coord1;
                coord1 = coord2;
                coord2 = t;
            }
            for (let i = coord1.y; i <= coord2.y; i++) {
                this.grid.add(`${coord1.x},${i}`);
            }
        } else {
            if (coord2.x < coord1.x) {
                let t = coord1;
                coord1 = coord2;
                coord2 = t;
            }
            for (let i = coord1.x; i <= coord2.x; i++) {
                this.grid.add(`${i},${coord1.y}`);
            }
        }
        if (this.lowery < coord2.y) this.lowery = coord2.y
    }

    abstract startFlow(): number;
}