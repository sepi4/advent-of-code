use std::fs;

fn read_input(file_path: &str) -> String {
    let text = fs::read_to_string(file_path).expect("ERROR IN READING FILE");
    text
}

fn part1(steps: &Vec<(&str, i32)>) {
    let mut index = 50;
    let mut stoped_at_zero = 0;

    for s in steps {
        if s.0 == "L" {
            index -= s.1;
        } else {
            index += s.1;
        }
        while index < 0 {
            index += 100;
        }
        while index > 99 {
            index -= 100;
        }
        if index == 0 {
            stoped_at_zero += 1;
        }
        // println!("index: {:?}", index);
    }
    println!("{}", stoped_at_zero);
}

fn part2(steps: &Vec<(&str, i32)>) {
    let mut index = 50;
    let mut clicked_zero = 0;

    for s in steps {
        let mut v = s.1 as i32;

        while v > 0 {
            if s.0 == "L" {
                index -= 1;
            } else {
                index += 1;
            }
            while index < 0 {
                index += 100;
            }
            while index > 99 {
                index -= 100;
            }
            if index == 0 {
                clicked_zero += 1;
            }
            // println!("index: {:?}", index);
            v -= 1;
        }
    }
    println!("{}", clicked_zero);
}

fn main() {
    let text = read_input("../example_01.txt");
    // let text = read_input("../input_01.txt");
    let vv: Vec<&str> = text.split("\n").filter(|x| x.len() > 0).collect();

    let mut steps: Vec<(&str, i32)> = Vec::new();
    for v in vv {
        let a = &v[..1];
        let b = &v[1..];
        let b: i32 = b.parse().expect("NOT NUMBER IN PARSE");
        steps.push((a, b));
    }
    println!("part1---");
    part1(&steps);

    println!("part2---");
    part2(&steps);
}
