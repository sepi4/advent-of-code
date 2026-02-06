use std::fs;

fn read_input(file_path: &str) -> String {
    let text = fs::read_to_string(file_path).expect("virhe tiedoston lukussa");
    text
}

fn part1(nums: &Vec<(i64, i64)>) {
    let mut invalid_nums: Vec<i64> = Vec::new();
    for pair in nums {
        let mut i = pair.0;
        while i <= pair.1 {
            let s = i.to_string();
            let a = &s[0..s.len()/2];
            let b = &s[s.len()/2..];

            if a == b {
                invalid_nums.push(s.parse().expect("virhe"));
            }
            
            i += 1;
        }
    }

    let sum: i64 = invalid_nums.iter().sum();
    // println!("invalid: {:?}", invalid_nums);
    println!("sum: {:?}", sum);
}

fn is_valid(num: i64) -> bool {
    let num_str = num.to_string();
    let mut l = num_str.len() / 2;
    while l > 0 {
        if num_str.len() % l != 0 {
            l -= 1;
            continue;
        }

        if repeating(&num_str, l) {
            return false;
        }
        l -= 1;
    }
    return true;
}

fn repeating(num_str: &String, l: usize) -> bool {
    let mut pre = &num_str[0..l];
    let mut start = l;
    let mut end = start + l;
    while end <= num_str.len() {
        let v = &num_str[start..end];
        if pre != v {
            return false;
        }
        pre = v;

        start += l;
        end += l;
    }
    return true;
}

fn part2(nums: &Vec<(i64, i64)>) {
    let mut invalid_nums: Vec<i64> = Vec::new();
    for pair in nums {
        let mut num = pair.0;
        while num <= pair.1 {
            if !is_valid(num) {
                invalid_nums.push(num);
            }
            num += 1;
        }
    }
    let sum: i64 = invalid_nums.iter().sum();
    println!("sum: {:?}", sum);
}

fn main() {
    // let text = read_input("../example_02.txt");
    let text = read_input("../input_02.txt");
    
    let nums: Vec<(i64, i64)> = text
        .split(',')
        .map(|x| {
            let (a, b) = x.split_once('-').expect("missing -");
            return (
                a.parse().expect("virhe 1. parsinnassa"),
                b.parse().expect("virhe 2. parsinnassa"),
            );
        }).collect();
    
    println!("part1 -----------------------------------");
    part1(&nums);
    println!("part2 -----------------------------------");
    part2(&nums);
}
