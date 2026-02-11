import fs from 'node:fs/promises';
// const fs = require('node:fs/promises');

async function read(path: string) {
  const data = await fs.readFile(path, { encoding: 'utf8' });
  return data
}

class Circuit {
  boxes: Set<Box>

  constructor(boxes: Box[]) {
    this.boxes = new Set(boxes)
  }

  private addBoxes(newBoxes: Set<Box>) {
    for (let nb of newBoxes) {
      this.boxes.add(nb)
    }
  }

  public combineCircuits(c: Circuit) {
    this.addBoxes(c.boxes)
    c = this
  }
}

class Box {
  x: number
  y: number
  z: number
  circuit: Circuit

  constructor(x: number, y: number, z: number) {
    this.x = x
    this.y = y
    this.z = z
    this.circuit = new Circuit([this])
  }
}

function getDistance(a: Box, b: Box): number {
	let dx =  (a.x - b.x) * (a.x - b.x)
	let dy = (a.y - b.y) * (a.y - b.y)
	let dz = (a.z - b.z) * (a.z - b.z)
	return Math.sqrt(dx + dy + dz)
}

async function main() {
  let amount = 10
  let text: string = await read('example.txt')

  // let amount = 1000
  // let text: string = await read('input.txt')


  let rows = text.split('\n')
    .filter(row => row.length > 0)
  
  let boxes = rows.map(row => row.split(',')
  .map(str => Number(str)))
  .map(v => new Box(v[0], v[1], v[2]))

  let distances: {dist: number, a: Box, b: Box}[] = []

  for (let i = 0; i < boxes.length-1; i++) {
    for (let j = i + 1; j < boxes.length; j++) {
      let a = boxes[i]
      let b = boxes[j]
      distances.push({
        dist: getDistance(a, b),
        a,
        b,
      })
    }
  }

  distances.sort((v1, v2) => -(v1.dist - v2.dist)) // reversed

  // for (let i =0; i < 10; i++) {
  //   console.log(distances[distances.length - 1 - i])
  // }



  for (let i = 0; i < amount; i++) {
    let d = distances.pop()
    console.log(d?.dist)
    let ac = d?.a.circuit
    let bc = d?.b.circuit
    if (ac === undefined || bc === undefined) {
      throw new Error("undefined circuit");
    }
    ac.combineCircuits(bc)
  }

  let s = new Set(
    boxes.map(b => b.circuit)
  )
  console.log(s)



  // // console.log(boxes[0].circuit.boxes.size)
  // let xx = boxes.sort((a, b) => -(a.circuit.boxes.size-b.circuit.boxes.size))
  //   .map(v => v.circuit.boxes)
  // console.log(xx)

  // let aa = ['1', '2', '3']
  // let bb = ['2', '1', '3']

  // let setAa = new Set(aa)
  // console.log(setAa)
  // let setBb = new Set(bb)
  // console.log(setBb)

}

main()