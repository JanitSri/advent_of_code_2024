use std::collections::HashMap;

#[allow(unused)]

use crate::util::read_file;

fn aoc01b() -> i32 {

    let contents = read_file("../inputs/01a.txt");

    let mut first: Vec<i32> = vec!();
    let mut second: Vec<i32> = vec!();
    contents.split("\n")
        .for_each(|line| {
            let items: Vec<&str> = line.split("   ").collect();
            first.push(str::parse(items[0]).expect("could not parse first str to i32"));
            second.push(str::parse(items[1]).expect("could not parse second str to i32"));
        });

    let mut m: HashMap<i32, i32>  = HashMap::new();
    for num in second {
        *m.entry(num).or_insert(0) += 1
    }

    let total: i32 = first.iter()
    .map(|&num| m.get(&num).unwrap_or(&0) * num)
    .sum();

    total
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_aoc01b() {
        let result  = aoc01b();
        assert_eq!(23741109, result)
    }
}