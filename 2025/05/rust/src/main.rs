use std::fs;

fn read_input(file_path: &str) -> String {
    let text = fs::read_to_string(file_path).expect("virhe tiedoston luvussa");
    text
}

fn part1(fresh_ranges: &Vec<(isize, isize)>, ingredients: &Vec<isize>) -> isize {
    let mut fresh_count: isize = 0;
    for ing in ingredients {
        for (min, max) in fresh_ranges {
            if ing >= min && ing <= max {
                fresh_count += 1;
                break;
            }
        }
    }
    println!("fresh amoutn: {}", fresh_count);
    return fresh_count;
}

fn inside(a: (isize, isize), b: (isize, isize)) -> bool {
	if a.0 >= b.0 && a.0 <= b.1 {
		return true
	}
	if a.1 >= b.0 && a.1 <= b.1 {
		return true
	}
	if b.0 >= a.0 && b.0 <= a.1 {
		return true
	}
	if b.1 >= a.0 && b.1 <= a.1 {
		return true
	}
	return false;
}

fn new_min_max(a: (isize, isize), b: (isize, isize)) -> (isize, isize) {
    let vv: Vec<isize> = Vec::from([a.0, a.1, b.0, b.1]);
    let x = vv.iter().min().unwrap();
    let y = vv.iter().max().unwrap();
    return (x.clone(), y.clone());
}

fn part2(fresh_ranges: &Vec<(isize, isize)>) -> isize {
    let mut ranges = fresh_ranges.clone();
    println!("{:?}", ranges);
    loop {
        let mut added = false;
        let mut i = 0;
        while i < ranges.len() - 1 {
            let mut j = i + 1;
            while j < ranges.len() {
                let a = ranges[i];
                let b = ranges[j];
                // println!("{},{}", i, j);
                if inside(a, b) {
                    let (min, max) = new_min_max(a, b);
                    ranges.swap_remove(j);
                    ranges.swap_remove(i);
                    ranges.push((min, max));
                    added = true;
                    // println!("inside");
                }
                j += 1;
            }
            i += 1;
        }
        
        if !added  {
            break;
        }
    }
    // println!("{:?}", ranges);
    let sum: isize = ranges.iter().map(|r| {
        let diff = r.1 - r.0;
        return diff + 1;
    }).sum();
    println!("sum: {:?}", sum);
    return sum;
}

fn main() {
    // let text = read_input("../example.txt");
    let text = read_input("../input.txt");

    let (aa, bb) = text.split_once("\n\n").expect("not found");
    
    let fresh_ranges: Vec<(isize, isize)> = aa.split("\n")
        .map(|row| row.split_once("-").expect("not found '-'"))
        .map(|(a, b)| (
            a.parse().expect("error in parse range"),
            b.parse().expect("error in parse range"),
        )).collect::<Vec<(isize, isize)>>();
    
    let ingredients: Vec<isize> = bb.split("\n")
        .filter(|row| !row.is_empty())
        .map(|row| row.parse().expect("error in parse"))
        .collect();

    println!("part1 ============================");
    part1(&fresh_ranges, &ingredients);

    // println!("{:?}", new_min_max((123, 234), (188, 222)));

    println!("part2 ============================");
    part2(&fresh_ranges);

}
