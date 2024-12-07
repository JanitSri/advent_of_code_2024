use crate::util::read_file;
use std::str::FromStr;

fn aoc01a() -> i32 {
    let contents = read_file("../inputs/01a.txt");

    let mut first: Vec<i32> = vec![];
    let mut second: Vec<i32> = vec![];
    contents.split("\n").for_each(|line| {
        let items: Vec<&str> = line.split("   ").collect();
        first.push(FromStr::from_str(items[0]).expect("could not parse first str to i32"));
        second.push(FromStr::from_str(items[1]).expect("could not parse second str to i32"));
    });

    first.sort();
    second.sort();

    let mut total = 0;

    for i in 0..first.len() {
        let diff = first[i] - second[i];
        total += i32::abs(diff);
    }

    total
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_aoc01a() {
        let result = aoc01a();
        assert_eq!(2166959, result)
    }
}
