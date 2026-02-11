use std::{fs};

fn read_input(file_path: &str) -> String {
    let text = fs::read_to_string(file_path).expect("virhe tiedoston lukussa");
    text
}

fn part1() {}
fn part2() {}

#[derive(Debug)]
struct Box {
    x: usize,
    y: usize,
    z: usize,
    circuit: Option<Circuit>,
}

#[derive(Debug)]
struct Distance<'a> {
    dist: f64,
    a: &'a Box,
    b: &'a Box,
}

#[derive(Debug)]
struct Circuit {
    boxes: Vec<Box>
}

fn get_distance(a: &Box, b: &Box) -> f64 {
    fn erotus(v1: usize, v2: usize) -> f64 {
        return v1 as f64 - v2 as f64;
    }
    let inside = erotus(a.x, b.x).powi(2)
        + erotus(a.y, b.y).powi(2)
        + erotus(a.z, b.z).powi(2);
    
    return inside.sqrt();
}
    
fn main() {
    let text = read_input("../example.txt");
    // let text = read_input("../input.txt");

    let boxes: Vec<Box> = text.split('\n')
        .filter(|row| row.len() > 0)
        .map(|row| row
            .trim()
            .split(',')
            .map(|x| x.parse().expect("xxx!!!!"))
            .collect::<Vec<usize>>()
        )
        .map(|ss| Box{ x: ss[0], y: ss[1], z: ss[2], circuit: None })
        .collect();
    
    let mut distances: Vec<Distance> = vec![];
    for i in 0..boxes.len()-1 {
        for j in i+1..boxes.len() {
            distances.push(Distance{
                dist: get_distance(&boxes[i], &boxes[j]),
                a: &boxes[i],
                b: &boxes[j],
            });
        }
    }

    distances.sort_by(|a, b| a.dist
        .partial_cmp(&b.dist)
        .unwrap()
        .reverse()
    );

    // make connections
    for _ in 0..10 {
        let d = distances.pop().unwrap();
        if d.a.circuit.is_some()  {
            let v = d.a.circuit.take();
        }
    }

    
    
    println!("{:?}", distances);

    // // println!("{:?}", text);
    // let a = Box{ x: 162, y: 817, z: 812 };
    // let b = Box{ x: 425, y: 690, z: 689 };
    // let d = distance(a, b);
    // println!("{:?}", d);
    



}
