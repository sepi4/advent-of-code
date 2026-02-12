use std::fs;

fn read_input(file_path: &str) -> String {
    let text = fs::read_to_string(file_path).expect("virhe tiedoston lukussa");
    text
}

#[derive(Debug)]
struct Box {
    x: usize,
    y: usize,
    z: usize,
    circuit: i32,
}

#[derive(Debug)]
struct Distance {
    dist: f64,
    a_id: usize,
    b_id: usize,
}

fn get_distance(a: &Box, b: &Box) -> f64 {
    fn diff(v1: usize, v2: usize) -> f64 {
        v1 as f64 - v2 as f64
    }
    let dx = diff(a.x, b.x).powi(2);
    let dy = diff(a.y, b.y).powi(2);
    let dz = diff(a.z, b.z).powi(2);
    return (dx + dy + dz).sqrt();
}

fn main() {
    let amount = 10;
    let text = read_input("../example.txt");

    // let amount = 1000;
    // let text = read_input("../input.txt");

    let mut boxes = text
        .split('\n')
        .map(|row| {
            row.split(',')
                .map(|s| s.parse().expect("error in parse num"))
                .collect::<Vec<usize>>()
        })
        .map(|b| Box {
            x: b[0],
            y: b[1],
            z: b[2],
            circuit: -1,
        })
        .collect::<Vec<Box>>();

    let mut dd: Vec<Distance> = vec![];
    for i in 0..boxes.len() - 1 {
        for j in (i + 1)..boxes.len() {
            dd.push(Distance {
                dist: get_distance(&boxes[i], &boxes[j]),
                a_id: i,
                b_id: j,
            })
        }
    }

    dd.sort_by(|a, b| a.dist.partial_cmp(&b.dist).unwrap());

    let mut circuit_id = 0;
    for i in 0..dd.len() {
        let ai = dd[i].a_id;
        let bi = dd[i].b_id;

        if boxes[ai].circuit == -1 && boxes[bi].circuit == -1 {
            // molemmat eivat ole piirissa
            circuit_id += 1;
            boxes[ai].circuit = circuit_id;
            boxes[bi].circuit = circuit_id;
        } else if boxes[ai].circuit == -1 && boxes[bi].circuit != -1 {
            // b ei ole piirissa
            boxes[ai].circuit = boxes[bi].circuit;
        } else if boxes[ai].circuit != -1 && boxes[bi].circuit == -1 {
            // a ei ole  piirissa
            boxes[bi].circuit = boxes[ai].circuit;
        } else if boxes[ai].circuit != -1 && boxes[bi].circuit != -1 {
            // molemmat ovat piirissa
            let old_id = boxes[bi].circuit;
            let new_id = boxes[ai].circuit;
            boxes[bi].circuit = new_id;
            change_circuit_ids(&mut boxes, old_id, new_id);
        }

        if all_in_same_circuit(&boxes) {
            println!("part2 ================");
            let result = boxes[ai].x * boxes[bi].x;
            println!("result: {:?}", result);
            break;
        }
    }
}

fn all_in_same_circuit(boxes: &Vec<Box>) -> bool {
    let mut pre: i32 = -100;
    for i in 0..boxes.len() {
        if pre == -100 {
            pre = boxes[i].circuit;
        } else if boxes[i].circuit != pre {
            return false;
        }
    }
    return true;
}

fn change_circuit_ids(boxes: &mut Vec<Box>, old_id: i32, new_id: i32) {
    // change all old_id to new_id
    for i in 0..boxes.len() {
        if boxes[i].circuit == old_id {
            boxes[i].circuit = new_id;
        }
    }
}
