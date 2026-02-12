import fs from 'node:fs/promises';
// const fs = require('node:fs/promises');

async function read(path: string) {
  const data = await fs.readFile(path, { encoding: 'utf8' });
  return data
}

type Distance = { dist: number, a: Box, b: Box }

class Box {
  x: number
  y: number
  z: number
  circuit: number

  constructor(x: number, y: number, z: number) {
    this.x = x
    this.y = y
    this.z = z
    this.circuit = -1
  }
}

function getDistance(a: Box, b: Box): number {
  let dx = Math.pow(a.x - b.x, 2)
  let dy = Math.pow(a.y - b.y, 2)
  let dz = Math.pow(a.z - b.z, 2)
  return Math.sqrt(dx + dy + dz)
}

function changeCircuitId(boxes: Box[], oldId: number, newId: number) {
  boxes.forEach(box => {
    if (box.circuit === oldId) {
      box.circuit = newId
    }
  })
}

function part1(boxes: Box[]) {
  let cc: { [k: string]: Box[] } = {}
  for (let box of boxes) {
    if (box.circuit !== -1) {
      if (!cc[box.circuit]) {
        cc[box.circuit] = [box]
      } else {
        cc[box.circuit].push(box)
      }
    }
  }

  let arr = Object.values(cc).sort((a: Box[], b: Box[]) => -(a.length - b.length))
  let result = 1
  for (let i = 0; i < 3; i++) {
    result *= arr[i].length
  }
  console.log('result:',result)
}

async function main() {
  // let amount = 10
  // let text: string = await read('example.txt')

  let amount = 1000
  let text: string = await read('input.txt')

  let rows = text.split('\n')
    .filter(row => row.length > 0)

  let boxes = rows.map(row => row.split(',')
    .map(str => Number(str)))
    .map(v => new Box(v[0], v[1], v[2]))

  let distances: Distance[] = []

  for (let i = 0; i < boxes.length - 1; i++) {
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

  distances.sort((v1, v2) => (v1.dist - v2.dist))

  let circuitId = 0
  for (let i = 0; i < distances.length; i++) {
    let d = distances[i]

    if (d.a.circuit === -1 && d.b.circuit === -1) { // molemmat eivat ole piirissa
      circuitId++
      d.a.circuit = circuitId
      d.b.circuit = circuitId
    } else if (d.a.circuit == -1 && d.b.circuit != -1) { // b ei ole piirissa
      d.a.circuit = d.b.circuit
    } else if (d.a.circuit != -1 && d.b.circuit == -1) { // a ei ole  piirissa
      d.b.circuit = d.a.circuit
    } else if (d.a.circuit != -1 && d.b.circuit != -1) { // molemmat ovat piirissa
      let oldId = d.b.circuit
      d.b.circuit = d.a.circuit
      changeCircuitId(boxes, oldId, d.a.circuit)
    }

    if (i + 1 === amount) {
      console.log('part1 ====================')
      part1(boxes)
    }
    if (allInSameCircuit(boxes)) {
      console.log('part2 ====================')
      console.log('result: ', d.a.x * d.b.x)
      break
    }
  }

}

function allInSameCircuit(boxes: Box[]) {
  let pre = null
  for (let b of boxes) {
    if (pre === null) {
      pre = b.circuit
    } else if (b.circuit !== pre) {
      return false
    }
  }
  return true
}

main()