use std::{
    collections::{HashMap, HashSet},
    str::FromStr,
};

use crate::util;

fn aoc05a() -> i32 {
    let contents = util::read_file("../inputs/05.txt");
    let mut split_contents = contents.split("\n\n");

    let c1 = split_contents
        .next()
        .expect("could not read page ordering rules")
        .split("\n");
    let c2 = split_contents
        .next()
        .expect("could not read page numbers")
        .split("\n");

    let mut po: HashMap<i32, Vec<i32>> = HashMap::new();
    for l in c1 {
        let mut parts = l.split("|");
        let first_str = parts
            .next()
            .expect("could get the first part of the page ordering rules");
        let first: i32 = FromStr::from_str(first_str).expect("could not parse into i32");
        let second_str = parts
            .next()
            .expect("could get the second part of the page ordering rules");
        let second: i32 = FromStr::from_str(second_str).expect("could not parse into i32");

        let v = po.entry(first).or_insert(Vec::default());
        v.push(second);
    }

    let c2: Vec<Vec<i32>> = c2
        .map(|i| {
            i.split(",")
                .map(|ii| FromStr::from_str(ii).expect("could not parse into i32"))
                .collect()
        })
        .collect();
    let mut total = 0;
    for l in c2 {
        if valid(&l, &po) {
            total += mid(&l);
        }
    }

    total
}

fn mid(l: &[i32]) -> i32 {
    return l[0 + (l.len() - 1) / 2];
}

fn valid(list: &[i32], po: &HashMap<i32, Vec<i32>>) -> bool {
    let mut seen: HashSet<i32> = HashSet::new();
    let d: Vec<i32> = Vec::default();
    for r in list {
        let l = po.get(r).unwrap_or_else(|| &d);

        for i in l {
            let s = seen.get(i);
            if let Some(_) = s {
                return false;
            }
        }

        seen.insert(*r);
    }

    true
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_aoc05a() {
        assert_eq!(5452, aoc05a());
    }
}
